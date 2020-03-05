package user

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	userEntity "go-tutorial-2020/internal/entity/user"
	"go-tutorial-2020/pkg/response"
)

// IUserSvc is an interface to User Service
type IUserSvc interface {
	InsertUsers(ctx context.Context, user userEntity.User) error
	GetAllUsers(ctx context.Context) ([]userEntity.User, error)
	GetUserByNIP(ctx context.Context, NIP string) (userEntity.User, error)
	UpdateUserByNIP(ctx context.Context, NIP string, user userEntity.User) (userEntity.User, error)
	DeleteUserByNIP(ctx context.Context, NIP string, user userEntity.User) (userEntity.User, error)
	GetUserFromFireBase(ctx context.Context) ([]userEntity.User, error)
}

type (
	// Handler ...
	Handler struct {
		userSvc IUserSvc
	}
)

// New for user domain handler initialization
func New(is IUserSvc) *Handler {
	return &Handler{
		userSvc: is,
	}
}

// UserHandler will return user data
func (h *Handler) UserHandler(w http.ResponseWriter, r *http.Request) {
	var (
		resp     *response.Response
		metadata interface{}
		result   interface{}
		err      error
		errRes   response.Error
		user     userEntity.User
	)
	// Make new response object
	resp = &response.Response{}
	body, _ := ioutil.ReadAll(r.Body)
	// Defer will be run at the end after method finishes
	defer resp.RenderJSON(w, r)

	switch r.Method {
	// Check if request method is GET
	case http.MethodGet:
		// Cek query di URL
		_, nipOK := r.URL.Query()["NIP"]
		if nipOK {
			// Ambil user by NIP
			result, err = h.userSvc.GetUserByNIP(context.Background(), r.FormValue("NIP"))
		} else {
			// Ambil semua data user
			if _, fireOK := r.URL.Query()["firebasedb"]; fireOK {
				result, err = h.userSvc.GetUserFromFireBase(context.Background())
			}
			result, err = h.userSvc.GetAllUsers(context.Background())

		}

	case http.MethodPost:
		json.Unmarshal(body, &user)
		err = h.userSvc.InsertUsers(context.Background(), user)

	case http.MethodPut:
		_, nipOK := r.URL.Query()["NIP"]
		if nipOK {
			result, err = h.userSvc.GetUserByNIP(context.Background(), r.FormValue("NIP"))
		}
		json.Unmarshal(body, &user)
		_, updateOK := r.URL.Query()["NIP"]
		if updateOK {
			result, err = h.userSvc.UpdateUserByNIP(context.Background(), r.FormValue("NIP"), user)
		}
	case http.MethodDelete:

		json.Unmarshal(body, &user)
		_, deleteOK := r.URL.Query()["NIP"]
		if deleteOK {
			result, err = h.userSvc.DeleteUserByNIP(context.Background(), r.FormValue("NIP"), user)
		}

	default:
		err = errors.New("400")
	}

	// If anything from service or data return an error
	if err != nil {
		// Error response handling
		errRes = response.Error{
			Code:   101,
			Msg:    "Data Not Found",
			Status: true,
		}
		// If service returns an error
		if strings.Contains(err.Error(), "service") {
			// Replace error with server error
			errRes = response.Error{
				Code:   201,
				Msg:    "Failed to process request due to server error",
				Status: true,
			}
		}

		// Logging
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		resp.Error = errRes
		return
	}

	// Inserting data to response
	resp.Data = result
	resp.Metadata = metadata
	// Logging
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
}
