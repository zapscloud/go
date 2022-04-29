package database

import (
	"fmt"
	"log"
)

func init() {
}

// FindByCode - Find By Code
func (p *ZapsDB) StartTransaction() (map[string]interface{}, error) {
	log.Println("startTransaction::  Begin ")

	requestURL := "/transactions"
	log.Println("startTransaction one ", requestURL)

	responseData, err := p.HttpPost(requestURL, map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	return responseData, err
}

// GetOne - Get Single Record
func (p *ZapsDB) CommitTransaction(zapstxn string) (map[string]interface{}, error) {
	log.Println("commitTransaction::  Begin ", zapstxn)

	requestURL := fmt.Sprintf("/transactions/%s/commit", zapstxn)
	log.Println("commitTransaction ", requestURL)

	responseData, err := p.HttpPost(requestURL, nil)
	if err != nil {
		return nil, err
	}

	return responseData, err
}

// DeleteOne - Delete Key Record
func (p *ZapsDB) RollbackTransaction(zapstxn string) (map[string]interface{}, error) {
	log.Println("rollbackTransaction::  Begin ", zapstxn)

	requestURL := fmt.Sprintf("/transactions/%s/rollback", zapstxn)

	log.Println("rollbackTransaction", requestURL)

	responseData, err := p.HttpPost(requestURL, map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	return responseData, err
}
