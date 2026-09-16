// Package scw is the SecureCodeWarrior consumption facade for numbat.
//
// It contains only type aliases, constant re-exports, and function re-exports
// of symbols from numbat's internal packages so an external module can embed
// numbat's hook capture, rule engine, pipeline, and installers as a library.
// It deliberately carries no logic: every behavior lives either upstream or in
// the consuming module. Upstream renames surface here as compile errors.
package scw

import (
	"github.com/perplexityai/numbat/internal/finding"
	"github.com/perplexityai/numbat/internal/hook"
	"github.com/perplexityai/numbat/internal/model"
	"github.com/perplexityai/numbat/internal/output"
	"github.com/perplexityai/numbat/internal/pipeline"
	"github.com/perplexityai/numbat/internal/redact"
	"github.com/perplexityai/numbat/internal/rule"
	"github.com/perplexityai/numbat/internal/sequence"
	"github.com/perplexityai/numbat/internal/state"
)

// model

type (
	Event               = model.Event
	Finding             = model.Finding
	EnforcementDecision = model.EnforcementDecision
	EventType           = model.EventType
)

const (
	SchemaVersion = model.SchemaVersion

	EventSessionStart        = model.EventSessionStart
	EventSessionEnd          = model.EventSessionEnd
	EventPromptUser          = model.EventPromptUser
	EventMessageAssistant    = model.EventMessageAssistant
	EventToolCall            = model.EventToolCall
	EventToolResult          = model.EventToolResult
	EventCommandExec         = model.EventCommandExec
	EventCommandResult       = model.EventCommandResult
	EventFileRead            = model.EventFileRead
	EventFileWrite           = model.EventFileWrite
	EventFileDelete          = model.EventFileDelete
	EventPermissionRequested = model.EventPermissionRequested
	EventPermissionApproved  = model.EventPermissionApproved
	EventPermissionDenied    = model.EventPermissionDenied
	EventConfigAgent         = model.EventConfigAgent
	EventConfigMCP           = model.EventConfigMCP
	EventNetworkIndicator    = model.EventNetworkIndicator
	EventMessageReasoning    = model.EventMessageReasoning

	EnforcementDecisionNoOverride = model.EnforcementDecisionNoOverride
	EnforcementDecisionDeny       = model.EnforcementDecisionDeny

	EnforcementModeMonitor = model.EnforcementModeMonitor
	EnforcementModeEnforce = model.EnforcementModeEnforce

	EnforcementReasonMonitorMode            = model.EnforcementReasonMonitorMode
	EnforcementReasonNoEnforceEligibleMatch = model.EnforcementReasonNoEnforceEligibleMatch
	EnforcementReasonFailOpen               = model.EnforcementReasonFailOpen
	EnforcementReasonEnforceRuleMatch       = model.EnforcementReasonEnforceRuleMatch
)

// hook: live capture and installers

type (
	Lifecycle      = hook.Lifecycle
	InstallOptions = hook.InstallOptions
	InstallReport  = hook.InstallReport
)

