// entitys
package model

import "fmt"
import "strconv"

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	IdRoom int    `json:"idRoom"`
	Phone  string `json:"phone"`
}

func (u *User) GetTableName() string {
	return "users"
}
func (u *User) GetTableFields() string {
	return "name, id_room, phone"
}
func (u *User) GetFields() string {
	return fmt.Sprintf("'%s', %s, '%s'", u.Name, strconv.Itoa(u.IdRoom), u.Phone)
}

type Room struct {
	Id     int
	Number int
	Adress string
}

func (r *Room) GetTableName() string {
	return "room"
}
func (r *Room) GetTableFields() string {
	return "number, adress"
}
func (r *Room) GetFields() string {
	return fmt.Sprintf("%s, '%s'", strconv.Itoa(r.Number), r.Adress)
}

type Cost struct {
	Id       int
	Amount   float32
	IdAuthor int
	Object   string
}

func (c *Cost) GetTableName() string {
	return "cost"
}
func (c *Cost) GetTableFields() string {
	return "amount, id_author, object"
}
func (c *Cost) GetFields() string {
	return fmt.Sprintf("%s, %s, '%s'", fmt.Sprintf("%.2f", c.Amount), strconv.Itoa(c.IdAuthor), c.Object)
}

type Debit struct {
	PyingUser int
	IdCost    int
}

func (d *Debit) GetTableName() string {
	return "debit"
}
func (d *Debit) GetTableFields() string {
	return "pying_user, id_coast"
}
func (d *Debit) GetFields() string {
	return fmt.Sprintf("%s, %s", strconv.Itoa(d.PyingUser), strconv.Itoa(d.IdCost))
}
