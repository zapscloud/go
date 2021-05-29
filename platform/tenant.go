package app

import (
	"fmt"
	"log"
	"net/url"
)

func init() {
}

// FindByCode - Find By Code
func (p *ZapsPlatform) GetTenantList(filterquery string, sortquery string, skip int64, limit int64) (map[string]interface{}, error) {
	log.Println("GetTenantList::  Begin ", filterquery)

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

	requestURL := fmt.Sprintf("/tenants%s", queryparam)

	log.Println("Get one ", requestURL)

	responseData, err := p.HttpGet(requestURL)
	if err != nil {
		return nil, err
	}

	return responseData, err
}

// GetOne - Get Single Record
func (p *ZapsPlatform) GetTenant(tenantid string) (map[string]interface{}, error) {
	log.Println("GetOne::  Begin ", tenantid)

	requestURL := fmt.Sprintf("/tenants/%s", tenantid)

	log.Println("Get one ", requestURL)

	responseData, err := p.HttpGet(requestURL)
	if err != nil {
		return nil, err
	}

	return responseData, err
}

// CreateTenant - Create collection
func (p *ZapsPlatform) CreateTenant(tenantid, tenantname, tenantregion string) (map[string]interface{}, error) {
	log.Println("GetOne::  Begin ", tenantname)

	requestURL := "/tenants"

	log.Println("Get one ", requestURL)

	reqbody := []byte(`{
		"tenant_id": "` + tenantid + `",
		"tenant_name": "` + tenantname + `",
		"tenant_region": "` + tenantregion + `"
	}`)

	responseData, err := p.HttpPost(requestURL, reqbody)
	if err != nil {
		return nil, err
	}

	return responseData, err
}

// DeleteOne - Delete Key Record
func (p *ZapsPlatform) RemoveTenant(tenantid string) (map[string]interface{}, error) {
	log.Println("UpdateOne::  Begin ", tenantid)

	requestURL := fmt.Sprintf("/tenants/%s", tenantid)

	log.Println("Get one ", requestURL)

	responseData, err := p.HttpDelete(requestURL)
	if err != nil {
		return nil, err
	}

	return responseData, err
}
