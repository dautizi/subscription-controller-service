package entity

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	ID               bson.ObjectId `bson:"_id" json:"id"`
	AccountID        string        `bson:"accountId" json:"accountId"`
	FirstName        string        `bson:"firstName" json:"firstName"`
	LastName         string        `bson:"lastName" json:"lastName"`
	Company          string        `bson:"companyName" json:"companyName"`
	Email            string        `bson:"email" json:"email"`
	LastModifiedDate time.Time     `bson:"lastModifiedDate" json:"lastModifiedDate"`
}
