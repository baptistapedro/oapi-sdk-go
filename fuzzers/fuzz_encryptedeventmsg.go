package myfuzz

import (
	"github.com/larksuite/oapi-sdk-go/v3/core"
	"context"
)

func Fuzz(data []byte) int {
	buf := string(data)
	_, err := larkcore.EncryptedEventMsg(context.Background(), map[string]string{"key1": buf}, buf)
	if err != nil {
		return 1
	}
	return 0

}
