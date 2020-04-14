package productcrawl

import (
	"time"

	modelProduct "github.com/tokopedia/price-comparator/internal/model/vasproduct"
)

// ProductCrawl ...
type ProductCrawl struct {
	ID                 int64             `json:"id" db:"id"`
	RequestID          string            `json:"request_id" db:"request_id"`
	NormalizedName     string            `json:"normalized_name" db:"normalized_name"`
	ShopID             int64             `json:"shop_id" db:"shop_id"`
	ProductID          int64             `json:"product_id" db:"product_id"`
	ProductTitle       string            `json:"product_title" db:"product_title"`
	ProductURL         string            `json:"product_url" db:"product_url"`
	ImageURL           string            `json:"image_url" db:"image_url"`
	Price              float64           `json:"price" db:"price"`
	DiscountedPrice    float64           `json:"discounted_price" db:"discounted_price"`
	PriceMax           float64           `json:"price_max" db:"price_max"`
	PriceMin           float64           `json:"price_min" db:"price_min"`
	CategoryL1ID       int64             `json:"category_l1_id" db:"category_l1_id"`
	CategoryL2ID       int64             `json:"category_l2_id" db:"category_l2_id"`
	CategoryL3ID       int64             `json:"category_l3_id" db:"category_l3_id"`
	CategoryL1IDStr    string            `json:"category_l1_id_str" db:"category_l1_id_str"`
	CategoryL2IDStr    string            `json:"category_l2_id_str" db:"category_l2_id_str"`
	CategoryL3IDStr    string            `json:"category_l3_id_str" db:"category_l3_id_str"`
	ProcessStep        int               `json:"process_step" db:"process_step"`
	FilterTemplateID   int64             `json:"filter_template_id" db:"filter_template_id"`
	MarketplaceReports []byte            `json:"marketplace_report" db:"marketplace_report"`
	CreateBy           int64             `json:"create_by" db:"create_by"`
	CreateTime         time.Time         `json:"create_time" db:"create_time"`
	UpdateBy           int64             `json:"update_by" db:"update_by"`
	UpdateTime         time.Time         `json:"update_time" db:"update_time"`
	Status             int               `json:"status" db:"status"`
	FullCount          int64             `json:"full_count" db:"full_count"`
	BaselinePrice      float64           `json:"baseline_price" db:"baseline_price"`
	CalculatedPrices   CalculatedPrices  `json:"calculated_price" db:"calculated_price"`
	VerifiedKeyword    string            `json:"verified_keyword" db:"verified_keyword"`
	TopProductHistory  TopProductHistory `json:"top_product_history" db:"top_product_history"`
}

// CalculatedPrices ...
type CalculatedPrices struct {
	PriceMeans  float64 `json:"price_means" `
	PriceMedian float64 `json:"price_median" `
	PriceMin    float64 `json:"price_min" `
	PriceMax    float64 `json:"price_max" `
	PriceCount  int     `json:"price_count" `
	Messages    string  `json:"messages" `
}

// TopProductHistory ...
type TopProductHistory struct {
	TIVHistory      time.Time `json:"tiv_history"`
	TopSalesHistory time.Time `json:"top_sales_history"`
}

