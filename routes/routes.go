package routes

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/kabaliserv/tmpfiles/controllers"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"
)

// Init : add new routes on router
func Init() *mux.Router {
	var router = mux.NewRouter()

	managers := controllers.GetManagers()

	// Download Route: Use to download files
	router.HandleFunc("/d/{id}", managers.DownloadManager).Methods("GET")

	// Make Path for api reqiest
	r := router.PathPrefix("/api/").Name("apipath").Subrouter()

	// Upload Route: Use for post new Upload
	r.HandleFunc("/upload", managers.UploadManager).Methods("POST")
	r.PathPrefix("/upload/cache/").Handler(managers.InitTusServer())

	// Meta Route: Use to get metadata files
	r.HandleFunc("/meta/{id}", managers.MetadataManager).Methods("GET")

	// Auth Route: Use to get Token
	r.HandleFunc("/auth", managers.AuthManager).Methods("POST")

	// static assets & 404 handler
	box := packr.NewBox("../client/dist")
	router.Path("/").Handler(http.FileServer(box))
	router.Path("/index.html").Handler(serverFile(&box, "/index.html"))
	router.Path("/favicon.ico").Handler(serverFile(&box, "/favicon.ico"))
	router.PathPrefix("/css").Handler(http.FileServer(box))
	router.PathPrefix("/font").Handler(http.FileServer(box))
	router.PathPrefix("/js").Handler(http.FileServer(box))
	router.NotFoundHandler = notFoundPath(&box)

	// Enumeration of all the routes
	if err := router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("ROUTE:", pathTemplate)
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			fmt.Println("Path regexp:", pathRegexp)
		}
		queriesTemplates, err := route.GetQueriesTemplates()
		if err == nil {
			fmt.Println("Queries templates:", strings.Join(queriesTemplates, ","))
		}
		queriesRegexps, err := route.GetQueriesRegexp()
		if err == nil {
			fmt.Println("Queries regexps:", strings.Join(queriesRegexps, ","))
		}
		methods, err := route.GetMethods()
		if err == nil {
			fmt.Println("Methods:", strings.Join(methods, ","))
		}

		fmt.Println()

		return nil

	}); err != nil {

		fmt.Println(err)

	}

	return router
}

func notFoundPath(box *packr.Box) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, err := box.Open("/index.html")
		if err != nil {
			w.WriteHeader(500)
			return
		}
		defer file.Close()
		w.WriteHeader(200)
		io.Copy(w, file)

	})
}

func serverFile(box *packr.Box, filename string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, err := box.Open(filename)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		defer file.Close()
		w.WriteHeader(200)
		io.Copy(w, file)
	})
}
