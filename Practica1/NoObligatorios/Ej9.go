/*Realice un programa que reciba una frase e imprima en pantalla
la misma frase reemplazando las ocurrencias de “jueves” por
“martes” respetando las letras minúsculas o mayúsculas de la
palabra original en su posición correspondiente. Por ejemplo, se
reemplaza “Jueves” por “Martes” o “jueveS” por “marteS”*/

package practica1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func esMayuscula(letra rune) bool {
	return unicode.IsUpper(letra)
}

func Ej9() {

	var frase string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Ingrese una frase: ")

	if scanner.Scan() {
		frase = scanner.Text()
	}

	aux := strings.FieldsFunc(frase, unicode.IsSpace)
	for i := 0; i < len(aux); i++ {
		if strings.EqualFold(aux[i], "Martes") {
			palabra := []rune(aux[i])
			palabra[0] = palabra[0] - 3
			palabra[1] = palabra[1] + 20
			palabra[2] = palabra[2] - 13
			palabra[3] = palabra[3] + 2
			aux[i] = string(palabra)
		} else if strings.EqualFold(aux[i], "Jueves") {
			palabra := []rune(aux[i])
			palabra[0] = palabra[0] + 3
			palabra[1] = palabra[1] - 20
			palabra[2] = palabra[2] + 13
			palabra[3] = palabra[3] - 2
			aux[i] = string(palabra)
		}
	}
	fmt.Println(strings.Join(aux, " "))
}
