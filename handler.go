package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (S *Server) SetRoutes() {
	Mux := S.Handler.(*http.ServeMux)
	Mux.Handle("/check", pingHandler())
	Mux.Handle("/post", postHandler())
}

func pingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Check Completed!!\n"))
	}
}

func postHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data, _ := io.ReadAll(r.Body)

		var cred Creds

		err := json.Unmarshal(data, &cred)
		if err != nil {
			Logger.Error(fmt.Sprintf("Error in Unmarshaling: %s", err))
			return
		}

		w.Write([]byte("Post Completed!\n"))

	}

}
