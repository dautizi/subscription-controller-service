package repository

import (
	"log"

	. "github.com/Accedo-Products/subscription-controller-service/data/entity"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type SubscriptionRepository struct {
	Server   string
	Database string
}

var subscriptionDB *mgo.Database

const (
	SubscriptionDatabaseID = "subscription"
	SubscriptionCollection = "subscription"
)

// Establish a connection to database
func (m *SubscriptionRepository) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	subscriptionDB = session.DB(m.Database)
}

// Find a subscription by its id
func (m *SubscriptionRepository) GetByID(id string) (Subscription, error) {
	var subscription Subscription
	err := subscriptionDB.C(SubscriptionCollection).FindId(bson.ObjectIdHex(id)).One(&subscription)
	return subscription, err
}

// Find all existing subscriptions
func (m *SubscriptionRepository) GetAll() ([]Subscription, error) {
	var subscriptions []Subscription
	err := subscriptionDB.C(SubscriptionCollection).Find(bson.M{}).All(&subscriptions)
	return subscriptions, err
}

// Create a new subscription into database
func (m *SubscriptionRepository) Insert(subscription Subscription) error {
	err := subscriptionDB.C(SubscriptionCollection).Insert(&subscription)
	return err
}

// Update an existing subscription
func (m *SubscriptionRepository) Update(subscription Subscription) error {
	err := subscriptionDB.C(SubscriptionCollection).UpdateId(subscription.ID, &subscription)
	return err
}

// Delete an existing subscription
func (m *SubscriptionRepository) Delete(subscription Subscription) error {
	err := subscriptionDB.C(SubscriptionCollection).Remove(&subscription)
	return err
}
