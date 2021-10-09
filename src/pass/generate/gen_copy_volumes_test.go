package pass

import (
	"compose-generator/model"
	"errors"
	"os"
	"testing"

	"github.com/briandowns/spinner"
	spec "github.com/compose-spec/compose-go/types"
	"github.com/otiai10/copy"
	"github.com/stretchr/testify/assert"
)

// ---------------------------------------- GenerateCopyVolumes ----------------------------------------

func TestGenerateCopyVolumes1(t *testing.T) {
	// Test data
	templatesPath := "../predefined-templates/path"
	project := &model.CGProject{
		Composition: &spec.Project{
			WorkingDir: "./",
			Services: spec.Services{
				{
					Name: "Service 1",
					Volumes: []spec.ServiceVolumeConfig{
						{
							Type:   spec.VolumeTypeBind,
							Source: templatesPath + "/type/template/volumes/volume1",
							Target: "/test/target/in/container",
						},
					},
				},
				{
					Name: "Service 2",
					Volumes: []spec.ServiceVolumeConfig{
						{
							Type:   spec.VolumeTypeBind,
							Source: templatesPath + "/type/template/volumes/volume2",
							Target: "/test/target/in/other/container",
						},
					},
				},
			},
		},
	}
	// Mock functions
	startProcess = func(text string) (s *spinner.Spinner) {
		assert.Equal(t, "Copying volumes ...", text)
		return nil
	}
	stopProcessCallCount := 0
	stopProcess = func(s *spinner.Spinner) {
		stopProcessCallCount++
		assert.Nil(t, s)
	}
	copyVolumeCallCount := 0
	copyVolumeMockable = func(volume *spec.ServiceVolumeConfig, srcPath, dstPath string) {
		copyVolumeCallCount++
		if copyVolumeCallCount == 1 {
			assert.Equal(t, project.Composition.Services[0].Volumes[0], *volume)
			assert.Equal(t, "../predefined-templates/path/type/template/volumes/volume1", srcPath)
			assert.Equal(t, "./volumes/volume1", dstPath)
		} else {
			assert.Equal(t, project.Composition.Services[1].Volumes[0], *volume)
			assert.Equal(t, "../predefined-templates/path/type/template/volumes/volume2", srcPath)
			assert.Equal(t, "./volumes/volume2", dstPath)
		}
	}
	copyBuildDirMockable = func(build *spec.BuildConfig, srcPath, dstPath string) {
		assert.Fail(t, "Unexpected call of copyBuildDir")
	}
	getPredefinedServicesPath = func() string {
		return templatesPath
	}
	// Execute test
	GenerateCopyVolumes(project)
	// Assert
	assert.Equal(t, 1, stopProcessCallCount)
	assert.Equal(t, 2, copyVolumeCallCount)
}

func TestGenerateCopyVolumes2(t *testing.T) {
	// Test data
	templatesPath := "../predefined-templates"
	project := &model.CGProject{
		Composition: &spec.Project{
			WorkingDir: "./",
			Services: spec.Services{
				{
					Name: "Service 1",
					Build: &spec.BuildConfig{
						Context:    templatesPath + "/type/template/frontend",
						Dockerfile: "Dockerfile",
					},
				},
			},
		},
	}
	// Mock functions
	startProcess = func(text string) (s *spinner.Spinner) {
		assert.Equal(t, "Copying volumes ...", text)
		return nil
	}
	stopProcessCallCount := 0
	stopProcess = func(s *spinner.Spinner) {
		stopProcessCallCount++
		assert.Nil(t, s)
	}
	copyVolumeMockable = func(volume *spec.ServiceVolumeConfig, srcPath, dstPath string) {
		assert.Fail(t, "Unexpected call of copyVolume")
	}
	copyBuildDirMockable = func(build *spec.BuildConfig, srcPath, dstPath string) {
		assert.Equal(t, project.Composition.Services[0].Build, build)
		assert.Equal(t, "../predefined-templates/type/template/frontend", srcPath)
		assert.Equal(t, "frontend", dstPath)
	}
	getPredefinedServicesPath = func() string {
		return templatesPath
	}
	// Execute test
	GenerateCopyVolumes(project)
	// Assert
	assert.Equal(t, 1, stopProcessCallCount)
}

// -------------------------------------------- copyVolume ---------------------------------------------

func TestCopyVolume1(t *testing.T) {
	// Test data
	templatesPath := "../predefined-templates"
	volume := &spec.ServiceVolumeConfig{
		Type:   spec.VolumeTypeBind,
		Source: templatesPath + "/type/template/volumes/volume1",
		Target: "/test/target/in/container",
	}
	srcPath := "../predefined-templates/path/type/template/volumes/volume1"
	dstPath := "volumes/volume1"
	// Mock functions
	fileExists = func(path string) bool {
		assert.Equal(t, srcPath, path)
		return false
	}
	mkdirAll = func(path string, perm os.FileMode) error {
		assert.Equal(t, dstPath, path)
		assert.Equal(t, os.FileMode(0777), perm)
		return nil
	}
	printWarning = func(description string) {
		assert.Fail(t, "Unexpected call of printWarning")
	}
	// Execute test
	copyVolume(volume, srcPath, dstPath)
	// Assert
	assert.Equal(t, dstPath, volume.Source)
}

