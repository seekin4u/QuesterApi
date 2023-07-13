package controllers

import (
	"QuesterApi/initializers"
	"QuesterApi/models"

	"github.com/gin-gonic/gin"
)

func GetQuests(c *gin.Context) {
	var qt []models.QuestTime
	initializers.DB.Preload("Quest.QuestReward").
		Limit(50).
		Order("id desc").
		Find(&qt)

	c.JSON(200, gin.H{
		"array": qt,
	})
}

func GetQuestsNpc(c *gin.Context) {
	npc := c.Param("npc")
	if len(npc) == 0 {
		npc = ""
	}

	var quests []models.QuestTime
	initializers.DB.
		Joins("JOIN quest_structures on quest_structures.id=quest_id").
		Joins("JOIN quest_descriptions on quest_descriptions.id=quest_reward_id").
		Distinct("questgiver_name").
		Where("questgiver_name = ?", npc).
		Preload("Quest.QuestReward").
		Find(&quests)

	c.JSON(200, gin.H{
		"array": quests,
	})

}

func GetQuestsQuality(c *gin.Context) {
	quality := c.Param("quality")
	if len(quality) == 0 {
		quality = ""
	}

	var quests []models.QuestTime
	initializers.DB.
		Joins("JOIN quest_structures on quest_structures.id=quest_id").
		Joins("JOIN quest_descriptions on quest_descriptions.id=quest_reward_id").
		Where("reward_local_quality = ?", quality).
		Preload("Quest.QuestReward").
		Find(&quests)

	c.JSON(200, gin.H{
		"array": quests,
	})

}

func GetQuestgivers(c *gin.Context) {
	var qgs []models.QuestDescription
	initializers.DB.
		Distinct("questgiver_name").
		Find(&qgs)

	var npcList []string

	for _, el := range qgs {
		if !contains(npcList, el.QuestgiverName) {
			npcList = append(npcList, el.QuestgiverName)
		}
	}

	var lpsum int
	initializers.DB.
		Raw("SELECT sum(quest_descriptions.reward_lp::numeric) from quest_descriptions WHERE quest_descriptions.reward_lp IS DISTINCT FROM ''").Scan(&lpsum)

	var expsum int
	initializers.DB.
		Raw("SELECT sum(quest_descriptions.reward_exp::numeric) from quest_descriptions WHERE quest_descriptions.reward_exp IS DISTINCT FROM ''").Scan(&expsum)

	c.JSON(200, gin.H{
		"qgs":  npcList,
		"tlp":  lpsum,
		"texp": expsum,
	})
}

func GetQuestgiver(c *gin.Context) {
	npc := c.Param("npc")
	if len(npc) == 0 {
		npc = ""
	}

	var qds []models.QuestDescription
	initializers.DB.Where("questgiver_name = ?", npc).
		Distinct("questgiver_name", "reward_local_quality").
		Find(&qds)

	var qualities []string
	for _, val := range qds {
		if len(val.RewardLocalQuality) != 0 && !contains(qualities, val.RewardLocalQuality) {
			qualities = append(qualities, val.RewardLocalQuality)
		}
	}

	var lpsum int
	initializers.DB.
		Raw("SELECT sum(quest_descriptions.reward_lp::numeric) from quest_descriptions WHERE quest_descriptions.questgiver_name = ? AND quest_descriptions.reward_lp IS DISTINCT FROM ''", npc).Scan(&lpsum)

	var expsum int
	initializers.DB.
		Raw("SELECT sum(quest_descriptions.reward_exp::numeric) from quest_descriptions WHERE quest_descriptions.questgiver_name = ? AND quest_descriptions.reward_exp IS DISTINCT FROM ''", npc).Scan(&expsum)

	c.JSON(200, gin.H{
		"qg":   npc,
		"ql":   qualities,
		"tlp":  lpsum,
		"texp": expsum,
	})
}

func GetQuestgiverQuests(c *gin.Context) {
	npc := c.Param("npc")
	if len(npc) == 0 {
		npc = ""
	}

	var quests []models.QuestTime
	initializers.DB.
		Joins("JOIN quest_structures on quest_structures.id=quest_id").
		Joins("JOIN quest_descriptions on quest_descriptions.id=quest_reward_id").
		Where("questgiver_name = ?", npc).
		Preload("Quest.QuestReward").
		Find(&quests)

	c.JSON(200, gin.H{
		"array": quests,
	})
}

func GetQuestgiverQualitiesQuests(c *gin.Context) {
	npc := c.Param("npc")
	if len(npc) == 0 {
		npc = ""
	}

	var qds []models.QuestDescription
	initializers.DB.
		Where("questgiver_name = ?", npc).
		Distinct("questgiver_name", "reward_local_quality", "reward_local_quality_additional").
		Find(&qds)

	var qualities []string
	for _, val := range qds {
		if len(val.RewardLocalQuality) != 0 && !contains(qualities, val.RewardLocalQuality) {
			qualities = append(qualities, val.RewardLocalQuality)
		}
	}

	c.JSON(200, gin.H{
		"qg": npc,
		"ql": qualities,
	})
}

