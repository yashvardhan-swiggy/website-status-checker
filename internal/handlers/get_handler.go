package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"websites-status-checker/internal/interfaces"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Accept", "application/json")
	params := r.URL.Query()
	queryWebsites, ok := params["name"]
	if ok {
		outputMap := make(map[string]interface{})
		for _, item := range queryWebsites {
			_, ok = interfaces.WebsitesMap[item]
			if !ok {
				w.Write([]byte(fmt.Sprintf("%s not present in database\n", item)))
				log.Printf("%s not present in map", item)
			} else {
				outputMap[item] = interfaces.WebsitesMap[item]
			}
		}
		w.Write([]byte(fmt.Sprint("Status Map for queried websites : ")))
		err := json.NewEncoder(w).Encode(outputMap)
		if err != nil {
			log.Fatal(err)
			return
		}
	} else {
		w.Write([]byte(fmt.Sprint("Status Map for queried websites : ")))
		err := json.NewEncoder(w).Encode(interfaces.WebsitesMap)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}
