package main

import (
	"net/http"

	configs "github.com/eltoncasacio/api-go/configs"
	_ "github.com/eltoncasacio/api-go/docs"
	"github.com/eltoncasacio/api-go/internal/entity"
	dbProduct "github.com/eltoncasacio/api-go/internal/infra/database/product"
	dbUser "github.com/eltoncasacio/api-go/internal/infra/database/user"
	"github.com/eltoncasacio/api-go/internal/infra/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title           Go API Example
// @version         1.0
// @description     API de produto com autenticação (JWT)
// @termsOfService  http://swagger.io/terms/
// @contact.name   Elton Casacio
// @contact.url    https://www.instagram.com/elton_casacio/
// @contact.email  eltoncasacio@hotmail.com.br
// @license.name   C3R Innovation
// @license.url    https://c3rinnovation.com
// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	configs, err := configs.LoadConfig("./cmd/.env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("./cmd/teste.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("jwtExperiesIn", configs.JwtExperesIn))

	productDB := dbProduct.NewProductRepository(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := dbUser.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userDB)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Get("/", productHandler.GetProducts)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)
		r.Post("/generate_token", userHandler.GetJWT)
	})

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", r)
}
