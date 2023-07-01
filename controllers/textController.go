package controllers

import (
	"QuesterApi/initializers"
	"QuesterApi/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetQgs(c *gin.Context) {
	var qgs []models.QuestDescription
	initializers.DB.Distinct("questgiver_name").Find(&qgs)

	var npcList []string

	for _, el := range qgs {
		if !contains(npcList, el.QuestgiverName) {
			npcList = append(npcList, el.QuestgiverName)
		}
	}

	fmt.Println(qgs)
	c.JSON(200, gin.H{
		"qgs": npcList,
	})
}

func GetQgQs(c *gin.Context) {
	npc := c.Param("npc")
	if len(npc) == 0 {
		npc = ""
	}

	var qds []models.QuestDescription
	initializers.DB.Where("questgiver_name = ?", npc).
		Distinct("questgiver_name", "reward_local_quality", "reward_local_quality_additional").
		Find(&qds)

	var qualities []string
	for _, val := range qds {
		if len(val.RewardLocalQuality) != 0 && !contains(qualities, val.RewardLocalQuality) {
			qualities = append(qualities, val.RewardLocalQuality)
		} else if len(val.RewardLocalQualityAdditional) != 0 && !contains(qualities, val.RewardLocalQualityAdditional) {
			qualities = append(qualities, val.RewardLocalQualityAdditional)
		}
	}

	c.JSON(200, gin.H{
		"qg": npc,
		"ql": qualities,
	})
}

func GetQgQlGeneric(c *gin.Context) {
	var qualities []models.QuestDescription
	initializers.DB.Where("questgiver_name = ?", "QGname").Distinct("questgiver_name", "reward_local_quality", "reward_local_quality_additional").Find(&qualities)

	var quality []string
	for _, val := range qualities {
		if len(val.RewardLocalQuality) != 0 && !contains(quality, val.RewardLocalQuality) {
			quality = append(quality, val.RewardLocalQuality)
		} else if len(val.RewardLocalQualityAdditional) != 0 && !contains(quality, val.RewardLocalQualityAdditional) {
			quality = append(quality, val.RewardLocalQualityAdditional)
		}
	}

	c.JSON(200, gin.H{
		"qg": "QGname",
		"ql": quality,
	})
}

func GetQgSpecial(c *gin.Context) {
	npc := c.Param("npc")
	if len(npc) == 0 {
		npc = ""
	}

	var qualities []models.QuestDescription
	initializers.DB.Where("questgiver_name = ?", npc).Distinct("questgiver_name", "reward_local_quality", "reward_local_quality_additional").Find(&qualities)

	c.JSON(200, gin.H{
		"ql": qualities,
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
