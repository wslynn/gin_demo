package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http serve failed, err:%v", err)
		return
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 读取模板文件
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("ParseFiles failed, err:%v", err)
		return
	}
	// 渲染模板, .表示当前对象
	err = t.Execute(w, "小王")
	if err != nil {
		fmt.Println("Execute failed, err:%v", err)
		return
	}
}

func tmplDemo(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	user := UserInfo{Name: "小王子", Gender: "男", Age: 18}
	tmpl.Execute(w, user)
}
