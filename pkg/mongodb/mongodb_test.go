package mongodb

import (
	"context"
	"testing"
	"github.com/stretchr/testify/assert"

)

func TestGetMongoClient(t *testing.T) {
	client := GetMongoClient()
	err := client.Ping(context.TODO(), nil)
	assert.Equal(t, err, nil, "client ping fail")
}
