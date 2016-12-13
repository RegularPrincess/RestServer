// users
package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-martini/martini"
	"github.com/regularprincess/RestServer/src/model"
)

func CreateU(w http.ResponseWriter, req *http.Request) {
	var u model.User
	//Проверяем тело на пустоту
	if req.Body == nil {
		w.WriteHeader(400)
		w.Write([]byte("Please send a request body"))
		return
	}
	//из джесон в обьект
	err := json.NewDecoder(req.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
	}
	//TODO Создавать отдельно
	//Создаем коннекшн
	db, err := model.NewPostgreDB("user=postgres password=root dbname=postgres sslmode=disable")
	//	defer db.Close()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	//Записываем в базу
	id, err := db.Create(&u)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	//Считываем из базы
	u, err = readU(id)
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	w.WriteHeader(201)
	w.Write(b.Bytes())

}

func ReadU(w http.ResponseWriter, params martini.Params) { //:id
	id, err := strconv.Atoi(params["id"])
	u, err := readU(id)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	w.WriteHeader(200)
	w.Write(b.Bytes())
}

func readU(id int) (model.User, error) {
	var u model.User
	//Создавать отдельно
	db, err := model.NewPostgreDB("user=postgres password=root dbname=postgres sslmode=disable")
	//	defer db.Close()
	if err != nil {
		return u, err
	}
	row, err := db.Query("SELECT * FROM users WHERE id = " + strconv.Itoa(id) + ";")
	row.Next()
	err = row.Scan(&u.Id, &u.Name, &u.IdRoom, &u.Phone)
	if err != nil {
		return u, err
	}
	return u, nil
}
