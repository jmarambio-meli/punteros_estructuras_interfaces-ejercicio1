/*Ejercicio 1 - Registro de estudiantes
Una universidad necesita registrar a los/as estudiantes y
generar una funcionalidad para imprimir el detalle de los datos de
cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos
brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables
Nombre, Apellido, DNI, Fecha y que tenga un método detalle
*/

package main

import "fmt"

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    string
}

func (a Alumno) detalle() {
	fmt.Printf(
		"El nombre del estudiante es: %s, %s, DNI: %s, Fecha de ingreso: %s\n", a.Nombre, a.Apellido, a.DNI, a.Fecha,
	)
}

func main() {
	alumno1 := Alumno{"Jean", "Marambio", "20777519", "19/12/2022"}
	alumno1.detalle()
}
