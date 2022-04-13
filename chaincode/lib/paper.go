package lib

import "strconv"

type Paper struct {
	ID          string    `json:"id"`       // SHA256(Uploader + UploadTime)
	Uploader    string    `json:"uploader"` // User email
	UploadTime  int64     `json:"upload_time"`
	Title       string    `json:"title"`
	Abstract    string    `json:"abstract"`
	Authors     []string  `json:"authors"`
	Keywords    []string  `json:"keywords,omitempty"`
	PeerReviews [3]string `json:"peer_reviews"` // PeerReview emails
	Status      string    `json:"status"`
	ReviewTime  int64     `json:"review_time"`
}

func (p Paper) Type() string {
	return ObjectTypePaper
}

func (p Paper) Keys() []string {
	return []string{p.Uploader + strconv.FormatInt(p.UploadTime, 10)}
}
