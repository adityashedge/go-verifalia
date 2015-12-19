# go-verifalia
Go library for accessing the Verifalia API http://verifalia.com/developers

Verifalia provides a simple HTTPS-based API for validating email addresses and checking whether or not they are deliverable.
This Golang package allows to communicate with the Verifalia API, scrubbing lists of email addresses in a couple of lines of code.
To learn more about Verifalia, please visit http://verifalia.com

## Usage ##

```go
import "github.com/adityashedge/go-verifalia/verifalia"
```

### Validate emails ###
Construct a new Verifalia client, then use the client to access different API of Verifalia.  
To create a job for validating emails, create a client and call the 'Validate' method on the client,
passing the emails as an argument.
```go
client := verifalia.NewClient("Account SID", "Auth Token")
emails := []string{"john.smith@example.com", "foo@example.net"}
resp, err := client.Validate(emails)
```

### Get status of queued job ###
Query the Email Validations API for a specific job result using the unique job ID.
Pass the unique ID as an argument to the 'Query' method on the client.
Response is same as 'Validate' method on the client.
```go
client := verifalia.NewClient("Account SID", "Auth Token")
resp, err := client.Query("Unique ID")
```

### Delete an existing job ###
Delete a specific validation job using the unique job ID.
Pass the unique ID as an argument to the 'Delete' method on the client.
Returns status 200(OK) when job deleted else returns status 406(Not Acceptable).
Returns status 404(Not Found) when job does not exist.
```go
client := verifalia.NewClient("Account SID", "Auth Token")
resp, err := client.Delete("Unique ID")
```

