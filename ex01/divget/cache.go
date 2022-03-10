package divget

import (
	"fmt"
	"os"
)

// size 5
// from 0 to 10
// ok 0-4, ng 5-10
func setEachDataRange(cf *config) (data []byteRange) {
	data = makeByteRangeArray(cf.fileSize, cf.parallelN)

	// cacheファイルがあればdataを書き換える
	for i := uint64(0); i < cf.parallelN; i++ {
		fi, err := os.Stat(fmt.Sprintf("./.cache/%s_%d", cf.filePath, i))
		if err != nil {
			continue
		}
		data[i].from = uint64(fi.Size())
	}
	return data
}
