package shared_types

const WORK_TYPE_SCRAPE_URL = "SCRAPE_URL"

type ScrapeUrl struct {
	MessagingBase
	Url   string
	Flags string
}
