package repositories

import (
	"errors"
	"fmt"
	"geometric/internal/models"
	"math"

	"github.com/go-pg/pg/v10"
)

func CreateTriangle(db *pg.DB, reqData *models.TriangleCreateReq) (err error) {

	t := models.Triangle{
		SideA:       reqData.SideA,
		SideB:       reqData.SideB,
		SideC:       reqData.SideC,
		InjectionAB: reqData.InjectionAB,
		InjectionBC: reqData.InjectionBC,
		InjectionAC: reqData.InjectionAC,
	}

	area, err := CalculateTriangleArea(&t)
	t.Area = area
	fmt.Println(t.Area)
	if err != nil {
		return
	}

	_, err = db.Model(&t).Insert()
	if err != nil {
		return
	}

	return nil
}

func CalculateTriangleArea(reqData *models.Triangle) (float64, error) {
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

func SelectDataTriangleById(db *pg.DB, param int64) (err error) {
	triangle := models.Triangle{Id: param}
	err = db.Model(&triangle).WherePK().Select()
	if err != nil {
		return
	}
	fmt.Println(&triangle)
	return
}
