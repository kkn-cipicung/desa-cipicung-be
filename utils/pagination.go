package utils

const (
	DefaultLimit = 10
	MaxLimit     = 100
)

func NormalizePagination(limit, index *int) {
	if *limit <= 0 {
		*limit = DefaultLimit
	}

	if *limit > MaxLimit {
		*limit = MaxLimit
	}

	if *index < 0 {
		*index = 0
	}
}
