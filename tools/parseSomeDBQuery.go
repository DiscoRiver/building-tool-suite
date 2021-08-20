package tools

import (
	"github.com/discoriver/building-tool-suite/internal/db"
	"github.com/discoriver/building-tool-suite/internal/terminator"
	"github.com/discoriver/building-tool-suite/pkg/log"
)

type ParseSomeDBQueryCommandFlag struct {
	ID string
}

func (cmd *ParseSomeDBQueryCommandFlag) ParseSomeDBQuery(ID string) {
	rows, err := db.QSomeDBQuery(ID)
	if err != nil {
		log.Error("ParseSomeDBQuery:: Couldn't get results: %s", err)
	}

	log.Debug("ParseSomeDBQuery:: Parsing %d rows", len(*rows))

	printer := terminator.NewTableOutput()
	printer.Print(rows)
}
