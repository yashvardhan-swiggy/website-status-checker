package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"websites-status-checker/internal/constant"
	"websites-status-checker/internal/interfaces"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	var post interfaces.PostRequestURLs
	err = json.Unmarshal(body, &post)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, website := range post.URLs {
		if _, ok := interfaces.WebsitesMap[website]; !ok {
			interfaces.WebsitesMap[website] = constant.StatusNotFetched
		}
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintln("Websites entered successfully")))
}