// ProductResponse ...
type ProductResponse struct {
	ID                 int64                        `json:"id"`
	ProductID          int64                        `json:"product_id"`
	ShopID             int64                        `json:"shop_id"`
	ShopName           string                       `json:"shop_name"`
	NormalizedName     string                       `json:"normalized_name" db:"normalized_name"`
	ProductName        string                       `json:"product_name"`
	ProductURL         string                       `json:"product_url"`
	ImageURL           string                       `json:"image_url" db:"image_url"`
	Price              float64                      `json:"price"`
	DiscountedPrice    float64                      `json:"discounted_price"`
	PriceMin           float64                      `json:"price_min"`
	PriceMax           float64                      `json:"price_max"`
	CategoryL1         CategoryL1                   `json:"category_l1"`
	CategoryL2         CategoryL2                   `json:"category_l2"`
	CategoryL3         CategoryL3                   `json:"category_l3"`
	ProcessStep        int                          `json:"process_step"`
	FilterTemplateID   int64                        `json:"filter_template_id"`
	CreateBy           int64                        `json:"create_by"`
	CreateTime         time.Time                    `json:"create_time"`
	UpdateTime         time.Time                    `json:"update_time"`
	Status             int                          `json:"status"`
	ClosestProductURL  string                       `json:"closest_product_url"`
	ClosestProductName string                       `json:"closest_product_name"`
	MarketplaceReports map[string]MarketplaceReport `json:"marketplace_reports,omitempty"`
	CalculatedPrices   CalculatedPrices             `json:"calculated_price" db:"calculated_price"`
	VerificationStatus VerificationStatus           `json:"verification_status" `
}

//VerificationStatus store verif info per input product
type VerificationStatus struct {
	CrawledProductCount int `json:"crawled_product_count"`
	TrueCount           int `json:"true_count"`
	FalseCount          int `json:"false_count"`
	UnlabeledCount      int `json:"unlabeled_count"`
}

type CrawledProductResponse struct {
	CrawledProduct []modelProduct.Product `json:"crawled_products"`
	Statistic      Statistic              `json:"statistic"`
	Meta           Meta                   `json:"meta"`
}

type SummaryResponse struct {
	Total               int   `json:"total"`
	TotalCompetitive    int   `json:"total_competitive"`
	VerifiedTrue        int   `json:"verified_true"`
	VerifiedFalse       int   `json:"verified_false"`
	TotalCrawledProduct int64 `json:"total_crawled_product"`
}

type Statistic struct {
	Marketplace  Marketplace  `json:"marketplace"`
	OSType       OSType       `json:"os_type"`
	Verification Verification `json:"verification"`
	PriceType    PriceType    `json:"price_type"`
}

type Marketplace struct {
	Shopee    int `json:"shopee"`
	Lazada    int `json:"lazada"`
	Tokopedia int `json:"tokopedia"`
}

type OSType struct {
	OS  int `json:"os"`
	C2C int `json:"c2c"`
}

type Verification struct {
	True      int `json:"true"`
	False     int `json:"false"`
	Unlabeled int `json:"unlabeled"`
}

type PriceType struct {
	SlashPrice  int `json:"slash_price"`
	NormalPrice int `json:"normal_price"`
}

type Meta struct {
	TotalProduct int64 `json:"total_product"`
	Limit        int   `json:"limit"`
	Offset       int   `json:"offset"`
}

type MarketplaceReport struct {
	Success bool     `json:"success"`
	Errors  []string `json:"errors,omitempty"`
}

type CategoryL1 struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type CategoryL2 struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type CategoryL3 struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

//FilterProductCrawl ...
type FilterProductCrawl struct {
	Offset     int
	Limit      int
	Title      string
	ProductID  int64
	ShopIDs    []int64
	CategoryL1 int64
	CategoryL2 int64
	CategoryL3 int64
}

//IrisProduct is a struct to save iris msg
type IrisProduct struct {
	CatLevel1    int64  `json:"cat_level_1"`
	CatLevel2    int64  `json:"cat_level_2"`
	CatLevel3    int64  `json:"cat_level_3"`
	ProductName  string `json:"product_name"`
	ProductURL   string `json:"product_url"`
	ProductPrice int64  `json:"product_price"`
	LazadaURL    string `json:"lazada_url"`
	LazadaPrice  int64  `json:"lazada_price"`
	ShopeeURL    string `json:"shopee_url"`
	ShopeePrice  int64  `json:"shopee_price"`
	TkpdURL      string `json:"tkpd_url"`
	TkpdPrice    int64  `json:"tkpd_price"`
	CreateTime   string `json:"create_time"`
	Type         int    `json:"type"`
}

