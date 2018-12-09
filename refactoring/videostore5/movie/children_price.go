package movie

type ChildrenPrice struct {
}

func (c ChildrenPrice) GetPriceCode() int {
	return CHILDRES
}

func (c ChildrenPrice) GetCharge(daysRented int) float64 {
	result := 1.5
	if daysRented > 3 {
		result += float64(daysRented-3) * float64(2)
	}
	return result
}

func (c ChildrenPrice) GetFrequentRenterPoints(daysRented int) int {
	return 1
}
