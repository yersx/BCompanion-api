package utils

import (
	"bytes"

	"github.com/kyokomi/cloudinary"
	"golang.org/x/net/context"
)

type CloudynaryInfo struct {
	FilePath string
	Err      error
}

var (
	CtxCloudinary  = NewCloudinary()
	CloudinaryAuth = "cloudinary://YOUR_KEY@login"
	path           = "user_images/"
)

func NewCloudinary() context.Context {
	ctx := context.Background()
	ctxCloud := cloudinary.NewContext(ctx, CloudinaryAuth)
	return ctxCloud
}

func UploadImage(nameFile string, buff []byte) chan CloudynaryInfo {
	readFileCopied := bytes.NewBuffer(buff)
	chanInfo := make(chan CloudynaryInfo)
	go func() {
		err := cloudinary.UploadStaticImage(CtxCloudinary, path+nameFile, readFileCopied)
		chanInfo <- CloudynaryInfo{cloudinary.ResourceURL(CtxCloudinary, path+nameFile), err}
	}()
	return chanInfo
}
