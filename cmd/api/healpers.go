package main

import (
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type envelope map[string]interface{}

func (app *application) readIDParam(r *http.Request) (int64, error) {
	// 获取 POST 请求参数
	params := httprouter.ParamsFromContext(r.Context())

	// 转换 id 类型为整型
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, header http.Header) error {
	// 将数据序列化为 JSON
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	// 遍历 Header 信息写入响应
	js = append(js, '\n')
	for key, value := range header {
		w.Header()[key] = value
	}

	// 修改为 JSON 的数据类型
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
