/*
*@Time        : 2021/06/25
*@Creator     : lu.kaicheng 10261316
*@File        : adapter.go
*@Description : 适配器模式:类适配器，通过继承实现适配
 */

package adapter

import "fmt"

// 期望接口
type Taget interface {
	Create() error
}

// Adaptee 实现类
type AWS struct{}

func (aws AWS) Run() string {
	return "aws server run"
}

type Aliyun struct{}

func (aliyun Aliyun) Start(param string) {
	fmt.Println("aliyun server start", param)
}

// Adapter
type AWSAdapter struct {
	aws AWS
}

func NewAWSAdapter() *AWSAdapter {
	return &AWSAdapter{aws: AWS{}}
}

func (adapter AWSAdapter) Create() error {
	fmt.Println(adapter.aws.Run())
	return nil
}

// Adapter
type AliyunAdapter struct {
	aliyun Aliyun
}

func NewAliyunAdater() *AliyunAdapter {
	return &AliyunAdapter{aliyun: Aliyun{}}
}

func (adapter AliyunAdapter) Create() error {
	adapter.aliyun.Start("param")
	return nil
}
