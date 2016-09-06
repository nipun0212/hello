package shivsubhashendpoints

import (
	"errors"
	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"log"
	"net/http"
	//	"reflect"
)

type AddEmployeeReq struct {
	OrganizationKey *datastore.Key `json:"OrganizationKey"`
	EmpID           int64          `json:"EmpID"`
	EmailID         string         `json:"EmailID"`
	Name            string         `json:"Name"`
	Age             int            `json:"Age"`
	Gender          bool           `json:"Gender"`
}

type UpdateEmployeeRoleReq struct {
	EmployeeKey *datastore.Key `json:"EmployeeKey"`
	IsPrincipal bool           `json:"IsPrincipal"`
	IsTeacher   bool           `json:"IsTeacher"`
	IsClerk     bool           `json:"IsClerk"`
}

func (ss *ShivSubhashSchoolAPI) AddEmployee(r *http.Request, req *AddEmployeeReq) error {
	c := endpoints.NewContext(r)
	u, err := getCurrentUser(c)
	if err != nil {
		return err
	}
	var empKey *datastore.Key
	var userKey *datastore.Key
	var userExist bool
	user := User{}
	emp := Employee{}
	org := Organization{}
	userKey = datastore.NewKey(c, "User", u.ID, 0, nil)
	err = datastore.Get(c, userKey, &user)
	if err == nil {
		userExist = true
	} else {
		userExist = false
	}
	if userExist == false {
		return errors.New("You are not authorized to add Employee")
	}

	if !(user.IsOwner == true || user.IsPrincipal == true || user.IsClerk == true) {
		return errors.New("You are not authorized to add Employee")
	}

	if user.OrganizationKey.String() != req.OrganizationKey.String() {
		return errors.New("You are not authorized to add Employee")
	}

	//q := datastore.NewQuery("Employee").Filter("organizationKey =", req.OrganizationKey)
	//count, err := q.Count(c)
	//log.Println("count")
	//log.Println(count)
	err = datastore.Get(c, user.OrganizationKey, &org)
	if err != nil {
		return err
	}
	log.Println(emp.EmpID)
	log.Println("count")
	log.Println(org.CountEmployee)
	emp.Age = req.Age
	emp.Name = req.Name
	emp.EmailID = req.EmailID
	emp.OrganizationKey = user.OrganizationKey
	empKey = datastore.NewIncompleteKey(c, "Employee", nil)

	err = datastore.RunInTransaction(c, func(ctx context.Context) error {
		emp.EmpID = org.CountEmployee + 1
		org.CountEmployee = org.CountEmployee + 1
		empKey, err = datastore.Put(ctx, empKey, &emp)
		if err != nil {
			return err
		}
		_, err = datastore.Put(ctx, user.OrganizationKey, &org)
		return err

	}, &datastore.TransactionOptions{XG: true})

	return err
}
func (ss *ShivSubhashSchoolAPI) UpdateEmployeeRole(r *http.Request, req *UpdateEmployeeRoleReq) error {
	c := endpoints.NewContext(r)
	u, err := getCurrentUser(c)
	if err != nil {
		return err
	}
	var empKey *datastore.Key
	var userKey *datastore.Key
	var userExist bool
	user := User{}
	emp := Employee{}
	org := Organization{}
	empKey = req.EmployeeKey
	err = datastore.Get(c, empKey, &emp)
	if err != nil {
		return errors.New("You are a registered Employee")
	}
	userKey = datastore.NewKey(c, "User", u.ID, 0, nil)
	err = datastore.Get(c, userKey, &user)
	if err == nil {
		userExist = true
	} else {
		userExist = false
	}
	if userExist == false {
		return errors.New("You are not authorized to add Employee")
	}

	if !(user.IsOwner == true || user.IsPrincipal == true || user.IsClerk == true) {
		return errors.New("You are not authorized to add Employee")
	}

	if user.OrganizationKey.String() != emp.OrganizationKey.String() {
		return errors.New("You are not authorized to add Employee")
	}
	switch req.IsPrincipal {
	case true:
		emp.IsPrincipal = true
	}
	switch req.IsClerk {
	case true:
		emp.IsClerk = true
	}
	switch req.IsTeacher {
	case true:
		emp.IsTeacher = true
	}

	err = datastore.Get(c, user.OrganizationKey, &org)
	if err != nil {
		return err
	}
	empKey, err = datastore.Put(c, empKey, &emp)
	return err
}

//datastore.n
//	emp := Employee{}
//	emp1 := Employee{}
//	x := make([]int64, 4, 4)
//	x = append(x, 1, 2, 3, 4)
//	emp.EmpID = x
//
//	key := datastore.NewIncompleteKey(c, "Employee", nil)
//	key, err = datastore.Put(c, key, &emp)
//	x = append(x, 5, 6, 7, 8)
//	emp.EmpID = x
//	key, err = datastore.Put(c, key, &emp)
//
//	err = datastore.Get(c, key, &emp1)
//	log.Println("emp1")
//	log.Println(len(emp1.EmpID))
//	x2 := emp1.EmpID
//	for i, v := range x2 {
//		if v == 5 {
//			//a = a[:i+copy(a[i:], a[i+1:])]
//			x2 = x2[:i+copy(x2[i:], x2[i+1:])]
//			log.Println("x2")
//			log.Println(x2)
//		}
//
//	}
//	return err
//}
