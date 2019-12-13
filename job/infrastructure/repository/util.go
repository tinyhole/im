package repository

const (
	defaultLimit = 30
)

func calcPage(page, pageSize int) (offset, limit int) {
	offset = (page - 1) * pageSize
	if offset <= 0 {
		offset = 0
	}
	limit = pageSize
	if limit <= 0 || limit > defaultLimit {
		limit = defaultLimit
	}
	return
}
