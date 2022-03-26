package divget

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

func ConfigDivN(divN uint64) bool {
	if divN < uint64(runtime.NumCPU()) {
		return false
	}
	msg := fmt.Sprintf("WARNING: The number of divisions specified is %d, but we recommend %d, which is the same as the number of CPU cores. Would you like to continue? [y/N]: ", divN, runtime.NumCPU())
	fmt.Fprintf(os.Stderr, msg)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		switch scanner.Text() {
		case "y":
			return false
		case "N":
			return true
		}
		fmt.Fprintf(os.Stderr, "[y/N]?: ")
	}
	return true
}
