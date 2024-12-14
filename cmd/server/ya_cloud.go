package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gennadyterekhov/auth-microservice/internal/app"
	"github.com/gennadyterekhov/auth-microservice/internal/logger"
	"github.com/gennadyterekhov/auth-microservice/internal/models/requests"
	"github.com/gennadyterekhov/auth-microservice/internal/models/responses"
	"github.com/gennadyterekhov/auth-microservice/internal/project/config"
)

// Handler for yandex cloud https://yandex.cloud/ru/docs/functions/lang/golang/handler
func Handler(ctx context.Context, request *requests.YandexCloudRequest) (*responses.YandexCloudResponse, error) {
	// он будет собирать приложение на каждый запрос ?! ну да а как иначе....
	logger.Debugln("main.Handler")

	_, appInstance, err := getAppInstance()
	if err != nil {
		return nil, err
	}

	logger.Debugln("got app instance")

	// 1. Парсим запрос от API Gateway

	// 2. Преобразуем запрос в *http.Request
	req, err := http.NewRequest(request.HTTPMethod, request.Path, strings.NewReader(request.Body))
	if err != nil {
		return &responses.YandexCloudResponse{StatusCode: 500, Body: err.Error()}, err
	}

	logger.Debugln("created http.NewRequest from requests.YandexCloudRequest")

	// Добавляем заголовки
	for k, v := range request.Headers {
		req.Header.Set(k, v)
	}

	logger.Debugln("set Headers")

	// Добавляем query параметры
	q := req.URL.Query()
	for k, v := range request.QueryString {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	logger.Debugln("set QueryString")

	// 3. Создаем ResponseRecorder для перехвата ответа
	rr := httptest.NewRecorder()
	logger.Debugln("created ResponseRecorder")

	logger.Debugln("launching appInstance.Router().ServeHTTP")

	// 4. Вызываем роутер
	appInstance.Router().ServeHTTP(rr, req)

	logger.Debugln("creating responses.YandexCloudResponse")
	logger.Debugln("rr.Code", rr.Code)
	logger.Debugln("rr.Body.String()", rr.Body.String())

	// 5. Преобразуем ответ в формат API Gateway
	resp := &responses.YandexCloudResponse{
		StatusCode: rr.Code,
		Headers:    rr.Header(),
		Body:       rr.Body.String(),
	}

	return resp, nil
}

func getAppInstance() (*config.Config, *app.App, error) {
	serverConfig, err := config.New()
	if err != nil {
		return nil, nil, err
	}

	appInstance, err := app.New(serverConfig.DBDsn)
	return serverConfig, appInstance, err
}
