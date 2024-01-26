package models

type Skill struct {
	SkillId   int64  `json:"skillId" db:"skillId"`
	Category  string `json:"category" db:"category"`
	SkillName string `json:"skillName" db:"skillName"`
}