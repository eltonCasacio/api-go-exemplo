package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/eltoncasacio/api-go/internal/dto"
	"github.com/eltoncasacio/api-go/internal/entity"
	database "github.com/eltoncasacio/api-go/internal/infra/database/user"
	"github.com/go-chi/jwtauth"
)

type Error struct {
	ErrorMessage string `json:"error-message"`
}

type UserHandler struct {
	userDB database.UserRepositoryInterface
}

func NewUserHandler(db database.UserRepositoryInterface) *UserHandler {
	return &UserHandler{userDB: db}
}

// GetJWT godoc
// @Summary      Buscar usuário JWT
// @Description  Buscar usuário JWT
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request   body     dto.GetJWTInput  true  "user credentials"
// @Success      200  {object}  dto.GetJWTOutput
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /users/generate_token [post]
func (uh *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExperiesIn := r.Context().Value("jwtExperiesIn").(int)

	var user dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{ErrorMessage: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	u, err := uh.userDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := Error{ErrorMessage: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExperiesIn)).Unix(),
	})
	accessToken := dto.GetJWTOutput{AccessToken: tokenString}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// Create user godoc
// @Summary      Criar usuário
// @Description  Criar usuário
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request     body      dto.CreateUserInput  true  "user request"
// @Success      201
// @Failure      500         {object}  Error
// @Router       /users [post]
func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := dto.CreateUserInput{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{ErrorMessage: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	err = uh.userDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{ErrorMessage: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
