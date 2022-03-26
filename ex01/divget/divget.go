package divget

import (
	"context"
	"os"

	"github.com/cheggaaa/pb/v3"
	"golang.org/x/sync/errgroup"
	"golang.org/x/term"
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

	// progress
	// makeProgressbar
	width, _, err := term.GetSize(0)
	width /= 4
	if err != nil {
		return err
	}
	// var bars []*pb.ProgressBar
	// for _, d := range data {
	// 	// pbi := pb.New(int(d.to - d.from)).SetMaxWidth(width)
	// 	pbi := pb.Start64(int64(d.to-d.from)).SetMaxWidth(width).SetWriter(os.Stderr).Set(pb.Bytes, true)
	// 	bars = append(bars, pbi)
	// }
	// fmt.Println(len(bars))
	// os.Exit(1)
	bar := pb.Start64(int64(cf.fileSize)).SetMaxWidth(width).SetWriter(os.Stderr).Set(pb.Bytes, true)

	eg, ctx := errgroup.WithContext(ctx)
	for i := uint64(0); i < cf.divN; i++ {
		i := i
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				if err := divDownload(bar, &data[i], cf.url, cf.filePath, i); err != nil {
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
