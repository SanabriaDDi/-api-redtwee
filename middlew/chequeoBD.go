package middlew

import (
	"github.com/SanabriaDDi/api-redtwee/bd"
	"net/http"
)

/*ChequeoBD es el middleware que permite conocer el estado de la Base de Datos*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Conexión perdida con la Base de Datos", 500)
			return
		}

		next.ServeHTTP(w, r)
	}
}