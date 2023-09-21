package cli

import (
	"app/internal/database"
	"flag"
	"fmt"
	"os"
)

func Main() {
	flag.Parse()

	switch flag.Arg(0) {
	case "migrate":
		database.Migrate()
	case "migrate:test":
		database.TestLastMigration()
	default:
		fmt.Println("Please specify a command")
		os.Exit(2)
	}

}
