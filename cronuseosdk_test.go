package cronuseogosdk

import (
	"fmt"
	"testing"
)

func TestUsernameCheckPermission(t *testing.T) {

	cronuseo := Cronuseo("http://localhost:8080/api/v1", "super", "g+6bFqAg+y1EswfWlxGeDSRh3+WlxsrFEBH30fK8HHg=")
	allow, err := cronuseo.CheckPermission("shashimal", "read", "doc")
	if err != nil {
		fmt.Println(err)
	}
	if allow {
		fmt.Println("Allow")
	} else {
		fmt.Println("Deny")
	}
}

func TestUsernameCheckPermissions(t *testing.T) {

	cronuseo := Cronuseo("http://localhost:8080/api/v1", "super", "g+6bFqAg+y1EswfWlxGeDSRh3+WlxsrFEBH30fK8HHg=")
	grantedScopes, err := cronuseo.CheckPermissions("shashimal", []string{"read", "create"}, "doc")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(grantedScopes)
}
