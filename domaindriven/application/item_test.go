package application_test

import (
	"github.com/stretchr/testify/assert"
	"magical.dev/tesing/domaindriven/application"
	"testing"
)

func TestItemApplicationService_CreateItem(t *testing.T) {
	itemService := application.ItemApplicationService{}

	itemService.CreateItem()

	assert.Len(t, itemService.GetAllItems(), 1)
}
