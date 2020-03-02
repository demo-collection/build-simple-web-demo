package core

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)

	if r.Method == "GET" {
		t, _ := template.ParseFiles("../../login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		//请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}
