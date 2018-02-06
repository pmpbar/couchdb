package main

import (
	"bytes"
	// "encoding/json"
	"github.com/pmpbar/log"
	"io/ioutil"
	"net/http"
	// "net/url"
	// "strings"
)

var l logger.Logger

type CouchDB struct {
	url string
	DB  DB
}

func NewCouchDB(url string) CouchDB {
	return CouchDB{url, DB{url, []string{}}}
}

func (cdb *CouchDB) Find(database string, selector string) {
	q := []byte(`{ "selector": ` + selector + ` }`)
	req, err := http.NewRequest("POST", cdb.url+database+"/_find", bytes.NewBuffer(q))
	if err != nil {
		l.Error("%v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
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
	/* var dbs []string
	err = json.Unmarshal(body, &dbs)
	if err != nil {
		l.Error("%v", err)
	} */
}

func (cdb *CouchDB) FindAll(database string) {
	resp, err := http.Get(cdb.url + database + "/_all_docs")
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
	/* var dbs []string
	err = json.Unmarshal(body, &dbs)
	if err != nil {
		l.Error("%v", err)
	} */
}

func main() {
	l = logger.NewLogger(logger.LEVELVERBOSE)
	cdb := NewCouchDB("http://localhost:5984/")
	_ = cdb.DB.GetAll()
	cdb.Find("test", `{ "test": { "$eq": "asdf" } }`)
	// cdb.FindAll("test")
	// l.Debug("%s", cdb.DB.Names)
	// cdb.AddDB("test1")
}
