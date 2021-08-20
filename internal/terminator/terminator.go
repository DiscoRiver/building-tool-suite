package terminator

import (
	"github.com/kataras/tablewriter"
	"github.com/landoop/tableprinter"
	"github.com/spf13/viper"
	"os"
)

func NewTableOutput() *tableprinter.Printer {
	printer := tableprinter.New(os.Stdout)

	printer.BorderTop, printer.BorderBottom, printer.BorderLeft, printer.BorderRight = true, true, true, true
	printer.CenterSeparator = "|"
	printer.ColumnSeparator = "|"
	printer.RowSeparator = "-"

	// Set table colours
	printer.HeaderFgColor, printer.HeaderBgColor = setDBTableHeaderColour()

	return printer
}

func setDBTableHeaderColour() (int, int) {
	backColour := viper.GetString("core.DBTableHeaderBackgroundColour")
	foreColour := viper.GetString("core.DBTableHeaderForegroundColour")

	foreColourMap := map[string]int{
		"red":     tablewriter.FgRedColor,
		"green":   tablewriter.FgGreenColor,
		"yellow":  tablewriter.FgYellowColor,
		"blue":    tablewriter.FgBlueColor,
		"cyan":    tablewriter.FgCyanColor,
		"magenta": tablewriter.FgMagentaColor,
		"white":   tablewriter.FgWhiteColor,
	}

	backColourMap := map[string]int{
		"red":     tablewriter.BgRedColor,
		"green":   tablewriter.BgGreenColor,
		"yellow":  tablewriter.BgYellowColor,
		"blue":    tablewriter.BgBlueColor,
		"cyan":    tablewriter.BgCyanColor,
		"magenta": tablewriter.BgMagentaColor,
		"white":   tablewriter.BgWhiteColor,
	}

	return foreColourMap[foreColour], backColourMap[backColour]
}
