package entity

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// It represents a mongo entity
type Subscription struct {
	ID               bson.ObjectId `bson:"_id" json:"id"`
	AccountID        string        `bson:"accountId" json:"accountId"`
	TypeID           string        `bson:"subscriptionTypeId" json:"subscriptionTypeId"`
	Frequency        string        `bson:"frequency" json:"frequency"`
	Active           bool          `bson:"active" json:"active"`
	LastModifiedDate time.Time     `bson:"lastModifiedDate" json:"lastModifiedDate"`
}
