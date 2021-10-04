package cmd

import (
	"compose-generator/model"
	"errors"
	"io/fs"
	"io/ioutil"
	"testing"

	"github.com/briandowns/spinner"
	spec "github.com/compose-spec/compose-go/types"
	"github.com/otiai10/copy"
	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------- askForTemplate ----------------------------------------------------------------

func TestAskForTemplate1(t *testing.T) {
	// Test data
	templateList := map[string]*model.CGProjectMetadata{
		"Template 1": {
			Name:           "Template 1",
			LastModifiedAt: 1632579623478000000,
		},
		"Template 2": {
			Name:           "Template 2",
			LastModifiedAt: 1632579661510000000,
		},
	}
	// Mock functions
	startProcess = func(text string) *spinner.Spinner {
		assert.Equal(t, "Loading template list ...", text)
		return nil
	}
	getTemplateMetadataListMockable = func() map[string]*model.CGProjectMetadata {
		return templateList
	}
	stopProcess = func(s *spinner.Spinner) {
		assert.Nil(t, s)
	}
	pelCallCount := 0
	pel = func() {
		pelCallCount++
	}
	menuQuestionIndex = func(label string, items []string) int {
		assert.Equal(t, "Which template do you want to load?", label)
		assert.ElementsMatch(t, []string{"Template 1 (Saved at: Sep-25-21 2:20:23 PM)", "Template 2 (Saved at: Sep-25-21 2:21:01 PM)"}, items)
		if items[0] == "Template 1 (Saved at: Sep-25-21 2:20:23 PM)" {
			return 0
		}
		return 1
	}
	printError = func(description string, err error, exit bool) {
		assert.Fail(t, "Unexpected call of printError")
	}
	// Execute test
	result := askForTemplate()
	// Assert
	assert.Equal(t, "Template 1", result)
	assert.Equal(t, 1, pelCallCount)
}

func TestAskForTemplate2(t *testing.T) {
	// Test data
	templateList := map[string]*model.CGProjectMetadata{}
	// Mock functions
	startProcess = func(text string) *spinner.Spinner {
		assert.Equal(t, "Loading template list ...", text)
		return nil
	}
	getTemplateMetadataListMockable = func() map[string]*model.CGProjectMetadata {
		return templateList
	}
	stopProcess = func(s *spinner.Spinner) {
		assert.Nil(t, s)
	}
	pelCallCount := 0
	pel = func() {
		pelCallCount++
	}
	menuQuestionIndex = func(label string, items []string) int {
		assert.Fail(t, "Unexpected call of menuQuestionIndex")
		return 1
	}
	printError = func(description string, err error, exit bool) {
		assert.Equal(t, "No templates found. Use \"$ compose-generator save <template-name>\" to save one.", description)
		assert.Nil(t, err)
		assert.True(t, exit)
	}
	// Execute test
	result := askForTemplate()
	// Assert
	assert.Equal(t, "", result)
	assert.Equal(t, 1, pelCallCount)
}

// ---------------------------------------------------------------- showTemplateList ---------------------------------------------------------------

func TestShowTemplateList1(t *testing.T) {
	// Test data
	templateList := map[string]*model.CGProjectMetadata{
		"Template 1": {
			Name:           "Template 1",
			LastModifiedAt: 1632579623478000000,
		},
		"Template 2": {
			Name:           "Template 2",
			LastModifiedAt: 1632579661510000000,
		},
	}
	// Mock functions
	startProcess = func(text string) *spinner.Spinner {
		assert.Equal(t, "Loading template list ...", text)
		return nil
	}
	getTemplateMetadataListMockable = func() map[string]*model.CGProjectMetadata {
		return templateList
	}
	stopProcess = func(s *spinner.Spinner) {
		assert.Nil(t, s)
	}
	pelCallCount := 0
	pel = func() {
		pelCallCount++
	}
	printHeadingCallCount := 0
	printHeading = func(text string) {
		printHeadingCallCount++
		assert.Equal(t, "List of all templates:", text)
	}
	// Execute test
	showTemplateList()
	// Assert
	assert.Equal(t, 2, pelCallCount)
	assert.Equal(t, 1, printHeadingCallCount)
}

func TestShowTemplateList2(t *testing.T) {
	// Test data
	templateList := map[string]*model.CGProjectMetadata{}
	// Mock functions
	startProcess = func(text string) *spinner.Spinner {
		assert.Equal(t, "Loading template list ...", text)
		return nil
	}
	getTemplateMetadataListMockable = func() map[string]*model.CGProjectMetadata {
		return templateList
	}
	stopProcess = func(s *spinner.Spinner) {
		assert.Nil(t, s)
	}
	pelCallCount := 0
	pel = func() {
		pelCallCount++
	}
	printHeading = func(text string) {
		assert.Fail(t, "Unexpected call of printHeading")
	}
	printError = func(description string, err error, exit bool) {
		assert.Equal(t, "No templates found. Use \"$ compose-generator save <template-name>\" to save one.", description)
		assert.Nil(t, err)
		assert.True(t, exit)
	}
	// Execute test
	showTemplateList()
	// Assert
	assert.Equal(t, 1, pelCallCount)
}

// ------------------------------------------------------------- getTemplateMetadataList -----------------------------------------------------------

func TestGetTemplateMetatdataList1(t *testing.T) {
	// Test data
	customTemplatesPath := "../../.github/test-files/templates"
	// Mock functions
	readDir = func(dirname string) ([]fs.FileInfo, error) {
		assert.Equal(t, customTemplatesPath, dirname)
		return ioutil.ReadDir(dirname)
	}
	getCustomTemplatesPath = func() string {
		return customTemplatesPath
	}
	printError = func(description string, err error, exit bool) {
		assert.Fail(t, "Unexpected call of printError")
	}
	// Execute test
	result := getTemplateMetadataList()
	// Assert
	assert.Equal(t, "template-1", result["../../.github/test-files/templates/Template 1"].ContainerName)
	assert.Equal(t, "template-2", result["../../.github/test-files/templates/Template 2"].ContainerName)
}

func TestGetTemplateMetatdataList2(t *testing.T) {
	// Test data
	customTemplatesPath := "../templates"
	// Mock functions
	readDir = func(dirname string) ([]fs.FileInfo, error) {
		assert.Equal(t, customTemplatesPath, dirname)
		return nil, errors.New("Error message")
	}
	getCustomTemplatesPath = func() string {
		return customTemplatesPath
	}
	printError = func(description string, err error, exit bool) {
		assert.Equal(t, "Cannot access directory for custom templates", description)
		assert.Equal(t, "Error message", err.Error())
		assert.True(t, exit)
	}
	// Execute test
	result := getTemplateMetadataList()
	// Assert
	assert.Nil(t, result)
}

// ------------------------------------------------------------- copyVolumesFromTemplate -----------------------------------------------------------

func TestCopyVolumesFromTemplate1(t *testing.T) {
	// Test data
	project := &model.CGProject{
		Composition: &spec.Project{
			Services: spec.Services{
				{
					Volumes: []spec.ServiceVolumeConfig{
						{
							Source: "./volume1",
							Type:   spec.VolumeTypeBind,
						},
						{
							Source: "./volume1/volume2",
							Type:   spec.VolumeTypeBind,
						},
					},
				},
				{
					Volumes: []spec.ServiceVolumeConfig{
						{
							Source: "../volume3",
							Type:   spec.VolumeTypeBind,
						},
						{
							Source: "../volume4",
							Type:   spec.VolumeTypeBind,
						},
					},
				},
			},
		},
	}
	sourceDir := "."
	// Mock functions
	absCallCount := 0
	abs = func(path string) (string, error) {
		absCallCount++
		switch absCallCount {
		case 1:
			assert.Equal(t, ".", path)
			return "/usr/lib/compose-generator/templates", nil
		case 2:
			assert.Equal(t, "./volume1", path)
			return "/usr/lib/compose-generator/templates/Template 1/volume1", nil
		case 3:
			assert.Equal(t, "../volume3", path)
			return "/usr/lib/compose-generator/templates/Template 2/volume3", nil
		case 4:
			assert.Equal(t, "../volume4", path)
			return "", errors.New("Error message")
		}
		return "", nil
	}
	rel = func(basepath, targpath string) (string, error) {
		assert.Equal(t, "/usr/lib/compose-generator/templates", basepath)
		switch targpath {
		case "/usr/lib/compose-generator/templates/Template 1/volume1":
			assert.Equal(t, "/usr/lib/compose-generator/templates/Template 1/volume1", targpath)
			return "./Template 1/volume1", nil
		case "/usr/lib/compose-generator/templates/Template 2/volume3":
			assert.Equal(t, "/usr/lib/compose-generator/templates/Template 2/volume3", targpath)
			return "", errors.New("Error message 1")
		}
		return "", nil
	}
	copyDirCallCount := 0
	copyDir = func(src, dest string, opt ...copy.Options) error {
		copyDirCallCount++
		assert.Zero(t, len(opt))
		if copyDirCallCount == 1 {
			assert.Equal(t, "././Template 1/volume1", src)
			assert.Equal(t, "./volume1", dest)
			return nil
		} else {
			assert.Equal(t, "././Template 2/volume3", src)
			assert.Equal(t, "../volume3", dest)
		}
		return errors.New("Error message")
	}
	printErrorCallCount := 0
	printError = func(description string, err error, exit bool) {
		printErrorCallCount++
		if printErrorCallCount == 1 {
			assert.Equal(t, "Could not copy volume '../volume3'", description)
			assert.Equal(t, "Error message 1", err.Error())
			assert.False(t, exit)
		} else {
			assert.Equal(t, "Could not find absolute path of volume dir", description)
			assert.Equal(t, "Error message", err.Error())
			assert.True(t, exit)
		}
	}
	printWarning = func(description string) {
		assert.Equal(t, "Could not copy volumes from '././Template 2/volume3' to '../volume3'", description)
	}
	// Execute test
	copyVolumesAndBuildContextsFromTemplate(project, sourceDir)
	// Assert
	assert.Equal(t, 4, absCallCount)
	assert.Equal(t, 1, copyDirCallCount)
	assert.Equal(t, 2, printErrorCallCount)
}

func TestCopyVolumesFromTemplate2(t *testing.T) {
	// Test data
	project := &model.CGProject{}
	sourceDir := "."
	// Mock functions
	abs = func(path string) (string, error) {
		assert.Equal(t, ".", path)
		return "", errors.New("Error message")
	}
	printError = func(description string, err error, exit bool) {
		assert.Equal(t, "Could not find absolute path of current dir", description)
		assert.Equal(t, "Error message", err.Error())
		assert.True(t, exit)
	}
	// Execute test
	copyVolumesAndBuildContextsFromTemplate(project, sourceDir)
}
