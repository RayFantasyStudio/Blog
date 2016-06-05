package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/RayFantasyStudio/blog/utils"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/astaxie/beego/context"
)

var accessToken = make(map[string]int)

type User struct {
	Id   int `form:"_" json:"id"`
	Name string `orm:"size(16);unique" form:"name" json:"name" valid:"Required"`
	Pwd  string `orm:"size(32)" form:"pwd" json:"_" valid:"MinSize(8);MaxSize(32)"`
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
	accessToken[token] = user.Id

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

	// 如果发现重复的Token，重新生成
	if accessToken[token] != 0 {
		return generateToken()
	}

	return token
}

// 根据Token取得Uid
func GetUidByToken(token string) (uid int, err error) {
	uid = accessToken[token]
	if uid == 0 {
		err = fmt.Errorf("token: \"%s\"无效", token)
		return
	}
	return
}

// 根据Token获取一个仅有Id的User
func GetEmptyUserByToken(token string) (eUser *User, err error) {
	var uid int
	if uid, err = GetUidByToken(token); err != nil {
		return
	}

	eUser = &User{Id:uid}
	return
}

// 根据Token取得User
func GetUserByToken(token string) (user *User, err error) {
	user, err = GetEmptyUserByToken(token)
	if err != nil {
		return
	}

	if err = user.ReadById(); err != nil {
		return
	}

	return
}

func GetEmptyUserFromContext(ctx *context.Context) (eUser *User, err error) {
	token := ctx.GetCookie("token")
	eUser, err = GetEmptyUserByToken(token)
	return
}

func GetUserFromContext(ctx *context.Context) (user *User, err error) {
	token := ctx.GetCookie("token")
	user, err = GetUserByToken(token)
	return
}

func GetUidFromContext(ctx *context.Context) (uid int, err error) {
	token := ctx.GetCookie("token")
	uid, err = GetUidByToken(token)
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
