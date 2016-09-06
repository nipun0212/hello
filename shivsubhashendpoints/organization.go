package shivsubhashendpoints

import (
	"errors"
	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
	//"github.com/nipun0212/shivsubhashmodel"
	//"shivsubhash/shivsubhashmodel"
	"golang.org/x/net/context"
	//"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	//"google.golang.org/appengine/user"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"
)

type ShivSubhashSchoolAPI struct {
}

type CreateOrganizationReq struct {
	Name          string `json:"organiztionName" endpoints:"required"`
	Address       string `json:"organizationAddress" endpoints:"req"`
	ContactNumber string `json:"contactNumber" endpoints:"required"`
}
type UpdateOrganizationReq struct {
	OrganizationKey *datastore.Key `json:"organizationKey"`
	Name            string         `json:"organiztionName"`
	Address         string         `json:"organizationAddress"`
	ContactNumber   string         `json:"ContactNumber"`
}

func (ss *ShivSubhashSchoolAPI) CreateOrganization(r *http.Request, req *CreateOrganizationReq) error {
	c := endpoints.NewContext(r)

	u, err := getCurrentUser(c)
	if err != nil {
		return err
	}
	var orgKey *datastore.Key
	var userKey *datastore.Key

	user := User{}
	org := Organization{}

	userKey = datastore.NewKey(c, "User", u.ID, 0, nil)
	userExist := datastore.Get(c, userKey, &user)
	if userExist != nil || user.IsOwner == false {
		org.CreatedDate = time.Now()
		user.CreatedDate = time.Now()
		orgKey = datastore.NewIncompleteKey(c, "Organization", nil)
	} else {
		return errors.New("Organization already attahced to user")
	}

	org.ContactNumber = req.ContactNumber
	org.Name = req.Name
	org.Address = req.Address
	org.UpdatedDate = time.Now()
	org.UserID = userKey
	org.CountEmployee = 0
	org.CountStudent = 0
	user.UserId = userKey
	user.IsOwner = true

	err = datastore.RunInTransaction(c, func(ctx context.Context) error {
		key, err := datastore.Put(ctx, orgKey, &org)
		if err != nil {
			return err
		}
		user.OrganizationKey = key
		key, err = datastore.Put(ctx, userKey, &user)
		return err

	}, &datastore.TransactionOptions{XG: true})

	return err
}

func (ss *ShivSubhashSchoolAPI) UpdateOrganization(r *http.Request, req *UpdateOrganizationReq) error {
	c := endpoints.NewContext(r)
	var orgKey *datastore.Key
	var userKey *datastore.Key
	user := User{}
	org := Organization{}
	u, err := getCurrentUser(c)
	if err != nil {
		return err
	}
	userKey = datastore.NewKey(c, "User", u.ID, 0, nil)
	err = datastore.Get(c, userKey, &user)
	if user.OrganizationKey.String() != req.OrganizationKey.String() || user.IsOwner == false || err != nil {
		return errors.New("You are not allowed to update the Organization Data")
	}
	orgKey = user.OrganizationKey
	err = datastore.Get(c, orgKey, &org)
	if err != nil {
		return err
	}
	if strings.Compare(req.ContactNumber, "") != 0 {
		org.ContactNumber = req.ContactNumber
	}
	if strings.Compare(req.Name, "") != 0 {
		org.Name = req.Name
	}
	if strings.Compare(req.Address, "") != 0 {
		org.Address = req.Address
	}
	org.UpdatedDate = time.Now()
	_, err = datastore.Put(c, orgKey, &org)
	return err
}

func (ss *ShivSubhashSchoolAPI) GetOrganizationWithKey(r *http.Request, req *OrgGetReq) (*OrgRes, error) {
	c := endpoints.NewContext(r)
	_, err := getCurrentUser(c)
	if err != nil {
		return nil, err
	}
	o := Organization{}
	o1, err := o.GetOrgWithKey(c, req.OrganizationKey)
	//res := shivsubhashmodel.OrgRes{}
	res := o1.ToOrganizationJSON(nil)
	return res, err
}

func (ss *ShivSubhashSchoolAPI) GetAllOrganization(r *http.Request) (*OrganizationList, error) {
	c := endpoints.NewContext(r)
	_, err := getCurrentUser(c)
	if err != nil {
		return nil, err
	}
	o := Organization{}
	x, err := o.GetAllOrganization(c)
	//y := make([]*shivsubhashmodel.OrgRes, 0, 10)
	//resp.Items = make([]*ScoreRespMsg, len(scores))
	//	for i, score := range scores {
	//		resp.Items[i] = score.toMessage(nil)
	//	}
	//resp := shivsubhashmodel.OrganizationList{}
	//z := shivsubhashmodel.Organization{}
	resp := OrganizationList{}
	resp.Org1 = make([]*OrgRes, len(x))

	for i, v := range x {
		//resp.Org1[i] = v.toOrganizationJSON(nil)
		log.Println(reflect.ValueOf(v))
		//resp.Org1[i] = v
		log.Println("venus")
		log.Println(i)
		log.Println(v)
		//k := v.ToOrganizationJSON(nil)
		log.Println("k")
		log.Println(resp)
		log.Println("reflect.ValueOf(k)")
		log.Println(reflect.ValueOf(resp))
		resp.Org1[i] = v.ToOrganizationJSON(nil)
		//	_, _ = v.shivsubhashmodel.toOrganizationJSON(nil)
	}
	log.Println(x)
	log.Println(resp)
	return &resp, nil
}
