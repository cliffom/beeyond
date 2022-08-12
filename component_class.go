package main

import "github.com/google/uuid"

const (
	CATEGORY_ENEMY  = "enemy"
	CATEGORY_PLAYER = "player"
	CATEGORY_STATIC = "static"
)

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

func NewClass(category string) Class {
	return Class{
		id:       uuid.New(),
		category: category,
	}
}

// NewPlayerClass returns a class of category player
func NewPlayerClass() Class {
	return NewClass(CATEGORY_PLAYER)
}

// NewEnemyClass returns a class of category enemy
func NewEnemyClass() Class {
	return NewClass(CATEGORY_ENEMY)
}

// NewStaticClass returns a class of category static
func NewStaticClass() Class {
	return NewClass(CATEGORY_STATIC)
}
