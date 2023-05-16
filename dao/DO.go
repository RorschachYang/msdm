package dao

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	OpenID    string `gorm:"not null;unique"`
	NickName  string
	AvatarURL string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Deck struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string `gorm:"size:512"`
	Code        string `gorm:"size:1024"`
	Cards       []Card `gorm:"many2many:deck_cards"`
	Author      User
	AuthorID    uint
	CopiedTimes uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Collection struct {
	ID        uint `gorm:"primaryKey"`
	User      User
	UserID    uint
	Cards     []Card `gorm:"many2many:collection_cards"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Card struct {
	ID            uint   `gorm:"primaryKey"`
	Name          string `gorm:"not null"`
	NameZh        string
	Description   string
	DescriptionZh string
	Source        string
	Power         int
	Cost          int
	ImageURLName  string //用于拼接成图片URL
	Defid         string //用于做卡组识别等
	Variants      []Variant
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type Variant struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	ImageURLName string    //用于拼接成图片URL
	Artist       []*Artist `gorm:"many2many:artist_variants"`
	Tags         []*Tag    `gorm:"many2many:variant_tags"`
	Rarity       string
	CardID       uint
	Released     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type Tag struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	NameZh    string
	Variant   []*Variant `gorm:"many2many:variant_tags"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Artist struct {
	ID        uint       `gorm:"primaryKey"`
	Name      string     `gorm:"not null"`
	Varinat   []*Variant `gorm:"many2many:artist_variants"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Location struct {
	ID            uint `gorm:"primaryKey"`
	Name          string
	NameZh        string
	Description   string
	DescriptionZh string
	ImageURLName  string //用于拼接成图片URL
	DefID         string
	Released      bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type Title struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	NameZh    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
