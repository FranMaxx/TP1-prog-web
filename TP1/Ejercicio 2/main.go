package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

// Plantilla HTML del formulario (Punto 1 y 2)
const formularioHTML = `<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <title>Contacto</title>
</head>
<body>
    <h1>Formulario de Contacto</h1>
    {{if .Error}}
        <p style="color: red;"><strong>Error:</strong> {{.Error}}</p>
    {{end}}
    <!-- Formulario enviado con método POST hacia /contacto (Punto 1) -->
    <form action="/contacto" method="POST">
        <div>
            <label for="nombre">Nombre:</label><br>
            <input type="text" id="nombre" name="nombre" value="{{.Nombre}}">
        </div><br>
        <div>
            <label for="email">Email:</label><br>
            <input type="email" id="email" name="email" value="{{.Email}}">
        </div><br>
        <div>
            <label for="mensaje">Mensaje:</label><br>
            <textarea id="mensaje" name="mensaje" rows="4">{{.Mensaje}}</textarea>
        </div><br>
        <button type="submit">Enviar</button>
    </form>
</body>
</html>`

// Plantilla HTML de confirmación (Punto 2)
const confirmacionHTML = `<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <title>Confirmación</title>
</head>
<body>
    <h1>¡Datos Recibidos con Éxito!</h1>
    <p><strong>Nombre:</strong> {{.Nombre}}</p>
    <p><strong>Email:</strong> {{.Email}}</p>
    <p><strong>Mensaje:</strong> {{.Mensaje}}</p>
    <br>
    <a href="/contacto">Volver al formulario</a>
</body>
</html>`

// Estructura para pasar datos a la plantilla
type FormData struct {
	Nombre  string
	Email   string
	Mensaje string
	Error   string
}

func contactoHandler(w http.ResponseWriter, r *http.Request) {
	tmplForm := template.Must(template.New("form").Parse(formularioHTML))
	tmplConfirm := template.Must(template.New("confirm").Parse(confirmacionHTML))

	switch r.Method {
	// Mostrar el formulario en GET /contacto (Punto 2)
	case http.MethodGet:
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tmplForm.Execute(w, FormData{})

	// Procesar los datos en POST /contacto (Punto 2)
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
			return
		}

		nombre := strings.TrimSpace(r.FormValue("nombre"))
		email := strings.TrimSpace(r.FormValue("email"))
		mensaje := strings.TrimSpace(r.FormValue("mensaje"))

		// Validar que los campos no estén vacíos (Punto 2)
		if nombre == "" || email == "" || mensaje == "" {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusBadRequest)
			tmplForm.Execute(w, FormData{
				Nombre:  nombre,
				Email:   email,
				Mensaje: mensaje,
				Error:   "Todos los campos son obligatorios.",
			})
			return
		}

		// Mostrar los datos recibidos en la página de confirmación (Punto 2)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tmplConfirm.Execute(w, FormData{
			Nombre:  nombre,
			Email:   email,
			Mensaje: mensaje,
		})

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/contacto", contactoHandler)

	port := ":8080"
	fmt.Printf("Servidor escuchando en http://localhost%s/contacto\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
	}
}