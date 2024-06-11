package main

import "net/http"

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, struct{}{})
}

func handleErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusBadRequest, "Something went wrong")
}
