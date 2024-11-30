package model
type Blog struct {
    ID       int    `json:"id"`
    Title    string `json:"title"`
    CoverURL string `json:"coverURL"`
    Body     string `json:"body"`
}