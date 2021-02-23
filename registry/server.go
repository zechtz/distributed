package registry

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

// ServerPort the port the server will run on
const ServerPort = ":3000"

// ServiceURL the url the server will run on
const ServiceURL = "http://localhost" + ServerPort + "/services"

type registry struct {
	registrations []Registration
	mutex         *sync.Mutex
}

func (r *registry) add(reg Registration) error {
	r.mutex.Lock()
	r.registrations = append(r.registrations, reg)
	r.mutex.Unlock()
	return nil
}

var reg = registry{
	registrations: make([]Registration, 0),
	mutex:         new(sync.Mutex),
}

// RegistryService struct
type RegistyService struct{}

func (s RegistyService) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Println("Request received")
	switch req.Method {
	case http.MethodPost:
		dec := json.NewDecoder(req.Body)
		var r Registration
		err := dec.Decode(&r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("Adding service: %v with URL: %v\n", r.ServiceName, r.ServiceURL)
		err = reg.add(r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
