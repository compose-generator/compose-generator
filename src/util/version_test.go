/*
Copyright © 2021-2023 Compose Generator Contributors
All rights reserved.
*/

package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsDevVersion(t *testing.T) {
	result := IsDevVersion()
	if Version == "dev" {
		assert.True(t, result)
	} else {
		assert.False(t, result)
	}
}

func TestIsPreRelease(t *testing.T) {
	result := IsPreRelease()
	assert.False(t, result)
}

// -------------------------------------------- BuildVersion -------------------------------------------

func TestBuildVersion(t *testing.T) {
	for name, tt := range map[string]struct {
		version, commit, date, builtBy string
		out                            string
	}{
		"all empty": {
			out: "",
		},
		"complete": {
			version: "1.2.3",
			date:    "12/12/12",
			commit:  "aaaa",
			builtBy: "me",
			out:     "1.2.3\ncommit: aaaa\nbuilt at: 12/12/12\nbuilt by: me",
		},
		"only version": {
			version: "1.2.3",
			out:     "1.2.3",
		},
		"version and date": {
			version: "1.2.3",
			date:    "12/12/12",
			out:     "1.2.3\nbuilt at: 12/12/12",
		},
		"version, date, built by": {
			version: "1.2.3",
			date:    "12/12/12",
			builtBy: "me",
			out:     "1.2.3\nbuilt at: 12/12/12\nbuilt by: me",
		},
	} {
		tt := tt
		t.Run(name, func(t *testing.T) {
			require.Equal(t, tt.out, BuildVersion(tt.version, tt.commit, tt.date, tt.builtBy))
		})
	}
}
