package example_module

import (
	"fmt"
	"net/http"
)

// ExampleModule is an example module
func ExampleModule(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Example Module!")
}
