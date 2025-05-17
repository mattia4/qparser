package file

import "errors"

const (
	errCannotCreateDirMsg               = "can't create dir"
	errStdInMustProvideAtLeastOneArgMsg = "must provide at least one arg from stdin"
)

func GetErrCannotCreateDir() error {

	return errors.New("err -> " + errCannotCreateDirMsg)

}

func GetErrStdInMustProvideAtLeastOneArg() error {

	return errors.New("err -> " + errStdInMustProvideAtLeastOneArgMsg)

}
