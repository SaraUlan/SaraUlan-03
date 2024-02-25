package domain

import "fmt"

type Group struct {
	ID       int
	Name     string
}

func NewGroup(id int, name string) (*Group, error) {
	if len(name) > 250 {
		return nil, fmt.Errorf("group name exceeds maximum length of 250 characters")
	}

	group := &Group{
		ID:   id,
		Name: name,
	}

	return group, nil
}
