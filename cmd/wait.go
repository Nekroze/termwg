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
	"github.com/Nekroze/termwg/twg"
	"github.com/spf13/cobra"
)

var waitCmd = &cobra.Command{
	Use:   "wait [<CHANNEL_NAME>]",
	Short: "wait for a the given channel name (or 'default') counter to reach 0",
	Long: `wait for a the given channel name (or 'default') counter to reach 0.
If the channel counter is currently 0, wait will add 1 to the couter for convenience.`,
	Run: func(cmd *cobra.Command, args []string) {
		channel := defaultChannel
		if len(args) > 0 {
			channel = args[0]
		}
		twg.WaitGroup{
			Name: channel,
		}.Wait()
	},
}

func init() {
	rootCmd.AddCommand(waitCmd)
}
