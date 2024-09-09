package application_test

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	p := &application.Product{
		Price: 10,
	}
	err := p.Enable()
	require.Nil(t, err)
}