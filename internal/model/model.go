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