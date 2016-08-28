package shivsubhashendpoints

import (
	"errors"
	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
	//"github.com/nipun0212/shivsubhashmodel"
	//"shivsubhash/shivsubhashmodel"

	"golang.org/x/net/context"
	//	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/user"
	"log"
	//"net/http"
	//"reflect"
)

func getCurrentUser(c context.Context) (*user.User, error) {
	u, err := endpoints.CurrentUser(c, scopes, audiences, clientIds)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, errors.New("Unaithorized user: Please, sign in")
	}
	c.Done()
	return u, nil
}

func RegisterService() (*endpoints.RPCService, error) {
	api := &ShivSubhashSchoolAPI{}
	rpcService, err := endpoints.RegisterService(api,
		"registerorganization", "v1", "Rgisters the Organization", true)
	log.Println("nipun")
	log.Println(err)
	if err != nil {
		return nil, err
	}

	info := rpcService.MethodByName("CreateOrganization").Info()
	info.Path, info.HTTPMethod, info.Name = "createOrganization", "POST", "createOrganization"
	info.Scopes, info.ClientIds, info.Audiences = scopes, clientIds, audiences

	info = rpcService.MethodByName("UpdateOrganization").Info()
	info.Path, info.HTTPMethod, info.Name = "updateOrganization", "POST", "updateOrganization"
	info.Scopes, info.ClientIds, info.Audiences = scopes, clientIds, audiences

	info = rpcService.MethodByName("GetOrganizationWithKey").Info()
	info.Path, info.HTTPMethod, info.Name = "getOrganization", "POST", "getOrganization"
	info.Scopes, info.ClientIds, info.Audiences = scopes, clientIds, audiences

	info = rpcService.MethodByName("GetAllOrganization").Info()
	info.Path, info.HTTPMethod, info.Name = "getALlOrganization", "POST", "getAllorganization"
	info.Scopes, info.ClientIds, info.Audiences = scopes, clientIds, audiences

	api1 := &ShivSubhashSchool1API{}
	rpcService, err = endpoints.RegisterService(api1,
		"registerorganization1", "v2", "Rgisters the Organization", true)
	log.Println("nipun")
	log.Println(err)
	if err != nil {
		return nil, err
	}

	info = rpcService.MethodByName("SaveOrganization").Info()
	info.Path, info.HTTPMethod, info.Name = "saveOrganization1", "POST", "save.organization"
	info.Scopes, info.ClientIds, info.Audiences = scopes, clientIds, audiences
	return rpcService, nil
}
