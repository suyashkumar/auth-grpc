# auth-grpc [WIP]
A simple (but opinionated) Golang authentication gRPC microservice that implements very simple interface (below) that comes with batteries included (permissions, hashing, validation, etc). This is a gRPC wrapper around my core [suyashkumar/auth](https://github.com/suyashkumar/auth) library.

```go
type Auth interface {
	Register(user User, password string) error
	GetToken(email string, password string, reqPermissions Permissions) (token string, err error)
	Validate(token string) (*Claims, error)
}
```

You only need to set a database `connectionString` and `signingKey`, and everything else is taken care of for you including:
* table and database setup (including uniqueness constraints and useful indicies)
* hashing passwords using `bcrypt` on register
* comparing hashed passwords on login
* validation of new user fields like "Email" (TBD)
* encoding and extraction of key fields stored in the JSON Web Token (JWT)
* ensuring that a token's requested permissions does not exceed the user's maximum permission level
