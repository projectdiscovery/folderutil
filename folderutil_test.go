package folderutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetFiles(t *testing.T) {
	// get files from current folder
	files, err := GetFiles(".")
	require.Nilf(t, err, "couldn't retrieve the list of files: %s", err)

	// we check only if the number of files is bigger than zero
	require.Positive(t, len(files), "no files could be retrieved: %s", err)
}

func TestPathInfo(t *testing.T) {
	got, err := NewPathInfo("/a/b/c")
	assert.Nil(t, err)
	gotPaths, err := got.Paths()
	assert.Nil(t, err)
	assert.EqualValues(t, []string{"/", "/a", "/a/b", "/a/b/c"}, gotPaths)
	gotMeshPaths, err := got.MeshWith("test.txt")
	assert.Nil(t, err)
	assert.EqualValues(t, []string{"/test.txt", "/a/test.txt", "/a/b/test.txt", "/a/b/c/test.txt"}, gotMeshPaths)
}
