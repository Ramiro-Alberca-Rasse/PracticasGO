/* Implemente una serie de funciones para manejar slices de enteros que
estadísticamente tienen muchas rachas de números repetidos. Utilice
una estructura (que con el objetivo de ahorrar memoria) tenga en cada
elemento el número entero y la cantidad de ocurrencias. Implemente:
func New(s slice) OptimumSlice
func IsEmpty(o OptimumSlice) bool
func Len(o OptimumSlice) int
func FrontElement(o OptimumSlice) int
func LastElement(o OptimumSlice) int
func Average(o OptimumSlice) float64
func Occurrences(o OptimumSlice, element int) int
func IndexOf(o OptimumSlice, element int) int //Buscar
primera aparición de un valor
func Mode(o OptimumSlice) int //El número que más se
repite
func Occurrences(o OptimumSlice, element int) int
func Insert(o OptimumSlice, element int, position int)
OptimumSlice
func SliceArray(o OptimumSlice) []int
Por ejemplo, si se invoca Insert con o =
que sería la representación del arreglo:
{3,3,3,3,3,1,1,1,1,1,1,1,23,23,23,23,23,23,3,3,3,3,3,3,3,3,7,5,5,5}
PRÁCTICA 2
y donde X[Y], X es el elemento e Y es la cantidad de ocurrencias
consecutivas
element = 9
position = 6
el resultado sería:
Nota: no se permite realizar el Insert convirtiendo el OptimunSlice a un slice,
insertar,  y luego volver a convertirlo.
Este último ejercicio es el que deben entregar*/

package main
import "fmt"
type slice []int

type racha struct {
	Valor       int
	Ocurrencias int
}


type OptimumSlice struct {
	sliceComprimido []racha
}

func New(s slice) OptimumSlice {

	if len(s) == 0 {
		return OptimumSlice{}
	}

	var listaRachas []racha
	valor := s[0]
	cont := 1

	for i := 1; i < len(s); i++ { // recorre el slice y separa los valores
		if s[i] == valor {
			cont += 1
		} else {
			// el numero cambio
			a := racha{Valor: valor, Ocurrencias: cont}
			listaRachas = append(listaRachas, a)
			valor = s[i]
			cont = 1
		}
	}

	a := racha{Valor: valor, Ocurrencias: cont}
	listaRachas = append(listaRachas, a)

	return OptimumSlice{sliceComprimido: listaRachas}
}

func IsEmpty(o OptimumSlice) bool {
	if len(o.sliceComprimido) == 0 {
		return true
	}
	return false
}

func Len(o OptimumSlice) int {
	total := 0
	for i := range o.sliceComprimido {
		total += o.sliceComprimido[i].Ocurrencias
	}
	return total
}

func FrontElement(o OptimumSlice) int {
	if IsEmpty(o) {
		return 0
	}
	return o.sliceComprimido[0].Valor
}
func LastElement(o OptimumSlice) int {
	if IsEmpty(o) {
		return 0
	}
	return o.sliceComprimido[len(o.sliceComprimido)-1].Valor
}
func Average(o OptimumSlice) float64 {
	if IsEmpty(o) {
		return 0.0
	}
	resultado := 0.0
	contador := 0.0 // cantidad de elementos
	for _, valor := range o.sliceComprimido { //
		resultado += float64((valor.Valor * valor.Ocurrencias)) // suma de todos los elementos
		contador += float64(valor.Ocurrencias)
	}
	resultado = resultado / contador
	return resultado
}
func Occurrences(o OptimumSlice, element int) int {
	contador := 0
	for _, valor := range o.sliceComprimido {
		if valor.Valor == element { // suma las ocurrencias de ese numero
			contador += valor.Ocurrencias
		}
	}
	return contador
}
func IndexOf(o OptimumSlice, element int) int {
	posicion := -1
	for _, valor := range o.sliceComprimido {
		posicion += valor.Ocurrencias
		if valor.Valor == element {
			return posicion - valor.Ocurrencias + 1 // devuelve la posicion de la primera aparicion del numero
		}
	}
	return -1 // si no se encuentra devuelve -1
}

func Mode(o OptimumSlice) int {

	if IsEmpty(o) {
		return 0
	}
	maximo := -1
	numeroMaximo := 0
	frecuencias := make(map[int]int)

	for _, racha := range o.sliceComprimido { // crear el map con todos los numeros
		frecuencias[racha.Valor] += racha.Ocurrencias
	}

	for numero, totalOcurrencias := range frecuencias { // recorre el mapa y busca el numero con mas ocurrencias
		if totalOcurrencias > maximo {
			maximo = totalOcurrencias
			numeroMaximo = numero
		}
	}

	return numeroMaximo
}

