package useragent

import (
	_ "embed"
)

//go:embed data/final.txt
var userAgentsFile string