func GetQgQlGeneric(c *gin.Context) {
	var qualities []models.QuestDescription
	initializers.DB.Where("questgiver_name = ?", "QGname").Distinct("questgiver_name", "reward_local_quality").Find(&qualities)

	var quality []string
	for _, val := range qualities {
		if len(val.RewardLocalQuality) != 0 && !contains(quality, val.RewardLocalQuality) {
			quality = append(quality, val.RewardLocalQuality)
		}
	}

	c.JSON(200, gin.H{
		"qg": "QGname",
		"ql": quality,
	})
}

func GetQuestgiver1(c *gin.Context) {
	npc := c.Param("npc")
	if len(npc) == 0 {
		npc = ""
	}

	var qualities []models.QuestDescription
	initializers.DB.Where("questgiver_name = ?", npc).Distinct("questgiver_name", "reward_local_quality").Find(&qualities)

	c.JSON(200, gin.H{
		"ql": qualities,
	})
}

func GetQualities(c *gin.Context) {

	var qls []models.QuestDescription
	initializers.DB. //TODO: remove qualities string and give QuestReward set?
				Distinct("reward_local_quality").
				Find(&qls)

	var qualities []string

	for _, el := range qls {
		if len(el.RewardLocalQuality) != 0 && !contains(qualities, el.RewardLocalQuality) {
			qualities = append(qualities, el.RewardLocalQuality)
		}
	}

	c.JSON(200, gin.H{
		"qls": qualities,
	})

}

func GetQualityQuestgiversSum(c *gin.Context) {
	quality := c.Param("quality")
	if len(quality) == 0 {
		quality = ""
	}

	/*SELECT DISTINCT quest_descriptions.questgiver_name, quest_descriptions.reward_local_quality
	  FROM quest_descriptions WHERE quest_descriptions.reward_local_quality IS DISTINCT FROM ''
	  AND quest_descriptions.questgiver_name = 'QGname'
	  UNION
	  SELECT DISTINCT quest_descriptions.questgiver_name, quest_descriptions.reward_local_quality_additional
	  FROM quest_descriptions WHERE quest_descriptions.reward_local_quality_additional IS DISTINCT FROM ''
	  AND quest_descriptions.questgiver_name = 'QGname'

	SELECT DISTINCT quest_descriptions.questgiver_name FROM quest_descriptions
	WHERE  quest_descriptions.reward_local_quality = 'Badger'
	UNION
	SELECT DISTINCT quest_descriptions.questgiver_name FROM quest_descriptions
	WHERE  quest_descriptions.reward_local_quality_additional = 'Badger' */

	/*Select("questgiver_name").
	Where("reward_local_quality IS DISTINCT FROM ''").
	Distinct("reward_local_quality").
	Find(&qls)*/

	/*WITH wt AS(
	SELECT DISTINCT quest_descriptions.id, quest_descriptions.reward_local_quality, quest_descriptions.questgiver_name FROM quest_descriptions WHERE  quest_descriptions.reward_local_quality = 'Badger'
	UNION
	SELECT DISTINCT quest_descriptions.id, quest_descriptions.reward_local_quality_additional, quest_descriptions.questgiver_name FROM quest_descriptions WHERE  quest_descriptions.reward_local_quality_additional = 'Badger')
	select quest_times.id, quest_structures.id, wt.id,  quest_structures.character from quest_times inner join
	quest_structures  on quest_times.quest_id = quest_structures.id inner join
	wt on quest_structures.quest_reward_id = wt.id*/
	var qgs []models.QuestDescription
	initializers.DB. //TODO: remove qualities string and give QuestReward set?
				Distinct("questgiver_name").
				Where("reward_local_quality = ?", quality).
				Find(&qgs)

	var qualsum int
	initializers.DB.
		Raw("SELECT DISTINCT sum(quest_descriptions.reward_by::numeric) from quest_descriptions WHERE quest_descriptions.reward_local_quality = ? group by quest_descriptions.reward_local_quality", quality).Scan(&qualsum)

	c.JSON(200, gin.H{
		"ql":  quality,
		"qgs": qgs,
		"ups": qualsum,
	})

}

func GetQualityQuests(c *gin.Context) {

	var qls []models.QuestDescription
	initializers.DB. //TODO: remove qualities string and give QuestReward set?
				Distinct("reward_local_quality").
				Find(&qls)

	var qualities []string

	for _, el := range qls {
		if len(el.RewardLocalQuality) != 0 && !contains(qualities, el.RewardLocalQuality) {
			qualities = append(qualities, el.RewardLocalQuality)
		}
	}

	c.JSON(200, gin.H{
		"qls": qualities,
	})

}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
