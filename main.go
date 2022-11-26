package main

import (
	"database/sql"
	//"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

func conexionDB() (conexion *sql.DB) {
	Driver := "mysql"
	User := "root"
	Password := ""
	Name := "agenciaf"

	conexion, err := sql.Open(Driver, User+":"+Password+"@tcp(127.0.0.1)/"+Name)
	if err != nil {
		panic(err.Error())
	}
	return conexion
}

var templates = template.Must(template.ParseGlob("templates/*"))

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/indexHotel", indexHotel)
	http.HandleFunc("/createHotel", createHotel)
	http.HandleFunc("/insertHotel", insertHotel)
	http.HandleFunc("/eliminarHotel", eliminarHotel)
	http.HandleFunc("/editarHotel", editarHotel)
	http.HandleFunc("/updateHotel", updateHotel)
	http.HandleFunc("/indexSucursal", indexSucursal)
	http.HandleFunc("/createSucursal", createSucursal)
	http.HandleFunc("/insertSucursal", insertSucursal)
	http.HandleFunc("/eliminarSucursal", eliminarSucursal)
	http.HandleFunc("/editarSucursal", editarSucursal)
	http.HandleFunc("/updateSucursal", updateSucursal)
	http.HandleFunc("/indexVuelo", indexVuelo)
	http.HandleFunc("/createVuelo", createVuelo)
	http.HandleFunc("/insertVuelo", insertVuelo)
	http.HandleFunc("/eliminarVuelo", eliminarVuelo)
	http.HandleFunc("/editarVuelo", editarVuelo)
	http.HandleFunc("/updateVuelo", updateVuelo)
	http.HandleFunc("/indexCliente", indexCliente)
	http.HandleFunc("/createCliente", createCliente)
	http.HandleFunc("/insertCliente", insertCliente)
	http.HandleFunc("/eliminarCliente", eliminarCliente)
	http.HandleFunc("/editarCliente", editarCliente)
	http.HandleFunc("/updateCliente", updateCliente)
	log.Println("servidor corriendo")
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index", nil)
}

type Hotel struct {
	IdHotel      int
	Nombre       string
	Telefono     string
	NumeroPlazas int
	Calle        string
	Colonia      string
	Cp           string
	Ciudad       string
	Estado       string
	Pais         string
	Estatus      int
}

func indexHotel(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola mundo")
	conexionEst := conexionDB()
	registros, err := conexionEst.Query("SELECT * FROM hotel WHERE estatus=1")
	if err != nil {
		panic(err.Error())
	}
	hotel := Hotel{}
	arrayHotel := []Hotel{}
	for registros.Next() {
		var idHotel, numeroPlazas, estatus int
		var nombre, telefono, calle, colonia, cp, ciudad, estado, pais string
		err = registros.Scan(&idHotel, &telefono, &nombre, &calle, &colonia, &cp, &ciudad, &estado, &pais, &numeroPlazas, &estatus)
		if err != nil {
			panic(err.Error())
		}
		hotel.IdHotel = idHotel
		hotel.Nombre = nombre
		hotel.Telefono = telefono
		hotel.NumeroPlazas = numeroPlazas
		hotel.Calle = calle
		hotel.Colonia = colonia
		hotel.Cp = cp
		hotel.Ciudad = ciudad
		hotel.Estado = estado
		hotel.Pais = pais
		hotel.Estatus = estatus
		arrayHotel = append(arrayHotel, hotel)
	}

	templates.ExecuteTemplate(w, "indexHotel", arrayHotel)
}

func createHotel(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "createHotel", nil)
}

func insertHotel(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nombre := r.FormValue("nombre")
		numeroPlazas := r.FormValue("numeroPlazas")
		telefono := r.FormValue("telefono")
		calle := r.FormValue("calle")
		colonia := r.FormValue("colonia")
		cp := r.FormValue("cp")
		ciudad := r.FormValue("ciudad")
		estado := r.FormValue("estado")
		pais := r.FormValue("pais")

		conexionEst := conexionDB()
		createRegister, err := conexionEst.Prepare("INSERT INTO hotel(nombre, telefono, numeroPlazas, calle, colonia, cp, ciudad, estado, pais) VALUES(?,?,?,?,?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		createRegister.Exec(nombre, telefono, numeroPlazas, calle, colonia, cp, ciudad, estado, pais)

		http.Redirect(w, r, "/indexHotel", 301)
	}
}

