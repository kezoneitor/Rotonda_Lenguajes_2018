package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//-------------------------------------------Variables Globales-------------------------------------------

var Rotonda []Calle
var RutaCQ []Carro
var RutaSC []Carro
var RutaSR []Carro
var Consola []string
var Historial []string

//-------------------------------------------Estructuras-------------------------------------------

type Carro struct {
	id        int
	velocidad int
	dest      string
	orig      string
	tTra      int
}

type Calle struct {
	entrada *[]Carro
	salida  string
	carro   Carro
}

//-------------------------------------------Funciones-------------------------------------------

/*
Funcion nombre: CrearCalle
	Variables	| tipo 		| E/S 	|
	------------+-----------+-------|
	entrada		| *[]Carro	| E		|
	salida		| string	| E		|
	carro		| Carro		| E		|
	calle		| Calle		| S		|

	return:
		Struct tipo Calle

	Descripcion:
		Crear un struct tipo calle con los parametros recibidos
*/
func CrearCalle(entrada *[]Carro, salida string, carro Carro) Calle {
	calle := Calle{entrada, salida, carro}
	return calle
}

/*
Funcion nombre: CrearCarro
	Variables	| tipo 		| E/S 	|
	------------+-----------+-------|
	id			| int		| E		|
	velocidad	| int		| E		|
	destino		| string	| E		|
	origen		| string	| E		|
	tTrayecto	| int		| E		|
	carro		| Carro		| S		|

	return:
		Struct tipo Carro

	Descripcion:
		Crear un struct tipo carro con los parametros recibidos
*/
func CrearCarro(id int, velocidad int, destino string, origen string, tTrayecto int) Carro {
	carro := Carro{id, velocidad, destino, origen, tTrayecto}
	return carro
}

/*
Funcion nombre: EntradaRotonda
	Variables	| tipo 		| E/S 	|
	------------+-----------+-------|
	calle		| *Calle	| E		|
	ruta		| string	| E		|

	return:
		-

	Descripcion:
		El metodo realiza una insercion a la rotonda si el campo esta vacio.
		(se envían mensajes a la lista Consola)
*/
func EntrarRotonda(calle *Calle, ruta string) {
	listaRuta := *calle.entrada

	if len(listaRuta) != 0 {
		if calle.carro.id != 1000 {
			MsgConsoleAdd("Carro " + strconv.Itoa(listaRuta[0].id) + " esperando a entrar desde la " + ruta)
		} else {
			calle.carro = listaRuta[0]
			MsgConsoleAdd("Carro " + strconv.Itoa(listaRuta[0].id) + " entró desde la " + ruta)
			listaRuta = listaRuta[1:]
			*calle.entrada = listaRuta
		}
	}
}

/*
Funcion nombre: MovRotonda
	Variables	| tipo 		| E/S 	|
	------------+-----------+-------|
	temp		| Carro		| -		|

	return:
		-

	Descripcion:
		Mueve los carros de la rotonda un campo hacia adelante

*/
func MovRotonda() {
	temp := Rotonda[0].carro
	Rotonda[0].carro = Rotonda[len(Rotonda)-1].carro
	for i := len(Rotonda) - 2; i > 0; i-- {

		Rotonda[i+1].carro = Rotonda[i].carro

	}
	Rotonda[1].carro = temp
}

/*
Funcion nombre: SalirRotonda
	Variables	| tipo 		| E/S 	|
	------------+-----------+-------|

	return:
		-

	Descripcion:
		Saca de la rotonda los autos que van hacia esa calle

*/
func SalirRotonda(calle *Calle, ruta string) {
	if calle.carro.dest == calle.salida {
		MsgConsoleAdd("Carro " + strconv.Itoa(calle.carro.id) + " salió por la " + ruta)
		calle.carro = CrearCarro(1000, 0, "---", "---", 0)
	}
}

/*
Funcion nombre: CarRand
	Variables	| tipo 		| E/S 	|
	------------+-----------+-------|
	id			| int		| E/S	|
	rutas		| []string	| E		|

	return:
		el siguiente id de carro

	Descripcion:
		Crea carros de manera aleatoria y los agrega de igual forma
		a las rutas por la que entrará
*/
func CarRand(id int, rutas []string) int {
	entrada := rand.Intn(3)
	salida := rand.Intn(3)
	velocidad := rand.Intn(1000) + 1
	if entrada == 0 {
		carRand := CrearCarro(id, velocidad, rutas[salida], rutas[entrada], 0)
		RutaCQ = append(RutaCQ, carRand)
	} else if entrada == 1 {
		carRand := CrearCarro(id, velocidad, rutas[salida], rutas[entrada], 0)
		RutaSC = append(RutaSC, carRand)
	} else {
		carRand := CrearCarro(id, velocidad, rutas[salida], rutas[entrada], 0)
		RutaSR = append(RutaSR, carRand)
	}
	return id + 1
}

