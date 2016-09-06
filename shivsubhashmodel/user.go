package shivsubhashmodel

import (
	//	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	//"google.golang.org/appengine/user"
	//	"log"
)

type User struct {
	UserId          *datastore.Key `datastore:"UserId"`
	Name            string         `datastore:"Name"`
	Address         string         `datastore:"Address"`
	ContactNumber   string         `datastore:"ContactNumber"`
	RoleName        string         `datastore:"RoleName"`
	OrganizationKey *datastore.Key `datastore:"organizationKey"`
}

type UserSaveReq struct {
	Name            string `json:"organiztionName" endpoints:"required"`
	Address         string `json:"organizationAddress" endpoints:"req"`
	ContactNumber   string `json:"contactNumber" endpoints:"required"`
	OrganizationKey string `json:"organizationKey"`
}
