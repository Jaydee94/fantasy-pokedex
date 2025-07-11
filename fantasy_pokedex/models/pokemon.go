package models

type Pokemon struct {
	ID           uint     `json:"id" gorm:"primaryKey"`
	Name         string   `json:"name"`
	Types        []string `json:"types" gorm:"type:text[]"`
	Category     string   `json:"category"`
	HeightM      float64  `json:"height_m"`
	WeightKg     float64  `json:"weight_kg"`
	Ability      string   `json:"ability"`
	Appearance   string   `json:"appearance"`
	Attacks      []string `json:"attacks" gorm:"type:text[]"`
	PokedexEntry string   `json:"pokedex_entry"`
	ImageURL     string   `json:"image_url"`
}
