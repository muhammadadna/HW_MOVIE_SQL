package structs

import (
	"github.com/jinzhu/gorm"
)

type Movie struct {
	gorm.Model
	Title       string `gorm: "type:char(255); not null"`
	Slug        string `gorm: "type:char(255); unique; not null"`
	Description string `gorm: "type:text; not null;"`
	Duration    int    `gorm: "type:int(5); not null;"`
	Image       string `gorm: "type:char(255); not null"`
}
