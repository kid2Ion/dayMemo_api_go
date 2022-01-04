package gastorage

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/labstack/echo"
	"google.golang.org/api/option"
)

func UploadFile(bucket string, object string, imgBase64 string) error {
	// bucket := "bucket-name"  storageのバケット名
	// object := "object-neme"   アップロード後のファイル名、自分で決める
	credentialFilePath := "./daymemo-d5cdab55b3ae.json"
	ctx := context.Background()

	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFilePath))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "cannot set up gstorage client",
		}
	}
	defer client.Close()

	// decodeで[]byte型になる
	decodedImage, err := base64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "imagebase64 cannot decode",
		}
	}

	// io.reader型にする
	decodedReader := bytes.NewReader(decodedImage)
	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)

	// bucket,objectを指定し、reader型の画像をアップロード
	if _, err = io.Copy(wc, decodedReader); err != nil {
		fmt.Errorf("io.Copy:%v", err)
	}
	if err := wc.Close(); err != nil {
		fmt.Errorf("wc.Close:%v", err)
	}

	fmt.Printf("Blob %v uploaded \n", object)

	// if err := makePublic(bucket, object); err != nil {
	// 	return err
	// }

	return nil
}

// もし公開でなくするするなら以下の実装が必要になる

// func makePublic(bucket string, object string) error {
// 	credentialFilePath := "./daymemo-d5cdab55b3ae.json"
// 	ctx := context.Background()
// 	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFilePath))
// 	if err != nil {
// 		return &echo.HTTPError{
// 			Code:    http.StatusBadRequest,
// 			Message: "cannot set up gstorage client",
// 		}
// 	}
// 	defer client.Close()

// 	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
// 	defer cancel()
// 	acl := client.Bucket(bucket).Object(object).ACL()
// 	if err := acl.Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
// 		fmt.Println("nityaa")
// 		return fmt.Errorf("ACLHandle.Set: %v", err)
// 	}

// 	fmt.Printf("Blob %v is now publicly accessible.\n", object)

// 	return nil
// }
