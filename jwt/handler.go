package jwt

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	RequestErr     = "Error in request"
	CreateTokenErr = "Error while Create the token"
	Unauthorized   = "Unauthorized access to this resource"
	Authorized     = "Gained access to protected resource"
)

type UserCredentials struct {
	UserName string `json:"username"`
	ID       int64  `json:"id"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user UserCredentials
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		writeStatus(w, http.StatusForbidden, RequestErr)
		return
	}

	responseToken, err := CreateToken(user.UserName, user.ID)
	if err != nil {
		writeStatus(w, http.StatusInternalServerError, CreateTokenErr)
		return
	}
	JsonResponse(responseToken, w)
}

// ValidateTokenHandler
func ValidateTokenHandler(w http.ResponseWriter, r *http.Request) {
	//token, err := ParseToken(r.Header.Get("Authorization")[7:])
	// 过期、无法识别等其他错误
	if !CheckToken(r.Header.Get("Authorization")[7:]) {
		writeUnauthorizedStatus(w, Unauthorized)
		return
	}
	// 满足权限
	writeStatus(w, http.StatusOK, Authorized)
}

// 写入状态
func writeStatus(w http.ResponseWriter, status int, reason string) {
	w.WriteHeader(status)
	fmt.Fprint(w, reason)
}

// writeUnauthorizedStatus 无授权状态
func writeUnauthorizedStatus(w http.ResponseWriter, reason string) {
	writeStatus(w, http.StatusUnauthorized, reason)
}

// JsonResponse
func JsonResponse(response interface{}, w http.ResponseWriter) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func StartServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", LoginHandler)
	mux.HandleFunc("/auth", ValidateTokenHandler)

	log.Println("Now jwt is listening...")
	http.ListenAndServe(":8080", mux)
}
