package controller

import (
	"bytes"
	"challenge-go/models"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

var token string

//Login log into
func Login(c *gin.Context) {
	url := "http://interview.agileengine.com/auth"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(`
	{
	    "apiKey": "23567b218376f79d9415"
	}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	//req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	token = string(body)
}

//GetPictures ... Get all
func GetPictures(c *gin.Context) {
	var pic models.Pictures

	err := models.GetAllPictures(&pic, token)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, pic)
	}
}

//CreatePicture ... Create
func CreatePicture(c *gin.Context) {
	var pic models.Pictures
	c.BindJSON(&pic)
	err := models.CreatePictures(&pic)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, pic)
	}
}

//GetPictureByID ... Get the pic by id
func GetPictureByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var pic models.Pictures
	err := models.GetPicturesByID(&pic, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, pic)
	}
}

//UpdatePicture ... Update the pic information
func UpdatePicture(c *gin.Context) {
	var pic models.Pictures
	id := c.Params.ByName("id")
	err := models.GetPicturesByID(&pic, id)
	if err != nil {
		c.JSON(http.StatusNotFound, pic)
	}
	c.BindJSON(&pic)
	err = models.UpdatePictures(&pic, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, pic)
	}
}

//DeletePicture ... Delete
func DeletePicture(c *gin.Context) {
	var pic models.Pictures
	id := c.Params.ByName("id")
	err := models.DeletePictures(&pic, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
