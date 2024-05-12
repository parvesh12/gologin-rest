package api

import (
	"encoding/json"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {

	var user TblUser

	jerr := json.NewDecoder(r.Body).Decode(&user)

	if jerr != nil {

		http.Error(w, jerr.Error(), http.StatusInternalServerError)

		return
	}

	if err := db.Model(TblUser{}).Create(&user).Error; err != nil {

		resp := map[string]interface{}{
			"status":  500,
			"message": err.Error(),
		}

		respjson, _ := json.Marshal(resp)

		w.Write(respjson)

		return
	}

	resp := map[string]interface{}{
		"status":  200,
		"message": "Registered Successfully",
	}

	respjson, _ := json.Marshal(resp)

	w.Write(respjson)
}
