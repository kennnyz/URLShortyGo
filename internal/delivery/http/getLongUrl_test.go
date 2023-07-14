package http

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	. "ozonTech/muhtarov/internal/delivery/mock"
	"ozonTech/muhtarov/internal/models"
	"testing"
)

func TestHandle_GetLongURL(t *testing.T) {
	type mockBehavior func(shorty *MockURLShorty, shortURL string)
	testTable := []struct {
		name               string
		testId             int
		inputShortURL      string
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedLongURL    string
	}{
		{
			name:          "OK",
			testId:        1,
			inputShortURL: "1wyI4Mkj7L",
			mockBehavior: func(shorty *MockURLShorty, shortURL string) {
				shorty.EXPECT().GetFullUrl(shortURL).Return(models.UrlStruct{
					LongUrl:  "www.google.com",
					ShortUrl: shortURL,
					Id:       26413135951547755,
				}, nil)
			},
			expectedLongURL:    "www.google.com",
			expectedStatusCode: 200,
		},
		{
			name:               "Empty short url",
			testId:             2,
			inputShortURL:      "",
			mockBehavior:       func(shorty *MockURLShorty, shortURL string) {},
			expectedLongURL:    models.NotValidUrlErr.Error() + "\n",
			expectedStatusCode: 400,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			urlShorty := NewMockURLShorty(ctrl)
			testCase.mockBehavior(urlShorty, testCase.inputShortURL)

			handler := NewHandler(urlShorty)
			mux := http.NewServeMux()
			mux.HandleFunc("/get-long-url", handler.getLongUrlByShort)

			req := httptest.NewRequest(http.MethodGet, "/get-long-url?url="+testCase.inputShortURL, nil)
			recorder := httptest.NewRecorder()

			mux.ServeHTTP(recorder, req)
			assert.Equal(t, testCase.expectedStatusCode, recorder.Code)
			assert.Equal(t, testCase.expectedLongURL, recorder.Body.String())
		})
	}
}
