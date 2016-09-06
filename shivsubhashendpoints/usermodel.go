package shivsubhashendpoints

import (
	//	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	//"google.golang.org\':/ppengine/user"
	//	"log"
	"time"
)

type User struct {
	UserId          *datastore.Key `json:"organiztionName" datastore:"UserId"`
	OrganizationKey *datastore.Key `json:"organiztionName" datastore:"organizationKey"`
	Name            string         `json:"organiztionName" datastore:"Name"`
	Address         string         `json:"organiztionName" json:"organiztionName" datastore:"Address"`
	ContactNumber   string         `json:"organiztionName" datastore:"ContactNumber"`
	CreatedDate     time.Time      `json:"createdDate" datastore:"createdDate"`
	UpdatedDate     time.Time      `json:"updatedDate" datastore:"updatedDate"`
	IsOwner         bool           `json:"isOwner" datastore:"isOwner" endpoints:"d=false"`
	IsPrincipal     bool           `json:"isPrincipal" datastore:"IsPrincipal" endpoints:"d=false"`
	IsTeacher       bool           `json:"isTeacher" datastore:"isTeacher" endpoints:"d=false"`
	IsClerk         bool           `json:"isClerk " datastore:"isClerk" endpoints:"d=false"`
	IsStudent       bool           `json:"isStudent" datastore:"isStudent" endpoints:"d=false"`
}

type UserSaveReq struct {
	Name            string `json:"organiztionName" endpoints:"required"`
	Address         string `json:"organizationAddress" endpoints:"req"`
	ContactNumber   string `json:"contactNumber" endpoints:"required"`
	OrganizationKey string `json:"organizationKey"`
}
