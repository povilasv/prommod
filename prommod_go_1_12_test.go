// +build go1.12

package prommod

import (
	"fmt"
)

func ExamplePrint() {
	version = map[string]string{
		"pkg1": "v1.1.1",
		"pkg2": "v1.1.1",
	}

	fmt.Println(Print("app_name"))
	// Output: app_name
	//        pkg1: v1.1.1
	//        pkg2: v1.1.1
}
