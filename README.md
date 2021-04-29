# rbac-client-go
A Go client for the cloud.redhat.com RBAC API

[![Go Report Card](https://goreportcard.com/badge/github.com/RedHatInsights/rbac-client-go)](https://goreportcard.com/report/github.com/RedHatInsights/rbac-client-go) [![Go Reference](https://pkg.go.dev/badge/github.com/RedHatInsights/rbac-client-go.svg)](https://pkg.go.dev/github.com/RedHatInsights/rbac-client-go)

## Features
This client is an evolving work and generally implements functionality as needed.
* [x] Get permitted access for a principal
* [ ] List principals
* [ ] List permissions
* [ ] View and manage groups
* [ ] View and manage policies
* [ ] View and manage roles

## Usage
A client is created given a base URL and application name. A new `http.Client` is generated, but can be overriden
if needed for a custom Transport, etc.

```go
c := rbac.NewClient("https://foo.bar/api/rbac/v1", "app")
```

Most operations are a method of the client.

```go
acl, err := c.GetAccess(identity, "")
```

An `AccessList` contains a method for testing permissions.
```go
if acl.IsAllowed("chipotle", "burrito_bowl", "order") {
    fmt.Printf("yay!")
}
```

## License
Apache 2.0

See LICENCE to see the full text.
