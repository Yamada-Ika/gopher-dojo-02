package divget

import (
	"context"
	"errors"

	"golang.org/x/sync/errgroup"
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

var eg errgroup.Group

func Run(ctx context.Context, url string, divN uint64) error {
	// config（divget内で使う設定データ）
	cf, err := makeConfig(url, divN)
	if err != nil {
		return err
	}

	// 前準備
	data := loadCache(cf)

	eg, ctx := errgroup.WithContext(ctx)
	for i := uint64(0); i < cf.divN; i++ {
		i := i
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return errors.New("Canceled by signal")
			default:
				if err := divDownload(&data[i], cf.url, cf.filePath, i); err != nil {
					return err
				}
				return nil
			}
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}

	// データマージ
	return mergeData(cf)
}
