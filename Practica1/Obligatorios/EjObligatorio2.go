package Practica1

/* Realice un programa que reciba una frase e imprima en pantalla la
misma frase con cada una de las palabras invertidas siempre que su
ubicación sea impar en la frase comenzando a contar las palabras
desde 1, por ejemplo, si la frase ingresada es:
Qué lindo día es hoy.
El programa imprimirá:
éuQ lindo aíd es yoh.*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func EjObligatorio2() {
	var frase string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Ingrese una frase: ")

	if scanner.Scan() {
		frase = scanner.Text()
	}

	// separa las palabras de la frase
	aux := strings.FieldsFunc(frase, unicode.IsSpace) //frase

	for i := 1; i <= len(aux); i++ {
		println(i % 2)
		if (i)%2 != 0 { // no es par
			palabra := []rune(aux[i-1]) // porque i empieza en 1 y no en 0
			nuevaPalabra := make([]rune, len(palabra))
			for j := 0; j < len(palabra); j++ {
				nuevaPalabra[j] = palabra[len(palabra)-(j+1)] //porque el indice len(palabra) - 0 no existe
			}
			aux[i-1] = string(nuevaPalabra)
		}
	}
	fmt.Println(strings.Join(aux, " "))
}


