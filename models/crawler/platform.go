package models

type Platform struct {
  ID int64 `json:"id"`
  Name string `json:"name"`
  ReadableName string `json:"readable_name"`
}
