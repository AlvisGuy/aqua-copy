package installpackage

import (
	"context"
)

type Executor interface {
	HdiutilAttach(ctx context.Context, dmgPath, mountPoint string) (int, error)
	HdiutilDetach(ctx context.Context, mountPath string) (int, error)
	UnarchivePkg(ctx context.Context, pkgFilePath, dest string) (int, error)
	Exec(ctx context.Context, exePath string, args ...string) (int, error)
	ExecWithEnvs(ctx context.Context, exePath string, args, envs []string) (int, error)
}
