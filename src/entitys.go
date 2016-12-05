// entitys
package main

type User struct {
	id     int
	name   string
	idRoom int
	phone  string
}

type Room struct {
	id     int
	number int
	adress string
}

type Cost struct {
	id       int
	amount   float32
	idAuthor int
	object   string
}

type Debit struct {
	pyingUser int
	idCost    int
}
