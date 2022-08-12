package main

import "github.com/google/uuid"

type Class struct {
	id       uuid.UUID
	category string
}

// GetID returns the id of a class
func (c *Class) GetID() (id uuid.UUID) {
	return c.id
}

// GetCategory returns the category of a class
func (c *Class) GetCategory() (category string) {
	return c.category
}

// NewEnemyClass returns a class of category enemy
func NewEnemyClass() Class {
	return Class{
		id:       uuid.New(),
		category: "enemy",
	}
}
