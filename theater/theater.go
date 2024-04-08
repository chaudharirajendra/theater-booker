// theater package
package theater

import (
	"fmt"
	"math"
	"strings"

	"github.com/chaudharirajendra/theater-booker/utils"
)

// SeatPricer calculates the price for a given seat
type SeatPricer interface {
	CalculateSeatPrice(seat Seat) int
}

// SeatAvailabilityChecker checks if a seat is available
type SeatAvailabilityChecker interface {
	IsSeatAvailable(show Show, seat Seat) bool
}

// BookingCalculator calculates the subtotal and taxes for a booking
type BookingCalculator interface {
	CalculateBookingCost(subtotal int) (total, serviceTax, swachhCess, krishiCess float64)
}

// TheaterBooker books seats for a show
type TheaterBooker interface {
	BookSeats(showNo int, seats []Seat) bool
}

// Show represents a movie show with available seats
type Show struct {
	Name           string
	AvailableSeats []Seat
}

// Seat represents a seat in the theater
type Seat struct {
	Row  string
	Seat string
}

// Theater represents a theater with multiple shows
type Theater struct {
	Shows []Show
}

// NewTheater creates a new theater with provided shows
func NewTheater(shows []Show) *Theater {
	return &Theater{Shows: shows}
}

// PrintAvailableSeats prints all available seats for a given show
func (t *Theater) PrintAvailableSeats(showNo int) {
	show := t.Shows[showNo-1]
	fmt.Printf("Available Seats for Show %d:\n", showNo)

	for _, seat := range show.AvailableSeats {

		if seat.Row == "A" {
			fmt.Printf("%s %s", seat.Row, seat.Seat)
		}

	}
}

// BookSeats books the provided seats for a given show
func (t *Theater) BookSeats(showNo int, seats []Seat) bool {
	show := t.Shows[showNo-1]
	for _, seat := range seats {
		if !t.IsSeatAvailable(show, seat) {
			fmt.Printf("%s%s Not available, Please select different seats\n", seat.Row, seat.Seat)
			return false
		}
	}

	// Calculate subtotal
	subtotal := 0
	for _, seat := range seats {
		subtotal += t.CalculateSeatPrice(seat)
	}

	// Calculate taxes
	var serviceTax, swachhCess, krishiCess float64
	_, serviceTax, swachhCess, krishiCess = t.CalculateBookingCost(subtotal)

	// Calculate total
	total := subtotal + int(math.Round(serviceTax+swachhCess+krishiCess))

	// Print booking details
	fmt.Printf("Successfully Booked - Show %d\n", showNo)
	fmt.Printf("Subtotal: Rs. %d\n", subtotal)
	fmt.Printf("Service Tax @14%%: Rs. %.2f\n", serviceTax)
	fmt.Printf("Swachh Bharat Cess @0.5%%: Rs. %.2f\n", swachhCess)
	fmt.Printf("Krishi Kalyan Cess @0.5%%: Rs. %.2f\n", krishiCess)
	fmt.Printf("Total: Rs. %d\n", total)

	// Update available seats
	for i := range show.AvailableSeats {
		for _, seat := range seats {
			if show.AvailableSeats[i] == seat {
				show.AvailableSeats[i] = Seat{} // Mark seat as booked
			}
		}
	}

	return true
}

// CalculateSeatPrice calculates the price for a given seat
func (t *Theater) CalculateSeatPrice(seat Seat) int {
	switch strings.ToUpper(seat.Row) {
	case "A":
		return utils.PlatinumPrice
	case "B":
		return utils.GoldPrice
	case "C":
		return utils.SilverPrice
	default:
		return 0
	}
}

// IsSeatAvailable checks if a seat is available
func (t *Theater) IsSeatAvailable(show Show, seat Seat) bool {
	for _, availableSeat := range show.AvailableSeats {
		if availableSeat == seat {
			return true
		}
	}
	return false
}

// CalculateBookingCost calculates the subtotal and taxes for a booking
func (t *Theater) CalculateBookingCost(subtotal int) (total, serviceTax, swachhCess, krishiCess float64) {

	tax := Tax{Name: "Service Tax", Percentage: utils.ServiceTax}

	serviceTax = tax.Calculate(subtotal)

	tax = Tax{Name: "SwachhCess", Percentage: utils.SwachhCess}

	swachhCess = tax.Calculate(subtotal)

	tax = Tax{Name: "KrishiCess", Percentage: utils.KrishiCess}

	krishiCess = tax.Calculate(subtotal)

	total = float64(subtotal) + serviceTax + swachhCess + krishiCess
	return total, serviceTax, swachhCess, krishiCess
}
