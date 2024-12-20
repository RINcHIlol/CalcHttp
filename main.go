package main

import "calc_http/internal/application"

func main() {
	err := application.RunServer()
	if err != nil {
		panic(err)
	}
}
