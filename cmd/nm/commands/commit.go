/*
*	Copyright (C) 2022  Kendall Tauser
*
*	This program is free software; you can redistribute it and/or modify
*	it under the terms of the GNU General Public License as published by
*	the Free Software Foundation; either version 2 of the License, or
*	(at your option) any later version.
*
*	This program is distributed in the hope that it will be useful,
*	but WITHOUT ANY WARRANTY; without even the implied warranty of
*	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
*	GNU General Public License for more details.
*
*	You should have received a copy of the GNU General Public License along
*	with this program; if not, write to the Free Software Foundation, Inc.,
*	51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
 */

package commands

import (
	"github.com/fire833/notemanager/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type commitCmdOpts struct {
	debug *bool
}

var (
	commitCmd *cobra.Command = &cobra.Command{
		Use:     "commit",
		Aliases: []string{"c"},
		Short:   "Generate a commit for the latest changes to your notes context",
		Version: pkg.VersionStr,
		Long: `
`,
		SuggestFor: []string{},
		RunE:       commitExec,
	}

	commitOpts *commitCmdOpts
)

func commitExec(cmd *cobra.Command, args []string) error {
	return nil
}

func init() {
	set := pflag.NewFlagSet(commitCmd.Short, pflag.ExitOnError)

	o := &newCmdOpts{
		debug: set.BoolP("debug", "d", false, "Specify whether to run this command in debug mode. This means outputs should not be used, but analysed for debugging purposes."),
	}

	NMCmd.Flags().AddFlagSet(set)
	newOpts = o
}
