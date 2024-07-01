package main

import (
	"log"
)

func main() {
	db, err := setupDB()
	if err != nil {
		log.Fatal(err)
	}

	cids, err := readCSV("data/ipfs_cids.csv")
	if err != nil {
		log.Fatal(err)
	}

	for _, cid := range cids {
		metadata, err := fetchMetadata(cid)
		if err != nil {
			log.Printf("Failed to fetch metadata for CID %s: %v", cid, err)
			continue
		}
		err = insertMetadata(db, cid, metadata)
		if err != nil {
			log.Printf("Failed to insert metadata for CID %s: %v", cid, err)
		}
	}

	startAPI(db)
}
