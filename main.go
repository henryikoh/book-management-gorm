package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/book-crud/pkg/myhttp"
	"github.com/book-crud/pkg/repositery"
)

func main() {

	// init DAO
	dao := repositery.InitDAO()

	// init server
	srv := myhttp.Newserver(*dao)

	server := &http.Server{
		Addr:         ":8090",
		Handler:      srv,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Starting server at port 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
