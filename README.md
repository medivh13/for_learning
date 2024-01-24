# Project ini saya buat untuk sharing kepada rekan - rekan semua
- pada setiap materi yang saya share, mungkin akan saya tempatkan pada branch yang berbeda

# cara menjalankan setelah clone
- cd for_learning
- go mod tidy (untuk mengambil package yang dibutuhkan)
- copy .env.example ke file baru .env lalu bisa disesuaikan dengan kebutuhan untuk configurasinya
- go run main.go

# end point
```
curl --location --request GET 'http://localhost:8080/api/books?subject=love'
```