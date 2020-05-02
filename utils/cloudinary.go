package utils

import (
	"bytes"

	"golang.org/x/net/context"

	"github.com/kyokomi/cloudinary"
)

type CloudynaryInfo struct {
	FilePath string
	Err      error
}

var (
	CtxCloudinary  = NewCloudinary()
	CloudinaryAuth = "cloudinary://365527915797683:u_kri0We3qcCmD0ojDkU9GhPetw@yersx"
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
		err := cloudinary.UploadStaticImage(CtxCloudinary, nameFile, readFileCopied)
		chanInfo <- CloudynaryInfo{cloudinary.ResourceURL(CtxCloudinary, nameFile), err}
	}()
	return chanInfo
}
