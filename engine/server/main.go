package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	mw "github.com/go-chi/chi/middleware"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/cors"

	chess "github.com/Vincent-Carrier/libchess"
)

var (
	PORT, _ = strconv.Atoi(os.Getenv("PORT"))
)

//type Server struct {
//}

//var srv = Server{}

func main() {
	r := chi.NewRouter()
	r.Use(
		mw.Logger,
		mw.Recoverer,
		mw.Timeout(15*time.Second),
		cors.Default().Handler,
	)

	r.Get("/game/{fen}", func(w http.ResponseWriter, r *http.Request) {
		fen := chi.URLParam(r, "fen")
		game, err := chess.NewGame(fen)
		RenderJson(w, map[string]interface{}{
			"square": chess.Square("e4"),
		})
	})

	fmt.Printf("Application started on http://localhost:%d\n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", PORT), r))
}

func RenderJson(w http.ResponseWriter, payload interface{}) {
	body, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(422)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(body)
}
