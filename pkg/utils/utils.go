package utils

import (
	"crypto/sha256"
	"encoding/json"
	"io"
	"os"
	baseresp "sayo_framework/pkg/base_resp"
	sayolog "sayo_framework/pkg/sayo_log"
	"strings"

	"github.com/kataras/iris/v12"
)

func StringPlus(segments ...string) string {
	var builder strings.Builder

	for _, seg := range segments {
		builder.WriteString(seg)
	}

	return builder.String()
}

func SHA256(filePath string) (res string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	hash := sha256.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return
	}

	return string(hash.Sum(nil)), nil
}

func JSON(filePath string, dst interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	bts, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bts, dst); err != nil {
		return err
	}

	return nil
}

type HandlerFunc func(iris.Context)

func IrisCtxJSONWrap(f func(ctx iris.Context) (*baseresp.BaseResp, error)) HandlerFunc {
	return func(ctx iris.Context) {
		resp, err := f(ctx)
		if err != nil {
			sayolog.Err(err).Error()
		}
		ctx.JSON(resp)
	}
}
