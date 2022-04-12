package lib

type User struct {
	objectType string
	ID         string `json:"id"` // CompositeKey(Email)
	Email      string `json:"email"`
	Name       string `json:"name"`
	Passwd     string `json:"passwd"`
	IsAdmin    bool   `json:"is_admin"`
	IsReviewer bool   `json:"is_reviewer"`
}
