package doc

import (
	"fmt"
	"github.com/penlook/mgo"
	//"github.com/penlook/mgo/bson"
	//"log"
)

type Doc struct {
	Name     string
	Server   []string
	connection *mgo.Session
}

type Json map[string]interface {}

type Database struct {
	*mgo.Database
}

type Collection struct {
	*mgo.Collection
}

func (doc Doc) New(bson map[string]string) {
	fmt.Println(bson)
}

func (doc Doc) getServers() string {
	servers := ""
	for _, server := range doc.Server {
		servers += server + ","
	}
	return servers
}

func (doc *Doc) SetConnection(connection *mgo.Session) {
	doc.connection = connection
}

func (doc Doc) GetConnection() *mgo.Session{
	return doc.connection
}

func (doc Doc) Connect() Doc {
	connection, err := mgo.Dial(doc.getServers())

	if err != nil {
		panic(err)
	}

	doc.SetConnection(connection)
	doc.connection.SetMode(mgo.Monotonic, true)
	return doc
}

func (doc Doc) Close() {
	doc.GetConnection().Close()
}

func (doc Doc) Database(name string) Database {
	database := Database{doc.connection.DB(name)}
	return database
}

func (db Database) Collection(name string) Collection {
	collection := Collection{db.C(name)}
	return collection
}
