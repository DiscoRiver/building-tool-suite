package tools

import (
	"github.com/discoriver/building-tool-suite/internal/db"
	"github.com/discoriver/building-tool-suite/internal/terminator"
	"github.com/discoriver/building-tool-suite/pkg/log"
)

type ParseSomeDBQueryCommandFlag struct {
	ID string
}

func (cmd *ParseSomeDBQueryCommandFlag) ParseSomeDBQuery() {
	rows, err := db.QSomeDBQuery(cmd.ID)
	if err != nil {
		log.Error("ParseSomeDBQuery:: Couldn't get results: %s", err)
	}

	log.Debug("ParseSomeDBQuery:: Parsing %d rows", len(*rows))

	printer := terminator.NewTableOutput()
	printer.Print(rows)
}