//InputRequest ...
type InputRequest struct {
	ProductID        int64 `json:"product_id"`
	FilterTemplateID int64 `json:"filter_template_id"`
}
type VerificationParam struct {
	ID               int64  `json:"id"`
	ProductID        int64  `json:"product_id"`
	CrawledProductID string `json:"crawled_product_id"`
	Status           int    `json:"status"`
	CreateBy         int64  `json:"create_by"`
	UpdateBy         int64  `json:"update_by"`
	UserID           int64  `json:"user_id"`
}

//AnalyseResponse for API response
type AnalyseResponse struct {
	Result         ProductResponse      `json:"result"`
	Meta           AnalyseMeta          `json:"meta"`
	NearestProduct modelProduct.Product `json:"nearest_product"`
}

//AnalyseMeta for transparancy purpose for client, how we got the result
type AnalyseMeta struct {
	CrawledProductCount  int64 `json:"crawled_product_count"`
	VerifiedProductCount int64 `json:"verified_product_count"`
	FilteredProductCount int64 `json:"filtered_product_count"`
	FormulaCode          int   `json:"formula_code"`
}

//KeywordVerificationParam for decode from request body
type KeywordVerificationParam struct {
	ID      int64  `json:"id"`
	Keyword string `json:"verified_keyword"`
}

// PriceData ...
type PriceData struct {
	PriceChangeCount int `json:"price_change_count" db:"price_change_count"`
	TotalProduct     int `json:"total_product" db:"total_product"`
}

// ProductSummaryParam ...
type ProductSummaryParam struct {
	ProcessStep []int
	Time        time.Time
	ShopeeCd    int
	LazadaCd    int
	TkpdCd      int
}

// ProductCountPerMarketplace ...
type ProductCountPerMarketplace struct {
	SuccessShopee int64  `json:"success_shopee" db:"success_shopee"`
	FailedShopee  int64  `json:"failed_shopee" db:"failed_shopee"`
	SuccessLazada int64  `json:"success_lazada" db:"success_lazada"`
	FailedLazada  int64  `json:"failed_lazada" db:"failed_lazada"`
	SuccessTkpd   int64  `json:"success_tkpd" db:"success_tkpd"`
	FailedTkpd    int64  `json:"failed_tkpd" db:"failed_tkpd"`
	ErrorShopee   string `json:"shopee_error" db:"shopee_error"`
	ErrorLazada   string `json:"lazada_error" db:"lazada_error"`
	ErrorTkpd     string `json:"tkpd_error" db:"tkpd_error"`
}

// TotalProductSummary ...
type TotalProductSummary struct {
	TotalProduct     int64 `json:"total_product" db:"total_product"`
	TotalPriceChange int64 `json:"total_price_change" db:"total_price_change"`
}

// WeeklyCrawlingSummary ...
type WeeklyCrawlingSummary struct {
	Date          string               `json:"date"`
	CopyrightYear string               `json:"copyright_year"`
	Data          []WeeklyCrawlingData `json:"data"`
}

// WeeklyCrawlingData ...
type WeeklyCrawlingData struct {
	CategoryName                  string `json:"category_name"`
	TotalProductCrawled           string `json:"total_product_crawled"`
	TotalCompetitive              string `json:"total_competitive"`
	PercentageCompetitive         string `json:"percentage_competitive"`
	TotalTIV                      string `json:"total_tiv"`
	PercentageTIV                 string `json:"percentage_tiv"`
	TotalTopOrder                 string `json:"total_top_order"`
	PercentageTopOrder            string `json:"percentage_top_order"`
	TotalVerified                 string `json:"total_verified"`
	TotalVerifiedCompetitive      string `json:"total_verified_competitive"`
	PercentageVerifiedCompetitive string `json:"percentage_verified_competitive"`
}
