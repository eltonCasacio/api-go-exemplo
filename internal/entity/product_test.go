package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("produto 1", 10.0)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "produto 1", p.Name)
	assert.Equal(t, 10.0, p.Price)
}
func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 10)
	assert.Nil(t, p)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("produto 1", 0)
	assert.Nil(t, p)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWhenInvalidPrice(t *testing.T) {
	p, err := NewProduct("produto 1", -10)
	assert.Nil(t, p)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestProductValidate(t *testing.T) {
	p, err := NewProduct("produto 1", 10)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
	assert.NotNil(t, p.ID)
}

func TestProductValidate_WithoutID(t *testing.T) {
	p := Product{}
	err := p.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrIDIsRequired)
}
