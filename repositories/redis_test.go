package repositories

import (
	"fmt"
	"github.com/fzzy/radix/redis"
	"github.com/sp-lorenzo-arribas/event_validator/domain"
	"github.com/sp-lorenzo-arribas/event_validator/test"
	"os"
	"testing"
)

func TestRedisRepository(t *testing.T) {
	host, port := os.Getenv("EV_TEST_REDIS_HOST"), os.Getenv("EV_TEST_REDIS_PORT")
	domain.Current.NewRepository = func() domain.Repository {
		return NewRedisRepository(host, port)
	}
	test.GenericRepositoryTest(t, func() {

	}, func() {
		// Clear Redis
		cl, _ := redis.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
		cl.Cmd("FLUSHDB")
	})
}
