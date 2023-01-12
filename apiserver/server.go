package apiserver

import (
	"WEBAPP/apiserver/config"
	"WEBAPP/apiserver/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	Config *config.Config
	Router *mux.Router
	Logger *logrus.Logger
}

func New() *Server {
	return &Server{
		Config: config.New(),
		Router: mux.NewRouter(),
		Logger: logrus.New(),
	}
}

func (s *Server) Start() error {
	s.Logger = s.ConfigureLogger()
	s.Router = s.ConfigureRouter()
	logrus.Info("Server started at port ", s.Config.Server.Port)
	logrus.Info("Server bind addr: ", "http://"+s.Config.Server.Host+s.Config.Server.Port)
	return http.ListenAndServe(s.Config.Server.Port, s.Router)
}

func (s *Server) Stop() error {
	return nil
}

func (s *Server) ConfigureServer(path string) (*Server, error) {
	var err error
	s.Config, err = config.TakeFromFile(path)
	if err != nil {
		return nil, err
	}
	s.Logger.Trace("Server is configurated")
	return s, nil
}

func (s *Server) ConfigureLogger() *logrus.Logger {
	s.Logger = logrus.New()
	s.Logger.Level = logrus.Level(s.Config.Server.LogLevel)
	s.Logger.Trace("Logger is configurated")
	return s.Logger
}

func (s *Server) ConfigureRouter() *mux.Router {
	handlers.InitHandlersLogger(s.Config.Server.LogLevel)
	s.Router = mux.NewRouter()

	s.Router.HandleFunc("/", handlers.Index)
	s.Router.HandleFunc("/auth", handlers.Auth)
	s.Router.HandleFunc("/reset", handlers.Reset)
	s.Router.HandleFunc("/signup", handlers.Signup)
	s.Router.HandleFunc("/singleBlog", handlers.SingleBlog)
	s.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("./front/")))
	s.Logger.Trace("Router is configurated")
	return s.Router
}
