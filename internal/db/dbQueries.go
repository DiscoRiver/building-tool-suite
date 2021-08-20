package db

import (
	"github.com/discoriver/building-tool-suite/pkg/log"
)

type SomeDBQueryColumns struct {
	Name string `db:"name" header:"name"`
	Type string `db:"type" header:"type"`
}

type SomeOtherDBQueryColumns struct {
	Name             string `db:"name" header:"Name"`
	ProductCode      string `db:"product_code" header:"Product Code"`
	ProductName      string `db:"product_name" header:"Product Name"`
	BuyPrice         string `db:"buy_price" header:"Buy Price"`
	MarkupPercentage string `db:"markup_percentage" header:"Markup %"`
}

func QSomeDBQuery(ID string) (*[]SomeDBQueryColumns, error) {
	query := `select
				name
				,type
				from some_table
				where ID == $1`

	log.Debug("SQL:: Query: %s", query)

	db := connect()

	var SomeQueryResult []SomeDBQueryColumns
	if err := db.Select(&SomeQueryResult, query, ID); err != nil {
		return nil, err
	}
	return &SomeQueryResult, nil
}

func QSomeOtherDBQuery(UID string) (*[]SomeOtherDBQueryColumns, error) {
	query := `select
				name
				,product_code
				,product_name
				,buy_price
				,markup_percentage
				from some_other_table
				where UID == $1`

	log.Debug("SQL:: Query: %s", query)

	db := connect()

	var SomeQueryResult []SomeOtherDBQueryColumns
	if err := db.Select(&SomeQueryResult, query, UID); err != nil {
		return nil, err
	}
	return &SomeQueryResult, nil
}
