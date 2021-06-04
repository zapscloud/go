package database

import (
	"fmt"
	"log"
	"net/url"
)

func init() {
}

// GetRelationList - Get List of relations
func (p *ZapsDB) GetRelationList(collection string, filterquery string, sortquery string, skip int64, limit int64) (map[string]interface{}, error) {
	log.Println("GetRelationList::  Begin ", filterquery)

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

	requestURL := fmt.Sprintf("/collections/%s/relations%s", collection, queryparam)

	log.Println("Get one ", requestURL)

	responseData, err := p.HttpGet(requestURL)
	if err != nil {
		return nil, err
	}

	return responseData, err
}

// GetRelation - Get Relation
func (p *ZapsDB) GetRelation(collection string, relation string) (map[string]interface{}, error) {
	log.Println("GetOne::  Begin ", collection)

	requestURL := fmt.Sprintf("/collections/%s/relations/%s", collection, relation)

	log.Println("Get one ", requestURL)

	responseData, err := p.HttpGet(requestURL)
	if err != nil {
		return nil, err
	}

	return responseData, err
}

// CreateRelation - Create collection
func (p *ZapsDB) CreateRelation(collectionname string, fieldname string, foreign_collection string, foreign_key string) (map[string]interface{}, error) {
	log.Println("GetOne::  Begin ", collectionname)

	requestURL := fmt.Sprintf("/collections/%s/relations", collectionname)

	log.Println("Get one ", requestURL)

	reqbody := map[string]interface{}{
		"fieldname":          fieldname,
		"foreign_collection": foreign_collection,
		"foreign_key":        foreign_key,
	}

	responseData, err := p.HttpPost(requestURL, reqbody)
	if err != nil {
		return nil, err
	}

	return responseData, err
}

// DeleteOne - Delete Key Record
func (p *ZapsDB) RemoveRelation(collection string, relation string) (map[string]interface{}, error) {
	log.Println("UpdateOne::  Begin ", collection)

	requestURL := fmt.Sprintf("/collections/%s/relations/%s", collection, relation)

	log.Println("Get one ", requestURL)

	responseData, err := p.HttpDelete(requestURL)
	if err != nil {
		return nil, err
	}

	return responseData, err
}
