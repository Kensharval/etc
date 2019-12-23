package etc

import (
	"fmt"
	"github.com/qamarian-dtp/err"
	errLib "github.com/qamarian-lib/err"
	"gopkg.in/qamarian-lib/str.v3"
	"testing"
)

func TestDecompress (t *testing.T) {
	str.PrintEtr ("Test started...", "std", "TestDecompress ()")

	someStr := []string {"a9dcddd2", "p[11", "s22", "@@2vdd4lm", "a3e3i33o3u3", "q2w2r2t2y2i2o2p2s2d2", "88", "467589ss3"}

	for _, aStr := range someStr {
		o, errX := Decompress (aStr)
		if errX != nil {
			errY := err.New ("Decompression failed.", nil, nil, errX)
			str.PrintEtr (errLib.Fup (errY), "std", "TestDecompress ()")
			continue
		}

		write := fmt.Sprintf ("Decompression success!\nCompressed:%s\nDecompresssed:%s", aStr, o)
		str.PrintEtr (write, "std", "TestDecompress ()")
	}

	str.PrintEtr ("Test failed!", "std", "TestDecompress ()")
}
