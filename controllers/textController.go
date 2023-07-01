package controllers

import (
	"QuesterApi/initializers"
	"QuesterApi/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetQGs(c *gin.Context) {
	var qgs []models.QuestDescription
	initializers.DB.Distinct("questgiver_name").Find(&qgs)

	fmt.Println(qgs)
	c.JSON(200, gin.H{
		"qgs": qgs,
	})
}

// передавать[] из []models.QuestDescription
// каждый элемент в себе будет нести список QuestDescription
// с разным инаградами, .questGiverName везде будет один
func GetQGsQs(c *gin.Context) {
	var qg []models.QuestDescription
	initializers.DB.Distinct("questgiver_name", "reward_local_quality", "reward_local_quality_additional").Find(&qg)

	c.JSON(200, gin.H{
		"qgsqs": qg,
	})
}

func GetQGGeneral(c *gin.Context) {
	var questgiver models.QuestDescription
	initializers.DB.Where("questgiver_name = ?", "QGname").Find(&questgiver)

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
		"qg": questgiver.QuestgiverName,
		"ql": quality,
	})
}

func GetQGSpecial(c *gin.Context) {
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
