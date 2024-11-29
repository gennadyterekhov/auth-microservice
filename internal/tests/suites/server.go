package suites

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gennadyterekhov/auth-microservice/internal/interfaces"
	"github.com/gennadyterekhov/auth-microservice/internal/tests/inits"
	"github.com/stretchr/testify/require"
)

type WithServer struct {
	WithFactory
	server *httptest.Server
}

var _ interfaces.WithServer = &WithServer{}

func (s *WithServer) SetupSuite() {
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
	require.NoError(s.T(), err)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	response, err := s.GetServer().Client().Do(req)
	require.NoError(s.T(), err)

	bodyAsBytes, err := getBodyAsBytes(response.Body)
	response.Body.Close()
	require.NoError(s.T(), err)

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
	require.NoError(s.T(), err)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	response, err := s.GetServer().Client().Do(req)
	require.NoError(s.T(), err)

	bodyAsBytes, err := getBodyAsBytes(response.Body)
	response.Body.Close()
	require.NoError(s.T(), err)

	return response.StatusCode, bodyAsBytes
}

func getBodyAsBytes(reader io.Reader) ([]byte, error) {
	readBytes, err := io.ReadAll(reader)
	if err != nil {
		return []byte{}, err
	}

	return readBytes, nil
}
