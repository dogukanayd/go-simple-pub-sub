package req

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Requester ..
type Requester struct {
	ID string `json:"id"`
}

// Send ..
func (r *Requester) Send() {
	data, _ := json.Marshal(r)

	req, _ := http.NewRequest("POST", "https://webhook.site/1a71823e-07a4-4f50-af9d-0f9f2649b07c", bytes.NewReader(data))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Insider")
	req.Close = true

	_, _ = (&http.Client{}).Do(req)
}
