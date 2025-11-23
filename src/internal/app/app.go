package app

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/fancyinnovations/fancyspaces/src/internal/auth"
	"github.com/fancyinnovations/fancyspaces/src/internal/spaces"
	sqliteSpacesDB "github.com/fancyinnovations/fancyspaces/src/internal/spaces/database/sqlite"
	spacesHandler "github.com/fancyinnovations/fancyspaces/src/internal/spaces/handler"
	"github.com/fancyinnovations/fancyspaces/src/internal/versions"
	sqliteVersionsDB "github.com/fancyinnovations/fancyspaces/src/internal/versions/database/sqlite"
	localVersionFileStorage "github.com/fancyinnovations/fancyspaces/src/internal/versions/filestorage/local"
	versionsHandler "github.com/fancyinnovations/fancyspaces/src/internal/versions/handler"
)

const apiPrefix = "/api/v1"

type Configuration struct {
	Mux    *http.ServeMux
	SQLite *sql.DB
}

func Start(cfg Configuration) {
	// Spaces
	spacesDB := sqliteSpacesDB.New(sqliteSpacesDB.Configuration{
		Conn: cfg.SQLite,
	})
	if err := spacesDB.Setup(); err != nil {
		panic(fmt.Errorf("failed to setup spaces database: %w", err))
	}
	spacesStore := spaces.New(spaces.Configuration{
		DB: spacesDB,
	})
	sh := spacesHandler.New(spacesHandler.Configuration{
		Store:       spacesStore,
		UserFromCtx: auth.UserFromContext,
	})
	sh.Register(apiPrefix, cfg.Mux)

	// Versions
	versionsDB := sqliteVersionsDB.New(sqliteVersionsDB.Configuration{
		Conn: cfg.SQLite,
	})
	if err := versionsDB.Setup(); err != nil {
		panic(fmt.Errorf("failed to setup versions database: %w", err))
	}
	versionFileStorage := localVersionFileStorage.New()
	versionsStore := versions.New(versions.Configuration{
		DB:          versionsDB,
		FileStorage: versionFileStorage,
	})
	vh := versionsHandler.New(versionsHandler.Configuration{
		Store:       versionsStore,
		Spaces:      spacesStore,
		UserFromCtx: auth.UserFromContext,
	})
	vh.Register(apiPrefix, cfg.Mux)
}