func eliminarHotel(w http.ResponseWriter, r *http.Request) {
	idHotel := r.URL.Query().Get("id")

	conexionEst := conexionDB()
	deleteRegister, err := conexionEst.Prepare("UPDATE hotel SET estatus=0 WHERE idHotel=?")
	if err != nil {
		panic(err.Error())
	}
	deleteRegister.Exec(idHotel)
	http.Redirect(w, r, "/indexHotel", 301)
}

func editarHotel(w http.ResponseWriter, r *http.Request) {
	idHotel := r.URL.Query().Get("id")

	conexionEst := conexionDB()
	registro, err := conexionEst.Query("SELECT * FROM hotel WHERE idHotel=?", idHotel)
	if err != nil {
		panic(err.Error())
	}
	hotel := Hotel{}
	for registro.Next() {
		var idHotel, numeroPlazas, estatus int
		var nombre, telefono, calle, colonia, cp, ciudad, estado, pais string
		err = registro.Scan(&idHotel, &telefono, &nombre, &calle, &colonia, &cp, &ciudad, &estado, &pais, &numeroPlazas, &estatus)
		if err != nil {
			panic(err.Error())
		}
		hotel.IdHotel = idHotel
		hotel.Nombre = nombre
		hotel.Telefono = telefono
		hotel.NumeroPlazas = numeroPlazas
		hotel.Calle = calle
		hotel.Colonia = colonia
		hotel.Cp = cp
		hotel.Ciudad = ciudad
		hotel.Estado = estado
		hotel.Pais = pais
		hotel.Estatus = estatus
	}
	templates.ExecuteTemplate(w, "editarHotel", hotel)
}

func updateHotel(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		idHotel := r.FormValue("idHotel")
		nombre := r.FormValue("nombre")
		numeroPlazas := r.FormValue("numeroPlazas")
		telefono := r.FormValue("telefono")
		calle := r.FormValue("calle")
		colonia := r.FormValue("colonia")
		cp := r.FormValue("cp")
		ciudad := r.FormValue("ciudad")
		estado := r.FormValue("estado")
		pais := r.FormValue("pais")
		conexionEst := conexionDB()
		updateRegister, err := conexionEst.Prepare("UPDATE hotel SET nombre=?, telefono=?, numeroPlazas=?, calle=?, colonia=?, cp=?, ciudad=?, estado=?, pais=? WHERE idHotel=?")
		if err != nil {
			panic(err.Error())
		}
		updateRegister.Exec(nombre, telefono, numeroPlazas, calle, colonia, cp, ciudad, estado, pais, idHotel)

		http.Redirect(w, r, "/indexHotel", 301)
	}

}

type Sucursal struct {
	IdSucursal   int
	Nombre       string
	Telefono     string
	NumeroPlazas int
	Calle        string
	Colonia      string
	Cp           string
	Estatus      int
}

func indexSucursal(w http.ResponseWriter, r *http.Request) {
	conexionEst := conexionDB()
	registros, err := conexionEst.Query("SELECT * FROM sucursal WHERE estatus=1")
	if err != nil {
		panic(err.Error())
	}
	sucursal := Sucursal{}
	arraySucursal := []Sucursal{}
	for registros.Next() {
		var idSucursal, numeroPlazas, estatus int
		var nombre, telefono, calle, colonia, cp string
		err = registros.Scan(&idSucursal, &nombre, &telefono, &calle, &colonia, &cp, &numeroPlazas, &estatus)
		if err != nil {
			panic(err.Error())
		}
		sucursal.IdSucursal = idSucursal
		sucursal.Nombre = nombre
		sucursal.Telefono = telefono
		sucursal.NumeroPlazas = numeroPlazas
		sucursal.Calle = calle
		sucursal.Colonia = colonia
		sucursal.Cp = cp
		sucursal.Estatus = estatus
		arraySucursal = append(arraySucursal, sucursal)
	}

	templates.ExecuteTemplate(w, "indexSucursal", arraySucursal)
}

