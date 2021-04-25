package domain

import "time"

type CryptoAddresses []CryptoAddress

type CryptoAddress struct {
	ID        int    `gorm:"primary_key;auto_increment"`
	Label     string `gorm:"type:varchar(100)"`
	Content   string `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
