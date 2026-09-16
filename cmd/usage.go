package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/spf13/cobra"
)

var (
	summaryFlag bool
	yearFlag    int
	monthFlag   int
	dayFlag     int
)

var usageCmd = &cobra.Command{
	Use:   "usage",
	Short: "Show Copilot premium request usage summary",
	Long: `Show the Copilot premium request usage summary for the authenticated user.

By default, outputs the full JSON response from the GitHub API.

With the --summary flag, only the total gross quantity is displayed.

Optional flags --year (-y), --month (-m), and --day (-d) allow filtering
billing usage data by specific time periods:
  gh vibe usage --year 2024 --month 11
  gh vibe usage --year 2024 --month 11 --day 15
  gh vibe usage --summary --year 2024 --month 11`,
	Args: cobra.NoArgs,
	RunE: runUsage,
}

func init() {
	rootCmd.AddCommand(usageCmd)
	usageCmd.Flags().BoolVarP(&summaryFlag, "summary", "s", false, "Show total gross quantity instead of full JSON")
	usageCmd.Flags().IntVarP(&yearFlag, "year", "y", 0, "Specify the year (e.g., 2024)")
	usageCmd.Flags().IntVarP(&monthFlag, "month", "m", 0, "Specify the month (e.g., 1-12)")
	usageCmd.Flags().IntVarP(&dayFlag, "day", "d", 0, "Specify the day (e.g., 1-31)")
}

// UsageResponse represents the billing usage response from the GitHub API
type UsageResponse struct {
	TimePeriod TimePeriod  `json:"timePeriod"`
	User       string      `json:"user"`
	UsageItems []UsageItem `json:"usageItems"`
}

// TimePeriod represents the billing period
type TimePeriod struct {
	Year  int `json:"year"`
	Month int `json:"month,omitempty"`
	Day   int `json:"day,omitempty"`
}

// UsageItem represents a single usage item
type UsageItem struct {
	Product          string  `json:"product"`
	SKU              string  `json:"sku"`
	Model            string  `json:"model"`
	UnitType         string  `json:"unitType"`
	PricePerUnit     float64 `json:"pricePerUnit"`
	GrossQuantity    float64 `json:"grossQuantity"`
	GrossAmount      float64 `json:"grossAmount"`
	DiscountQuantity float64 `json:"discountQuantity"`
	DiscountAmount   float64 `json:"discountAmount"`
	NetQuantity      float64 `json:"netQuantity"`
	NetAmount        float64 `json:"netAmount"`
}

// buildUsageEndpoint constructs the API endpoint with optional query parameters.
func buildUsageEndpoint(username string, year, month, day int) string {
	endpoint := fmt.Sprintf("users/%s/settings/billing/ai_credit/usage", username)
	var queryParams []string
	if year > 0 {
		queryParams = append(queryParams, fmt.Sprintf("year=%d", year))
	}
	if month > 0 {
		queryParams = append(queryParams, fmt.Sprintf("month=%d", month))
	}
	if day > 0 {
		queryParams = append(queryParams, fmt.Sprintf("day=%d", day))
	}
	if len(queryParams) > 0 {
		endpoint = fmt.Sprintf("%s?%s", endpoint, strings.Join(queryParams, "&"))
	}
	return endpoint
}

func runUsage(_ *cobra.Command, _ []string) error {
	client, err := api.DefaultRESTClient()
	if err != nil {
		return fmt.Errorf("failed to create REST client: %w", err)
	}

	// Get the authenticated user
	username, err := getAuthenticatedUser(client)
	if err != nil {
		return fmt.Errorf("failed to get authenticated user: %w", err)
	}

	// Fetch the usage data
	var response UsageResponse
	endpoint := buildUsageEndpoint(username, yearFlag, monthFlag, dayFlag)
	err = client.Get(endpoint, &response)
	if err != nil {
		return fmt.Errorf("failed to fetch usage data: %w", err)
	}

	if summaryFlag {
		fmt.Printf("%.1f\n", response.TotalGrossQuantity())
	} else {
		// Output full JSON response
		output, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}
		fmt.Println(string(output))
	}

	return nil
}

// TotalGrossQuantity returns the sum of GrossQuantity across all usage items.
func (r UsageResponse) TotalGrossQuantity() float64 {
	var total float64
	for _, item := range r.UsageItems {
		total += item.GrossQuantity
	}
	return total
}

func getAuthenticatedUser(client *api.RESTClient) (string, error) {
	var user struct {
		Login string `json:"login"`
	}
	err := client.Get("user", &user)
	if err != nil {
		return "", err
	}
	return user.Login, nil
}
