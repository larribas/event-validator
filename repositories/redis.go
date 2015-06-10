package repositories

import (
	"fmt"
	"github.com/fzzy/radix/redis"
	"github.com/sp-lorenzo-arribas/event_validator/domain"
	"log"
)

// RedisNamespace is used in all calls to Redis to provide a particular namespace for the Event Validator,
// in the event that the same Redis cluster is being shared by multiple applications
const RedisNamespace = "github.com/socialpoint/ulog/validator"

type RedisRepository struct {
	host, port string
	client     *redis.Client
}

func NewRedisRepository(host, port string) domain.Repository {
	return &RedisRepository{
		host: host,
		port: port,
	}
}

func (r *RedisRepository) Create(validator *domain.Validator) (version int) {
	nextVersion := r.GetNextVersion(validator.Type)
	r.getRedis().Cmd("HSET", keyFor(validator.Type), nextVersion, validator.Rules)

	validator.Version = nextVersion
	return validator.Version
}

func (r *RedisRepository) GetNextVersion(_type string) int {
	reply := r.getRedis().Cmd("HLEN", keyFor(_type))
	nextVersion, _ := reply.Int()
	return nextVersion
}

func (r *RedisRepository) Inspect(_type string, version int) (*domain.Validator, error) {
	reply := r.getRedis().Cmd("HGET", keyFor(_type), version)
	if reply.Type == redis.NilReply {
		return nil, domain.ErrValidatorDoesNotExist{_type, version}
	}

	rules, _ := reply.Bytes()
	return &domain.Validator{
		Type:    _type,
		Version: version,
		Rules:   rules,
	}, nil
}

func (r *RedisRepository) getRedis() *redis.Client {
	if r.client == nil {
		cl, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", r.host, r.port))
		if err != nil {
			log.Fatalf("Couldn't connect to Redis. Error: %s\n", err.Error())
		}

		r.client = cl
	}

	return r.client
}

func keyFor(_type string) string {
	return fmt.Sprintf("%s__%s", RedisNamespace, _type)
}
