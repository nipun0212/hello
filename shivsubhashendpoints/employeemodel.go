package shivsubhashendpoints

import (
	"google.golang.org/appengine/datastore"
)

type Employee struct {
	key             *datastore.Key `json:"organizationKey" datastore:"-"`
	OrganizationKey *datastore.Key `json:"organiztionName" datastore:"organizationKey"`
	EmpID           int            `json:"EmpID" datastore:"EmpID"`
	EmailID         string         `json:"EmailID" datastore:"EmailID"`
	Name            string         `json:"Name" datastore:"Name"`
	Age             int            `json:"Age" datastore:"Age"`
	Gender          bool           `json:"Gender" datastore:"Gender"`
	IsOwner         bool           `json:"isOwner" datastore:"isOwner" endpoints:"d=false"`
	IsPrincipal     bool           `json:"isPrincipal" datastore:"IsPrincipal" endpoints:"d=false"`
	IsTeacher       bool           `json:"isTeacher" datastore:"isTeacher" endpoints:"d=false"`
	IsClerk         bool           `json:"isClerk " datastore:"isClerk" endpoints:"d=false"`
	IsApproved      bool           `json:"isApproved " datastore:"isApproved" endpoints:"d=false"`
}
