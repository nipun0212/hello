package tictactoe

import (
	"time"

	 "google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/urlfetch"
	"google.golang.org/appengine/user"
	//"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
	"golang.org/x/net/context"
	"log"
	"github.com/Workiva/go-datastructures/merge"
)

const (
	TIME_LAYOUT = "Jan 2, 2006 15:04:05 AM"
	SCORE_KIND  = "Score"
)

// Score is an entity to store scores that have been inserted by users.
type Score struct {
	key     *datastore.Key
	Outcome string    `datastore:"outcome"`
	Played  time.Time `datastore:"played"`
	Player  string    `datastore:"player"`
}
type Organization struct {
	key           *datastore.Key
	Name          string `datastore:"organiztionName"`
	Address       string `datastore:"organizationAddress" endpoints:"req"`
	ContactNumber string `datastore:"contactNumber"`
}

// Turns the Score struct/entity into a ScoreRespMsg which is then used
// as an API response.
func (s *Score) toMessage(req *ScoreReqMsg) *ScoreReqMsg {
	msg := &ScoreRespMsg{}
	log.Println("s")
	log.Println(s.Outcome)
	msg.Id = s.key.IntID()
	msg.Outcome = s.Outcome
	msg.Played = s.timestamp()
	log.Println("msg")
	log.Println(req)
	c := appengine.NewContext(r)
    client := urlfetch.Client(c)
    client.Get("www.google.com")
	return req
}

// timestamp formats date/time of the score.
func (s *Score) timestamp() string {
	return s.Played.Format(TIME_LAYOUT)
	
}

// put stores the score in the Datastore.
func (s *Score) put(c context.Context) (err error) {
	key := s.key
	if key == nil {
		key = datastore.NewIncompleteKey(c, SCORE_KIND, nil)
	}
	key, err = datastore.Put(c, key, s)
	if err == nil {
		s.key = key
	}
	return
}

func (o *Organization) put(c context.Context) (err error) {
	key := o.key
	if key == nil {
		key = datastore.NewIncompleteKey(c, "Organization", nil)
	}
	key, err = datastore.Put(c, key, o)
	if err == nil {
		o.key = key
	}
	return err
}
func (o *Organization) get(c context.Context) *Organization {
	key := o.key
	ox := new(Organization)
	datastore.NewQuery(Organization).
	key1 := datastore.Get(c, key, ox)
	return key1
}

// newScore returns a new Score ready to be stored in the Datastore.
func newScore(outcome string, u *user.User) *Score {
	return &Score{Outcome: outcome, Played: time.Now(), Player: userId(u)}
}

// newUserScoreQuery returns a Query which can be used to list all previous
// games of a user.
func newUserScoreQuery(u *user.User) *datastore.Query {
	return datastore.NewQuery(SCORE_KIND).Filter("player =", userId(u))
}

// fetchScores runs Query q and returns Score entities fetched from the
// Datastore.
func fetchScores(c context.Context, q *datastore.Query, limit int) (
	[]*Score, error) {

	scores := make([]*Score, 0, limit)
	keys, err := q.Limit(limit).GetAll(c, &scores)
	if err != nil {
		return nil, err
	}
	for i, score := range scores {
		score.key = keys[i]
	}
	return scores, nil
}

// userId returns a string ID of the user u to be used as Player of Score.
func userId(u *user.User) string {
	return u.String()
}
