// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateUserInput struct {
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Bio      *string `json:"bio,omitempty"`
	Location *string `json:"location,omitempty"`
}

type Mutation struct {
}

type Profile struct {
	ID       string  `json:"id"`
	Bio      *string `json:"bio,omitempty"`
	Location *string `json:"location,omitempty"`
}

type Query struct {
}

type User struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Email   string   `json:"email"`
	Profile *Profile `json:"profile,omitempty"`
}