package shivsubhashendpoints

import (
	//"errors"
	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
	//"github.com/nipun0212/shivsubhashmodel"
	//"shivsubhash/shivsubhashmodel"
	"time"
	//"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	//"google.golang.org/appengine/user"
	//"log"
	"net/http"
	//"reflect"
)

func prepareRoles() []Roles {
	roles := make([]Roles, 0, 10)
	role := Roles{}
	role.Name = "owner"
	role.ID = "1"
	role.CreatedDate = time.Now()
	roles = append(roles, role)
	role.Name = "principal"
	role.ID = "2"
	role.CreatedDate = time.Now()
	roles = append(roles, role)
	role.Name = "teacher"
	role.ID = "3"
	role.CreatedDate = time.Now()
	roles = append(roles, role)
	return roles

}
func (ss *ShivSubhashSchoolAPI) CreateRoles(r *http.Request) error {
	c := endpoints.NewContext(r)
	var err error
	roles := prepareRoles()
	for _, role := range roles {
		roleKey := datastore.NewKey(c, "Roles", role.ID, 0, nil)
		err = datastore.Get(c, roleKey, &role)
		if err == nil {
			role.UpdatedDate = time.Now()
		}
		_, err = datastore.Put(c, roleKey, &role)
	}

	return err
}
