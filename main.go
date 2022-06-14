package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"
	_"github.com/go-sql-driver/mysql"
)

func conexionBD() (conexion *sql.DB) {
	Driver := "mysql"
	Usuario := "root"
	Contrasenia := ""
	Nombre := "sistema-empleados"

	conexion, err := sql.Open(Driver, Usuario+":"+Contrasenia+"@tcp(127.0.0.1)/"+Nombre)
	if err != nil {
		panic(err.Error())
	}

	return conexion
}

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {
	http.HandleFunc("/", Inicio)
	http.HandleFunc("/crear", Crear)
	http.HandleFunc("/insertar", Insertar)

	log.Println("Servidor corriendo...")
	http.ListenAndServe(":8080", nil)
}

type Empleadxs struct {
	Id int
	Nombre string
	Correo string
}


func Inicio (w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola Karen")

	conexionEstablecida := conexionBD()

	registros, err := conexionEstablecida.Query("SELECT * FROM empleadxs")

	if err != nil {
		panic(err.Error())
	}


	empleadx := Empleadxs{}
	arregloEmpleadxs:=[]Empleadxs{}

	for registros.Next() {
		var id int
		var nombre, correo string
		err= registros.Scan(&id,&nombre,&correo)

		if err != nil {
			panic(err.Error())
		}

		empleadx.Id=id
		empleadx.Nombre=nombre 
		empleadx.Correo=correo

		arregloEmpleadxs=append(arregloEmpleadxs, empleadx)

	}

	fmt.Println(arregloEmpleadxs)
	

	plantillas.ExecuteTemplate(w, "inicio", nil)

}
func Crear (w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "crear", nil)

}

func Insertar(w http.ResponseWriter, r *http.Request) {
	if r.Method=="POST" {

		nombre:= r.FormValue("nombre")
		correo:= r.FormValue("correo")

		conexionEstablecida := conexionBD()

		insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO empleadxs(nombre, correo) VALUES(?,?)")

		if err != nil {
			panic(err.Error())
		}

		insertarRegistros.Exec(nombre, correo)

		http.Redirect(w,r,"/",301)

	}

}


