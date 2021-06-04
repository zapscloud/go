package auth

import (
	"fmt"
	"log"
	"net/url"
)

func init() {
}

// GetUserList -Get list of users
func (p *ZapsAuth) GetUserList(filterquery string, sortquery string, skip int64, limit int64) (map[string]interface{}, error) {
	log.Println("Auth::GetUserList:  Begin ", filterquery)

	queryparam := ""
	if filterquery != "" {
		queryparam = "?" + url.QueryEscape(filterquery)
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

	requestURL := fmt.Sprintf("/users%s", queryparam)

	log.Println("Get one ", requestURL)

	responseData, err := p.HttpGet(requestURL)
	if err != nil {
		return nil, err
	}

	log.Println("Auth::GetUserList:  End ")

	return responseData, err
}

// GetUser - Get User Details
func (p *ZapsAuth) GetUser(tenantid string) (map[string]interface{}, error) {
	log.Println("Auth::GetUser:  Begin ", tenantid)

	requestURL := fmt.Sprintf("/users/%s", tenantid)

	log.Println("Get one ", requestURL)

	responseData, err := p.HttpGet(requestURL)
	if err != nil {
		return nil, err
	}

	log.Println("Auth::GetUser:  End ")
	return responseData, err
}

// CreateUser - Create new user
func (p *ZapsAuth) CreateUser(reqbody interface{}) (map[string]interface{}, error) {
	log.Println("Auth::CreateUser:  Begin ")

	requestURL := "/users"

	log.Println("Get one ", requestURL)
	responseData, err := p.HttpPost(requestURL, reqbody)
	if err != nil {
		return nil, err
	}

	return responseData, err
}

// UpdateUser - Update user
func (p *ZapsAuth) UpdateUser(userid string, reqbody interface{}) (map[string]interface{}, error) {
	log.Println("Auth::UpdateUser:  Begin ", userid)

	requestURL := fmt.Sprintf("/users/%s", userid)

	log.Println("Update one ", requestURL)
	responseData, err := p.HttpPut(requestURL, reqbody)
	if err != nil {
		return nil, err
	}

	log.Println("Auth::UpdateUser:  End ")

	return responseData, err
}

// RemoveUser - Delete Key Record
func (p *ZapsAuth) RemoveUser(tenantid string) (map[string]interface{}, error) {
	log.Println("UpdateOne::  Begin ", tenantid)

	requestURL := fmt.Sprintf("/users/%s", tenantid)

	log.Println("Get one ", requestURL)

	responseData, err := p.HttpDelete(requestURL)
	if err != nil {
		return nil, err
	}

	return responseData, err
}

// ValidateUser - Validate user
func (p *ZapsAuth) ValidateUser(userid, usersecret string) (map[string]interface{}, error) {
	log.Println("Auth::ValidateUser:  Begin ", userid)

	requestURL := "/users/validate"
	log.Println("Get one ", requestURL)

	reqbody := map[string]interface{}{
		"user_id":     userid,
		"user_secret": usersecret,
	}

	responseData, err := p.HttpPost(requestURL, reqbody)
	if err != nil {
		return nil, err
	}
	log.Println("Auth::ValidateUser:  End ")

	return responseData, err
}

// ValidateUser - Validate user
func (p *ZapsAuth) ResetPassword(reqbody interface{}) (map[string]interface{}, error) {
	log.Println("Auth::ValidateUser:  Begin ")

	requestURL := "/users/secret"
	log.Println("Get one ", requestURL)

	responseData, err := p.HttpPut(requestURL, reqbody)
	if err != nil {
		return nil, err
	}
	log.Println("Auth::ValidateUser:  End ")

	return responseData, err
}
