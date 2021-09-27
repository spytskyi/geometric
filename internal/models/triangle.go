package models

type Triangle struct {
	tableName   struct{} `pg:"triangles"`
	Id          int64    `pg:"id,pk"`
	SideA       float64  `pg:"side_a"`
	SideB       float64  `pg:"side_b"`
	SideC       float64  `pg:"side_c"`
	InjectionAB float64  `pg:"injection_ab"`
	InjectionBC float64  `pg:"injection_bc"`
	InjectionAC float64  `pg:"injection_ac"`
	Square      float64  `pg:"square"`
}

type TriangleCreateReq struct {
	SideA       float64 `json:"side_a"`
	SideB       float64 `json:"side_b"`
	SideC       float64 `json:"side_c"`
	InjectionAB float64 `json:"injection_ab"`
	InjectionBC float64 `json:"injection_bc"`
	InjectionAC float64 `json:"injection_ac"`
	Square      float64 `json:"square"`
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
