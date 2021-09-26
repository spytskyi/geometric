package main

import (
	"context"
	"geometric/conf"
	"geometric/internal/delivery"
	"geometric/postgres"

	"github.com/kataras/iris/v12"
)

func main() {

	conf.InitConfig()

	db := postgres.InitPostgres()

	defer db.Close()

	ctx := context.Background()

	if err := db.Ping(ctx); err != nil {
		panic(err)
	}

	app := iris.Default()

	app = delivery.NewRouter(app, db)
	// This handler will match /user/john but will not match /user/ or /user

	app.Listen(":8080")
}

//func (t Triangle) String() string {
//	return fmt.Sprintf("User<%d %f %f %f %f %f %f>", t.Id, t.sideA, t.sideB, t.sideC, t.injectionAB, t.injectionBC, t.injectionAC)
//}

//func (t *Triangle) calculateSquareTriangle() (float64, error) {
//	if (t.a <= 0) && (t.b <= 0) && (t.c <= 0) {
//		return float64(0), errors.New("Сторона не может равняться нулю либо быть отрицательным числом")
//	}
//	hp := (t.a + t.b + t.c) / 2
//	return math.Sqrt(hp * (hp - t.a) * (hp - t.b) * (hp - t.c)), nil
//}

//type Story struct {
//	Id       int64
//	Title    string
//	AuthorId int64
//	Author   *User `pg:"rel:has-one"`
//}
//
//func (s Story) String() string {
//	return fmt.Sprintf("Story<%d %s %s>", s.Id, s.Title, s.Author)
//}
//func (t *Triangle) sAllSide() (float64, error) {
//	if (t.a <= 0) && (t.b <= 0) && (t.c <= 0) {
//		return float64(0), errors.New("Сторона не может равняться нулю либо быть отрицательным числом")
//	}
//	hp := (t.a + t.b + t.c) / 2
//	return math.Sqrt(hp * (hp - t.a) * (hp - t.b) * (hp - t.c)), nil
//}
//
//func (t *Triangle) sTriangle() {
//	value, err := t.sAllSide()
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	fmt.Println("S triangle: ", value)
//}
//func foo(n int) error {
//	v := reflect.ValueOf(n)
//	if v.Kind() != reflect.Int {
//		return errors.New("Передаваться должно число")
//	}
//	return nil
//}
//
//func printFoo(n int) {
//	err := foo(n)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	for i := 0; i <= n; i++ {
//		fmt.Println("2 в", i, "степени =", math.Pow(float64(2), float64(i)))
//	}
//}
//
//func calculateRadiusCircle(radius int) (float64, error) {
//	if radius <= 0 {
//		return float64(0), errors.New("Радиус не может быть отрецательным либо равняться 0")
//	}
//	return float64(radius) * float64(radius) * pi, nil
//}
//
//func printValueRadiusCircle(radius int) {
//	circleArea, err := calculateRadiusCircle(radius)
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	fmt.Println("Radius is: ", circleArea)
//}
//
////S = 0,5 * a * b⋅sin(α)
//func squareOfTriangle(injection float64, a float64, b float64) (float64, error) {
//	if (injection <= 0) || (a <= 0) || (b <= 0) {
//		return float64(0), errors.New("Передаваемые значения не могут быть меньше 0")
//	}
//	sinAB := float64(math.Sin(injection))
//	square := 0.5 * a * b * sinAB
//	return square, nil
//}

//func printSquareTriangle(injection float64, a float64, b float64) {
//	result, err := squareOfTriangle(injection, a, b)
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	fmt.Println("Square of triangle: ", result)
//	return
//}
//
//func squareOfRectangle(a float64, b float64) (float64, error) {
//	if (a <= 0) || (b <= 0) {
//		return float64(0), errors.New("Передаваемые значения не могут быть меньше 0")
//	}
//	square := a * b
//	return float64(square), nil
//}
//
//func printSquareRectangle(a float64, b float64) {
//	square, err := squareOfRectangle(a, b)
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	fmt.Println("Square of rectangle: ", square)
//	return
//}
//
//func perimeterRectangle(a float64, b float64) (float64, error) {
//	if (a <= 0) || (b <= 0) {
//		return float64(0), errors.New("Сторона не может равняться 0 либо быть отрицательным значением")
//	}
//	pRect := (2 * a) + (2 * b)
//	return float64(pRect), nil
//}
//
//func printPerimetrRectangle(a float64, b float64) {
//	pRect, err := perimeterRectangle(a, b)
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	fmt.Println("Периметр прямоугольника: ", pRect)
//	return
//}
//
//func ExampleDB_Model() {
//	db := pg.Connect(&pg.Options{
//		User: "postgres",
//	})
//	defer db.Close()
//
//	err := createSchema(db)
//	if err != nil {
//		panic(err)
//	}
//
//	user1 := &User{
//		Name:   "admin",
//		Emails: []string{"admin1@admin", "admin2@admin"},
//	}
//	_, err = db.Model(user1).Insert()
//	if err != nil {
//		panic(err)
//	}
//
//	_, err = db.Model(&User{
//		Name:   "root",
//		Emails: []string{"root1@root", "root2@root"},
//	}).Insert()
//	if err != nil {
//		panic(err)
//	}
//
//	story1 := &Story{
//		Title:    "Cool story",
//		AuthorId: user1.Id,
//	}
//	_, err = db.Model(story1).Insert()
//	if err != nil {
//		panic(err)
//	}
//
//	// Select user by primary key.
//	user := &User{Id: user1.Id}
//	err = db.Model(user).WherePK().Select()
//	if err != nil {
//		panic(err)
//	}
//
//	// Select all users.
//	var users []User
//	err = db.Model(&users).Select()
//	if err != nil {
//		panic(err)
//	}
//
//	// Select story and associated author in one query.
//	story := new(Story)
//	err = db.Model(story).
//		Relation("Author").
//		Where("story.id = ?", story1.Id).
//		Select()
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Println(user)
//	fmt.Println(users)
//	fmt.Println(story)
//	// Output: User<1 admin [admin1@admin admin2@admin]>
//	// [User<1 admin [admin1@admin admin2@admin]> User<2 root [root1@root root2@root]>]
//	// Story<1 Cool story User<1 admin [admin1@admin admin2@admin]>>
//}
//
//func createSchema(db *pg.DB) error {
//	models := []interface{}{
//		(*User)(nil),
//		(*Story)(nil),
//	}
//
//	for _, model := range models {
//		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
//			Temp: true,
//		})
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}
