package repository

type AutoIncrRepository interface {
	NextID(key string) (int64, error)
}
