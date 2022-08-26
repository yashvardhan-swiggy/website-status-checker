package service

import (
	"log"
	"net/http"
	"sync"
	"websites-status-checker/internal/constant"
	"websites-status-checker/internal/interfaces"
)

func Traverse() {
	if len(interfaces.WebsitesMap) == 0 {
		return
	}
	var wg sync.WaitGroup
	var lock sync.Mutex
	wg.Add(len(interfaces.WebsitesMap))
	for website := range interfaces.WebsitesMap {
		website := website
		go func() {
			defer wg.Done()
			resp, err := http.Get("http://" + website)
			lock.Lock()
			if err != nil {
				interfaces.WebsitesMap[website] = constant.StatusDown
				log.Println(err)
			} else {
				log.Printf("Get %v : %v\n", "http://"+website, resp.StatusCode)
				if resp.StatusCode == http.StatusOK {
					interfaces.WebsitesMap[website] = constant.StatusUp
				} else {
					interfaces.WebsitesMap[website] = constant.StatusDown
				}
			}
			lock.Unlock()
		}()
	}
	wg.Wait()
}