func createSucursal(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "createSucursal", nil)
}

func insertSucursal(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nombre := r.FormValue("nombre")
		numeroPlazas := r.FormValue("numeroPlazas")
		telefono := r.FormValue("telefono")
		calle := r.FormValue("calle")
		colonia := r.FormValue("colonia")
		cp := r.FormValue("cp")

		conexionEst := conexionDB()
		createRegister, err := conexionEst.Prepare("INSERT INTO sucursal(nombre, telefono, numeroPlazas, calle, colonia, cp) VALUES(?,?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		createRegister.Exec(nombre, telefono, numeroPlazas, calle, colonia, cp)

		http.Redirect(w, r, "/indexSucursal", 301)
	}
}

func eliminarSucursal(w http.ResponseWriter, r *http.Request) {
	idSucursal := r.URL.Query().Get("id")

	conexionEst := conexionDB()
	deleteRegister, err := conexionEst.Prepare("UPDATE sucursal SET estatus=0 WHERE idSucursal=?")
	if err != nil {
		panic(err.Error())
	}
	deleteRegister.Exec(idSucursal)
	http.Redirect(w, r, "/indexSucursal", 301)
}

func editarSucursal(w http.ResponseWriter, r *http.Request) {
	idSucursal := r.URL.Query().Get("id")

	conexionEst := conexionDB()
	registro, err := conexionEst.Query("SELECT * FROM sucursal WHERE idSucursal=?", idSucursal)
	if err != nil {
		panic(err.Error())
	}
	sucursal := Sucursal{}
	for registro.Next() {
		var idSucursal, numeroPlazas, estatus int
		var nombre, telefono, calle, colonia, cp string
		err = registro.Scan(&idSucursal, &nombre, &telefono, &calle, &colonia, &cp, &numeroPlazas, &estatus)
		if err != nil {
			panic(err.Error())
		}
		sucursal.IdSucursal = idSucursal
		sucursal.Nombre = nombre
		sucursal.Telefono = telefono
		sucursal.NumeroPlazas = numeroPlazas
		sucursal.Calle = calle
		sucursal.Colonia = colonia
		sucursal.Cp = cp
		sucursal.Estatus = estatus
	}
	templates.ExecuteTemplate(w, "editarSucursal", sucursal)
}

func updateSucursal(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		idSucursal := r.FormValue("idSucursal")
		nombre := r.FormValue("nombre")
		numeroPlazas := r.FormValue("numeroPlazas")
		telefono := r.FormValue("telefono")
		calle := r.FormValue("calle")
		colonia := r.FormValue("colonia")
		cp := r.FormValue("cp")
		conexionEst := conexionDB()
		updateRegister, err := conexionEst.Prepare("UPDATE sucursal SET nombre=?, telefono=?, numeroPlazas=?, calle=?, colonia=?, cp=? WHERE idSucursal=?")
		if err != nil {
			panic(err.Error())
		}
		updateRegister.Exec(nombre, telefono, numeroPlazas, calle, colonia, cp, idSucursal)

		http.Redirect(w, r, "/indexSucursal", 301)
	}
}

type Vuelo struct {
	IdVuelo       int
	Fecha         string
	Hora          string
	PlazasTotales int
	CiudadOrigen  string
	EstadoOrigen  string
	PaisOrigen    string
	CiudadDestino string
	EstadoDestino string
	PaisDestino   string
	Estatus       int
}

