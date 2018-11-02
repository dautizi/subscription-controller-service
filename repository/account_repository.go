package repository

import (
	"log"

	. "github.com/Accedo-Products/subscription-controller-service/data/entity"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AccountRepository struct {
	Server   string
	Database string
}

var accountDB *mgo.Database

const (
	AccountDatabaseID = "account"
	AccountCollection = "account"
)

// Establish a connection to database
func (m *AccountRepository) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	accountDB = session.DB(m.Database)
}

// Find a account by its id
func (m *AccountRepository) GetByID(id string) (Account, error) {
	var account Account
	err := accountDB.C(AccountCollection).FindId(bson.ObjectIdHex(id)).One(&account)
	return account, err
}

// Find a account by its accountId
func (m *AccountRepository) GetByAccountID(accountId string) (Account, error) {
	var account Account
	err := accountDB.C(AccountCollection).Find(bson.M{accountId: accountId}).One(&account)
	return account, err
}
