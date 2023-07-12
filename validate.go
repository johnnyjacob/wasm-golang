package main

import (
	"context"
	"fmt"

	"syscall/js"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
)

func jsRegoValidateWrapper() js.Func {  
        jsRegoValidateFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
                if len(args) != 1 {
                        return "Invalid no of arguments passed"
                }
                inputPolicy := args[0].String()
                fmt.Printf("input %s\n", inputPolicy)

		// Validate Rego

		ctx := context.Background()

		r := rego.New(
			rego.Query("data.example.p"),
			rego.Module("example_error.rego",
				inputPolicy,
			))

		_, err := r.Eval(ctx)

		switch err := err.(type) {
		case ast.Errors:
			for _, e := range err {
				fmt.Println("code:", e.Code)
				fmt.Println("row:", e.Location.Row)
			}
		default:
			// Some other error occurred.
		}

                if err != nil {
                        fmt.Printf("unable to convert to json %s\n", err)
                        return err.Error()
                }
                return "Error Message"
        })

        return jsRegoValidateFunc
}

func main() {
        fmt.Println("Go Web Assembly")
	js.Global().Set("validateRegoPolicy", jsRegoValidateWrapper())
        <-make(chan bool)
}
