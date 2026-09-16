package cmd

import (
	"encoding/json"
	"testing"
)

func TestRootCommandExists(t *testing.T) {
	if rootCmd == nil {
		t.Error("rootCmd should not be nil")
	}
}

func TestRootCommandUse(t *testing.T) {
	if rootCmd.Use != "gh-vibe" {
		t.Errorf("Expected rootCmd.Use to be 'gh-vibe', got '%s'", rootCmd.Use)
	}
}

func TestInitCommandExists(t *testing.T) {
	if initCmd == nil {
		t.Error("initCmd should not be nil")
	}
}

func TestInitCommandUse(t *testing.T) {
	if initCmd.Use != "init [owner/repo]" {
		t.Errorf("Expected initCmd.Use to be 'init [owner/repo]', got '%s'", initCmd.Use)
	}
}

func TestInitCommandRegistered(t *testing.T) {
	found := false
	for _, cmd := range rootCmd.Commands() {
		if cmd.Use == "init [owner/repo]" {
			found = true
			break
		}
	}
	if !found {
		t.Error("init command should be registered with root command")
	}
}

func TestInitCommandAcceptsNoArgs(t *testing.T) {
	if initCmd.Args == nil {
		t.Error("initCmd.Args should not be nil")
		return
	}

	err := initCmd.Args(initCmd, []string{})
	if err != nil {
		t.Errorf("initCmd should accept no arguments, got error: %v", err)
	}
}

func TestInitCommandAcceptsOneArg(t *testing.T) {
	if initCmd.Args == nil {
		t.Error("initCmd.Args should not be nil")
		return
	}

	err := initCmd.Args(initCmd, []string{"owner/repo"})
	if err != nil {
		t.Errorf("initCmd should accept one argument, got error: %v", err)
	}
}

func TestInitCommandRejectsTwoArgs(t *testing.T) {
	if initCmd.Args == nil {
		t.Error("initCmd.Args should not be nil")
		return
	}

	err := initCmd.Args(initCmd, []string{"owner/repo", "extra"})
	if err == nil {
		t.Error("initCmd should reject more than one argument")
	}
}

func TestInitCommandHasBranchFlag(t *testing.T) {
	flag := initCmd.Flags().Lookup("branch")
	if flag == nil {
		t.Error("initCmd should have a 'branch' flag")
		return
	}
	if flag.DefValue != "main" {
		t.Errorf("Expected branch flag default to be 'main', got '%s'", flag.DefValue)
	}
	if flag.Shorthand != "b" {
		t.Errorf("Expected branch flag shorthand to be 'b', got '%s'", flag.Shorthand)
	}
}

func TestReadyMergeCommandExists(t *testing.T) {
	if readyMergeCmd == nil {
		t.Error("readyMergeCmd should not be nil")
	}
}

func TestReadyMergeCommandUse(t *testing.T) {
	if readyMergeCmd.Use != "ready-merge [PR number]" {
		t.Errorf("Expected readyMergeCmd.Use to be 'ready-merge [PR number]', got '%s'", readyMergeCmd.Use)
	}
}

func TestReadyMergeCommandRegistered(t *testing.T) {
	found := false
	for _, cmd := range rootCmd.Commands() {
		if cmd.Use == "ready-merge [PR number]" {
			found = true
			break
		}
	}
	if !found {
		t.Error("ready-merge command should be registered with root command")
	}
}

func TestReadyMergeCommandAcceptsNoArgs(t *testing.T) {
	if readyMergeCmd.Args == nil {
		t.Error("readyMergeCmd.Args should not be nil")
		return
	}

	err := readyMergeCmd.Args(readyMergeCmd, []string{})
	if err != nil {
		t.Errorf("readyMergeCmd should accept no arguments, got error: %v", err)
	}
}

func TestReadyMergeCommandAcceptsOneArg(t *testing.T) {
	if readyMergeCmd.Args == nil {
		t.Error("readyMergeCmd.Args should not be nil")
		return
	}

	err := readyMergeCmd.Args(readyMergeCmd, []string{"123"})
	if err != nil {
		t.Errorf("readyMergeCmd should accept one argument, got error: %v", err)
	}
}

