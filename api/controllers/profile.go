package controllers

import (
	"crud/api/model"
	"crud/config/respon"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var (
	respn1 = map[string]interface{}{"status": true, "message": "Succes", "code": 200}
)

func (a *App) Profile(w http.ResponseWriter, r *http.Request) {

	profile := &model.Profile{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respon.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = json.Unmarshal(body, &profile)
	if err != nil {
		respon.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = profile.Validate("")
	if err != nil {
		respon.ERROR(w, http.StatusBadRequest, err)
		return
	}

	userCreated, err := profile.SaveProfile(a.DB)
	if err != nil {
		respon.ERROR(w, http.StatusBadRequest, err)
		return
	}
	respn1["data"] = userCreated
	respon.JSON(w, http.StatusCreated, respn1)
	return

}

func (a *App) Update(w http.ResponseWriter, r *http.Request) {

	profile := &model.Profile{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respon.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = json.Unmarshal(body, &profile)
	if err != nil {
		respon.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = profile.Validate("update")
	if err != nil {
		respon.ERROR(w, http.StatusBadRequest, err)
		return
	}
	userCreated, err := profile.Update(profile.ID, a.DB)
	if err != nil {
		respon.ERROR(w, http.StatusBadRequest, err)
		return
	}
	respn1["data"] = userCreated
	respon.JSON(w, http.StatusCreated, respn1)
	return
}

//untuk fungsi delete/hapus
func (a *App) Delete(w http.ResponseWriter, r *http.Request) {
	profile := &model.Profile{}
	userId := r.URL.Query().Get("id")
	users, err := profile.Delete(a.DB, userId)
	if err != nil {
		respon.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	respn1["data"] = users
	respon.JSON(w, http.StatusOK, respn1)
	return
}

//untuk join tabel 92-103
// func (a *App) JoinTabel(w http.ResponseWriter, r *http.Request) {
// 	var respn1 = map[string]interface{}{"status": true, "message": "Sukses", "code": 200}
// 	userId := r.URL.Query().Get("cari") // untuk searching
// 	data := &model.Profile{}
// 	users, err := data.Delete(a.DB, userId)
// 	if err != nil {
// 		respon.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	respn1["data"] = users
// 	respon.JSON(w, http.StatusOK, respn1)
// 	return
// }
