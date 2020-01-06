package practice_gorm

import (
	"github.com/jinzhu/gorm"
)

type Animal struct {
	ID   int64
	Name string `gorm:"default:'galeone'"`
	Age  int64
}

type animalRepository struct {
	db *gorm.DB
}

func (a animalRepository) Create() {
	var animal = Animal{Age: 99, Name: ""}
	a.db.Create(animal)
}
