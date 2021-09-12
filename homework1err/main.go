package main

import (
	"fmt"

	"github.com/pkg/errors"
)


var ErrRecordNotFound = errors.New("record not found")
var sqlNoRowsErr = errors.New("sql.ErrNoRows")


func main()  {
	err := findUserLogic()
	if err != nil {
		fmt.Println(fmt.Sprintf("错误：%+v", err))
	}
}


func findUser() error {

	// 模拟而已别当真
	virErr := sqlNoRowsErr
	if errors.Is(virErr, sqlNoRowsErr) {
		return ErrRecordNotFound
	}
	return virErr
}

func findUserLogic() error {
	// do logic...
	err := findUser()
	if err != nil && errors.Is(err, ErrRecordNotFound)  {
		return errors.Wrapf(err, "user not found", )
	}
	return err
}
