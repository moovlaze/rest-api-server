package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIserver struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Функция с помощью которой мы будем запускать сервер, подключаться к базе данных и т.д.
func (a *APIserver) Start() error {
	if err := a.configureLogger(); err != nil {
		return nil
	}

	a.configureRouter()

	a.logger.Info("Start API server")

	return http.ListenAndServe(a.config.Bindaddr, a.router)
}

func (a *APIserver) configureLogger() error {
	level, err := logrus.ParseLevel(a.config.LogLevel)
	if err != nil {
		return nil
	}

	logrus.SetLevel(level)

	return nil
}

func (a *APIserver) configureRouter() {

	a.router.HandleFunc("/hello", a.handleHello())
}

func (a *APIserver) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world")
	}
}