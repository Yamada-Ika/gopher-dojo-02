package divget

import (
	"fmt"
	"os"
)

type byteRange struct {
	from uint64
	to   uint64
}

// size : 100, div : 4 -> 0-24, 25-49, 50-74, 75-100
func makeByteRangeArray(fileSize, divNum uint64) []byteRange {
	var array []byteRange

	if divNum == 1 {
		elem := byteRange{0, fileSize - 1}
		array = append(array, elem)
		return array
	}

	delta := (uint64)(fileSize / divNum)
	from := uint64(0)
	to := delta

	for i := uint64(0); i < divNum; i++ {
		elem := byteRange{from, to}
		array = append(array, elem)
		from = to + 1
		if i == divNum-2 {
			to = fileSize - 1
		} else {
			to = from + delta
		}
	}
	return array
}

// size 5
// from 0 to 10
// ok 0-4, ng 5-10
func loadCache(cf *config) (data []byteRange) {
	data = makeByteRangeArray(cf.fileSize, cf.divN)

	// cacheファイルがあればdataを書き換える
	for i := uint64(0); i < cf.divN; i++ {
		fi, err := os.Stat(fmt.Sprintf("./.cache/%s_%d", cf.filePath, i))
		if err != nil {
			continue
		}
		data[i].from = uint64(fi.Size())
	}
	return data
}
