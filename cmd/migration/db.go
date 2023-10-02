package migration

import (
	"flag"
	"fmt"
	"github.com/pressly/goose/v3"
	"log"
	"migration/config"
	driver "migration/db"
	"os"
)

var (
	flags         = flag.NewFlagSet("db:migrate", flag.ExitOnError)
	usageCommands = `
Commands:
    up                   Migrate the DB to the most recent version available
    down                 Roll back the version by 1
    reset                Roll back all migrations
`
	dirDestination = flags.String("dir destination", "db/destination", "directory with migration destination")
	dirSource      = flags.String("dir source", "db/source", "directory with migration source")
)

// RunMigration running auto migration
func RunMigration(cfg config.Config) {
	// assign var to flags
	flags.Usage = usage
	flags.Parse(os.Args[2:])

	args := flags.Args()
	if len(args) == 0 {
		flags.Usage()
		return
	}

	command := args[0]
	arguments := []string{}
	if len(args) > 3 {
		arguments = append(arguments, args[3:]...)
	}

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Failed to set dialect: %v", err)
	}

	// source database
	{
		dbSrc, err := driver.NewDatabase(cfg.SrcDB)
		if err != nil {
			log.Fatalf(err.Error())
		}

		//close connection
		defer func() {
			if err := dbSrc.Close(); err != nil {
				log.Fatalf("db migrate: failed to close DB: %v\n", err)
			}
		}()

		// running migration in destination folder
		if err := goose.Run(command, dbSrc, *dirSource, arguments...); err != nil {
			log.Fatalf("db migrate run: %v", err)
		}
	}

	// destination database
	{
		dbDes, err := driver.NewDatabase(cfg.DesDB)
		if err != nil {
			log.Fatalf(err.Error())
		}

		//close connection
		defer func() {
			if err := dbDes.Close(); err != nil {
				log.Fatalf("db migrate: failed to close DB: %v\n", err)
			}
		}()

		// running migration in source folder
		if err := goose.Run(command, dbDes, *dirDestination, arguments...); err != nil {
			log.Fatalf("db migrate run: %v", err)
		}
	}

}

// usage print list of command
func usage() {
	fmt.Println(usageCommands)
}
