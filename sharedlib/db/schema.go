package db

import (
	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid"`

	TokenSalt        []byte
	StrengthendToken []byte

	Aliases []Alias  `gorm:"foreignKey:UserID"`
	Domains []Domain `gorm:"foreignKey:UserID"`
}

type Domain struct {
	ID uuid.UUID `gorm:"type:uuid"`

	UserID uuid.UUID `gorm:"type:uuid"`
	User   User

	Domain string

	Disabled bool `gorm:"default:false"`
}

type Alias struct {
	ID uuid.UUID `gorm:"type:uuid"`

	UserID uuid.UUID `gorm:"type:uuid"`
	User   User

	DomainID uuid.UUID `gorm:"type:uuid"`
	Domain   Domain

	Alias string

	IconPresent  bool `gorm:"default:false"`
	IconMimeType string
	IconData     []byte

	Disabled bool `gorm:"default:false"`
}
