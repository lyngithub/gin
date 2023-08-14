package notify

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"xx/app/api"
	"xx/app/domain/notify"
	"xx/forms"
)

func Status(ctx *gin.Context) {
	fmt.Println(44)

	form := &forms.UserStatus{}
	fmt.Println(55)
	if err := ctx.ShouldBindJSON(form); err != nil {
		api.BadRequest(ctx, err.Error())
		panic(err.Error())
	}

	fmt.Printf("%#v", form)
	if err := notify.SetChan(form); err != nil {
		api.InternalServerError(ctx, 1)
	}

	api.Ok(ctx, 1)
}
