package application_test

import (
	"testing"
	"github.com/codeedu/go-hexagonal/application"
	"github.com/stretchr/testify/require"
	uuid "github.com/satori/go.uuid"
)


func TestProduct_Enable(t *testing.T){
	product := application.Product{}
	product.Name = "product"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()

	require.Equal(t, "the price must be greather than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T){
	product := application.Product{}
	product.Name = "product"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()

	require.Equal(t, "the price must be zero in order to have the product disabled", err.Error())
}

func TestProduct_IsValid(t *testing.T){
	product := application.Product{}
	product.ID = string(uuid.NewV4().String())
	product.Name = "product"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"

	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	product.Price = -1

	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal to zero", err.Error())


}



