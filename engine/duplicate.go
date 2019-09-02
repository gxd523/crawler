package engine

func IsDuplicateUrl(url string) bool { // TODO md5
	if isVisited := visitedUrlMap[url]; isVisited {
		return true
	} else {
		visitedUrlMap[url] = true
		return false
	}
}
