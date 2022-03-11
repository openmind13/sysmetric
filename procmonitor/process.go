package procmonitor

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
)

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

	var wg sync.WaitGroup
	for _, file := range files {
		wg.Add(1)

		go func() {
			defer wg.Done()
			if !file.IsDir() {
				return
			}

			pid, err := strconv.Atoi(file.Name())
			if err != nil {
				return
			}

			f, err := os.Open(file.Name() + "/stat")
			if err != nil {
				return
			}
			defer f.Close()

			reader := bufio.NewReader(f)
			scanner := bufio.NewScanner(reader)
			scanner.Split(bufio.ScanWords)
			for scanner.Scan() {
				if strings.Contains(scanner.Text(), processName) {
					fmt.Println(pid)
					return
				}
			}
		}()
	}
	wg.Wait()

	return pids, nil
}
