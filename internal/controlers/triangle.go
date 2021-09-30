package controlers

import (
	"geometric/internal/models"
	"geometric/internal/repositories"

	"github.com/go-pg/pg/v10"
	"github.com/kataras/iris/v12"
)

func SetArea(ctx iris.Context, db *pg.DB) {
	var request models.TriangleCreateReq
	err := ctx.ReadBody(&request)
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Triangle creation failure").DetailErr(err))
		return
	}

	err = repositories.CreateTriangle(db, &request)
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

func GetArea(ctx iris.Context, db *pg.DB) {
	request, err := ctx.Params().GetInt64("id")
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Triangle creation failure").DetailErr(err))
		return
	}
	err = repositories.SelectDataTriangleById(db, request)
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

func GetAllArea(ctx iris.Context, db *pg.DB) {
	err := repositories.SelectAllTriangles(db)
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
