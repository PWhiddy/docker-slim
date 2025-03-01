package profile

import (
	"github.com/docker-slim/docker-slim/pkg/app/master/commands"

	"github.com/c-bata/go-prompt"
)

var CommandSuggestion = prompt.Suggest{
	Text:        Name,
	Description: Usage,
}

var CommandFlagSuggestions = &commands.FlagSuggestions{
	Names: []prompt.Suggest{
		{Text: commands.FullFlagName(commands.FlagTarget), Description: commands.FlagTargetUsage},
		{Text: commands.FullFlagName(commands.FlagPull), Description: commands.FlagPullUsage},
		{Text: commands.FullFlagName(commands.FlagShowPullLogs), Description: commands.FlagShowPullLogsUsage},
		{Text: commands.FullFlagName(commands.FlagShowContainerLogs), Description: commands.FlagShowContainerLogsUsage},
		{Text: commands.FullFlagName(commands.FlagCRORuntime), Description: commands.FlagCRORuntimeUsage},
		{Text: commands.FullFlagName(commands.FlagCROHostConfigFile), Description: commands.FlagCROHostConfigFileUsage},
		{Text: commands.FullFlagName(commands.FlagCROSysctl), Description: commands.FlagCROSysctlUsage},
		{Text: commands.FullFlagName(commands.FlagCROShmSize), Description: commands.FlagCROShmSizeUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPProbeOff), Description: commands.FlagHTTPProbeOffUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPProbe), Description: commands.FlagHTTPProbeUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPProbeCmd), Description: commands.FlagHTTPProbeCmdUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPProbeCmdFile), Description: commands.FlagHTTPProbeCmdFileUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPProbeStartWait), Description: commands.FlagHTTPProbeStartWaitUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPProbeRetryCount), Description: commands.FlagHTTPProbeRetryCountUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPProbeRetryWait), Description: commands.FlagHTTPProbeRetryWaitUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPProbePorts), Description: commands.FlagHTTPProbePortsUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPProbeFull), Description: commands.FlagHTTPProbeFullUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPProbeExitOnFailure), Description: commands.FlagHTTPProbeExitOnFailureUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPProbeCrawl), Description: commands.FlagHTTPProbeCrawlUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPCrawlMaxDepth), Description: commands.FlagHTTPCrawlMaxDepthUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPCrawlMaxPageCount), Description: commands.FlagHTTPCrawlMaxPageCountUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPCrawlConcurrency), Description: commands.FlagHTTPCrawlConcurrencyUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPMaxConcurrentCrawlers), Description: commands.FlagHTTPMaxConcurrentCrawlersUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPProbeAPISpec), Description: commands.FlagHTTPProbeAPISpecUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPProbeAPISpecFile), Description: commands.FlagHTTPProbeAPISpecFileUsage},
		{Text: commands.FullFlagName(commands.FlagPublishPort), Description: commands.FlagPublishPortUsage},
		{Text: commands.FullFlagName(commands.FlagPublishExposedPorts), Description: commands.FlagPublishExposedPortsUsage},
		{Text: commands.FullFlagName(commands.FlagHostExec), Description: commands.FlagHostExecUsage},
		{Text: commands.FullFlagName(commands.FlagHostExecFile), Description: commands.FlagHostExecFileUsage},
		//{Text: commands.FullFlagName(commands.FlagKeepPerms), Description: commands.FlagKeepPermsUsage},
		{Text: commands.FullFlagName(commands.FlagRunTargetAsUser), Description: commands.FlagRunTargetAsUserUsage},
		{Text: commands.FullFlagName(commands.FlagCopyMetaArtifacts), Description: commands.FlagCopyMetaArtifactsUsage},
		{Text: commands.FullFlagName(commands.FlagRemoveFileArtifacts), Description: commands.FlagRemoveFileArtifactsUsage},
		{Text: commands.FullFlagName(commands.FlagUser), Description: commands.FlagUserUsage},
		{Text: commands.FullFlagName(commands.FlagEntrypoint), Description: commands.FlagEntrypointUsage},
		{Text: commands.FullFlagName(commands.FlagCmd), Description: commands.FlagCmdUsage},
		{Text: commands.FullFlagName(commands.FlagWorkdir), Description: commands.FlagWorkdirUsage},
		{Text: commands.FullFlagName(commands.FlagEnv), Description: commands.FlagEnvUsage},
		{Text: commands.FullFlagName(commands.FlagLabel), Description: commands.FlagLabelUsage},
		{Text: commands.FullFlagName(commands.FlagVolume), Description: commands.FlagVolumeUsage},
		{Text: commands.FullFlagName(commands.FlagLink), Description: commands.FlagLinkUsage},
		{Text: commands.FullFlagName(commands.FlagEtcHostsMap), Description: commands.FlagEtcHostsMapUsage},
		{Text: commands.FullFlagName(commands.FlagContainerDNS), Description: commands.FlagContainerDNSUsage},
		{Text: commands.FullFlagName(commands.FlagContainerDNSSearch), Description: commands.FlagContainerDNSSearchUsage},
		{Text: commands.FullFlagName(commands.FlagNetwork), Description: commands.FlagNetworkUsage},
		{Text: commands.FullFlagName(commands.FlagHostname), Description: commands.FlagHostnameUsage},
		{Text: commands.FullFlagName(commands.FlagExpose), Description: commands.FlagExposeUsage},
		{Text: commands.FullFlagName(commands.FlagExcludeMounts), Description: commands.FlagExcludeMountsUsage},
		{Text: commands.FullFlagName(commands.FlagExcludePattern), Description: commands.FlagExcludePatternUsage},
		{Text: commands.FullFlagName(commands.FlagMount), Description: commands.FlagMountUsage},
		{Text: commands.FullFlagName(commands.FlagContinueAfter), Description: commands.FlagContinueAfterUsage},
		{Text: commands.FullFlagName(commands.FlagUseLocalMounts), Description: commands.FlagUseLocalMountsUsage},
		{Text: commands.FullFlagName(commands.FlagUseSensorVolume), Description: commands.FlagUseSensorVolumeUsage},
		{Text: commands.FullFlagName(commands.FlagSensorIPCMode), Description: commands.FlagSensorIPCModeUsage},
		{Text: commands.FullFlagName(commands.FlagSensorIPCEndpoint), Description: commands.FlagSensorIPCEndpointUsage},
	},
	Values: map[string]commands.CompleteValue{
		commands.FullFlagName(commands.FlagPull):                   commands.CompleteBool,
		commands.FullFlagName(commands.FlagShowPullLogs):           commands.CompleteBool,
		commands.FullFlagName(commands.FlagTarget):                 commands.CompleteTarget,
		commands.FullFlagName(commands.FlagShowContainerLogs):      commands.CompleteBool,
		commands.FullFlagName(commands.FlagPublishExposedPorts):    commands.CompleteBool,
		commands.FullFlagName(commands.FlagHTTPProbeOff):           commands.CompleteBool,
		commands.FullFlagName(commands.FlagHTTPProbe):              commands.CompleteTBool,
		commands.FullFlagName(commands.FlagHTTPProbeCmdFile):       commands.CompleteFile,
		commands.FullFlagName(commands.FlagHTTPProbeFull):          commands.CompleteBool,
		commands.FullFlagName(commands.FlagHTTPProbeExitOnFailure): commands.CompleteTBool,
		commands.FullFlagName(commands.FlagHTTPProbeCrawl):         commands.CompleteTBool,
		commands.FullFlagName(commands.FlagHTTPProbeAPISpecFile):   commands.CompleteFile,
		commands.FullFlagName(commands.FlagHostExecFile):           commands.CompleteFile,
		//commands.FullFlagName(commands.FlagKeepPerms):              commands.CompleteTBool,
		commands.FullFlagName(commands.FlagRunTargetAsUser):     commands.CompleteTBool,
		commands.FullFlagName(commands.FlagRemoveFileArtifacts): commands.CompleteBool,
		commands.FullFlagName(commands.FlagNetwork):             commands.CompleteNetwork,
		commands.FullFlagName(commands.FlagExcludeMounts):       commands.CompleteTBool,
		//commands.FullFlagName(commands.FlagPathPermsFile):          commands.CompleteFile,
		//commands.FullFlagName(commands.FlagIncludePathFile):        commands.CompleteFile,
		//commands.FullFlagName(commands.FlagIncludeShell):           commands.CompleteBool,
		commands.FullFlagName(commands.FlagContinueAfter):   commands.CompleteContinueAfter,
		commands.FullFlagName(commands.FlagUseLocalMounts):  commands.CompleteBool,
		commands.FullFlagName(commands.FlagUseSensorVolume): commands.CompleteVolume,
		//commands.FullFlagName(commands.FlagKeepTmpArtifacts):       commands.CompleteBool,
		commands.FullFlagName(commands.FlagCROHostConfigFile): commands.CompleteFile,
		commands.FullFlagName(commands.FlagSensorIPCMode):     commands.CompleteIPCMode,
	},
}
