package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"migration/cmd/migration"
	"migration/config"
	"migration/db"
	"migration/internal/app"
	"migration/internal/repository"
)

func Start() {
	cfg := config.ReadConfig()
	// root command
	root := &cobra.Command{}

	// command allowed
	cmds := []*cobra.Command{
		{
			Use:   "db:migrate",
			Short: "database migration",
			Run: func(cmd *cobra.Command, args []string) {
				migration.RunMigration(cfg)
			},
		},
		{
			Use:   "db:seeding",
			Short: "database seeding",
			Run: func(cmd *cobra.Command, args []string) {
				err := migration.SeedingData(cfg)
				if err != nil {
					log.Fatal(err.Error())
				}
			},
		},
		{
			Use:   "api",
			Short: "run api server",
			Run: func(cmd *cobra.Command, args []string) {
				dbSource, err := db.NewDatabase(cfg.SrcDB)
				if err != nil {
					log.Fatal(err)
				}
				dbDest, err := db.NewDatabase(cfg.DesDB)
				if err != nil {
					log.Fatal(err)
				}
				sourceRepo := repository.NewSourceRepository(dbSource)
				destRepo := repository.NewDestinationRepository(dbDest)
				app.Run(destRepo, sourceRepo)
			},
		},
	}
	root.AddCommand(cmds...)
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
