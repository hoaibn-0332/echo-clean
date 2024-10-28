package service

import (
	sysErrors "echo-clean/domain/error"
	"errors"
	"github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type BaseService struct{}

func (b *BaseService) HandleDbError(err error) []sysErrors.Error {
	log.Debug().Msgf("Database error: %v", err)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return []sysErrors.Error{sysErrors.E10001}
			case "foreign_key_violation":
				return []sysErrors.Error{sysErrors.E10002}
			case "not_null_violation":
				return []sysErrors.Error{sysErrors.E10003}
			case "check_violation":
				return []sysErrors.Error{sysErrors.E10004}
			case "string_data_right_truncation":
				return []sysErrors.Error{sysErrors.E10009}
			default:
				return []sysErrors.Error{sysErrors.E10001}
			}
		}

		return []sysErrors.Error{sysErrors.E10001}
	}
	return []sysErrors.Error{sysErrors.E10001}
}
