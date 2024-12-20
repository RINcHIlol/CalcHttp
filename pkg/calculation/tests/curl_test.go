package tests

import (
	"bytes"
	"calc_http/internal/application"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalcAPI(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    map[string]string
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name:           "Valid Expression",
			requestBody:    map[string]string{"expression": "2+2*2"},
			expectedStatus: http.StatusOK,
			expectedBody:   map[string]interface{}{"result": 6.0},
		},
		{
			name:           "Invalid Expression",
			requestBody:    map[string]string{"expression": "ikujgf"},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   map[string]interface{}{"error": "Expression is not valid"},
		},
		{
			name:           "Internal server error",
			requestBody:    map[string]string{"expression": ""},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   map[string]interface{}{"error": "Internal server error"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаём тестовый HTTP-запрос
			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			// Создаём тестовый HTTP-ответ
			rec := httptest.NewRecorder()

			// Вызываем обработчик
			handler := http.HandlerFunc(application.Calc)
			handler.ServeHTTP(rec, req)

			// Проверяем статус ответа
			if rec.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rec.Code)
			}

			// Проверяем тело ответа
			respBody, err := ioutil.ReadAll(rec.Body)
			if err != nil {
				t.Fatalf("Failed to read response body: %v", err)
			}

			var respData map[string]interface{}
			err = json.Unmarshal(respBody, &respData)
			if err != nil {
				t.Fatalf("Failed to unmarshal response body: %v", err)
			}

			if len(tt.expectedBody) != len(respData) {
				t.Errorf("Expected body length %d, got %d", len(tt.expectedBody), len(respData))
			}

			for key, expectedValue := range tt.expectedBody {
				if respData[key] != expectedValue {
					t.Errorf("Expected %s: %v, got: %v", key, expectedValue, respData[key])
				}
			}
		})
	}
}
