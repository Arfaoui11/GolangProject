package models

type Adrress struct {
	Stat    string `json:"stat" bson:"stat"`
	City    string `json:"city" bson:"city"`
	Pincode int    `json:"pincode" bson:"pincode"`
}

type User struct {
	Name    string  `json:"name" bson:"name"`
	Age     int     `json:"age" bson:"age"`
	Adrress Adrress `json:"address" bson:"address"`
}
