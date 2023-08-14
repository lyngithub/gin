package forms

type PassWordLoginForm struct {
	Mobile    string `form:"mobile" json:"mobile" validate:"required,mobile"` //手机号码格式有规范可寻， 自定义validator
	PassWord  string `form:"password" json:"password" validate:"required,min=3,max=20"`
	Captcha   string `form:"captcha" json:"captcha" validate:"required,min=5,max=5"`
	CaptchaId string `form:"captcha_id" json:"captcha_id" validate:"required"`
}

type RegisterForm struct {
	Mobile   string `form:"mobile" json:"mobile" validate:"required,mobile"` //手机号码格式有规范可寻， 自定义validator
	PassWord string `form:"password" json:"password" validate:"required,min=3,max=20"`
	//Code     string `form:"code" json:"code" binding:"required,min=6,max=6"`
}
