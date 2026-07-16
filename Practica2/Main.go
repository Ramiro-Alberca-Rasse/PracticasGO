package main

import (
	"fmt"
	Practica2 "example.com/miapp/Practica2/Obligatorios"
)

func main() {
	

	//Practica2

	Practica2.Obligatorio1()
	//ej2
	bc := &Practica2.Blockchain{}

	wallet1 := Practica2.CrearBilletera("1", "Lionel", "Messi")
	wallet2 := Practica2.CrearBilletera("2", "Emiliano", "Martinez")

	Trans := Practica2.CrearTransaccion("admin", wallet1.ID, 1000.0)
	bc.InsertarBloque(Trans) //daria error porque no hay saldo suficiente, pero como es admin se agrega el saldo igual

	trans2 := Practica2.CrearTransaccion(wallet2.ID, wallet1.ID, 5000.0) // da error
	err := bc.InsertarBloque(trans2)
	if err != nil {
		fmt.Println("Error al insertar transacción: ", err)
	}

	fmt.Println("Saldo ", wallet1.Nombre, ": $", bc.ObtenerSaldo(wallet1.ID))
	fmt.Println("Saldo ", wallet2.Nombre, ": $", bc.ObtenerSaldo(wallet2.ID))

	fmt.Println("Validando cadena: ", bc.ValidarCadena()) // true o false

	//ej3

	base := []int{3, 3, 3, 3, 3, 7, 7, 7, 5, 5}
	o := Practica2.New(base)
	fmt.Println("Estructura:", o.SliceComprimido)
	fmt.Println("Len:", Practica2.Len(o))
	fmt.Println("Array:", Practica2.SliceArray(o))
	fmt.Println()

	o = Practica2.Insert(o, 9, 0)
	fmt.Println("Insert inicio:", Practica2.SliceArray(o))
	fmt.Println()

	o = Practica2.Insert(o, 1, 4)
	fmt.Println("Insert medio (rompe):", Practica2.SliceArray(o))
	fmt.Println("Bloques:", o.SliceComprimido)
	fmt.Println()

	idxSiete := Practica2.IndexOf(o, 7)
	o = Practica2.Insert(o, 7, idxSiete)
	fmt.Println("Insert fusión:", Practica2.SliceArray(o))
	fmt.Println("Bloques fusionados:", o.SliceComprimido)
	fmt.Println()

	o = Practica2.Insert(o, 8, Practica2.Len(o))
	fmt.Println("Insert final:", Practica2.SliceArray(o))
	fmt.Println()

	fmt.Println("Front:", Practica2.FrontElement(o))
	fmt.Println("Last:", Practica2.LastElement(o))
	fmt.Println("Average:", Practica2.Average(o))
	fmt.Println("Mode:", Practica2.Mode(o))

}	


