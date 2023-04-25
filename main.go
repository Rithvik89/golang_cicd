package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Printf("server succesfully started on port 3000")
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		panic(err)
	}

}
