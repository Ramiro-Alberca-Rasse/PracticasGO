/*Realizar un programa que lea el punto cardinal (como caracter o 
string) del cual viene el viento (‘N’, ‘S’, ‘E’, ‘O’) y envíe a la salida 
estándar hacia cuál se dirigiría. 
Sub-objetivo: Uso de case con la opción por default. E/S 
caracteres o strings. 
a. ¿Cómo se escribe el default en el case de otros lenguajes?*/

package practica1

import "fmt"
import "strings"

func Ej8() {
	var viento string
	fmt.Printf("Ingrese el punto cardinal del viento (N, S, E, O): ")
	fmt.Scan(&viento)
	viento = strings.ToUpper(viento)

	switch {
	case viento == "S":
		fmt.Printf("El viento viene del Sur, se dirige hacia el Norte\n")
	case viento == "N":
		fmt.Printf("El viento viene del Norte, se dirige hacia el Sur\n")
	case viento == "E":
		fmt.Printf("El viento viene del Este, se dirige hacia el Oeste\n")
	case viento == "O":
		fmt.Printf("El viento viene del Oeste, se dirige hacia el Este\n")
	default:
		fmt.Printf("Punto cardinal inválido\n")
	}

}