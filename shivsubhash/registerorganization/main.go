//// This package performs initialization during startup.
////
//// It is separated from tictactoe package so that the latter can be used
//// in another app, e.g. http://github.com/crhym3/go-endpoints.appspot.com.
//
//package app
//
//import (
//	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
//	"github.com/nipun0212/go-tictactoe-master/tictactoe"
//)
//
//func init() {
//	if _, err := tictactoe.RegisterService(); err != nil {
//		panic(err.Error())
//	}
//	endpoints.HandleHTTP()
//}

package registerorganization

import (
	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
	"github.com/nipun0212/shivsubhashendpoints"
	//"golang.org/x/net/context"
	//"log"
	//"net/http"
)

func init() {
	//	if _, err := RegisterService(); err != nil {
	//		panic(err.Error())
	//	}
	if _, err := shivsubhashendpoints.RegisterService(); err != nil {
		panic(err.Error())
	}
	endpoints.HandleHTTP()
}
