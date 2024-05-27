package main

import (
	"fmt"
	"sync"
)

// Variables globales
var turno = "hombres"           // Indica de quién es el turno (inicialmente hombres)
var hombres = 0                 // Contador de hombres que han entrado
var mujeres = 0                 // Contador de mujeres que han entrado
var auxiliar = make([]bool, 10) // Array para verificar si un ID específico ya ha entrado

// Función que representa a un hombre intentando entrar al Rotor
func hombre(id int, cedula string, m *sync.Mutex) {
	for {
		m.Lock() // Adquiere el bloqueo del mutex para exclusión mutua

		if turno == "hombres" && !auxiliar[id] {
			// Si es el turno de los hombres y este hombre no ha entrado
			hombres++
			auxiliar[id] = true // Marca que este hombre ya ha entrado
			fmt.Printf("Hombre con cédula: %s está entrando al Rotor.\n", cedula)
		}

		if hombres > 9 {
			// Si 10 hombres han entrado, resetea el contador y el array auxiliar
			fmt.Println()
			hombres = 0
			auxiliar = make([]bool, 10)
			turno = "mujeres" // Cambia el turno a mujeres
		}

		m.Unlock() // Libera el bloqueo del mutex
	}
}

// Función que representa a una mujer intentando entrar al Rotor
func mujer(id int, cedula string, m *sync.Mutex) {
	for {
		m.Lock() // Adquiere el bloqueo del mutex para exclusión mutua

		if turno == "mujeres" && !auxiliar[id] {
			// Si es el turno de las mujeres y esta mujer no ha entrado
			mujeres++
			auxiliar[id] = true // Marca que esta mujer ya ha entrado
			fmt.Printf("Mujer con cédula: %s está entrando al Rotor.\n", cedula)
		}

		if mujeres > 9 {
			// Si 10 mujeres han entrado, resetea el contador y el array auxiliar
			fmt.Println()
			mujeres = 0
			auxiliar = make([]bool, 10)
			turno = "hombres" // Cambia el turno a hombres
		}

		m.Unlock() // Libera el bloqueo del mutex
	}
}

func main() {
	m := new(sync.Mutex) // Crea un nuevo mutex

	// Arrays con las cédulas de hombres y mujeres
	hombresCedula := [10]string{"4260866621", "4238919234", "5935507212", "9161328856", "7098612938", "1050361178", "6640608712", "2605717910", "6008472523", "7346755923"}
	mujeresCedula := [10]string{"0480239123", "6113403321", "8912399478", "1282513345", "1898616056", "4106773589", "6358033210", "4272866912", "9185910834", "9333883132"}

	for id := 0; id < 10; id++ {
		// Lanza una goroutine para cada hombre y cada mujer
		go hombre(id, hombresCedula[id], m)
		go mujer(id, mujeresCedula[id], m)
	}

	var esperar string
	fmt.Scanln(&esperar) // Espera a que el usuario presione Enter para evitar que el programa termine inmediatamente
}
