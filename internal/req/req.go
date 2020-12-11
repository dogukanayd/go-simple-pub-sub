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

	url := "" // put here to your webhook own url you can use https://webhook.site/
	req, _ := http.NewRequest("POST", url, bytes.NewReader(data))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Insider")
	req.Close = true

	_, _ = (&http.Client{}).Do(req)
}
