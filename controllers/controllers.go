package controllers

import (
	"github.com/gin-gonic/gin"
)

// func methodChecker(w http.ResponseWriter, r *http.Request, allowedMethod string) {
// 	if r.Method != allowedMethod {
// 		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
// 		return
// 	}
// }

func errorMsg(err error) gin.H {
	return gin.H{"error": err.Error()}
}
