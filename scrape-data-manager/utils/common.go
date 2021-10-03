package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// GetStructFromRequest takes json as input and converts it to go struct.
func GetStructFromRequest(w http.ResponseWriter, r *http.Request, arg interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	if len(body) == 0 || err != nil {
		log.Println("Empty body")
		RespondWithError(w, 400, "empty request body!")
	} else {
		err = json.Unmarshal(body, arg)
		if err != nil {
			log.Println("Unmarshalling error: ", err)
			RespondWithError(w, 400, "cannot unmarshal request!")
		}
	}
}

// RespondWithError sends Bad Request Response.
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func CheckError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// RespondWithJSON takes struct and response writer as input and responds back as JSON.
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	CheckError(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(response)
	CheckError(err)
}

// GetEnv gets value from env or sets defaultValue if env variable is not found.
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
