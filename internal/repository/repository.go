package repository

import (
	"time"

	"github.com/tsawler/bookings/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
	SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error)
	SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)
	GetRoomByID(id int) (models.Room, error)

	// Mr.Dev Hai NguyenVan
	InsertProduct(res models.Product) (int, error)
	InsertMicroservicesApi(res models.MicroservicesDB) (int, error)
}
