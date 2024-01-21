package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
)

// 一种通用的读取文件，并发处理每行内容的方案
func ReadFileLineAndFn(path string, concurrence int, fn func(string) error) (uint64, uint64, error) {
	if concurrence < 1 {
		concurrence = 1
	}
	c := make(chan string)
	go readFileLines(path, c)

	var wg sync.WaitGroup
	limitChan := make(chan bool, concurrence)
	count := uint64(0)
	errCount := uint64(0)
	for line := range c {
		limitChan <- true
		wg.Add(1)
		go func(l string) {
			defer func() {
				wg.Done()
				<-limitChan
			}()
			if err := fn(l); err != nil {
				fmt.Println("fn err", err)
				atomic.AddUint64(&errCount, 1)
				return
			}
			atomic.AddUint64(&count, 1)
		}(line)
	}
	wg.Wait()
	return count, errCount, nil
}

func readFileLines(path string, c chan string) error {
	defer close(c)
	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		c <- line
	}
	return nil
}

func yourFn(line string) error {
	fmt.Println(line)
	// words := strings.Split(line, "\t")
	// a := words[0]
	// b := words[1]
	// fmt.Println(a, b)
	return nil
}

// go run common_read_file.go data.csv 1
func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println("please input filepath and concurrence")
		return
	}
	filepath := args[1]
	concurrence, _ := strconv.Atoi(args[2])
	count, errCount, err := ReadFileLineAndFn(filepath, concurrence, yourFn)
	fmt.Printf("success_count=%v,err_count=%v,err=%v\n", count, errCount, err)
}
