package auth

import (
	"fmt"
	"log"

	"github.com/zapscloud/golib/zaps"
)

// ZapsAuth - Document Service structure
type ZapsAuth struct {
	zaps    *zaps.ZapsCloud
	baseURL string
}

func init() {
}

// NewZapsAuth - Construct ZapsAuth
func NewZapsAuth(zc *zaps.ZapsCloud) (*ZapsAuth, error) {
	p := ZapsAuth{}
	p.zaps = zc

	log.Println("NewZapsAuth ", p)

	p.baseURL = fmt.Sprintf("https://auth.api%s.zapscloud.com", zc.Stage())

	return &p, nil
}

// EndZapsAuth - Close all the ZapsAuth
func (p *ZapsAuth) EndZapsAuth() {
	log.Printf("EndZapsAuth ")
	return
}

// HttpGet - HTTP Get Request
func (p *ZapsAuth) HttpGet(requrl string) (map[string]interface{}, error) {

	log.Println("ZapsAuth:HttpGet::  Begin ", requrl)

	requestURL := p.baseURL + requrl
	resp, err := p.zaps.HttpGet(requestURL)

	log.Println("ZapsAuth:HttpGet::  End ", resp, err)

	return resp, err
}

// HttpGet - HTTP Get Request
func (p *ZapsAuth) HttpPost(requrl string, body []byte) (map[string]interface{}, error) {

	log.Println("ZapsAuth:HttpPost::  Begin ", requrl)

	requestURL := p.baseURL + requrl
	resp, err := p.zaps.HttpPost(requestURL, body)

	log.Println("ZapsAuth:HttpPost::  End ", resp, err)

	return resp, err
}

// HttpPut - HTTP Put Request
func (p *ZapsAuth) HttpPut(requrl string, body []byte) (map[string]interface{}, error) {

	log.Println("ZapsAuth:HttpPut::  Begin ", requrl)

	requestURL := p.baseURL + requrl
	resp, err := p.zaps.HttpPut(requestURL, body)

	log.Println("ZapsAuth:HttpPut::  End ", resp, err)

	return resp, err
}

// HttpDelete - HTTP Delete Request
func (p *ZapsAuth) HttpDelete(requrl string) (map[string]interface{}, error) {

	log.Println("ZapsAuth:HttpDelete::  Begin ", requrl)

	requestURL := p.baseURL + requrl
	resp, err := p.zaps.HttpDelete(requestURL)

	log.Println("ZapsAuth:HttpDelete::  End ", resp, err)

	return resp, err
}
