package creature

import "gorm.io/gorm"

// Creature represents a creature in the Zelda's Sheikah Slate compendium
type Creature struct {
	gorm.Model
	ID      int64  `json:"id"`
	Picture string `json:"picture"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	HP      int    `json:"hp"`
}
