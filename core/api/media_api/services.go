package media_api

import (
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/daqing/waterway/lib/repo"
	"github.com/daqing/waterway/lib/utils"
)

func SaveFile(
	userId repo.IdType,
	filename string,
	mime string,
	bytes int64,
) (*MediaFile, error) {
	return SaveFileExpiredAt(userId, filename, mime, bytes, repo.NeverExpires)
}

func SaveFileExpiredAt(
	userId repo.IdType,
	filename string,
	mime string,
	bytes int64,
	expiredAt time.Time,
) (*MediaFile, error) {

	return repo.Insert[MediaFile]([]repo.KVPair{
		repo.KV("user_id", userId),
		repo.KV("filename", filename),
		repo.KV("mime", mime),
		repo.KV("bytes", bytes),
		repo.KV("expired_at", expiredAt),
	})
}

// replace filename with part and origin extension
// replace("foo.pdf", "bar") -> "bar.pdf"
func replace(filename string, part string) string {
	return part + filepath.Ext(filename)
}

func hashDirPath(prefix string, path string) string {
	return assetFullPath(prefix, DirParts(path))
}

func DirParts(path string) string {
	if len(path) < 4 {
		return path
	}

	p1 := path[0:2]
	p2 := path[0:4]

	return "/" + p1 + "/" + p2
}

func assetFullPath(assetDir string, path string) string {
	return assetDir + path
}

func AssetHostPath(filename string) string {
	if filename == repo.EMPTY_STRING {
		return repo.EMPTY_STRING
	}

	subPath := "/public/assets" + filename

	host := utils.GetEnv("WATERWAY_ASSET_HOST")
	if host == utils.EMPTY_STRING {
		return subPath
	}

	return host + "/" + subPath
}

func DestFilePath(fileHeader *multipart.FileHeader) (destPath string, filePath string, err error) {
	file, err := fileHeader.Open()
	if err != nil {
		return
	}

	hash, err := utils.MD5SumFile(file)
	if err != nil {
		return
	}

	newFilename := replace(fileHeader.Filename, hash)

	parts := DirParts(newFilename)

	destDir := assetFullPath(AssetStorageDir(), parts)
	if err = os.MkdirAll(destDir, 0755); err != nil {
		return
	}

	destPath = destDir + "/" + newFilename
	filePath = parts + "/" + newFilename

	return destPath, filePath, nil
}

func AssetStorageDir() string {
	assetDir := utils.GetEnv("WATERWAY_STORAGE_DIR")
	if assetDir == utils.EMPTY_STRING {
		// env not set
		return utils.GetEnvMust("WATERWAY_PWD") + "/public/assets"
	}

	return assetDir
}
