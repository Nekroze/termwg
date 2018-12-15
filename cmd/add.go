// Copyright © 2018 Taylor Lawson <nekroze.lawson@gmail.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"strconv"

	"github.com/Nekroze/termwg/twg"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [<DELTA>] [<CHANNEL_NAME>]",
	Short: "Add <DELTA> (defaulting to 1) to the counter of the given <CHANNEL_NAME> or 'default' if not provided",
	Long:  `Add <DELTA> (defaulting to 1) to the counter of the given <CHANNEL_NAME> or 'default' if not provided.`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		delta := int64(1)
		if len(args) > 0 {
			delta, err = strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				panic(err)
			}
		}

		channel := "default"
		if len(args) > 1 {
			channel = args[1]
		}

		twg.WaitGroup{channel}.Add(int(delta))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
