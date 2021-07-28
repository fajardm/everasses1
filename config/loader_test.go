package config_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/everasses1/config"
	"github.com/stretchr/testify/assert"
)

func TestLoad_PanicError(t *testing.T) {
	res, err := config.LoadFile("xxx")

	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestLoad_ReadInConfig_Error(t *testing.T) {
	mockYML := `
	app:
		name: app
		version: v1.0
		timezone: Asia/Jakarta
		env: dev
	`
	mockPath := "./config.unittest.yml"

	err := ioutil.WriteFile(mockPath, []byte(mockYML), 0644)
	assert.NoError(t, err)

	res, err := config.LoadFile(mockPath)

	assert.Nil(t, res)
	assert.Error(t, err)

	err = os.Remove(mockPath)

	assert.NoError(t, err)
}

func TestLoad_Unmarshal_Error(t *testing.T) {
	mockYML := `
app:
 name: app
 version: v1.0
 timezone: Asia/Jakarta
 env: dev
http:
 port: xxxx`
	mockPath := "./config.unittest.yml"

	err := ioutil.WriteFile(mockPath, []byte(mockYML), 0644)
	assert.NoError(t, err)

	res, err := config.LoadFile(mockPath)

	assert.Nil(t, res)
	assert.Error(t, err)

	err = os.Remove(mockPath)
	assert.NoError(t, err)
}

func TestLoad_NoError(t *testing.T) {
	mockYML := `
app:
 name: app
 version: v1.0
 timezone: Asia/Jakarta
 env: dev`
	mockPath := "./config.unittest.yml"

	err := ioutil.WriteFile(mockPath, []byte(mockYML), 0644)
	assert.NoError(t, err)

	res, err := config.LoadFile(mockPath)

	assert.NoError(t, err)
	assert.Equal(t, "dev", res.App.Env)
	assert.Equal(t, "app", res.App.Name)
	assert.NoError(t, err)

	err = os.Remove(mockPath)
	assert.NoError(t, err)
}
