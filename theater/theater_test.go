package theater

import (
	"testing"

	"github.com/chaudharirajendra/theater-booker/utils"
	"github.com/stretchr/testify/assert"
)

func TestTheater(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name          string
		theater       *Theater
		showNo        int
		seats         []Seat
		expectedError bool
	}{
		{
			name: "Valid booking",
			theater: NewTheater([]Show{
				{
					Name: "Test Show",
					AvailableSeats: []Seat{
						{"A", "1"}, {"A", "2"}, {"A", "3"},
					},
				},
			}),
			showNo:        1,
			seats:         []Seat{{"A", "1"}, {"A", "2"}},
			expectedError: false, // No error expected
		},
		{
			name: "InValid booking",
			theater: NewTheater([]Show{
				{
					Name: "Test Show",
					AvailableSeats: []Seat{
						{"A", "1"}, {"A", "2"}, {"A", "3"},
					},
				},
			}),
			showNo:        1,
			seats:         []Seat{{"A", "1"}, {"A", "4"}},
			expectedError: true, //expected error
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			success := tc.theater.BookSeats(tc.showNo, tc.seats)
			if tc.expectedError {
				assert.False(t, success, "Expected booking to fail but it succeeded")
			} else {
				assert.True(t, success, "Expected booking to succeed but it failed")
			}
		})
	}
}

func TestCalculateSeatPrice(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name     string
		row      string
		expected int
	}{
		{
			name:     "Platinum seat",
			row:      "A",
			expected: utils.PlatinumPrice,
		},
		{
			name:     "Gold seat",
			row:      "B",
			expected: utils.GoldPrice,
		},
		{
			name:     "Silver seat",
			row:      "C",
			expected: utils.SilverPrice,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			theaterObj := NewTheater(nil) // We don't need actual theater object for this test

			// Call the function
			price := theaterObj.CalculateSeatPrice(Seat{Row: tc.row})

			// Verify result
			assert.Equal(t, tc.expected, price, "Seat price does not match expected price")
		})
	}
}

func TestCalculate(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name     string
		tax      Tax
		price    float64
		expected float64
	}{
		{
			name:     "GST",
			tax:      Tax{Name: "GST", Percentage: 10},
			price:    50,
			expected: 5,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the function
			price := tc.tax.Calculate(tc.price)
			if price != tc.expected {
				t.Errorf("Expected =%f and got %f", tc.expected, price)
			}

		})
	}
}
