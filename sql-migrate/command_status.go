package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"

	migrate "github.com/rubenv/sql-migrate"
)

type StatusCommand struct{}

func (*StatusCommand) Help() string {
	helpText := `
Usage: sql-migrate status [options] ...

  Show migration status.

Options:

  -config=dbconfig.yml   Configuration file to use.
  -env="development"     Environment.

`
	return strings.TrimSpace(helpText)
}

func (*StatusCommand) Synopsis() string {
	return "Show migration status"
}

func (c *StatusCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("status", flag.ContinueOnError)
	cmdFlags.Usage = func() { log.Println(c.Help()) }
	ConfigFlags(cmdFlags)

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	env, err := GetEnvironment()
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not parse config: %s", err))
		return 1
	}

	db, dialect, err := GetConnection(env)
	if err != nil {
		log.Fatal(err.Error())
		return 1
	}
	defer db.Close()

	source := migrate.FileMigrationSource{
		Dir: env.Dir,
	}
	migrations, err := source.FindMigrations()
	if err != nil {
		log.Print(err.Error())
		return 1
	}

	records, err := migrate.GetMigrationRecords(db, dialect)
	if err != nil {
		log.Print(err.Error())
		return 1
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Migration", "Applied"})
	table.SetColWidth(60)

	rows := make(map[string]*statusRow)

	for _, m := range migrations {
		rows[m.Id] = &statusRow{
			Id:       m.Id,
			Migrated: false,
		}
	}

	for _, r := range records {
		if rows[r.Id] == nil {
			log.Println(fmt.Sprintf("Could not find migration file: %v", r.Id))
			continue
		}

		rows[r.Id].Migrated = true
		rows[r.Id].AppliedAt = r.AppliedAt
	}

	for _, m := range migrations {
		if rows[m.Id] != nil && rows[m.Id].Migrated {
			table.Append([]string{
				m.Id,
				rows[m.Id].AppliedAt.String(),
			})
		} else {
			table.Append([]string{
				m.Id,
				"no",
			})
		}
	}

	table.Render()

	return 0
}

type statusRow struct {
	Id        string
	Migrated  bool
	AppliedAt time.Time
}
