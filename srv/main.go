package main

import (
	"fmt"
	"github.com/0sax/apiHelpers"
	"github.com/0sax/blacklist"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

func main() {

	// Load env vars
	err := godotenv.Load("srv/vars.env")
	if err != nil {
		log.Println("couldn't load env vars because:", err)
	}

	router := httprouter.New()

	router.GET("/crc/:bvn", crcHandler())
	router.GET("/blacklist/:bvn", blacklistHandler())
	router.GlobalOPTIONS = corsHandler()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}

func crcHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		bvn := p.ByName("bvn")

		if bvn == "" {
			apiHelpers.WriteErrorJSONResponse(w, 400, "invalid bvn")
			return
		}

		bl := blacklist.NewBlackListClient(
			os.Getenv("BLACKLIST_BASE_URL"),
			os.Getenv("BLACKLIST_API_KEY"))

		crc, err := bl.SearchCRCFull(bvn)
		if err != nil {
			apiHelpers.WriteError(w, err, 500, "internal error")
			return
		}

		apiHelpers.WriteOKJSONResponse(w, 200, "crc search done", crc)
		return

	}
}

func blacklistHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		bvn := p.ByName("bvn")

		if bvn == "" {
			apiHelpers.WriteErrorJSONResponse(w, 400, "invalid bvn")
			return
		}

		bl := blacklist.NewBlackListClient(
			os.Getenv("BLACKLIST_BASE_URL"),
			os.Getenv("BLACKLIST_API_KEY"))

		blr, err := bl.SearchBlacklistFull(bvn)
		if err != nil {
			apiHelpers.WriteError(w, err, 500, "internal error")
			return
		}

		apiHelpers.WriteOKJSONResponse(w, 200, fmt.Sprintf("%v blacklist record(s) found", len(blr)), blr)
		return

	}
}

//Cors Handler
func corsHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Add("Access-Control-Allow-Headers", "Content-Type")
			header.Set("Content-Type", "application/json")
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			//header.Set("Access-Control-Allow-Origin", os.Getenv("BACKEND_CLIENT_URL"))
			header.Set("Access-Control-Allow-Origin", "*")
		}

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	}
}
