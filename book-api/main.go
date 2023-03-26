package main

import "book-api/app/handler"

func main() {
	var PORT = ":8080"

	handler.StartServer().Run(PORT)
}