func indexVuelo(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola mundo")
	conexionEst := conexionDB()
	registros, err := conexionEst.Query("SELECT * FROM vuelo WHERE estatus=1")
	if err != nil {
		panic(err.Error())
	}
	vuelo := Vuelo{}
	arrayVuelo := []Vuelo{}
	for registros.Next() {
		var idVuelo, plazasTotales, estatus int
		var fecha, hora, ciudadOrigen, estadoOrigen, paisOrigen, ciudadDestino, estadoDestino, paisDestino string
		err = registros.Scan(&idVuelo, &ciudadOrigen, &estadoOrigen, &paisOrigen, &ciudadDestino, &estadoDestino, &paisDestino, &plazasTotales, &fecha, &hora, &estatus)
		if err != nil {
			panic(err.Error())
		}
		vuelo.IdVuelo = idVuelo
		vuelo.Fecha = fecha
		vuelo.Hora = hora
		vuelo.PlazasTotales = plazasTotales
		vuelo.CiudadOrigen = ciudadOrigen
		vuelo.EstadoOrigen = estadoOrigen
		vuelo.PaisOrigen = paisOrigen
		vuelo.CiudadDestino = ciudadDestino
		vuelo.EstadoDestino = estadoDestino
		vuelo.PaisDestino = paisDestino
		vuelo.Estatus = estatus
		arrayVuelo = append(arrayVuelo, vuelo)
	}

	templates.ExecuteTemplate(w, "indexVuelo", arrayVuelo)
}

func createVuelo(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "createVuelo", nil)
}

func insertVuelo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fecha := r.FormValue("fecha")
		plazasTotales := r.FormValue("plazasTotales")
		hora := r.FormValue("hora")
		ciudadOrigen := r.FormValue("ciudadOrigen")
		estadoOrigen := r.FormValue("estadoOrigen")
		paisOrigen := r.FormValue("paisOrigen")
		ciudadDestino := r.FormValue("ciudadDestino")
		estadoDestino := r.FormValue("estadoDestino")
		paisDestino := r.FormValue("paisDestino")

		conexionEst := conexionDB()
		createRegister, err := conexionEst.Prepare("INSERT INTO vuelo(fecha, hora, plazasTotales, ciudadOrigen, estadoOrigen, paisOrigen, ciudadDestino, estadoDestino, paisDestino) VALUES(?,?,?,?,?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		createRegister.Exec(fecha, hora, plazasTotales, ciudadOrigen, estadoOrigen, paisOrigen, ciudadDestino, estadoDestino, paisDestino)

		http.Redirect(w, r, "/indexVuelo", 301)
	}
}

func eliminarVuelo(w http.ResponseWriter, r *http.Request) {
	idVuelo := r.URL.Query().Get("id")

	conexionEst := conexionDB()
	deleteRegister, err := conexionEst.Prepare("UPDATE vuelo SET estatus=0 WHERE idVuelo=?")
	if err != nil {
		panic(err.Error())
	}
	deleteRegister.Exec(idVuelo)
	http.Redirect(w, r, "/indexVuelo", 301)
}

func editarVuelo(w http.ResponseWriter, r *http.Request) {
	idVuelo := r.URL.Query().Get("id")

	conexionEst := conexionDB()
	registro, err := conexionEst.Query("SELECT * FROM vuelo WHERE idVuelo=?", idVuelo)
	if err != nil {
		panic(err.Error())
	}
	vuelo := Vuelo{}
	for registro.Next() {
		var idVuelo, plazasTotales, estatus int
		var fecha, hora, ciudadOrigen, estadoOrigen, paisOrigen, ciudadDestino, estadoDestino, paisDestino string
		err = registro.Scan(&idVuelo, &ciudadOrigen, &estadoOrigen, &paisOrigen, &ciudadDestino, &estadoDestino, &paisDestino, &plazasTotales, &fecha, &hora, &estatus)
		if err != nil {
			panic(err.Error())
		}
		vuelo.IdVuelo = idVuelo
		vuelo.Fecha = fecha
		vuelo.Hora = hora
		vuelo.PlazasTotales = plazasTotales
		vuelo.CiudadOrigen = ciudadOrigen
		vuelo.EstadoOrigen = estadoOrigen
		vuelo.PaisOrigen = paisOrigen
		vuelo.CiudadDestino = ciudadDestino
		vuelo.EstadoDestino = estadoDestino
		vuelo.PaisDestino = paisDestino
		vuelo.Estatus = estatus
	}
	templates.ExecuteTemplate(w, "editarVuelo", vuelo)
}

