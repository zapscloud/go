package database

import (
	"fmt"
	"log"

	"github.com/zapscloud/golib/zaps"
)

// ZapsDB - Document Service structure
type ZapsDB struct {
	zaps    *zaps.ZapsCloud
	baseURL string
}

func init() {
}

// NewZapsDB - Construct ZapsDB
func NewZapsDB(zc *zaps.ZapsCloud) (*ZapsDB, error) {
	p := ZapsDB{}
	p.zaps = zc

	log.Println("NewZapsDB ", p)

	p.baseURL = fmt.Sprintf("https://database.api%s.zapscloud.com", zc.Stage())

	return &p, nil
}

// EndZapsDB - Close all the ZapsDB
func (p *ZapsDB) EndZapsDB() {
	log.Printf("EndZapsDB ")
}

// HttpGet - HTTP Get Request
func (p *ZapsDB) HttpGet(requrl string) (map[string]interface{}, error) {

	log.Println("ZapsDB:HttpGet::  Begin ", requrl)

	requestURL := p.baseURL + requrl
	resp, err := p.zaps.HttpGet(requestURL)

	log.Println("ZapsDB:HttpGet::  End ", resp, err)

	return resp, err
}

// HttpGet - HTTP Get Request
func (p *ZapsDB) HttpPost(requrl string, body map[string]interface{}) (map[string]interface{}, error) {

	log.Println("ZapsDB:HttpPost::  Begin ", requrl)

	requestURL := p.baseURL + requrl
	resp, err := p.zaps.HttpPost(requestURL, body)

	log.Println("ZapsDB:HttpPost::  End ", resp, err)

	return resp, err
}

// HttpPut - HTTP Put Request
func (p *ZapsDB) HttpPut(requrl string, body map[string]interface{}) (map[string]interface{}, error) {

	log.Println("ZapsDB:HttpPut::  Begin ", requrl)

	requestURL := p.baseURL + requrl
	resp, err := p.zaps.HttpPut(requestURL, body)

	log.Println("ZapsDB:HttpPut::  End ", resp, err)

	return resp, err
}

// HttpDelete - HTTP Delete Request
func (p *ZapsDB) HttpDelete(requrl string) (map[string]interface{}, error) {

	log.Println("ZapsDB:HttpDelete::  Begin ", requrl)

	requestURL := p.baseURL + requrl
	resp, err := p.zaps.HttpDelete(requestURL)

	log.Println("ZapsDB:HttpDelete::  End ", resp, err)

	return resp, err
}
