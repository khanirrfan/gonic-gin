package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID          uint         // Unique identifier
	Username    string       // Username
	Email       string       // Email address
	Password    string       // Password (You should consider using a secure password hashing library)
	CompanyID   uint         // Foreign key referencing the company ID
	Company     Company      // Company associated with this user
	Experiences []Experience // Experiences associated with this user
}

// Company represents a company model
type Company struct {
	ID    uint   // Unique identifier
	Name  string // Company name
	Users []User // Users belonging to this company
}

// Experience represents a user's experience model
type Experience struct {
	ID       uint   // Unique identifier
	UserID   uint   // Foreign key referencing the user ID
	User     User   // User associated with this experience
	Title    string // Experience title
	Duration int    // Duration of experience in months or years
}

// AutoMigrate will create tables and foreign key relationships
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Company{}, &Experience{})
	db.Model(&User{}).Association("Experiences")
}
