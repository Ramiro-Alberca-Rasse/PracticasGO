/*2) Implemente una blockchain para que sea soporte de una cryptomoneda
que incluya la creación de billeteras para los clientes. Una blockchain, o
cadena de bloques, es un sistema digital distribuído que funciona como
un libro de contabilidad público e inmutable. Almacena información
sobre transacciones de forma segura y descentralizada, sin necesidad
de intermediarios. Cada transacción se agrupa en un bloque, que se
enlaza con el bloque anterior (de manera similar a una lista enlazada),
creando una cadena irrompible.
Utilice structs para representar la transacción (con el monto, el
identificador de quien envía dinero, el identificador de quien recibe ese
dinero y la fecha/hora de la transacción - timestamp -), el bloque (que
tienen el hash, el hash previo, la transacción (data) y la fecha/hora de
creación de dicho bloque), la cadena de bloques y la billetera (con el
identificador, nombre y apellido del cliente).
Diagrama de la cadena:
Tip: puede utilizar la librería crypto/sha256 para crear el hash del
bloque anterior.
a) Defina todos los tipos de datos que va a utilizar.
b) Programe funciones para:
i)
Crear una billetera
PRÁCTICA 2
ii)
iii)
iv)
v)
vi)
i)
Enviar una transacción
Insertar un bloque en la cadena
Obtener el saldo de un usuario (recorriendo toda la
cadena)
Realizar una función que valide que la cadena sea
consistente recorriéndola y verificando que el hash
almacenado del bloque anterior es válido
Si utilizó un slice para la estructura de la cadena de
bloques cambie la implementación con una lista enlazada.
Puede reutilizar la implementación del ejercicio 9. ¿Cuál
fue el impacto que tuvo en su programa?
¿Cómo validar que la transacción solo se pueda hacer si la
billetera que va a enviar los fondos tiene saldo suficiente? */

package Practica2

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"time"
)

type Billetera struct {
	ID       string
	Nombre   string
	Apellido string
}

type Transaccion struct {
	Monto     float64
	Origen    string
	Destino   string
	Timestamp time.Time
}

type Bloque struct {
	Hash       string
	HashPrevio string
	Data       Transaccion
	Timestamp  time.Time
}

type NodoBloque struct {
	Bloque *Bloque
	Sig    *NodoBloque
}

type Blockchain struct {
	Cabeza *NodoBloque // Primero nodo
	Cola   *NodoBloque // ultimo nodo
}

func CrearBilletera(id string, nombre string, apellido string) Billetera {
	billetera := Billetera{
		ID:       id,
		Nombre:   nombre,
		Apellido: apellido,
	}
	return billetera
}

func CrearTransaccion(origen string, destino string, monto float64) Transaccion {
	transaccion := Transaccion{
		Monto:     monto,
		Origen:    origen,
		Destino:   destino,
		Timestamp: time.Now(),
	}
	return transaccion
}

func calcularHash(hashPrevio string, trans Transaccion, timestamp time.Time) string { // le pasas todos los datos del bloque y crea una clave unica

	record := fmt.Sprintln(hashPrevio, trans.Origen, trans.Destino, trans.Monto, timestamp)
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
} // si el bloque cambia entonces la clave tambien

func (bc *Blockchain) InsertarBloque(trans Transaccion) error {

	if trans.Origen != "admin" { // porque sino nunca se podria hacer una transaccion, porque no hay saldo inicial
		saldoDisponible := bc.ObtenerSaldo(trans.Origen)
		if saldoDisponible < trans.Monto {
			err := errors.New("saldo insuficiente para la transacción")
			return err
		}
	}

	nuevoBloque := &Bloque{ //puntero
		Data:      trans,
		Timestamp: time.Now(),
	}

	// Es el primer bloque
	if bc.Cabeza == nil {
		nuevoBloque.HashPrevio = "0"
		nuevoBloque.Hash = calcularHash(nuevoBloque.HashPrevio, nuevoBloque.Data, nuevoBloque.Timestamp)

		nuevoNodo := &NodoBloque{Bloque: nuevoBloque, Sig: nil}
		bc.Cabeza = nuevoNodo // primero
		bc.Cola = nuevoNodo   // ultimo
	} else {
		nuevoBloque.HashPrevio = bc.Cola.Bloque.Hash // en hasprevi va el actual ultimo bloque
		nuevoBloque.Hash = calcularHash(nuevoBloque.HashPrevio, nuevoBloque.Data, nuevoBloque.Timestamp)

		nuevoNodo := &NodoBloque{Bloque: nuevoBloque, Sig: nil}
		bc.Cola.Sig = nuevoNodo //actualiza el ultimo actual para que apunte al nuevo
		bc.Cola = nuevoNodo     // ahora cola apunta al nuevo
	}

	return nil
}

func (bc *Blockchain) ObtenerSaldo(idUsuario string) float64 { //muy poco eficiente, se deberia tener una estructura de datos que guarde el saldo de cada usuario como un map
	saldo := 0.0
	actual := bc.Cabeza // primero

	for actual != nil {
		trans := actual.Bloque.Data // transaccion del primer bloque

		if trans.Origen == idUsuario {
			saldo -= trans.Monto
		} else if trans.Destino == idUsuario {
			saldo += trans.Monto
		} // SI EL USUARIO ENVIO SE RESTA Y SI RECIBE SE SUMA
		actual = actual.Sig
	}

	return saldo
}

func (bc *Blockchain) ValidarCadena() bool { // el hash anterior tiene que coincidir
	if bc.Cabeza == nil || bc.Cabeza.Sig == nil {
		return true // Si esta vacia o solo tiene 1 es valida
	}

	anterior := bc.Cabeza
	actual := bc.Cabeza.Sig

	for actual != nil {
		if actual.Bloque.HashPrevio != anterior.Bloque.Hash {
			return false //hay un fallo
		}

		// Se tiene que recalcular porque si alguien hubiera cambiado el contenido del bloque al recalcular el hash no coincidiria
		hashRecalculado := calcularHash(actual.Bloque.HashPrevio, actual.Bloque.Data, actual.Bloque.Timestamp)
		if actual.Bloque.Hash != hashRecalculado {
			return false
		}

		anterior = actual
		actual = actual.Sig
	}

	return true
}
