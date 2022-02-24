package divget

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"sync"

	"golang.org/x/sync/errgroup"
)

func genFileUrlToPath(url string) string {
	urls := strings.Split(url, "/")
	return urls[len(urls)-1]
}

type byteRange struct {
	from int
	to   int
}

var bodies [][]byte
var newBodies map[int][]byte = make(map[int][]byte)
var mutex sync.Mutex

func makeRequest(dataRange *byteRange, fileUrl string, index int) error {
	req, err := http.NewRequest("GET", fileUrl, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", dataRange.from, dataRange.to))

	dump, _ := httputil.DumpRequestOut(req, true)
	fmt.Printf("%s\n", dump)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	mutex.Lock()
	newBodies[index] = body
	mutex.Unlock()
	defer resp.Body.Close()
	return nil
}

func makeByteRangeArray(contentLength, divNum int) []byteRange {
	var array []byteRange

	if divNum == 1 {
		elem := byteRange{0, contentLength - 1}
		array = append(array, elem)
		return array
	}

	delta := (int)(contentLength / divNum)
	from := 0
	to := delta

	for i := 0; i < divNum; i++ {
		elem := byteRange{from, to}
		array = append(array, elem)
		from = to + 1
		if i == divNum-2 {
			to = contentLength - 1
		} else {
			to = from + delta
		}
	}
	return array
}

var eg errgroup.Group
var args []string

func validateArgs() error {
	args = os.Args
	if len(args) != 2 {
		return errors.New("error: Invalid argument")
	}
	return nil
}

type contentInfo struct {
	fileUrl        string
	contentLength  int
	canRangeAccess bool
}

func getContentInfo() (*contentInfo, error) {
	var config contentInfo

	config.fileUrl = args[1]
	resp, err := http.Head(config.fileUrl)
	if err != nil {
		return nil, err
	}
	config.contentLength = (int)(resp.ContentLength)
	if resp.Header["Accept-Ranges"] != nil {
		config.canRangeAccess = true
	} else {
		config.canRangeAccess = false
	}
	defer resp.Body.Close()

	return &config, nil
}

func calcParallelNum(config *contentInfo) int {
	if config.canRangeAccess {
		return 2
	}
	return 1
}

func Start() error {
	if err := validateArgs(); err != nil {
		return err
	}

	config, err := getContentInfo()
	if err != nil {
		return err
	}

	parallelNum := calcParallelNum(config)
	data := makeByteRangeArray(config.contentLength, parallelNum)
	for i := 0; i < parallelNum; i++ {
		i := i
		eg.Go(func() error {
			fmt.Println(i)
			err := makeRequest(&data[i], config.fileUrl, i)
			if err != nil {
				return err
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}

	f1, err := os.OpenFile(genFileUrlToPath(config.fileUrl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer f1.Close()

	for i := 0; i < len(newBodies); i++ {
		f1.Write(newBodies[i])
	}
	return nil
}
