package Practica2

/* Usando la estructura de datos definida en el ejercicio 9 resolver el
siguiente problema. Se dispone de una lista con la información de los
ingresantes a la Facultad del año anterior. De cada ingresante se
conoce: apellido, nombre, ciudad de origen, fecha de nacimiento (día,
mes, año), si presentó el título del colegio secundario y el código de la
carrera en la que se inscribe (APU, LI, LS). Con esta información de los
ingresantes se pide que recorra el listado una vez para:
a) Informar el nombre y apellido de los ingresantes cuya ciudad
origen es “Bariloche”.
b) Calcular e informar el año en que más ingresantes nacieron.
c) Informar la carrera con la mayor cantidad de inscriptos.
d) Eliminar de la lista aquellos ingresantes que no presentaron el
título */

//estructura del ej 9 es

import (
	"fmt"
)

type Fecha struct {
	Dia int
	Mes int
	Año int
}

type Ingresante struct {
	Apellido string
	Nombre   string
	Ciudad   string
	FechaNac Fecha
	Titulo   bool
	Carrera  string
}

type Nodo struct {
	Valor Ingresante
	Sig   *Nodo
}

type Lista struct {
	Nodo   *Nodo
	Tamaño int
}

// Agregar un elemento al final de la lista
func (l *Lista) PushBack(elem Ingresante) {
	nuevoNodo := &Nodo{Valor: elem, Sig: nil}

	if l.Nodo == nil {
		l.Nodo = nuevoNodo
	} else {
		actual := l.Nodo
		for actual.Sig != nil {
			actual = actual.Sig
		}
		actual.Sig = nuevoNodo
	}
	l.Tamaño++
}

func (l *Lista) Imprimir() {
	actual := l.Nodo
	for actual != nil {
		println("-", actual.Valor.Apellido, ",", actual.Valor.Nombre, "(Carrera:", actual.Valor.Carrera, ", Título:", actual.Valor.Titulo, ")")
		actual = actual.Sig
	}
}

func Obligatorio1() {
	listaIngresantes := &Lista{}

	listaIngresantes.PushBack(Ingresante{"Perez", "Juan", "Bariloche", Fecha{10, 5, 2002}, true, "APU"})
	listaIngresantes.PushBack(Ingresante{"Gomez", "Maria", "La Plata", Fecha{15, 8, 2003}, false, "LI"})
	listaIngresantes.PushBack(Ingresante{"Lopez", "Carlos", "Bariloche", Fecha{2, 1, 2002}, true, "LS"})
	listaIngresantes.PushBack(Ingresante{"Alberca Rasse", "Ramiro", "General Villegas", Fecha{17, 5, 2004}, true, "LS"})
	listaIngresantes.PushBack(Ingresante{"Diaz", "Ana", "Cordoba", Fecha{20, 11, 2004}, true, "APU"})
	listaIngresantes.PushBack(Ingresante{"Ruiz", "Pedro", "Mendoza", Fecha{5, 3, 2003}, false, "APU"})

	añosCont := make(map[int]int) //para cada año un contador
	carrerasCont := make(map[string]int)

	actual := listaIngresantes.Nodo
	var anterior *Nodo

	for actual != nil {
		if actual.Valor.Ciudad == "Bariloche" {
			fmt.Println(actual.Valor.Nombre, actual.Valor.Apellido)
		}

		añosCont[actual.Valor.FechaNac.Año]++

		carrerasCont[actual.Valor.Carrera]++

		if !actual.Valor.Titulo {
			if anterior == nil {
				listaIngresantes.Nodo = actual.Sig //pasamos al siguiente
			} else {
				anterior.Sig = actual.Sig
			}
			listaIngresantes.Tamaño--

			// se avanza el actual pero no el anterior porque el que estaba en medio se elimina
			actual = actual.Sig
		} else {
			//aca si se avanzan los dos
			anterior = actual
			actual = actual.Sig
		}
	}

	añoMax, cantAñoMax := 0, 0
	for año, cant := range añosCont { // 2 variales pero año no es indice, porque es de tipo map no array
		if cant > cantAñoMax {
			cantAñoMax = cant
			añoMax = año
		}
	}

	carreraMax, cantCarreraMax := "", 0
	for carrera, cant := range carrerasCont {
		if cant > cantCarreraMax {
			cantCarreraMax = cant
			carreraMax = carrera
		}
	}

	fmt.Println("Año con más nacimientos")
	fmt.Println("Año: ", añoMax, "Cantidad de nacimientos: ", cantAñoMax)

	fmt.Println("Carrera con más inscriptos: ", carreraMax, "Cantidad de inscriptos: ", cantCarreraMax)

	fmt.Println("Lista de ingresantes después de eliminar los que no presentaron el título:")
	listaIngresantes.Imprimir()
}
