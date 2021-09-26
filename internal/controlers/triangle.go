package controlers

import (
	"fmt"
	"geometric/internal/models"
	"geometric/internal/repositories"

	"github.com/go-pg/pg/v10"
	"github.com/kataras/iris/v12"
)

func GetSquare(ctx iris.Context, db *pg.DB) {
	var request models.TriangleCreateReq
	err := ctx.ReadBody(&request)
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Triangle creation failure").DetailErr(err))
		return
	}
	fmt.Println(request.Name)
	err = repositories.CreateTriangle(db, request)
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Triangle creation failure").DetailErr(err))
		return
	}
	ctx.StatusCode(iris.StatusCreated)
	_, err = ctx.JSON(iris.Map{"message": "OK"})
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Triangle creation failure").DetailErr(err))
		return
	}
	return
}
