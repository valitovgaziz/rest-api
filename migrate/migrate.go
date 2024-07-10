package migrate

import (
	"log/slog"

	"github.com/valitovgaziz/rest-api/initializers"
	"github.com/valitovgaziz/rest-api/models"
)

func Migrate() {
	initializers.DB.AutoMigrate(
		&models.DepthOrder{},
		&models.HistoryOrder{},
		&models.OrderBook{},
		&models.Client{},
	)

	slog.Info("Migration complited")

}
