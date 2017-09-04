package main

import (
	"flag"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// MovieQuote represents a quote from a movie
type MovieQuote struct {
	Movie     string `json:"movie"`
	Quote     string `json:"quote"`
	Character string `json:"character"`
	Actor     string `json:"actor"`
	Year      uint   `json:"year"`
}

func main() {
	var movieQuotes []MovieQuote

	// Initialize command line parameters
	verbose := flag.Bool("v", false, "Be verbose")
	help := flag.Bool("h", false, "Display usage")
	flag.Parse()

	// Print usage text if binary is called with -h
	if *help {
		flag.Usage()
		os.Exit(0)
	}

	initLogging(*verbose)

	Logger.DEBUG.Println("Initializing random number generator...")
	rand.Seed(time.Now().Unix())

	movieQuotes = append(movieQuotes, MovieQuote{
		Movie:     "Star Wars: Episode V - The Empire Strikes Back",
		Quote:     "Luke, I am your father!",
		Character: "Darth Vader",
		Actor:     "David Prowse",
		Year:      1980,
	})

	// TODO: Add quotes from moviequotes.json
	// HINT: Use ioutil.ReadFile to read from a file
	// HINT: Use json.Unmarshal to convert to a proper type

	http.HandleFunc("/v1/moviequotes", func(resp http.ResponseWriter, req *http.Request) {

		// Switch according to HTTP Request Method
		switch req.Method {
		case "GET":
			// Set MIME type application/json to HTTP Response
			resp.Header().Add("Content-Type", "application/json")

			// TODO: Convert movieQuotes slice to json string
			// Use json.Marshal to convert structs to a []byte
			data := []byte("{}")
			resp.Write(data)
		case "POST":
			/*
				Test the POST method with the following curl command
				curl -v \
				  -H "Content-Type: application/json" \
				  -d '{"movie":"Terminator 2: Judgment Day","quote":"Hasta La Vista, Baby","actor":"Arnold Schwarzenegger","character":"The Terminator T-101", "year": 1991 }' \
				  http://localhost:1323/v1/moviequotes
			*/

			// TODO: Read HTTP POST Body
			// HINT: ioutil.ReadAll can be used to read streams

			// TODO: Append to movieQuotes the quote from the body

			// Send back HTTP Status 204
			resp.WriteHeader(http.StatusNoContent)
		default:
			// Return HTTP Status 405 method allowed code if the user tries other http verbs like PATCH,PUT,DELETE
			Logger.WARN.Printf("Denied %v request to %v", req.Method, req.RequestURI)
			resp.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/v1/moviequotes/random", func(resp http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET":
			// TODO: return a random movie quote from moviequotes variable
			resp.Header().Add("Content-Type", "application/json")
			data := []byte("{\"test\": true}")
			resp.Write(data)
		default:
			// Return HTTP Status 405 method allowed code if the user tries other http verbs like PATCH,PUT,DELETE,POST
			resp.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	Logger.INFO.Println("net/http listening at http://localhost:1323")
	Logger.ERROR.Fatalln(http.ListenAndServe(":1323", nil))
}

// Logger is a basic logging facility
var Logger struct {
	INFO  *log.Logger
	ERROR *log.Logger
	WARN  *log.Logger
	DEBUG *log.Logger
}

func initLogging(verbose bool) {
	Logger.INFO = log.New(os.Stdout, "[INFO] - ", log.LstdFlags)
	Logger.ERROR = log.New(os.Stderr, "[ERROR] - ", log.LstdFlags)
	Logger.WARN = log.New(os.Stderr, "[WARN] - ", log.LstdFlags)
	if verbose {
		Logger.DEBUG = log.New(os.Stdout, "[DEBUG] - ", log.LstdFlags)
	} else {
		Logger.DEBUG = log.New(ioutil.Discard, "[DEBUG] - ", log.LstdFlags)
	}
}
