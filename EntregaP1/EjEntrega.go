/*Realice un programa que reciba una palabra como argumento y lee de
la entrada una frase. Luego, el programa debe imprimir la frase que
leyó con cada una de las ocurrencias de la palabra con las mayúsculas
y minúsculas invertidas. Por ejemplo, si la frase es:
“Parece peqUEño, pero no es tan pequeÑo el PEQUEÑO”
y la palabra es “PEQUEÑO” entonces el programa imprimirá:
“Parece PEQueÑO, pero no es tan PEQUEñO el pequeño”
Tenga en cuenta que la palabra a buscar puede ser ingresada con
mayúsculas y minúsculas mezcladas.*/

package practica1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var frase string

func invertirMayusculas(s string) string { //Convierte en mayúscula lo que es minúscula y viceversa
	runas := []rune(s)
	for i := 0; i < len(runas); i++ {
		if unicode.IsUpper(runas[i]) {
			runas[i] = unicode.ToLower(runas[i])
		} else {
			runas[i] = unicode.ToUpper(runas[i])
		}
	}	
	s = string(runas)
	return s
}

/*func Separar(r rune) bool {  //Función para separar la frase en palabras
	if r == ' ' || r == '.' || r == ',' || r == '¿' || r == '?' || r == '¡' || r == '!' {
	    return true
	} else {	
		return false
	}
}

 func EjEntrega(palabra string) string {

	scanner := bufio.NewScanner(os.Stdin) // Scanner que lee desde terminal
	fmt.Println("Ingrese una frase: ")
	if scanner.Scan() { // Si el usuario ingresa una frase, se guarda en la variable frase
		frase = scanner.Text() 
	}
	aux := strings.FieldsFunc(frase, Separar) // Separa la frase usando el criterio de la funcion Separar
	for i := 0; i < len(aux); i++ {
		if strings.ToUpper(aux[i]) == strings.ToUpper(palabra) { // Compara las palabras en mayusuclas para que no importe como se ingresó la palabra
			aux[i] = invertirMayusculas(aux[i]) // Si la palabra coincide, se invierten las mayúsculas y minúsculas usando la función
		}
	}
	frase = strings.Join(aux, " ") // Se vuelve a unir la frase con las palabras modificadas
	return frase
} */


func EjEntrega(palabra string) string {
    scanner := bufio.NewScanner(os.Stdin) 
    fmt.Println("Ingrese una frase: ")
    if scanner.Scan() { 
        frase = scanner.Text() 
    }

    resultado := ""
    aux := frase
    palabraBusqueda := strings.ToLower(palabra)

    for { // porque puede haber mas de 1 palabra que cumpla
        // indice de la palabra
        idx := strings.Index(strings.ToLower(aux), palabraBusqueda)
        
        if idx == -1 { // Si no encuentra más
            resultado += aux // agregamos el resto de la frase
            break
        }


        resultado += aux[:idx]
        
        palabraEncontrada := aux[idx : idx+len(palabra)]
        
        resultado += invertirMayusculas(palabraEncontrada)
        
        // le agregamos desde la palabra hasta el final
        aux = aux[idx+len(palabra):]
    }
    
    return resultado
}

func main() {
	var palabra string
    if len(os.Args) < 2 { // Significa que no se paso nada extra como parametro
        fmt.Println("Falta ingresar la palabra como argumento")
        return
    }
    palabra = os.Args[1] // 1 porqaue el 0 es el nombre del programa
	frase := EjEntrega(palabra)
	fmt.Println(frase)
}

