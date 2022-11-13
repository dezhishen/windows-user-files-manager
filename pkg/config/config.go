package config

import (
	"os"
	"sync"

	"git.mills.io/prologic/bitcask"
	"github.com/dezhishen/windows-user-files-manager/pkg/storage"
)

const (
	// DefaultConfigPath is the default path to the configuration file
	DefaultConfigPath = "./data/config"
)

var _db *bitcask.Bitcask
var _sync sync.RWMutex

func init() {
	os.Mkdir("./data", os.ModePerm)
	if _db == nil {
		_sync.Lock()
		defer _sync.Unlock()
		if _db == nil {
			var err error
			_db, err = storage.GetDatabase(DefaultConfigPath)
			if err != nil {
				panic(err)
			}
		}
	}
}

func Unlock() {
	_sync.Lock()
	defer _sync.Unlock()
	if _db != nil {
		_db.Close()
		_db = nil
	}
}

func GetConfigValue(key string) (string, error) {
	_sync.RLock()
	defer _sync.RUnlock()
	val, err := storage.GetWithDb(_db, []byte(key))
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func PutConfigValue(key, value string) error {
	_sync.RLock()
	defer _sync.RUnlock()
	return storage.PutWithDb(_db, []byte(key), []byte(value))
}

func PutAllValues(values map[string]string) error {
	_sync.RLock()
	defer _sync.RUnlock()
	for key, value := range values {
		_db.Put([]byte(key), []byte(value))
	}
	return nil
}

func GetConfigAllValues() (map[string]string, error) {
	_sync.RLock()
	defer _sync.RUnlock()
	var result = make(map[string]string)
	c := _db.Keys()
	for key := range c {
		val, err := _db.Get(key)
		if err != nil {
			return nil, err
		}
		result[string(key)] = string(val)
	}
	return result, nil
}
