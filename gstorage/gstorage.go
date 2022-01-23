package gastorage

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"time"

	"cloud.google.com/go/storage"
	"github.com/labstack/echo"
	"google.golang.org/api/option"
)

func UploadFile(bucket string, object string, imgBase64 string) error {
	// bucket := "bucket-name"  storageのバケット名
	// object := "object-neme"   アップロード後のファイル名、自分で決める
	credentialFilePath := "./gcs-sdk.json"
	ctx := context.Background()

	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFilePath))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "cannot set up gstorage client",
		}
	}
	defer client.Close()

	decodedImage, err := base64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "imagebase64 cannot decode",
		}
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	decodedReader := bytes.NewReader(decodedImage)
	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)

	if _, err = io.Copy(wc, decodedReader); err != nil {
		fmt.Errorf("io.Copy:%v", err)
	}
	if err := wc.Close(); err != nil {
		fmt.Errorf("wc.Close:%v", err)
	}

	fmt.Printf("Blob %v uploaded \n", object)

	return nil
}

func DeleteFile(bucket string, object string) error {
	credentialFilePath := "./gcs-sdk.json"
	ctx := context.Background()

	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFilePath))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "cannot set up gstorage client",
		}
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	o := client.Bucket(bucket).Object(object)
	if err := o.Delete(ctx); err != nil {
		return fmt.Errorf("object(%q) .delere: %v", object, err)
	}
	fmt.Printf("Blob %v delete \n", object)

	return nil
}