const (
	AgentClaude      = hook.AgentClaude
	AgentCodex       = hook.AgentCodex
	AgentGemini      = hook.AgentGemini
	AgentCursor      = hook.AgentCursor
	AgentWindsurf    = hook.AgentWindsurf
	AgentCopilot     = hook.AgentCopilot
	AgentVSCode      = hook.AgentVSCode
	AgentOpenCode    = hook.AgentOpenCode
	AgentOpenClaw    = hook.AgentOpenClaw
	AgentAntigravity = hook.AgentAntigravity
	AgentFactory     = hook.AgentFactory
	AgentGrok        = hook.AgentGrok
	AgentDevin       = hook.AgentDevin
	AgentHermes      = hook.AgentHermes
	AgentPi          = hook.AgentPi
	AgentKimi        = hook.AgentKimi
	AgentQwen        = hook.AgentQwen
	AgentCline       = hook.AgentCline
	AgentAmp         = hook.AgentAmp
	AgentAuggie      = hook.AgentAuggie
	AgentKiro        = hook.AgentKiro
	AgentGoose       = hook.AgentGoose
	AgentKilo        = hook.AgentKilo
	AgentOpenHands   = hook.AgentOpenHands
	AgentCrush       = hook.AgentCrush
	AgentJunie       = hook.AgentJunie

	LifecycleSessionStart     = hook.LifecycleSessionStart
	LifecyclePromptSubmit     = hook.LifecyclePromptSubmit
	LifecyclePreTool          = hook.LifecyclePreTool
	LifecyclePostTool         = hook.LifecyclePostTool
	LifecyclePermission       = hook.LifecyclePermission
	LifecyclePermissionDenied = hook.LifecyclePermissionDenied
	LifecycleStop             = hook.LifecycleStop
	LifecycleSessionEnd       = hook.LifecycleSessionEnd
	LifecycleFileRead         = hook.LifecycleFileRead
	LifecycleFileWrite        = hook.LifecycleFileWrite
	LifecycleMCPCall          = hook.LifecycleMCPCall
	LifecycleCommandExec      = hook.LifecycleCommandExec
	LifecycleCommandResult    = hook.LifecycleCommandResult
	LifecycleAssistant        = hook.LifecycleAssistant

	LifecycleCopilotPreTool         = hook.LifecycleCopilotPreTool
	LifecycleCopilotPostTool        = hook.LifecycleCopilotPostTool
	LifecycleCopilotPermission      = hook.LifecycleCopilotPermission
	LifecycleCursorPreTool          = hook.LifecycleCursorPreTool
	LifecycleCursorPostTool         = hook.LifecycleCursorPostTool
	LifecycleCursorPostToolFailure  = hook.LifecycleCursorPostToolFailure
	LifecycleVSCodePreTool          = hook.LifecycleVSCodePreTool
	LifecycleVSCodePostTool         = hook.LifecycleVSCodePostTool
	LifecycleCodexPreTool           = hook.LifecycleCodexPreTool
	LifecycleCodexPostTool          = hook.LifecycleCodexPostTool
	LifecycleCodexPermission        = hook.LifecycleCodexPermission
	LifecycleGeminiPreTool          = hook.LifecycleGeminiPreTool
	LifecycleGeminiPostTool         = hook.LifecycleGeminiPostTool
	LifecycleOpenCodePreTool        = hook.LifecycleOpenCodePreTool
	LifecycleOpenCodePostTool       = hook.LifecycleOpenCodePostTool
	LifecycleAntigravityPreTool     = hook.LifecycleAntigravityPreTool
	LifecycleAntigravityPostTool    = hook.LifecycleAntigravityPostTool
	LifecycleFactoryPreTool         = hook.LifecycleFactoryPreTool
	LifecycleFactoryPostTool        = hook.LifecycleFactoryPostTool
	LifecycleGrokPreTool            = hook.LifecycleGrokPreTool
	LifecycleGrokPostTool           = hook.LifecycleGrokPostTool
	LifecycleDevinPreTool           = hook.LifecycleDevinPreTool
	LifecycleDevinPostTool          = hook.LifecycleDevinPostTool
	LifecycleDevinPermission        = hook.LifecycleDevinPermission
	LifecycleHermesPreTool          = hook.LifecycleHermesPreTool
	LifecycleHermesPostTool         = hook.LifecycleHermesPostTool
	LifecycleHermesPermission       = hook.LifecycleHermesPermission
	LifecycleHermesPermissionResult = hook.LifecycleHermesPermissionResult
)

