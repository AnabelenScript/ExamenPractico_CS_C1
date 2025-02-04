package main

import (
	"examenPractico/server"
	"log"
)

func main() {
	go func() {
		log.Println("Iniciando servidor principal en el puerto 8080...")
		server.StartServer("8080")
	}()

	go func() {
		log.Println("Iniciando servidor de replicación en el puerto 8081...")
		server.StartServer("8081")
	}()

	select {} 
}
