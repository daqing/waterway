package blog

import "github.com/daqing/waterway/lib/utils"

func BlogTitle() string {
	return utils.GetEnvMust("AW_BLOG_TITLE")
}

func BlogTagline() string {
	tagline := utils.GetEnv("AW_BLOG_TAGLINE")
	if tagline == utils.EMPTY_STRING {
		return utils.EMPTY_STRING
	}

	return tagline
}
