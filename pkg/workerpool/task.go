package workerpool

import (
	"bufio"
	"github.com/goblinus/filestat/pkg/util"
	"os"
)

type Task struct {
	Err      error
	Filename string
	Result   map[rune]int
}

func NewTask(filename string) Task {
	return Task{Filename: filename}
}

func proccess(task *Task) {
	if file, err := os.Open(task.Filename); err == nil {
		defer file.Close()
		task.Result, task.Err = scanFile(file)
		//fmt.Println("file:", task.Filename, "task:", task.Result)
	} else {
		task.Err = err
	}
}

//scanFile реализует логику посимвольному сканированию файла и
//подготовку итогового словаря со статистикой
func scanFile(file *os.File) (map[rune]int, error) {
	result := make(map[rune]int)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		if r := util.ExtractRune(scanner.Text()); util.IsAscii(r) {
			if _, ok := result[r]; !ok {
				result[r] = 0
			}
			result[r]++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
