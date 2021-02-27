package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Vincent-Carrier/libchess"
	"github.com/go-chi/chi"
	mw "github.com/go-chi/chi/middleware"
	_ "github.com/joho/godotenv/autoload"
)

var (
	PORT, _ = strconv.Atoi(os.Getenv("PORT"))
)

type Server struct {
}

var srv = Server{}

func main() {
	r := chi.NewRouter()
	r.Use(mw.Logger, mw.Recoverer, mw.Timeout(15*time.Second))

	fmt.Printf("Application started on http://localhost:%d\n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", PORT), r))
}
