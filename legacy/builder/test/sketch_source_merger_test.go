/*
 * This file is part of Arduino Builder.
 *
 * Arduino Builder is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA  02110-1301  USA
 *
 * As a special exception, you may use this file as part of a free software
 * library without restriction.  Specifically, if other files instantiate
 * templates or use macros or inline functions from this file, or you compile
 * this file and link it with other files to produce an executable, this
 * file does not by itself cause the resulting executable to be covered by
 * the GNU General Public License.  This exception does not however
 * invalidate any other reasons why the executable file might be covered by
 * the GNU General Public License.
 *
 * Copyright 2015 Arduino LLC (http://www.arduino.cc/)
 * Copyright 2015 Matthijs Kooijman
 */

package test

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/arduino/arduino-cli/legacy/builder"
	"github.com/arduino/arduino-cli/legacy/builder/types"
	"github.com/arduino/go-paths-helper"
	"github.com/stretchr/testify/require"
)

func TestMergeSketch(t *testing.T) {
	ctx := &types.Context{
		SketchLocation: paths.New("sketch1", "sketch.ino"),
	}

	commands := []types.Command{
		&builder.SketchLoader{},
		&builder.SketchSourceMerger{},
	}

	for _, command := range commands {
		err := command.Run(ctx)
		NoError(t, err)
	}

	source := ctx.Source

	expected_source := LoadAndInterpolate(t, filepath.Join("sketch1", "merged_sketch.txt"), ctx)
	require.Equal(t, expected_source, strings.Replace(source, "\r\n", "\n", -1))
}
