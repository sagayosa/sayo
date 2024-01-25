package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"io"
	"net"
	"net/http"
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

func Post(URL string, data interface{}) (code int, body []byte, err error) {
	bts, err := json.Marshal(data)
	if err != nil {
		return
	}

	resp, err := http.Post(URL, "application/json", bytes.NewBuffer(bts))
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return resp.StatusCode, body, nil
}

func UnMarshalUnknownAny(source interface{}, dest interface{}) error {
	bts, err := json.Marshal(source)
	if err != nil {
		return err
	}

	return json.Unmarshal(bts, dest)
}

func ChangeRoutineWorkDir(workDir string) error {
	return os.Chdir(workDir)
}

// could there be a better way?
func GetAvailablePort() (int, error) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, err
	}
	defer listener.Close()
	addr := listener.Addr().(*net.TCPAddr)
	return addr.Port, nil
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
