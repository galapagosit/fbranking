package controllers

import "github.com/revel/revel"
import "github.com/revel/revel/modules/db/app"

type App struct {
	*revel.Controller
	db.Transactional
}

func (c App) Index() revel.Result {
	return c.Render()
}
