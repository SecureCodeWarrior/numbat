package scw

import (
	"context"
	"io"
	"io/fs"
	"testing"
	"time"

	bolt "go.etcd.io/bbolt"
)

// TestFacadeSignatures pins the exported surface consumers compile against.
// Each assignment fails to compile when upstream renames a symbol or changes
// a signature, so `go test ./...` in the fork flags drift before a downstream
// module does. It never runs any of the referenced functions.
func TestFacadeSignatures(t *testing.T) {
	t.Helper()
	_ = context.Background

	var (
		_ func(string) (string, error)                                        = ParseAgent
		_ func(string) (Lifecycle, error)                                     = ParseLifecycle
		_ func(string, string) (Lifecycle, error)                             = ResolveLifecycle
		_ func(Lifecycle, string, string, string, map[string]any) []Event     = MapEvents
		_ func(string, string, Lifecycle) bool                                = EnforceEligible
		_ func(string, string) (map[string]any, bool)                         = DenyResponse
		_ func(string, Lifecycle) map[string]any                              = PassthroughResponse
		_ func(string) bool                                                   = DenyUsesExitCode
		_ func() []string                                                     = AgentNames
		_ func() string                                                       = AgentUsage
		_ func(string) bool                                                   = AgentSupportsHooks
		_ func(string) bool                                                   = AgentSupportsEnforcement
		_ func(string) bool                                                   = AgentSupportsManagedHooks
		_ func() []string                                                     = InstallAgentNames
		_ func([]string) []string                                             = SortAgentNames
		_ func(string) (string, bool)                                         = ManagedHooksPath
		_ func(string, string, string, InstallOptions) (InstallReport, error) = InstallWithOptions
		_ func(string, string, string, InstallOptions) (InstallReport, error) = InstallManagedWithOptions
		_ func(string, string) (InstallReport, error)                         = Uninstall
		_ func(string, string) (InstallReport, error)                         = UninstallManaged
		_ func(string, string) (InstallReport, error)                         = StatusErr
		_ func(string, string) (InstallReport, error)                         = StatusManagedErr
		_ func(io.Reader, string) (RuleSource, error)                         = LoadSource
		_ func(fs.FS, string) ([]RuleSource, error)                           = LoadSourcesFS
		_ func(fs.FS, string) (CheckedExpressions, error)                     = LoadCheckedExpressionsFS
		_ func([]RuleSource) (*Engine, error)                                 = NewEngine
		_ func([]RuleSource, CheckedExpressions) (*Engine, error)             = NewEngineWithCheckedExpressions
		_ func(string) bool                                                   = IsRuleTestFile
		_ func(io.Writer, io.Writer, string, ...EmitterOption) *Emitter       = OutputNew
		_ func() EmitterOption                                                = WithFullContent
		_ func(string, time.Duration) (*StateDB, error)                       = StateOpen
		_ func() SequenceConfig                                               = SequenceDefaultConfig
		_ func(Event) Event                                                   = RedactEvent
	)

	// Pipeline construction and the sequence store take engine-derived types.
	var eng *Engine
	_ = func(db *bolt.DB) (*SequenceStore, error) {
		return SequenceNewStore(db, eng.SequenceRules(), SequenceDefaultConfig())
	}
	_ = func(em *Emitter) *Pipeline {
		return PipelineNew(eng, em, Selection{Events: true, Findings: true}, FindingOptions{Now: time.Now()}, nil)
	}

	// Per-agent default paths all share the (home string) string shape.
	for _, fn := range []func(string) string{
		ClaudeSettingsPath, CursorSettingsPath, WindsurfSettingsPath, CopilotSettingsPath, VSCodeHooksPath,
		CodexSettingsPath, GeminiSettingsPath, OpenCodePluginPath, OpenClawPluginPath, AntigravityHooksPath,
		FactorySettingsPath, GrokHooksPath, DevinUserConfigPath, HermesUserConfigPath, PiExtensionPath,
		KimiConfigPath, QwenSettingsPath, ClineHooksPath, AmpPluginPath, AuggieSettingsPath, KiroHooksPath,
		GoosePluginPath, KiloPluginPath, CrushConfigPath, JunieConfigPath,
	} {
		if fn == nil {
			t.Fatal("nil path helper")
		}
	}

	// Constants referenced by consumers.
	_ = []string{
		SchemaVersion,
		EnforcementDecisionNoOverride, EnforcementDecisionDeny,
		EnforcementModeMonitor, EnforcementModeEnforce,
		EnforcementReasonMonitorMode, EnforcementReasonNoEnforceEligibleMatch,
		EnforcementReasonFailOpen, EnforcementReasonEnforceRuleMatch,
		AgentClaude, AgentCodex, AgentGemini, AgentCursor, AgentWindsurf, AgentCopilot, AgentVSCode,
		AgentOpenCode, AgentOpenClaw, AgentAntigravity, AgentFactory, AgentGrok, AgentDevin, AgentHermes,
		AgentPi, AgentKimi, AgentQwen, AgentCline, AgentAmp, AgentAuggie, AgentKiro, AgentGoose, AgentKilo,
		AgentOpenHands, AgentCrush, AgentJunie,
	}
	_ = []Lifecycle{
		LifecycleSessionStart, LifecyclePromptSubmit, LifecyclePreTool, LifecyclePostTool, LifecycleStop,
		LifecycleSessionEnd, LifecycleAssistant, LifecycleCursorPreTool, LifecycleCursorPostTool, LifecycleCodexPreTool,
		LifecycleGeminiPreTool, LifecycleVSCodePreTool, LifecycleCopilotPreTool,
	}
	var _ EventType = EventMessageReasoning
	var _ Rule
	var _ Finding
	var _ EnforcementDecision
	var _ EnforceDecision
}
