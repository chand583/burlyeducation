package models

import (
	"burlyeducation/log"
	"context"
	"encoding/json"

	//"fmt"

	"strconv"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
)

const listkey string = "s"

type UserModel struct{}

type User struct {
	Id        int       `orm:"column(id)"`
	UserName  string    `json:"user_name"`
	Email     string    `json:"email"`
	Mobile     string    `json:"mobile"`
	Password  string    `json:"password"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type UserApiRes struct {
	Id        int       `orm:"column(id)"`
	UserName  string    `json:"user_name"`
	Email     string    `json:"email"`
	Mobile     string    `json:"mobile"`
	Password  string    `json:"password"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func init() {
	//orm.RegisterModel(new(Article))
}
func AddUser(m *User) (id int64, err error) {

	o := orm.NewOrm()
	id, err = o.Insert(m)
	if err != nil {
		log.Error(1001, map[string]interface{}{"error_details": err})
	}

	return
}

/*
 * Method to return save Posts
 */
func (a UserModel) SaveUpdateUser(user interface{}, userId int) (id int64, err error) {
	if userId == 0 {
		id, err = a.SaveUser(user)
	} else {
		id, err = a.UpdateUser(user)
	}
	return
}

/*
 * Method to return save Posts
 */
func (a UserModel) SaveUser(article interface{}) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(article)
	if err != nil {
		log.Error(1001, map[string]interface{}{"error_details": err})
	}
	return
}

/*
 * Method to return Update Article
 */
func (a UserModel) UpdateUser(article interface{}) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Update(article)
	if err != nil {
		log.Error(1001, map[string]interface{}{"error_details": err})
	}
	return
}

/*
 * Method to return all Posts which is published and not expired
 */
func (a UserModel) GetAll() interface{} {
	var articles []User
	articleList := make([]interface{}, 0)
	//Read From Cache if enabled
	if isCacheEnable {
		articles = a.getUserListFromCache()
	} else {
		_, articles = a.getUserListFromDB(0)
	}
	if len(articles) > 0 {
		for _, article := range articles {
			objArticles := UserApiRes{}
			objArticles = a.prepareData(article)
			articleList = append(articleList, objArticles)
		}
		return articleList
	} else {
		return articleList
	}

}

/*
 * Read data from Cache if exist otherwise read from DB and store it to Cache for futher use
 * TO DO - Keep only latest 50 articles in Cache [in to do task list]
 */
func (a UserModel) getUserListFromCache() []User {
	var users []User
	var totalArticles int64
	val, _ := bm.Get(context.Background(), listkey)
	if val == nil {
		totalArticles, users = a.getUserListFromDB(0)
		if totalArticles > 0 {
			caheExpiryTime, _ := config.Int("cache::cache_expiry_time_insec")
			strToCache, err := json.Marshal(users)
			if err != nil {
				log.Error(1508, map[string]interface{}{"error_details": err})
			}

			bm.Put(context.Background(), listkey, strToCache, time.Duration(caheExpiryTime)*time.Second)
		}
	} else {
		//fmt.Println("From Cache")
		totalArticles = 1
		resUint8 := val.([]uint8)
		strResponse := string(resUint8)
		json.Unmarshal([]byte(strResponse), &users)
	}
	return users
}

/*
 * Read data from db and return the result
 * @params - int [LegacyId], Pass 0 if need all articles
 * @return
 * 	- int [totalArticles]
 * 	- Article [allPublishedArticles - either all articles or Article related to legacyId]
 */
func (a UserModel) getUserListFromDB(id int) (int64, []User) {
	var users []User
	var totalUsers int64
	currentTimestamp := time.Time.UTC(time.Now()).Format(time.RFC3339)
	sql := "SELECT * from users"
	o := orm.NewOrm()
	usersSet, err := o.Raw(sql, currentTimestamp).QueryRows(&users)

	totalUsers = usersSet
	if err == nil && totalUsers > 0 {
		return totalUsers, users
	}
	if err != nil {
		log.Warning(1001, map[string]interface{}{"error_details": err})
	}
	return 0, users
}

/*
 * Method to check if Article published date is less than 24 hours then return true otherwise false
 * @param - updatedDate time
 * @return - boolean
 * Convert article updated date and current time to same TIME ZONE and then do the compare
 */

func (a UserModel) isNewUser(updatedDate time.Time) bool {
	userUpdatedTimestamp := updatedDate.Local().Unix()
	currentTimestamp := time.Now().Local().Unix()
	diffSec := currentTimestamp - userUpdatedTimestamp

	if diffSec < 86400 {
		return true
	} else {
		return false
	}
	//pastTime := updatedDate.Local().Add(-time.Hour * 24)
	//fmt.Println(pastTime)
	return false

}

/*
 * Prepare API response data as per ArticleApiRes structure
 */
func (a UserModel) prepareData(user User) UserApiRes {
	objUsers := UserApiRes{}

	return objUsers
}

/*
 * Read data from Cache if exist otherwise read from DB and store it to Cache for futher use
 */
func (a UserModel) getUserDetailFromCache(id int) []User {
	var userDetails []User
	var key string = "burly-userd-" + strconv.Itoa(id)
	var totalArticles int64
	val, err := bm.Get(context.Background(), key)
	if err != nil {
		log.Error(1052, map[string]interface{}{"error_details": err})
	}
	if val == nil {
		totalArticles, userDetails = a.getUserListFromDB(id)
		if totalArticles > 0 {
			caheExpiryTime, _ := config.Int("cache::cache_expiry_time_insec")
			strToCache, err := json.Marshal(userDetails)
			if err != nil {
				log.Error(1508, map[string]interface{}{"error_details": err})
			}
			bm.Put(context.Background(), key, strToCache, time.Duration(caheExpiryTime)*time.Second)
		}
	} else {
		//fmt.Println("From Cache")
		totalArticles = 1
		resUint8 := val.([]uint8)
		strResponse := string(resUint8)
		json.Unmarshal([]byte(strResponse), &userDetails)
	}
	return userDetails
}
