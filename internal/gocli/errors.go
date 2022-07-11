package gocli

import "github.com/Tauhoo/adon-desktop/internal/errors"

var (
	ReadDirFailCode      errors.Code = "READ_DIR_FAIL"
	GetGoBinPathFailCode errors.Code = "GET_GO_BIN_PATH_FAIL"
	SetPATHEnvFailCode   errors.Code = "SET_PATH_ENV_FAIL"
)
