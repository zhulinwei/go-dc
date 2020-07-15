package util

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"gopkg.in/go-playground/validator.v8"

	"github.com/zhulinwei/go-dc/pkg/model"
	"gopkg.in/yaml.v3"
)

func CommonContent() (ctx context.Context) {
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}

func IsPathExists(path string) (bool, error) {
	// 若返回的错误为nil,说明文件或文件夹存在
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	// 若IsNotExist判断为true,说明文件或文件夹不存在
	if os.IsNotExist(err) {
		return false, nil
	}
	// 若返回的错误为其它类型,则不确定是否在存在，可视为不存在
	return false, err
}

func ParseServerConfig(filePath string) (model.ServerConfig, error) {
	var err error
	var fileBytes []byte
	var serverConfig model.ServerConfig
	if fileBytes, err = ioutil.ReadFile(filePath); err != nil {
		return serverConfig, err
	}
	if err = yaml.Unmarshal(fileBytes, &serverConfig); err != nil {
		return serverConfig, err
	}
	return serverConfig, nil
}

func ParserErrorMsg(err error) string {
	// unmarshal type error
	if unmarshalTypeErr, ok := err.(*json.UnmarshalTypeError); ok {
		return fmt.Sprintf("unmarshal type fail, field '%s' type must be '%s', not '%s'", unmarshalTypeErr.Field, unmarshalTypeErr.Type, unmarshalTypeErr.Value)
	}

	// validator validation error
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, validationErr := range validationErrs {
			fmt.Println(validationErr.Param)
			fmt.Println(validationErr.Value)
			fmt.Println(validationErr.Name)
			fmt.Println(validationErr.Kind)
			fmt.Println(validationErr.Field)
			switch validationErr.Tag {
			case "required":
				return fmt.Sprintf("field '%s' was required, can not be empty.\n", strings.ToLower(validationErr.Name))
			case "min":
				return fmt.Sprintf("field '%s' must be greater than %s.\n", strings.ToLower(validationErr.Name), validationErr.Param)
			case "max":
				return fmt.Sprintf("field '%s' must be less than %s.\n", strings.ToLower(validationErr.Name), validationErr.Param)
			}
		}
	}
	return err.Error()
}
