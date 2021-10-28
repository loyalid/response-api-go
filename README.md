# API Response for Go

## Installation
`go get github.com/loyalid/response-api-go`

## Usage
```
import "github.com/loyalid/response-api-go"

response := new(response.API)
response.Simple("200", "OK", "Hello World!")
```