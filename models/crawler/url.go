package models

type URL struct {
  ID int64 `json:"id"`
  PlatformID int64 `json:"platform_id"`
  URL string `json:"url"`
  CreatedAt int64 `json:"created_at"`
  LastCrawled int64 `json:"last_crawled"`
}

type UserURL struct {
  ID int64 `json:"id"`
  UserID int64 `json:"user_id"`
  URLID int64 `json:"url_id"`
}
