package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	// TODO: добавить методы
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for i, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Printf("ошибка парсинга элемента %d: %v", i, err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("ошибка генерации информации для элемента %d: %v", i, err)
			continue
		}

		fmt.Println(info) // Просто выводим info с \n в конце
	}
}
