package handler

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	gdatabase "github.com/pilinux/gorest/database"
	gmodel "github.com/pilinux/gorest/database/model"

	"github.com/pilinux/gorest/example/database/model"
)

// GetHobbies handles jobs for controller.GetHobbies
func GetHobbies() (httpResponse gmodel.HTTPResponse, httpStatusCode int) {
	db := gdatabase.GetDB()
	hobbies := []model.Hobby{}

	if err := db.Find(&hobbies).Error; err != nil {
		log.WithError(err).Error("error code: 1251")
		httpResponse.Result = "internal server error"
		httpStatusCode = http.StatusInternalServerError
		return
	}

	if len(hobbies) == 0 {
		httpResponse.Result = "no hobby found"
		httpStatusCode = http.StatusNotFound
		return
	}

	httpResponse.Result = hobbies
	httpStatusCode = http.StatusOK
	return
}