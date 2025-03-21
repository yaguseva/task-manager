package repository

import (
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"os"
	"taskManager/internal/config"
	"taskManager/internal/entity"
)

var Storage = make(map[uuid.UUID]entity.Task)

func LoadData() (err error) {
	file, err := os.Open(config.Config.FileName)
	if err != nil {
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	decoder := json.NewDecoder(file)
	var task entity.Task
	for {
		if err = decoder.Decode(&task); err != nil {
			if err.Error() == "EOF" {
				return nil
			}
			return err
		}
		Storage[task.ID] = task
	}
}
