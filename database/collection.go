package database

import (
	"fmt"
	"log"
	"net/url"
)

func init() {
}

// FindByCode - Find By Code
func (p *ZapsDB) GetCollectionList(filterquery string, sortquery string, skip int64, limit int64) (map[string]interface{}, error) {
	log.Println("GetCollectionList::  Begin ", filterquery)

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

	requestURL := fmt.Sprintf("/collections%s", queryparam)

	log.Println("Get one ", requestURL)

	responseData, err := p.HttpGet(requestURL)
	if err != nil {
		return nil, err
	}

	return responseData, err
}

// GetOne - Get Single Record
func (p *ZapsDB) GetCollection(collectionname string) (map[string]interface{}, error) {
	log.Println("GetOne::  Begin ", collectionname)

	requestURL := fmt.Sprintf("/collections/%s", collectionname)

	log.Println("Get one ", requestURL)

	responseData, err := p.HttpGet(requestURL)
	if err != nil {
		return nil, err
	}

	return responseData, err
}

// CreateCollection - Create collection
func (p *ZapsDB) CreateCollection(collectionname string, collectionkey string, description string) (map[string]interface{}, error) {
	log.Println("GetOne::  Begin ", collectionname)

	requestURL := fmt.Sprintf("/collections")

	log.Println("Get one ", requestURL)

	reqbody := []byte(`{
		"collection_id": "` + collectionname + `",
		"collection_name": "` + description + `",
		"collection_key": "` + collectionkey + `"
	}`)

	responseData, err := p.HttpPost(requestURL, reqbody)
	if err != nil {
		return nil, err
	}

	return responseData, err
}

// DeleteOne - Delete Key Record
func (p *ZapsDB) RemoveCollection(collectionname string) (map[string]interface{}, error) {
	log.Println("UpdateOne::  Begin ", collectionname)

	requestURL := fmt.Sprintf("/collections/%s", collectionname)

	log.Println("Get one ", requestURL)

	responseData, err := p.HttpDelete(requestURL)
	if err != nil {
		return nil, err
	}

	return responseData, err
}
