package notify

import (
	"testing"
	"xx/forms"
)

//func TestSetChan(t *testing.T) {
//	form := &forms.UserStatus{
//
//	}
//	if err := SetChan(form); err != nil {
//		t.Errorf(err.Error())
//	}
//}

func TestSetChan(t *testing.T) {
	// 创建一个 UserStatus 对象
	status := &forms.UserStatus{
		Timestamp: 1600060847294,
		User:      "t_13434@101961.com/android_b069b852-79a3-3c9e-9d08-ee5176b95df5",
		Status:    "online",
	}

	// 调用 SetChan 函数
	err := SetChan(status)
	if err != nil {
		t.Errorf("SetChan returned an error: %v", err)
	}
}
