package letter

import (
	"fmt"
	"luvletter/app/mood"
	"luvletter/app/tag"
	"luvletter/app/user"
	"luvletter/custom"
	"net/http"
	"strconv"

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
	if offset != "" && position != "" {
		positionInt64, err := strconv.ParseInt(position, 10, 64)
		offsetInt64, err := strconv.ParseInt(offset, 10, 64)
		if err != nil {
			return custom.BadRequestError("querying parameters error", err)
		}
		if all, err = FindPage(positionInt64, offsetInt64); err != nil {
			return custom.BadRequestError("querying letters error", err)
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

	_, err = user.TrackUserAction(account, "create mood", "")
	if err != nil {
		return custom.HTTPTrackError(err)
	}
	if all, err = FindAll(); err != nil {
		return custom.BadRequestError("querying all letters error", err)
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
	type lengthRes struct {
		Length int64 `json:"length"`
	}
	length, err := FindNumber()
	fmt.Println(length)
	if err != nil {
		return custom.InternalServerError("querying letters number error", err)
	}

	return c.JSON(http.StatusOK, &lengthRes{length})
}
