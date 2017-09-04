package controllers

import (
	"github.com/revel/revel"
	"github.com/bamboV/media_watcher/app/models"
)

type Media struct {
	GormController
}

func (m Media) Index() revel.Result {
	result := []models.Media{}
	m.Txn.Find(&result)
	return m.RenderJSON(result)
}

func (m Media) Create() revel.Result {
	newMedia := models.Media {}
	m.Params.BindJSON(&newMedia)
	m.Txn.Save(&newMedia)
	return m.RenderJSON(newMedia)
}

func (m Media) Update(id uint) revel.Result {
	updateMedia := models.Media{}
	m.Txn.First(&updateMedia, models.Media{ID: id})
	m.Params.BindJSON(&updateMedia)
	m.Txn.Save(&updateMedia)
	return m.RenderJSON(updateMedia)
}