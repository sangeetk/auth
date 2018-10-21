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
	"git.urantiatech.com/auth/auth/user"
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
	h "github.com/urantiatech/kit/transport/http"
)

func main() {

	var port int
	var key string
	flag.IntVar(&port, "port", 8080, "Port number")
	flag.StringVar(&key, "key", "NEW", "Signing key")
	flag.StringVar(&user.DBPath, "dbpath", "db", "User database directory")
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
		s.SigningKey = []byte(s.RandomToken(16))
	} else if key != "" {
		s.SigningKey = []byte(key)
	}
	log.Println("SIGNING_KEY: ", string(s.SigningKey))

	// Create a cache with TTL for storing logged-out tokens
	s.AccessTokenValidity = 1 * time.Hour
	s.RefreshTokenValidity = 24 * time.Hour
	s.BlacklistTokens = cache.New(s.AccessTokenValidity, 2*s.AccessTokenValidity)

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
	r.Handle("/recover", h.NewServer(s.MakeRecoverEndpoint(svc), s.DecodeRecoverRequest, s.EncodeResponse))
	r.Handle("/reset", h.NewServer(s.MakeResetEndpoint(svc), s.DecodeResetRequest, s.EncodeResponse))

	http.ListenAndServe(fmt.Sprintf(":%d", port), r)

}
