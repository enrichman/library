package book

import (
	"errors"
	"fmt"
)

type User struct {
	Age Age
}

type Age struct {
	age int
}

func ParseAge(age int) (Age, error) {
	if age < 0 {
		return Age{}, errors.New("negative age")
	}
	return Age{age}, nil
}

func Book(user User) error {
	fmt.Println("booking trip", user)
	return nil
}
