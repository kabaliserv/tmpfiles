package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"

	"github.com/kabaliserv/tmpfiles/config"
	"github.com/kabaliserv/tmpfiles/controllers"
	"github.com/kabaliserv/tmpfiles/models"
	"github.com/kabaliserv/tmpfiles/routes"
	"github.com/kabaliserv/tmpfiles/storage"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/mux"
	"github.com/urfave/cli/v2"
)

var (
	// FlagConf : Description conf Flag
	FlagConf = &cli.StringFlag{
		Name:     "conf",
		Aliases:  []string{"c"},
		Usage:    "TOML config file",
		Required: true,
	}
	// StartApp : App start function
	StartApp = startApp
)

// App start function
func startApp(c *cli.Context) error {

	// Get Conf Arg
	pathConf, _ := filepath.Abs(c.String("conf"))

	// Get Config file to array bytes
	src, err := ioutil.ReadFile(pathConf)
	if err != nil {
		return err
	}

	// Parse Config file
	conf := config.Conf{}
	err = toml.Unmarshal(src, &conf)
	if err != nil {
		return err
	}
	if err := conf.Store.Valid(); err != nil {
		return err
	}

	// Parse and get listen address from config
	var listenAddress string
	if conf.Web.Address == "" {
		listenAddress += "localhost"
	} else {
		listenAddress += conf.Web.Address
	}
	if conf.Web.Port == 0 {
		listenAddress += fmt.Sprintf(":%v", "3000")
	} else {
		listenAddress += fmt.Sprintf(":%v", conf.Web.Port)
	}

	// init New Store Files
	store, err := storage.NewStore(conf.Store.Path)
	if err != nil {
		return err
	}

	// Get DataBase
	db, err := models.NewSqliteDB(store.GetDataPath())
	if err != nil {
		return err
	}

	// Init Controller
	controllers := controllers.NewController(store, db)

	// Get root path from config
	var urlpath = "/"
	if conf.Web.Path != "" {
		urlpath = conf.Web.Path
	}

	// Get new router
	var handler = mux.NewRouter()

	// Add root route path on router
	r := handler.PathPrefix(urlpath).Name("rootpath").Subrouter()

	// Add all routes controlleur on router
	routes.AddRoutes(r, controllers)

	// Init web server
	srv := &http.Server{
		Handler:      handler,
		Addr:         listenAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Show listen port
	fmt.Printf("Listen Port: \":%v\"\n", conf.Web.Port)

	return srv.ListenAndServe()
}
