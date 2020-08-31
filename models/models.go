package models

type Pictures struct {
	Pic []Items `json:"pictures"`
}

type Items struct {
	Id              uint   `json:"id"`
	Cropped_picture string `json:"cropped_picture"`
}


//AddItem it add
func (box *Pictures) AddItem(item Items) {
        box.Pic = append(box.Pic, item)
}

type Token struct {
	Auth  bool
	token string
}
