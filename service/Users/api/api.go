package api

import (
	"Saavedra/service/Users/service"
	"Saavedra/service/Users/types"
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type EndpointHandler struct {
	service service.Service
}

func New(s service.Service) *EndpointHandler {
	return &EndpointHandler{service: s}
}

// HELPER -> DECODES JSON TO STRUCT PARSING ID FROM STR TO INT
// RETURNS User POINTER.
func parseStruct(r *http.Request) (*types.User, error) {
	var userStr types.UserStr
	err := json.NewDecoder(r.Body).Decode(&userStr)
	if err != nil {
		return nil, err
	}

	idInt, err := strconv.Atoi(userStr.Id)
	if err != nil {
		return nil, err
	}

	user := types.User{
		Id:       idInt,
		Name:     userStr.Name,
		Email:    userStr.Email,
		Password: userStr.Password,
	}
	return &user, nil
}

// ROUTE: /users -> ONLY RENDERS HTML
func (e EndpointHandler) CallUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		tpl, err := template.ParseFiles("service/Users/views/users.html")
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Couldn't parse users HTML template file.", http.StatusInternalServerError)
			return
		}

		err = tpl.Execute(w, nil)
		if err != nil {
			log.Print("Error when rendering users template.")
			http.Error(w, "Render Error", http.StatusInternalServerError)
		}

	default:
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}
}

// ROUTE: /users/records?page=x -> ONLY RETURNS TYPE(TABLE) -> RECORDS & METADATA
func (e EndpointHandler) CallUsersRecords(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		page := r.URL.Query().Get("page")
		if page != "" {
			page = "1"
		}

		intPage, err := strconv.Atoi(page)
		if err != nil {
			intPage = 1
		}

		data, err := e.service.ListUsers(intPage)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Fetching users data error...", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)

	default:
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return

	}

}

// ROUTE: /users/new -> REDIRECTS -> /users/record?id=x
func (e EndpointHandler) CreateUserInDB(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		tpl, err := template.ParseFiles("service/Users/views/usersCreate.html")

		if err != nil {
			log.Print("Error parsing usersCreate HTML template...")
			http.Error(w, "Parsing error...", http.StatusInternalServerError)
			return
		}

		err = tpl.Execute(w, nil)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Loading error...", http.StatusInternalServerError)
			return
		}

	case http.MethodPost:
		var userRequest types.NoIdUser
		err := json.NewDecoder(r.Body).Decode(&userRequest)
		if err != nil {
			log.Printf("Error parsing struct '/users/new'...(%v)", err.Error())
			http.Error(w, "Error parsing struct '/users/new'.", http.StatusInternalServerError)
			return
		}
		user := types.User{
			Id:       0,
			Name:     userRequest.Name,
			Email:    userRequest.Email,
			Password: userRequest.Password,
		}

		newUser, err := e.service.NewUser(&user)
		if err != nil {
			log.Printf("Error creating record 'users/new'...(%v)", err.Error())
			http.Error(w, "Error creating record", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(newUser)

	case http.MethodHead:
		w.WriteHeader(http.StatusOK)

	default:
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}
}

// ROUTE /users/record?id=x
func (e EndpointHandler) CallUsersRecord(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		// GET USER ID
		strId := r.URL.Query().Get("id")
		id, err := strconv.Atoi(strId)
		if err != nil {
			log.Printf("Invalid 'user id': (%v)...", strId)
			http.Redirect(w, r, "/users", http.StatusSeeOther)
			return
		}

		// QUERY USER DATA
		user, err := e.service.SelectUser(id)
		if err != nil {
			log.Printf("usersRecord query error... (%v)", err.Error())
			http.Error(w, "userRecord error", http.StatusInternalServerError)
			return
		}

		// PARSE HTML TEMPLATE
		tpl, err := template.ParseFiles("service/Users/views/usersRecord.html")
		if err != nil {
			log.Printf("usersRecord parse error... (%v)", err.Error())
			http.Error(w, "userRecord error", http.StatusInternalServerError)
			return
		}

		err = tpl.Execute(w, user)
		if err != nil {
			log.Printf("usersRecord load error... (%v)", err.Error())
			http.Error(w, "userRecord error", http.StatusInternalServerError)
			return
		}

	case http.MethodPut:
		user, err := parseStruct(r)
		if err != nil {
			log.Printf("Error parsing struct... (%v)", err.Error())
			http.Error(w, "Error parsing struct", http.StatusInternalServerError)
			return
		}

		updUser, err := e.service.UpdateUser(user)
		if err != nil {
			log.Printf("usersRecord query error... (%v)", err.Error())
			http.Error(w, "usersRecord json errro", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updUser)

	case http.MethodDelete:
		user, err := parseStruct(r)
		if err != nil {
			log.Printf("Error parsing struct... (%v)", err.Error())
			http.Error(w, "Error parsing struct", http.StatusInternalServerError)
			return
		}
		err = e.service.DeleteUser(user)
		if err != nil {
			if errors.Is(err, types.ErrAdminProtection) {
				log.Printf("Impossible delete admin user: %v", user.Id)
				http.Error(w, "Cannot delete standard admin user", http.StatusForbidden)
				return
			}
			log.Printf("Error deleting record: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

	case http.MethodHead:
		w.WriteHeader(http.StatusOK)

	default:
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}
}
