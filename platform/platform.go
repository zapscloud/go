package app

import (
	"fmt"
	"log"

	"github.com/zapscloud/golib/zaps"
)

// ZapsPlatform - Document Service structure
type ZapsPlatform struct {
	zaps    *zaps.ZapsCloud
	baseURL string
}

func init() {
}

// NewZapsPlatform - Construct ZapsPlatform
func NewZapsPlatform(zc *zaps.ZapsCloud) (*ZapsPlatform, error) {
	p := ZapsPlatform{}
	p.zaps = zc

	log.Println("NewZapsPlatform ", p)

	p.baseURL = fmt.Sprintf("https://platform.api%s.zapscloud.com", zc.Stage())

	return &p, nil
}

// EndZapsPlatform - Close all the ZapsPlatform
func (p *ZapsPlatform) EndZapsPlatform() {
	log.Printf("EndZapsPlatform ")
}

// HttpGet - HTTP Get Request
func (p *ZapsPlatform) HttpGet(requrl string) (map[string]interface{}, error) {

	log.Println("ZapsPlatform:HttpGet::  Begin ", requrl)

	requestURL := p.baseURL + requrl
	resp, err := p.zaps.HttpGet(requestURL)

	log.Println("ZapsPlatform:HttpGet::  End ", resp, err)

	return resp, err
}

// HttpGet - HTTP Get Request
func (p *ZapsPlatform) HttpPost(requrl string, body interface{}) (map[string]interface{}, error) {

	log.Println("ZapsPlatform:HttpPost::  Begin ", requrl)

	requestURL := p.baseURL + requrl
	resp, err := p.zaps.HttpPost(requestURL, body)

	log.Println("ZapsPlatform:HttpPost::  End ", resp, err)

	return resp, err
}

// HttpPut - HTTP Put Request
func (p *ZapsPlatform) HttpPut(requrl string, body interface{}) (map[string]interface{}, error) {

	log.Println("ZapsPlatform:HttpPut::  Begin ", requrl)

	requestURL := p.baseURL + requrl
	resp, err := p.zaps.HttpPut(requestURL, body)

	log.Println("ZapsPlatform:HttpPut::  End ", resp, err)

	return resp, err
}

// HttpDelete - HTTP Delete Request
func (p *ZapsPlatform) HttpDelete(requrl string) (map[string]interface{}, error) {

	log.Println("ZapsPlatform:HttpDelete::  Begin ", requrl)

	requestURL := p.baseURL + requrl
	resp, err := p.zaps.HttpDelete(requestURL)

	log.Println("ZapsPlatform:HttpDelete::  End ", resp, err)

	return resp, err
}
