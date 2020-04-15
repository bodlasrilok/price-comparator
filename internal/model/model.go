package model

type Quartile struct {
	Q1 float64
	Q2 float64
	Q3 float64
	Q4 float64
}
type QuartileForSolds struct {
	Q1 int64
	Q2 int64
	Q3 int64
	Q4 int64
}
type Limit struct {
	Upper float64
	Lower float64
}
type LimitForSolds struct {
	Upper int64
	Lower int64
}
type HistoryData struct {
	ProductID string  `json:"product_id" db:"product_id"`
	Price     float64 `json:"price" db:"price"`
	Stock     int64   `json:"stock" db:"stock"`
	Sold      int64   `json:"sold" db:"sold"`
	View      int64   `json:"view" db:"view"`
	CNI       float64
}

type Output struct {
	ProductID      int64   `csv:"product_id" db:"product_id"`
	NormalizedName string  `csv:"normalized_name" db:"normalized_name"`
	Price          float64 `csv:"price" db:"normalized_name"`
	Price1         float64 `csv:"price1" db:"price"`
	Price2         float64 `csv:"price2" db:"price"`
	Price3         float64 `csv:"price3" db:"price"`
}
