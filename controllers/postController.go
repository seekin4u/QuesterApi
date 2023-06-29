package controllers

import (
	"QuesterApi/initializers"
	"QuesterApi/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {

	questTime := models.QuestTime{
		Time: time.Now().Unix(),
		Quest: models.QuestStructure{Content: "content", Character: "Character1",
			QuestReward: models.QuestDescription{
				QuestgiverName:               "QGname",
				RewardLp:                     "1024",
				RewardExp:                    "2048",
				RewardLocalQuality:           "Auroch",
				RewardLocalQualityAdditional: "",
				RewardBy:                     "1",
				RewardItem:                   "",
			}},
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

func GetAll(c *gin.Context) {
	var qt []models.QuestTime
	initializers.DB.Preload("Quest.QuestReward").Find(&qt)

	c.JSON(200, gin.H{
		"all": qt,
	})

}

func GetAllQ(c *gin.Context) {
	var qd []models.QuestDescription
	initializers.DB.Where("questgiver_name = ?", "QGname").Find(&qd)

	c.JSON(200, gin.H{
		"GCname": qd,
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
	var qd []models.QuestStructure
	initializers.DB.Find(&qd)

	c.JSON(200, gin.H{
		"descriptions": qd,
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
	fmt.Println("---------")

	fmt.Println("Time is " + string(strconv.FormatInt(questTime.Time, 10)))
	fmt.Println("	 Character:" + questTime.Quest.Character)
	fmt.Println("	 Content:" + questTime.Quest.Content)
	if len(questTime.Quest.QuestReward.QuestgiverName) != 0 {
		fmt.Println("	 	QG:" + questTime.Quest.QuestReward.QuestgiverName)
	}
	if len(questTime.Quest.QuestReward.RewardLp) != 0 {
		fmt.Println("		LP:" + questTime.Quest.QuestReward.RewardLp)
	}
	if len(questTime.Quest.QuestReward.RewardExp) != 0 {
		fmt.Println("		EXP:" + questTime.Quest.QuestReward.RewardExp)
	}
	if len(questTime.Quest.QuestReward.RewardLocalQuality) != 0 {
		fmt.Println("		LocalQ:" + questTime.Quest.QuestReward.RewardLocalQuality)
	}
	if len(questTime.Quest.QuestReward.RewardLocalQualityAdditional) != 0 {
		fmt.Println("		LocalQAdd:" + questTime.Quest.QuestReward.RewardLocalQualityAdditional)
	}
	if len(questTime.Quest.QuestReward.RewardBy) != 0 {
		fmt.Println("		LocalQ by:" + questTime.Quest.QuestReward.RewardBy)
	}
	if len(questTime.Quest.QuestReward.RewardItem) != 0 {
		fmt.Println("		Item:" + questTime.Quest.QuestReward.RewardItem)
	}

	fmt.Println("---------")
}

func HandleJson(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	var receivedQuest models.QuestStructure
	err = json.Unmarshal(body, &receivedQuest)
	checkError(err)
	w.WriteHeader(http.StatusOK)

	var questTime models.QuestTime
	questTime.Time = time.Now().Unix()
	questTime.Quest = receivedQuest

	fmt.Println("---------")

	fmt.Println("Time is " + string(strconv.FormatInt(questTime.Time, 10)))
	fmt.Println("	 Character:" + questTime.Quest.Character)
	fmt.Println("	 Content:" + questTime.Quest.Content)
	if len(questTime.Quest.QuestReward.QuestgiverName) != 0 {
		fmt.Println("	 	QG:" + questTime.Quest.QuestReward.QuestgiverName)
	}
	if len(questTime.Quest.QuestReward.RewardLp) != 0 {
		fmt.Println("		LP:" + questTime.Quest.QuestReward.RewardLp)
	}
	if len(questTime.Quest.QuestReward.RewardExp) != 0 {
		fmt.Println("		EXP:" + questTime.Quest.QuestReward.RewardExp)
	}
	if len(questTime.Quest.QuestReward.RewardLocalQuality) != 0 {
		fmt.Println("		LocalQ:" + questTime.Quest.QuestReward.RewardLocalQuality)
	}
	if len(questTime.Quest.QuestReward.RewardLocalQualityAdditional) != 0 {
		fmt.Println("		LocalQAdd:" + questTime.Quest.QuestReward.RewardLocalQualityAdditional)
	}
	if len(questTime.Quest.QuestReward.RewardBy) != 0 {
		fmt.Println("		LocalQ by:" + questTime.Quest.QuestReward.RewardBy)
	}
	if len(questTime.Quest.QuestReward.RewardItem) != 0 {
		fmt.Println("		Item:" + questTime.Quest.QuestReward.RewardItem)
	}

	fmt.Println("---------")

	//save to DB
	result := initializers.DB.Create(&questTime)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
