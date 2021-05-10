package app

import (
	"fmt"
	"log"

	"github.com/zapscloud/golib/zaps"
)

// ZapsApp - Document Service structure
type ZapsApp struct {
	zaps    *zaps.ZapsCloud
	baseURL string
}

func init() {
}

// NewZapsApp - Construct ZapsApp
func NewZapsApp(zc *zaps.ZapsCloud) (*ZapsApp, error) {
	p := ZapsApp{}
	p.zaps = zc

	log.Println("NewZapsApp ", p)

	p.baseURL = fmt.Sprintf("https://setup.%sapi.zapscloud.com", zc.Stage())

	return &p, nil
}

// EndZapsApp - Close all the ZapsApp
func (p *ZapsApp) EndZapsApp() {
	log.Printf("EndZapsApp ")
	return
}

// HttpGet - HTTP Get Request
func (p *ZapsApp) HttpGet(requrl string) (map[string]interface{}, error) {

	log.Println("ZapsApp:HttpGet::  Begin ", requrl)

	requestURL := p.baseURL + requrl
	resp, err := p.zaps.HttpGet(requestURL)

	log.Println("ZapsApp:HttpGet::  End ", resp, err)

	return resp, err
}

// HttpGet - HTTP Get Request
func (p *ZapsApp) HttpPost(requrl string, body []byte) (map[string]interface{}, error) {

	log.Println("ZapsApp:HttpPost::  Begin ", requrl)

	requestURL := p.baseURL + requrl
	resp, err := p.zaps.HttpPost(requestURL, body)

	log.Println("ZapsApp:HttpPost::  End ", resp, err)

	return resp, err
}

// HttpPut - HTTP Put Request
func (p *ZapsApp) HttpPut(requrl string, body []byte) (map[string]interface{}, error) {

	log.Println("ZapsApp:HttpPut::  Begin ", requrl)

	requestURL := p.baseURL + requrl
	resp, err := p.zaps.HttpPut(requestURL, body)

	log.Println("ZapsApp:HttpPut::  End ", resp, err)

	return resp, err
}

// HttpDelete - HTTP Delete Request
func (p *ZapsApp) HttpDelete(requrl string) (map[string]interface{}, error) {

	log.Println("ZapsApp:HttpDelete::  Begin ", requrl)

	requestURL := p.baseURL + requrl
	resp, err := p.zaps.HttpDelete(requestURL)

	log.Println("ZapsApp:HttpDelete::  End ", resp, err)

	return resp, err
}
