package folderutil

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var separator = string(os.PathSeparator)

// GetFiles within a folder
func GetFiles(root string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		matches = append(matches, path)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

// PathInfo about a folder
type PathInfo struct {
	IsAbsolute         bool
	RootPath           string
	Parts              []string
	PartsWithSeparator []string
}

// NewPathInfo returns info about a given path
func NewPathInfo(path string) (PathInfo, error) {
	var pathInfo PathInfo
	path = filepath.Clean(path)
	pathInfo.RootPath = filepath.VolumeName(path)
	if filepath.IsAbs(path) {
		if isUnixOS() {
			if pathInfo.RootPath == "" {
				pathInfo.IsAbsolute = true
				pathInfo.RootPath = "/"
			}
		}
	}

	for _, part := range strings.Split(path, separator) {
		if part != "" {
			pathInfo.Parts = append(pathInfo.Parts, part)
		}
	}

	for i, pathItem := range pathInfo.Parts {
		if i == 0 && pathInfo.IsAbsolute {
			pathInfo.PartsWithSeparator = append(pathInfo.PartsWithSeparator, pathInfo.RootPath)
		} else if pathInfo.PartsWithSeparator[len(pathInfo.PartsWithSeparator)-1] != separator {
			pathInfo.PartsWithSeparator = append(pathInfo.PartsWithSeparator, separator)
		}
		pathInfo.PartsWithSeparator = append(pathInfo.PartsWithSeparator, pathItem)
	}
	return pathInfo, nil
}

// Returns all possible combination of the various levels of the path parts
func (pathInfo PathInfo) Paths() ([]string, error) {
	var combos []string
	for i := 0; i <= len(pathInfo.Parts); i++ {
		computedPath := filepath.Join(pathInfo.Parts[:i]...)
		if pathInfo.IsAbsolute && pathInfo.RootPath != "" {
			computedPath = pathInfo.RootPath + computedPath
		}
		combos = append(combos, filepath.Clean(computedPath))
	}

	return combos, nil
}

// MeshWith combine all values from Path with another provided path
func (pathInfo PathInfo) MeshWith(anotherPath string) ([]string, error) {
	allPaths, err := pathInfo.Paths()
	if err != nil {
		return nil, err
	}
	var combos []string
	for _, basePath := range allPaths {
		combinedPath := filepath.Join(basePath, anotherPath)
		combos = append(combos, filepath.Clean(combinedPath))
	}

	return combos, nil
}

func isUnixOS() bool {
	switch runtime.GOOS {
	case "android", "darwin", "freebsd", "ios", "linux", "netbsd", "openbsd", "solaris":
		return true
	default:
		return false
	}
}

func isWindowsOS() bool {
	return runtime.GOOS == "windows"
}
