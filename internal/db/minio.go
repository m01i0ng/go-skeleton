package db

import (
  "context"

  "github.com/google/wire"
  "github.com/kataras/golog"
  "github.com/m01i0ng/go-skeleton/internal/config"
  "github.com/minio/minio-go/v7"
  "github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioProvider = wire.NewSet(NewMinio)

func NewMinio(c *config.Config) (*minio.Client, error) {
  m := c.Minio
  client, err := minio.New(m.Endpoint, &minio.Options{
    Creds:  credentials.NewStaticV4(m.AccessKey, m.SecretKey, ""),
    Secure: m.Tls,
  })

  if err != nil {
    return nil, err
  }

  ctx := context.Background()

  err = client.MakeBucket(ctx, m.Bucket, minio.MakeBucketOptions{})
  if err != nil {
    bucketExists, err := client.BucketExists(ctx, m.Bucket)
    if err == nil && bucketExists {
      golog.Infof("bucket exists: %s", m.Bucket)
    } else {
      return nil, err
    }
  } else {
    golog.Infof("bucket make ok: %s", m.Bucket)
  }

  return client, err
}
