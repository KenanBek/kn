package crawler

import (
	"fmt"
	"kn/se/internal/domain"
	"log"
	"os"
	"testing"

	"github.com/ory/dockertest"
	"github.com/stretchr/testify/assert"
)

var mr *MongoRepository

func GetMongodbEnv() (string, string) {
	return os.Getenv(hostEnv), os.Getenv(portEnv)
}

func SetMongodbEnv(host, port string) error {
	var err error

	err = os.Setenv("KN_MONGODB_HOST", "localhost")
	if err != nil {
		return err
	}

	err = os.Setenv("KN_MONGODB_PORT", port)
	if err != nil {
		return err
	}

	return nil
}

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("can not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("mongo", "latest", []string{})
	if err != nil {
		log.Fatalf("can not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error

		host := "localhost"
		port := resource.GetPort("27017/tcp")

		err = SetMongodbEnv(host, port)
		if err != nil {
			return err
		}

		mr = NewMongoRepository()

		err = mr.Ping()
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		log.Fatalf("can not connect to docker: %s", err)
	}

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("can not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestMongoClient(t *testing.T) {
	t.Run("Ping", func(t *testing.T) {
		err := mr.Ping()
		assert.Equal(t, nil, err)
	})
	t.Run("AddLink", func(t *testing.T) {
		var err error
		for i := 1; i < 10; i++ {
			err = mr.SaveLink(&domain.Link{
				Hash:      fmt.Sprintf("abcde%v", i),
				URL:       fmt.Sprintf("http://example.com/abcde%v", i),
				IsArticle: false,
			})
			if err != nil {
				log.Fatalf("SaveLing failed with error: %s", err)
			}
		}

		assert.Nil(t, err)
		assert.Equal(t, true, mr.HasLink("abcde1"))
	})
	t.Run("HasLink", func(t *testing.T) {
		var err error
		for i := 10; i < 20; i++ {
			err = mr.SaveLink(&domain.Link{
				Hash:      fmt.Sprintf("abcd%v", i),
				URL:       fmt.Sprintf("http://example.com/abcd%v", i),
				IsArticle: false,
			})
			if err != nil {
				log.Fatalf("SaveLing failed with error: %s", err)
			}
		}

		assert.Nil(t, err)

		assert.Equal(t, true, mr.HasLink("abcde1"))
		assert.Equal(t, true, mr.HasLink("abcd10"))
		assert.Equal(t, false, mr.HasLink("abcd20"))
	})
}

func TestMongoRepository_NoConnection(t *testing.T) {
	var err error

	h, p := GetMongodbEnv()
	err = SetMongodbEnv("wrong-host", "wrong-port")
	assert.Nil(t, err, fmt.Sprintf("setting env variables failed with error %v", err))

	failMR := NewMongoRepository()

	err = failMR.Ping()
	assert.Error(t, err, "ping on a failed connection should raise an error")
	assert.Contains(t, err.Error(), "database ping error")

	err = failMR.SaveLink(&domain.Link{
		Hash:      "abc",
		URL:       "abc",
		IsArticle: false,
	})
	assert.Error(t, err, "saving link on a failed connection should raise an error")
	assert.Contains(t, err.Error(), "database error on save link")

	hasLink := failMR.HasLink("abc")
	assert.False(t, hasLink, "checking link on a failed connection should return false as a result")

	err = SetMongodbEnv(h, p)
	assert.Nil(t, err, fmt.Sprintf("setting env variables failed with error %v", err))
}
