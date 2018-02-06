package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (cdb *CouchDB) getUUIDs(amount int) []string {
	resp, err := http.Get(cdb.url + "_uuids?count=" + strconv.Itoa(amount))
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
	l.Verbose(string(body))
	type UUIDs struct {
		UUIDS []string
	}
	u := UUIDs{}
	err = json.Unmarshal(body, &u)
	if err != nil {
		l.Error("%v", err)
	}
	return u.UUIDS
}
