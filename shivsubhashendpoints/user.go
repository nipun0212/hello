package shivsubhashendpoints

import (
	//"errors"
	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
	"github.com/nipun0212/shivsubhashmodel"
	//"shivsubhash/shivsubhashmodel"

	//"golang.org/x/net/context"
	//	"google.golang.org/appengine/datastore"
	//"google.golang.org/appengine/user"
	//"log"
	"net/http"
	//"reflect"
)

type ShivSubhashSchool1API struct {
}

//type zz1 struct {1
//	oooo *zz
//}
// type OrgSaveReq struct {
// 	Name          string `json:"organiztionName" endpoints:"required"`
// 	Address       string `json:"organizationAddress" endpoints:"req"`
// 	ContactNumber string `json:"contactNumber" endpoints:"required"`
// 	CYZ           string `json:"cyz"`
// }

func (ss *ShivSubhashSchool1API) SaveOrganization(r *http.Request, req *shivsubhashmodel.OrgSaveReq) error {
	c := endpoints.NewContext(r)
	_, err := getCurrentUser(c)
	if err != nil {
		return err
	}
	o := shivsubhashmodel.Organization{}
	o.Update(c, req)
	err = o.Put(c)
	return err
}
