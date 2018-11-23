package avatar

import (
	"io"
	"luvletter/app/user"
	"luvletter/conf"
	"luvletter/custom"
	"net/http"
	"os"
	"path"

	"github.com/labstack/echo"
)

// GetAvatar ...
func GetAvatar(c echo.Context) error {
	account := c.Param("account")

	accountAvatarName, err := GetAccountAvatarName(account)
	if err != nil {
		return err
	}
	return c.File(conf.Conf.Assets.Avatar + accountAvatarName)
}

// UploadAvatar ...
func UploadAvatar(c echo.Context) error {
	//-----------
	// Read file
	//-----------

	account := c.Param("account")
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	UniqUserAvatar(account)

	// Destination
	dst, err := os.Create(conf.Conf.Assets.Avatar + account + path.Ext(file.Filename))
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	if _, err = user.TrackUserAction(account, "upload avatar", ""); err != nil {
		return custom.HTTPTrackError(err)
	}

	return c.JSON(http.StatusOK, nil)
}
