package main

import (
	"fmt"
	"healthtracker/Database"
	"healthtracker/Router"
	"healthtracker/Templater"
	"log"
	"net/http"
	"time"
)

func main() {
	err := Templater.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = Database.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	s := &http.Server{
		Addr:           ":8080",
		Handler:        Router.Register(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("Started at: ", s.Addr)

	log.Fatal(s.ListenAndServe())
}
