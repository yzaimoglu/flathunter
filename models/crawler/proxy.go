package models

type Proxy struct {
  ID int64 `json:"id"`
  IP string `json:"ip"`
  Username string `json:"username"`
  Password string `json:"password"`
  Port int `json:"port"`
}
