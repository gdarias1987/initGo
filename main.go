package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	// hay que referenciar desde src ->
	"github.com/gdarias1987/DEV/initGo/utils"
)

// Afuera del main solo se pueden declarar como en JS
var test = "test"

// Declaracion funcion anonima como JS
func sumador() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// STRUCT EJEMPLO
type Usuario struct {
	id     int
	email  string
	status bool
	// firstName, lastName, gender string -> en una sola linea
}

// GETTER (VALUE RECEIVER)
func (user Usuario) toJSON() string {
	aux := "{\"id\":\"" + strconv.Itoa(user.id) + "\",\"id\":\"" + user.email + "\",\"id\":\"" + strconv.FormatBool(user.status) + "\"}"
	return aux
}

// SETTER (POINTER RECEIVER)
func (user *Usuario) deactivateUser() {
	user.status = false
}

func (user *Usuario) activateUser() {
	user.status = true
}

//DECLARACION INTERFACE
type Forma interface {
	area() float64
}

type Circulo struct {
	x, y, radio float64
}

type Rectangulo struct {
	ancho, alto float64
}

func (c Circulo) area() float64 {
	return (math.Pi * c.radio * c.radio)
}

func (r Rectangulo) area() float64 {
	return r.ancho * r.alto
}

func getArea(f Forma) float64 {
	return f.area()
}

// IMPLEMENTACION WEB

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Willi3")
}

func main() {

	// var name string = "TES1T"
	// name := "name"
	// var age int = 13

	name, age := "name", 14

	switch name {
	case "name":
		fmt.Println("-----> Ejemplo Switch")
	default:
		fmt.Println("Ejemplo Default")
	}

	fmt.Printf("%T\n", age)
	fmt.Println(name, test)
	fmt.Println(math.Floor(2.4))
	fmt.Println(utils.Reverse("olleh"))

	fmt.Println("----------------------------------------------------------\n")
	fmt.Println("Array\n")
	// // Array fixo
	// var arrTest [2]string
	// arrTest[0] = "PRIMERO"
	// arrTest[1] = "SEGUNDO"

	// arrTest := [2]string{"PRIMERO", "SEGUNDO"}

	fmt.Println("----------------------------------------------------------\n")
	fmt.Println("Slice\n")
	fmt.Println("[]string{\"PRIMERO\", \"SEGUNDO\", \"TERCERO\", \"CUARTO\"}\n")

	// Sin index en la declaracion -> Slices
	arrTest := []string{"PRIMERO", "SEGUNDO", "TERCERO", "CUARTO"}

	fmt.Println(arrTest)
	fmt.Println(arrTest[1:2])

	fmt.Println("----------------------------------------------------------\n")
	fmt.Println("MAP\n")
	fmt.Println("user := map[string]string{\"id\": \"1\", \"email\": \"g@g.com\", \"toDelete\": \"delete\"}\n")

	// Map forma 1
	// user := make(map[string]string)

	// user["id"] = "1"
	// user["email"] = "g@g.com"
	// user["toDelete"] = "detele"

	// Map form 2
	user := map[string]string{"id": "1", "email": "g@g.com", "toDelete": "delete"}

	fmt.Println(user, "id: ", user["id"], " - Lenght: ", len(user))
	delete(user, "toDelete")
	fmt.Println(user, "id: ", user["id"], " - Lenght: ", len(user))

	fmt.Println("----------------------------------------------------------\n")
	fmt.Println("Range\n")
	fmt.Println("ids := []int{23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33}\n")

	ids := []int{23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33}

	// for _, id := range ids {
	for i, id := range ids {
		fmt.Printf("i:[%d] - id: %d \n", i, id)
	}
	fmt.Println("\n")
	fmt.Println("----------------------------------------------------------\n")
	fmt.Println("Range with Maps\n")

	for k, v := range user {
		fmt.Printf("Key: %s - Value: %s\n", k, v)
	}
	fmt.Println("\n")

	for k := range user {
		fmt.Printf("Key: %s\n", k)
	}

	fmt.Println("----------------------------------------------------------\n")
	fmt.Println("Pointer\n")

	valor := 987
	direccion := &valor

	fmt.Printf("Type: %T\n", direccion)
	fmt.Println("direccion, *direccion")
	fmt.Println(direccion, *direccion)

	fmt.Println("----------------------------------------------------------\n")
	fmt.Println("Funciones anonimas\n")

	funcion := sumador()

	for i := 0; i < 10; i++ {
		fmt.Println(i, funcion(i))
	}

	fmt.Println("----------------------------------------------------------\n")
	fmt.Println("Estructuras\n")

	usuario := Usuario{1, "g@g.com", true}
	fmt.Println(usuario)
	fmt.Printf("Type: %T\n", usuario)
	fmt.Printf("Usuario.email: %s\n", usuario.email)

	fmt.Println(usuario.toJSON())
	// listUsuario := []Usuario{2, "g2@g.com", false}
	// fmt.Println(listUsuario)

	usuario.deactivateUser()

	fmt.Println("----------------------------------------------------------\n")
	fmt.Println("Interfaces\n")

	circulo := Circulo{x: 0, y: 0, radio: 2}
	rectangulo := Rectangulo{10, 20}

	fmt.Printf("Area del circulo: %f\n", getArea(circulo))
	fmt.Printf("Area del rectangulo: %f\n\n", getArea(rectangulo))

	fmt.Println("----------------------------------------------------------\n")
	fmt.Println("Web\n")
	fmt.Println("Se inicio un servicio web en el puerto 3000 -> http://localhost:3000/")
	http.HandleFunc("/", Index)
	http.ListenAndServe(":3000", nil)

}
