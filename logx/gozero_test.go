package logx

import (
	"testing"

	zerologx "github.com/zeromicro/go-zero/core/logx"

	"github.com/stretchr/testify/assert"
	"github.com/wwqdrh/logger"
)

func TestGozeroLog(t *testing.T) {
	l, err := NewZeroWriter(logger.DefaultLogger.Logger)
	assert.Nil(t, err)

	zerologx.Must(err)
	zerologx.SetWriter(l)
	zerologx.Info("test hhh")
}
