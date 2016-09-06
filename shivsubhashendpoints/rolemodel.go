package shivsubhashendpoints

import (
	//	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	//"google.golang.org/appengine/user"
	//	"log"
	"time"
)

type Roles struct {
	key         *datastore.Key
	Name        string    `json:"Name" datastore:"Name"`
	ID          string    `json:"ID" datastore:"ID"`
	Description string    `json:"Description" datastore:"Description"`
	CreatedDate time.Time `json:"CreatedDate" datastore:"CreatedDate"`
	UpdatedDate time.Time `json:"UpdatedDate" datastore:"UpdatedDate"`
}
