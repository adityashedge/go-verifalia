# go-verifalia
Go library for accessing the Verifalia API http://verifalia.com/developers

Verifalia provides a simple HTTPS-based API for validating email addresses and checking whether or not they are deliverable.
This Golang package allows to communicate with the Verifalia API, scrubbing lists of email addresses in a couple of lines of code.
To learn more about Verifalia, please visit http://verifalia.com

## Usage ##

```go
import "github.com/adityashedge/go-verifalia/verifalia"
```

### Validate Emails ###
Construct a new Verifalia client, then use the client to access different API of Verifalia.  
To create a job for validating emails, create a client and call the 'Validate' method on the client,
passing the emails as an argument.
```go
client := verifalia.NewClient("Account SID", "Auth Token")
emails := []string{"john.smith@example.com", "foo@example.net"}
resp, err := client.Validate(emails)
```
