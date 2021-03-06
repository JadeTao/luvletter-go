package letter

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/JadeTao/luvletter-go/app/mood"
	"github.com/JadeTao/luvletter-go/app/tag"
	"github.com/JadeTao/luvletter-go/app/user"
	"github.com/JadeTao/luvletter-go/conf"
	"github.com/JadeTao/luvletter-go/custom"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// GetPage ...
func GetPage(c echo.Context) error {
	var (
		all []Letter
		err error
	)
	params := c.QueryParams()

	offset := params.Get("offset")
	position := params.Get("position")

	userInfo := c.Get("user").(*jwt.Token)
	claims := userInfo.Claims.(*user.JwtCustomClaims)
	account := claims.Account

	if offset == "" {
		return custom.BadRequestError("querying parameters error", err)
	}
	if position == "" {
		return custom.BadRequestError("querying parameters error", err)
	}
	if offset != "" && position != "" {
		positionInt64, err := strconv.ParseInt(position, 10, 64)
		offsetInt64, err := strconv.ParseInt(offset, 10, 64)
		if err != nil {
			return custom.BadRequestError("querying parameters error", err)
		}
		if all, err = FindPage(positionInt64, offsetInt64); err != nil {
			return custom.InternalServerError("querying letters error", err)
		}
		_, err = user.TrackUserAction(account, "query letter", fmt.Sprintf("查询从%s到%s的letter", position, offset))
		if err != nil {
			return custom.HTTPTrackError(err)
		}
		return c.JSON(http.StatusOK, all)
	}
	if all, err = FindAll(); err != nil {
		return custom.BadRequestError("querying letters error", err)
	}
	return c.JSON(http.StatusOK, all)
}

// GetAll ...
func GetAll(c echo.Context) error {
	var (
		all []Letter
		err error
	)
	userInfo := c.Get("user").(*jwt.Token)
	claims := userInfo.Claims.(*user.JwtCustomClaims)
	account := claims.Account

	_, err = user.TrackUserAction(account, "get letters", "")
	if err != nil {
		return custom.HTTPTrackError(err)
	}
	if all, err = FindAll(); err != nil {
		return custom.BadRequestError("get all letters error", err)
	}
	return c.JSON(http.StatusOK, all)
}

// Save ...
func Save(c echo.Context) error {
	var (
		l     Letter
		trace user.TrackAction
		err   error
	)

	if err = c.Bind(&l); err != nil {
		return custom.BadRequestError("binding parameters error", err)
	}

	trace, err = user.TrackUserAction(l.Account, "create letter", "")

	l.CreateTime = trace.Time

	err = SaveLetter(&l)
	if err != nil {
		return custom.InternalServerError("saving letter error", err)
	}

	// mood、tag计数
	_ = tag.AddCountInBatch(l.Tags)
	_ = mood.AddCount(l.Mood)

	return c.JSON(http.StatusOK, l)
}

// GetLength ...
func GetLength(c echo.Context) error {
	userInfo := c.Get("user").(*jwt.Token)
	claims := userInfo.Claims.(*user.JwtCustomClaims)
	account := claims.Account

	_, err := user.TrackUserAction(account, "get the number of letter", "")
	if err != nil {
		return custom.HTTPTrackError(err)
	}
	type result struct {
		Size   int64 `json:"size"`
		Number int64 `json:"number"`
	}
	length, err := FindNumber()

	res := &result{
		Size:   conf.Conf.Letter.Size,
		Number: length/conf.Conf.Letter.Size + 1,
	}
	if length%conf.Conf.Letter.Size == 0 {
		res.Number = length / conf.Conf.Letter.Size
	}
	if err != nil {
		return custom.InternalServerError("querying letters number error", err)
	}

	return c.JSON(http.StatusOK, res)
}
