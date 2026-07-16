/*  Desarrolla un programa que implemente un sistema de planificación
(scheduler) utilizando 5 goroutines (el main y 4 más). El programa debe
generar una serie de números enteros aleatorios, cada uno con una
prioridad aleatoria entre 0 y 3 (donde 0 es la prioridad más alta y 3 la
más baja).
El scheduler debe seguir las siguientes reglas:
a) El scheduler debe procesar los datos en orden de prioridad,
comenzando por los de prioridad 0, luego 1, 2 y 3.
b) Mientras haya datos de prioridad 0, el scheduler debe
procesarlos exclusivamente.
c) Si no hay datos de prioridad 0 y hay goroutines disponibles, el
scheduler puede asignarles datos de menor prioridad para su
procesamiento.
PRÁCTICA 3
d) Una vez que no haya datos de prioridad 0, el scheduler debe
pasar a procesar los datos de prioridad 1, y así sucesivamente,
utilizando las goroutines disponibles de forma dinámica.
e) El programa principal debe generar los datos numéricos
aleatorios con sus respectivas prioridades aleatorias y distribuir el
trabajo a las goroutines disponibles manteniendo el orden en el
que llegan los datos por prioridad.
f) Con los datos se debe:
i)
Prioridad 0: Sumar los dígitos del número y almacenar el
resultado en un archivo llamado "prioridad0.txt" en formato
de par ordenado (0, resultado).
ii)
iii)
iv)
Prioridad 1: Invertir los dígitos del número y almacenar el
resultado en un archivo llamado "prioridad1.txt" en formato
de par ordenado (1, resultado).
Prioridad 2: Multiplicar el número por un valor constante
(por ejemplo, 10) e imprimir el resultado en la consola.
Prioridad 3: Acumular los números y mostrar el valor
acumulado en la consola cada vez que se procesa un dato.
Tip: Puedes utilizar la librería math/rand para generar números
aleatorios. */

package Practica3

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

var candado sync.Mutex
var acumuladorPrioridad3 int
var candadoAcumulador sync.Mutex


type Numeros struct {
	Numero    int
	Prioridad int
}

// Función DE GORUTINES
func Rutinas(canal chan Numeros, wg *sync.WaitGroup, wgNivel *sync.WaitGroup) {
	defer wg.Done() // Avisa que termino
	var num Numeros
	for num = range canal {

		if num.Prioridad == 0 {
			suma := 0
			numero := num.Numero
			for numero > 0 {
				suma += numero % 10
				numero = numero / 10
			}
			GuardarEnArchivo(num.Prioridad, suma, "prioridad0.txt")
		} else if num.Prioridad == 1 {
			resultado := ""
			numero := num.Numero
			if numero == 0 {
				resultado = "0"
			}

			for numero > 0 {
				resultado += strconv.Itoa(numero % 10)
				numero = numero / 10
			}
			resultadoInt, _ := strconv.Atoi(resultado)
			GuardarEnArchivo(num.Prioridad, resultadoInt, "prioridad1.txt")
		} else if num.Prioridad == 2 {
			resultado := num.Numero * 10
			fmt.Printf("Resultado de multiplicar por 10: %d\n", resultado)

		} else if num.Prioridad == 3 {
			candadoAcumulador.Lock()
			acumuladorPrioridad3 += num.Numero
			fmt.Printf("Prioridad 3 - Nuevo valor acumulado: %d\n", acumuladorPrioridad3)
			candadoAcumulador.Unlock()
		}

		wgNivel.Done()
	}
}

func GuardarEnArchivo(prioridad int, resultado int, nombreArchivo string) {
	candado.Lock()
	defer candado.Unlock()

	textoAGuardar := fmt.Sprintf("(%d, %d)\n", prioridad, resultado)
	archivo, _ := os.OpenFile(nombreArchivo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer archivo.Close()

	archivo.WriteString(textoAGuardar)
}


