package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/goblinus/filestat/pkg/util"
	"github.com/goblinus/filestat/pkg/workerpool"
)

//внешний цикл: читаем файлы директории
//обработка ведется порациями (см. batchCapacity): обработка в воркер-пуле
//После получения результатов в тасках - отрправляем в итоговую статистику, см. как работать с освободившимся тасками

type Result struct {
	mutex sync.Mutex
	Data  map[rune]int
}

func main() {
	//1. Готовим переменные для работы и открываем директорию с файлами
	batchCapacity := 2
	workDir := "test"
	workPath, err := util.Pwd()
	if err != nil {
		log.Fatal(err)
	}

	files, err := os.ReadDir(fmt.Sprintf("%s/%s", workPath, workDir))
	if err != nil {
		log.Fatal(err)
	}

	var counter int
	var batch []*workerpool.Task
	var wg sync.WaitGroup

	//2. Основной цикл обработки данных из файла
	result := make(map[rune]int)
	for _, file := range files {
		counter++
		if file.IsDir() {
			continue
		}

		pathToFile := fmt.Sprintf("%s/%s/%s", workPath, workDir, file.Name())
		task := workerpool.NewTask(pathToFile)
		batch = append(batch, &task)
		if counter >= batchCapacity {
			pool := workerpool.NewPool(batch, len(batch))
			wg.Add(1)
			go func() {
				defer wg.Done()
				pool.Run()
			}()
			wg.Wait()

			//Собираем результат обработки файлов
			for _, item := range batch {
				result = util.Reduce(result, item.Result)
			}

			counter = 0
			batch = make([]*workerpool.Task, 0)
		}
	}

	util.DrawHistogramm(result)
	//fmt.Println(result)
	//fmt.Println(len(result))
	//fmt.Println(len(make(map[rune]int)))
}
