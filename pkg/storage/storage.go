package storage

func Get(databasePath string, key []byte) ([]byte, error) {
	_db, err := GetDatabase(databasePath)
	if err != nil {
		return nil, err
	}
	defer _db.Close()
	return _db.Get(key)
}

func Put(databasePath string, key, value []byte) error {
	_db, err := GetDatabase(databasePath)
	if err != nil {
		return err
	}
	defer _db.Close()
	return _db.Put(key, value)
}

func GetKeys(databasePath string) ([][]byte, error) {
	_db, err := GetDatabase(databasePath)
	if err != nil {
		return nil, err
	}
	defer _db.Close()
	c := _db.Keys()
	var keys [][]byte
	for key := range c {
		keys = append(keys, key)
	}
	return keys, nil
}
