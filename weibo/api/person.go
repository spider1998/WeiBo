package api

import (
	"WeiPro/weibo/entity"
	"WeiPro/weibo/util"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"reflect"
	"strconv"
	"time"
)

//获取微博个人信息接口
func GetWeibo(user, token string) (person entity.Person, err error) {
	urls := "https://api.weibo.com/2/users/show.json"
	args := make(map[string]string)
	args["screen_name"] = user
	args["access_token"] = token
	res, err := GetMethond(args, urls)
	if err != nil {
		return
	}
	v := reflect.ValueOf(&person).Elem()
	for i, re := range res.(map[string]interface{}) {
		if re == nil {
			continue
		}
		fil := util.CamelString(i)
		if v.FieldByName(fil).String() == "<invalid Value>" {
			continue
		}
		v.FieldByName(fil).Set(reflect.ValueOf(re))
	}

	/*err = mapstructure.Decode(res, &person)
	if err != nil {
		fmt.Println(err)
	}*/
	return
}

type Com struct {
	Comments []entity.Comment
}

//获取评论接口
func GetComments(weiID, token, page_count, count string) (comments []entity.Comment, err error) {
	var lens int
	countInt, err := strconv.Atoi(count)
	if err != nil {
		return
	}
	urls := "https://api.weibo.com/2/comments/show.json"
	args := make(map[string]string)
	args["id"] = weiID
	args["access_token"] = token
	args["count"] = page_count
	for i := 1; ; i++ {
		args["page"] = strconv.Itoa(i)
		res, err1 := GetMethond(args, urls)
		if err1 != nil {
			if err1.Error() == "User requests out of rate limit!" {
				err = err1
				break
			}
			err = err1
			return
		}
		var coms Com
		err = mapstructure.Decode(res, &coms)
		if err != nil {
			fmt.Println(err)
		}
		comments = append(comments, coms.Comments...)
		//如果大于等于预定条数或数量饱和则退出爬取
		if len(comments) >= countInt || len(comments) == lens {
			break
		}
		lens = len(comments)
		fmt.Println(strconv.Itoa(lens) + "load over...")
		//休眠，防止机器识别
		if i == 10 {
			break
		}
		time.Sleep(10 * time.Second)
	}
	return
}
