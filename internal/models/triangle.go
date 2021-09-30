package models

type Triangle struct {
	tableName  struct{} `pg:"triangles"`
	Id         int64    `pg:"id,pk"`
	SideA      float64  `pg:"side_a"`
	SideB      float64  `pg:"side_b"`
	SideC      float64  `pg:"side_c"`
	InjectionA float64  `pg:"injection_a"`
	InjectionB float64  `pg:"injection_b"`
	InjectionC float64  `pg:"injection_a"`
	Area       float64  `pg:"area"`
}

type TriangleCreateReq struct {
	SideA      float64 `json:"side_a"`
	SideB      float64 `json:"side_b"`
	SideC      float64 `json:"side_c"`
	InjectionA float64 `json:"injection_a"`
	InjectionB float64 `json:"injection_b"`
	InjectionC float64 `json:"injection_c"`
	Area       float64 `json:"area"`
}

type TriangleGetReq struct {
	Id int64 `pg:"id,pk"`
}

//type TriangleCreateReq struct {
//	Id          int64 `json:"ref"`
//	sideA       float64
//	sideB       float64
//	sideC       float64
//	injectionAB float64
//	injectionBC float64
//	injectionAC float64
//	square      float64
//}
