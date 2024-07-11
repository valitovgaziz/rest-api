package migrate

import (
	"log/slog"

	"github.com/valitovgaziz/rest-api/initializers"
	"github.com/valitovgaziz/rest-api/models"
)

func Migrate() {
	initializers.DB.AutoMigrate(
		&models.HistoryOrder{},
		&models.OrderBook{},
		&models.DepthOrderAsks{},
		&models.DepthOrderBids{},
		&models.Client{},
	)

	slog.Info("Migration complited")
}
