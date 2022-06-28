package logx

import (
	"testing"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/stretchr/testify/assert"
	"github.com/wwqdrh/logger"
)

func TestGozeroLog(t *testing.T) {
	l, err := NewZeroWriter(logger.DefaultLogger)
	assert.Nil(t, err)

	logx.Must(err)
	logx.SetWriter(l)
	logx.Info("test hhh")
}
