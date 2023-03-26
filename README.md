# book-api
Repo ini adalah sebuah api serderhana menggunakan GO (gin web framework) dan PostgreSQL. Untuk menjalankan silakan ketik "go run main.go" di terminal.

# ubah config database
Jangan lupa untuk menyesuaikan config database dengan mengubah data di bagian
pkg/common/const.go

# test postman
Berikut ini adalah HTTP method beserta route-nya

GET "/books/"            //get all books
GET "/books/:bookID"     //get book by id
POST "/books/            //create book
PUT "/books/:bookID"     //update book by id
DELETE "/books/:bookID"  //delete book by id
