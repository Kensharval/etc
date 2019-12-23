package etc

import (
	"github.com/qamarian-dtp/err"
	"strconv"
	"strings"
)

// Function Decompress () decompresses a string.
//	
// 	o, e := Decompress ("a4bbcd3")
//	fmt.Println (o) // Output: aaaabbcddd
//
// If an empty string is supplied, an empty string gets returned.
// A number can not be the first character of the string, or be zero.
// Numbers can not exceed: (1 << 31) - 1
func Decompress (i string) (o string, e error) {
	var (
		number string
	)

	// ..1.. {
	if i == "" {
		o = i
		return
	}
	// ..1.. }

	// ..1.. {
	e1 := string (i [0])
	if strings.Contains ("0123456789", e1) {
		e = err.New ("Invalid input. [The first character can not be a number.]", nil, nil)
		return
	}
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
			n, errX := strconv.Atoi (number)
			if errX != nil {
				e = err.New ("Broken dependency. Ref: 0", nil, nil, errX)
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
		n, errZ := strconv.Atoi (number)
		if errZ != nil {
			e = err.New ("Broken dependency. Ref: 0", nil, nil, errZ)
		}

		lastElem := string (o [len (o) - 1])
		for y := 1; y <= n - 1; y ++ {
			o += lastElem
		}
	}
	// ..1.. }

	return
}
