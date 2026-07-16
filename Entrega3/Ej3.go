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

package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"sync"
)

type Numeros struct {
	numero    int
	prioridad int
}

// Función DE GORUTINES
func rutinas(canales []chan Numeros, wg *sync.WaitGroup) {
	defer wg.Done() // Avisa que termino

	for {
		var num Numeros

		// Lo que hace cada prioridad
		select {
		case num = <-canales[0]:
			if num.prioridad == -1 { 
				return
			}
			suma := 0
			numero := num.numero
			for numero > 0 {
				suma += numero % 10
				numero = numero / 10
			}
			guardarEnArchivo(num.prioridad, suma, "prioridad0.txt")

		case num = <-canales[1]:
			if num.prioridad == -1 { 
				return
			}
			resultado := ""
			numero := num.numero
			if numero == 0 { 
				resultado = "0"
			}

			for numero > 0 {
				resultado += strconv.Itoa(numero % 10)
				numero = numero / 10
			}

			resultadoInt, _ := strconv.Atoi(resultado)
			guardarEnArchivo(num.prioridad, resultadoInt, "prioridad1.txt")

		case num = <-canales[2]:
			if num.prioridad == -1 { 
				return
			}
			resultado := num.numero * 10
			fmt.Printf("Resultado de multiplicar por 10: %d\n", resultado)

		case num = <-canales[3]:
			if num.prioridad == -1 {
				return
			}
			candadoAcumulador.Lock()
			acumuladorPrioridad3 += num.numero
			fmt.Printf("Prioridad 3 - Nuevo valor acumulado: %d\n", acumuladorPrioridad3)
			candadoAcumulador.Unlock()
		}
	}
}


func guardarEnArchivo(prioridad int, resultado int, nombreArchivo string) {
	candado.Lock()
	defer candado.Unlock()

	textoAGuardar := fmt.Sprintf("(%d, %d)\n", prioridad, resultado)
	archivo, _ := os.OpenFile(nombreArchivo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer archivo.Close()

	archivo.WriteString(textoAGuardar)
}

var candado sync.Mutex
var acumuladorPrioridad3 int
var candadoAcumulador sync.Mutex
var numeros []Numeros
var wg sync.WaitGroup

func main() {
	chPrioridad0 := make(chan Numeros)
	chPrioridad1 := make(chan Numeros)
	chPrioridad2 := make(chan Numeros)
	chPrioridad3 := make(chan Numeros)

	// genera de 20 datos aleatorios
	for i := 0; i < 20; i++ {
		num := rand.Intn(1000)
		prioridad := rand.Intn(4)
		numeros = append(numeros, Numeros{numero: num, prioridad: prioridad})
	}

	// llamar a las 4 goroutines
	wg.Add(4)
	for i := 0; i < 4; i++ {
		go rutinas([]chan Numeros{chPrioridad0, chPrioridad1, chPrioridad2, chPrioridad3}, &wg)
	}

	// Ordena por prioridad.
	sort.Slice(numeros, func(i, j int) bool {
		return numeros[i].prioridad < numeros[j].prioridad
	})

	for _, num := range numeros {
		switch num.prioridad {
		case 0:
			chPrioridad0 <- num
		case 1:
			chPrioridad1 <- num
		case 2:
			chPrioridad2 <- num
		case 3:
			chPrioridad3 <- num
		}
	}

	// Envio señal para finalizar
	chPrioridad0 <- Numeros{numero: 0, prioridad: -1}
	chPrioridad1 <- Numeros{numero: 0, prioridad: -1}
	chPrioridad2 <- Numeros{numero: 0, prioridad: -1}
	chPrioridad3 <- Numeros{numero: 0, prioridad: -1}

	wg.Wait() // Espera que todas las goroutines terminen

	// --- PRINTS AGREGADOS AL FINAL PARA TESTEO ---
	fmt.Println("\n=================================================")
	fmt.Println("PROGRAMA FINALIZADO. RESUMEN DE DATOS PROCESADOS:")
	fmt.Println("=================================================")
	
	// Muestro la lista ordenada para que puedas verificar los resultados en los TXT
	fmt.Println("\n1. Lista de números generados (ordenados por el Scheduler):")
	for _, n := range numeros {
		fmt.Printf("   -> Número: %4d | Prioridad: %d\n", n.numero, n.prioridad)
	}

	fmt.Println("\n2. Resultados de las operaciones:")
	fmt.Println("   [Prioridad 0] -> Guardados en 'prioridad0.txt'")
	fmt.Println("   [Prioridad 1] -> Guardados en 'prioridad1.txt'")
	fmt.Println("   [Prioridad 2] -> Impresos arriba durante la ejecución")
	fmt.Printf("   [Prioridad 3] -> VALOR ACUMULADO FINAL: %d\n", acumuladorPrioridad3)
	fmt.Println("=================================================")
}