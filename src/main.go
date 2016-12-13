// main
package main

import (
	"github.com/go-martini/martini"
	"github.com/regularprincess/RestServer/src/controllers"
)

func main() {
	m := martini.Classic()

	m.Get("/", func() string {
		return "Hello!"
	})

	m.Post("/room", controllers.CreateR)
	m.Get("/room/:id", controllers.ReadR)
	m.Post("/user", controllers.CreateU)
	m.Get("/user/:id", controllers.ReadU)
	m.Run()
}

//TODO
// поля моделей сделать Паблик
// Формат json {"id":12, "name":"Юрий", "idRoom":1, "phone":"8-800"}
