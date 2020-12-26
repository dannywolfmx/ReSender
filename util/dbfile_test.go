package util

import (
	"os"
	"testing"
)

//TestCreateFile
// * Get a path where the file will be create
// * Create and check if the file is correct
// * Fail when the file already exist
// * Return bool ("ok" if the file was successfully created), and error if with details
func TestCreateDBFolder(t *testing.T){
	dirPath :="/tmp/dbPrueba";
	createIfNotExistDirPath(dirPath);

	if err := os.Chdir(dirPath); err != nil{
		t.Error(err)
	}
}
