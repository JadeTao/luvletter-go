package avatar

import (
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"luvletter/conf"
	"luvletter/util"
	"os"

	"github.com/chai2010/webp"
)

// SaveAvatar ...
func SaveAvatar() {

}

// Png2Webp ...
func Png2Webp(source io.Reader, target io.Writer) error {
	img, err := png.Decode(source)
	if err != nil {
		return err
	}
	err = webp.Encode(target, img, &webp.Options{Lossless: true})
	return err
}

// Jpg2Webp ...
func Jpg2Webp(source io.Reader, target io.Writer) error {
	img, err := jpeg.Decode(source)
	if err != nil {
		return err
	}
	err = webp.Encode(target, img, &webp.Options{Lossless: true})
	return err
}

// UniqUserAvatar ...
func UniqUserAvatar(account string) error {
	files, err := ioutil.ReadDir(conf.Conf.Assets.Avatar)
	if err != nil {
		return err
	}

	for _, file := range files {
		if util.GetFileWithoutSuffix(file.Name()) == account {
			if err := os.Remove(conf.Conf.Assets.Avatar + file.Name()); err != nil {
				return err
			}
		}
	}
	return err
}

// GetAccountAvatarName ...
func GetAccountAvatarName(account string) (string, error) {
	files, err := ioutil.ReadDir(conf.Conf.Assets.Avatar)

	if err != nil {
		return "", err
	}

	for _, file := range files {
		if util.GetFileWithoutSuffix(file.Name()) == account {
			return file.Name(), nil
		}
	}
	return "", err
}
