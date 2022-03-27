package divget

import (
	"os"

	"github.com/cheggaaa/pb/v3"
)

func makeProgressbar(data []byteRange, cf *config) *pb.ProgressBar {
	bar := pb.Start64(int64(cf.fileSize)).SetWriter(os.Stderr).Set(pb.Bytes, true)
	var total uint64
	for _, d := range data {
		total += d.to - d.from
	}
	bar.SetCurrent(int64(cf.fileSize - total))
	return bar
}
