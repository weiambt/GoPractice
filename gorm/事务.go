package main

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type ContextTx struct{}

type Data struct {
	ApolloDb *gorm.DB
}

func (d *Data) ApolloDB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(ContextTx{}).(*gorm.DB)
	if ok {
		fmt.Println("-")
		return tx
	}
	return d.ApolloDb
}

func main() {
	// 假设 ApolloDb 已经初始化
	defaultDB := &gorm.DB{}
	fmt.Println(defaultDB)

	data := Data{ApolloDb: defaultDB}

	ctxWithTx := context.WithValue(context.Background(), ContextTx{}, &gorm.DB{})
	ctxWithoutTx := context.Background()

	fmt.Println(data.ApolloDB(ctxWithTx))    // 打印/gorm.DB实例指针 (事务)
	fmt.Println(data.ApolloDB(ctxWithoutTx)) // 打印默认的ApolloDb指针 (没有事务)
}
