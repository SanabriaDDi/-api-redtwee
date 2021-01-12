package routers

import (
	"encoding/json"
	"github.com/SanabriaDDi/api-redtwee/bd"
	"github.com/SanabriaDDi/api-redtwee/models"
	"net/http"
)

/*Registro es la función para crear en la BD el registro del usuario*/
func Registro(w http.ResponseWriter, r *http.Request)  {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
	}
	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de al menos 6 caracteres", 400)
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya esxiste un usuario registrado con ese email", 400)
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro de usuario "+err.Error(), 400)
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario", 400)
	}

	w.WriteHeader(http.StatusCreated)
}
