package main

import (
	"fmt"
	"net/http"
)

func main() {
	staticDir := "./static"
	fileServer := http.FileServer(http.Dir(staticDir))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 1. Validar que la petición sea únicamente GET (Punto 1)
		if r.Method != http.MethodGet {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		// 2. Establecer el encabezado Content-Type (Punto 1)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		// 3. Control de Rutas (Puntos 1 y 2)
		switch r.URL.Path {
		    case "/":
			    fileServer.ServeHTTP(w, r)

		    case "/about":
			    r.URL.Path = "/about.html"
			    fileServer.ServeHTTP(w, r)

		    default:
			    http.NotFound(w, r)
		}
	})

	port := ":8080"
	fmt.Printf("Servidor escuchando en http://localhost%s\n", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
	}
}