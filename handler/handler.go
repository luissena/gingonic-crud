package handler

import "github.com/luissena/golang-crud/config"

var (
	logger *config.Logger
)

func InitializeHandler() {
	logger = config.GetLogger("handler")

}
