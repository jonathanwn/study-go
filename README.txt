GO FUNDAMENTALS DAN FIRST API
GDGoC Universitas Brawijaya - Web Development Study Jam

Repository ini berisi contoh kode dari sesi Go Fundamentals dan First API.
Setiap materi berada di foldernya masing-masing dan bisa dijalankan sendiri.


1. PERSIAPAN

   a. Install Go (versi 1.22 atau lebih baru)
      Unduh di https://go.dev/dl/

      Cek instalasi dengan perintah:
      go version

   b. Install Git
      Unduh di https://git-scm.com/downloads

   c. Install Postman (untuk mengetes API)
      Unduh di https://www.postman.com/downloads/

   d. Disarankan memakai Visual Studio Code dengan ekstensi "Go"


2. CARA MENDOWNLOAD (CLONE)

   Buka terminal, lalu jalankan:

   git clone [link repository]
   cd [nama folder repository]


3. CARA MENJALANKAN

   Setiap folder adalah program yang berdiri sendiri.
   Masuk ke folder yang ingin dijalankan, lalu jalankan:

   cd variable
   go run .

   Ganti "variable" dengan nama folder lain sesuai kebutuhan.

   Contoh:
   cd looping
   go run .

   PENTING: jalankan "go run ." di dalam folder materinya,
   bukan dari folder paling luar.


4. DAFTAR FOLDER

     Go Fundamentals
   variable       Variable dan tipe data
   condition      Kondisi (if else dan switch)
   looping        Perulangan (for)
   function       Function
   arraySlice     Array dan slice

   Materi tambahan
   map            Map
   struct         Struct
   pointer        Pointer
   errorHandling  Penanganan error


6. KENDALA UMUM

   - "go: command not found" atau "'go' is not recognized"
     Go belum terinstall atau belum masuk PATH.
     Tutup lalu buka ulang terminal setelah instalasi.

   - "package ... is not a main package"
     Baris pertama file harus: package main

   - "function main is undeclared in the main package"
     Nama fungsi harus "main" dengan huruf kecil semua.

   - Server langsung berhenti atau muncul error "address already in use"
     Port 8080 sedang dipakai program lain.
     Hentikan program tersebut atau ganti port di kode,
     misalnya ":8081".

   - Data hilang setelah server dijalankan ulang
     Normal. Data disimpan di slice (memori), bukan di database.


