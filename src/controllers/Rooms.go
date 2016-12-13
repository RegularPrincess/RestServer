// rooms
package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-martini/martini"
	"github.com/regularprincess/RestServer/src/model"
)

func CreateR(w http.ResponseWriter, req *http.Request) {
	var r model.Room
	//Проверяем тело на пустоту
	if req.Body == nil {
		w.WriteHeader(400)
		w.Write([]byte("Please send a request body"))
		return
	}
	//из джесон в обьект
	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
	}
	//TODO Создавать отдельно
	//Создаем коннекшн
	db, err := model.NewPostgreDB("user=postgres password=root dbname=postgres sslmode=disable")
	//defer db.Close()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	//Записываем в базу
	id, err := db.Create(&r)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	//Считываем из базы
	r, err = readR(id)
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)
	w.WriteHeader(201)
	w.Write(b.Bytes())

}

func ReadR(w http.ResponseWriter, params martini.Params) { //:id
	id, err := strconv.Atoi(params["id"])
	r, err := readR(id)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)
	w.WriteHeader(200)
	w.Write(b.Bytes())
}

func readR(id int) (model.Room, error) {
	var r model.Room
	//Создавать отдельно
	db, err := model.NewPostgreDB("user=postgres password=root dbname=postgres sslmode=disable")
	//defer db.Close()
	if err != nil {
		return r, err
	}
	row, err := db.Query("SELECT * FROM room WHERE id = " + strconv.Itoa(id) + ";")
	row.Next()
	err = row.Scan(&r.Id, &r.Number, &r.Adress)
	if err != nil {
		return r, err
	}
	return r, nil
}