func updateVuelo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		idVuelo := r.FormValue("idVuelo")
		fecha := r.FormValue("fecha")
		plazasTotales := r.FormValue("plazasTotales")
		hora := r.FormValue("hora")
		ciudadOrigen := r.FormValue("ciudadOrigen")
		estadoOrigen := r.FormValue("estadoOrigen")
		paisOrigen := r.FormValue("paisOrigen")
		ciudadDestino := r.FormValue("ciudadDestino")
		estadoDestino := r.FormValue("estadoDestino")
		paisDestino := r.FormValue("paisDestino")
		conexionEst := conexionDB()
		updateRegister, err := conexionEst.Prepare("UPDATE vuelo SET fecha=?, hora=?, plazasTotales=?, ciudadOrigen=?, estadoOrigen=?, paisOrigen=?, ciudadDestino=?, estadoDestino=?, paisDestino=? WHERE idVuelo=?")
		if err != nil {
			panic(err.Error())
		}
		updateRegister.Exec(fecha, hora, plazasTotales, ciudadOrigen, estadoOrigen, paisOrigen, ciudadDestino, estadoDestino, paisDestino, idVuelo)

		http.Redirect(w, r, "/indexVuelo", 301)
	}

}

type Cliente struct {
	IdCliente       int
	Nombre          string
	ApellidoPaterno string
	ApellidoMaterno string
	Telefono        string
	Calle           string
	Colonia         string
	Cp              string
	IdHotel         int
	RegimenHotel    string
	IdSucursal      int
	IdVuelo         int
	ClaseVuelo      string
	Estatus         int
}

func indexCliente(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola mundo")
	conexionEst := conexionDB()
	registros, err := conexionEst.Query("SELECT * FROM cliente WHERE estatus=1")
	if err != nil {
		panic(err.Error())
	}
	cliente := Cliente{}
	arrayCliente := []Cliente{}
	for registros.Next() {
		var idCliente, idHotel, idVuelo, idSucursal, estatus int
		var nombre, telefono, calle, colonia, cp, apellidoPaterno, apellidoMaterno, regimenHotel, claseVuelo string
		err = registros.Scan(&idCliente, &nombre, &apellidoPaterno, &apellidoMaterno, &telefono, &calle, &colonia, &cp, &idHotel, &regimenHotel, &idSucursal, &idVuelo, &claseVuelo, &estatus)
		if err != nil {
			panic(err.Error())
		}
		cliente.IdCliente = idCliente
		cliente.Nombre = nombre
		cliente.Telefono = telefono
		cliente.IdHotel = idHotel
		cliente.Calle = calle
		cliente.Colonia = colonia
		cliente.Cp = cp
		cliente.ApellidoPaterno = apellidoPaterno
		cliente.ApellidoMaterno = apellidoMaterno
		cliente.RegimenHotel = regimenHotel
		cliente.IdSucursal = idSucursal
		cliente.IdVuelo = idVuelo
		cliente.ClaseVuelo = claseVuelo
		cliente.Estatus = estatus
		arrayCliente = append(arrayCliente, cliente)
	}

	templates.ExecuteTemplate(w, "indexCliente", arrayCliente)
}

func createCliente(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "createCliente", nil)
}

