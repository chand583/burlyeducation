package models

import (
	//"fmt"
	"burlyeducation/log"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

const listkey string = "buly-allusers"

type UserModel struct{}

type User struct {
	Id        int       `orm:"column(id)"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Mobile    string    `json:"mobile"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}

func init() {
	orm.RegisterModel(new(User))
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
	// if err == nil {
	// 	var key string = "einfo-post-" + strconv.Itoa(legacyId)
	// 	err1 := bm.Delete(context.Background(), key)
	// 	err2 := bm.Delete(context.Background(), listkey)
	// 	if err1 != nil {
	// 		log.Error(1055, map[string]interface{}{"error_details": err1, "payload_details": key})
	// 	}
	// 	if err2 != nil {
	// 		log.Error(1055, map[string]interface{}{"error_details": err2, "payload_detils": listkey})
	// 	}

	// }
	return
}

/*
 * Method to return save Posts
 */
func (a UserModel) SaveUser(user interface{}) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(user)
	if err != nil {
		log.Error(1001, map[string]interface{}{"error_details": err})
	}
	return
}

/*
 * Method to return Update Article
 */
func (a UserModel) UpdateUser(user interface{}) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Update(user)
	if err != nil {
		log.Error(1001, map[string]interface{}{"error_details": err})
	}
	return
}
