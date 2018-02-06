package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type DB struct {
	url   string
	Names []string
}

func (cdb *DB) GetAll() []string {
	resp, err := http.Get(cdb.url + "_all_dbs")
	if err != nil {
		l.Error("%v", err)
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			l.Error("%v", err)
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		l.Error("%v", err)
	}
	l.Verbose("Body: %v", string(body))
	var dbs []string
	err = json.Unmarshal(body, &dbs)
	if err != nil {
		l.Error("%v", err)
	}
	cdb.Names = dbs
	return dbs
}

func (cdb *DB) Add(name string) bool {
	client := &http.Client{}
	request, err := http.NewRequest("PUT", cdb.url+name, nil)
	if err != nil {
		l.Error("%v", err)
	}
	resp, err := client.Do(request)
	if err != nil {
		l.Error("%v", err)
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			l.Error("%v", err)
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		l.Error("%v", err)
	}
	l.Verbose("Body: %v", string(body))

	var res map[string]bool
	err = json.Unmarshal(body, &res)
	if err != nil {
		l.Error("%v", err)
	}
	if val, ok := res["ok"]; ok {
		l.Debug("Add DB successful")
		return val
	}
	return false
}
