package Practica1

/*Realice las modificaciones necesarias al ejercicio anterior para que en
lugar de reemplazar la palabra “jueves” por “martes” ahora se
reemplace “miércoles” por “automóvil”. Piense qué impacto tuvieron
esas modificaciones en el programa que había realizad*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// mejor forma de hacerlo que como lo hice en el ejercicio 9
func EjObligatorio1() {
	var frase string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Ingrese una frase: ")

	if scanner.Scan() {
		frase = scanner.Text()
	}

	// separa las palabras de la frase
	aux := strings.FieldsFunc(frase, unicode.IsSpace) //frase
	palabraReemplazo := []rune("automóvil")

	for i := 0; i < len(aux); i++ {

		if strings.EqualFold(aux[i], "miércoles") {
			palabraOriginal := []rune(aux[i]) // guardamos en formato de runas
			nuevaPalabra := make([]rune, len(palabraReemplazo))

			for j := 0; j < len(palabraOriginal) && j < len(palabraReemplazo); j++ {

				if unicode.IsUpper(palabraOriginal[j]) {
					nuevaPalabra[j] = unicode.ToUpper(palabraReemplazo[j])
				} else {
					nuevaPalabra[j] = unicode.ToLower(palabraReemplazo[j])
				}

			}
			aux[i] = string(nuevaPalabra) // frase actualizada
		}
	}

	fmt.Println(strings.Join(aux, " ")) // unir con espacios
}
