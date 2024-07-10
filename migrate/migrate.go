package migrate

import (
	"log/slog"

	"github.com/valitovgaziz/rest-api/initializers"
	"github.com/valitovgaziz/rest-api/models"
)

func Migrate() {
	initializers.DB.AutoMigrate(
		&models.OrderBook{},
		&models.DepthOrderAsks{},
		&models.DepthOrderBids{},
	)

	slog.Info("Migration complited")

}
