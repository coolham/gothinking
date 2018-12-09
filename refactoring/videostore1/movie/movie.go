package movie

const (
	REGULAR     = iota //0  普通电影
	NEW_RELEASE        //1  新发行电影
	CHILDRES           //2  儿童电影
)

type Movie struct {
	Title     string //片名
	PriceCode int    //价格代号
}
