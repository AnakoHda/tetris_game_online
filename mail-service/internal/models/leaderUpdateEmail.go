package models

type LeaderUpdateEmail struct {
	Email     string `json:"email"`
	Nickname  string `json:"nickname"`
	NewLeader string `json:"new_leader"`
	NewScore  int    `json:"new_score"`
}
