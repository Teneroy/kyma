package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func prepareMockDirectories(testDir string) (string, error) {
	dirPath := testDir + "/test-data"
	err := os.Mkdir(dirPath, 0700)
	if err != nil {
		return "", err
	}

	emitters := []string{"emitter1", "emitter2", "emitter3"}
	for i, emitterName := range emitters {
		err = prepareMockDirectory(dirPath, emitterName, int64(i*100))
		if err != nil {
			return "", err
		}
	}

	return dirPath, err
}

func prepareMockDirectory(dirPath string, dirName string, size int64) error {
	const fileName string = "test.txt"

	err := os.Mkdir(dirPath+"/"+dirName, 0700)
	if err != nil {
		return err
	}

	var newFile *os.File = nil
	newFile, err = os.Create(dirPath + "/" + dirName + "/" + fileName)
	if err != nil {
		return err
	}

	err = os.Truncate(dirPath+"/"+dirName+"/"+fileName, size)
	if err != nil {
		newFile.Close()
		return err
	}
	newFile.Close()

	return err
}

func TestListDir(t *testing.T) {
	dirPath, errDirs := prepareMockDirectories(t.TempDir())
	assert.NoError(t, errDirs)

	expectedDirectories := []directory{
		{name: "emitter1", size: int64(0)},
		{name: "emitter2", size: int64(100)},
		{name: "emitter3", size: int64(200)},
	}

	directories, err := listDirs(dirPath)
	assert.NoError(t, err)

	isTrue := (len(directories) == len(expectedDirectories))
	for i, dir := range directories {
		if dir != expectedDirectories[i] {
			isTrue = false
			break
		}
	}

	require.True(t, isTrue)
}

func TestDirSize(t *testing.T) {
	dirPath, errDirs := prepareMockDirectories(t.TempDir())
	assert.NoError(t, errDirs)

	size, err := dirSize(dirPath)
	assert.NoError(t, err)

	require.Equal(t, int64(300), size)
}
