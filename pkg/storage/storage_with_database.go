package storage

import "git.mills.io/prologic/bitcask"

func GetDatabase(path string) (*bitcask.Bitcask, error) {
	return bitcask.Open(path)
}

func GetWithDb(db *bitcask.Bitcask, key []byte) ([]byte, error) {
	return db.Get(key)
}

func PutWithDb(db *bitcask.Bitcask, key, value []byte) error {
	return db.Put(key, value)
}

func GetKeysWithDb(db *bitcask.Bitcask) ([][]byte, error) {
	c := db.Keys()
	var keys [][]byte
	for key := range c {
		keys = append(keys, key)
	}
	return keys, nil
}
