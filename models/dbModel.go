package models

type Account struct {
	Username string `gorm:"type:varchar(20)" json:"username"`
	UID  int `json:"uid" gorm:"type:int;primary_key"`
	Pwd string `json:"pwd" gorm:"type:varchar(255)"`
}
