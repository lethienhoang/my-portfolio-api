package models

type Token struct {
	AccessToken string
	AccessUUID  string
	AtExpires   int64
}
