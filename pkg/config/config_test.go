package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfigFromFile(t *testing.T) {
	// Set up a test YAML file
	tempFile, err := os.CreateTemp(".", "example-app-test-config")
	assert.NoError(t, err)
	defer tempFile.Close()
	defer os.Remove(tempFile.Name())
	cfgFileName := tempFile.Name()

	_, err = tempFile.Write([]byte(`
jdbc_url: "jdbc:postgresql://localhost:5432/testdb"
username: "testuser"
password: "testpassword"
`))
	assert.NoError(t, err)

	// Load the config from the file
	LoadConfig(&cfgFileName)

	// Assert the config values
	config := GetConfig()
	assert.NotNil(t, config)
	assert.Equal(t, "postgresql", config.DatasourceUrl.Scheme)
	assert.Equal(t, "localhost:5432", config.DatasourceUrl.Host)
	assert.Equal(t, "/testdb", config.DatasourceUrl.Path)
	assert.Equal(t, "testuser", config.DatasourceUrl.User.Username())
	password, ok := config.DatasourceUrl.User.Password()
	assert.True(t, ok)
	assert.Equal(t, "testpassword", password)
}

func TestLoadConfigFromEnv(t *testing.T) {
	// Set the DATABASE_URL environment variable
	os.Setenv(dbUrlEnv, "jdbc:postgresql://testuser:testpassword@localhost:5432/testdb")
	defer os.Unsetenv(dbUrlEnv)

	// tempFile, err := os.CreateTemp(".", "example-app-test-config")
	// assert.NoError(t, err)
	// defer tempFile.Close()
	// defer os.Remove(tempFile.Name())

	// Load the config from the environment variable (cfg is empty)
	LoadConfig(nil)

	// Assert the config values
	config := GetConfig()
	assert.NotNil(t, config)
	assert.Equal(t, "postgresql", config.DatasourceUrl.Scheme)
	assert.Equal(t, "localhost:5432", config.DatasourceUrl.Host)
	assert.Equal(t, "/testdb", config.DatasourceUrl.Path)
	assert.Equal(t, "testuser", config.DatasourceUrl.User.Username())
	password, ok := config.DatasourceUrl.User.Password()
	assert.True(t, ok)
	assert.Equal(t, "testpassword", password)
}

// TODO: Implement test for live reload of config file