func TestReadyMergeCommandRejectsTwoArgs(t *testing.T) {
	if readyMergeCmd.Args == nil {
		t.Error("readyMergeCmd.Args should not be nil")
		return
	}

	err := readyMergeCmd.Args(readyMergeCmd, []string{"123", "456"})
	if err == nil {
		t.Error("readyMergeCmd should reject more than one argument")
	}
}

func TestSetupAICommandExists(t *testing.T) {
	if setupAICmd == nil {
		t.Error("setupAICmd should not be nil")
	}
}

func TestSetupAICommandUse(t *testing.T) {
	if setupAICmd.Use != "setup-ai" {
		t.Errorf("Expected setupAICmd.Use to be 'setup-ai', got '%s'", setupAICmd.Use)
	}
}

func TestSetupAICommandRegistered(t *testing.T) {
	found := false
	for _, cmd := range rootCmd.Commands() {
		if cmd.Use == "setup-ai" {
			found = true
			break
		}
	}
	if !found {
		t.Error("setup-ai command should be registered with root command")
	}
}

func TestSetupAICommandRejectsArgs(t *testing.T) {
	if setupAICmd.Args == nil {
		t.Error("setupAICmd.Args should not be nil")
		return
	}

	err := setupAICmd.Args(setupAICmd, []string{"unexpected-arg"})
	if err == nil {
		t.Error("setupAICmd should reject unexpected arguments")
	}
}

func TestSetupAICommandAcceptsNoArgs(t *testing.T) {
	if setupAICmd.Args == nil {
		t.Error("setupAICmd.Args should not be nil")
		return
	}

	err := setupAICmd.Args(setupAICmd, []string{})
	if err != nil {
		t.Errorf("setupAICmd should accept no arguments, got error: %v", err)
	}
}

func TestUsageCommandExists(t *testing.T) {
	if usageCmd == nil {
		t.Error("usageCmd should not be nil")
	}
}

func TestUsageCommandUse(t *testing.T) {
	if usageCmd.Use != "usage" {
		t.Errorf("Expected usageCmd.Use to be 'usage', got '%s'", usageCmd.Use)
	}
}

func TestUsageCommandRegistered(t *testing.T) {
	found := false
	for _, cmd := range rootCmd.Commands() {
		if cmd.Use == "usage" {
			found = true
			break
		}
	}
	if !found {
		t.Error("usage command should be registered with root command")
	}
}

func TestUsageCommandRejectsArgs(t *testing.T) {
	if usageCmd.Args == nil {
		t.Error("usageCmd.Args should not be nil")
		return
	}

	err := usageCmd.Args(usageCmd, []string{"unexpected-arg"})
	if err == nil {
		t.Error("usageCmd should reject unexpected arguments")
	}
}

func TestUsageCommandAcceptsNoArgs(t *testing.T) {
	if usageCmd.Args == nil {
		t.Error("usageCmd.Args should not be nil")
		return
	}

	err := usageCmd.Args(usageCmd, []string{})
	if err != nil {
		t.Errorf("usageCmd should accept no arguments, got error: %v", err)
	}
}

func TestUsageCommandHasSummaryFlag(t *testing.T) {
	flag := usageCmd.Flags().Lookup("summary")
	if flag == nil {
		t.Error("usageCmd should have a 'summary' flag")
		return
	}
	if flag.DefValue != "false" {
		t.Errorf("Expected summary flag default to be 'false', got '%s'", flag.DefValue)
	}
	if flag.Shorthand != "s" {
		t.Errorf("Expected summary flag shorthand to be 's', got '%s'", flag.Shorthand)
	}
}

