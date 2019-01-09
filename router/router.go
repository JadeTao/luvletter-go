package router

import (
	"github.com/JadeTao/luvletter-go/app/avatar"
	"github.com/JadeTao/luvletter-go/app/mood"
	"github.com/JadeTao/luvletter-go/app/tag"

	"github.com/labstack/echo"

	"github.com/JadeTao/luvletter-go/app/letter"
	"github.com/JadeTao/luvletter-go/app/user"
)

// Prefix ...
var Prefix = "/api/v1"

// PrefixMapper ...
func PrefixMapper(router map[string]echo.HandlerFunc, prefix string) map[string]echo.HandlerFunc {
	var withPrefixRouter = make(map[string]echo.HandlerFunc)
	for key, value := range router {
		withPrefixRouter[prefix+key] = value
	}
	return withPrefixRouter
}

// GETRouters RouterConfig for GET.
var GETRouters = PrefixMapper(map[string]echo.HandlerFunc{
	APILetters:       letter.GetPage,
	APILettersLength: letter.GetLength,
	APIMoods:         mood.GetAll,
	APITags:          tag.GetAll,
	APIAvatar:        avatar.GetAvatar,
}, Prefix)

// POSTRouters RouterConfig for POST.
var POSTRouters = PrefixMapper(map[string]echo.HandlerFunc{
	APILogin:    user.Login,
	APIRegister: user.Register,
	APILetters:  letter.Save,
	APITags:     tag.Save,
	APIMoods:    mood.Save,
	APIAvatar:   avatar.UploadAvatar,
}, Prefix)
