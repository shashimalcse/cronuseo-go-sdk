package cronuseogosdk

import (
	"fmt"
	"testing"
)

func TestUsernameCheckPermission(t *testing.T) {

	cronuseo := Cronuseo("http://localhost:8080/api/v1", "super", "JLE+1Z3c/jIQL+i+ORhI+jLbM5pXvdxNrKvIcrKVFss=")
	allow, err := cronuseo.CheckPermission("shashimal", "write", "doc")
	if err != nil {
		fmt.Println(err)
	}
	if allow {
		fmt.Println("Allow")
	} else {
		fmt.Println("Deny")
	}
}
