package etc

import (
	"github.com/qamarian-dtp/err"
	"math/big"
	"regexp"
	"strconv"
	"strings"
)

// Function Decompress () decompresses a string.
//	
// 	someStr, err := Decompress ("a4bbcd3")
//	fmt.Println (someStr) // Output: aaaabbcddd
//
// A number, in the string, can not be:
//
// 	zero;
// 	one;	
// 	exceed: (1 << 31) - 1;
// 	the first character of the string.
func Decompress (i string) (o string, e error) {
	var (
		number string
	)

	// ..1.. {
	if decompress_InputPattern.MatchString (i) == false {
		e = err.New ("Invalid input.", nil, nil)
		return
	}
	// ..1.. }

	// ..1.. {
	e1 := string (i [0])
	o += e1
	// ..1.. }

	// ..1.. {
	if len (i) == 1 {
		return
	}
	// ..1.. }

	// ..1.. {
	e2 := string (i [1])
	if ! strings.Contains ("0123456789", e2) {
		o += e2
	} else {
		number += e2
	}
	// ..1.. }

	// ..1.. {
	for x := 3; x <= len (i); x ++ {
		nextElem := string (i [x - 1])

		if strings.Contains ("0123456789", nextElem) {
			number += nextElem
			continue
		}

		if number != "" {
			num, okT := big.NewInt (0).SetString (number, 0)
			if okT == false {
				e = err.New ("Bug detected: possibly due to broken " +
					"dependency; ref: 0.", nil, nil)
				return
			}

			max := big.NewInt ((1 << 31) - 1)
			if num.Cmp (max) == 1 {
				e = err.New (`A number exceeds "(1 << 31) - 1"`, nil, nil)
				return
			}

			n, errX := strconv.Atoi (number)
			if errX != nil {
				e = err.New (`Bug detected: possibly due to broken " +
					"dependency; ref: 1.`, nil, nil, errX)
				return
			}

			lastElem := string (o [len (o) - 1])
			for y := 1; y <= n - 1; y ++ {
				o += lastElem
			}
		}

		number = ""

		o += nextElem
	}
	// ..1.. }

	// ..1.. {
	if number != "" {
		num, okT := big.NewInt (0).SetString (number, 0)
		if okT == false {
			e = err.New ("Bug detected: possibly due to broken dependency; " +
				"ref: 0.", nil, nil)
			return
		}

		max := big.NewInt ((1 << 31) - 1)
		if num.Cmp (max) == 1 {
			e = err.New (`A number exceeds "(1 << 31) - 1"`, nil, nil)
			return
		}

		n, errZ := strconv.Atoi (number)
		if errZ != nil {
			e = err.New ("Bug detected: possibly due to broken dependency; " +
				"ref: 2.", nil, nil, errZ)
		}

		lastElem := string (o [len (o) - 1])
		for y := 1; y <= n - 1; y ++ {
			o += lastElem
		}
	}
	// ..1.. }

	return
}
var (
	decompress_InputPattern *regexp.Regexp
)
func init () {
	if initReport != nil {
		return
	}

	var errX error
	decompress_InputPattern, errX = regexp.Compile (`^(\D([2-9]|[1-9]\d+)?)+$`)
	if errX != nil {
		initReport = err.New ("Buggy package: unable to compile regular " +
			"expression; ref: Decompress ()-1.", nil, nil, errX)
	}
}
