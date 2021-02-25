package log

import (
	"bytes"
	"fmt"
	stlog "log"
	"net/http"

	"github.com/zechtz/distributed/registry"
)

type clientLogger struct {
	url string
}

// SetClientLogger sets client's logger file
func SetClientLogger(serviceURL string, clientService registry.ServiceName) {
	stlog.SetPrefix(fmt.Sprintf("[%v] - ", clientService))
	stlog.SetFlags(0)
	stlog.SetOutput(&clientLogger{url: serviceURL})
}

func (cl clientLogger) Write(data []byte) (int, error) {
	b := bytes.NewBuffer([]byte(data))
	res, err := http.Post(cl.url+"/log", "text/plain", b)
	if err != nil {
		return 0, err
	}
	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("Failed to send log message. Servive reponded with an error")
	}

	return len(data), nil
}
