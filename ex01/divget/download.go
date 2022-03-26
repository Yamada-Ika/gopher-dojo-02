package divget

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/cheggaaa/pb/v3"
)

func save(bar *pb.ProgressBar, path string, dataRange *byteRange, resp *http.Response) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Seek(0, 2)
	if err != nil {
		return err
	}

	// display progress bar
	barReader := bar.NewProxyReader(resp.Body)
	io.Copy(f, barReader)
	bar.Finish()

	return nil
}

func divDownload(bar *pb.ProgressBar, dataRange *byteRange, url, filePath string, index uint64) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", dataRange.from, dataRange.to))

	// dump, _ := httputil.DumpRequestOut(req, true)
	// fmt.Printf("%s\n", dump)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return save(bar, fmt.Sprintf("./.cache/%s_%d", filePath, index), dataRange, resp)
}