func Insert(o OptimumSlice, element int, position int) OptimumSlice {
	contador := 0         // contador para saber cuántas veces se repite el número en la racha
	contadorPosicion := 0 // posicion como si fuera un arreglo
	encontrado := false

	if position < 0 || position > Len(o) {
        return o 
    }
	if position == 0 && IsEmpty(o) { // caso base, si el slice esta vacio y se inserta en la posicion 0
		o.sliceComprimido = append(o.sliceComprimido, racha{Valor: element, Ocurrencias: 1})
		return o
	}	
	for i, valor := range o.sliceComprimido { // recorre el slice comprimido y busca la racha donde se encuentra la posicion
		contador = 0
		for contador < valor.Ocurrencias {
			if contadorPosicion == position {
				encontrado = true
				if valor.Valor == element {
					o.sliceComprimido[i].Ocurrencias += 1
					return o
				}
				break
			}
			contador += 1
			contadorPosicion += 1
		}
		if contadorPosicion == position { // por si la posicion es la ultima
			encontrado = true
			if valor.Valor == element {
				o.sliceComprimido[i].Ocurrencias += 1
				return o
			}
		}
		if encontrado {
			valorOriginal := valor.Valor
			ocurrenciasOriginales := valor.Ocurrencias
			aux := ocurrenciasOriginales - contador     // cantidad de veces que se repite el número en la racha después de la posición
			o.sliceComprimido[i].Ocurrencias = contador // cantidad de veces que se repite el número en la racha antes de la posición
			if contador == 0 && aux > 0 {               // en caso que se divida en 2
				o.sliceComprimido = append(o.sliceComprimido, racha{})
				copy(o.sliceComprimido[i+1:], o.sliceComprimido[i:])
				o.sliceComprimido[i] = racha{element, 1}
				o.sliceComprimido[i+1] = racha{Valor: valorOriginal, Ocurrencias: aux}
				break
			} else if aux > 0 && contador != 0 { // en caso que se divida en 3
				o.sliceComprimido = append(o.sliceComprimido, racha{})
				o.sliceComprimido = append(o.sliceComprimido, racha{})
				copy(o.sliceComprimido[i+3:], o.sliceComprimido[i+1:])
				o.sliceComprimido[i] = racha{Valor: valorOriginal, Ocurrencias: contador}
				o.sliceComprimido[i+1] = racha{Valor: element, Ocurrencias: 1}
				o.sliceComprimido[i+2] = racha{Valor: valorOriginal, Ocurrencias: aux}
				break
			} else { // en caso que no se divida
				if i+1 < len(o.sliceComprimido) && o.sliceComprimido[i+1].Valor == element {
					o.sliceComprimido[i+1].Ocurrencias += 1
					break
				}
				o.sliceComprimido = append(o.sliceComprimido, racha{})
				copy(o.sliceComprimido[i+2:], o.sliceComprimido[i+1:])
				o.sliceComprimido[i+1] = racha{Valor: element, Ocurrencias: 1}
				break
			}
		}
	}
	return o
}

func SliceArray(o OptimumSlice) []int { // convierte a slice
	var slice []int
	for _, racha := range o.sliceComprimido {
		for j := 0; j < racha.Ocurrencias; j++ {
			slice = append(slice, racha.Valor)
		}
	}
	return slice
}

func main() {
	base := slice{3, 3, 3, 3, 3, 7, 7, 7, 5, 5}
	o := New(base)
	fmt.Printf("Estructura: %v\n", o.sliceComprimido)
	fmt.Printf("Len: %d\n", Len(o))
	fmt.Printf("Array: %v\n\n", SliceArray(o))

	o = Insert(o, 9, 0)
	fmt.Printf("Insert inicio: %v\n\n", SliceArray(o))

	o = Insert(o, 1, 4)
	fmt.Printf("Insert medio (rompe): %v\n", SliceArray(o))
	fmt.Printf("Bloques: %v\n\n", o.sliceComprimido)

	idxSiete := IndexOf(o, 7)
	o = Insert(o, 7, idxSiete)
	fmt.Printf("Insert fusión: %v\n", SliceArray(o))
	fmt.Printf("Bloques fusionados: %v\n\n", o.sliceComprimido)

	o = Insert(o, 8, Len(o))
	fmt.Printf("Insert final: %v\n\n", SliceArray(o))

	fmt.Printf("Front: %d\n", FrontElement(o))
	fmt.Printf("Last: %d\n", LastElement(o))
	fmt.Printf("Average: %.2f\n", Average(o))
	fmt.Printf("Mode: %d\n", Mode(o))
}