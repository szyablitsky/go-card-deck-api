package main

import "github.com/szyablitsky/go-card-deck-api/app"

var a app.App

func main()  {
	a.Initialize()
	a.Run(":8080")
}