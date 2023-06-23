package db

import (
	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid" json:"-"`

	TokenSalt        []byte `json:"-"`
	StrengthendToken []byte `json:"-"`

	Aliases []Alias  `gorm:"foreignKey:UserID" json:"aliases"`
	Domains []Domain `gorm:"foreignKey:UserID" json:"domains"`
}

type Domain struct {
	ID uuid.UUID `gorm:"type:uuid" json:"id"`

	UserID uuid.UUID `gorm:"type:uuid" json:"-"`
	User   User      `json:"-"`

	Domain string `json:"domain"`

	Disabled bool `gorm:"default:false" json:"disabled"`
}

type Alias struct {
	ID uuid.UUID `gorm:"type:uuid" json:"-"`

	UserID uuid.UUID `gorm:"type:uuid" json:"-"`
	User   User      `json:"-"`

	DomainID uuid.UUID `gorm:"type:uuid" json:"domainId"`
	Domain   Domain    `json:"-"`

	Alias string `json:"alias"`

	IconPresent  bool   `gorm:"default:false" json:"iconPresent"`
	IconMimeType string `json:"iconMimeType"`
	IconData     []byte `json:"iconData"`

	Disabled bool `gorm:"default:false" json:"disabled"`
}
