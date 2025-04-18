package utils

import (
	"regexp"
	"strings"
)

func Slugify(name string) string {
	slug := strings.ToLower(name)
	slug = strings.TrimSpace(slug)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "_", "-")
	reg := regexp.MustCompile(`[^a-z0-9\-]+`)
	slug = reg.ReplaceAllString(slug, "")
	slug = strings.Trim(slug, "-")
	return slug
}
