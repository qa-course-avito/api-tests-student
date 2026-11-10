package myAdvertisements

import (
	"io"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/client/http/myAdvertisements"
)

// GetMyAdvertisements возвращает список объявлений авторизованного пользователя и проверяет ожидаемый статус-код
func GetMyAdvertisements(t *testing.T, token string, expectedStatusCode int, limit int) string {
	myAdvertisementsRawResponse := myAdvertisements.HttpGetMyAdvertisements(t, token, limit)
	defer myAdvertisementsRawResponse.Body.Close()

	require.Equalf(t, expectedStatusCode, myAdvertisementsRawResponse.StatusCode,
		"HTTP status code должен быть %d", expectedStatusCode)

	myAdvertisementsResponse, errResponse := io.ReadAll(myAdvertisementsRawResponse.Body)
	require.NoError(t, errResponse)

	return string(myAdvertisementsResponse)
}
func GetAdvertisementsList(t *testing.T, token string, text string) string {
	t.Helper()

	// 1. Добавили имя пакета myAdvertisements перед функцией клиента
	raw := myAdvertisements.HttpGetAdvertisements(t, token, text)
	defer raw.Body.Close()

	// 2. Читаем сырые байты ответа сервера
	body, err := io.ReadAll(raw.Body)
	require.NoError(t, err, "Не удалось прочитать тело ответа")

	// 3. Просто возвращаем весь JSON-ответ как строку string (как на строке 25)
	return string(body)
}
