package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/dataleodev/registry"
	"github.com/dataleodev/registry/api"
	"github.com/dataleodev/registry/bcrypt"
	"github.com/dataleodev/registry/jwt"
	"github.com/dataleodev/registry/logger"
	"github.com/dataleodev/registry/pkg/random"
	"github.com/dataleodev/registry/pkg/uuid"
	"github.com/dataleodev/registry/postgres"
	"github.com/dataleodev/registry/postgres/keys"
	"github.com/dataleodev/registry/postgres/nodes"
	"github.com/dataleodev/registry/postgres/regions"
	"github.com/dataleodev/registry/postgres/users"
	"github.com/go-kit/kit/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
	)
	flag.Parse()

	var l log.Logger
	{
		l = log.NewLogfmtLogger(os.Stderr)
		l = log.With(l, "ts", log.DefaultTimestampUTC)
		l = log.With(l, "caller", log.DefaultCaller)
	}

	log, err := logger.New(os.Stderr, "info")
	if err != nil {
		panic(err)
	}

	dbConfig := postgres.DBConfig{
		Hostname: "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		DBName:   "postgres",
		SslMode:  "disable",
	}

	db := connectToDB(dbConfig,log)
	defer db.Close()

	tokenizer := jwt.NewTokenizer()
	hasher := bcrypt.New()
	up := uuid.New()
	randomizer := random.New([]rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ123456789!@#$%&abcdefghijklmnopqrstuvwxyz"))
	userStore := users.NewRepository(db)
	nodeStore := nodes.NewRepository(db)
	regionStore := regions.NewRepository(db)
	keyStore := keys.NewRepository(db)

	var s registry.Service
	{
		s = registry.NewService(userStore,nodeStore,regionStore,keyStore,up,hasher,l,tokenizer,randomizer)
		s = api.LoggingMiddleware(l)(s)
	}

	var h http.Handler
	{
		h = api.MakeHTTPHandler(s, l)
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {

		log.Info(fmt.Sprintf("registry started on port %v", *httpAddr))
		//l.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, h)
	}()

	log.Info(fmt.Sprintf("exit due to: %v", <-errs))
}

func connectToDB(cfg postgres.DBConfig, logger logger.Logger) *sql.DB {
	db, err := postgres.New(cfg)
	if err != nil {
		logger.Info(fmt.Sprintf("Failed to connect to postgres: %s", err))
		os.Exit(1)
	}
	return db
}

