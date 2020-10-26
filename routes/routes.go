package routes

import (
	"fmt"
	"strings"

	"github.com/kabaliserv/tmpfiles/controllers"

	"github.com/gorilla/mux"
)

// NewRoutes :
func NewRoutes(controllers *controllers.Controller) *mux.Router {

	router := mux.NewRouter()

	//u := router.PathPrefix("/upload").Subrouter()

	// Upload Route: Use for post new Upload
	router.HandleFunc("/upload", controllers.PostUploads).Methods("POST")
	router.PathPrefix("/cache").Handler(controllers.NewCacheUpload("toto"))

	return router
}

// AddRoutes :
func AddRoutes(router *mux.Router, controllers *controllers.Controller) {

	//u := router.PathPrefix("/upload").Subrouter()

	// Get Root Path url
	url, err := router.Get("rootpath").URLPath()

	// Upload Route: Use for post new Upload
	router.HandleFunc("/upload", controllers.PostUploads).Methods("POST")
	router.PathPrefix("/upload/cache/").Handler(controllers.NewCacheUpload(url.Path))

	// Meta Route: Use to get metadata files
	router.HandleFunc("/meta/{id}", controllers.GetMeta).Methods("GET")

	// Auth Route: Use to get Token
	router.HandleFunc("/auth", controllers.GetAuth).Methods("POST")

	// Download Route: Use to download files
	router.HandleFunc("/d/{id}", controllers.GetFiles).Methods("GET")

	// Enumeration of all the routes
	err = router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
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
	})

	if err != nil {
		fmt.Println(err)
	}
}