var (
	ParseAgent                = hook.ParseAgent
	ParseLifecycle            = hook.ParseLifecycle
	ResolveLifecycle          = hook.ResolveLifecycle
	MapEvents                 = hook.MapEvents
	EnforceEligible           = hook.EnforceEligible
	DenyResponse              = hook.DenyResponse
	PassthroughResponse       = hook.PassthroughResponse
	DenyUsesExitCode          = hook.DenyUsesExitCode
	AgentNames                = hook.AgentNames
	AgentUsage                = hook.AgentUsage
	AgentSupportsHooks        = hook.AgentSupportsHooks
	AgentSupportsEnforcement  = hook.AgentSupportsEnforcement
	AgentSupportsManagedHooks = hook.AgentSupportsManagedHooks
	InstallAgentNames         = hook.InstallAgentNames
	SortAgentNames            = hook.SortAgentNames
	ManagedHooksPath          = hook.ManagedHooksPath
	InstallWithOptions        = hook.InstallWithOptions
	InstallManagedWithOptions = hook.InstallManagedWithOptions
	Uninstall                 = hook.Uninstall
	UninstallManaged          = hook.UninstallManaged
	StatusErr                 = hook.StatusErr
	StatusManagedErr          = hook.StatusManagedErr

	// Per-agent default install paths (the set cmd/numbat's settingsPathFor uses).
	ClaudeSettingsPath   = hook.ClaudeSettingsPath
	CursorSettingsPath   = hook.CursorSettingsPath
	WindsurfSettingsPath = hook.WindsurfSettingsPath
	CopilotSettingsPath  = hook.CopilotSettingsPath
	VSCodeHooksPath      = hook.VSCodeHooksPath
	CodexSettingsPath    = hook.CodexSettingsPath
	GeminiSettingsPath   = hook.GeminiSettingsPath
	OpenCodePluginPath   = hook.OpenCodePluginPath
	OpenClawPluginPath   = hook.OpenClawPluginPath
	AntigravityHooksPath = hook.AntigravityHooksPath
	FactorySettingsPath  = hook.FactorySettingsPath
	GrokHooksPath        = hook.GrokHooksPath
	DevinUserConfigPath  = hook.DevinUserConfigPath
	HermesUserConfigPath = hook.HermesUserConfigPath
	PiExtensionPath      = hook.PiExtensionPath
	KimiConfigPath       = hook.KimiConfigPath
	QwenSettingsPath     = hook.QwenSettingsPath
	ClineHooksPath       = hook.ClineHooksPath
	AmpPluginPath        = hook.AmpPluginPath
	AuggieSettingsPath   = hook.AuggieSettingsPath
	KiroHooksPath        = hook.KiroHooksPath
	GoosePluginPath      = hook.GoosePluginPath
	KiloPluginPath       = hook.KiloPluginPath
	CrushConfigPath      = hook.CrushConfigPath
	JunieConfigPath      = hook.JunieConfigPath
)

// rules, engine, pipeline, output, state

type (
	Rule               = rule.Rule
	RuleSource         = rule.Source
	Engine             = rule.Engine
	CheckedExpressions = rule.CheckedExpressions
	Pipeline           = pipeline.Pipeline
	Selection          = pipeline.Selection
	EnforceDecision    = pipeline.EnforceDecision
	FindingOptions     = finding.Options
	Emitter            = output.Emitter
	EmitterOption      = output.EmitterOption
	StateDB            = state.DB
	SequenceStore      = sequence.Store
	SequenceConfig     = sequence.Config
)

var (
	LoadSource                      = rule.LoadSource
	LoadSourcesFS                   = rule.LoadSourcesFS
	LoadCheckedExpressionsFS        = rule.LoadCheckedExpressionsFS
	NewEngine                       = rule.NewEngine
	NewEngineWithCheckedExpressions = rule.NewEngineWithCheckedExpressions
	IsRuleTestFile                  = rule.IsRuleTestFile
	PipelineNew                     = pipeline.New
	OutputNew                       = output.New
	WithFullContent                 = output.WithFullContent
	StateOpen                       = state.Open
	SequenceNewStore                = sequence.NewStore
	SequenceDefaultConfig           = sequence.DefaultConfig
	RedactEvent                     = redact.Event
)
