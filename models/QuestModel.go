package models

type QuestTime struct {
	ID      uint           `gorm:"primarykey"`
	Time    int64          `json:"time"`
	Quest   QuestStructure `json:"questStructure"`
	QuestID uint
}

type QuestStructure struct {
	ID            uint             `gorm:"primarykey"`
	Content       string           `json:"content"`
	Character     string           `json:"character"`
	QuestReward   QuestDescription `json:"quest"`
	QuestRewardID uint
}

type QuestDescription struct {
	ID                 uint   `gorm:"primarykey"`
	QuestgiverName     string `json:"questgiverName"`
	RewardLp           string `json:"rewardLp,omitempty"`
	RewardExp          string `json:"rewardExp,omitempty"`
	RewardLocalQuality string `json:"rewardLocalQuality,omitempty"`
	RewardBy           string `json:"rewardBy,omitempty"`
	RewardItem         string `json:"rewardItem,omitempty"`
}
