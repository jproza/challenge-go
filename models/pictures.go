package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func sanitizeToken(token string) string {
	byt := []byte(token)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	auth := dat["auth"].(bool)
	fmt.Println(auth)

	tk := dat["token"].(string)
	fmt.Println(tk)
	return tk
}

//GetAllPictures Fetch all pictures data
func GetAllPictures(pic *Pictures, token string) (err error) {
	url := "http://interview.agileengine.com/images"
	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + sanitizeToken(token)

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string([]byte(body)))

	s := string([]byte(body))
	data := Pictures{}
	json.Unmarshal([]byte(s), &data)

	//p := []*Pictures{}
	for item := 0; item < len(data.Pic); item++ {

		//fmt.Println(data.Pic[item].Id)
		//fmt.Println(data.Pic[item].Cropped_picture)

		//var rect1 = Pictures{items 1, 1}
		//fmt.Println(rect1)

		//pi := Pictures{}
		//pi.Pic[item].Id = data.Pic[item].Id
		//pi.Pic[item].Cropped_picture = data.Pic[item].Cropped_picture
		//pic = append(make(Pictures , 20), pi)
		//fmt.Println(p)
		//fmt.Println(pi.Pic[item].Id)
		//fmt.Println(pi.Pic[item].Cropped_picture
		ite := Items{Id: data.Pic[item].Id,
			Cropped_picture: data.Pic[item].Cropped_picture}
		pic.AddItem(ite)
		//pic.append(p) //pic = p
	}

	//}

	//fmt.Println("id: ", data.Pic.items.Id)
	//fmt.Printf("pic: %s", data.Cropped_picture)

	return nil
}

//CreatePictures ... Insert New data
func CreatePictures(pic *Pictures) (err error) {
	//if err = Config.DB.Create(user).Error; err != nil {
	//	return err
	//}
	return nil
}

//GetPicturesByID ... Fetch only one pic by Id
func GetPicturesByID(pic *Pictures, id string) (err error) {
	//if err = Config.DB.Where("id = ?", id).First(user).Error; err != nil {
	//	return err
	//}
	return nil
}

//UpdatePictures ... Update
func UpdatePictures(pic *Pictures, id string) (err error) {
	fmt.Println(pic)
	////Config.DB.Save(user)
	return nil
}

//DeletePictures ... Delete
func DeletePictures(pic *Pictures, id string) (err error) {
	///Config.DB.Where("id = ?", id).Delete(user)
	return nil
}
