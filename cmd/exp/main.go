package main

import (
	"errors"
	"fmt"
	"html/template"
)

type User struct {
	Name string
	Bio  template.HTML
	Age  int
}

func main() {
	err := CreateOrg()
	fmt.Println(err)
}

func Connect() error {
	return errors.New("Connect failed")
}

func CreateUser() error {
	err := Connect()
	if err != nil {
		return fmt.Errorf("Create User: %w", err)
	}
	return nil
}

func CreateOrg() error {
	err := CreateUser()
	if err != nil {
		return fmt.Errorf("Create Org: %w", err)
	}
	return nil
}
