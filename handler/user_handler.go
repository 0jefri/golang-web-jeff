package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-web/common"
	"github.com/golang-web/model"
	"github.com/golang-web/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var payload model.User

	payload.ID = common.GenerateUUID()

	// Decode JSON request body ke dalam struct User
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Println(err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Panggil metode RegisterNewUser dari UserService
	if err := h.userService.RegisterNewUser(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Jika berhasil, kirimkan response sukses
	response := model.Response{
		StatusCode: http.StatusCreated,
		Message:    "User registered successfully",
		Data:       payload,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")

	// Validasi ID
	idInt, err := strconv.Atoi(id)
	if err != nil || idInt <= 0 {
		badResponse := model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid user ID",
			Data:       nil,
		}
		w.WriteHeader(http.StatusBadRequest) // Set status code
		json.NewEncoder(w).Encode(badResponse)
		return
	}

	// Panggil CustomerByID dari userService
	user, err := h.userService.UserByID(idInt)
	if err != nil {
		badResponse := model.Response{
			StatusCode: http.StatusNotFound,
			Message:    "User Not Found",
			Data:       nil,
		}
		w.WriteHeader(http.StatusNotFound) // Set status code
		json.NewEncoder(w).Encode(badResponse)
		return
	}

	// Siapkan response untuk pengguna yang ditemukan
	response := model.Response{
		StatusCode: http.StatusOK,
		Message:    "success",
		Data:       user,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		log.Printf("error retrieving users: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Failed to encode users", http.StatusInternalServerError)
		log.Printf("error encoding users: %v", err)
	}
}

func (h *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {

	// bodyBytes, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	json.NewEncoder(w).Encode(model.Response{
	// 		StatusCode: http.StatusInternalServerError,
	// 		Messagge:   "Failed to read request body",
	// 		Data:       nil,
	// 	})
	// 	return
	// }
	// fmt.Println(string(bodyBytes))
	// r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	var req model.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request payload",
			Data:       nil,
		})
		return
	}
	defer r.Body.Close()

	user, err := h.userService.Login(req.Username, req.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
			Data:       nil,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.Response{
		StatusCode: http.StatusOK,
		Message:    "Login successful",
		Data:       user,
	})
}
