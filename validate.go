package main

import (
	"context"
	"fmt"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
)

func main() {

	ctx := context.Background()

	r := rego.New(
		rego.Query("data.example.p"),
		rego.Module("example_error.rego",
			`package example

p = true { not q[x] }
q = {1, 2, 3} { true }`,
		))

	_, err := r.Eval(ctx)

	switch err := err.(type) {
	case ast.Errors:
		for _, e := range err {
			fmt.Println("code:", e.Code)
			fmt.Println("row:", e.Location.Row)
			fmt.Println("filename:", e.Location.File)
		}
	default:
		// Some other error occurred.
	}

}
