package web

import (
	"time"

	jsontime "github.com/liamylian/jsontime/v2/v2"
)

// replace ecoding/json lib to encode/decode json.
var json = jsontime.ConfigWithCustomTimeFormat

func init() {
	var (
		dateFormat     = "2006-01-02"           // yyyy-MM-dd
		dateTimeFormat = "2006-01-02T15:04:05Z" // yyyy-MM-ddThh:mm:ssZ
	)

	// Set alias to use with tag: time_format
	jsontime.AddTimeFormatAlias("date", dateFormat)
	jsontime.AddTimeFormatAlias("date_time", dateTimeFormat)

	// Set default date format
	jsontime.SetDefaultTimeFormat(dateTimeFormat, time.Local)

	// Load time zone America/Sao_Paulo
	timeZoneSaoPaulo, _ := time.LoadLocation("America/Sao_Paulo")

	// Add alias to use eith tag: time_location
	jsontime.AddLocaleAlias("sao_paulo", timeZoneSaoPaulo)
}
