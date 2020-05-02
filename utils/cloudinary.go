package utils

import (
	"bytes"
	"log"

	"golang.org/x/net/context"

	"github.com/kyokomi/cloudinary"
)

type CloudynaryInfo struct {
	FilePath string
	Err      error
}

var (
	CtxCloudinary  = NewCloudinary()
	CloudinaryAuth = "cloudinary://365527915797683:u_kri0We3qcCmD0ojDkU9GhPetw@yers"
	path           = "user_images/"
)

func NewCloudinary() context.Context {
	ctx := context.Background()
	ctxCloud := cloudinary.NewContext(ctx, CloudinaryAuth)
	return ctxCloud
}

func UploadImage(fileName string, image []byte) chan CloudynaryInfo {
	ctx := context.Background()
	ctx = cloudinary.NewContext(ctx, "cloudinary://365527915797683:u_kri0We3qcCmD0ojDkU9GhPetw@yersx")

	cloudinary.UploadStaticImage(ctx, path+fileName, bytes.NewBuffer(image))

	log.Printf("creation succeded")

	chanInfo := make(chan CloudynaryInfo)
	go func() {
		err := cloudinary.UploadStaticImage(ctx, path+fileName, bytes.NewBuffer(image))
		chanInfo <- CloudynaryInfo{cloudinary.ResourceURL(ctx, path+fileName), err}
	}()
	return chanInfo
}
