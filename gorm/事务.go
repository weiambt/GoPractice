package main

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type TxTest struct {
	ApolloDb *gorm.DB
}

type ContextTx struct{}

func (d *TxTest) ApolloDB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(ContextTx{}).(*gorm.DB)
	if ok {
		fmt.Println("-")
		return tx
	}
	return d.ApolloDb
}

func main() {
	defaultDB := &gorm.DB{}
	data := TxTest{ApolloDb: defaultDB}

	fmt.Println(defaultDB)
	ctxWithoutTx := context.Background()
	ctxWithTx := context.WithValue(context.Background(), ContextTx{}, &gorm.DB{})

	fmt.Println(data.ApolloDB(ctxWithoutTx)) // 打印默认的ApolloDb指针 (没有事务)
	fmt.Println(data.ApolloDB(ctxWithTx))    // 打印/gorm.DB实例指针 (事务)

}
