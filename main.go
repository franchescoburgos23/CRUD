package main

import (
	"database/sql"
	"fmt"
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

type Employer struct {
	Id    int
	Name  string
	Email string
}

func Home(w http.ResponseWriter, r *http.Request) {

	establishconnexion := ConnetionBD()
	records, err := establishconnexion.Query("SELECT *FROM Empleados")

	if err != nil {
		panic(err.Error())
	}

	employer := Employer{}
	aremployer := []Employer{}
	for records.Next() {
		var id int
		var name, email string
		err = records.Scan(&id, &name, &email)
		if err != nil {
			panic(err.Error())
		}
		employer.Id = id
		employer.Name = name
		employer.Email = email

		aremployer = append(aremployer, employer)
	}

	//fmt.Println(aremployer)

	templt.ExecuteTemplate(w, "home", aremployer)

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

func Delete(w http.ResponseWriter, r *http.Request) {

	idemployer := r.URL.Query().Get("id")
	fmt.Println(idemployer)

	establishconnexion := ConnetionBD()
	deleterecords, err := establishconnexion.Prepare("DELETE * FROM Empleados WHERE id= ?")

	if err != nil {
		panic(err.Error())
	}

	deleterecords.Exec(idemployer)

	http.Redirect(w, r, "/", 301)

}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/delete", Delete)

	log.Println("Server Start")

	http.ListenAndServe(":8080", nil)

}
