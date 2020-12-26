package util

import (
	"errors"
	"log"
	"os"
)

//ConfigDBSQLite will execute the next task
// * Create a directory to contain the SQLIte files
func ConfigDBSQLite(path string){
	//createDirPath will check if the path exist, if not this will be created.
	if err := createIfNotExistDirPath("db/data"); err != nil{
		//Something goes wrong
		log.Panic(err);
	}
}

//createIfNotExistDirPath try to create a path if this do not exist
func createIfNotExistDirPath(dirPath string) error{
	fileMode := os.FileMode(0755);
	//Try to create the dir and subdir from the path
	if err := os.MkdirAll(dirPath,fileMode); err != nil{
		if !errors.Is(err, os.ErrExist){
			return err
		}
	}
	//All was ok
	return nil;
}
