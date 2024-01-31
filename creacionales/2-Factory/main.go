package main

import (
	"fmt"
	"github.com/thebigyovadiaz/design-pattern-go/creacionales/2-Factory/factory"
	"os"
)

func main() {
	var t int
	fmt.Print("Conectar con: 1 Postgres - 2 MySQL >> ")

	_, err := fmt.Scan(&t)
	if err != nil {
		fmt.Printf("Error al leer la opción de conexión: %v", err)
		os.Exit(1)
	}

	conn := factory.Factory(t)
	if conn == nil {
		fmt.Println("Motor de DB no válido", t)
		os.Exit(1)
	}

	err = conn.Connect()
	if err != nil {
		fmt.Printf("No se pudo conectar a la DB: %v", err)
		os.Exit(1)
	}

	now, err := conn.GetNow()
	if err != nil {
		fmt.Printf("No se pudo consultar la fecha: %v", err)
		os.Exit(1)
	}

	fmt.Println(now.Format("2006-01-02"))
	err = conn.Close()
	if err != nil {
		fmt.Printf("No se pudo cerrar la conexión: %v", err)
	}
}
