// dictionary/routes.go
package dictionary

import "github.com/gorilla/mux"

// RegisterHandlers enregistre les gestionnaires de routes avec le routeur Gorilla Mux.
func RegisterHandlers(r *mux.Router) {
    r.HandleFunc("/add", HandleAdd).Methods("POST")
    r.HandleFunc("/define/{word}", HandleDefine).Methods("GET")
    r.HandleFunc("/remove/{word}", HandleRemove).Methods("DELETE")
    r.HandleFunc("/list", HandleList).Methods("GET")
}
