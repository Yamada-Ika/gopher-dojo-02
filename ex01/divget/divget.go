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
	wg.Done()
	return nil
}

func makeByteRangeArray(contentLength, divNum int) []byteRange {
	var array []byteRange
	delta := (int)(contentLength / 2)
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

var wg sync.WaitGroup

func Start() error {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("error")
		return errors.New("error: Invalid argument")
	}
	fileUrl := args[1]
	resp0, err := http.Head(fileUrl)
	if err != nil {
		return err
	}
	contentLength := (int)(resp0.ContentLength)
	defer resp0.Body.Close()

	data := makeByteRangeArray(contentLength, 2)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go makeRequest(&data[i], fileUrl, i)
	}
	wg.Wait()

	f1, err := os.OpenFile("out.gif", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer f1.Close()

	for i := 0; i < len(newBodies); i++ {
		f1.Write(newBodies[i])
	}
	return nil
}
