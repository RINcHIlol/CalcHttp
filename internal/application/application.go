package application

import (
	"calc_http/pkg/calculation"
	"encoding/json"
	"net/http"
)

type Request struct {
	Expression string `json:"expression"`
}

type Result struct {
	Result float64 `json:"result"`
}

func RunServer() error {
	http.HandleFunc("/api/v1/calculate", Calc)

	return http.ListenAndServe(":8080", nil)
}

func Calc(w http.ResponseWriter, r *http.Request) {
	var request Request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		calculation.ResponseError(w, "Expression is not valid", http.StatusUnprocessableEntity)
		return
	}

	calcResult, err := calculation.Calc(request.Expression)

	if err == nil {
		result := Result{Result: calcResult}

		resultJson, err := json.Marshal(result)

		if err != nil {
			calculation.ResponseError(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(resultJson)
		return
	}

	switch err.Error() {
	case "Expression is not valid":
		calculation.ResponseError(w, "Expression is not valid", http.StatusUnprocessableEntity)
		return
	case "Internal server error":
		calculation.ResponseError(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