func insertCliente(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nombre := r.FormValue("nombre")
		idHotel := r.FormValue("idHotel")
		telefono := r.FormValue("telefono")
		calle := r.FormValue("calle")
		colonia := r.FormValue("colonia")
		cp := r.FormValue("cp")
		apellidoPaterno := r.FormValue("apellidoPaterno")
		apellidoMaterno := r.FormValue("apellidoMaterno")
		regimenHotel := r.FormValue("regimenHotel")
		idVuelo := r.FormValue("idVuelo")
		idSucursal := r.FormValue("idSucursal")
		claseVuelo := r.FormValue("claseVuelo")

		conexionEst := conexionDB()
		createRegister, err := conexionEst.Prepare("INSERT INTO cliente(nombre, telefono, idHotel, calle, colonia, cp, apellidoPaterno, apellidoMaterno, regimenHotel, idVuelo, idSucursal, claseVuelo) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		createRegister.Exec(nombre, telefono, idHotel, calle, colonia, cp, apellidoPaterno, apellidoMaterno, regimenHotel, idVuelo, idSucursal, claseVuelo)

		http.Redirect(w, r, "/indexCliente", 301)
	}
}

func eliminarCliente(w http.ResponseWriter, r *http.Request) {
	idCliente := r.URL.Query().Get("id")

	conexionEst := conexionDB()
	deleteRegister, err := conexionEst.Prepare("UPDATE cliente SET estatus=0 WHERE idCliente=?")
	if err != nil {
		panic(err.Error())
	}
	deleteRegister.Exec(idCliente)
	http.Redirect(w, r, "/indexCliente", 301)
}

func editarCliente(w http.ResponseWriter, r *http.Request) {
	idCliente := r.URL.Query().Get("id")

	conexionEst := conexionDB()
	registro, err := conexionEst.Query("SELECT * FROM cliente WHERE idCliente=?", idCliente)
	if err != nil {
		panic(err.Error())
	}
	cliente := Cliente{}
	for registro.Next() {
		var idCliente, idHotel, idSucursal, idVuelo, estatus int
		var nombre, telefono, calle, colonia, cp, apellidoPaterno, apellidoMaterno, regimenHotel, claseVuelo string
		err = registro.Scan(&idCliente, &nombre, &apellidoPaterno, &apellidoMaterno, &telefono, &calle, &colonia, &cp, &idHotel, &regimenHotel, &idSucursal, &idVuelo, &claseVuelo, &estatus)
		if err != nil {
			panic(err.Error())
		}
		cliente.IdCliente = idCliente
		cliente.Nombre = nombre
		cliente.Telefono = telefono
		cliente.IdHotel = idHotel
		cliente.Calle = calle
		cliente.Colonia = colonia
		cliente.Cp = cp
		cliente.ApellidoPaterno = apellidoPaterno
		cliente.ApellidoMaterno = apellidoMaterno
		cliente.RegimenHotel = regimenHotel
		cliente.IdVuelo = idVuelo
		cliente.IdSucursal = idSucursal
		cliente.ClaseVuelo = claseVuelo
		cliente.Estatus = estatus
	}
	templates.ExecuteTemplate(w, "editarCliente", cliente)
}

func updateCliente(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		idCliente := r.FormValue("idCliente")
		nombre := r.FormValue("nombre")
		idHotel := r.FormValue("idHotel")
		telefono := r.FormValue("telefono")
		calle := r.FormValue("calle")
		colonia := r.FormValue("colonia")
		cp := r.FormValue("cp")
		apellidoPaterno := r.FormValue("apellidoPaterno")
		apellidoMaterno := r.FormValue("apellidoMaterno")
		regimenHotel := r.FormValue("regimenHotel")
		idVuelo := r.FormValue("idVuelo")
		idSucursal := r.FormValue("idSucursal")
		claseVuelo := r.FormValue("claseVuelo")
		conexionEst := conexionDB()
		updateRegister, err := conexionEst.Prepare("UPDATE cliente SET nombre=?, telefono=?, idHotel=?, calle=?, colonia=?, cp=?, apellidoPaterno=?, apellidoMaterno=?, regimenHotel=?, idVuelo=?, idSucursal=?, claseVuelo=? WHERE idCliente=?")
		if err != nil {
			panic(err.Error())
		}
		updateRegister.Exec(nombre, telefono, idHotel, calle, colonia, cp, apellidoPaterno, apellidoMaterno, regimenHotel, idVuelo, idSucursal, claseVuelo, idCliente)

		http.Redirect(w, r, "/indexCliente", 301)
	}

}
