package group

type CreateGroupRequest struct {
	Privacy     string `json:"privacy"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Location    string `json:"location"`
	Avatar      string `json:"avatar"`
	Cover       string `json:"cover"`
	Description string `json:"description"`
	Terms       string `json:"terms"`
}

type UpdateGroupRequest struct {
	Privacy     string `json:"privacy"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Location    string `json:"location"`
	Avatar      string `json:"avatar"`
	Cover       string `json:"cover"`
	Description string `json:"description"`
	Terms       string `json:"terms"`
}
