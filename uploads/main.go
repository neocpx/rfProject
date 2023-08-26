package main

import "os"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5432"
	}
	server := newServer(":" + port)
	server.Run()
}
