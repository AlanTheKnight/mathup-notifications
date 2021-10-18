package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

type LocalDB struct {
	FilePath  string
	Data      map[string]ExpoPushToken
	TimeStamp int64
}

func LoadLocalDB(filePath string) (*LocalDB, bool, error) {
	var localDB LocalDB = LocalDB{
		FilePath:  filePath,
		Data:      make(map[string]ExpoPushToken),
		TimeStamp: time.Now().Unix(),
	}
	var data []byte

	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		data, err = ioutil.ReadFile(filePath)
		if err != nil {
			return nil, false, err
		}

		err = json.Unmarshal(data, &localDB)
		if err != nil {
			return nil, false, err
		}
		return &localDB, true, nil
	}
	return &localDB, false, nil
}

func SaveLocalDB(db *LocalDB) error {
	file, err := json.MarshalIndent(db, "", " ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(db.FilePath, file, 0644)
}

func (db *LocalDB) Size() int {
	return len(db.Data)
}
