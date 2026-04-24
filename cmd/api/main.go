package main

import (
	"my-project/internal/config"
	"net/http"
	"os"

	"github.com/TAhirr01/cliflags"
	"github.com/TAhirr01/confmaker"
	"github.com/joho/godotenv"
)

func main() {
	path := cliflags.RegisterFlag("ENV")
	cliflags.ParseFlags()
	godotenv.Load(path.Value)

	var cfg config.Config
	if err := confmaker.Load(os.Getenv("FILE_LOCATION"), &cfg); err != nil {
		panic(err)
	}

	//Dependecy Injection
	server := DepedencyInjection()

	//Custom server Multiplexer
	mux := http.NewServeMux()

	//
	CreteRautes(mux, server)

	// log.Fatal(http.ListenAndServe(":8080", mux))
}
