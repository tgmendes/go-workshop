package main

import (
	"encoding/json"
	"time"
)

// UnmarshalUser will read the data passed as input and return an instance of a User.
// Note that we are returning a pointer to a user.
func UnmarshalUser(data []byte) (*User, error) {
	var u User
	err := json.Unmarshal(data, &u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// User is a struct that holds information about a user of the system.
// It's used here as a simple demonstration of using a struct and different possible types.
type User struct {
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	// The type below is a slice: it has an undefined number of elements.
	// Go also has the concept of arrays, which have a fixed number of elements.
	Entitlements []string `json:"entitlements"`

	Metadata map[string]any `json:"metadata"`

	// This is an example of an unexported field (notice the lowercase)
	age int
}

// Age returns the age for a given user from their date of birth.
// This is a method because it is "attached" to a struct.
// Notice the method is called `Age` instead of `GetAge` - Go favours simplicity in method naming, and
// using the context in which things are being used to clearly know the purpose of methods, functions and variables.
// Method receiver names should also be short (u instead of User) (Check go code review comments)
func (u *User) Age() int {
	// A naive time calculation for academic purposes
	return time.Now().Year() - u.DateOfBirth.Year()
}

// HasAccess checks if a user has access to a given entitlement.
func (u *User) HasAccess(name string) bool {
	// This is how to search for an element in a Go slice (there is no 'in' or 'contains' out of the box)
	for _, r := range u.Entitlements {
		// If we find an element in the slice we return
		if r == name {
			return true
		}
	}
	// If we didn't find anything, we just return false
	return false
}

// SetAge is purely for demonstration purposes of pointers.
func (u User) SetAge() {
	// A naive time calculation for academic purposes
	u.age = time.Now().Year() - u.DateOfBirth.Year()
}

// GetSetAge is also purely for demonstration purposes. The name of the method is not a good example.
func (u User) GetSetAge() int {
	return u.age
}
