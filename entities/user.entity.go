package models

import "time"

type User struct {
	ID        uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name      string `gorm:"type:varchar(255)" json:"name"`
	Email     string `gorm:"type:varchar(255)" json:"email"`
	Rfid      string `gorm:"uniqueIndex;type:varchar(255)" json:"rfid"`
	Status    bool   `gorm:"bool" json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Rfid string `gorm:"->; <-;not null" json:"-"` uniqueIndex
// Gender        string      `gorm:"type:varchar(255)" json:"gender"`
// Location      string      `gorm:"type:varchar(255)" json:"location"`
// PhoneNumber   string      `gorm:"uniqueIndex;type:varchar(255)" json:"phonenumber" binding:"e164"`
// AccountStatus bool        `gorm:"default:true" json:"accountstatus"`
// DeviceName    string      `gorm:"type:varchar(255)" json:"device-name"`
// DeviceModel   string      `gorm:"type:varchar(255)" json:"device-model"`
// Ip            string      `gorm:"type:varchar(255)" json:"Ip"`
// Category      *[]Category `json:"categories,omitempty"`
// Token     string

//ActivatedAt sql.NullTime
