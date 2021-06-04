package database

import (
	"fmt"
	"log"
	"net/url"
)

func init() {
}

// FindByCode - Find By Code
func (p *ZapsDB) GetMany(collectionname string, filterquery string, sortquery string, skip int64, limit int64) (map[string]interface{}, error) {
	log.Println("GetMany::  Begin ", collectionname, filterquery)

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

	requestURL := fmt.Sprintf("/documents/%s%s", collectionname, queryparam)

	log.Println("Get one ", requestURL)

	responseData, err := p.HttpGet(requestURL)
	if err != nil {
		return nil, err
	}

	return responseData, err
}

// GetOne - Get Single Record
func (p *ZapsDB) GetOne(collectionname string, collectionkey string, lookups string) (map[string]interface{}, error) {
	log.Println("GetOne::  Begin ", collectionname, collectionkey, lookups)

	var lookupstring string
	if lookups == "" {
		lookupstring = ""
	} else {
		lookupstring = "?lookup=" + lookups
	}

	requestURL := fmt.Sprintf("/documents/%s/%s%s", collectionname, collectionkey, lookupstring)

	log.Println("Get one ", requestURL)

	responseData, err := p.HttpGet(requestURL)
	if err != nil {
		return nil, err
	}

	return responseData, err
}

// Insert - Insert Single Record
func (p *ZapsDB) Insert(collectionname string, body map[string]interface{}) (map[string]interface{}, error) {
	log.Println("GetOne::  Begin ", collectionname)

	requestURL := fmt.Sprintf("/documents/%s", collectionname)

	log.Println("Get one ", requestURL)

	responseData, err := p.HttpPost(requestURL, body)
	if err != nil {
		return nil, err
	}

	return responseData, err
}

// UpdateOne - Update Key Record
func (p *ZapsDB) UpdateOne(collectionname string, collectionkey string, body map[string]interface{}) (map[string]interface{}, error) {
	log.Println("UpdateOne::  Begin ", collectionname)

	requestURL := fmt.Sprintf("/documents/%s/%s", collectionname, collectionkey)

	log.Println("Get one ", requestURL)

	responseData, err := p.HttpPut(requestURL, body)
	if err != nil {
		return nil, err
	}

	return responseData, err
}

// Update - Update Record
func (p *ZapsDB) UpdateMany(collectionname string, filterquery string, body map[string]interface{}) (map[string]interface{}, error) {
	log.Println("GetOne::  Begin ", collectionname)

	queryparam := ""
	if filterquery != "" {
		queryparam = "?filter=" + filterquery
	}

	requestURL := fmt.Sprintf("/documents/%s%s", collectionname, queryparam)

	log.Println("Get one ", requestURL)

	responseData, err := p.HttpPut(requestURL, body)
	if err != nil {
		return nil, err
	}

	return responseData, err
}

// DeleteOne - Delete Key Record
func (p *ZapsDB) DeleteOne(collectionname string, collectionkey string) (map[string]interface{}, error) {
	log.Println("UpdateOne::  Begin ", collectionname)

	requestURL := fmt.Sprintf("/documents/%s/%s", collectionname, collectionkey)

	log.Println("Get one ", requestURL)

	responseData, err := p.HttpDelete(requestURL)
	if err != nil {
		return nil, err
	}

	return responseData, err
}

// DeleteMany - Delete Many Records
func (p *ZapsDB) DeleteMany(collectionname string, filterquery string) (map[string]interface{}, error) {
	log.Println("GetOne::  Begin ", collectionname)

	queryparam := ""
	if filterquery != "" {
		queryparam = "?filter=" + url.QueryEscape(filterquery)
	}

	requestURL := fmt.Sprintf("/documents/%s%s", collectionname, queryparam)

	log.Println("Get one ", requestURL)

	responseData, err := p.HttpDelete(requestURL)
	if err != nil {
		return nil, err
	}

	return responseData, err
}

// Aggregate - Aggregate String
func (p *ZapsDB) Aggregate(collectionname string, filterquery string, aggquery string, sortquery string, skip int64, limit int64) (map[string]interface{}, error) {
	log.Println("GetMany::  Begin ", collectionname, filterquery)

	queryparam := ""
	if filterquery != "" {
		queryparam = "?filter=" + url.QueryEscape(filterquery)
	}

	if aggquery != "" {
		if queryparam != "" {
			queryparam = fmt.Sprintf("%s&aggregate=%s", queryparam, url.QueryEscape(aggquery))
		} else {
			queryparam = fmt.Sprintf("?aggregate=%s", url.QueryEscape(aggquery))
		}
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

	requestURL := fmt.Sprintf("/documents/%s%s", collectionname, queryparam)

	log.Println("Get one ", requestURL)

	responseData, err := p.HttpGet(requestURL)
	if err != nil {
		return nil, err
	}

	return responseData, err
}