/*
Funcion nombre: llenarRotonda
	Variables	| tipo 		| E/S 	|
	------------+-----------+-------|

	return:
		-

	Descripcion:
		Crear la rotonda, estos valores son por defecto por lo que
		siempre va a empezar de la misma manera
*/
func llenarRotonda() {
	Rotonda[0] = CrearCalle(nil, "RCQ", Carro{1000, 0, "---", "---", 0})
	Rotonda[1] = CrearCalle(nil, "nop", Carro{1000, 0, "---", "---", 0})
	Rotonda[2] = CrearCalle(&RutaCQ, "nop", Carro{1000, 0, "---", "---", 0})
	Rotonda[3] = CrearCalle(nil, "nop", Carro{1000, 0, "---", "---", 0})
	Rotonda[4] = CrearCalle(nil, "nop", Carro{1000, 0, "---", "---", 0})
	Rotonda[5] = CrearCalle(nil, "nop", Carro{1000, 0, "---", "---", 0})
	Rotonda[6] = CrearCalle(nil, "nop", Carro{1000, 0, "---", "---", 0})
	Rotonda[7] = CrearCalle(nil, "nop", Carro{1000, 0, "---", "---", 0})
	Rotonda[8] = CrearCalle(nil, "RSC", Carro{1000, 0, "---", "---", 0})
	Rotonda[9] = CrearCalle(nil, "nop", Carro{1000, 0, "---", "---", 0})
	Rotonda[10] = CrearCalle(&RutaSC, "nop", Carro{1000, 0, "---", "---", 0})
	Rotonda[11] = CrearCalle(nil, "nop", Carro{1000, 0, "---", "---", 0})
	Rotonda[12] = CrearCalle(nil, "RSR", Carro{1000, 0, "---", "---", 0})
	Rotonda[13] = CrearCalle(nil, "nop", Carro{1000, 0, "---", "---", 0})
	Rotonda[14] = CrearCalle(&RutaSR, "nop", Carro{1000, 0, "---", "---", 0})
	Rotonda[15] = CrearCalle(nil, "nop", Carro{1000, 0, "---", "---", 0})
}

/*
Funcion nombre: MsgConosolaAdd
	Variables	| tipo 		| E/S 	|
	------------+-----------+-------|
	msg			| string	| E		|

	return:
		-

	Descripcion:
		Agregar un mensaje a la lista Consola
*/
func MsgConsoleAdd(msg string) {
	Consola = append(Consola, msg)
	Historial = append(Historial, msg)
}

/*
Funcion nombre: MsgConosoleShow
	Variables	| tipo 		| E/S 	|
	------------+-----------+-------|
	msg			| string	| S		|

	return:
		Regresa el primer mensaje ingresado a la lista Consola

	Descripcion:
		Retorna el primer mensaje de la lista consola para ser
		mostrado.
*/
func MsgConsoleShow() string {
	if len(Consola) != 0 {
		msg := Consola[0]
		Consola = Consola[1:]
		return msg
	}
	return "Consola vacia"
}

