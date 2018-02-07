package main

import (
	"bytes"
	"encoding/json"
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

type FindRes struct {
	Docs []interface{}          `json:"docs"`
	_    map[string]interface{} `json:"-"`
}

func NewCouchDB(url string) CouchDB {
	return CouchDB{url, DB{url, []string{}}}
}

func (cdb *CouchDB) Insert(database string, doc string) error {
	uuid := cdb.getUUIDs(1)[0]

	q := []byte(doc)
	req, err := http.NewRequest("PUT", cdb.url+database+"/"+uuid, bytes.NewBuffer(q))
	if err != nil {
		l.Error("%v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
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
	return nil
}

func (cdb *CouchDB) Find(database string, selector string) (FindRes, error) {
	q := []byte(`{ "selector": ` + selector + ` }`)
	req, err := http.NewRequest("POST", cdb.url+database+"/_find", bytes.NewBuffer(q))
	if err != nil {
		l.Error("%v", err)
		return FindRes{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		l.Error("%v", err)
		return FindRes{}, err
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
		return FindRes{}, err
	}
	l.Verbose("Body: %v", string(body))
	var res FindRes
	err = json.Unmarshal(body, &res)
	if err != nil {
		l.Error("%v", err)
		return FindRes{}, err
	}
	return res, nil
}

func (cdb *CouchDB) FindAll(database string) error {
	resp, err := http.Get(cdb.url + database + "/_all_docs")
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
	return nil
}

func main() {
	l = logger.NewLogger(logger.LEVELINFO)
	l.Verbose("test")
	// cdb := NewCouchDB("http://localhost:5984/")
	// _ = cdb.DB.GetAll()
	// _, _ = cdb.Find("test", `{ "test": { "$eq": "asdf" } }`)
	// _ = cdb.Insert("test", `{ "test": "asdf" }`)
	// _ = cdb.FindAll("test")
	// l.Debug("%s", cdb.DB.Names)
	// _ = cdb.AddDB("test1")
}
