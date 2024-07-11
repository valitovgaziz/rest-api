package migrate

import (
	"log/slog"

	"github.com/valitovgaziz/rest-api/models"
	"github.com/valitovgaziz/rest-api/storage"
)

func Migrate() {
	storage.DB.AutoMigrate(
		&models.HistoryOrder{},
		&models.OrderBook{},
		&models.DepthOrderAsks{},
		&models.DepthOrderBids{},
		&models.Client{},
	)

	slog.Info("Migration complited")
}
