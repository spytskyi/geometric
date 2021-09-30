package repositories

import (
	"errors"
	"fmt"
	"geometric/internal/models"
	"math"

	"github.com/go-pg/pg/v10"
)

//func CreateTriangle ...
func CreateTriangle(db *pg.DB, reqData *models.TriangleCreateReq) (err error) {
	var area float64

	t := models.Triangle{
		SideA:      reqData.SideA,
		SideB:      reqData.SideB,
		SideC:      reqData.SideC,
		InjectionA: reqData.InjectionA,
		InjectionB: reqData.InjectionB,
		InjectionC: reqData.InjectionC,
	}

	if t.SideA > 0 && t.SideB > 0 && t.SideC > 0 {
		area, err = CalculatingAreaTriangleOnThreeSides(&t)
	}

	if t.SideA > 0 && t.InjectionB > 0 && t.InjectionC > 0 {
		area, err = CalculatingAreaTriangleOnSideAndTwoInjections(&t)
	}
	if t.SideA > 0 && t.SideB > 0 && t.InjectionC > 0 {
		area, err = CalculatingAreaTriangleOnTwoSidesAndInjectionBetweenThem(&t)
	}

	t.Area = area

	if err != nil {
		return
	}

	_, err = db.Model(&t).Insert()
	if err != nil {
		return
	}

	return nil
}

//func CalculatingAreaTriangleOnThreeSides ...
func CalculatingAreaTriangleOnThreeSides(reqData *models.Triangle) (float64, error) {
	var m = map[string]float64{
		"SideA": reqData.SideA,
		"SideB": reqData.SideB,
		"SideC": reqData.SideC,
	}

	for _, value := range m {
		if value <= 0 {
			return float64(0), errors.New("Data entry error")
		}
	}

	hp := (reqData.SideA + reqData.SideB + reqData.SideC) / 2

	for _, value := range m {
		if (hp - value) < 0 {
			return float64(0), errors.New("A triangle with such dimensions does not exist")
		}
		if (hp - value) == 0 {
			return float64(0), errors.New("The area of a triangle is zero")
		}
	}
	return math.Round(math.Sqrt(hp*(hp-reqData.SideA)*(hp-reqData.SideB)*(hp-reqData.SideC))*10) / 10, nil
}

//func CalculatingAreaTriangleOnSideAndTwoInjections ...
func CalculatingAreaTriangleOnSideAndTwoInjections(reqData *models.Triangle) (float64, error) {
	var m = map[string]float64{
		"SideA":      reqData.SideA,
		"InjectionB": reqData.InjectionB,
		"InjectionC": reqData.InjectionC,
	}

	for _, value := range m {
		if value <= 0 {
			return float64(0), errors.New("1Data entry error")
		}
	}
	if (m["InjectionB"] + m["InjectionC"]) >= 180 {
		return float64(0), errors.New("2Data entry error")
	}

	m["InjectionA"] = 180 - (m["InjectionB"] + m["InjectionC"])

	area := 0.5 * (m["SideA"] * m["SideA"]) * (math.Sin(m["InjectionB"]) * math.Sin(m["InjectionC"]) / math.Sin(m["InjectionA"]))

	return area, nil
}

//func CalculatingAreaTriangleOnTwoSidesAndInjectionBetweenThem ...
func CalculatingAreaTriangleOnTwoSidesAndInjectionBetweenThem(reqData *models.Triangle) (float64, error) {

	var m = map[string]float64{
		"SideA":      reqData.SideA,
		"SideB":      reqData.SideB,
		"InjectionC": reqData.InjectionC,
	}

	for _, value := range m {
		if value <= 0 {
			return float64(0), errors.New("Data entry error")
		}
	}

	if m["InjectionC"] > 179 {
		return float64(0), errors.New("Data entry error")
	}
	return 0.5 * m["SideA"] * m["SideB"] * math.Sin(m["InjectionC"]), nil
}

//func SelectDataTriangleById ...
func SelectDataTriangleById(db *pg.DB, param int64) (err error) {
	triangle := models.Triangle{Id: param}
	err = db.Model(&triangle).WherePK().Select()
	if err != nil {
		return
	}
	fmt.Println(&triangle)
	return
}

func SelectAllTriangles(db *pg.DB) (err error) {
	var triangles []models.Triangle
	err = db.Model(&triangles).Select()
	if err != nil {
		return
	}
	fmt.Println(&triangles)
	return
}
