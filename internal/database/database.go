package database

import (
	"github.com/rodrigosollberger/sheikah-slate/internal/app/creature"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateCreatureDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("sheikah-slate.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&creature.Creature{})
	if err != nil {
		return nil, err
	}

	// create creature seed DB
	db.Create(&creature.Creature{
		Picture: "https://www.zeldadungeon.net/wiki/images/b/bf/Black-Bokoblin-Model.png",
		Name:    "Bokoblin",
		Type:    "Gold",
		HP:      1080,
	})
	db.Create(&creature.Creature{
		Picture: "https://www.zeldadungeon.net/wiki/images/e/e6/Black-Moblin.png",
		Name:    "Moblin",
		Type:    "Black",
		HP:      360,
	})
	db.Create(&creature.Creature{
		Picture: "https://www.zeldadungeon.net/wiki/images/c/c8/Electric-Lizalfos.png",
		Name:    "Lizalfos",
		Type:    "Electric",
		HP:      288,
	})
	db.Create(&creature.Creature{
		Picture: "https://static.wikia.nocookie.net/zelda_gamepedia_en/images/c/c4/BotW_White-Maned_Lynel_Model.png/revision/latest?cb=20180410175300",
		Name:    "Lynel",
		Type:    "White",
		HP:      4000,
	})
	db.Create(&creature.Creature{
		Picture: "https://static.wikia.nocookie.net/zelda_gamepedia_en/images/9/9b/BotW_Guardian_Artwork.png/revision/latest/scale-to-width-down/1200?cb=20170423093246",
		Name:    "Guardian",
		Type:    "Stalker",
		HP:      1500,
	})
	db.Create(&creature.Creature{
		Picture: "https://www.zeldadungeon.net/wiki/images/8/80/Yiga-blademaster.jpg",
		Name:    "Yiga",
		Type:    "Blademaster",
		HP:      400,
	})
	db.Create(&creature.Creature{
		Picture: "https://static.wikia.nocookie.net/zelda_gamepedia_en/images/6/69/BotW_Igneo_Pebblit_Model.png/revision/latest?cb=20210316071605",
		Name:    "Pebblit",
		Type:    "Igneo",
		HP:      20,
	})
	db.Create(&creature.Creature{
		Picture: "https://static.wikia.nocookie.net/zelda_gamepedia_en/images/b/ba/BotW_Black_Hinox_Model.png/revision/latest?cb=20181115055839",
		Name:    "Hinox",
		Type:    "Black",
		HP:      1000,
	})
	db.Create(&creature.Creature{
		Picture: "https://static.wikia.nocookie.net/zelda_gamepedia_en/images/a/a1/BotW_Molduga_Model.png/revision/latest?cb=20171226215850",
		Name:    "Molduga",
		Type:    "Sand",
		HP:      1500,
	})
	db.Create(&creature.Creature{
		Picture: "https://static.wikia.nocookie.net/villains/images/7/7f/CalamityGanonConceptArt.png/revision/latest?cb=20201120184101",
		Name:    "Ganon",
		Type:    "Calamity",
		HP:      8000,
	})
	db.Create(&creature.Creature{
		Picture: "https://static.wikia.nocookie.net/zelda_gamepedia_en/images/0/0b/BotW_Dark_Beast_Ganon_Model.png/revision/latest?cb=20170430145024",
		Name:    "Ganon",
		Type:    "Dark Beast",
		HP:      8,
	})

	return db, nil

}
