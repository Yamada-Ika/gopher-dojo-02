package divget

import (
	"context"

	"github.com/cheggaaa/pb/v3"
	"golang.org/x/sync/errgroup"
)

func download(ctx context.Context, bar *pb.ProgressBar, data []byteRange, cf *config) error {
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
	return nil
}

func Run(ctx context.Context, url string, divN uint64) error {
	// new config
	cf, err := makeConfig(url, divN)
	if err != nil {
		return err
	}

	// load cache data
	data := loadCache(cf)

	// new progress bar
	bar := makeProgressbar(data, cf)

	// download
	if err := download(ctx, bar, data, cf); err != nil {
		return err
	}

	// bind data
	return bindData(cf)
}
