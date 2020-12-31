package util

import (
	"encoding/json"
	"github.com/scott-x/email/model"
	"io/ioutil"
)

func ParseConfig() (*model.EmailParam,error){
	bs, err:=ioutil.ReadFile("config.json")
	if err!=nil {
		return nil,err
	}
	ep := &model.EmailParam{}
	err=json.Unmarshal(bs,ep)
	if err!=nil {
		return nil, err
	}
	return ep, nil
}