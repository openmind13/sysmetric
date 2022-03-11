package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

var (
	names = []string{""}
)

func main() {
	for _, name := range names {
		pids, err := GetPidsByName(name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(name, pids)
	}

}

const (
	procDir = "/proc"
)

func GetPidsByName(processName string) ([]int, error) {
	pids := []int{}

	if err := os.Chdir(procDir); err != nil {
		return nil, err
	}

	files, err := ioutil.ReadDir(".")
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() {
			continue
		}
		pid, err := strconv.Atoi(file.Name())
		if err != nil {
			continue
		}
		f, err := os.Open(file.Name() + "/stat")
		if err != nil {
			log.Println(err)
			continue
		}

		r := bufio.NewReader(f)
		data, err := io.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(pid, string(data))
		// scanner := bufio.NewScanner(r)
		// scanner.Split(bufio.ScanWords)
		// for scanner.Scan() {
		// 	if scanner.Text() == processName {
		// 		pids = append(pids, )
		// 	}
		// 	if strings.Contains(scanner.Text(), processName) {
		// 		fmt.Println(pid)
		// 		// continue
		// 	}
		// }
	}

	return pids, nil
}
