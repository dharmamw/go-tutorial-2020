package arif

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	userEntity "tugas-arif/internal/entity/arif"
	"tugas-arif/pkg/response"
)

// IArifSvc is an interface to Arif Service
// Masukkan function dari service ke dalam interface ini
type IArifSvc interface {
	UpdateUserByID(ctx context.Context, ID string, user userEntity.User) (userEntity.User, error)
	DeleteUserByID(ctx context.Context, ID string) error
	InsertUser(ctx context.Context, user userEntity.User) error
	GetAllUser(ctx context.Context) ([]userEntity.User, error)
	GetUserByID(ctx context.Context, ID string) (userEntity.User, error)
	GetPrintPage(ctx context.Context, page int, length int) ([]userEntity.User, error)
}

type (
	// Handler ...
	Handler struct {
		arifSvc IArifSvc
	}
)

// New for bridging product handler initialization
func New(is IArifSvc) *Handler {
	return &Handler{
		arifSvc: is,
	}
}

// ArifHandler will receive request and return response . . .
func (h *Handler) ArifHandler(w http.ResponseWriter, r *http.Request) {
	var (
		resp     *response.Response
		result   interface{}
		metadata interface{}
		err      error
		errRes   response.Error
		user     userEntity.User
		page     int
		length   int
	)
	resp = &response.Response{}
	body, _ := ioutil.ReadAll(r.Body)

	defer resp.RenderJSON(w, r)

	switch r.Method {
	// Check if request method is GET
	case http.MethodGet:
		// Cek query di URL
		var _type string

		if _, idOK := r.URL.Query()["Get"]; idOK {
			_type = r.FormValue("Get")
		}
		switch _type {
		case "sqlID":
			//json.Unmarshal(body, &user)
			result, err = h.arifSvc.GetUserByID(context.Background(), r.FormValue("ID"))
		case "PrintPage":
			page, err = strconv.Atoi(r.FormValue("page"))
			length, err = strconv.Atoi(r.FormValue("length"))
			result, err = h.arifSvc.GetPrintPage(context.Background(), page, length)
		case "sqlall":
			//json.Unmarshal(body, &user)
			result, err = h.arifSvc.GetAllUser(context.Background())
		}

	// Check if request method is POST
	case http.MethodPost:

		var (
			_type string
		)
		if _, fireOk := r.URL.Query()["Insert"]; fireOk {
			_type = r.FormValue("Insert")
		}
		switch _type {

		case "sql":
			json.Unmarshal(body, &user)
			err = h.arifSvc.InsertUser(context.Background(), user)
		}
	// Check if request method is PUT
	case http.MethodPut:
		json.Unmarshal(body, &user)
		_, updateOK := r.URL.Query()["ID"]
		if updateOK {
			result, err = h.arifSvc.UpdateUserByID(context.Background(), r.FormValue("ID"), user)
		}

	// Check if request method is DELETE
	case http.MethodDelete:
		json.Unmarshal(body, &user)
		_, deleteOK := r.URL.Query()["ID"]
		if deleteOK {
			err = h.arifSvc.DeleteUserByID(context.Background(), r.FormValue("ID"))
		}

		var (
			_type string
			user  userEntity.User
		)
		if _, fireOk := r.URL.Query()["Insert"]; fireOk {
			_type = r.FormValue("Insert")
		}
		switch _type {
		case "sql":
			json.Unmarshal(body, &user)
			err = h.arifSvc.InsertUser(context.Background(), user)
		}

	default:
		err = errors.New("404")
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

		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		resp.Error = errRes
		return
	}

	resp.Data = result
	resp.Metadata = metadata
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
}
