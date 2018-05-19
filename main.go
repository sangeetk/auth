package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	httptransport "github.com/urantiatech/kit/transport/http"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	"git.urantiatech.com/auth/auth/middleware"
	"git.urantiatech.com/auth/auth/model"
	"git.urantiatech.com/auth/auth/service"
	"github.com/patrickmn/go-cache"
	"github.com/urfave/negroni"
)

func main() {
	var listen = flag.String("listen", ":9018", "HTTP listen address")
	var key = flag.String("key", "", "The Signing Key")

	// sqlite, mysql or postgres
	var driver = flag.String("driver", "mysql", "The Database driver")
	// sqlite
	var sqlitedb = flag.String("sqlitedb", "auth.db", "The SQLite database file")
	// mysql and postgres
	var dbuser = flag.String("dbuser", "auth", "The database username")
	var dbpass = flag.String("dbpass", "auth", "The database password")
	var dbname = flag.String("dbname", "auth", "The database name")

	var host = flag.String("host", "localhost", "The database host")
	var port = flag.Int("port", 0, "The database port")

	var create bool
	flag.BoolVar(&create, "create", false, "Create DB tables")

	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	if *key != "" {
		service.SigningKey = []byte(*key)
	} else {
		service.SigningKey = []byte(service.RandomToken(16))
	}

	log.Println("SigningKey : ", string(service.SigningKey))

	var conn string
	var err error

	switch *driver {
	case "sqlite":
		conn = *sqlitedb
		service.DB, err = gorm.Open("sqlite3", conn)

	case "mysql":
		if *port == 0 {
			*port = 3306
		}
		conn = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", *dbuser, *dbpass, *dbname)
		service.DB, err = gorm.Open("mysql", conn)

	case "postgres":
		if *port == 0 {
			*port = 5432
		}
		conn = fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s", *host, *port, *dbuser, *dbname, *dbpass)
		service.DB, err = gorm.Open("postgres", conn)
	}

	// log.Println("conn: ", conn)

	// Try to connect but continue even if database is down
	if err != nil {
		log.Panic(err.Error())
	}

	if create {
		service.DB.CreateTable(&model.User{})
		service.DB.CreateTable(&model.Role{})
		service.DB.CreateTable(&model.Address{})
		service.DB.CreateTable(&model.Profile{})
	} else {
		if !service.DB.HasTable(&model.User{}) ||
			!service.DB.HasTable(&model.Role{}) ||
			!service.DB.HasTable(&model.Address{}) ||
			!service.DB.HasTable(&model.Profile{}) {
			log.Println("Database does not has required tables")
			return
		}
	}

	// Create a cache with TTL for storing logged-out tokens
	service.TokenValidity = 10 * time.Minute
	service.BlacklistTokens = cache.New(service.TokenValidity, 2*service.TokenValidity)

	var svc service.Auth
	svc = service.Auth{}

	registerHandler := httptransport.NewServer(
		makeRegisterEndpoint(svc),
		decodeRegisterRequest,
		encodeResponse,
	)

	updateHandler := httptransport.NewServer(
		makeUpdateEndpoint(svc),
		decodeUpdateRequest,
		encodeResponse,
	)

	loginHandler := httptransport.NewServer(
		makeLoginEndpoint(svc),
		decodeLoginRequest,
		encodeResponse,
	)

	logoutHandler := httptransport.NewServer(
		makeLogoutEndpoint(svc),
		decodeLogoutRequest,
		encodeResponse,
	)

	identifyHandler := httptransport.NewServer(
		makeIdentifyEndpoint(svc),
		decodeIdentifyRequest,
		encodeResponse,
	)

	profileHandler := httptransport.NewServer(
		makeProfileEndpoint(svc),
		decodeProfileRequest,
		encodeResponse,
	)

	refreshHandler := httptransport.NewServer(
		makeRefreshEndpoint(svc),
		decodeRefreshRequest,
		encodeResponse,
	)

	confirmHandler := httptransport.NewServer(
		makeConfirmEndpoint(svc),
		decodeConfirmRequest,
		encodeResponse,
	)

	recoverHandler := httptransport.NewServer(
		makeRecoverEndpoint(svc),
		decodeRecoverRequest,
		encodeResponse,
	)

	n := negroni.Classic()
	r := mux.NewRouter()
	r.Handle("/", loginHandler)
	r.Handle("/login", loginHandler)
	r.Handle("/logout", logoutHandler)
	r.Handle("/register", registerHandler)
	r.Handle("/update", updateHandler)
	r.Handle("/identify", identifyHandler)
	r.Handle("/profile", profileHandler)
	r.Handle("/refresh", refreshHandler)
	r.Handle("/confirm", confirmHandler)
	r.Handle("/recover", recoverHandler)

	n.Use(middleware.Database())
	n.UseHandler(r)
	http.ListenAndServe(*listen, n)

	// CLose the database connection
	if service.DB != nil {
		service.DB.Close()
	}

}
