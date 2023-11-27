package shortener

type Redirect struct {
	Code      string `json:"code" bson:"code"` // I have to goes bson for mongoDB
	URL       string `json:"url" bson:"url"`
	CreatedAt int64  `json:"createdAt" bson:"createdAt"`
}
