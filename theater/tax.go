package theater

type Tax struct {
	Name       string
	Percentage float64
}

func (tax Tax) Calculate(price int) float64 {
	return float64(price) * tax.Percentage / 100
}
