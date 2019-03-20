package entity

const TablePerson = "person"

//人员数据库实体
type Person struct {
	ID               float64 `json:"id" gorm:"type:bigint"` //id
	ScreenName       string  `json:"screen_name"`           //昵称
	Name             string  `json:"name"`                  //友好显示名称
	Province         string  `json:"province"`              //用户所在省级ID
	City             string  `json:"city"`                  //用户所在城市ID
	Location         string  `json:"location"`              //用户所在地
	Description      string  `json:"description"`           //用户个人描述
	Url              string  `json:"url"`                   //用户博客地址
	ProfileImageUrl  string  `json:"profile_image_url"`     //用户头像地址
	ProfileUrl       string  `json:"profile_url"`           //用户的微博统一URL地址
	Domain           string  `json:"domain"`                //用户的个性化域名
	Gender           string  `json:"gender"`                //性别，m：男、f：女、n：未知
	FollowersCount   float64 `json:"followers_count"`       //粉丝数
	FriendsCount     float64 `json:"friends_count"`         //关注数
	StatusesCount    float64 `json:"statuses_count"`        //微博数
	FavouritesCount  float64 `json:"favourites_count"`      //收藏数
	CreatedAt        string  `json:"created_at"`            //用户创建（注册）时间
	BiFollowersCount float64 `json:"bi_followers_count"`    //用户的互粉数
	OnlineStatus     float64 `json:"online_status"`         //用户的在线状态，0：不在线、1：在线
	AvatarLarge      string  `json:"avatar_large"`          //用户头像地址（大图），180×180像素
	AvatarHd         string  `json:"avatar_hd"`             //用户头像地址（高清），高清头像原图
	AbilityTags      string  `json:"ability_tags"`          //标签
	Lang             string  `json:"lang"`                  //用户当前的语言版本，zh-cn：简体中文，zh-tw：繁体中文，en：英语
}

func (Person) TableName() string {
	return TablePerson
}
