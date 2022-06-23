package example

import (
	"context"
	"github.com/fallboss/persistence"
)

func main() {
	// Config
	var prop = &persistence.PostgresProp{}
	pgClient := persistence.GetPostgresClient(context.Background(), prop)

	var repo1 persistence.DBQuery
	repo1 = persistence.GetPgRepo(pgClient)

	// Secondary
	repo1.ExecuteQuery("", nil)
}
