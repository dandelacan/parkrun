package controllers

import (
	"fmt"
	"net/http"
	"parkrun/backend/models"
	"strconv"
)

func NewUsers(us *models.UserService) *Users {
	return &Users{
		us: us,
	}
}

type Users struct {
	us *models.UserService
}

func (u *Users) ViewUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.us.AllUsersse()
	if err != nil {
		panic(err)
	}
	message := ""

	for _, v := range users {
		fmt.Printf("%T\n", v.BarcodeID)
		message = message + v.Name + "  " + strconv.FormatInt(int64(v.BarcodeID), 10) + "\n"
	}
	fmt.Fprint(w, message)
}
