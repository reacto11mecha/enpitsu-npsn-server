# NPSN Server

Repositori ini adalah untuk server yang menampung url server berdasarkan npsn yang sudah ditetapkan. Client ulangan akan mengambil data ke `/api/school/[npsn]`.

## Deploy

1. Buat repositori baru menggunakan template, [klik disini](https://github.com/new?template_name=enpitsu-npsn-server&template_owner=reacto11mecha). Lengkapi nama dan klik buat.
2. Jika proses sudah selesai, edit file [`.gitignore`](.gitignore). Komentari baris terakhir seperti yang di contohkan seperti di bawah.
   ```diff
   - server/assets/schools.json
   + # server/assets/schools.json
   ```
3. Salin konten file [`server/assets/schools.json.sample`](server/assets/schools.json.sample) dan buat file baru di folder [`server/assets`](server/assets) dengan nama `schools.json`. Sesuaikan valuenya.
5. Host server ini di https://vercel.com dengan memilih repositori yang telah di buat, langsung deploy tanpa konfigurasi apapun.
