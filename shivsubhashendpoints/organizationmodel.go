package shivsubhashendpoints

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"log"
	"reflect"
	"time"
)

type Organization struct {
	key           *datastore.Key `json:"organizationKey" datastore:"-"`
	Name          string         `json:"organiztionName" datastore:"organiztionName"`
	Address       string         `json:"organizationAddress" datastore:"organizationAddress"`
	ContactNumber string         `json:"ContactNumber" datastore:"ContactNumber"`
	UserID        *datastore.Key `json:"UserID" datastore:"UserID"`
	CountEmployee int            `json:"CountEmployee" datastore:"CountEmployee"`
	CountStudent  int            `json:"CountStudent" datastore:"CountStudent"`
	CreatedDate   time.Time      `json:"CreatedDate" datastore:"CreatedDate"`
	UpdatedDate   time.Time      `json:"UpdatedDate" datastore:"UpdatedDate"`
}

type OrgSaveReq struct {
	Name          string `json:"organiztionName" endpoints:"required"`
	Address       string `json:"organizationAddress" endpoints:"req"`
	ContactNumber string `json:"ContactNumber" endpoints:"required"`
}
type OrgSaveReq1 struct {
	xx OrgSaveReq `json:"organiztionNameee" endpoints:"required"`
}
type OrgGetReq struct {
	OrganizationKey *datastore.Key `json:"organizationKey"`
}

type OrgSaveRes struct {
	Success string `json:"success" endpoints:"required"`
	Error   string `json:"error"`
}
type OrgRes struct {
	Name          string `json:"organiztionName"`
	Address       string `json:"organizationAddress"`
	ContactNumber string `json:"ContactNumber"`
}
type OrganizationList struct {
	Org1 []*OrgRes `json:"org1"`
}

func (o *Organization) Update(c context.Context, req *OrgSaveReq) {
	o.ContactNumber = req.ContactNumber
	o.Name = req.Name
	o.Address = req.Address
}

func (o *Organization) Put(c context.Context) (err error) {
	key := datastore.NewIncompleteKey(c, "Organization", nil)
	key, err = datastore.Put(c, key, o)
	return err
}

func (o *Organization) ToOrganizationJSON(org *OrgRes) *OrgRes {
	if org == nil {
		org = &OrgRes{}
	}
	org.Address = o.Address
	org.ContactNumber = o.ContactNumber
	org.Name = o.Name
	log.Println("org")
	log.Println(org)

	return org
}
func (o *Organization) GetOrgWithKey(c context.Context, key *datastore.Key) (*Organization, error) {
	//org := Organization{}
	err := datastore.Get(c, key, o)
	//	res := OrgSaveReq{}
	//	res.Address = org.Address
	//	res.ContactNumber = org.ContactNumber
	//	res.Name = org.Name
	log.Println("reflect.ValueOf(o)")
	log.Println(reflect.ValueOf(o))
	//res := o.OrganizationJSON(nil)
	//log.Println(res)
	return o, err
}

func (o *Organization) GetAllOrganization(c context.Context) ([]*Organization, error) {
	//z1 := zz{}
	//_, _ = datastore.NewQuery("organization").Limit(10).GetAll(c, org)
	//org := []Organization{}
	org := make([]*Organization, 0, 10)
	keys, err := datastore.NewQuery("Organization").GetAll(c, &org)
	//	var org []Organization
	//	_, err := datastore.NewQuery("Organization").GetAll(c, &org)
	//	log.Println(":ddd")
	//	log.Println(err)
	//	log.Println(":ddd")
	//log.Println(org1)
	//	//res = z1.oo
	//	//	x * Organization
	//	//	x = *Organization(org)
	//log.Println(reflect.TypeOf(&org))
	//	log.Println(OrganizationList{org1})
	for i, v := range keys {
		org[i].key = v
	}
	return org, err
}

func (o *Organization) GetAllOrganizationCursor(c context.Context) ([]*Organization, error) {
	//z1 := zz{}
	//_, _ = datastore.NewQuery("organization").Limit(10).GetAll(c, org)
	//org := []Organization{}
	org := make([]*Organization, 0, 10)
	keys, err := datastore.NewQuery("Organization").GetAll(c, &org)
	//datastore.NewQuery(kind).Start(c).GetAll(c, dst)
	//	var org []Organization
	//	_, err := datastore.NewQuery("Organization").GetAll(c, &org)
	//	log.Println(":ddd")
	//	log.Println(err)
	//	log.Println(":ddd")
	//log.Println(org1)
	//	//res = z1.oo
	//	//	x * Organization
	//	//	x = *Organization(org)
	//log.Println(reflect.TypeOf(&org))
	//	log.Println(OrganizationList{org1})
	for i, v := range keys {
		org[i].key = v
	}
	return org, err
}

//func (o *Organization) GetAllOrganization(c context.Context) (*Organization, error) {
//	//z1 := zz{}
//	//_, _ = datastore.NewQuery("organization").Limit(10).GetAll(c, org)
//	//org := []Organization{}
//	org1 := make([]*Organization, 0, 10)
//	_, err := datastore.NewQuery("Organization").GetAll(c, &org1)
//	//	var org []Organization
//	//	_, err := datastore.NewQuery("Organization").GetAll(c, &org)
//	//	log.Println(":ddd")
//	//	log.Println(err)
//	//	log.Println(":ddd")
//	log.Println(org1)
//	//	//res = z1.oo
//	//	//	x * Organization
//	//	//	x = *Organization(org)
//	log.Println(reflect.TypeOf(&org))
//	//	log.Println(OrganizationList{org1})
//	return &OrganizationList{org1}, err
//}

//func Update1(c context.Context, req *OrgSaveReq) *Organization {
//	o := Organization{}
//	o.ContactNumber = req.ContactNumber
//	o.Name = req.Name
//	o.Address = req.Address
//	o.cYZ = req.CYZ
//	return &o
//}

//
//func (o *Organization) Put(c context.Context) (err error) {
//	key := o.key
//	if key == nil {
//		key = datastore.NewIncompleteKey(c, "Organization", nil)
//	}
//	key, err = datastore.Put(c, key, o)
//	if err == nil {
//		o.key = key
//	}
//	log.Println("o.key")
//	log.Println(o.key)
//	return err
//
//}
