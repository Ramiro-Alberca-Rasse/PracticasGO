package Practica3

import (
	"math"
	"sync"
)

/*1) Realice un programa que acepte un número entero positivo N como
entrada desde la línea de comandos y calcule todos los números
primos menores o iguales a N.
a) Realice el programa con una única goroutine que muestre en
pantalla la lista de números primos encontrados.
b) Realice el programa utilizando más de una goroutine para dividir
el trabajo de comprobación de primos entre múltiples goroutines
en paralelo
i)
Cada goroutine debe recibir un rango de números a
comprobar y devolver una lista de los números primos
encontrados
ii)
El programa principal debe recibir el número N y dividir el
trabajo en goroutines, asignando a cada una un rango de
números a comprobar
iii)
Una vez que todas las goroutines hayan finalizado, el
programa principal debe recopilar los resultados y mostrar
en pantalla la lista de números primos encontrados.
c) Realice la ejecución con N igual a 1.000, 100.000 y 1.000.000
tanto del programa a) como del b). Para cada caso calcule el
speed-up con la siguiente fórmula:
S(p) =  T(1) / T(p)
donde, S(p) es el speed-up, T(1) es el tiempo que tarda la
ejecución con una única goroutine y T(p) es el tiempo de
ejecución con p goroutines.
*/

func EsPrimo(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	limite := int(math.Sqrt(float64(n))) // Calcula la raíz cuadrada de n
	for i := 3; i <= limite; i += 2 {    // recorre todos los numeros impares desde la raiz de n hasta n
		if n%i == 0 {
			return false
		}
	}
	return true
}

func BuscarPrimosEnRango(inicio int, fin int) []int {
	var primos []int
	for i := inicio; i <= fin; i++ {
		if EsPrimo(i) {
			primos = append(primos, i)
		}
	}
	return primos
}

func Rutina(ini int, fn int, canalResultados chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	primosEncontrados := BuscarPrimosEnRango(ini, fn)
	canalResultados <- primosEncontrados
}


