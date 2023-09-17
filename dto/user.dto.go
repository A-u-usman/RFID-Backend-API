package dto

//RegisterDTO is used when a client hit register POST request
type RegisterDTO struct {
	Name   string `json:"name" form:"name" binding:"required"`
	Email  string `json:"email" form:"email" binding:"required,email"`
	Rfid   string `json:"rfid" form:"rfid" binding:"required"`
	Status bool   `json:"status" form:"status" binding:"required"`
}

type GetDTO struct {
	// Name   string `json:"name" form:"name" binding:"required"`
	Rfid string `json:"rfid" form:"rfid" binding:"required"`
	// Email  string `json:"email" form:"email" binding:"required,email"`
	// Status bool   `json:"status" form:"status" binding:"required"`
}

type UpdateDTO struct {
	ID           string `json:"id" form:"id" binding:"required"`
	Name         string `json:"name" form:"name" binding:"required"`
	Rfid         string `json:"rfid" form:"rfid" binding:"required"`
	Email        string `json:"email" form:"email" binding:"required,email"`
	Status       bool   `json:"status" form:"status" binding:"required"`
	AccessStatus bool   `json:"access" form:"access" binding:"required"`
}

type DeleteDTO struct {
	ID string `json:"id" form:"id" binding:"required"`
	// Name   string `json:"name" form:"name" binding:"required"`
	Rfid string `json:"rfid" form:"rfid" binding:"required"`
	// Email  string `json:"email" form:"email" binding:"required,email"`
	// Status bool   `json:"status" form:"status" binding:"required"`
}

// PhoneNumber string `gorm:"uniqueIndex;type:varchar(255)" json:"phonenumber" form:"phonenumber" binding:"required,e164"`
// DeviceName  string `gorm:"type:varchar(255)" json:"device-name" form:"device-name" binding:"required"`
// DeviceModel string `gorm:"type:varchar(255)" json:"device-model" form:"device-model" binding:"required"`
// Ip          string `gorm:"type:varchar(255)" json:"Ip" form:"Ip" binding:"required"`
// Firstname   string `json:"firstname" form:"firstname" binding:"required"`
// ID        uint64 `gorm:"primary_key:auto_increment" json:"id"`
// Name      string `gorm:"type:varchar(255)" json:"name"`
// Email     string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
// Rfid      string `gorm:"rfid" json:"rfid"`
// Status    bool   `gorm:"status" json:"status"`
// CreatedAt time.Time
// UpdatedAt time.Time