/*
Funcion nombre: impRotonda
	Variables	| tipo 		| E/S 	|
	------------+-----------+-------|
	r			| []Calle	| E
	return:
		-

	Descripcion:
		Muestra la rotonda con sus valores actuales
*/
func impRotonda(r []Calle) {
	fmt.Println("X x x x x x x X x x x x x x X| | | | | | |X x x x x x x X| | | | | | |X x x x x x x X x x x x x x X")
	fmt.Println("x v v v v v v x v v v v 0   x    " + strconv.Itoa(r[14].carro.id) + "id   x    " + strconv.Itoa(r[13].carro.id) + "id   x    " + strconv.Itoa(r[12].carro.id) + "id   x   0 v v v v x v v v v v v x")
	fmt.Println("x v v v v v v x v v v 0     x    " + r[14].carro.dest + "des   x    " + r[13].carro.dest + "des   x    " + r[12].carro.dest + "des   x     0 v v v x v v v v v v x")
	fmt.Println("x v v v v v v x v v 0       x    " + r[14].carro.orig + "ori   x    " + r[13].carro.orig + "ori   x    " + r[12].carro.orig + "ori   x       0 v v x v v v v v v x")
	fmt.Println("x v v v v v v x v 0         x             x             x             x         0 v x v v v v v v x")
	fmt.Println("x v v v v v v x 0           x             x             x             x           0 x v v v v v v x")
	fmt.Println("X x x x x x x X x x x x x x X x x x x x x X x x x x x x X x x x x x x X x x x x x x X x x x x x x X")
	fmt.Println("x v v v v v 0 x    " + strconv.Itoa(r[15].carro.id) + "id   x         0 v x v v v v v v x v 0         x    " + strconv.Itoa(r[11].carro.id) + "id   x 0 v v v v v x")
	fmt.Println("x v v v v 0   x    " + r[15].carro.dest + "des   x       0 v v x v v v v v v x v v 0       x    " + r[11].carro.dest + "des   x   0 v v v v x")
	fmt.Println("x v v v 0     x    " + r[15].carro.orig + "ori   x     0 v v v x v v v v v v x v v v 0     x    " + r[11].carro.orig + "ori   x     0 v v v x")
	fmt.Println("x v v 0       x             x   0 v v v v x v v v v v v x v v v v 0   x             x       0 v v x")
	fmt.Println("x v 0         x             x 0 v v v v v x v v v v v v x v v v v v 0 x             x         0 v x")
	fmt.Println("X x x x x x x X x x x x x x X x x x x x x X x x x x x x X x x x x x x X x x x x x x X x x x x x x X")
	fmt.Println("_    " + strconv.Itoa(r[0].carro.id) + "id   x           0 x v v v v v v x v v v v v v x v v v v v v x 0           x    " + strconv.Itoa(r[10].carro.id) + "id   _")
	fmt.Println("_    " + r[0].carro.dest + "des   x         0 v x v v v v v v x v v v v v v x v v v v v v x v 0         x    " + r[10].carro.dest + "des   _")
	fmt.Println("_    " + r[0].carro.orig + "ori   x       0 v v x v v v v v v x v v v v v v x v v v v v v x v v 0       x    " + r[10].carro.orig + "ori   _")
	fmt.Println("_             x     0 v v v x v v v v v v x v v v v v v x v v v v v v x v v v 0     x             _")
	fmt.Println("_             x   0 v v v v x v v v v v v x v v v v v v x v v v v v v x v v v v 0   x             _")
	fmt.Println("X x x x x x x X x x x x x x X x x x x x x X x x x x x x X x x x x x x X x x x x x x X x x x x x x X")
	fmt.Println("x    " + strconv.Itoa(r[1].carro.id) + "id   x v v v v v v x v v v v v v x v v v v v v x v v v v v v x v v v v v v x    " + strconv.Itoa(r[9].carro.id) + "id   x")
	fmt.Println("x    " + r[1].carro.dest + "des   x v v v v v v x v v v v v v x v v v v v v x v v v v v v x v v v v v v x    " + r[9].carro.dest + "des   x")
	fmt.Println("x    " + r[1].carro.orig + "ori   x v v v v v v x v v v v v v x v v v v v v x v v v v v v x v v v v v v x    " + r[9].carro.orig + "ori   x")
	fmt.Println("x             x v v v v v v x v v v v v v x v v v v v v x v v v v v v x v v v v v v x             x")
	fmt.Println("x             x v v v v v v x v v v v v v x v v v v v v x v v v v v v x v v v v v v x             x")
	fmt.Println("X x x x x x x X x x x x x x X x x x x x x X x x x x x x X x x x x x x X x x x x x x X x x x x x x X")
	fmt.Println("_    " + strconv.Itoa(r[2].carro.id) + "id   x   0 v v v v x v v v v v v x v v v v v v x v v v v v v x v v v v 0   x    " + strconv.Itoa(r[8].carro.id) + "id   _")
	fmt.Println("_    " + r[2].carro.dest + "des   x     0 v v v x v v v v v v x v v v v v v x v v v v v v x v v v 0     x    " + r[8].carro.dest + "des   _")
	fmt.Println("_    " + r[2].carro.orig + "ori   x       0 v v x v v v v v v x v v v v v v x v v v v v v x v v 0       x    " + r[8].carro.orig + "ori   _")
	fmt.Println("_             x         0 v x v v v v v v x v v v v v v x v v v v v v x v 0         x             _")
	fmt.Println("_             x           0 x v v v v v v x v v v v v v x v v v v v v x 0           x             _")
	fmt.Println("X x x x x x x X x x x x x x X x x x x x x X x x x x x x X x x x x x x X x x x x x x X x x x x x x X")
	fmt.Println("x v 0         x    " + strconv.Itoa(r[3].carro.id) + "id   x 0 v v v v v x v v v v v v x v v v v v 0 x    " + strconv.Itoa(r[7].carro.id) + "id   x         0 v x")
	fmt.Println("x v v 0       x    " + r[3].carro.dest + "des   x   0 v v v v x v v v v v v x v v v v 0   x    " + r[7].carro.dest + "des   x       0 v v x")
	fmt.Println("x v v v 0     x    " + r[3].carro.orig + "ori   x     0 v v v x v v v v v v x v v v 0     x    " + r[7].carro.orig + "ori   x     0 v v v x")
	fmt.Println("x v v v v 0   x             x       0 v v x v v v v v v x v v 0       x             x   0 v v v v x")
	fmt.Println("x v v v v v 0 x             x         0 v x v v v v v v x v 0         x             x 0 v v v v v x")
	fmt.Println("X x x x x x x X x x x x x x X x x x x x x X x x x x x x X x x x x x x X x x x x x x X x x x x x x X")
	fmt.Println("x v v v v v v x 0           x    " + strconv.Itoa(r[4].carro.id) + "id   x    " + strconv.Itoa(r[5].carro.id) + "id   x    " + strconv.Itoa(r[6].carro.id) + "id   x           0 x v v v v v v x")
	fmt.Println("x v v v v v v x v 0         x    " + r[4].carro.dest + "des   x    " + r[5].carro.dest + "des   x    " + r[6].carro.dest + "des   x         0 v x v v v v v v x")
	fmt.Println("x v v v v v v x v v 0       x    " + r[4].carro.orig + "ori   x    " + r[5].carro.orig + "ori   x    " + r[6].carro.orig + "ori   x       0 v v x v v v v v v x")
	fmt.Println("x v v v v v v x v v v 0     x             x             x             x     0 v v v x v v v v v v x")
	fmt.Println("x v v v v v v x v v v v 0   x             x             x             x   0 v v v v x v v v v v v x")
	fmt.Println("X x x x x x x X x x x x x x X x x x x x x X x x x x x x X x x x x x x X x x x x x x X x x x x x x X")
	for i := 0; i < len(Consola); i++ {
		fmt.Println(MsgConsoleShow())
	}
}

