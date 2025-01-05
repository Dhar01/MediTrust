package controllers

import "net/http"

func methodChecker(w http.ResponseWriter, r *http.Request, allowedMethod string) {
	if r.Method != allowedMethod {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}