func TestUsageResponseUnmarshal(t *testing.T) {
	raw := `{
		"timePeriod": {"year": 2025, "month": 7},
		"user": "octocat",
		"usageItems": [
			{
				"product": "copilot",
				"sku": "premium",
				"model": "gpt-4",
				"unitType": "tokens",
				"pricePerUnit": 0.01,
				"grossQuantity": 100.5,
				"grossAmount": 1.005,
				"discountQuantity": 10.0,
				"discountAmount": 0.1,
				"netQuantity": 90.5,
				"netAmount": 0.905
			}
		]
	}`

	var resp UsageResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		t.Fatalf("Failed to unmarshal UsageResponse: %v", err)
	}

	if resp.User != "octocat" {
		t.Errorf("Expected user 'octocat', got '%s'", resp.User)
	}
	if resp.TimePeriod.Year != 2025 {
		t.Errorf("Expected year 2025, got %d", resp.TimePeriod.Year)
	}
	if resp.TimePeriod.Month != 7 {
		t.Errorf("Expected month 7, got %d", resp.TimePeriod.Month)
	}
	if len(resp.UsageItems) != 1 {
		t.Fatalf("Expected 1 usage item, got %d", len(resp.UsageItems))
	}
	item := resp.UsageItems[0]
	if item.GrossQuantity != 100.5 {
		t.Errorf("Expected GrossQuantity 100.5, got %f", item.GrossQuantity)
	}
	if item.NetQuantity != 90.5 {
		t.Errorf("Expected NetQuantity 90.5, got %f", item.NetQuantity)
	}
}

func TestUsageResponseTotalGrossQuantity(t *testing.T) {
	resp := UsageResponse{
		UsageItems: []UsageItem{
			{GrossQuantity: 10.0},
			{GrossQuantity: 20.5},
			{GrossQuantity: 5.25},
		},
	}

	expected := 35.75
	if got := resp.TotalGrossQuantity(); got != expected {
		t.Errorf("Expected TotalGrossQuantity %.2f, got %.2f", expected, got)
	}
}

func TestUsageCommandHasYearFlag(t *testing.T) {
	flag := usageCmd.Flags().Lookup("year")
	if flag == nil {
		t.Error("usageCmd should have a 'year' flag")
		return
	}
	if flag.DefValue != "0" {
		t.Errorf("Expected year flag default to be '0', got '%s'", flag.DefValue)
	}
	if flag.Shorthand != "y" {
		t.Errorf("Expected year flag shorthand to be 'y', got '%s'", flag.Shorthand)
	}
}

func TestUsageCommandHasMonthFlag(t *testing.T) {
	flag := usageCmd.Flags().Lookup("month")
	if flag == nil {
		t.Error("usageCmd should have a 'month' flag")
		return
	}
	if flag.DefValue != "0" {
		t.Errorf("Expected month flag default to be '0', got '%s'", flag.DefValue)
	}
	if flag.Shorthand != "m" {
		t.Errorf("Expected month flag shorthand to be 'm', got '%s'", flag.Shorthand)
	}
}

func TestUsageCommandHasDayFlag(t *testing.T) {
	flag := usageCmd.Flags().Lookup("day")
	if flag == nil {
		t.Error("usageCmd should have a 'day' flag")
		return
	}
	if flag.DefValue != "0" {
		t.Errorf("Expected day flag default to be '0', got '%s'", flag.DefValue)
	}
	if flag.Shorthand != "d" {
		t.Errorf("Expected day flag shorthand to be 'd', got '%s'", flag.Shorthand)
	}
}

