package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/RayFantasyStudio/blog/utils"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/astaxie/beego"
	"time"
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego/context"
	"encoding/json"
)

// 登陆有效期60天
const UserTokenPeriod = 60 * 24 * time.Hour

type User struct {
	Id   int `form:"_" json:"id"`
	Name string `orm:"size(16);unique" form:"name" json:"name" valid:"Required"`
	Pwd  string `orm:"size(32)" form:"pwd" json:"_" valid:"MinSize(8);MaxSize(32)"`
}

type userToken struct {
	User       User
	ExpireTime time.Time
	IsExpire   bool
}

var UserTokenPrefix string

func initUser() {
	UserTokenPrefix = beego.AppConfig.String("appname") + "_usertoken:"
}

// 新建用户
func (u *User) New() error {
	if err := checkUserValid(u); err != nil {
		return err
	}

	o := orm.NewOrm()

	// 寻找用户名相同的用户是否已经存在
	o.Read(u, "Name")
	if u.Id != 0 {
		return fmt.Errorf("用户名为\"%s\"的用户已存在", u.Name)
	}

	if _, err := o.Insert(u); err != nil {
		return err
	}

	return nil
}

// 用户登录
func (u *User) Login() (token string, err error) {
	if err = checkUserValid(u); err != nil {
		return
	}

	o := orm.NewOrm()

	user := User{Name:u.Name}
	o.Read(&user, "Name")
	if user.Id == 0 {
		err = fmt.Errorf("找不到用户名为\"%s\"的用户", u.Name)
		return
	}

	if user.Pwd != u.Pwd {
		err = fmt.Errorf("输入的与密码用户\"%s\"的密码不符", u.Name)
		return
	}

	token = generateToken()
	userToken := &userToken{User:user, ExpireTime:time.Now().Add(UserTokenPeriod)}
	err = userToken.store(redisPool.Get(), token)
	if err != nil {
		return
	}

	return
}

// 从仅有Id的User中读出Name, Pwd
func (u *User) ReadById() error {
	o := orm.NewOrm()
	return o.Read(u)
}

// 生成用作Token的随机串
func generateToken() string {
	token := utils.GetRandomString(16)

	isExist, _ := redis.Bool(redisPool.Get().Do("EXISTS", UserTokenPrefix + token))

	// 如果发现重复的Token，重新生成
	if isExist {
		return generateToken()
	}

	return token
}

// 根据Token取得User
func GetUserByToken(token string) (user *User, err error) {
	conn := redisPool.Get()

	userToken, err := getUserToken(conn, token)
	if err != nil {
		return
	}

	if userToken.IsExpire {
		err = fmt.Errorf("Token:\"%s\"已过期", token)
		return
	}

	if userToken.ExpireTime.Before(time.Now()) {
		userToken.IsExpire = true
		userToken.store(conn, token)
		err = fmt.Errorf("Token:\"%s\"已过期", token)
		return
	}

	user = &userToken.User

	return
}

func GetUserFromContext(ctx *context.Context) (user *User, err error) {
	token := ctx.GetCookie("token")
	user, err = GetUserByToken(token)
	if err != nil {
		ctx.SetCookie("token", "")
	}
	return
}

func ExpireToken(token string) error {
	conn := redisPool.Get()
	userToken, err := getUserToken(conn, token)
	if err != nil {
		return err
	}
	userToken.IsExpire = true
	err = userToken.store(conn, token)
	if err != nil {
		return err
	}
	return nil
}

func (ut *userToken) store(conn redis.Conn, token string) (err error) {
	tokenStr := UserTokenPrefix + token

	serialized, err := json.Marshal(ut)
	if err != nil {
		return
	}
	conn.Do("set", tokenStr, serialized)

	return
}

func getUserToken(conn redis.Conn, token string) (ut *userToken, err error) {
	tokenStr := UserTokenPrefix + token
	raw, err := conn.Do("get", tokenStr)
	if err != nil {
		return
	}
	serialized, ok := raw.([]byte)
	if !ok {
		err = fmt.Errorf("UserToken转换失败")
		return
	}

	ut = new(userToken)
	if err = json.Unmarshal(serialized, ut); err != nil {
		err = fmt.Errorf("UserToken转换失败")
		return
	}

	return
}

func checkUserValid(user *User) error {
	valid := validation.Validation{}
	b, err := valid.Valid(user)
	if err != nil {
		return err
	}
	if !b {
		for _, err := range valid.Errors {
			return fmt.Errorf("%s %s", err.Key, err.Message)
		}
	}

	return nil
}

func GetUserNameById(uid int) (string,error){
	o := orm.NewOrm()
	qs := o.QueryTable("user").Filter("id",uid)
	user := User{}
	err := qs.One(&user)
	if err != nil {
		return "",err
	}
	return user.Name,err
}