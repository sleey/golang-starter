package main

import (
	flag "github.com/spf13/pflag"
)

type Flag struct {
	SkipMigration bool
	MigrationUp   bool
	MigrationDown bool
	DryRun        bool
}

var f Flag

func init() {
	flag.BoolVar(&f.SkipMigration, "skip-migration", false, "Skip running migrations")
	flag.BoolVar(&f.MigrationUp, "migration-up", false, "Run migrations up")
	flag.BoolVar(&f.MigrationDown, "migration-down", false, "Run migrations down")
	flag.BoolP("dry-run", "d", false, "dry run will only initialize all dependencies but not running the API. You still need to add skip-migration flag to skip migration")

	flag.Parse()
}
