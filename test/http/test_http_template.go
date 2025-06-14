package http01

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type User struct {
	ID       int
	Name     string
	Email    string
	IsActive bool
}

func Http_template_main() {
	fmt.Println("Http_template_main......start")
	defer fmt.Println("Http_template_main......end")

	workDir, err := os.Getwd()
	if err != nil {
		log.Fatalln("os.Getwd() ..err=", err)
	} else {
		log.Println("workDir=", workDir)
	}

	// http://127.0.0.1:8080/users
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		// 模拟数据
		users := []User{
			{ID: 1, Name: "Alice", Email: "alice@example.com", IsActive: true},
			{ID: 2, Name: "Bob", Email: "bob@example.com", IsActive: false},
		}

		// 解析模板
		tmpl := template.Must(template.ParseFiles(workDir + "/test/http/templates/users.html"))
		// tmpl := template.Must(template.ParseFiles("ginchat/test/http/templates/users.html"))

		// 执行模板，传递数据
		tmpl.Execute(w, users)
	})

	http.ListenAndServe(":8080", nil)
}