func TestCopyVolume2(t *testing.T) {
	// Test data
	templatesPath := "../predefined-templates"
	volume := &spec.ServiceVolumeConfig{
		Type:   spec.VolumeTypeBind,
		Source: templatesPath + "/type/template/volumes/volume1",
		Target: "/test/target/in/container",
	}
	srcPath := "../predefined-templates/path/type/template/volumes/volume1"
	dstPath := "volumes/volume1"
	// Mock functions
	fileExists = func(path string) bool {
		assert.Equal(t, srcPath, path)
		return false
	}
	mkdirAll = func(path string, perm os.FileMode) error {
		assert.Equal(t, dstPath, path)
		assert.Equal(t, os.FileMode(0777), perm)
		return errors.New("MkdirAll error")
	}
	printWarning = func(description string) {
		assert.Equal(t, "Could not create volume dir", description)
	}
	// Execute test
	copyVolume(volume, srcPath, dstPath)
	// Assert
	assert.Equal(t, dstPath, volume.Source)
}

func TestCopyVolume3(t *testing.T) {
	// Test data
	templatesPath := "../predefined-templates"
	volume := &spec.ServiceVolumeConfig{
		Type:   spec.VolumeTypeBind,
		Source: templatesPath + "/type/template/volumes/volume1",
		Target: "/test/target/in/container",
	}
	srcPath := "../predefined-templates/path/type/template/volumes/volume1"
	dstPath := "volumes/volume1"
	// Mock functions
	fileExists = func(path string) bool {
		assert.Equal(t, srcPath, path)
		return true
	}
	mkdirAll = func(path string, perm os.FileMode) error {
		assert.Equal(t, dstPath, path)
		assert.Equal(t, os.FileMode(0777), perm)
		return errors.New("MkdirAll error")
	}
	copyFile = func(src, dest string, opt ...copy.Options) error {
		assert.Equal(t, srcPath, src)
		assert.Equal(t, dstPath, dest)
		return nil
	}
	// Execute test
	copyVolume(volume, srcPath, dstPath)
	// Assert
	assert.Equal(t, dstPath, volume.Source)
}

// ------------------------------------------- copyBuildDir --------------------------------------------

func TestCopyBuildDir1(t *testing.T) {
	// Test data
	templatesPath := "../predefined-templates"
	buildDir := &spec.BuildConfig{
		Context:    templatesPath + "/type/template/backend",
		Dockerfile: "Dockerfile",
	}
	srcPath := "../predefined-templates/path/type/template/volumes/volume1"
	dstPath := "volumes/volume1"
	// Mock functions
	fileExists = func(path string) bool {
		assert.Equal(t, srcPath, path)
		return false
	}
	mkdirAll = func(path string, perm os.FileMode) error {
		assert.Equal(t, dstPath, path)
		assert.Equal(t, os.FileMode(0777), perm)
		return nil
	}
	printWarning = func(description string) {
		assert.Fail(t, "Unexpected call of printWarning")
	}
	// Execute test
	copyBuildDir(buildDir, srcPath, dstPath)
	// Assert
	assert.Equal(t, dstPath, buildDir.Context)
}

func TestCopyBuildDir2(t *testing.T) {
	// Test data
	templatesPath := "../predefined-templates"
	buildDir := &spec.BuildConfig{
		Context:    templatesPath + "/type/template/backend",
		Dockerfile: "Dockerfile",
	}
	srcPath := "../predefined-templates/path/type/template/volumes/volume1"
	dstPath := "volumes/volume1"
	// Mock functions
	fileExists = func(path string) bool {
		assert.Equal(t, srcPath, path)
		return false
	}
	mkdirAll = func(path string, perm os.FileMode) error {
		assert.Equal(t, dstPath, path)
		assert.Equal(t, os.FileMode(0777), perm)
		return errors.New("MkdirAll error")
	}
	printWarning = func(description string) {
		assert.Equal(t, "Could not create volume dir", description)
	}
	// Execute test
	copyBuildDir(buildDir, srcPath, dstPath)
	// Assert
	assert.Equal(t, dstPath, buildDir.Context)
}

func TestCopyBuildDir3(t *testing.T) {
	// Test data
	templatesPath := "../predefined-templates"
	buildDir := &spec.BuildConfig{
		Context:    templatesPath + "/type/template/backend",
		Dockerfile: "Dockerfile",
	}
	srcPath := "../predefined-templates/path/type/template/volumes/volume1"
	dstPath := "volumes/volume1"
	// Mock functions
	fileExists = func(path string) bool {
		assert.Equal(t, srcPath, path)
		return true
	}
	mkdirAll = func(path string, perm os.FileMode) error {
		assert.Equal(t, dstPath, path)
		assert.Equal(t, os.FileMode(0777), perm)
		return errors.New("MkdirAll error")
	}
	copyFile = func(src, dest string, opt ...copy.Options) error {
		assert.Equal(t, srcPath, src)
		assert.Equal(t, dstPath, dest)
		return nil
	}
	// Execute test
	copyBuildDir(buildDir, srcPath, dstPath)
	// Assert
	assert.Equal(t, dstPath, buildDir.Context)
}
