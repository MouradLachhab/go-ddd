package domain

import (
	"errors"
)

const (
	LegalAge = 16
)

type Owner struct {
	firstName string
	lastName  string
	age       int
}

func CreateOwner(firstName string, lastName string, age int) (*Owner, error) {
	if age < LegalAge {
		return nil, errors.New("Owner is not of legal age")
	}

	return &Owner{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}, nil
}
