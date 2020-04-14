package vasproduct

import "time"

// CrawlFilter ...
type CrawlFilter struct {
	RequestID string `json:"request_id"`
	Size      int64  `json:"size"`
	Page      int64  `json:"page"`
}

// CrawlResult ...
type CrawlResult struct {
	Data          CrawlResultData `json:"data"`
	StatusCode    int             `json:"status"`
	NeedReattempt bool            //used by grpc
	Success       bool            //used by grpc
}

// CrawlResultData ...
type CrawlResultData struct {
	CreatedTime time.Time `json:"created_time"`
	Products    []Product `json:"products"`
	Total       int64     `json:"total"`
}

// Product ...
type Product struct {
	MarketplaceID       int           `json:"marketplace_id"`
	ProductID           string        `json:"product_id"`
	Index               int64         `json:"index"`
	Name                string        `json:"name"`
	NormalizedName      string        `json:"normalized_name"`
	SearchRelevance     float64       `json:"search_relevance"`
	Variant             bool          `json:"has_variant"`
	Description         string        `json:"description"`
	DisplayPrice        float64       `json:"display_price"`
	Price               float64       `json:"price"`
	PriceBeforeDiscount float64       `json:"price_before_discount"`
	CategoryL1ID        int64         `json:"category_l1_id"`
	CategoryL2ID        int64         `json:"category_l2_id"`
	CategoryL3ID        int64         `json:"category_l3_id"`
	CategoryL1IDStr     string        `json:"category_l1_id_str"`
	CategoryL2IDStr     string        `json:"category_l2_id_str"`
	CategoryL3IDStr     string        `json:"category_l3_id_str"`
	ProductCondition    string        `json:"product_condition"` //enumerated?
	ProductURL          string        `json:"product_url"`
	OfficialShop        bool          `json:"from_official_shop"`
	BrandName           string        `json:"brand_name"`
	Verified            bool          `json:"verified_by_marketplace"`
	PowerMerchant       int           `json:"power_merchant"`
	ViewCount           int64         `json:"view_count"` // -1 if doesn't exist
	LikeCount           int64         `json:"like_count"` // -1 if doesn't exist
	Review              Review        `json:"review"`
	DiscussionCount     int64         `json:"discussion_count"` // -1 if doesn't exist
	Sold                Sold          `json:"sold"`
	WishlistCount       int64         `json:"wishlist_count"` // -1 if doesn't exist
	Rating              Rating        `json:"rating"`
	Preorder            bool          `json:"is_preorder"`
	MainImageURL        string        `json:"main_image_url"`
	Stock               int64         `json:"stock"`
	Bulks               []Bulk        `json:"bulk"`
	ShopID              int64         `json:"shop_id"`
	ImageURL            []string      `json:"image_url"`
	VariantInfos        []VariantInfo `json:"variant_info"`
	CreatedTime         time.Time     `json:"created_time"`
	LastUpdated         time.Time     `json:"last_updated"`
	CrawlTime           time.Time     `json:"crawl_time"`
	Active              bool          `json:"active"`
	ShopInfo            ShopInfo      `json:"shop_info"`
	Verification        Verification  `json:"verification"` // information about verification dependent with input productID
}

//Verification parse information about verification dependent with input productID
type Verification struct {
	ID               int       `json:"id" db:"id"`
	ProductID        int64     `json:"product_id" db:"product_id"`
	CrawledProductID string    `json:"crawled_product_id" db:"crawled_product_id"`
	Status           int       `json:"status" db:"status"`
	CreatedTime      time.Time `json:"created_time" db:"create_time"`
	UpdateTime       time.Time `json:"update_time" db:"update_time"`
	CreateBy         int64     `json:"create_by" db:"create_by"`
	UpdateBy         int64     `json:"update_by" db:"update_by"`
}

// Review ...
type Review struct {
	Count         int64 `json:"count"`
	StarReview    int64 `json:"star_review"`
	WrittenReview int64 `json:"written_review"`
	ImageReview   int64 `json:"image_review"`
}

// Sold ...
type Sold struct {
	Sold                int64 `json:"sold"`
	SoldCountLast30Days int64 `json:"sold_count_last30days"`
}

// Rating ...
type Rating struct {
	Value       float64  `json:"value"`
	RatingCount [5]int64 `json:"rating_count"`
}

// Bulk ...
type Bulk struct {
	Price      float64 `json:"price"`
	MinimumBuy int64   `json:"minimum_buy"`
}

// VariantInfo (duplicated) ...
type VariantInfo struct {
	VariantName         []string `json:"variant_name"`
	VariantValue        []string `json:"variant_value"`
	Price               float64  `json:"price"`
	PriceBeforeDiscount float64  `json:"price_before_discount"`
	StockQuantity       int64    `json:"stock_quantity"`
}

// ShopInfo ...
type ShopInfo struct {
	ID            string       `json:"id"`
	Name          string       `json:"name"`
	Username      string       `json:"domain"`
	Rating        ShopRating   `json:"rating"`
	Follower      ShopFollower `json:"follower"`
	PowerMerchant bool         `json:"power_merchant"`
	Verified      bool         `json:"verified_by_marketplace"`
	Place         string       `json:"place"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
	LastActive    time.Time    `json:"last_active"`
	Logo          string       `json:"logo"`
	ProductCount  int64        `json:"product_count"`
	// TopProducts   []int64      `json:"top_products"` //product id?
}

// ShopRating ...
type ShopRating struct {
	Status bool    `json:"status"`
	Value  float64 `json:"value"`
}

// ShopFollower ...
type ShopFollower struct {
	Status bool  `json:"status"`
	Value  int64 `json:"value"`
}

type CrawledProductFilter struct {
	MarketplaceID []int
	OStype        int //0,1,2
	MinSoldCount  int
	CategoryIDs   []int //lv1,lv2,lv3
	Keyword       string
}
type ProductResponse struct {
	Total    int64     `json:"total,omitempty"`
	MaxPrice float64   `json:"max_price,omitempty"`
	MinPrice float64   `json:"min_price,omitempty"`
	AvgPrice float64   `json:"avg_price,omitempty"`
	Products []Product `json:"products,omitempty"`
}
