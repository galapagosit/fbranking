package controllers

import (
	"crypto/sha1"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/galapagosit/fbranking/app/models"
	"github.com/revel/revel"
)
import "github.com/revel/revel/modules/db/app"

type Score struct {
	*revel.Controller
	db.Transactional
}

type AddResponse struct {
	Result string `json:"result"`
}

func isValidParams(id string, score int, score_best int, sec_key string) bool {

	var ary = []string{
		id,
		strconv.Itoa(score),
		strconv.Itoa(score_best)}
	joinedStr := strings.Join(ary, "db")

	h := sha1.New()
	h.Write([]byte(joinedStr))
	bs := h.Sum(nil)
	s := fmt.Sprintf("%x", bs)
	log.Printf("sha1 hash is:" + s)

	if sec_key == s {
		return true
	} else {
		return false
	}
}

func (c Score) Add(id string, score int, score_best int, sec_key string) revel.Result {
	if !isValidParams(id, score, score_best, sec_key) {
		res := &AddResponse{
			Result: "ng",
		}
		return c.RenderJson(res)
	}

	models.CreateUser(c.Txn, id)
	models.RegisterScore(c.Txn, id, score, score_best)

	res := &AddResponse{
		Result: "ok",
	}
	return c.RenderJson(res)
}

func (c Score) List(ids string) revel.Result {
	idList := strings.Split(ids, ",")
	userScores := models.GetScores(c.Txn, idList)

	return c.RenderJson(userScores)
}
