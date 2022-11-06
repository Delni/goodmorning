package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlug(t *testing.T) {
	app1 := App{
		label: "Test",
	}
	app2 := App{
		label: "Test Espace",
	}
	app3 := App{
		label: "Test Trimmed ",
	}

	assert.Equal(t, "test", app1.slug())
	assert.Equal(t, "test_espace", app2.slug())
	assert.Equal(t, "test_trimmed", app3.slug())
}
