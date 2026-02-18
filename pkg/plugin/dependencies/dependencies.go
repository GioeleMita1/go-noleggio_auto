package dependencies

import (
	"noleggio_auto/pkg/repository"
	"noleggio_auto/pkg/service"
	"noleggio_auto/pkg/config/database"
	"os"
)

type AppDependencies struct {
	AutoService service.AutoService
}

func Init() *AppDependencies {
	
	db := database.MongoConnect(
		os.Getenv("MONGO_URI"), 
		os.Getenv("MONGO_DB_NAME"),
	)

	autoRepository := repository.NewAutoRepository(db)

	autoService := service.NewAutoService(autoRepository)

	return &AppDependencies{
		AutoService: autoService,
	}
}