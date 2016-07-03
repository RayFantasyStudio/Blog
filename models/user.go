package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/RayFantasyStudio/blog/utils"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego/context"
)

// 登陆有效期60天         day  hour  min  sec
const UserTokenPeriod_s = 60 * 24 * 60 * 60

type User struct {
	Id   int `form:"_" json:"id" redis:"id"`
	Name string `orm:"size(16);unique" form:"name" json:"name" valid:"Required" redis:"name"`
	Pwd  string `orm:"size(32)" form:"pwd" json:"_" valid:"MinSize(8);MaxSize(32)" redis:"pwd"`
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
	if err = storeUserToToken(&user, token); err != nil {
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
	user, err = getUserFromRedis(token)
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

func DelUserToken(token string) error {
	tokenStr := UserTokenPrefix + token
	conn := redisPool.Get()
	_, err := conn.Do("DEL", tokenStr)
	return err
}

func storeUserToToken(u *User, token string) (err error) {
	conn := redisPool.Get()
	tokenStr := UserTokenPrefix + token

	if _, err = conn.Do("HMSET", redis.Args{}.Add(tokenStr).AddFlat(*u)...); err != nil {
		return
	}

	if _, err = conn.Do("EXPIRE", tokenStr, UserTokenPeriod_s); err != nil {
		return
	}

	return
}

func getUserFromRedis(token string) (u *User, err error) {
	conn := redisPool.Get()
	tokenStr := UserTokenPrefix + token

	u = new(User)

	values, err := redis.Values(conn.Do("HGETALL", tokenStr))
	if err != nil {
		return
	}

	redis.ScanStruct(values, u)

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

func GetUserNameById(uid int) (string, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("user").Filter("id", uid)
	user := User{}
	err := qs.One(&user)
	if err != nil {
		return "", err
	}
	return user.Name, err
}
func ModifyUserInfo(uid int64,username string,pwd string) error{
	o := orm.NewOrm()
	qs := o.QueryTable("user").Filter("Id",uid)
	var user User
	err := qs.One(&user)
	if err != nil {
		return err
	}
	user.Name = username
	user.Pwd = pwd
	o.Update(&user)
	return err
}