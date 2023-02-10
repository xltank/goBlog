package number

import (
	"fmt"
	"math"
	"strconv"
)

const byteUnit string = "BKMGTP"

func FormatByte(n int) string {
	const step = 1024

	if n < step {
		return strconv.Itoa(n) + "B"
	}

	c := n / step
	logN := 1
	for ; c > step; c /= step {
		logN++
	}

	return fmt.Sprintf("%.2v%s", float64(n)/math.Pow(step, float64(logN)), string(byteUnit[logN]))
}
