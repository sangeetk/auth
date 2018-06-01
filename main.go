package main

import (
	"flag"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"git.urantiatech.com/auth/auth/service"
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
	httptransport "github.com/urantiatech/kit/transport/http"
)

func main() {
	var listen = ":9018"
	var port = os.Getenv("PORT")

	if port != "" {
		listen = ":" + port
		log.Println("Auth service listening on", listen)
	} else {
		log.Println("PORT environment variable not provided, listening on", listen)
	}

	var key = os.Getenv("SIGNING_KEY")

	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	if key == "" {
		log.Fatalln("SIGNING_KEY environment variable not provided, exiting.")
	} else if key == "NEW" {
		service.SigningKey = []byte(service.RandomToken(16))
		log.Println("Generating new SIGNING_KEY: ", string(service.SigningKey))
	} else {
		service.SigningKey = []byte(key)
		log.Println("SIGNING_KEY : ", string(service.SigningKey))
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

	http.ListenAndServe(listen, r)

}
