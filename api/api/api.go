package api

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func JsonResponse(w http.ResponseWriter, data interface{}) {
	m, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(m))
}
