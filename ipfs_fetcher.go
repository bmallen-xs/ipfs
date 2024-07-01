package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Metadata struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

func fetchMetadata(cid string) (*Metadata, error) {
	url := fmt.Sprintf("https://ipfs.io/ipfs/%s", cid)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch metadata: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var metadata Metadata
	if err := json.Unmarshal(data, &metadata); err != nil {
		return nil, err
	}

	return &metadata, nil
}
