package main

import (
	"fmt"

	"github.com/adityashedge/go-verifalia/verifalia"
)

func main() {
	// add a valid account subscription ID and auth token(or sub-account password)
	c := verifalia.NewClient("Account SID", "Auth Token")

	emails := []string{"john.smith@example.com", "foo@example.net"}

	fmt.Println("-------------> Validate")
	resp, err := c.Validate(emails)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response:\t%+v\n", resp)
	fmt.Printf("Server response:\t%+v\n", resp.Data)

	fmt.Println("-------------> Query")
	resp, err = c.Query(resp.UniqueID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response:\t%+v\n", resp)
	fmt.Printf("Server response:\t%+v\n", resp.Data)

	fmt.Println("-------------> Delete")
	resp, err = c.Delete(resp.UniqueID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response:\t%+v\n", resp)
}
