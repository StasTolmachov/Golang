package model_test

import (
	"testing"

	"github.com/Golang/HTTP-REST-API/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestUser_Validate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.Validate())
}

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}
