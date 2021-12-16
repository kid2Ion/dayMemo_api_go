package gastorage

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
)

func UploadFile(w io.Writer, bucket, object string) error {
	// bucket := "bucket-name"  storageのバケット名
	// object := "object-neme"   アップロード後のファイル名

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Errorf("NewClient:%v", err)
	}
	defer client.Close()

	decodedImage, err := base64.RawStdEncoding.DecodeString("aaa")
	if err != nil {
		fmt.Errorf("decode:%v", err)
	}
	decodedReader := bytes.NewReader(decodedImage)

	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)
	if _, err = io.Copy(wc, decodedReader); err != nil {
		fmt.Errorf("io.Copy:%v", err)
	}
	if err := wc.Close(); err != nil {
		fmt.Errorf("wc.Close:%v", err)
	}

	fmt.Fprintf(w, "Blob %v uploaded \n", object)
	return nil
}
