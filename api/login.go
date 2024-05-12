package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/parvesh12/gologin-rest/config"
	"golang.org/x/crypto/bcrypt"
)

type Loginreq struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type TblUser struct {
	Id        int    `gorm:"primarykey" json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	UserName  string `json:"username"`
	Email     string `json:"emailid"`
	Password  string `json:"password"`
}

var db = config.SetupDB()

// login api
func Login(w http.ResponseWriter, r *http.Request) {

	var user Loginreq

	decodeerr := json.NewDecoder(r.Body).Decode(&user)

	if decodeerr != nil {

		http.Error(w, decodeerr.Error(), http.StatusBadRequest)

		return
	}

	var duser TblUser

	if err := db.Model(TblUser{}).Where("user_name=?", user.UserName).First(&duser).Error; err != nil {

		fmt.Printf("Error Getting user: %s\n", err)

		resp := map[string]interface{}{
			"status":  404,
			"message": err.Error(),
		}

		respjson, _ := json.Marshal(resp)

		w.Write(respjson)

		return

	}

	perr := bcrypt.CompareHashAndPassword([]byte(duser.Password), []byte(user.Password))

	if perr != nil {

		http.Error(w, perr.Error(), http.StatusUnauthorized)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	resp := map[string]interface{}{
		"status":  200,
		"message": "logged in",
	}

	respjson, _ := json.Marshal(resp)

	w.Write(respjson)

}
