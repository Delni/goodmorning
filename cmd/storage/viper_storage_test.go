package storage

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var vSM = ViperStorageManager{}

func TestSaveApp(t *testing.T) {
	viper.AddConfigPath(".")
	vSM.saveApp(App{
		label:   "Test",
		command: "test",
		active:  true,
	})

	assert.NotNil(t, viper.Get("apps.test"))
	assert.Equal(t, "Test", viper.GetString("apps.test.label"))
	assert.Equal(t, "test", viper.GetString("apps.test.command"))
	assert.True(t, viper.GetBool("apps.test.active"))
}

func TestParseApps(t *testing.T) {
	viper.AddConfigPath(".")
	vSM.saveApp(App{
		label:   "Test",
		command: "test",
		active:  true,
	})

	apps := vSM.getApps()

	assert.Len(t, apps, 1)
	testApp := apps[0]

	assert.Equal(t, App{
		label:   "Test",
		command: "test",
		active:  true,
	}, testApp)
}
