package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtil(t *testing.T) {

	assert.Equal(t, "util", util())

}

func TestFeature1(t *testing.T) {

	assert.Equal(t, "feature1", feature1())
}
