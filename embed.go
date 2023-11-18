package useragent

import (
	_ "embed"
)

//go:embed agents/final.txt
var userAgentsFile string
