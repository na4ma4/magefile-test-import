//go:build mage

package main

import (
	//mage:import
	"github.com/na4ma4/magefile-test-import/megapack"
)

var Default = megapack.Test
