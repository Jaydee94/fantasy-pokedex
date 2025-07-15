package models

import (
	"github.com/lib/pq"
)

type Pokemon struct {
	Name         string         `gorm:"primaryKey"`
	Types        pq.StringArray `gorm:"type:text[]"`
	PokedexEntry string
	ImageData    []byte `gorm:"type:bytea"`
	Appearance   string
	Attacks      pq.StringArray `gorm:"type:text[]"`
	Ability      string
	HeightM      float64
	WeightKg     float64
	Category     string
}
