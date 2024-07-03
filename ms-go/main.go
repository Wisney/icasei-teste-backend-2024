package main

import (
	"ms-go/app/consumers"
	_ "ms-go/db"
	"ms-go/router"
)

func main() {
	go func(){
		consumers.ListenProduct()
	}()
	
	router.Run()
}
