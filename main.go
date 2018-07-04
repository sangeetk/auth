package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"git.urantiatech.com/auth/auth/service"
	"git.urantiatech.com/auth/auth/user"
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
	h "github.com/urantiatech/kit/transport/http"
)

func main() {
	var port int
	var key string
	flag.IntVar(&port, "port", 9999, "Port for running the UI")
	flag.StringVar(&key, "key", "NEW", "Signing key")
	flag.StringVar(&user.DBPath, "dbpath", "db", "Directory for storing user database")
	flag.Parse()

	if os.Getenv("PORT") != "" {
		p, err := strconv.ParseInt(os.Getenv("PORT"), 10, 32)
		if err != nil {
			port = int(p)
		}
	}

	rand.Seed(time.Now().UnixNano())

	if os.Getenv("SIGNING_KEY") != "" {
		key = os.Getenv("SIGNING_KEY")
	}

	if key == "NEW" {
		service.SigningKey = []byte(service.RandomToken(16))
	} else if key != "" {
		service.SigningKey = []byte(key)
	}
	log.Println("SIGNING_KEY: ", string(service.SigningKey))

	// Create a cache with TTL for storing logged-out tokens
	service.AccessTokenValidity = 1 * time.Hour
	service.RefreshTokenValidity = 24 * time.Hour
	service.BlacklistTokens = cache.New(service.AccessTokenValidity, 2*service.AccessTokenValidity)

	var svc service.Auth
	svc = service.Auth{}

	r := mux.NewRouter()

	// r.Handle("/", h.NewServer(makeLoginEndpoint(svc), decodeLoginRequest, encodeResponse))
	r.Handle("/login", h.NewServer(makeLoginEndpoint(svc), decodeLoginRequest, encodeResponse))
	r.Handle("/logout", h.NewServer(makeLogoutEndpoint(svc), decodeLogoutRequest, encodeResponse))
	r.Handle("/register", h.NewServer(makeRegisterEndpoint(svc), decodeRegisterRequest, encodeResponse))
	r.Handle("/update", h.NewServer(makeUpdateEndpoint(svc), decodeUpdateRequest, encodeResponse))
	r.Handle("/identify", h.NewServer(makeIdentifyEndpoint(svc), decodeIdentifyRequest, encodeResponse))
	r.Handle("/profile", h.NewServer(makeProfileEndpoint(svc), decodeProfileRequest, encodeResponse))
	r.Handle("/refresh", h.NewServer(makeRefreshEndpoint(svc), decodeRefreshRequest, encodeResponse))
	r.Handle("/confirm", h.NewServer(makeConfirmEndpoint(svc), decodeConfirmRequest, encodeResponse))
	r.Handle("/recover", h.NewServer(makeRecoverEndpoint(svc), decodeRecoverRequest, encodeResponse))

	http.ListenAndServe(fmt.Sprintf(":%d", port), r)

}
