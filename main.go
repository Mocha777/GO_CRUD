package main

import (
	//"github.com/kataras/iris"
	"strconv"

	"github.com/kataras/iris/context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var students []Student
var student Student

func rootHandler(ctx iris.Context) {
	ctx.JSON(context.Map{"message": "welcome to CRUD service"})
}

func getStudents(ctx iris.Context) {
	ctx.JSON(students)
}

func getStudent(ctx iris.Context) {
	params, err := strconv.Atoi(ctx.Params().Get("id"))
	if err != nil {
		ctx.JSON(context.Map{"message": "ID ERROR"})
		return
	}
	for _, item := range students {
		if item.ID == params {
			ctx.JSON(item)
			return
		}
	}
}

func createStudent(ctx iris.Context) {

	err := ctx.ReadJSON(&student)
	if err != nil {
		ctx.JSON(context.Map{"message": "ERROR"})
		return
	}
	student.ID = students[len(students)-1].ID + 1
	students = append(students, student)
}

func updateStudent(ctx iris.Context) {
	//get id
	params, err := strconv.Atoi(ctx.Params().Get("id"))
	//get student name
	_ = ctx.ReadJSON(&student)
	if err != nil {
		ctx.JSON(context.Map{"message": "ERROR"})
		return
	}
	//update the name
	for index, item := range students {
		if params == item.ID {
			students[index].Name = student.Name
			return
		}
	}
}

func deleteStudent(ctx iris.Context) {
	params, err := strconv.Atoi(ctx.Params().Get("id"))
	if err != nil {
		ctx.JSON(context.Map{"message": "ERROR"})
		return
	}
	//find out and delete
	for index, item := range students {
		if item.ID == params {
			students = append(students[:index], students[index+1:]...)
			return
		}
	}
}

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	app.Logger().SetLevel("debug")

	//set original data
	students = append(students, Student{ID: 1, Name: "Train"})
	students = append(students, Student{ID: 2, Name: "Todd"})

	app.Handle("GET", "/", rootHandler)
	app.Handle("GET", "/students", getStudents)
	app.Handle("GET", "/students/{id: int}", getStudent)
	app.Handle("POST", "/students", createStudent)
	app.Handle("PUT", "/students/{id: int}", updateStudent)
	app.Handle("DELETE", "/students/{id: int}", deleteStudent)
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
