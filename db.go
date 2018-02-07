package main

import (
	"encoding/json"
	"errors"
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

func (cdb *DB) Add(name string) error {
	client := &http.Client{}
	request, err := http.NewRequest("PUT", cdb.url+name, nil)
	if err != nil {
		l.Error("%v", err)
		return err
	}
	resp, err := client.Do(request)
	if err != nil {
		l.Error("%v", err)
		return err
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
		return err
	}
	l.Verbose("Body: %v", string(body))

	var res map[string]interface{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		l.Error("%v", err)
		return err
	}
	if _, ok := res["ok"].(bool); ok {
		l.Debug("Add DB successful")
		return nil
	}
	return errors.New(res["error"].(string))
}
