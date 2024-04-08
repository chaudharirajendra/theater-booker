package main

import (
	"fmt"
	"strings"

	"github.com/chaudharirajendra/theater-booker/report"
	"github.com/chaudharirajendra/theater-booker/theater"
	"github.com/chaudharirajendra/theater-booker/utils"
)

func main() {
	// Initialize theater with shows and available seats
	theaterObj := theater.NewTheater([]theater.Show{
		{
			Name: "1 Running in Audi 1",
			AvailableSeats: []theater.Seat{
				{Row: "A", Seat: "1"}, {Row: "A", Seat: "2"}, {Row: "A", Seat: "3"}, {Row: "A", Seat: "4"}, {Row: "A", Seat: "5"}, {Row: "A", Seat: "6"}, {Row: "A", Seat: "7"}, {Row: "A", Seat: "8"}, {Row: "A", Seat: "9"},
				{Row: "B", Seat: "1"}, {Row: "B", Seat: "2"}, {Row: "B", Seat: "3"}, {Row: "B", Seat: "4"}, {Row: "B", Seat: "5"}, {Row: "B", Seat: "6"},
				{Row: "C", Seat: "1"}, {Row: "C", Seat: "2"}, {Row: "C", Seat: "3"}, {Row: "C", Seat: "4"}, {Row: "C", Seat: "5"}, {Row: "C", Seat: "6"}, {Row: "C", Seat: "7"},
			},
		},
		{
			Name: "2 Running in Audi 2",
			AvailableSeats: []theater.Seat{
				{Row: "A", Seat: "1"}, {Row: "A", Seat: "2"}, {Row: "A", Seat: "3"}, {Row: "A", Seat: "4"}, {Row: "A", Seat: "5"}, {Row: "A", Seat: "6"}, {Row: "A", Seat: "7"},
				{Row: "B", Seat: "1"}, {Row: "B", Seat: "2"}, {Row: "B", Seat: "3"}, {Row: "B", Seat: "4"}, {Row: "B", Seat: "5"}, {Row: "B", Seat: "6"},
				{Row: "C", Seat: "1"}, {Row: "C", Seat: "2"}, {Row: "C", Seat: "3"}, {Row: "C", Seat: "4"}, {Row: "C", Seat: "5"}, {Row: "C", Seat: "6"}, {Row: "C", Seat: "7"},
			},
		},
		{
			Name: "3 Running in Audi 3",
			AvailableSeats: []theater.Seat{
				{Row: "A", Seat: "1"}, {Row: "A", Seat: "2"}, {Row: "A", Seat: "3"}, {Row: "A", Seat: "4"}, {Row: "A", Seat: "5"}, {Row: "A", Seat: "6"}, {Row: "A", Seat: "7"},
				{Row: "B", Seat: "1"}, {Row: "B", Seat: "2"}, {Row: "B", Seat: "3"}, {Row: "B", Seat: "4"}, {Row: "B", Seat: "5"}, {Row: "B", Seat: "6"}, {Row: "B", Seat: "7"}, {Row: "B", Seat: "8"},
				{Row: "C", Seat: "1"}, {Row: "C", Seat: "2"}, {Row: "C", Seat: "3"}, {Row: "C", Seat: "4"}, {Row: "C", Seat: "5"}, {Row: "C", Seat: "6"}, {Row: "C", Seat: "7"}, {Row: "C", Seat: "8"}, {Row: "C", Seat: "9"},
			},
		},
	})

	var totalRevenue, totalServiceTax, totalSwachhCess, totalKrishiCess float64

	for {
		var showNo int
		fmt.Print("Enter Show no: ")
		fmt.Scanln(&showNo)

		theaterObj.PrintAvailableSeats(showNo)

		var seatInput string
		fmt.Print("Enter seats: ")
		fmt.Scanln(&seatInput)

		// Parse seatInput into individual seats
		seatStrings := strings.Split(seatInput, ",")
		var seats []theater.Seat
		for _, seatStr := range seatStrings {
			seatStr = strings.TrimSpace(seatStr) // Trim spaces
			if len(seatStr) < 2 {
				fmt.Println("Invalid seat format:", seatStr)
				return
			}
			row := strings.ToUpper(seatStr[:1])
			seatNum := seatStr[1:]
			seats = append(seats, theater.Seat{Row: row, Seat: seatNum})
		}

		if theaterObj.BookSeats(showNo, seats) {
			// Update total revenue and taxes
			subtotal := 0
			for _, seat := range seats {
				subtotal += theaterObj.CalculateSeatPrice(seat)
			}
			totalRevenue += float64(subtotal)
			totalServiceTax += float64(subtotal) * utils.ServiceTax
			totalSwachhCess += float64(subtotal) * utils.SwachhCess
			totalKrishiCess += float64(subtotal) * utils.KrishiCess
		}

		var continueBooking string
		fmt.Print("Would you like to continue (Yes/No): ")
		fmt.Scanln(&continueBooking)
		if strings.ToLower(continueBooking) != "yes" {
			break
		}
	}

	// Print total sales report using the reporting package
	report.PrintSalesReport(totalRevenue, totalServiceTax, totalSwachhCess, totalKrishiCess)
}
