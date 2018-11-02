package entity

import (
	"time"
)

// It represents a "One" Account on presentation side
type OneAccount struct {
	ID               string    `bson:"_id" json:"id"`
	FirstName        string    `bson:"firstName" json:"firstName"`
	LastName         string    `bson:"lastName" json:"lastName"`
	Company          string    `bson:"companyName" json:"companyName"`
	Email            string    `bson:"email" json:"email"`
	CreatedDate      time.Time `bson:"createdDate" json:"createdDate"`
	LastModifiedDate time.Time `bson:"lastModifiedDate" json:"lastModifiedDate"`
}
