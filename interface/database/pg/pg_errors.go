package pg

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/airbusgeo/geocube/internal/utils"
	"github.com/lib/pq"
)

/* http://www.postgresql.org/docs/9.3/static/errcodes-appendix.html */
const (
	noError             = "00000"
	connectionFailure   = "08006"
	foreignKeyViolation = "23503"
	uniqueViolation     = "23505"
	checkViolation      = "23514"
	noDataFound         = "P0002"

	notPqError = "X"
)

func extractKeyValueFromDetail(err *pq.Error) (string, string) {
	if err != nil {
		re := regexp.MustCompile(`\((.*)\)=\((.*)\)`)
		if value := re.FindStringSubmatch(err.Detail); len(value) == 3 {
			return value[1], value[2]
		}
	}
	return "", ""
}

func pqErrorCode(err error) pq.ErrorCode {
	if err == nil {
		return noError
	}
	var pqerr *pq.Error
	if errors.As(err, &pqerr) {
		return pqerr.Code
	}
	return notPqError
}

func pqErrorFormat(format string, err error) error {
	ferr := fmt.Errorf(format, err)
	switch pqErrorCode(err) {
	case connectionFailure:
		return utils.MakeTemporary(ferr)
	}
	retriable := []string{"connection refused", "connection reset"}
	for _, s := range retriable {
		if strings.Contains(err.Error(), s) {
			return utils.MakeTemporary(ferr)
		}
	}
	return ferr
}
