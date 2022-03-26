package divget

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

var eg errgroup.Group

func Run(ctx context.Context, url string, divN uint64) error {
	// config（divget内で使う設定データ）
	cf, err := makeConfig(url, divN)
	if err != nil {
		return err
	}

	// 前準備
	data := loadCache(cf)

	fmt.Println(data)

	eg, ctx := errgroup.WithContext(ctx)
	for i := uint64(0); i < cf.divN; i++ {
		i := i
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return ctx.Err()
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
