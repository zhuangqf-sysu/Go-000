答：

dao层中产生的sql.ErrNoRows，是查询不到数据，严格意义上不属于错误或者异常，是正常的业务结果。

在处理上，我更喜欢用自定义的错误包装NotFound结果，warp之后将这个错误抛给上层业务处理

上层业务（service or controller） 可以根据业务需求，使用默认值或者返回错误，灵动

具体操作见error.go
