package main

import (
	"log"
	"strings"
	"time"

	"github.com/jinzhu/now"
	"github.com/samber/lo"
	"github.com/spf13/cobra"

	ms "github.com/struckchure/monitoring-spirit"
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Create report",
	Run: func(cmd *cobra.Command, args []string) {
		from, _ := cmd.Flags().GetString("from")
		to, _ := cmd.Flags().GetString("to")
		since, _ := cmd.Flags().GetString("since")
		until, _ := cmd.Flags().GetString("until")
		email, _ := cmd.Flags().GetString("email")

		filter := ms.FilterArgs{
			From: lo.TernaryF(
				lo.IsEmpty(from),
				func() *string { return nil },
				func() *string { return &from },
			),
			To: lo.TernaryF(
				lo.IsEmpty(to),
				func() *string { return nil },
				func() *string { return &to },
			),
			Since: lo.TernaryF(
				lo.IsEmpty(since),
				func() *time.Time { return nil },
				func() *time.Time { return lo.ToPtr(now.MustParse(since)) },
			),
			Until: lo.TernaryF(
				lo.IsEmpty(until),
				func() *time.Time { return nil },
				func() *time.Time { return lo.ToPtr(now.MustParse(until)) },
			),
			Email: lo.TernaryF(
				lo.IsEmpty(email),
				func() *string { return nil },
				func() *string { return &email },
			),
		}

		output, err := msService.Report(filter)
		if err != nil {
			log.Panic(err)
		}

		println(ms.RenderMarkdown(strings.Join(output, "\n")))
	},
}

func init() {
	reportCmd.Flags().String("from", "", "Starting commit hash")
	reportCmd.Flags().String("to", "", "Ending commit hash")
	reportCmd.Flags().String("since", "", "Start date (RFC3339 format)")
	reportCmd.Flags().String("until", "", "End date (RFC3339 format)")
	reportCmd.Flags().String("email", "", "Filter by author email")

	rootCmd.AddCommand(reportCmd)
}
