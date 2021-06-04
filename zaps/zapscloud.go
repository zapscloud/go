package zaps

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/zapscloud/golib/common"
)

// ZapsCloud - Document Service structure
type ZapsCloud struct {
	application string
	tenant      string
	stage       string
	credentials string
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)
	// log.SetOutput(ioutil.Discard)

}

func (p *ZapsCloud) SetStage(stage string) {
	p.stage = stage
}

func (p *ZapsCloud) Stage() string {
	return p.stage
}

func (p *ZapsCloud) SetTenant(tenant string) {
	p.tenant = tenant
}

func (p *ZapsCloud) Tenant() string {
	return p.tenant
}

// NewZapsCloud - Construct ZapsCloud
func NewZapsCloud(authKey string, authSecret, application string) (*ZapsCloud, error) {
	p := ZapsCloud{}

	p.application = application
	auth := authKey + ":" + authSecret
	p.credentials = "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))

	log.Println("NewZapsCloud ", p)
	return &p, nil
}

// EndZapsCloud - Close all the ZapsCloud
func (p *ZapsCloud) EndZapsCloud() {
	log.Printf("EndZapsCloud ")
}

func (p *ZapsCloud) makeReqeust(method string, requestURL string, body io.Reader) (*http.Request, error) {
	// Create a new request using http
	req, err := http.NewRequest(method, requestURL, body)

	// Set application/json content type
	req.Header.Set("Content-Type", "application/json")

	// add authorization header to the req
	req.Header.Add("Authorization", p.credentials)
	req.Header.Add("Application", p.application)
	if p.tenant != "" {
		req.Header.Add("Tenant", p.tenant)
	}
	log.Println("req_value", req)

	return req, err
}

// HttpGet - HTTP Get Request
func (p *ZapsCloud) HttpGet(requrl string) (map[string]interface{}, error) {

	log.Println("HttpGet::  Begin ", requrl)
	req, err := p.makeReqeust("GET", requrl, nil)
	if err != nil {
		return nil, err
	}

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)

	// check err
	log.Println(err)
	defer resp.Body.Close() // we have to close Body if we want to reuse the connection

	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	log.Println("Request status ", resp.StatusCode, resp.Status)

	if resp.StatusCode == 200 {
		responseData, errlog := ioutil.ReadAll(resp.Body)
		log.Println("Error Log### ", errlog)
		if errlog != nil {
			return nil, errlog
		}
		json.Unmarshal([]byte(responseData), &data)
		log.Println("unmarshal ", &data)

	} else {
		err = &common.AppError{ErrorCode: (string)(resp.StatusCode), ErrorMsg: resp.Status, ErrorDetail: resp.Status}
		data = nil
	}

	return data, err
}

// HttpGet - HTTP Get Request
func (p *ZapsCloud) HttpPost(requrl string, body map[string]interface{}) (map[string]interface{}, error) {

	log.Println("HttpPost::  Begin ", requrl)

	jsonstring, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodydata := bytes.NewBuffer(jsonstring)
	log.Println("HttpPost:: Body", requrl, bodydata)

	req, err := p.makeReqeust("POST", requrl, bodydata)
	if err != nil {
		return nil, err
	}

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)

	// check err
	log.Println(err)
	defer resp.Body.Close() // we have to close Body if we want to reuse the connection

	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	log.Println("Request status ", resp.StatusCode, resp.Status)

	if resp.StatusCode == 200 {
		responseData, errlog := ioutil.ReadAll(resp.Body)
		log.Println("Error Log### ", errlog)
		if errlog != nil {
			return nil, errlog
		}
		json.Unmarshal([]byte(responseData), &data)
		log.Println("unmarshal ", &data)

	} else {
		err = &common.AppError{ErrorCode: (string)(resp.StatusCode), ErrorMsg: resp.Status, ErrorDetail: resp.Status}
		data = nil
	}

	return data, err
}

// HttpPut - HTTP Put Request
func (p *ZapsCloud) HttpPut(requrl string, body map[string]interface{}) (map[string]interface{}, error) {

	log.Println("HttpPut::  Begin ", requrl)
	jsonstring, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodydata := bytes.NewBuffer(jsonstring)
	log.Println("HttpPost:: Body", requrl, bodydata)

	req, err := p.makeReqeust("PUT", requrl, bodydata)
	if err != nil {
		return nil, err
	}

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)

	// check err
	log.Println(err)
	defer resp.Body.Close() // we have to close Body if we want to reuse the connection

	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	log.Println("Request status ", resp.StatusCode, resp.Status)

	if resp.StatusCode == 200 {
		responseData, errlog := ioutil.ReadAll(resp.Body)
		log.Println("HttpPut :: Error Log### ", errlog)
		if errlog != nil {
			return nil, errlog
		}
		json.Unmarshal([]byte(responseData), &data)
		log.Println("HttpPut :: unmarshal ", &data)

	} else {
		err = &common.AppError{ErrorCode: (string)(resp.StatusCode), ErrorMsg: resp.Status, ErrorDetail: resp.Status}
		data = nil
	}

	return data, err
}

// HttpDelete - HTTP Delete Request
func (p *ZapsCloud) HttpDelete(requrl string) (map[string]interface{}, error) {

	log.Println("HttpPut::  Begin ", requrl)

	req, err := p.makeReqeust("DELETE", requrl, nil)
	if err != nil {
		return nil, err
	}

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)

	// check err
	log.Println(err)
	defer resp.Body.Close() // we have to close Body if we want to reuse the connection

	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	log.Println("Request status ", resp.StatusCode, resp.Status)

	if resp.StatusCode == 200 {
		responseData, errlog := ioutil.ReadAll(resp.Body)
		log.Println("HttpPut :: Error Log### ", errlog)
		if errlog != nil {
			return nil, errlog
		}
		json.Unmarshal([]byte(responseData), &data)
		log.Println("HttpPut :: unmarshal ", &data)

	} else {
		err = &common.AppError{ErrorCode: (string)(resp.StatusCode), ErrorMsg: resp.Status, ErrorDetail: resp.Status}
		data = nil
	}

	return data, err
}