func TestBuildUsageEndpoint(t *testing.T) {
	tests := []struct {
		name     string
		username string
		year     int
		month    int
		day      int
		expected string
	}{
		{
			name:     "no parameters",
			username: "octocat",
			year:     0,
			month:    0,
			day:      0,
			expected: "users/octocat/settings/billing/ai_credit/usage",
		},
		{
			name:     "year only",
			username: "octocat",
			year:     2024,
			month:    0,
			day:      0,
			expected: "users/octocat/settings/billing/ai_credit/usage?year=2024",
		},
		{
			name:     "month only",
			username: "octocat",
			year:     0,
			month:    11,
			day:      0,
			expected: "users/octocat/settings/billing/ai_credit/usage?month=11",
		},
		{
			name:     "day only",
			username: "octocat",
			year:     0,
			month:    0,
			day:      15,
			expected: "users/octocat/settings/billing/ai_credit/usage?day=15",
		},
		{
			name:     "year and month",
			username: "octocat",
			year:     2024,
			month:    11,
			day:      0,
			expected: "users/octocat/settings/billing/ai_credit/usage?year=2024&month=11",
		},
		{
			name:     "year and day",
			username: "octocat",
			year:     2024,
			month:    0,
			day:      15,
			expected: "users/octocat/settings/billing/ai_credit/usage?year=2024&day=15",
		},
		{
			name:     "month and day",
			username: "octocat",
			year:     0,
			month:    11,
			day:      15,
			expected: "users/octocat/settings/billing/ai_credit/usage?month=11&day=15",
		},
		{
			name:     "year, month, and day",
			username: "octocat",
			year:     2024,
			month:    11,
			day:      15,
			expected: "users/octocat/settings/billing/ai_credit/usage?year=2024&month=11&day=15",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := buildUsageEndpoint(tc.username, tc.year, tc.month, tc.day)
			if result != tc.expected {
				t.Errorf("Expected endpoint %q, got %q", tc.expected, result)
			}
		})
	}
}

func TestUsageResponseUnmarshalWithDay(t *testing.T) {
	raw := `{
		"timePeriod": {"year": 2025, "month": 7, "day": 15},
		"user": "octocat",
		"usageItems": [
			{
				"product": "copilot",
				"sku": "premium",
				"model": "gpt-4",
				"unitType": "tokens",
				"pricePerUnit": 0.01,
				"grossQuantity": 100.5,
				"grossAmount": 1.005,
				"discountQuantity": 10.0,
				"discountAmount": 0.1,
				"netQuantity": 90.5,
				"netAmount": 0.905
			}
		]
	}`

	var resp UsageResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		t.Fatalf("Failed to unmarshal UsageResponse: %v", err)
	}

	if resp.TimePeriod.Year != 2025 {
		t.Errorf("Expected year 2025, got %d", resp.TimePeriod.Year)
	}
	if resp.TimePeriod.Month != 7 {
		t.Errorf("Expected month 7, got %d", resp.TimePeriod.Month)
	}
	if resp.TimePeriod.Day != 15 {
		t.Errorf("Expected day 15, got %d", resp.TimePeriod.Day)
	}
}

func TestUsageFlagsParsing(t *testing.T) {
	// Save initial flag values
	origSummary := summaryFlag
	origYear := yearFlag
	origMonth := monthFlag
	origDay := dayFlag
	defer func() {
		summaryFlag = origSummary
		yearFlag = origYear
		monthFlag = origMonth
		dayFlag = origDay
	}()

	// Test long flag names
	err := usageCmd.ParseFlags([]string{"--year", "2024", "--month", "11", "--day", "15", "--summary"})
	if err != nil {
		t.Fatalf("Failed to parse flags: %v", err)
	}
	if yearFlag != 2024 {
		t.Errorf("Expected yearFlag 2024, got %d", yearFlag)
	}
	if monthFlag != 11 {
		t.Errorf("Expected monthFlag 11, got %d", monthFlag)
	}
	if dayFlag != 15 {
		t.Errorf("Expected dayFlag 15, got %d", dayFlag)
	}
	if !summaryFlag {
		t.Error("Expected summaryFlag to be true")
	}

	// Test shorthand flag names
	err = usageCmd.ParseFlags([]string{"-y", "2023", "-m", "5", "-d", "1", "-s"})
	if err != nil {
		t.Fatalf("Failed to parse shorthand flags: %v", err)
	}
	if yearFlag != 2023 {
		t.Errorf("Expected yearFlag 2023, got %d", yearFlag)
	}
	if monthFlag != 5 {
		t.Errorf("Expected monthFlag 5, got %d", monthFlag)
	}
	if dayFlag != 1 {
		t.Errorf("Expected dayFlag 1, got %d", dayFlag)
	}
	if !summaryFlag {
		t.Error("Expected summaryFlag to be true")
	}
}
