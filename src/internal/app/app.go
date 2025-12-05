package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/fancyinnovations/fancyspaces/src/internal/analytics"
	analyticsCache "github.com/fancyinnovations/fancyspaces/src/internal/analytics/cache"
	analyticsDatabase "github.com/fancyinnovations/fancyspaces/src/internal/analytics/database/clickhouse"
	"github.com/fancyinnovations/fancyspaces/src/internal/auth"
	"github.com/fancyinnovations/fancyspaces/src/internal/spaces"
	fakeSpacesDB "github.com/fancyinnovations/fancyspaces/src/internal/spaces/database/fake"
	spacesHandler "github.com/fancyinnovations/fancyspaces/src/internal/spaces/handler"
	"github.com/fancyinnovations/fancyspaces/src/internal/versions"
	mongoVersionsDB "github.com/fancyinnovations/fancyspaces/src/internal/versions/database/mongo"
	minioVersionFileStorage "github.com/fancyinnovations/fancyspaces/src/internal/versions/filestorage/minio"
	versionsHandler "github.com/fancyinnovations/fancyspaces/src/internal/versions/handler"
	"github.com/minio/minio-go/v7"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const apiPrefix = "/api/v1"

type Configuration struct {
	Mux        *http.ServeMux
	Mongo      *mongo.Database
	ClickHouse driver.Conn
	MinIO      *minio.Client
}

func Start(cfg Configuration) {
	// Analytics
	aDB := analyticsDatabase.NewDB(&analyticsDatabase.Configuration{
		CH: cfg.ClickHouse,
	})
	if err := aDB.Setup(context.Background()); err != nil {
		panic(fmt.Errorf("could not setup analytics database: %w", err))
	}
	ac := analyticsCache.NewCache()
	as := analytics.New(analytics.Configuration{
		DB:    aDB,
		Cache: ac,
	})

	// Spaces
	spacesStore := spaces.New(spaces.Configuration{
		DB: seedSpacesDB(),
	})
	sh := spacesHandler.New(spacesHandler.Configuration{
		Store:       spacesStore,
		UserFromCtx: auth.UserFromContext,
	})
	sh.Register(apiPrefix, cfg.Mux)

	// Versions
	versionsDB := mongoVersionsDB.NewDB(&mongoVersionsDB.Configuration{
		Mongo: cfg.Mongo,
	})
	versionFileStorage := minioVersionFileStorage.NewStorage(cfg.MinIO)
	if err := versionFileStorage.Setup(context.Background()); err != nil {
		panic(fmt.Errorf("could not setup version file storage: %w", err))
	}
	versionsStore := versions.New(versions.Configuration{
		DB:          versionsDB,
		FileStorage: versionFileStorage,
		Analytics:   as,
	})
	vh := versionsHandler.New(versionsHandler.Configuration{
		Store:       versionsStore,
		Spaces:      spacesStore,
		UserFromCtx: auth.UserFromContext,
	})
	vh.Register(apiPrefix, cfg.Mux)
}

func seedSpacesDB() *fakeSpacesDB.DB {
	db := fakeSpacesDB.New()

	fancynpcsSpace := &spaces.Space{
		ID:          "fn",
		Slug:        "fancynpcs",
		Title:       "FancyNpcs",
		Description: "Simple, lightweight and feature-rich NPC plugin for Paper and Folia servers using packets.",
		Categories:  []spaces.Category{spaces.CategoryMinecraftPlugin},
		IconURL:     "",
		Status:      spaces.StatusApproved,
		CreatedAt:   time.Date(2025, 12, 3, 20, 0, 0, 0, time.UTC),
		Members: []spaces.Member{
			{
				UserID: "admin-1",
				Role:   spaces.RoleOwner,
			},
		},
	}
	if err := db.Create(fancynpcsSpace); err != nil {
		panic(fmt.Errorf("could not seed spaces db: %w", err))
	}

	return db
}
