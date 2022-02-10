package models

import "time"

// User is the user model
type User struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Room is the room model
type Room struct {
	ID        int
	RoomName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Restriction is the restriction model
type Restriction struct {
	ID              int
	RestrictionName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Reservation is the reservation model
type Reservation struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	StartDate time.Time
	EndDate   time.Time
	RoomID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Room      Room
}

// RoomRestriction is the room restriction model
type RoomRestriction struct {
	ID            int
	StartDate     time.Time
	EndDate       time.Time
	RoomID        int
	ReservationID int
	RestrictionID int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Room          Room
	Reservation   Reservation
	Restriction   Restriction
}

// Dev.Mr_Hai
// Product is the product model
type Product struct {
	ID             int
	Name           string
	Desc           string
	Material       string
	Color          string
	Gender         string
	Size           string
	Brand          string
	Origin         string
	TaxRate        float64
	PriceRegular   float64
	PriceDiscount  float64
	Currency       string
	SKU            string
	ShippingWidth  int
	ShippingHeight int
	ShippingWeight int
	ShippingFee    float64
	Category       string
	CategorySub    string
	Tag            string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// Dev.Mr_Hai ===END OF COMMENT FOR EDITED CODE===
