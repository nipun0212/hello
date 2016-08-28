package tictactoe

import (
	"errors"
	"fmt"
	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
	"golang.org/x/net/context"
	"google.golang.org/appengine/user"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const clientId = "522011478849-a1thjunen0u36rs4qoksoih24vdlabtn.apps.googleusercontent.com"

var (
	scopes    = []string{endpoints.EmailScope}
	clientIds = []string{clientId, endpoints.APIExplorerClientID}
	// in case we'll want to use TicTacToe API from an Android app
	audiences = []string{clientId}
)

type BoardMsg struct {
	State string `json:"state" endpoints:"required"`
}

type OrganizationReqSave struct {
	Name          string `json:"organiztionName"endpoints:"required"`
	Address       string `json:"organizationAddress"endpoints:"required"`
	ContactNumber string `json:"contactNumber"endpoints:"required"`
}
type ScoreRespMsg struct {
	Id      int64  `json:"id" endpoints:"required"`
	Outcome string `json:"outcomeeee"`
	Played  string `json:"played"`
}

type ScoresListReq struct {
	Limit int `json:"limit"`
}

type ScoresListResp struct {
	Items []*ScoreRespMsg `json:"items"`
}

// TicTacToe API service
type TicTacToeApi struct {
}

// BoardGetMove simulates a computer move in tictactoe.
// Exposed as API endpoint
func (ttt *TicTacToeApi) SaveOrganization(r *http.Request,
	req *OrganizationReqSave) error {
	c := endpoints.NewContext(r)
	u, err := getCurrentUser(c)
	if err != nil {
		return err
	}
	organization := &Organization{Name: req.Name, Address: req.Address, ContactNumber: req.ContactNumber}
	if err := organization.put(c); err != nil {
		return err
	}

	return nil
}

// ScoresList queries scores for the current user.
// Exposed as API endpoint
func (ttt *TicTacToeApi) ScoresList(r *http.Request,
	req *ScoresListReq, resp *ScoresListResp) error {

	c := endpoints.NewContext(r)
	u, err := getCurrentUser(c)
	if err != nil {
		return err
	}
	q := newUserScoreQuery(u)
	if req.Limit <= 0 {
		req.Limit = 10
	}

	scores, err := fetchScores(c, q, req.Limit)
	if err != nil {
		return err
	}
	resp.Items = make([]*ScoreRespMsg, len(scores))
	//for i, score := range scores {
	//resp.Items[i] = score.toMessage()
	//}
	return nil
}

// ScoresInsert inserts a new score for the current user.
func (ttt *TicTacToeApi) ScoresInsert(r *http.Request, req *ScoreReqMsg, r2 *ScoreReqMsg) error {

	c := endpoints.NewContext(r)
	u, err := getCurrentUser(c)
	if err != nil {
		return err
	}
	score := newScore(req.Outcome, u)
	if err := score.put(c); err != nil {
		return err
	}
	//resp.Played = "nipun"
	log.Println("score")
	log.Println(score.timestamp())
	score.toMessage(req)
	return nil
}

// getCurrentUser retrieves a user associated with the request.
// If there's no user (e.g. no auth info present in the request) returns
// an "unauthorized" error.
func getCurrentUser(c context.Context) (*user.User, error) {
	u, err := endpoints.CurrentUser(c, scopes, audiences, clientIds)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, errors.New("Unauthorized: Please, sign in.")
	}
	//c.Value("Current user: %#v", u)
	//	c.Debugf("Current user: %#v", u)
	c.Done()
	return u, nil
}

// RegisterService exposes TicTacToeApi methods as API endpoints.
//
// The registration/initialization during startup is not performed here but
// in app package. It is separated from this package (tictactoe) so that the
// service and its methods defined here can be used in another app,
// e.g. http://github.com/crhym3/go-endpoints.appspot.com.
func RegisterService() (*endpoints.RPCService, error) {
	api := &TicTacToeApi{}
	rpcService, err := endpoints.RegisterService(api,
		"tictactoe", "v1", "Tic Tac Toe API", true)
	if err != nil {
		return nil, err
	}

	info := rpcService.MethodByName("BoardGetMove").Info()
	info.Path, info.HTTPMethod, info.Name = "board", "POST", "board.getmove"
	info.Scopes, info.ClientIds, info.Audiences = scopes, clientIds, audiences

	info = rpcService.MethodByName("ScoresList").Info()
	info.Path, info.HTTPMethod, info.Name = "scores", "GET", "scores.list"
	info.Scopes, info.ClientIds, info.Audiences = scopes, clientIds, audiences

	info = rpcService.MethodByName("ScoresInsert").Info()
	info.Path, info.HTTPMethod, info.Name = "scores", "POST", "scores.insert"
	info.Scopes, info.ClientIds, info.Audiences = scopes, clientIds, audiences

	return rpcService, nil
}
