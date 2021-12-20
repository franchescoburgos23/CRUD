package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func ConnetionBD() (conexion *sql.DB) {

	Driver := "mysql"
	User := "root"
	Password := ""
	Name := "sistema"

	conexion, err := sql.Open(Driver, User+":"+Password+"@tcp(127.0.0.1)/"+Name)

	if err != nil {
		panic(err.Error())
	}

	return conexion

}

var templt = template.Must(template.ParseGlob("templates/*"))

func Home(w http.ResponseWriter, r *http.Request) {

	establishconnexion := ConnetionBD()
	insertconnexion, err := establishconnexion.Prepare("INSERT INTO Empleados(nombre,correo) VALUES ('Juan','correo@gmail.com')")

	if err != nil {
		panic(err.Error())
	}

	insertconnexion.Exec()

	templt.ExecuteTemplate(w, "home", nil)

}

func Create(w http.ResponseWriter, r *http.Request) {

	templt.ExecuteTemplate(w, "creater", nil)

}

func Insert(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		name := r.FormValue("Name")
		email := r.FormValue("Email")

		establishconnexion := ConnetionBD()
		insertconnexion, err := establishconnexion.Prepare("INSERT INTO Empleados(nombre,correo) VALUES (?,?)")

		if err != nil {
			panic(err.Error())
		}

		insertconnexion.Exec(name, email)

		http.Redirect(w, r, "/", 301)

	}
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/insert", Insert)

	log.Println("Server Start")

	http.ListenAndServe(":8080", nil)

}
