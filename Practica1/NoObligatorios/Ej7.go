/*Las temperaturas de los pacientes de un hospital se dividen en 3 
grupos: alta (mayor de 37.5), normal (entre 36 y 37.5) y baja 
(menor de 36). Se deben leer 10 temperaturas de pacientes e 
informar el porcentaje de pacientes de cada grupo. Luego se 
debe imprimir el promedio entero entre la temperatura máxima y 
la temperatura mínima. 
a. ¿Se puede utilizar el case para tipos reales en otros 
lenguajes? 
b. ¿ Cómo se realizan las conversiones entre reales (punto 
flotante) y enteros en otros lenguajes ? 
Sub-objetivo: El tipado fuerte, usar casting. Operaciones y  
E/S con float. Casting en otros lenguajes. */

package practica1

import "fmt"

func Ej7() {
	var temp float64
	var alta, normal, baja int
	var pAlta, pNormal, pBaja float64
	var maxTemp, minTemp float64 = 1, 1000
	var promedio int

	for i := 0; i < 10; i++ {
		fmt.Printf("Ingrese la temperatura del paciente: ")
		fmt.Scan(&temp)

		if(temp < 36) {
			baja += 1
			if (temp < minTemp) {
				minTemp = temp
			}
		} else if (temp > 36 && temp < 37.5) {
			normal += 1
			if (temp < minTemp) {
				minTemp = temp
			}
			if (temp > maxTemp) {
				maxTemp = temp
			}
		} else {
			alta += 1
			if (temp > maxTemp) {
				maxTemp = temp
			}
		}
	}
	promedio = int(maxTemp + minTemp) / 2
	pAlta = float64(alta) / 10 * 100
	pNormal = float64(normal) / 10 * 100
	pBaja = float64(baja) / 10 * 100
	fmt.Printf("Porcentaje de pacientes con temperatura alta: %.2f%%\n", pAlta)
	fmt.Printf("Porcentaje de pacientes con temperatura normal: %.2f%%\n", pNormal)
	fmt.Printf("Porcentaje de pacientes con temperatura baja: %.2f%%\n", pBaja)
	fmt.Printf("Promedio entero entre la temperatura máxima y mínima: %d\n", promedio)
}
