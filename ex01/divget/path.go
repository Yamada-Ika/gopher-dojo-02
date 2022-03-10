package divget

import "path/filepath"

func genurlToPath(url string) string {
	return filepath.Base(url)
}
