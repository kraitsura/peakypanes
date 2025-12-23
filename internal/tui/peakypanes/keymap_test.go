package peakypanes

import (
	"strings"
	"testing"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/regenrek/peakypanes/internal/layout"
)

func TestBuildDashboardKeyMapOverride(t *testing.T) {
	cfg := layout.DashboardKeymapConfig{
		ProjectLeft: []string{"ctrl+h"},
		SessionUp:   []string{"ctrl+k"},
	}
	km, err := buildDashboardKeyMap(cfg)
	if err != nil {
		t.Fatalf("buildDashboardKeyMap() error: %v", err)
	}
	if !key.Matches(tea.KeyMsg{Type: tea.KeyCtrlH}, km.projectLeft) {
		t.Error("projectLeft binding should match ctrl+h")
	}
	if !key.Matches(tea.KeyMsg{Type: tea.KeyCtrlK}, km.sessionUp) {
		t.Error("sessionUp binding should match ctrl+k")
	}
}

func TestBuildDashboardKeyMapInvalidKey(t *testing.T) {
	cfg := layout.DashboardKeymapConfig{
		ProjectLeft: []string{"ctrl+nope"},
	}
	_, err := buildDashboardKeyMap(cfg)
	if err == nil {
		t.Fatal("expected error for invalid key")
	}
	if !strings.Contains(err.Error(), "dashboard.keymap.project_left") {
		t.Fatalf("expected field name in error, got %q", err.Error())
	}
}

func TestBuildDashboardKeyMapDuplicateKey(t *testing.T) {
	cfg := layout.DashboardKeymapConfig{
		ProjectLeft:  []string{"ctrl+a"},
		ProjectRight: []string{"ctrl+a"},
	}
	_, err := buildDashboardKeyMap(cfg)
	if err == nil {
		t.Fatal("expected error for duplicate key")
	}
	if !strings.Contains(err.Error(), "already bound") {
		t.Fatalf("expected duplicate error, got %q", err.Error())
	}
}
