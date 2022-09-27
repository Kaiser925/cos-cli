package oss

import (
	"io/fs"
	"net/http"
	"net/url"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
)

const timeout = 10 * time.Second

type COS struct {
	cli *cos.Client
}

func NewCOS(URL string, secretID string, secretKey string) *COS {
	u, _ := url.Parse(URL)
	b := &cos.BaseURL{BucketURL: u}
	return &COS{
		cli: cos.NewClient(b, &http.Client{
			Timeout: timeout,
			Transport: &cos.AuthorizationTransport{
				SecretID:  secretID,
				SecretKey: secretKey,
			},
		}),
	}
}

func (c *COS) Open(name string) (fs.File, error) {
	// TODO implement me
	panic("implement me")
}
