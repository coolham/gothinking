package movie

const (
	REGULAR     = iota //0
	NEW_RELEASE        //1
	CHILDREN           // 2
)

type Movie struct {
	Title     string //片名
	PriceCode int    //价格代号
}

func (m Movie) GetTitle() string {
	return m.Title
}

func (m Movie) GetPriceCode() int {
	return m.PriceCode
}

func (m Movie) GetCharge(daysRented int) float64 {
	result := 0.0
	switch m.GetPriceCode() {
	case REGULAR:
		result += 2
		if daysRented > 2 {
			result += float64(daysRented-2.0) * 1.5
		}
	case NEW_RELEASE:
		result += float64(daysRented) * 3
	case CHILDREN:
		result += 1.5
		if daysRented > 3 {
			result += float64(daysRented-3) * 1.5
		}
	}
	return result
}

func (m Movie) GetFrequentRenterPoints(daysRented int) int {
	if m.GetPriceCode() == NEW_RELEASE && daysRented > 1 {
		return 2
	} else {
		return 1
	}
}
