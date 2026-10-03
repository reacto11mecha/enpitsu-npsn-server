package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type School struct {
	NPSN   int    `json:"npsn"`
	Origin string `json:"origin"`
	Name   string `json:"name"`
}

type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

var schools []School

func init() {
	file, err := os.ReadFile("data/schools.json")
	if err != nil {
		log.Fatal("Could not read schools.json: ", err)
	}
	if err := json.Unmarshal(file, &schools); err != nil {
		log.Fatal("Error parsing schools.json: ", err)
	}
	fmt.Printf("Loaded %d schools\n", len(schools))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/school/", handleSchoolRequest)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Printf("Server starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func handleSchoolRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,PUT,PATCH,POST,DELETE")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	npsnStr := strings.TrimPrefix(r.URL.Path, "/school/")

	if npsnStr == "" {
		sendError(w, "NPSN tidak boleh kosong!", http.StatusBadRequest)
		return
	}

	if len(npsnStr) != 8 {
		sendError(w, "Format NPSN tidak sesuai!", http.StatusBadRequest)
		return
	}

	npsn, err := strconv.Atoi(npsnStr)
	if err != nil {
		sendError(w, "Format NPSN tidak sesuai!", http.StatusBadRequest)
		return
	}

	var foundSchool *School
	for _, s := range schools {
		if s.NPSN == npsn {
			foundSchool = &s
			break
		}
	}

	if foundSchool == nil {
		sendError(w, "Sekolah tidak ditemukan!", http.StatusNotFound)
		return
	}

	// 1 month
	w.Header().Set("Cache-Control", "public, max-age=2592000")

	json.NewEncoder(w).Encode(Response{Data: foundSchool})
}

func sendError(w http.ResponseWriter, message string, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Response{Message: message})
}
