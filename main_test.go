package main

import (
	"github.com/pmpbar/log"
	"testing"
)

func TestInsert(t *testing.T) {
	l = logger.NewLogger(logger.LEVELINFO)
	cdb := NewCouchDB("http://localhost:5984/")
	err := cdb.Insert("test", `{ "test": "asdf" }`)
	if err != nil {
		t.Error(err)
	} else {
		l.Test("Insert success")
	}
}

func TestFind(t *testing.T) {
	l = logger.NewLogger(logger.LEVELVERBOSE)
	cdb := NewCouchDB("http://localhost:5984/")
	err := cdb.Insert("test", `{ "test": "asdf" }`)
	if err != nil {
		t.Error(err)
	} else {
		res, err := cdb.Find("test", `{ "test": { "$eq": "asdf" } }`)
		if err != nil {
			t.Error(err)
		}
		if len(res.Docs) == 0 {
			t.Error("No documents found")
		} else {
			l.Test("Find success")
		}
	}
}

func TestGetDBs(t *testing.T) {}
