package vasproduct

// CrawlResponse ...
type CrawlResponse struct {
	Header Header            `json:"header"`
	Data   CrawlResponseData `json:"data"`
}

// Header ...
type Header struct {
	ProcessTime float64  `json:"process_time"`
	Messages    []string `json:"messages"`
	Reason      string   `json:"reason"`
	ErrorCode   string   `json:"error_code"`
}

// CrawlResponseData ...
type CrawlResponseData struct {
	RequestID string `json:"request_id"`
	Status    bool   `json:"status"`
}

// CrawlRequestParams ...
type CrawlRequestParams struct {
	Keyword         string
	CategoryID      string //for category crawl
	MarketplaceID   []int
	MinPrice        int64
	MaxPrice        int64
	Limit           int64
	SortBy          int
	Prefix          string
	MarketplaceFlag map[int]int
	Proxy           string
}
