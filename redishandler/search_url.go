package redishandler

// Returns true if url was created before, ver = "longurl" or "shorturl"
func SearchURL(url, ver string) bool {
	// s can be "[val [url val]]"" or "[val []]""
	s := RDB.Do(Ctx, "zscan", ver, "0", "match", url).String()
	s = s[len(s)-3:]
	if s == "[]]" {
		return false
	}
	return true
}
