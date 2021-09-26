package models

type Triangle struct {
	Id          int64
	sideA       float64
	sideB       float64
	sideC       float64
	injectionAB float64
	injectionBC float64
	injectionAC float64
}

type TriangleCreateReq struct {
	Name string
}
