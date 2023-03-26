// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package controller

import (
	"context"
	"github.com/aquaproj/aqua/v2/pkg/checksum"
	"github.com/aquaproj/aqua/v2/pkg/config"
	"github.com/aquaproj/aqua/v2/pkg/config-finder"
	"github.com/aquaproj/aqua/v2/pkg/config-reader"
	"github.com/aquaproj/aqua/v2/pkg/controller/cp"
	exec2 "github.com/aquaproj/aqua/v2/pkg/controller/exec"
	"github.com/aquaproj/aqua/v2/pkg/controller/generate"
	"github.com/aquaproj/aqua/v2/pkg/controller/generate-registry"
	"github.com/aquaproj/aqua/v2/pkg/controller/generate/output"
	"github.com/aquaproj/aqua/v2/pkg/controller/initcmd"
	"github.com/aquaproj/aqua/v2/pkg/controller/initpolicy"
	"github.com/aquaproj/aqua/v2/pkg/controller/install"
	"github.com/aquaproj/aqua/v2/pkg/controller/list"
	"github.com/aquaproj/aqua/v2/pkg/controller/updateaqua"
	"github.com/aquaproj/aqua/v2/pkg/controller/updatechecksum"
	"github.com/aquaproj/aqua/v2/pkg/controller/which"
	"github.com/aquaproj/aqua/v2/pkg/cosign"
	"github.com/aquaproj/aqua/v2/pkg/download"
	"github.com/aquaproj/aqua/v2/pkg/exec"
	"github.com/aquaproj/aqua/v2/pkg/github"
	"github.com/aquaproj/aqua/v2/pkg/install-registry"
	"github.com/aquaproj/aqua/v2/pkg/installpackage"
	"github.com/aquaproj/aqua/v2/pkg/link"
	"github.com/aquaproj/aqua/v2/pkg/policy"
	"github.com/aquaproj/aqua/v2/pkg/runtime"
	"github.com/aquaproj/aqua/v2/pkg/slsa"
	"github.com/aquaproj/aqua/v2/pkg/unarchive"
	"github.com/spf13/afero"
	"github.com/suzuki-shunsuke/go-osenv/osenv"
	"io"
	"net/http"
)

// Injectors from wire.go:

func InitializeListCommandController(ctx context.Context, param *config.Param, httpClient *http.Client, rt *runtime.Runtime) *list.Controller {
	fs := afero.NewOsFs()
	configFinder := finder.NewConfigFinder(fs)
	configReaderImpl := reader.New(fs, param)
	repositoriesService := github.New(ctx)
	httpDownloader := download.NewHTTPDownloader(httpClient)
	gitHubContentFileDownloader := download.NewGitHubContentFileDownloader(repositoriesService, httpDownloader)
	executor := exec.New()
	downloader := download.NewDownloader(repositoriesService, httpDownloader)
	verifierImpl := cosign.NewVerifier(executor, fs, downloader, param)
	executorImpl := slsa.NewExecutor(executor, param)
	slsaVerifierImpl := slsa.New(downloader, fs, executorImpl)
	installerImpl := registry.New(param, gitHubContentFileDownloader, fs, rt, verifierImpl, slsaVerifierImpl)
	controller := list.NewController(configFinder, configReaderImpl, installerImpl, fs)
	return controller
}

func InitializeGenerateRegistryCommandController(ctx context.Context, param *config.Param, httpClient *http.Client, stdout io.Writer) *genrgst.Controller {
	fs := afero.NewOsFs()
	repositoriesService := github.New(ctx)
	outputter := output.New(stdout, fs)
	controller := genrgst.NewController(fs, repositoriesService, outputter)
	return controller
}

func InitializeInitCommandController(ctx context.Context, param *config.Param) *initcmd.Controller {
	repositoriesService := github.New(ctx)
	fs := afero.NewOsFs()
	controller := initcmd.New(repositoriesService, fs)
	return controller
}

func InitializeInitPolicyCommandController(ctx context.Context) *initpolicy.Controller {
	fs := afero.NewOsFs()
	controller := initpolicy.New(fs)
	return controller
}

func InitializeGenerateCommandController(ctx context.Context, param *config.Param, httpClient *http.Client, rt *runtime.Runtime) *generate.Controller {
	fs := afero.NewOsFs()
	configFinder := finder.NewConfigFinder(fs)
	configReaderImpl := reader.New(fs, param)
	repositoriesService := github.New(ctx)
	httpDownloader := download.NewHTTPDownloader(httpClient)
	gitHubContentFileDownloader := download.NewGitHubContentFileDownloader(repositoriesService, httpDownloader)
	executor := exec.New()
	downloader := download.NewDownloader(repositoriesService, httpDownloader)
	verifierImpl := cosign.NewVerifier(executor, fs, downloader, param)
	executorImpl := slsa.NewExecutor(executor, param)
	slsaVerifierImpl := slsa.New(downloader, fs, executorImpl)
	installerImpl := registry.New(param, gitHubContentFileDownloader, fs, rt, verifierImpl, slsaVerifierImpl)
	fuzzyFinder := generate.NewFuzzyFinder()
	versionSelector := generate.NewVersionSelector()
	controller := generate.New(configFinder, configReaderImpl, installerImpl, repositoriesService, fs, fuzzyFinder, versionSelector)
	return controller
}

