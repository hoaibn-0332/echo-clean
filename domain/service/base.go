package service

import (
	serviceErrors "echo-clean/domain/error"
	"errors"
	"github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type BaseService struct{}

func (b *BaseService) HandleDbError(err error) []serviceErrors.Error {
	log.Debug().Msgf("Database error: %v", err)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return []serviceErrors.Error{serviceErrors.E10001}
			case "foreign_key_violation":
				return []serviceErrors.Error{serviceErrors.E10002}
			case "not_null_violation":
				return []serviceErrors.Error{serviceErrors.E10003}
			case "check_violation":
				return []serviceErrors.Error{serviceErrors.E10004}
			case "string_data_right_truncation":
				return []serviceErrors.Error{serviceErrors.E10009}
			default:
				return []serviceErrors.Error{serviceErrors.E10001}
			}
		}

		return []serviceErrors.Error{serviceErrors.E10001}
	}
	return []serviceErrors.Error{serviceErrors.E10001}
}
