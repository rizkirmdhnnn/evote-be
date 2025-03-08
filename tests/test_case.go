package tests

import (
	"github.com/goravel/framework/testing"

	"evote-be/bootstrap"
)

func init() {
	bootstrap.Boot()
}

type TestCase struct {
	testing.TestCase
}
