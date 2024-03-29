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

	s "git.urantiatech.com/auth/auth/service"
	t "git.urantiatech.com/auth/auth/token"
	"git.urantiatech.com/auth/auth/user"
	h "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
)

func main() {

	var port int
	var key string
	flag.IntVar(&port, "port", 8080, "Port number")
	flag.StringVar(&key, "key", "NEW", "Signing key")
	flag.StringVar(&user.DBPath, "dbpath", "db", "User database directory")
	flag.Parse()

	log.SetFlags(log.Lshortfile)

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
		t.SigningKey = []byte(t.RandomToken(16))
	} else if key != "" {
		t.SigningKey = []byte(key)
	}
	log.Println("SIGNING_KEY: ", string(t.SigningKey))

	// Create a cache with TTL for storing logged-out tokens
	t.AccessTokenValidity = 1 * time.Hour
	t.RefreshTokenValidity = 3 * time.Hour

	t.RememberMeAccessTokenValidity = 24 * time.Hour
	t.RememberMeRefreshTokenValidity = 30 * 24 * time.Hour

	t.ResetTokenValidity = 24 * time.Hour
	t.ConfirmTokenValidity = 7 * 24 * time.Hour
	t.UpdateTokenValidity = 15 * time.Minute

	// Create cache for holding invalid tokens
	t.BlacklistAccessTokens = cache.New(t.RememberMeAccessTokenValidity, 2*t.RememberMeAccessTokenValidity)
	t.BlacklistRefreshTokens = cache.New(t.RememberMeRefreshTokenValidity, 2*t.RememberMeRefreshTokenValidity)

	// Create cache for temporary registration
	user.TemporaryRegistrationValidity = 1 * time.Hour
	user.TemporaryRegistration = cache.New(user.TemporaryRegistrationValidity, 2*user.TemporaryRegistrationValidity)

	var svc s.Auth
	svc = s.Auth{}

	r := mux.NewRouter()

	// r.Handle("/", h.NewServer(s.MakeLoginEndpoint(svc), s.DecodeLoginRequest, s.EncodeResponse))
	r.Handle("/authorize", h.NewServer(s.MakeAuthorizeEndpoint(svc), s.DecodeAuthorizeRequest, s.EncodeResponse))
	r.Handle("/login", h.NewServer(s.MakeLoginEndpoint(svc), s.DecodeLoginRequest, s.EncodeResponse))
	r.Handle("/logout", h.NewServer(s.MakeLogoutEndpoint(svc), s.DecodeLogoutRequest, s.EncodeResponse))
	r.Handle("/register", h.NewServer(s.MakeRegisterEndpoint(svc), s.DecodeRegisterRequest, s.EncodeResponse))
	r.Handle("/update", h.NewServer(s.MakeUpdateEndpoint(svc), s.DecodeUpdateRequest, s.EncodeResponse))
	r.Handle("/delete", h.NewServer(s.MakeDeleteEndpoint(svc), s.DecodeDeleteRequest, s.EncodeResponse))
	r.Handle("/identify", h.NewServer(s.MakeIdentifyEndpoint(svc), s.DecodeIdentifyRequest, s.EncodeResponse))
	r.Handle("/profile", h.NewServer(s.MakeProfileEndpoint(svc), s.DecodeProfileRequest, s.EncodeResponse))
	r.Handle("/refresh", h.NewServer(s.MakeRefreshEndpoint(svc), s.DecodeRefreshRequest, s.EncodeResponse))
	r.Handle("/confirm", h.NewServer(s.MakeConfirmEndpoint(svc), s.DecodeConfirmRequest, s.EncodeResponse))
	r.Handle("/forgot", h.NewServer(s.MakeForgotEndpoint(svc), s.DecodeForgotRequest, s.EncodeResponse))
	r.Handle("/reset", h.NewServer(s.MakeResetEndpoint(svc), s.DecodeResetRequest, s.EncodeResponse))

	http.ListenAndServe(fmt.Sprintf(":%d", port), r)

}
