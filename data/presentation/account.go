package entity

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// It represents an Account on presentation side
type Account struct {
	ID               bson.ObjectId `bson:"_id" json:"id"`
	FirstName        string        `bson:"firstName" json:"firstName"`
	LastName         string        `bson:"lastName" json:"lastName"`
	Company          string        `bson:"companyName" json:"companyName"`
	isSuper          bool          `bson:"isSuper" json:"isSuper"`
	AccountID        string        `bson:"accountId" json:"accountId"`
	LastModifiedDate time.Time     `bson:"lastModifiedDate" json:"lastModifiedDate"`
}
