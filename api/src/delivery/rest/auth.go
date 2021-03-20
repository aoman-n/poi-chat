package rest

import (
	"fmt"
	"net/http"

	"github.com/laster18/poi/api/src/delivery"
)

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	userSession, err := delivery.GetUserSession(r)
	if err != nil {
		handleInvalidSessionErr(w, err)
		return
	}

	if err := userSession.RemoveUser(r, w); err != nil {
		handleSaveOrRemoveSessionErr(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Success!!")
}
