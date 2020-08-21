package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/zioufang/mltrackapi/pkg/api/model"
)

// Server is the struct for the server
type Server struct {
	db     *gorm.DB
	router *chi.Mux
}

// Init initialize the server with database and router
// TODO adds postgres, mysql support
func (s *Server) Init(DbDriver, DbName string) {
	var err error
	switch DbDriver {
	case "sqlite3":
		s.db, err = gorm.Open(DbDriver, DbName)
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal(fmt.Errorf("%s is not a supported database", DbDriver))
	}
	// TODO add foreign key & index when necessary
	s.db.AutoMigrate(&model.Model{}, &model.ModelRun{})
	s.router = chi.NewRouter()
	// A good base middleware stack
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	// set routes
	s.SetRoutes()
}

// SetRoutes sets the routs for the server
func (s *Server) SetRoutes() {
	r := s.router
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Mltrack\n")
	})

	// model endpoints
	r.Route("/models", func(r chi.Router) {
		r.Post("/", s.CreateModel)
		r.Get("/", s.GetModel)
		r.Delete("/", s.DeleteModel)
		r.Get("/all", s.GetAllModels)
	})

	// model run endpoints
	r.Route("/runs", func(r chi.Router) {
		r.Post("/", s.CreateModelRun)
		r.Get("/", s.GetAllModelRuns)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", s.GetModelRunByID)
			r.Delete("/", s.DeleteModelRunByID)
		})
	})

}

// Run runs the server
func (s *Server) Run(port uint) {
	fmt.Printf("Listening to port %d\n", port)
	addr := ":" + fmt.Sprint(port)
	log.Fatal(http.ListenAndServe(addr, s.router))
}
