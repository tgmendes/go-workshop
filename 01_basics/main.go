package main

import "fmt"

// EntitlementChecker is an interface to check if a given entitlement is
// available. It's meant as a simple example of an interface.
type EntitlementChecker interface {
	HasAccess(name string) bool
}

func main() {
	data := `{
	"name": "John",
	"date_of_birth": "1990-03-01T00:00:00Z",
	"entitlements": ["A", "B"],
	"metadata": {"isAdmin": false}
}
`
	user, err := UnmarshalUser([]byte(data))
	if err != nil {
		panic(err)
	}

	fmt.Printf("User's age is: %d\n", user.Age())

	// Because "User" has the 'HasAccess' method, we can pass it as an argument
	// to the function.
	hasAllEntitlements := HasAllEntitlements(user)
	fmt.Printf("Has access to all entitlements? %t\n", hasAllEntitlements)

	// Should be 0 - the Go default for int
	fmt.Printf("The set age is: %d\n", user.GetSetAge())

	user.SetAge()
	// Should still be 0
	fmt.Printf("The set age is: %d\n", user.GetSetAge())
}

func HasAllEntitlements(ec EntitlementChecker) bool {
	all := []string{"A", "B", "C"}

	for _, ent := range all {
		// If the user doesn't have access to at least one entitlement, then
		// we fail
		if !ec.HasAccess(ent) {
			return false
		}
	}

	return true
}
