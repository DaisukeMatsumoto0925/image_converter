package convert

import "path/filepath"

func absPath(path string) (string, error) {
	if !filepath.IsAbs(path) {
		absPath, err := filepath.Abs(path)
		if err != nil {
			return "", err
		}
		return absPath, nil
	}
	return path, nil
}

func removeFileExt(path string) string {
	fileNameLen := len(path)
	extLen := len(filepath.Ext(path))
	return path[:fileNameLen-extLen]
}
