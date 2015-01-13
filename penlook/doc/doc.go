// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Loi Nguyen <loint@penlook.com>

package doc

import (
	"github.com/penlook/mgo"
)

type Doc struct {
	Name     string
	Server   []string
	connection *mgo.Session
}

// JSON Format for input data
type Json map[string]interface {}

// MongoDB Database
type Database struct {
	*mgo.Database
}

// MongoDB Collection
type Collection struct {
	*mgo.Collection
}

// Get multiple servers IP & Port
func (doc Doc) getServers() string {
	servers := ""
	for _, server := range doc.Server {
		servers += server + ","
	}
	return servers
}

// Set connection
func (doc *Doc) SetConnection(connection *mgo.Session) {
	doc.connection = connection
}

// Get connection
func (doc Doc) GetConnection() *mgo.Session{
	return doc.connection
}

// Establish new connection
func (doc Doc) Connect() Doc {
	connection, err := mgo.Dial(doc.getServers())

	if err != nil {
		panic(err)
	}

	doc.SetConnection(connection)
	doc.connection.SetMode(mgo.Monotonic, true)
	return doc
}

// Close connection
func (doc Doc) Close() {
	doc.GetConnection().Close()
}

// Choose database
func (doc Doc) Database(name string) Database {
	database := Database{doc.connection.DB(name)}
	return database
}

// Choose collection
func (db Database) Collection(name string) Collection {
	collection := Collection{db.C(name)}
	return collection
}
