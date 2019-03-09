// +build !go1.12

package prommod

import (
	"fmt"
)

func ExamplePrint() {
	fmt.Println(Print("test_app_name"))
	// Output: test_app_name
}
