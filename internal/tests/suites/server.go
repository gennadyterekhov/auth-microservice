package suites

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gennadyterekhov/auth-microservice/internal/interfaces"
	"github.com/gennadyterekhov/auth-microservice/internal/tests/inits"
)

type WithServer struct {
	WithService
	server *httptest.Server
}

var _ interfaces.WithServer = &WithServer{}

func (s *WithServer) SetupSuite() {
	fmt.Println("(s *WithServer) SetupSuite() ")
	fmt.Println()
	inits.InitFactorySuite(s)
}

func (s *WithServer) SetServer(srv *httptest.Server) {
	s.server = srv
}

func (s *WithServer) GetServer() *httptest.Server {
	return s.server
}

func (s *WithServer) SendGet(
	path string,
	token string,
) (int, []byte) {
	req, err := http.NewRequest(http.MethodGet, s.GetServer().URL+path, strings.NewReader(""))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	response, err := s.GetServer().Client().Do(req)
	if err != nil {
		panic(err)
	}
	bodyAsBytes, err := getBodyAsBytes(response.Body)
	response.Body.Close()
	if err != nil {
		panic(err)
	}
	return response.StatusCode, bodyAsBytes
}

func (s *WithServer) SendPostWithoutToken(
	path string,
	requestBody *bytes.Buffer,
) int {
	code, _ := s.SendPostAndReturnBody(path, "", requestBody)

	return code
}

func (s *WithServer) SendPost(
	path string,
	token string,
	requestBody *bytes.Buffer,
) int {
	code, _ := s.SendPostAndReturnBody(path, token, requestBody)

	return code
}

func (s *WithServer) SendPostAndReturnBody(
	path string,
	token string,
	requestBody *bytes.Buffer,
) (int, []byte) {
	req, err := http.NewRequest(http.MethodPost, s.GetServer().URL+path, requestBody)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	response, err := s.GetServer().Client().Do(req)
	if err != nil {
		panic(err)
	}
	bodyAsBytes, err := getBodyAsBytes(response.Body)
	response.Body.Close()
	if err != nil {
		panic(err)
	}
	return response.StatusCode, bodyAsBytes
}

func getBodyAsBytes(reader io.Reader) ([]byte, error) {
	readBytes, err := io.ReadAll(reader)
	if err != nil {
		return []byte{}, err
	}

	return readBytes, nil
}
