package controllers

import (
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

func (c Score) Add(id string, score int, score_best int) revel.Result {
	models.CreateUser(c.Txn, id)
	models.RegisterScore(c.Txn, id, score, score_best)

	res := &AddResponse{
		Result: "ok",
	}
	return c.RenderJson(res)
}

func (c Score) List() revel.Result {
	return c.Render()
}
