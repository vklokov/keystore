package handlers

import "github.com/vklokov/keystore/utils"

type AppHandler struct {
	Config string
	Logger *utils.LogWriter
}

func New() *AppHandler {
	h := AppHandler{}
	return &h
}
