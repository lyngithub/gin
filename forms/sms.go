package forms

type SendSmsForm struct {
	Mobile string `form:"mobile" json:"mobile" validate:"required,mobile"` //手机号码格式有规范可寻， 自定义validator
	Type   uint   `form:"type" json:"type" validate:"required,oneof=1 2"`
}
