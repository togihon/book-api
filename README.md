# book-api
Repo ini adalah sebuah api serderhana menggunakan GO (gin web framework) dan PostgreSQL. Untuk menjalankan silakan ketik "go run main.go" di terminal.

# buat database postgreSQL
buatlah sebuah database postgresql (bisa menggunakan pgadmin) dan buatlah sebuah table dengan script  
  
CREATE TABLE IF NOT EXISTS public.books  
(  
    id_book integer NOT NULL DEFAULT 'nextval('books_id_book_seq'::regclass)',  
    title character varying(75) COLLATE pg_catalog."default" NOT NULL,  
    author character varying(75) COLLATE pg_catalog."default" NOT NULL,  
    description text COLLATE pg_catalog."default",  
    CONSTRAINT books_pkey PRIMARY KEY (id_book)  
)  

# ubah config database
Jangan lupa untuk menyesuaikan config database dengan mengubah data di bagian  
pkg/common/const.go

# body request postman untuk create dan update  
  ```
  {  
  title: "insert title"  
  author: "insert author"  
  description: "insert description"  
  }  
  ```
# test postman
Berikut ini adalah HTTP method beserta route-nya  
  
GET "/books/" &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;//get all books  
GET "/books/:bookID" &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;  //get book by id  
POST "/books/" &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; //create book  
PUT "/books/:bookID" &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;   //update book by id  
DELETE "/books/:bookID" &nbsp;&nbsp;  //delete book by id  
