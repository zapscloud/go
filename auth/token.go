package auth

import (
	"fmt"
	"log"
)

func init() {
}

// AccessToken - Access Token
func (p *ZapsAuth) AccessToken(reqbody []byte) (map[string]interface{}, error) {
	log.Println("Auth::AccessToken::  Begin ")

	requestURL := fmt.Sprintf("/token")

	log.Println("Get one ", requestURL)
	responseData, err := p.HttpPost(requestURL, reqbody)
	if err != nil {
		return nil, err
	}

	return responseData, err
}

// DeleteOne - Delete Key Record
func (p *ZapsAuth) ValidateToken(token string) (map[string]interface{}, error) {
	log.Println("UpdateOne::  Begin ", token)

	requestURL := fmt.Sprintf("/token/%s", token)

	log.Println("Get one ", requestURL)

	responseData, err := p.HttpPost(requestURL, []byte(""))
	log.Println("Validateion ", responseData)
	if err != nil {
		log.Println("Validateion  error ", err)
		return nil, err
	}

	return responseData, err
}
