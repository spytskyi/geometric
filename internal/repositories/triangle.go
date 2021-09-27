package repositories

import (
	"fmt"
	"geometric/internal/models"

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
		Square:      54,
	}
	_, err = db.Model(&t).Insert()
	fmt.Println(t.Id)
	if err != nil {
		return
	}
	return nil
}
