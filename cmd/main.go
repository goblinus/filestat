package main

import (
	"fmt"
	"log"
	"os"

	"stristan/gopro/counter/pkg/utils"
	"stristan/gopro/counter/pkg/workerpool"
)

const (
	defaultSourceDir     = "test"
	emptyValue           = 0
	defaultBatchCapacity = "10"
)

//внешний цикл: читаем файлы директории
//обработка ведется порациями (см. batchCapacity): обработка в воркер-пуле
//После получения результатов в тасках - отрправляем в итоговую статистику, см. как работать с освободившимся тасками

func main() {
	//1. Готовим переменные для работы и открываем директорию с файлами
	batchCapacity := os.Getenv("batch_capacity")
	if len(batchCapacity) == emptyValue {
		batchCapacity = defaultBatchCapacity
	}
	workDir := os.Getenv("source_dir")
	if len(workDir) == emptyValue {
		workDir = defaultSourceDir
	}
	workPath, err := utils.Pwd()
	if err != nil {
		log.Fatal(err)
	}
	files, err := os.ReadDir(fmt.Sprintf("%s/%s", workPath, workDir))
	if err != nil {
		log.Fatal(err)
	}

	//2. Основной цикл обработки данных из файла
	var counter int
	var batch []*workerpool.Task
	for _, file := range files {
		counter++
		if file.IsDir() {
			continue
		}
		pathToFile := fmt.Sprintf("%s/%s/%s", workPath, workDir, file.Name())
		task := workerpool.NewTask(pathToFile)
		batch = append(batch, &task)
	}

	pool := workerpool.NewPool(batch, 5)

	pool.Run()
	fmt.Printf("%c <%d>\n", 'Ё', 'Ё')
}
