package cmd

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/kabaliserv/tmpfiles/config"
	"github.com/kabaliserv/tmpfiles/controllers"
	"github.com/kabaliserv/tmpfiles/models"
	"github.com/kabaliserv/tmpfiles/modules/clean"
	"github.com/kabaliserv/tmpfiles/modules/scheduler"
	"github.com/kabaliserv/tmpfiles/routes"
	"github.com/kabaliserv/tmpfiles/storage"

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

	if err := initializedAPP(c.String("conf")); err != nil {
		panic(err)
	}	

	// Init Schedulers
	scheduler.Init()

	// Init Controller
	controllers.Init()

	// Get new router
	var handler = routes.Init()

	// Init web server
	srv := &http.Server{
		Handler:      handler,
		Addr:         config.GetWebAddr() + ":" + config.GetWebPort(),
		WriteTimeout: 3 * time.Hour,
		ReadTimeout:  15 * time.Second,
	}

	// Show listen port
	fmt.Printf("Listen Port: \":%v\"\n", config.GetWebPort())

	return srv.ListenAndServe()
}

// CleanUp :
func CleanUp(c *cli.Context) error {
	if err := initializedAPP(c.String("conf")); err != nil {
		panic(err)
	}
	clean.CleanupUpload()
	return nil
}


func initializedAPP(confPath string) error {
	// Get Conf Arg
	pathConf, _ := filepath.Abs(confPath)

	// Parse File Config
	if err := config.ParseConfigFile(pathConf); err != nil {
		panic(err)
	}

	// Validate Config Options
	if err := config.GetConfigOptions().ValidConf(); err != nil {
		panic(err)
	}

	// init New Store Files
	if err := storage.Init(); err != nil {
		return err
	}

	// Init connection to dataBase
	if err := models.InitDB(); err != nil {
		panic(err)
	}

	return nil
}
