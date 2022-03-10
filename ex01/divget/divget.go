package divget

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type byteRange struct {
	from uint64
	to   uint64
}

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

// TODO キャンセルの処理
func Run(url string, parallelN uint64) error {

	ctx, finish := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	ctx, _ = signal.NotifyContext(ctx, syscall.SIGINT)
	ctx, _ = signal.NotifyContext(ctx, syscall.SIGQUIT)
	defer finish()

	// config（divget内で使う設定データ）
	// configを扱うところ
	cf, err := getConfig(url, parallelN)
	if err != nil {
		return err
	}
	if !cf.canDivDownload() {
		cf.setParallelN(1)
	}

	// 前準備
	data := setEachDataRange(cf)

	// 分割ダウンロードしているところ
	for i := uint64(0); i < cf.parallelN; i++ {
		i := i
		eg.Go(func() error {
			if err := divDownload(&data[i], cf.url, cf.filePath, i); err != nil {
				return err
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}

	// signalが送信されたら
	// キャンセルの処理はここに書く？
	select {
	case <-ctx.Done():
		return errors.New("Canceled by signal")
	default:

	}

	// データマージ
	f, err := os.OpenFile(cf.filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	for i := uint64(0); i < cf.parallelN; i++ {
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
