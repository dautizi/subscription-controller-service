package main

import (
	"encoding/json"
	"log"
	"net/http"

	. "github.com/Accedo-Products/subscription-controller-service/config"
	. "github.com/Accedo-Products/subscription-controller-service/data/entity"
	. "github.com/Accedo-Products/subscription-controller-service/repository"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var mongoConfig = MongoConfig{}
var subscriptionRepository = SubscriptionRepository{}
var accountRepository = AccountRepository{}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	account, err := accountRepository.GetByAccountID(params["accountId"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Account ID")
		return
	}
	respondWithJson(w, http.StatusOK, account)
}

func AllSubscriptions(w http.ResponseWriter, r *http.Request) {
	subscriptions, err := subscriptionRepository.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, subscriptions)
}

func GetSubscription(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	subscription, err := subscriptionRepository.GetByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Subscription ID")
		return
	}
	respondWithJson(w, http.StatusOK, subscription)
}

func CreateSubscription(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var subscription Subscription

	// Deserialize subscription from request
	if err := json.NewDecoder(r.Body).Decode(&subscription); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	subscription.ID = bson.NewObjectId()
	if err := subscriptionRepository.Insert(subscription); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, subscription)
}

func UpdateSubscription(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var subscription Subscription
	if err := json.NewDecoder(r.Body).Decode(&subscription); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := subscriptionRepository.Update(subscription); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteSubscription(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var subscription Subscription
	if err := json.NewDecoder(r.Body).Decode(&subscription); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := subscriptionRepository.Delete(subscription); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	// "_" is called "blank identifier" and it avoids having to declare all
	// the variables for the returns values.
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	// mongoConfig.Read()
	Initialise()

	accountDBconfig, err := LookupDatabase(AccountDatabaseID)
	log.Printf("Loaded repository %s", accountDBconfig, err)

	subscriptionDBconfig, err := LookupDatabase(SubscriptionDatabaseID)
	log.Printf("Loaded repository %s", subscriptionDBconfig, err)

	accountRepository.Server = accountDBconfig.Server
	accountRepository.Database = accountDBconfig.Name
	accountRepository.Connect()

	subscriptionRepository.Server = subscriptionDBconfig.Server
	subscriptionRepository.Database = subscriptionDBconfig.Name
	subscriptionRepository.Connect()
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/accounts/{accountId}", GetAccount).Methods("GET")

	r.HandleFunc("/subscriptions", AllSubscriptions).Methods("GET")
	r.HandleFunc("/subscriptions", CreateSubscription).Methods("POST")
	r.HandleFunc("/subscriptions", UpdateSubscription).Methods("PUT")
	r.HandleFunc("/subscriptions", DeleteSubscription).Methods("DELETE")
	r.HandleFunc("/subscriptions/{id}", GetSubscription).Methods("GET")

	if err := http.ListenAndServe(":3001", r); err != nil {
		log.Fatal(err)
	}
}
