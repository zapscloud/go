package auth

import (
	"fmt"
	"log"
	"net/url"
)

func init() {
}

// GetClientList - Get clients list
func (p *ZapsAuth) GetClientList(filterquery string, sortquery string, skip int64, limit int64) (map[string]interface{}, error) {
	log.Println("Auth::GetClientList:  Begin ", filterquery)

	queryparam := ""
	if filterquery != "" {
		queryparam = "?filter=" + url.QueryEscape(filterquery)
	}

	if sortquery != "" {
		if queryparam != "" {
			queryparam = fmt.Sprintf("%s&sort=%s", queryparam, url.QueryEscape(sortquery))
		} else {
			queryparam = fmt.Sprintf("?sort=%s", url.QueryEscape(sortquery))
		}
	}
	if skip != 0 {
		if queryparam != "" {
			queryparam = fmt.Sprintf("%s&skip=%d", queryparam, skip)
		} else {
			queryparam = fmt.Sprintf("?skip=%d", skip)
		}
	}

	if limit != 0 {
		if queryparam != "" {
			queryparam = fmt.Sprintf("%s&limit=%d", queryparam, limit)
		} else {
			queryparam = fmt.Sprintf("?limit=%d", limit)
		}
	}

	requestURL := fmt.Sprintf("/clients%s", queryparam)

	log.Println("Get one ", requestURL)

	responseData, err := p.HttpGet(requestURL)
	if err != nil {
		return nil, err
	}

	return responseData, err
}

// GetClient - Get Single Record
func (p *ZapsAuth) GetClient(clientid string) (map[string]interface{}, error) {
	log.Println("Auth::GetClient:  Begin ", clientid)

	requestURL := fmt.Sprintf("/clients/%s", clientid)

	log.Println("Get one ", requestURL)

	responseData, err := p.HttpGet(requestURL)
	if err != nil {
		return nil, err
	}

	log.Println("Auth::GetClient:  End ")

	return responseData, err
}

// CreateClient - Create collection
func (p *ZapsAuth) CreateClient(reqbody interface{}) (map[string]interface{}, error) {
	log.Println("Auth::CreateClient:  Begin ")

	requestURL := "/clients"

	log.Println("Get one ", requestURL)
	responseData, err := p.HttpPost(requestURL, reqbody)
	if err != nil {
		return nil, err
	}
	log.Println("Auth::CreateClient:  End ")

	return responseData, err
}

// DeleteOne - Delete Key Record
func (p *ZapsAuth) RemoveClient(clientid string) (map[string]interface{}, error) {
	log.Println("Auth::RemoveClient:  Begin ", clientid)

	requestURL := fmt.Sprintf("/clients/%s", clientid)

	log.Println("Get one ", requestURL)

	responseData, err := p.HttpDelete(requestURL)
	if err != nil {
		return nil, err
	}
	log.Println("Auth::RemoveClient:  End ")

	return responseData, err
}

// ValidateClient - Validate client
func (p *ZapsAuth) ValidateClient(clientkey, clientsecret string) (map[string]interface{}, error) {
	log.Println("Auth::ValidateClient:  Begin ", clientkey)

	requestURL := "/clients/validate"
	log.Println("Get one ", requestURL)

	reqbody := map[string]interface{}{
		"client_key":    clientkey,
		"client_secret": clientsecret,
	}

	responseData, err := p.HttpPost(requestURL, reqbody)
	if err != nil {
		return nil, err
	}
	log.Println("Auth::ValidateClient:  End ")

	return responseData, err
}
