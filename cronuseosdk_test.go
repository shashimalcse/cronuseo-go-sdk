package cronuseogosdk

import (
	"fmt"
	"testing"
)

func TestUsernameCheck(t *testing.T) {

	cronuseo := Cronuseo("http://localhost:8080/api/v1", "super", "ZHP/IlS5nDDXoWLyqQMcds6VyUoWjl+3+MLevUSku0A=")
	allow, err := cronuseo.CheckUser("shashimal", "read", "doc")
	if err != nil {
		fmt.Println(err)
	}
	if allow {
		fmt.Println("Allow")
	} else {
		fmt.Println("Deny")
	}
}