func InitializeInstallCommandController(ctx context.Context, param *config.Param, httpClient *http.Client, rt *runtime.Runtime) *install.Controller {
	fs := afero.NewOsFs()
	configFinder := finder.NewConfigFinder(fs)
	configReaderImpl := reader.New(fs, param)
	repositoriesService := github.New(ctx)
	httpDownloader := download.NewHTTPDownloader(httpClient)
	gitHubContentFileDownloader := download.NewGitHubContentFileDownloader(repositoriesService, httpDownloader)
	executor := exec.New()
	downloader := download.NewDownloader(repositoriesService, httpDownloader)
	verifierImpl := cosign.NewVerifier(executor, fs, downloader, param)
	executorImpl := slsa.NewExecutor(executor, param)
	slsaVerifierImpl := slsa.New(downloader, fs, executorImpl)
	installerImpl := registry.New(param, gitHubContentFileDownloader, fs, rt, verifierImpl, slsaVerifierImpl)
	linker := link.New()
	checksumDownloaderImpl := download.NewChecksumDownloader(repositoriesService, rt, httpDownloader)
	calculator := checksum.NewCalculator()
	unarchiverImpl := unarchive.New(executor)
	checkerImpl := policy.NewChecker()
	installpackageInstallerImpl := installpackage.New(param, downloader, rt, fs, linker, executor, checksumDownloaderImpl, calculator, unarchiverImpl, checkerImpl, verifierImpl, slsaVerifierImpl)
	policyConfigReaderImpl := policy.NewConfigReader(fs)
	controller := install.New(param, configFinder, configReaderImpl, installerImpl, installpackageInstallerImpl, fs, rt, policyConfigReaderImpl)
	return controller
}

func InitializeWhichCommandController(ctx context.Context, param *config.Param, httpClient *http.Client, rt *runtime.Runtime) *which.ControllerImpl {
	fs := afero.NewOsFs()
	configFinder := finder.NewConfigFinder(fs)
	configReaderImpl := reader.New(fs, param)
	repositoriesService := github.New(ctx)
	httpDownloader := download.NewHTTPDownloader(httpClient)
	gitHubContentFileDownloader := download.NewGitHubContentFileDownloader(repositoriesService, httpDownloader)
	executor := exec.New()
	downloader := download.NewDownloader(repositoriesService, httpDownloader)
	verifierImpl := cosign.NewVerifier(executor, fs, downloader, param)
	executorImpl := slsa.NewExecutor(executor, param)
	slsaVerifierImpl := slsa.New(downloader, fs, executorImpl)
	installerImpl := registry.New(param, gitHubContentFileDownloader, fs, rt, verifierImpl, slsaVerifierImpl)
	osEnv := osenv.New()
	linker := link.New()
	controllerImpl := which.New(param, configFinder, configReaderImpl, installerImpl, rt, osEnv, fs, linker)
	return controllerImpl
}

func InitializeExecCommandController(ctx context.Context, param *config.Param, httpClient *http.Client, rt *runtime.Runtime) *exec2.Controller {
	repositoriesService := github.New(ctx)
	httpDownloader := download.NewHTTPDownloader(httpClient)
	downloader := download.NewDownloader(repositoriesService, httpDownloader)
	fs := afero.NewOsFs()
	linker := link.New()
	executor := exec.New()
	checksumDownloaderImpl := download.NewChecksumDownloader(repositoriesService, rt, httpDownloader)
	calculator := checksum.NewCalculator()
	unarchiverImpl := unarchive.New(executor)
	checkerImpl := policy.NewChecker()
	verifierImpl := cosign.NewVerifier(executor, fs, downloader, param)
	executorImpl := slsa.NewExecutor(executor, param)
	slsaVerifierImpl := slsa.New(downloader, fs, executorImpl)
	installerImpl := installpackage.New(param, downloader, rt, fs, linker, executor, checksumDownloaderImpl, calculator, unarchiverImpl, checkerImpl, verifierImpl, slsaVerifierImpl)
	configFinder := finder.NewConfigFinder(fs)
	configReaderImpl := reader.New(fs, param)
	gitHubContentFileDownloader := download.NewGitHubContentFileDownloader(repositoriesService, httpDownloader)
	registryInstallerImpl := registry.New(param, gitHubContentFileDownloader, fs, rt, verifierImpl, slsaVerifierImpl)
	osEnv := osenv.New()
	controllerImpl := which.New(param, configFinder, configReaderImpl, registryInstallerImpl, rt, osEnv, fs, linker)
	policyConfigReaderImpl := policy.NewConfigReader(fs)
	controller := exec2.New(param, installerImpl, controllerImpl, executor, osEnv, fs, policyConfigReaderImpl, checkerImpl)
	return controller
}

