package api

import (
	"os"
	"regexp"

	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/router/result"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mPath"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/gofiber/fiber/v2"
)

type UploadImgType struct {
	Url string
}

func Upload(c *fiber.Ctx) error {
	// 授权验证

	headerPath := c.Get("File-Path")

	if len(headerPath) > 0 {
		isRule := StrAllLetter(headerPath)
		if !isRule {
			return c.JSON(result.ErrUpload.With("File-Path,不符合规范！", headerPath))
		}
	} else {
		headerPath = "files"
	}

	savePath := mStr.Join(
		config.Dir.App,
		mStr.ToStr(os.PathSeparator),
		"cache",
		mStr.ToStr(os.PathSeparator),
		headerPath,
	)

	// 目录不存在则创建
	isSavePath := mPath.Exists(savePath)
	if !isSavePath {
		err := os.MkdirAll(savePath, os.FileMode(0o777))
		if err != nil {
			return c.JSON(result.ErrUpload.With("创建目录失败", err))
		}
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(result.ErrUpload.WithData(mStr.ToStr(err)))
	}

	fileName := mFile.GetName(mFile.GetNameOpt{
		FileName: file.Filename,
		SavePath: savePath,
	})

	filePath := mStr.Join(
		savePath,
		mStr.ToStr(os.PathSeparator),
		fileName,
	)

	err = c.SaveFile(file, filePath)
	if err != nil {
		return c.JSON(result.ErrUpload.WithData(mStr.ToStr(err)))
	}

	isFilePath := mPath.Exists(filePath)

	if !isFilePath {
		return c.JSON(result.ErrUpload.WithData("文件保存失败"))
	}

	imgUrl := filePath

	return c.JSON(result.Succeed.WithData(UploadImgType{
		Url: imgUrl,
	}))
}

func StrAllLetter(str string) bool {
	match, _ := regexp.MatchString(`^[0-9a-z/]+$`, str)
	return match
}
