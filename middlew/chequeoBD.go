package middlew

import (
	"net/http"

	"github.com/jaguirre728/red_social.git/bd"
)

/*ChequeoBD middelware que chequea la conexión a la bd*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexión perdida con la Base de Datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
