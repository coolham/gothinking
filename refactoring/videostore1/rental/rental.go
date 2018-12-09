package rental

import "gothinking/refactoring/videostore1/movie"

/*
Rental表示顾客租赁电影

*/

type Rental struct {
	AMovie     movie.Movie
	DaysRented int //租期
}
