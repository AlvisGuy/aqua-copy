package unarchive

import (
	"context"
	"fmt"

	"github.com/aquaproj/aqua/v2/pkg/util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
)

const FormatDMG string = "dmg"

type dmgUnarchiver struct {
	dest     string
	executor Executor
	fs       afero.Fs
}

type Executor interface {
	HdiutilAttach(ctx context.Context, dmgPath, mountPoint string) (int, error)
	HdiutilDetach(ctx context.Context, mountPath string) (int, error)
	UnarchivePkg(ctx context.Context, pkgFilePath, dest string) (int, error)
}

func (unarchiver *dmgUnarchiver) Unarchive(ctx context.Context, logE *logrus.Entry, src *File) error {
	if err := util.MkdirAll(unarchiver.fs, unarchiver.dest); err != nil {
		return fmt.Errorf("create a directory: %w", err)
	}

	tempFilePath, err := src.Body.GetPath()
	if err != nil {
		return fmt.Errorf("get a temporal file path: %w", err)
	}

	tmpMountPoint, err := afero.TempDir(unarchiver.fs, "", "")
	if err != nil {
		return fmt.Errorf("create a temporal file: %w", err)
	}

	if _, err := unarchiver.executor.HdiutilAttach(ctx, tempFilePath, tmpMountPoint); err != nil {
		if err := unarchiver.fs.Remove(tmpMountPoint); err != nil {
			logE.WithError(err).Warn("remove a temporal directory created to attach a DMG file")
		}
		return fmt.Errorf("hdiutil attach: %w", err)
	}
	defer func() {
		if _, err := unarchiver.executor.HdiutilDetach(ctx, tmpMountPoint); err != nil {
			logE.WithError(err).Warn("detach a DMG file")
		}
		if err := unarchiver.fs.Remove(tmpMountPoint); err != nil {
			logE.WithError(err).Warn("remove a temporal directory created to attach a DMG file")
		}
	}()

	if err := util.Copy(unarchiver.fs, tmpMountPoint, unarchiver.dest); err != nil {
		return fmt.Errorf("copy a directory: %w", err)
	}

	return nil
}
