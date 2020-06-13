package server

import (
	"bytes"
	"demo/jwt"
	"demo/user"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)



const (
	jwtLoginURL = "http://localhost:8080/login"
	jwtAuthURL  = "http://localhost:8080/auth"
)

const UnauthorizedStatus = "401 Unauthorized"

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	/*
		进行其他操作，比如调用微信的auth.code2Session
		换取用户唯一标识OpenID和会话密钥session_key
	*/

	// 根据用户名和id生成用户
	myUser := &user.User{ID: 123456, Name: "Mike"}

	// 检测用户是否已经注册
	if !myUser.IsExist() {
		/*
			用户不存在，就返回进行注册
		*/
		return
	}
	// 用户存在，就鉴权
	token, _ := getToken(myUser.Name, myUser.ID)
	fmt.Println(token)
	/*
		进行其他操作，比如将Token绑定用户等
	*/
}

func getToken(name string, id int64) (token string, err error) {

	user := jwt.UserCredentials{UserName: name, ID: id}
	s, _ := json.Marshal(user)

	contentType := "application/json"
	resp, err := http.Post(jwtLoginURL, contentType, bytes.NewBuffer(s))
	if err != nil {
		log.Println(err)
		return
	}
	var jToken jwt.Token
	if err = json.NewDecoder(resp.Body).Decode(&jToken); err != nil {
		log.Println(err)
		return
	}
	token = jToken.Token
	return
}

func ResourceHandler(w http.ResponseWriter, r *http.Request) {
	// 在用户的请求中解析出Token
	var token jwt.Token
	if err := json.NewDecoder(r.Body).Decode(&token); err != nil {
		log.Println(err)
	}
	fmt.Println("res:", token)

	// Token鉴权
	client := &http.Client{}
	req, _ := http.NewRequest("GET", jwtAuthURL, nil)
	value := "Bearer " + token.Token
	req.Header.Add("Authorization", value)
	resp, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Err in request check token")
		return
	}
	// 直接从状态判断是否有权限
	if resp.Status == UnauthorizedStatus {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Is unauthorized")
		return
	}
	/*
		有权限就获取请求资源并放回
	 */
	return
}

func RegisterHandler(w http.ResponseWriter, r *http.Request)  {

}


func StartServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", LoginHandler)
	mux.HandleFunc("/resource", ResourceHandler)
	mux.HandleFunc("/register", RegisterHandler)
	log.Println("Now listening...")
	http.ListenAndServe(":9090", nil)
}
