package divget

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/schollz/progressbar/v3"
)

func save(path string, dataRange *byteRange, resp *http.Response) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer f.Close()
	if err != nil {
		return err
	}
	_, err = f.Seek(0, 2)
	if err != nil {
		return err
	}
	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"downloading",
	)
	io.Copy(io.MultiWriter(f, bar), resp.Body)
	return nil
}

func divDownload(dataRange *byteRange, url, filePath string, index uint64) error {
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

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return err
	// }

	defer resp.Body.Close()
	return save(fmt.Sprintf("./.cache/%s_%d", filePath, index), dataRange, resp)
}
