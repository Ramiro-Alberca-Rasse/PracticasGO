package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"sync"
	"time"

	Practica3 "example.com/miapp/Practica3/Obligados"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Ingrese N y la cantidad de gorutines:")
		return
	}

	N, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error, N debe ser un número entero mayor a 0")
		return
	}
	cant, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Error, la cantidad de goroutines debe ser un número entero mayo a 0")
		return
	}

	if N <= 0 || cant <= 0 {
		fmt.Println("Error, los valores tienen que ser mayores a 0")
		return
	}

	start := time.Now()
	var primosFinales []int

	if cant == 1 {
		primosFinales = Practica3.BuscarPrimosEnRango(2, N) // 2 porque es el primer primo
	} else {

		var wg sync.WaitGroup // para saber cuando terminan

		canalResultados := make(chan []int, cant)

		rangoBase := N / cant // rango de elementos por gorutine
		inicio := 2
		fin := rangoBase - 1
		for i := 0; i < cant; i++ {

			if i == cant-1 {
				fin = N // La última goroutine se queda con el remanente si N no es divisible por cant
			}

			wg.Add(1)
			go Practica3.Rutina(inicio, fin, canalResultados, &wg)

			inicio = fin + 1
			fin = inicio + rangoBase - 1
		}

		wg.Wait()
		close(canalResultados)

		var subLista []int

		for subLista = range canalResultados { // cada gorutine devuelve un slice
			primosFinales = append(primosFinales, subLista...)
		}

	}

	duracion := time.Since(start)

	fmt.Println("N =", N, "| Goroutines (p) =", cant, "| Cantidad de primos =", len(primosFinales), "| Tiempo:", duracion)

	/* RESULTADOS DE PRUEBA

	N = 1000
	S(5) = T(1) / T(5) //  S(5) = 0s / 0s

	N = 100000
	S(5) = T(1) / T(5) //  S(5) = 7ms / 3,5ms // S(5) = 2 ms

	N = 1000000
	S(5) = T(1) / T(5) //  S(5) = 177ms / 49,5ms // S(5) = 3,57 ms

	*/

	// EJ 2

	numCajeros := 3
	numClientes := 20

	println("Simulando atención para", numClientes, "clientes con", numCajeros, "cajeros...")
	println()

	tiempoA := Practica3.SuperColaGlobal(numCajeros, numClientes)
	println("A) Cola Global:", tiempoA)

	tiempoB := Practica3.SuperVariasColas(numCajeros, numClientes)
	println("B) Round-Robin:", tiempoB)

	tiempoC := Practica3.SuperColaMasCorta(numCajeros, numClientes)
	println("C) Cola Más Corta:", tiempoC)

	// Ej3

	var acumuladorPrioridad3 int
	var numeros []Practica3.Numeros
	var wg sync.WaitGroup
	var wgNivel sync.WaitGroup

	canal := make(chan Practica3.Numeros)

	// genera de 20 datos aleatorios
	for i := 0; i < 20; i++ {
		num := rand.Intn(1000)
		prioridad := rand.Intn(4)
		numeros = append(numeros, Practica3.Numeros{Numero: num, Prioridad: prioridad})
	}

	// llamar a las 4 goroutines
	wg.Add(4)
	for i := 0; i < 4; i++ {
		go Practica3.Rutinas(canal, &wg, &wgNivel)
	}

	// Ordena por prioridad.
	sort.Slice(numeros, func(i, j int) bool {
		return numeros[i].Prioridad < numeros[j].Prioridad
	})
	prioridad := 0
	for _, num := range numeros {
		if num.Prioridad > prioridad {
			prioridad = num.Prioridad
			wgNivel.Wait() // Espera que todas las goroutines terminen el nivel de prioridad actual
		}
		wgNivel.Add(1)
		canal <- num

	}
	close(canal)
	wg.Wait() // Espera que todas las goroutines terminen

	// PRINTS PARA TESTEO 
	fmt.Println("\n=================================================")
	fmt.Println("PROGRAMA FINALIZADO. RESUMEN DE DATOS PROCESADOS:")
	fmt.Println("=================================================")

	// Muestro la lista ordenada para que puedas verificar los resultados en los TXT
	fmt.Println("\n1. Lista de números generados (ordenados por el Scheduler):")
	for _, n := range numeros {
		fmt.Printf("   -> Número: %4d | Prioridad: %d\n", n.Numero, n.Prioridad)
	}

	fmt.Println("\n2. Resultados de las operaciones:")
	fmt.Println("   [Prioridad 0] -> Guardados en 'prioridad0.txt'")
	fmt.Println("   [Prioridad 1] -> Guardados en 'prioridad1.txt'")
	fmt.Println("   [Prioridad 2] -> Impresos arriba durante la ejecución")
	fmt.Printf("   [Prioridad 3] -> VALOR ACUMULADO FINAL: %d\n", acumuladorPrioridad3)
	fmt.Println("=================================================")



}
