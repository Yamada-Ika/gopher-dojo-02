package divget

import (
	"net/http"
	"path/filepath"
)

type config struct {
	url       string
	filePath  string
	divN      uint64
	fileSize  uint64
	canDivGet bool
}

func makeConfig(url string, divN uint64) (*config, error) {
	cf, err := getConfig(url, divN)
	if err != nil {
		return noConfig, err
	}
	if !cf.canDivDownload() {
		cf.setdivN(1)
	}
	return cf, nil
}

func getConfig(url string, divN uint64) (*config, error) {
	resp, err := http.Head(url)
	if err != nil {
		return nil, err
	}
	canDivGet := false
	if resp.Header["Accept-Ranges"] != nil {
		canDivGet = true
	} else {
		canDivGet = false
	}
	defer resp.Body.Close()

	return &config{
		url:       url,
		filePath:  filepath.Base(url),
		divN:      divN,
		fileSize:  (uint64)(resp.ContentLength),
		canDivGet: canDivGet,
	}, nil
}

func (cf *config) canDivDownload() bool {
	return cf.canDivGet
}

func (cf *config) setdivN(n uint64) {
	cf.divN = n
}
