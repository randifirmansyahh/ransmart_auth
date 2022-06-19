package httpRequest

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog/log"
)

type HttpHeaders map[string]string

func HTTPResponse(method string, url string, jsondatastring string, arrheader map[string]string) (int, string, error) {

	var finalresult string

	var ioreader io.Reader
	if jsondatastring != "" {
		ioreader = bytes.NewBuffer([]byte(jsondatastring))
	}

	req, err := http.NewRequest(method, url, ioreader)
	if err != nil {
		log.Error().Msg("Error, NewRequest: " + err.Error())
	}

	if len(arrheader) > 0 {
		for k, v := range arrheader {
			req.Header.Set(k, v)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error().Msg("Error, client.Do: " + err.Error())
		finalresult = ""
		return 0, finalresult, err
	}
	defer resp.Body.Close()

	respbody, _ := ioutil.ReadAll(resp.Body)
	finalresult = string(respbody)
	return resp.StatusCode, finalresult, nil
}
