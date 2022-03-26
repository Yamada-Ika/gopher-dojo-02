package divget

import (
	"fmt"
	"io"
	"os"
)

func mergeData(cf *config) error {
	f, err := os.OpenFile(cf.filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	for i := uint64(0); i < cf.divN; i++ {
		cache, err := os.OpenFile(fmt.Sprintf("./.cache/%s_%d", cf.filePath, i), os.O_CREATE|os.O_RDONLY|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		defer cache.Close()
		data, err2 := io.ReadAll(cache)
		if err2 != nil {
			fmt.Println("hoge")
			return err2
		}
		f.Write(data)
		defer os.Remove(fmt.Sprintf("./.cache/%s_%d", cf.filePath, i))
	}
	return nil
}
