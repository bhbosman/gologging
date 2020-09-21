package gologging_test

import (
	"github.com/bhbosman/gologging"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestLoggingLogLevel(t *testing.T) {
	tempMap := make(map[string]*gologging.SubSystemLogger)
	fac := gologging.NewFactory(log.New(os.Stdout, "", 0))
	subSystemLoggerFromFactory := fac.Create("ddd")
	tempMap["ddd"] = subSystemLoggerFromFactory
	var memoryClone gologging.SubSystemLogger = *subSystemLoggerFromFactory
	fac.SetLogLevel("ddd", 12)
	assert.Equal(t, 12, tempMap["ddd"].GetLogLevel())
	assert.Equal(t, 12, subSystemLoggerFromFactory.GetLogLevel())
	assert.Equal(t, 12, memoryClone.GetLogLevel())

}
