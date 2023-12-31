package models

import "time"

type UserActivityLog struct {
	ID uint64 `gorm:"primary_key:auto_increment" json:"id"`
	// UserID       uint64 `gorm:"type:varchar(255)" json:"userId"`
	Name         string `gorm:"type:varchar(255)" json:"name"`
	Email        string `gorm:"unique:false;type:varchar(255)" json:"email"`
	Rfid         string `gorm:"unique:false;type:varchar(255)" json:"rfid"`
	Status       bool   `gorm:"bool" json:"status"`
	AccessStatus bool   `gorm:"bool" json:"access"`
	DoorStatus   string `gorm:"unique:false;type:varchar(255)" json:"door"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	// UserEmail string `gorm:"primaryKey;autoIncrment:false;uniqueIndex;type:varchar(255);not null" json:"email"`
	// Otp       string `gorm:"type:varchar(255)" json:"otp"`
	// CreatedAt time.Time
	// ExpiresAt time.Time
}
