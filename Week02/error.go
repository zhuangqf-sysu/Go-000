package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

// model层 自定义错误类型
type MyError struct {
	Code int
	Message string
}

func (err *MyError) Error() string  {
	return err.Message
}

// 包装方法生成http或者grpc错误信息，即对外错误
func (err *MyError) ToHttpError() int  {
	return err.Code
}

// 定义内部错误
var (
	ErrNotFound = &MyError{Code: 404, Message: "not found"}
	ErrDBServer = &MyError{Code: 500, Message: "db server err"}
	ErrRequest  = &MyError{Code: 400, Message: "参数错误"}
	// ...
)

// 自定义对象
type MyObject struct {

}

// Dao层
// GetMyObjectByID 通过id从数据库获取MyObject
func GetMyObjectByID(ctx context.Context, id int64) (*MyObject, error) {
	// 执行数据库查询
	selectSql := ""
	obj, err := db.query(ctx, selectSql, id)
	if err == sql.ErrNoRows {
		return nil, errors.Wrap(ErrNotFound, fmt.Sprintf("GetMyObjectByID(%d) not found", id))
	}
	if err != nil {
		return nil, errors.Wrap(ErrDBServer, fmt.Sprintf("GetMyObjectByID(%d) sql(%s) err(%v)", id, selectSql, err))
	}
	// 正确, 取数据
	return &MyObject{}, nil
}

// Service层
func GetMyObject(ctx context.Context, id int64) (*MyObject, error)  {
	// 处理业务, 调用dao层接口
	return GetMyObjectByID(ctx, id)
}

// Controller层接口
func GetMyObject(ctx context.Context, params map[string]interface{}) (data interface{}, err error) {
	// 执行鉴权、取参， 参数校验等操作
	// 假设取到参数id
	id := int64(1)
	obj, err := GetMyObject(ctx, id)
	if errors.Is(err, ErrNotFound) {
		// 按外部协议返回默认值，或404错误
	}

	// 按外部协议返回错误或data
	return obj, err
}






