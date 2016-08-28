package shivsubhashendpoints

import (
	"errors"
	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
	"github.com/nipun0212/shivsubhashmodel"
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
	ContactNumber string `json:"ContactNumber" endpoints:"required"`
}
type UpdateOrganizationReq struct {
	OrganizationKey *datastore.Key `json:"organizationKey"`
	Name            string         `json:"organiztionName"`
	Address         string         `json:"organizationAddress"`
	ContactNumber   string         `json:"ContactNumber"`
}

func (ss *ShivSubhashSchoolAPI) CreateOrganization(r *http.Request, req *CreateOrganizationReq) error {
	c := endpoints.NewContext(r)
	var orgKey *datastore.Key
	var userKey *datastore.Key
	user := shivsubhashmodel.User{}
	org := shivsubhashmodel.Organization{}
	u, err := getCurrentUser(c)
	if err != nil {
		return err
	}
	userKey = datastore.NewKey(c, "User", u.ID, 0, nil)
	userExist := datastore.Get(c, userKey, &user)
	if userExist != nil || user.RoleName != "owner" {
		org.CreatedDate = time.Now()
		orgKey = datastore.NewIncompleteKey(c, "Organization", nil)
	} else {
		return errors.New("Organization already attahced to user")
	}

	org.ContactNumber = req.ContactNumber
	org.Name = req.Name
	org.Address = req.Address
	org.UpdatedDate = time.Now()
	org.UserID = userKey
	user.UserId = userKey
	user.RoleName = "owner"
	err = datastore.RunInTransaction(c, func(ctx context.Context) error {
		key, err := datastore.Put(c, orgKey, &org)
		if err != nil {
			return err
		}
		user.OrganizationKey = key
		key, err = datastore.Put(c, userKey, &user)
		return err

	}, nil)
	return err
}

func (ss *ShivSubhashSchoolAPI) UpdateOrganization(r *http.Request, req *UpdateOrganizationReq) error {
	c := endpoints.NewContext(r)
	var orgKey *datastore.Key
	var userKey *datastore.Key
	user := shivsubhashmodel.User{}
	org := shivsubhashmodel.Organization{}
	u, err := getCurrentUser(c)
	if err != nil {
		return err
	}
	userKey = datastore.NewKey(c, "User", u.ID, 0, nil)
	err = datastore.Get(c, userKey, &user)
	if user.OrganizationKey.String() != req.OrganizationKey.String() || user.RoleName != "owner" || err != nil {
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

func (ss *ShivSubhashSchoolAPI) GetOrganizationWithKey(r *http.Request, req *shivsubhashmodel.OrgGetReq) (*shivsubhashmodel.OrgRes, error) {
	c := endpoints.NewContext(r)
	_, err := getCurrentUser(c)
	if err != nil {
		return nil, err
	}
	o := shivsubhashmodel.Organization{}
	o1, err := o.GetOrgWithKey(c, req.OrganizationKey)
	//res := shivsubhashmodel.OrgRes{}
	res := o1.ToOrganizationJSON(nil)
	return res, err
}

func (ss *ShivSubhashSchoolAPI) GetAllOrganization(r *http.Request) (*shivsubhashmodel.OrganizationList, error) {
	c := endpoints.NewContext(r)
	_, err := getCurrentUser(c)
	if err != nil {
		return nil, err
	}
	o := shivsubhashmodel.Organization{}
	x, err := o.GetAllOrganization(c)
	//y := make([]*shivsubhashmodel.OrgRes, 0, 10)
	//resp.Items = make([]*ScoreRespMsg, len(scores))
	//	for i, score := range scores {
	//		resp.Items[i] = score.toMessage(nil)
	//	}
	//resp := shivsubhashmodel.OrganizationList{}
	//z := shivsubhashmodel.Organization{}
	resp := shivsubhashmodel.OrganizationList{}
	resp.Org1 = make([]*shivsubhashmodel.OrgRes, len(x))

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

	//	//x := shivsubhashmodel.OrganizationList{}
	//	//	x, _ := o.GetAllOrganization(c)
	return &resp, nil
}

//	if err != nil {
//		return err
//	}
//	req.ContactNumber = "918971819883"
//	o1 := shivsubhashmodel.Update1(c, req)
//	log.Println("o1")
//	log.Println(o1)
//	err = o1.Put(c)
//	if err != nil {
//		return err
//	}
//o.Update(req)
//log.Println("o.key")
//log.Println(o.key)
//log.Println("o.get(c, o.key)")
//	o2, _ := o.get(c, o.key)
//	log.Println(o2)
//	log.Println("o2.key")
//	log.Println(o2.key)
//	//	z := []Organization{}
//	//	z1 := &z
//	z1 := zz{}
//	log.Println("reflect.TypeOf(z1)")
//	log.Println(reflect.TypeOf(z1.oo))
//	log.Println(datastore.NewQuery("Organization").Limit(22).GetAll(c, z1.oo))
//	//log.Println((*z1)[2].key)
//	u1 := User{}
//	//	log.Println("o.get(c, o.key)")
//	//	o = o.get(c, key)
//	//	log.Println(o.)
//	log.Println(o.key)
//	u1.OrganizationKey = o.key
//	u1.userId = datastore.NewKey(c, "User", u.ID, 0, nil)
//	_, err = datastore.Put(c, u1.userId, &u1)