func InitializeUpdateAquaCommandController(ctx context.Context, param *config.Param, httpClient *http.Client, rt *runtime.Runtime) *updateaqua.Controller {
	fs := afero.NewOsFs()
	repositoriesService := github.New(ctx)
	httpDownloader := download.NewHTTPDownloader(httpClient)
	downloader := download.NewDownloader(repositoriesService, httpDownloader)
	linker := link.New()
	executor := exec.New()
	checksumDownloaderImpl := download.NewChecksumDownloader(repositoriesService, rt, httpDownloader)
	calculator := checksum.NewCalculator()
	unarchiverImpl := unarchive.New(executor)
	checkerImpl := policy.NewChecker()
	verifierImpl := cosign.NewVerifier(executor, fs, downloader, param)
	executorImpl := slsa.NewExecutor(executor, param)
	slsaVerifierImpl := slsa.New(downloader, fs, executorImpl)
	installerImpl := installpackage.New(param, downloader, rt, fs, linker, executor, checksumDownloaderImpl, calculator, unarchiverImpl, checkerImpl, verifierImpl, slsaVerifierImpl)
	controller := updateaqua.New(param, fs, rt, repositoriesService, installerImpl)
	return controller
}

func InitializeCopyCommandController(ctx context.Context, param *config.Param, httpClient *http.Client, rt *runtime.Runtime) *cp.Controller {
	repositoriesService := github.New(ctx)
	httpDownloader := download.NewHTTPDownloader(httpClient)
	downloader := download.NewDownloader(repositoriesService, httpDownloader)
	fs := afero.NewOsFs()
	linker := link.New()
	executor := exec.New()
	checksumDownloaderImpl := download.NewChecksumDownloader(repositoriesService, rt, httpDownloader)
	calculator := checksum.NewCalculator()
	unarchiverImpl := unarchive.New(executor)
	checkerImpl := policy.NewChecker()
	verifierImpl := cosign.NewVerifier(executor, fs, downloader, param)
	executorImpl := slsa.NewExecutor(executor, param)
	slsaVerifierImpl := slsa.New(downloader, fs, executorImpl)
	installerImpl := installpackage.New(param, downloader, rt, fs, linker, executor, checksumDownloaderImpl, calculator, unarchiverImpl, checkerImpl, verifierImpl, slsaVerifierImpl)
	configFinder := finder.NewConfigFinder(fs)
	configReaderImpl := reader.New(fs, param)
	gitHubContentFileDownloader := download.NewGitHubContentFileDownloader(repositoriesService, httpDownloader)
	registryInstallerImpl := registry.New(param, gitHubContentFileDownloader, fs, rt, verifierImpl, slsaVerifierImpl)
	osEnv := osenv.New()
	controllerImpl := which.New(param, configFinder, configReaderImpl, registryInstallerImpl, rt, osEnv, fs, linker)
	policyConfigReaderImpl := policy.NewConfigReader(fs)
	controller := install.New(param, configFinder, configReaderImpl, registryInstallerImpl, installerImpl, fs, rt, policyConfigReaderImpl)
	cpController := cp.New(param, installerImpl, fs, rt, controllerImpl, controller, policyConfigReaderImpl)
	return cpController
}

func InitializeUpdateChecksumCommandController(ctx context.Context, param *config.Param, httpClient *http.Client, rt *runtime.Runtime) *updatechecksum.Controller {
	fs := afero.NewOsFs()
	configFinder := finder.NewConfigFinder(fs)
	configReaderImpl := reader.New(fs, param)
	repositoriesService := github.New(ctx)
	httpDownloader := download.NewHTTPDownloader(httpClient)
	gitHubContentFileDownloader := download.NewGitHubContentFileDownloader(repositoriesService, httpDownloader)
	executor := exec.New()
	downloader := download.NewDownloader(repositoriesService, httpDownloader)
	verifierImpl := cosign.NewVerifier(executor, fs, downloader, param)
	executorImpl := slsa.NewExecutor(executor, param)
	slsaVerifierImpl := slsa.New(downloader, fs, executorImpl)
	installerImpl := registry.New(param, gitHubContentFileDownloader, fs, rt, verifierImpl, slsaVerifierImpl)
	checksumDownloaderImpl := download.NewChecksumDownloader(repositoriesService, rt, httpDownloader)
	controller := updatechecksum.New(param, configFinder, configReaderImpl, installerImpl, fs, rt, checksumDownloaderImpl, downloader, gitHubContentFileDownloader)
	return controller
}
