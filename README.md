# cronuseo golang sdk

## The cronuseo go sdk for cronuseo authorization framework.

This package provides a simple way to check permissions based on username, resource and permission.

### Features:
- CheckUser function to check permissions with username
### How to use:
- Import the package
- Create an instance of Cronuseo interface with endpoint, organization name and API Key.
- Use the Check function to check permissions with username

### Check single permision
```
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
```

>Note: Before using this package in production, make sure you test it properly and also have a look at the possible errors and edge cases.

