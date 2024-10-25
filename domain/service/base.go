package service

import (
	sysErrors "echo-clean/domain/error"
	"errors"
	"github.com/lib/pq"
)

type BaseService struct{}

func (b *BaseService) HandleDbError(err error) error {
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return errors.New(sysErrors.E10001.ToJsonError())
			case "foreign_key_violation":
				return errors.New(sysErrors.E10002.ToJsonError())
			case "not_null_violation":
				return errors.New(sysErrors.E10003.ToJsonError())
			case "check_violation":
				return errors.New(sysErrors.E10004.ToJsonError())
			case "string_data_right_truncation":
				return errors.New(sysErrors.E10009.ToJsonError())
			default:
				return errors.New(sysErrors.E10001.ToJsonError())
			}
		}

		return errors.New(sysErrors.E10001.ToJsonError())
	}
	return errors.New(sysErrors.E10001.ToJsonError())
}
