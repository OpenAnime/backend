package constani

type User struct {
	Id string `json:"id" bson:"id"`

	Name string `json:"name" bson:"name"`

	Email string `json:"email" bson:"email"`

	Password string `json:"-"`

	CreatedAt string `json:"created_at"`

	Avatar string `json:"avatar"`

	Banner string `json:"banner"`

	PlayList []string `json:"play_list"`

	AnimeList []string `json:"anime_list"`

	IsAdmin bool `json:"is_admin"`

	Presence string `json:"presence"`
}
