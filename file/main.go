package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"sync"
	"syscall"
	"time"
)

type T struct {
	Tdemo []byte
}

func main() {
	f, err := os.OpenFile("./filename", syscall.O_RDONLY, 0)
	fmt.Println("err:", err)
	defer f.Close()

	r := bufio.NewReader(f)
	linesPool := sync.Pool{
		New: func() interface{} {
			lines := make([]byte, 1024)
			// lines := new(T)
			// lines.Tdemo = make([]byte, 0)
			return &lines
		},
	}
	stringPool := sync.Pool{New: func() interface{} {
		lines := " stringPool "
		return &lines
	}}
	// slicePool := sync.Pool{New: func() interface{} {
	// 	lines := make([]string, 100)
	// 	return lines
	// }}

	var wg sync.WaitGroup
	for {
		buf1 := linesPool.Get().(*[]byte)
		n, err := r.Read(*buf1)
		fmt.Println("n:", n)
		buf := *buf1
		buf = buf[:n]
		fmt.Println("buf:\n", string(buf))
		if n == 0 {
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("err:", err)
				break
			}
		}
		nextUntillNewline, err := r.ReadBytes('\n') //read entire line

		if err != io.EOF {
			buf = append(buf, nextUntillNewline...)
		}

		wg.Add(1)
		go func() {
			//process each chunk concurrently
			//start -> log start time, end -> log end time
			ProcessChunk(buf, &linesPool, &stringPool, time.Now(), time.Now().Add(2*time.Second))
			wg.Done()
		}()
		wg.Wait()
	}
}

func ProcessChunk(chunk []byte, linesPool *sync.Pool, stringPool *sync.Pool, start time.Time, end time.Time) {
	//another wait group to process every chunk further
	var wg2 sync.WaitGroup
	// logs := stringPool.Get().(string)
	logs := string(chunk)
	linesPool.Put(&chunk) //put back the chunk in pool
	//split the string by "\n", so that we have slice of logs
	logsSlice := strings.Split(logs, "\n")
	stringPool.Put(&logs) //put back the string pool
	chunkSize := 100      //process the bunch of 100 logs in thread
	n := len(logsSlice)
	noOfThread := n / chunkSize
	if n%chunkSize != 0 { //check for overflow
		noOfThread++
	}
	length := len(logsSlice)
	//traverse the chunk
	// wg2.Add(length / chunkSize)
	for i := 0; i < length; i += chunkSize {
		wg2.Add(1)
		//process each chunk in saperate chunk
		go func(s int, e int) {
			defer wg2.Done()
			for i := s; i < e; i++ {
				text := logsSlice[i]
				if len(text) == 0 {
					continue
				}
				logParts := strings.SplitN(text, ",", 2)
				logCreationTimeString := logParts[0]
				logCreationTime, err := time.Parse("2006-01-02 15:04:05", logCreationTimeString)
				if err != nil {
					fmt.Printf("\n Could not able to parse the time :%s \nfor log : %v", logCreationTimeString, text)
					return
				}
				// check if log's timestamp is inbetween our desired period
				if logCreationTime.After(start) && logCreationTime.Before(end) {
					fmt.Println("text:", text)
				}
			}
			// wg2.Done()

		}(i*chunkSize, int(math.Min(float64((i+1)*chunkSize), float64(len(logsSlice)))))
		//passing the indexes for processing
	}
	wg2.Wait() //wait for a chunk to finish
	logsSlice = nil
}

// 分段读取
func case1() {
	f, err := os.OpenFile("filename", syscall.O_RDONLY, 0)
	fmt.Println(err)
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		buf := make([]byte, 4*1024)
		n, err := r.Read(buf)
		fmt.Println(n)
		buf = buf[:n]
		fmt.Println(string(buf))
		if n == 0 {
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println(err)
				break
			}
		}
	}
}
