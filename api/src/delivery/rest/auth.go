package rest

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/laster18/poi/api/src/config"
	"github.com/laster18/poi/api/src/delivery"
	"github.com/olahol/go-imageupload"
	"github.com/pkg/errors"
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

	w.Header().Set("location", fmt.Sprintf("%s/login", config.Conf.FrontBaseURL))
	w.WriteHeader(http.StatusMovedPermanently)
}

// parameters
// name: string
// image: file
func guestLoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("name")

	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "required name parameter")
		return
	}

	if len(username) > 12 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "name is max 12 characters")
		return
	}

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "internal server error")
		return
	}

	uploadedAvatarURL, err := saveImageToLocal(r, "image")
	if err != nil {
		log.Println("failed to save image to local, err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "internal server error")
	}

	userSession, err := delivery.GetUserSession(r)
	if err != nil {
		log.Printf("failed to get user session, the cause was %v", err)
		handleInvalidSessionErr(w, err)
		return
	}

	userSession.SetUser(&delivery.User{
		ID:        uuid.NewString(),
		Name:      username,
		AvatarURL: uploadedAvatarURL,
	})
	if err := userSession.Save(r, w); err != nil {
		log.Print("failed to set user to session err:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	handleRedirectRoot(w)
}

func saveImageToLocal(r *http.Request, paramname string) (string, error) {
	img, err := imageupload.Process(r, "image")

	if err != nil {
		return "", err
	}

	thumb, err := imageupload.ThumbnailPNG(img, 300, 300)

	if err != nil {
		return "", errors.Wrap(err, "failed to resize iamge")
	}

	uploadLocalPath := "static/" + uuid.NewString() + ".png"

	if err := ioutil.WriteFile(uploadLocalPath, thumb.Data, 0600); err != nil {
		return "", errors.Wrap(err, "failed to write image")
	}

	return "http://localhost:" + config.Conf.Port + "/" + uploadLocalPath, nil
}

func uploadImageToS3() {
	panic("todo implement")
}
