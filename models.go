package main

type CountAPIResponse struct {
	Value int `json:"value"`
}

type StatusResponse struct {
	CachedHashes      int    `json:"cached_hashes"`
	ProcessedRequests int64  `json:"processed_requests"`
	Uptime            int64  `json:"uptime"`
	CodeRepository    string `json:"source_code_repository"`
}
