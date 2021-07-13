package rest

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"unicode/utf8"

	"github.com/google/uuid"
	"github.com/laster18/poi/api/src/config"
	"github.com/laster18/poi/api/src/domain/user"
	"github.com/laster18/poi/api/src/util/session"
	"github.com/olahol/go-imageupload"
)

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	userSession, err := session.GetUserSession(r)
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
func guestLoginHandler(userRepo user.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("recevied guestLogin!!!")

		username := r.FormValue("name")

		if username == "" {
			handleValidationErr(w, errors.New("required name parameter"))
			return
		}

		if utf8.RuneCountInString(username) > 12 {
			handleValidationErr(w, errors.New("name is max 12 characters"))
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
			if err == errNoFile {
				handleValidationErr(w, errors.New("required image"))
				return
			}

			log.Println("failed to save image to local, err:", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "internal server error")
		}

		userSession, err := session.GetUserSession(r)
		if err != nil {
			log.Printf("failed to get user session, the cause was %v", err)
			handleInvalidSessionErr(w, err)
			return
		}

		user := &user.User{
			UID:       uuid.NewString(),
			Name:      username,
			AvatarURL: uploadedAvatarURL,
			Provider:  user.ProviderGuest,
		}
		if err := userRepo.Save(context.Background(), user); err != nil {
			log.Print("failed to save user err:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		userSession.SetUser(user)
		if err := userSession.Save(r, w); err != nil {
			log.Print("failed to set user to session err:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}
}

var (
	errNoFile = errors.New("no file error")
)

func saveImageToLocal(r *http.Request, paramname string) (string, error) {
	img, err := imageupload.Process(r, "image")

	if err != nil {
		return "", errNoFile
	}

	thumb, err := imageupload.ThumbnailPNG(img, 300, 300)

	if err != nil {
		return "", fmt.Errorf("failed to resize iamge: %w", err)
	}

	uploadLocalPath := "static/" + uuid.NewString() + ".png"

	if err := ioutil.WriteFile(uploadLocalPath, thumb.Data, 0600); err != nil {
		return "", fmt.Errorf("failed to write image: %w", err)
	}

	return "http://localhost:" + config.Conf.Port + "/" + uploadLocalPath, nil
}

func uploadImageToS3() {
	panic("todo implement")
}
