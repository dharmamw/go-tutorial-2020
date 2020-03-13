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
	DeleteUserByNIP(ctx context.Context, NIP string) error
	GetUserFromFireBase(ctx context.Context) ([]userEntity.User, error)
	InsertUsersToFirebase(ctx context.Context, user userEntity.User) error
	InsertMany(ctx context.Context, userList []userEntity.User) error
	PublishUser(user userEntity.User) error
	UpdateByNipFirebase(ctx context.Context, nip string, user userEntity.User) error
	DeleteByNipFirebase(ctx context.Context, nip string) error
	//DeleteAllFirebase(ctx context.Context) error
	GetUserClient(ctx context.Context, headers http.Header) ([]userEntity.User, error)
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
		var _type string
		if _, getOK := r.URL.Query()["get"]; getOK {
			_type = r.FormValue("get")
		}
		switch _type {
		case "sqlall":
			result, err = h.userSvc.GetAllUsers(context.Background())
		case "firebaseall":
			result, err = h.userSvc.GetUserFromFireBase(context.Background())
		case "sqlNIP":
			result, err = h.userSvc.GetUserByNIP(context.Background(), r.FormValue("NIP"))
		case "userClient":
			result, err = h.userSvc.GetUserClient(context.Background(), r.Header)
		}

	case http.MethodPost:

		var (
			_type    string
			userList []userEntity.User
		)
		if _, postOK := r.URL.Query()["insert"]; postOK {
			_type = r.FormValue("insert")
		}
		switch _type {
		case "firebase":
			json.Unmarshal(body, &user)
			err = h.userSvc.InsertUsersToFirebase(context.Background(), user)
		case "sql":
			json.Unmarshal(body, &user)
			err = h.userSvc.InsertUsers(context.Background(), user)
		case "many":
			json.Unmarshal(body, &userList)
			err = h.userSvc.InsertMany(context.Background(), userList)
		case "kafka":
			json.Unmarshal(body, &user)
			err = h.userSvc.PublishUser(user)
		}

	case http.MethodPut:
		json.Unmarshal(body, &user)

		var _type string

		if _, putOK := r.URL.Query()["update"]; putOK {
			_type = r.FormValue("update")
		}
		switch _type {
		case "firebase":
			err = h.userSvc.UpdateByNipFirebase(context.Background(), r.FormValue("NIP"), user)
		}

		//result, err = h.userSvc.UpdateUserByNIP(context.Background(), r.FormValue("NIP"), user)

	case http.MethodDelete:
		// usersDelete?Delete=&NIP=
		var _type string
		if _, deleteOK := r.URL.Query()["Delete"]; deleteOK {
			_type = r.FormValue("Delete")
		}
		switch _type {
		case "sql":
			err = h.userSvc.DeleteUserByNIP(context.Background(), r.FormValue("NIP"))
		case "firebase":
			err = h.userSvc.DeleteByNipFirebase(context.Background(), r.FormValue("NIP"))
			// case "firebaseall":
			// 	err = h.userSvc.DeleteAllFirebase(context.Background())
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
