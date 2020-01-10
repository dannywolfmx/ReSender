//Original code: https://medium.com/@matryer/golang-advent-calendar-day-eleven-persisting-go-objects-to-disk-7caf1ee3d11d

package db

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"sync"
)

var lock sync.Mutex

//Marshal convertir objeto en bytes
var Marshal = func(v interface{}) (io.Reader, error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

var Unmarshal = func(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

func Save(path string, v interface{}) error {
	lock.Lock()
	defer lock.Unlock()
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	raw, err := Marshal(v)
	if err != nil {
		return err
	}
	_, err = io.Copy(file, raw)
	return nil
}

func Load(path string, v interface{}) error {
	lock.Lock()
	defer lock.Unlock()
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return Unmarshal(file, v)
}
