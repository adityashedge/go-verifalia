package main

import (
	"fmt"

	"github.com/adityashedge/go-verifalia/verifalia"
)

func main() {
	// add a valid account subscription ID and auth token(or sub-account password)
	c := verifalia.NewClient("Account SID", "Auth Token")
	emails := []string{"john.smith@example.com", "foo@example.net"}
	resp, err := c.Validate(emails)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Response:\t%+v\n", resp)
		fmt.Printf("Server response:\t%+v\n", resp.Data)
	}
}
