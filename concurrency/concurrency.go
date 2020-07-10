package concurrency

// WebsiteChecker - Checks websites
type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

//CheckWebsites func to check
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultsChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultsChannel
		results[result.string] = result.bool
	}

	return results
}
