# NPSN Server

Repositori ini adalah untuk server yang menampung url server berdasarkan npsn yang sudah ditetapkan. Client ulangan akan mengambil data ke `/school/[npsn]`.

## Deploy

1. clone this repo
2. `./create-local-image.sh` 
3. Salin konten file [`data/schools.json.sample`](data/schools.json.sample) dan buat file baru di folder [`data`](data) dengan nama `schools.json`. Sesuaikan valuenya.
4. `docker compose up -d`
