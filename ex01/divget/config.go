package divget

import "net/http"

type config struct {
	url       string
	filePath  string
	parallelN uint64
	fileSize  uint64
	canDivGet bool
}

func getConfig(url string, parallelN uint64) (*config, error) {
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
		filePath:  genurlToPath(url),
		parallelN: parallelN,
		fileSize:  (uint64)(resp.ContentLength),
		canDivGet: canDivGet,
	}, nil
}

func (cf *config) canDivDownload() bool {
	return cf.canDivGet
}

func (cf *config) setParallelN(n uint64) {
	cf.parallelN = n
}
