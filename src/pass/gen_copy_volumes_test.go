package pass

import (
	"compose-generator/model"
	"testing"

	"github.com/briandowns/spinner"
	"github.com/compose-spec/compose-go/types"
	"github.com/stretchr/testify/assert"
)



func TestGenerateCopyVolumes1(t *testing.T) {
	// Test data
	templatesPath := "../predefined-templates/path"
	project := &model.CGProject{
		Composition: &types.Project{
			Services: types.Services{
				{
					Name: "Service 1",
					Volumes: []types.ServiceVolumeConfig{
						{
							Type:   types.VolumeTypeBind,
							Source: templatesPath + "/type/template/volumes/volume1",
							Target: "/test/target/in/container",
						},
					},
				},
				{
					Name: "Service 2",
					Volumes: []types.ServiceVolumeConfig{
						{
							Type:   types.VolumeTypeBind,
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
	copyVolumeMockable = func(volume *types.ServiceVolumeConfig, srcPath, dstPath string) {
		copyVolumeCallCount++
		if copyVolumeCallCount == 1 {
			assert.Equal(t, project.Composition.Services[0].Volumes[0], *volume)
			assert.Equal(t, "../predefined-templates/path/type/template/volumes/volume1", srcPath)
			assert.Equal(t, "volumes/volume1", dstPath)
		} else {
			assert.Equal(t, project.Composition.Services[1].Volumes[0], *volume)
			assert.Equal(t, "../predefined-templates/path/type/template/volumes/volume2", srcPath)
			assert.Equal(t, "volumes/volume2", dstPath)
		}
	}
	copyBuildDirMockable = func(build *types.BuildConfig, srcPath, dstPath string) {
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
		Composition: &types.Project{
			Services: types.Services{
				{
					Name: "Service 1",
					Build: &types.BuildConfig{
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
	copyVolumeMockable = func(volume *types.ServiceVolumeConfig, srcPath, dstPath string) {
		assert.Fail(t, "Unexpected call of copyVolume")
	}
	copyBuildDirMockable = func(build *types.BuildConfig, srcPath, dstPath string) {
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
