// report package
package report

import (
	"fmt"
)

// PrintSalesReport prints the total sales report
func PrintSalesReport(totalRevenue, totalServiceTax, totalSwachhCess, totalKrishiCess float64) {
	fmt.Println("\nTotal Sales:")
	fmt.Printf("Revenue: Rs. %.0f\n", totalRevenue)
	fmt.Printf("Service Tax: Rs. %.2f\n", totalServiceTax)
	fmt.Printf("Swachh Bharat Cess: Rs. %.2f\n", totalSwachhCess)
	fmt.Printf("Krishi Kalyan Cess: Rs. %.2f\n", totalKrishiCess)
}
