package mongo

import (
	"github.com/Lafriakh/log"
	"github.com/go-kira/kon"
	mgo "gopkg.in/mgo.v2"
)

// Mongo - global mongo session.
var Mongo *mgo.Session

// WithMongo - Open return new mongodb session
func WithMongo(configs *kon.Kon) *mgo.Session {
	info := mgo.DialInfo{
		Addrs:    []string{configs.GetString("DB_HOST")},
		Database: configs.GetString("DB_DATABASE"),
		Username: configs.GetString("DB_USERNAME"),
		Password: configs.GetString("DB_PASSWORD"),
	}

	// Connect
	session, err := mgo.DialWithInfo(&info)
	if err != nil {
		log.Panic(err)
	}
	// Optional. Switch the Session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	// append session to global variable.
	Mongo = session

	// return mongo session.
	return session
}

// Insert - to insert data into collection.
func Insert(configs *kon.Kon, name string, docs ...interface{}) error {
	err := Mongo.DB(configs.GetString("DB_DATABASE")).C(name).Insert(docs...)
	return err
}

// Database - return mongo database.
func Database(configs *kon.Kon, database string) *mgo.Database {
	return Mongo.DB(configs.GetString("DB_DATABASE"))
}

// C - return mongo collection.
func C(configs *kon.Kon, name string) *mgo.Collection {
	return Mongo.DB(configs.GetString("DB_DATABASE")).C(name)
}
