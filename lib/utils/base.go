package utils

import "fmt"

const API_PREFIX = "waterway_"

func GenerateApiToken() string {
	return API_PREFIX + RandomHex(20)
}

const SLASH = "/"
const EMPTY_PATH = ""

// PathPrefix returns the path for an app,
// regarding the env variable `WATERWAY_MULTI_APP`
// if WATERWAY_MULTI_APP is "1", then call this function
// with "blog" will return "/blog", otherwise it returns
// empty path
func PathPrefix(app string) string {
	multiApp, _ := GetEnv("WATERWAY_MULTI_APP")

	if multiApp == "1" {
		return SLASH + app
	}

	return EMPTY_PATH
}

func FullPath(suffix string) string {
	pwd := GetEnvMust("WATERWAY_PWD")

	return fmt.Sprintf("%s/%s", pwd, suffix)
}
