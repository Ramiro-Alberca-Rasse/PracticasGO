package main
import (
	"fmt"
	"os"
	Practica1 "example.com/miapp/Practica1/Obligatorios"
)
func main() {
	Practica1.EjObligatorio1()
	Practica1.EjObligatorio2()

	//ej 3
	var palabra string
	if len(os.Args) < 2 { // Significa que no se paso nada extra como parametro
		fmt.Println("Falta ingresar la palabra como argumento")
		return
	}
	palabra = os.Args[1] // 1 porqaue el 0 es el nombre del programa
	frase := Practica1.Obligatorio3(palabra)
	fmt.Println(frase)
}
