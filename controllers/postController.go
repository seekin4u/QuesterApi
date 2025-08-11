package controllers

import (
	"QuesterApi/initializers"
	"QuesterApi/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func randomChoice[T any](options []T) T {
	return options[rand.Intn(len(options))]
}

func PostCreate(c *gin.Context) {

	rand.Seed(time.Now().UnixNano())

	questgivers := []string{
		"Elder Brynjolf",
		"Lady Mira",
		"Captain Rurik",
		"Thane Astrid",
	}
	rewardLps := []string{"5", "10", "15", "20"}
	rewardExps := []string{"100", "250", "500", "750"}
	rewardQualities := []string{
		"Stinging Nettle",
		"Dragon's Breath",
		"Shadow Silk",
		"Golden Oak",
	}

	questTime := models.QuestTime{
		Time: time.Now().Unix(),
		Quest: models.QuestStructure{
			Content:   "Test client",
			Character: "Nords of Skyrim",
			QuestReward: models.QuestDescription{
				QuestgiverName:     randomChoice(questgivers),
				RewardLp:           randomChoice(rewardLps),
				RewardExp:          randomChoice(rewardExps),
				RewardLocalQuality: randomChoice(rewardQualities),
				RewardBy:           "1",
				RewardItem:         "testItem",
			},
		},
	}

	result := initializers.DB.Create(&questTime)

	if result.Error != nil {
		fmt.Println("PostCreate error :" + result.Error.Error())
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"questTime": questTime,
	})
}

func GetDescriptions(c *gin.Context) {
	var qd []models.QuestDescription
	initializers.DB.Find(&qd)

	c.JSON(200, gin.H{
		"structures": qd,
	})
}

func GetDescription(c *gin.Context) {
	var st models.QuestDescription
	initializers.DB.Last(&st)

	c.JSON(200, gin.H{
		"description": st,
	})
}

func GetStructure(c *gin.Context) {
	var st models.QuestStructure
	initializers.DB.Preload("QuestReward").Last(&st)

	c.JSON(200, gin.H{
		"structure": st,
	})
}

func GetStructures(c *gin.Context) {
	var st []models.QuestStructure
	initializers.DB.Preload("QuestReward").Find(&st)

	c.JSON(200, gin.H{
		"structures": st,
	})
}

func PostCreateFromJson(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("Error on reading body")
		c.Status(400)
		return
	}

	var receivedQuest models.QuestStructure
	err = json.Unmarshal(body, &receivedQuest)
	if err != nil {
		fmt.Println("Error on unmarshalling")
		c.Status(400)
		return
	}

	if receivedQuest.QuestReward.RewardLocalQuality == "Wild Windsown Weed" {
		receivedQuest.QuestReward.RewardLocalQuality = "WWW"
	}

	var questTime models.QuestTime
	questTime.Time = time.Now().Unix()
	questTime.Quest = receivedQuest

	printRecievedQuest(questTime)

	result := initializers.DB.Create(&questTime)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return
	}

}

func printRecievedQuest(questTime models.QuestTime) {
	fmt.Print("{")

	fmt.Print(string(strconv.FormatInt(questTime.Time, 10)))
	fmt.Print(" " + questTime.Quest.Character)
	fmt.Print(":" + questTime.Quest.Content)
	fmt.Print(" {")
	if len(questTime.Quest.QuestReward.QuestgiverName) != 0 {
		fmt.Print("QG:" + questTime.Quest.QuestReward.QuestgiverName)
	}
	if len(questTime.Quest.QuestReward.RewardLp) != 0 {
		fmt.Print(" LP:" + questTime.Quest.QuestReward.RewardLp)
	}
	if len(questTime.Quest.QuestReward.RewardExp) != 0 {
		fmt.Print(" EXP:" + questTime.Quest.QuestReward.RewardExp)
	}
	if len(questTime.Quest.QuestReward.RewardLocalQuality) != 0 {
		fmt.Print(" LQ:" + questTime.Quest.QuestReward.RewardLocalQuality)
	}
	if len(questTime.Quest.QuestReward.RewardBy) != 0 {
		fmt.Print(" BY:" + questTime.Quest.QuestReward.RewardBy)
	}
	if len(questTime.Quest.QuestReward.RewardItem) != 0 {
		fmt.Print(" IT:" + questTime.Quest.QuestReward.RewardItem)
	}
	fmt.Print("} ")

	fmt.Print("}")
}
