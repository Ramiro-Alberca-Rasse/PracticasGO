package Practica3

import (
	"math/rand"
	"sync"
	"time"
)

/* Realice un programa que simule la atención de clientes en las cajas de
un supermercado. La atención de cada cliente por parte del cajero se
debe simular con un timer entre 0 y 1 segundo.
a) Realice el programa haciendo esperar a los clientes en un única
cola global y luego enviándolo a la caja para su atención cuando
esta se encuentre disponible
b) Realice el programa haciendo esperar a los clientes en colas
individuales por caja asignando la caja para su atención con un
algoritmo de distribución round-robin
c) Realice el programa haciendo esperar a los clientes en colas
individuales por caja asignando la caja para su atención a aquella
que tenga la cola más corta
d) Imprima los tiempos de ejecución de cada uno de los programas
implementados en a), b) y c)
*/

type Cliente struct {
	id    int
	senal chan int // cada cliente tiene un canal para saber si esta siendo atendido
}

func CrearCliente(id int) *Cliente {
	return &Cliente{
		id:    id,
		senal: make(chan int),
	}
}

// go rutine
func nuevoCliente(cliente *Cliente, wg *sync.WaitGroup) {
	defer wg.Done()
	<-cliente.senal
	<-cliente.senal
}

func cajero(cola <-chan *Cliente, wg *sync.WaitGroup) {
	defer wg.Done()
	for x := range cola {
		x.senal <- 1 //le avisa al client

		tiempoAtencion := time.Duration(rand.Intn(1)) * time.Second // Clcula tiempo random entre 0 y 1
		time.Sleep(tiempoAtencion)

		x.senal <- 0
	}
}

func SuperColaGlobal(numCajeros int, numClientes int) time.Duration {
	inicio := time.Now()
	var wgClientes, wgCajeros sync.WaitGroup

	colaGlobal := make(chan *Cliente, numClientes)

	for i := 0; i < numCajeros; i++ { //gorutne cajero
		wgCajeros.Add(1)
		go cajero(colaGlobal, &wgCajeros)
	}

	for i := 0; i < numClientes; i++ { //gorutne cliente
		cliente := CrearCliente(i)
		wgClientes.Add(1)
		go nuevoCliente(cliente, &wgClientes)
		colaGlobal <- cliente
	}

	close(colaGlobal) // Se cierra el super
	wgClientes.Wait()
	wgCajeros.Wait()

	return time.Since(inicio)
}

func SuperVariasColas(numCajeros int, numClientes int) time.Duration {
	inicio := time.Now()
	var wgClientes, wgCajeros sync.WaitGroup

	// array de canales
	colas := make([]chan *Cliente, numCajeros)
	for i := 0; i < numCajeros; i++ {
		canal := make(chan *Cliente, numClientes)
		colas[i] = canal
		wgCajeros.Add(1)
		go cajero(colas[i], &wgCajeros)
	}

	for i := 0; i < numClientes; i++ {
		cliente := CrearCliente(i)
		wgClientes.Add(1)
		go nuevoCliente(cliente, &wgClientes)

		cajaAsignada := i % numCajeros
		colas[cajaAsignada] <- cliente // cada cliente tiene una caja
	}

	for i := 0; i < numCajeros; i++ {
		close(colas[i])
	}
	wgClientes.Wait()
	wgCajeros.Wait()

	return time.Since(inicio)
}

func SuperColaMasCorta(numCajeros, numClientes int) time.Duration {
	inicio := time.Now()
	var wgClientes, wgCajeros sync.WaitGroup

	colas := make([]chan *Cliente, numCajeros)
	for i := 0; i < numCajeros; i++ {
		colas[i] = make(chan *Cliente, numClientes)
		wgCajeros.Add(1)
		go cajero(colas[i], &wgCajeros)
	}

	for i := 0; i < numClientes; i++ {
		cliente := CrearCliente(i)
		wgClientes.Add(1)
		go nuevoCliente(cliente, &wgClientes)

		cajaMasCorta := 0 // primero la 0
		minClientes := len(colas[0])

		for j := 1; j < numCajeros; j++ {
			if len(colas[j]) <= minClientes {
				minClientes = len(colas[j])
				cajaMasCorta = j
			}
		}
		colas[cajaMasCorta] <- cliente
	}

	for i := 0; i < numCajeros; i++ {
		close(colas[i])
	}
	wgClientes.Wait()
	wgCajeros.Wait()

	return time.Since(inicio)
}

