package initconst

import (
	"log"

	"github.com/fixifi/fixifi-go-backend/handlers"
)

func MustInit(mainHandler *handlers.MainHandler) {
	if err := mainHandler.InitializeCategories(); err != nil {
		log.Fatal("Failed to initialize categories:", err)
	}
}
