package handler

type AddSkillRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
	Kind        string `json:"kind"`
}

type UpdateSkillRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
}
