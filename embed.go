package useragent

import (
	_ "embed"
)

//go:embed agents_cleaned.txt
var UserAgentsFile string