//-------------------------------------------Main-------------------------------------------
func main() {

	//Inicializar las variables globales
	Rotonda = make([]Calle, 16, 16)
	RutaCQ = make([]Carro, 0, 30)
	RutaSC = make([]Carro, 0, 30)
	RutaSR = make([]Carro, 0, 30)
	Consola = make([]string, 0, 30)
	Historial = make([]string, 0, 200)

	//Iniciar variables a utilizar en el proyecto
	var cant int                           // cantidad de carros para el metodo
	id := 1001                             // "Placa" de cada carro
	rutas := []string{"RCQ", "RSC", "RSR"} // Rutas a asignar
	//	Crear carros random para ingresar en cada lista
	fmt.Println("Cantidad de carros Aleatorios: ")
	fmt.Scanf("%d", &cant)
	for n := 0; n < cant; n++ {
		id = CarRand(id, rutas)
	}

	//Llenar la rotonda por defecto
	llenarRotonda()
	impRotonda(Rotonda)
	/*
		Func Anoni 	| tiempos de ejecucción
					| x1	| x2	| x 4	| x10
		------------+-------+-------+-------+------
		Salir		| 1s	| 2s	| 4s	| 10s
		Mover		| 2s	| 4s	| 8s	| 20s
		Entrar		| 0.5s	| 1s	| 2s	| 5s
		Consola		| 0.5s	| 1s	| 2s	| 5s
		MsgShow		| 0.5s	| 1s	| 2s	| 5s


	*/
	tS := 2000 * time.Millisecond
	tM := 4000 * time.Millisecond
	tE := 2000 * time.Millisecond

	go func() {
		for {
			time.Sleep(tS)

			go SalirRotonda(&Rotonda[0], "Ruta CQ")

			go SalirRotonda(&Rotonda[8], "Ruta SC")

			go SalirRotonda(&Rotonda[12], "Ruta SR")

			impRotonda(Rotonda)
		}
	}()

	go func() {
		x := 1
		for {

			time.Sleep(tM)
			MovRotonda()
			impRotonda(Rotonda)
			x++
		}
	}()

	go func() {
		for {
			time.Sleep(tE)

			go EntrarRotonda(&Rotonda[2], "Ruta CQ")

			go EntrarRotonda(&Rotonda[10], "Ruta SC")

			go EntrarRotonda(&Rotonda[14], "Ruta SR")

			impRotonda(Rotonda)

		}
	}()

	ok := 1
	for {
		fmt.Scanf("%d", &ok)
		if ok != 1 {
			break
		}
	}

	for _, msg := range Historial {
		fmt.Println(msg)
	}
}
