package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/tsawler/bookings/internal/config"
	"github.com/tsawler/bookings/internal/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)

	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)
	mux.Get("/choose-room/{id}", handlers.Repo.ChooseRoom)
	mux.Get("/book-room", handlers.Repo.BookRoom)

	mux.Get("/contact", handlers.Repo.Contact)

	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Post("/make-reservation", handlers.Repo.PostReservation)
	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)

	// Dev.Mr_Hai
	mux.Get("/by-category", handlers.Repo.ByCategory)
	mux.Get("/holiday", handlers.Repo.Holiday)
	mux.Get("/by-product", handlers.Repo.ByProduct)

	mux.Get("/business-dashboard", handlers.Repo.BusinessDashboard)
	mux.Get("/business-all-products", handlers.Repo.BusinessAllProducts)
	mux.Get("/business-product-categories", handlers.Repo.BusinessProductCategories)
	mux.Get("/business-add-product", handlers.Repo.BusinessAddProduct)
	mux.Post("/business-add-product", handlers.Repo.PostProduct)
	mux.Get("/business-orders", handlers.Repo.BusinessOrders)
	mux.Get("/business-transactions", handlers.Repo.BusinessTransactions)
	mux.Get("/business-settings", handlers.Repo.BusinessSettings)
	mux.Get("/business-account-login", handlers.Repo.BusinessAccountLogin)
	mux.Get("/business-account-register", handlers.Repo.BusinessAccountRegister)
	// Dev.Mr_Hai === END OF COMMENT FOR EDITED CODE ===

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
