package myAdvertisements

import (
	"net/http"
	"strconv"
	"testing"

	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
)

// HttpGetMyAdvertisements выполняет GET-запрос для получения списка объявлений авторизованного пользователя
func HttpGetMyAdvertisements(t *testing.T, token string, limit int) *http.Response {
	return apiRunner.GetRunner().Auth(token).Create().Get(path.MyAdvertisementsPath).
		Query("limit", strconv.Itoa(limit)).
		Expect(t).
		End().Response
}

// HttpGetAdvertisements выполняет GET-запрос для поиска объявлений по названию
func HttpGetAdvertisements(t *testing.T, token string, textForSearch string) *http.Response {
	return apiRunner.GetRunner().Auth(token).Create().
		Get(path.AdvertisementsPath).
		Query("text", textForSearch).
		Expect(t).
		Status(http.StatusOK).
		End().Response
}

