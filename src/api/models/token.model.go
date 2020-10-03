package models

type Token struct {
	AccessToken string `json:"access_token"`
	AccessUUID  string `json:"access_id"`
	AtExpires   int64  `json:"at_expires"`
}
