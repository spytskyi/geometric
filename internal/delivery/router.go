package delivery

import (
	"geometric/internal/controlers"

	"github.com/go-pg/pg/v10"

	"github.com/kataras/iris/v12"
)

func NewRouter(app *iris.Application, db *pg.DB) *iris.Application {
	trianglesAPI := app.Party("/triangles")
	{
		// middleware
		trianglesAPI.Use(iris.Compression)

		trianglesAPI.Get("/get-square", func(ctx iris.Context) {
			controlers.GetSquare(ctx, db)
		})

		trianglesAPI.Post("/add-square", func(ctx iris.Context) {
			controlers.GetSquare(ctx, db)
		})
	}
	return app
}

//migrate -source file://postgres/migrations/ -database postgres://postgres:password@localhost:5432/geometric_figure?sslmode=disable up 1
