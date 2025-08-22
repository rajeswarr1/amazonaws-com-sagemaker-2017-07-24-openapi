package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// CreatePresignedNotebookInstanceUrlOutput represents the CreatePresignedNotebookInstanceUrlOutput schema from the OpenAPI specification
type CreatePresignedNotebookInstanceUrlOutput struct {
	Authorizedurl interface{} `json:"AuthorizedUrl,omitempty"`
}

// DescribeTrialResponse represents the DescribeTrialResponse schema from the OpenAPI specification
type DescribeTrialResponse struct {
	Trialname interface{} `json:"TrialName,omitempty"`
	Metadataproperties MetadataProperties `json:"MetadataProperties,omitempty"` // Metadata properties of the tracking entity, trial, or trial component.
	Source interface{} `json:"Source,omitempty"`
	Displayname interface{} `json:"DisplayName,omitempty"`
	Experimentname interface{} `json:"ExperimentName,omitempty"`
	Createdby interface{} `json:"CreatedBy,omitempty"`
	Lastmodifiedby interface{} `json:"LastModifiedBy,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Trialarn interface{} `json:"TrialArn,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
}

// AlgorithmValidationProfile represents the AlgorithmValidationProfile schema from the OpenAPI specification
type AlgorithmValidationProfile struct {
	Trainingjobdefinition interface{} `json:"TrainingJobDefinition"`
	Transformjobdefinition interface{} `json:"TransformJobDefinition,omitempty"`
	Profilename interface{} `json:"ProfileName"`
}

// GetSearchSuggestionsResponse represents the GetSearchSuggestionsResponse schema from the OpenAPI specification
type GetSearchSuggestionsResponse struct {
	Propertynamesuggestions interface{} `json:"PropertyNameSuggestions,omitempty"`
}

// ProfilerConfig represents the ProfilerConfig schema from the OpenAPI specification
type ProfilerConfig struct {
	Profilingparameters interface{} `json:"ProfilingParameters,omitempty"`
	S3outputpath interface{} `json:"S3OutputPath,omitempty"`
	Disableprofiler interface{} `json:"DisableProfiler,omitempty"`
	Profilingintervalinmilliseconds interface{} `json:"ProfilingIntervalInMilliseconds,omitempty"`
}

// ListInferenceRecommendationsJobsRequest represents the ListInferenceRecommendationsJobsRequest schema from the OpenAPI specification
type ListInferenceRecommendationsJobsRequest struct {
	Lastmodifiedtimebefore interface{} `json:"LastModifiedTimeBefore,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Lastmodifiedtimeafter interface{} `json:"LastModifiedTimeAfter,omitempty"`
	Modelnameequals interface{} `json:"ModelNameEquals,omitempty"`
	Modelpackageversionarnequals interface{} `json:"ModelPackageVersionArnEquals,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Statusequals interface{} `json:"StatusEquals,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
}

// DataProcessing represents the DataProcessing schema from the OpenAPI specification
type DataProcessing struct {
	Inputfilter interface{} `json:"InputFilter,omitempty"`
	Joinsource interface{} `json:"JoinSource,omitempty"`
	Outputfilter interface{} `json:"OutputFilter,omitempty"`
}

// ProfilerConfigForUpdate represents the ProfilerConfigForUpdate schema from the OpenAPI specification
type ProfilerConfigForUpdate struct {
	Profilingintervalinmilliseconds interface{} `json:"ProfilingIntervalInMilliseconds,omitempty"`
	Profilingparameters interface{} `json:"ProfilingParameters,omitempty"`
	S3outputpath interface{} `json:"S3OutputPath,omitempty"`
	Disableprofiler interface{} `json:"DisableProfiler,omitempty"`
}

// DescribeTrialComponentRequest represents the DescribeTrialComponentRequest schema from the OpenAPI specification
type DescribeTrialComponentRequest struct {
	Trialcomponentname interface{} `json:"TrialComponentName"`
}

// ListMonitoringSchedulesRequest represents the ListMonitoringSchedulesRequest schema from the OpenAPI specification
type ListMonitoringSchedulesRequest struct {
	Statusequals interface{} `json:"StatusEquals,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Lastmodifiedtimebefore interface{} `json:"LastModifiedTimeBefore,omitempty"`
	Monitoringjobdefinitionname interface{} `json:"MonitoringJobDefinitionName,omitempty"`
	Monitoringtypeequals interface{} `json:"MonitoringTypeEquals,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Endpointname interface{} `json:"EndpointName,omitempty"`
	Lastmodifiedtimeafter interface{} `json:"LastModifiedTimeAfter,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// HyperParameters represents the HyperParameters schema from the OpenAPI specification
type HyperParameters struct {
}

// UpdateProjectOutput represents the UpdateProjectOutput schema from the OpenAPI specification
type UpdateProjectOutput struct {
	Projectarn interface{} `json:"ProjectArn"`
}

// OnlineStoreConfig represents the OnlineStoreConfig schema from the OpenAPI specification
type OnlineStoreConfig struct {
	Ttlduration interface{} `json:"TtlDuration,omitempty"`
	Enableonlinestore interface{} `json:"EnableOnlineStore,omitempty"`
	Securityconfig interface{} `json:"SecurityConfig,omitempty"`
}

// CreateModelPackageOutput represents the CreateModelPackageOutput schema from the OpenAPI specification
type CreateModelPackageOutput struct {
	Modelpackagearn interface{} `json:"ModelPackageArn"`
}

// ModelDashboardModelCard represents the ModelDashboardModelCard schema from the OpenAPI specification
type ModelDashboardModelCard struct {
	Securityconfig interface{} `json:"SecurityConfig,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Modelcardarn interface{} `json:"ModelCardArn,omitempty"`
	Modelid interface{} `json:"ModelId,omitempty"`
	Riskrating interface{} `json:"RiskRating,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Modelcardname interface{} `json:"ModelCardName,omitempty"`
	Modelcardstatus interface{} `json:"ModelCardStatus,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Createdby UserContext `json:"CreatedBy,omitempty"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	Lastmodifiedby UserContext `json:"LastModifiedBy,omitempty"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	Modelcardversion interface{} `json:"ModelCardVersion,omitempty"`
}

// AsyncInferenceConfig represents the AsyncInferenceConfig schema from the OpenAPI specification
type AsyncInferenceConfig struct {
	Outputconfig interface{} `json:"OutputConfig"`
	Clientconfig interface{} `json:"ClientConfig,omitempty"`
}

// DescribeHumanTaskUiRequest represents the DescribeHumanTaskUiRequest schema from the OpenAPI specification
type DescribeHumanTaskUiRequest struct {
	Humantaskuiname interface{} `json:"HumanTaskUiName"`
}

// DescribeArtifactRequest represents the DescribeArtifactRequest schema from the OpenAPI specification
type DescribeArtifactRequest struct {
	Artifactarn interface{} `json:"ArtifactArn"`
}

// DescribeActionRequest represents the DescribeActionRequest schema from the OpenAPI specification
type DescribeActionRequest struct {
	Actionname interface{} `json:"ActionName"`
}

// HyperParameterTuningJobConfig represents the HyperParameterTuningJobConfig schema from the OpenAPI specification
type HyperParameterTuningJobConfig struct {
	Hyperparametertuningjobobjective interface{} `json:"HyperParameterTuningJobObjective,omitempty"`
	Parameterranges interface{} `json:"ParameterRanges,omitempty"`
	Randomseed interface{} `json:"RandomSeed,omitempty"`
	Resourcelimits interface{} `json:"ResourceLimits"`
	Strategy interface{} `json:"Strategy"`
	Strategyconfig interface{} `json:"StrategyConfig,omitempty"`
	Trainingjobearlystoppingtype interface{} `json:"TrainingJobEarlyStoppingType,omitempty"`
	Tuningjobcompletioncriteria interface{} `json:"TuningJobCompletionCriteria,omitempty"`
}

// HyperParameterTrainingJobDefinition represents the HyperParameterTrainingJobDefinition schema from the OpenAPI specification
type HyperParameterTrainingJobDefinition struct {
	Retrystrategy interface{} `json:"RetryStrategy,omitempty"`
	Tuningobjective HyperParameterTuningJobObjective `json:"TuningObjective,omitempty"` // Defines the objective metric for a hyperparameter tuning job. Hyperparameter tuning uses the value of this metric to evaluate the training jobs it launches, and returns the training job that results in either the highest or lowest value for this metric, depending on the value you specify for the <code>Type</code> parameter.
	Enableintercontainertrafficencryption interface{} `json:"EnableInterContainerTrafficEncryption,omitempty"`
	Rolearn interface{} `json:"RoleArn"`
	Enablemanagedspottraining interface{} `json:"EnableManagedSpotTraining,omitempty"`
	Hyperparameterranges ParameterRanges `json:"HyperParameterRanges,omitempty"` // <p>Specifies ranges of integer, continuous, and categorical hyperparameters that a hyperparameter tuning job searches. The hyperparameter tuning job launches training jobs with hyperparameter values within these ranges to find the combination of values that result in the training job with the best performance as measured by the objective metric of the hyperparameter tuning job.</p> <note> <p>The maximum number of items specified for <code>Array Members</code> refers to the maximum number of hyperparameters for each range and also the maximum for the hyperparameter tuning job itself. That is, the sum of the number of hyperparameters for all the ranges can't exceed the maximum number specified.</p> </note>
	Stoppingcondition interface{} `json:"StoppingCondition"`
	Enablenetworkisolation interface{} `json:"EnableNetworkIsolation,omitempty"`
	Vpcconfig interface{} `json:"VpcConfig,omitempty"`
	Algorithmspecification interface{} `json:"AlgorithmSpecification"`
	Hyperparametertuningresourceconfig interface{} `json:"HyperParameterTuningResourceConfig,omitempty"`
	Environment interface{} `json:"Environment,omitempty"`
	Inputdataconfig interface{} `json:"InputDataConfig,omitempty"`
	Definitionname interface{} `json:"DefinitionName,omitempty"`
	Resourceconfig interface{} `json:"ResourceConfig,omitempty"`
	Outputdataconfig interface{} `json:"OutputDataConfig"`
	Statichyperparameters interface{} `json:"StaticHyperParameters,omitempty"`
	Checkpointconfig CheckpointConfig `json:"CheckpointConfig,omitempty"` // Contains information about the output location for managed spot training checkpoint data.
}

// TrialComponent represents the TrialComponent schema from the OpenAPI specification
type TrialComponent struct {
	Endtime interface{} `json:"EndTime,omitempty"`
	Sourcedetail interface{} `json:"SourceDetail,omitempty"`
	Outputartifacts interface{} `json:"OutputArtifacts,omitempty"`
	Trialcomponentname interface{} `json:"TrialComponentName,omitempty"`
	Metadataproperties MetadataProperties `json:"MetadataProperties,omitempty"` // Metadata properties of the tracking entity, trial, or trial component.
	Status TrialComponentStatus `json:"Status,omitempty"` // The status of the trial component.
	Createdby interface{} `json:"CreatedBy,omitempty"`
	Displayname interface{} `json:"DisplayName,omitempty"`
	Lastmodifiedby UserContext `json:"LastModifiedBy,omitempty"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Lineagegrouparn interface{} `json:"LineageGroupArn,omitempty"`
	Metrics interface{} `json:"Metrics,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
	Parameters interface{} `json:"Parameters,omitempty"`
	Parents interface{} `json:"Parents,omitempty"`
	Trialcomponentarn interface{} `json:"TrialComponentArn,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Runname interface{} `json:"RunName,omitempty"`
	Source interface{} `json:"Source,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Inputartifacts interface{} `json:"InputArtifacts,omitempty"`
}

// ListModelQualityJobDefinitionsRequest represents the ListModelQualityJobDefinitionsRequest schema from the OpenAPI specification
type ListModelQualityJobDefinitionsRequest struct {
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Endpointname interface{} `json:"EndpointName,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
}

// AthenaDatasetDefinition represents the AthenaDatasetDefinition schema from the OpenAPI specification
type AthenaDatasetDefinition struct {
	Querystring string `json:"QueryString"` // The SQL query statements, to be executed.
	Workgroup string `json:"WorkGroup,omitempty"` // The name of the workgroup in which the Athena query is being started.
	Catalog string `json:"Catalog"` // The name of the data catalog used in Athena query execution.
	Database string `json:"Database"` // The name of the database used in the Athena query execution.
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Outputcompression string `json:"OutputCompression,omitempty"` // The compression used for Athena query results.
	Outputformat string `json:"OutputFormat"` // The data storage format for Athena query results.
	Outputs3uri interface{} `json:"OutputS3Uri"`
}

// CreateModelCardExportJobResponse represents the CreateModelCardExportJobResponse schema from the OpenAPI specification
type CreateModelCardExportJobResponse struct {
	Modelcardexportjobarn interface{} `json:"ModelCardExportJobArn"`
}

// DescribeSubscribedWorkteamResponse represents the DescribeSubscribedWorkteamResponse schema from the OpenAPI specification
type DescribeSubscribedWorkteamResponse struct {
	Subscribedworkteam interface{} `json:"SubscribedWorkteam"`
}

// ListEdgePackagingJobsRequest represents the ListEdgePackagingJobsRequest schema from the OpenAPI specification
type ListEdgePackagingJobsRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Lastmodifiedtimeafter interface{} `json:"LastModifiedTimeAfter,omitempty"`
	Modelnamecontains interface{} `json:"ModelNameContains,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Lastmodifiedtimebefore interface{} `json:"LastModifiedTimeBefore,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Statusequals interface{} `json:"StatusEquals,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
}

// LabelingJobS3DataSource represents the LabelingJobS3DataSource schema from the OpenAPI specification
type LabelingJobS3DataSource struct {
	Manifests3uri interface{} `json:"ManifestS3Uri"`
}

// CreateUserProfileRequest represents the CreateUserProfileRequest schema from the OpenAPI specification
type CreateUserProfileRequest struct {
	Usersettings interface{} `json:"UserSettings,omitempty"`
	Domainid interface{} `json:"DomainId"`
	Singlesignonuseridentifier interface{} `json:"SingleSignOnUserIdentifier,omitempty"`
	Singlesignonuservalue interface{} `json:"SingleSignOnUserValue,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Userprofilename interface{} `json:"UserProfileName"`
}

// DescribeModelPackageGroupOutput represents the DescribeModelPackageGroupOutput schema from the OpenAPI specification
type DescribeModelPackageGroupOutput struct {
	Createdby UserContext `json:"CreatedBy"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	Creationtime interface{} `json:"CreationTime"`
	Modelpackagegrouparn interface{} `json:"ModelPackageGroupArn"`
	Modelpackagegroupdescription interface{} `json:"ModelPackageGroupDescription,omitempty"`
	Modelpackagegroupname interface{} `json:"ModelPackageGroupName"`
	Modelpackagegroupstatus interface{} `json:"ModelPackageGroupStatus"`
}

// InferenceExecutionConfig represents the InferenceExecutionConfig schema from the OpenAPI specification
type InferenceExecutionConfig struct {
	Mode interface{} `json:"Mode"`
}

// GetModelPackageGroupPolicyOutput represents the GetModelPackageGroupPolicyOutput schema from the OpenAPI specification
type GetModelPackageGroupPolicyOutput struct {
	Resourcepolicy interface{} `json:"ResourcePolicy"`
}

// SelectedStep represents the SelectedStep schema from the OpenAPI specification
type SelectedStep struct {
	Stepname interface{} `json:"StepName"`
}

// CreatePresignedNotebookInstanceUrlInput represents the CreatePresignedNotebookInstanceUrlInput schema from the OpenAPI specification
type CreatePresignedNotebookInstanceUrlInput struct {
	Notebookinstancename interface{} `json:"NotebookInstanceName"`
	Sessionexpirationdurationinseconds interface{} `json:"SessionExpirationDurationInSeconds,omitempty"`
}

// ListInferenceRecommendationsJobStepsRequest represents the ListInferenceRecommendationsJobStepsRequest schema from the OpenAPI specification
type ListInferenceRecommendationsJobStepsRequest struct {
	Steptype interface{} `json:"StepType,omitempty"`
	Jobname interface{} `json:"JobName"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// DeleteSpaceRequest represents the DeleteSpaceRequest schema from the OpenAPI specification
type DeleteSpaceRequest struct {
	Domainid interface{} `json:"DomainId"`
	Spacename interface{} `json:"SpaceName"`
}

// OidcConfigForResponse represents the OidcConfigForResponse schema from the OpenAPI specification
type OidcConfigForResponse struct {
	Logoutendpoint interface{} `json:"LogoutEndpoint,omitempty"`
	Tokenendpoint interface{} `json:"TokenEndpoint,omitempty"`
	Userinfoendpoint interface{} `json:"UserInfoEndpoint,omitempty"`
	Authorizationendpoint interface{} `json:"AuthorizationEndpoint,omitempty"`
	Clientid interface{} `json:"ClientId,omitempty"`
	Issuer interface{} `json:"Issuer,omitempty"`
	Jwksuri interface{} `json:"JwksUri,omitempty"`
}

// QueryFilters represents the QueryFilters schema from the OpenAPI specification
type QueryFilters struct {
	Createdafter interface{} `json:"CreatedAfter,omitempty"`
	Createdbefore interface{} `json:"CreatedBefore,omitempty"`
	Lineagetypes interface{} `json:"LineageTypes,omitempty"`
	Modifiedafter interface{} `json:"ModifiedAfter,omitempty"`
	Modifiedbefore interface{} `json:"ModifiedBefore,omitempty"`
	Properties interface{} `json:"Properties,omitempty"`
	Types interface{} `json:"Types,omitempty"`
}

// DescribePipelineDefinitionForExecutionResponse represents the DescribePipelineDefinitionForExecutionResponse schema from the OpenAPI specification
type DescribePipelineDefinitionForExecutionResponse struct {
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Pipelinedefinition interface{} `json:"PipelineDefinition,omitempty"`
}

// UpdateHubResponse represents the UpdateHubResponse schema from the OpenAPI specification
type UpdateHubResponse struct {
	Hubarn interface{} `json:"HubArn"`
}

// DescribeHubContentResponse represents the DescribeHubContentResponse schema from the OpenAPI specification
type DescribeHubContentResponse struct {
	Hubcontentmarkdown interface{} `json:"HubContentMarkdown,omitempty"`
	Hubcontentdescription interface{} `json:"HubContentDescription,omitempty"`
	Hubname interface{} `json:"HubName"`
	Hubcontentarn interface{} `json:"HubContentArn"`
	Hubcontentversion interface{} `json:"HubContentVersion"`
	Hubcontentdisplayname interface{} `json:"HubContentDisplayName,omitempty"`
	Creationtime interface{} `json:"CreationTime"`
	Hubcontentsearchkeywords interface{} `json:"HubContentSearchKeywords,omitempty"`
	Hubcontentdocument interface{} `json:"HubContentDocument"`
	Documentschemaversion interface{} `json:"DocumentSchemaVersion"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Hubcontentstatus interface{} `json:"HubContentStatus"`
	Hubarn interface{} `json:"HubArn"`
	Hubcontentdependencies interface{} `json:"HubContentDependencies,omitempty"`
	Hubcontentname interface{} `json:"HubContentName"`
	Hubcontenttype interface{} `json:"HubContentType"`
}

// ProcessingS3Input represents the ProcessingS3Input schema from the OpenAPI specification
type ProcessingS3Input struct {
	S3compressiontype interface{} `json:"S3CompressionType,omitempty"`
	S3datadistributiontype interface{} `json:"S3DataDistributionType,omitempty"`
	S3datatype interface{} `json:"S3DataType"`
	S3inputmode interface{} `json:"S3InputMode,omitempty"`
	S3uri interface{} `json:"S3Uri"`
	Localpath interface{} `json:"LocalPath,omitempty"`
}

// TransformOutput represents the TransformOutput schema from the OpenAPI specification
type TransformOutput struct {
	S3outputpath interface{} `json:"S3OutputPath"`
	Accept interface{} `json:"Accept,omitempty"`
	Assemblewith interface{} `json:"AssembleWith,omitempty"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
}

// Trial represents the Trial schema from the OpenAPI specification
type Trial struct {
	Trialname interface{} `json:"TrialName,omitempty"`
	Experimentname interface{} `json:"ExperimentName,omitempty"`
	Source TrialSource `json:"Source,omitempty"` // The source of the trial.
	Trialarn interface{} `json:"TrialArn,omitempty"`
	Createdby interface{} `json:"CreatedBy,omitempty"`
	Displayname interface{} `json:"DisplayName,omitempty"`
	Lastmodifiedby UserContext `json:"LastModifiedBy,omitempty"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Metadataproperties MetadataProperties `json:"MetadataProperties,omitempty"` // Metadata properties of the tracking entity, trial, or trial component.
	Tags interface{} `json:"Tags,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Trialcomponentsummaries interface{} `json:"TrialComponentSummaries,omitempty"`
}

// ModelCardExportOutputConfig represents the ModelCardExportOutputConfig schema from the OpenAPI specification
type ModelCardExportOutputConfig struct {
	S3outputpath interface{} `json:"S3OutputPath"`
}

// ProductionVariantServerlessUpdateConfig represents the ProductionVariantServerlessUpdateConfig schema from the OpenAPI specification
type ProductionVariantServerlessUpdateConfig struct {
	Maxconcurrency interface{} `json:"MaxConcurrency,omitempty"`
	Provisionedconcurrency interface{} `json:"ProvisionedConcurrency,omitempty"`
}

// DeleteImageResponse represents the DeleteImageResponse schema from the OpenAPI specification
type DeleteImageResponse struct {
}

// DescribeDomainResponse represents the DescribeDomainResponse schema from the OpenAPI specification
type DescribeDomainResponse struct {
	Domainid interface{} `json:"DomainId,omitempty"`
	Homeefsfilesystemkmskeyid interface{} `json:"HomeEfsFileSystemKmsKeyId,omitempty"`
	Securitygroupidfordomainboundary interface{} `json:"SecurityGroupIdForDomainBoundary,omitempty"`
	Domainsettings interface{} `json:"DomainSettings,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Appnetworkaccesstype interface{} `json:"AppNetworkAccessType,omitempty"`
	Defaultspacesettings interface{} `json:"DefaultSpaceSettings,omitempty"`
	Subnetids interface{} `json:"SubnetIds,omitempty"`
	Authmode interface{} `json:"AuthMode,omitempty"`
	Domainname interface{} `json:"DomainName,omitempty"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Appsecuritygroupmanagement interface{} `json:"AppSecurityGroupManagement,omitempty"`
	Defaultusersettings interface{} `json:"DefaultUserSettings,omitempty"`
	Vpcid interface{} `json:"VpcId,omitempty"`
	Domainarn interface{} `json:"DomainArn,omitempty"`
	Url interface{} `json:"Url,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Singlesignonmanagedapplicationinstanceid interface{} `json:"SingleSignOnManagedApplicationInstanceId,omitempty"`
	Homeefsfilesystemid interface{} `json:"HomeEfsFileSystemId,omitempty"`
}

// FileSource represents the FileSource schema from the OpenAPI specification
type FileSource struct {
	S3uri interface{} `json:"S3Uri"`
	Contentdigest interface{} `json:"ContentDigest,omitempty"`
	Contenttype interface{} `json:"ContentType,omitempty"`
}

// FinalAutoMLJobObjectiveMetric represents the FinalAutoMLJobObjectiveMetric schema from the OpenAPI specification
type FinalAutoMLJobObjectiveMetric struct {
	Value interface{} `json:"Value"`
	Metricname interface{} `json:"MetricName"`
	Standardmetricname interface{} `json:"StandardMetricName,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// UpdateProjectInput represents the UpdateProjectInput schema from the OpenAPI specification
type UpdateProjectInput struct {
	Servicecatalogprovisioningupdatedetails interface{} `json:"ServiceCatalogProvisioningUpdateDetails,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Projectdescription interface{} `json:"ProjectDescription,omitempty"`
	Projectname interface{} `json:"ProjectName"`
}

// UpdateInferenceExperimentRequest represents the UpdateInferenceExperimentRequest schema from the OpenAPI specification
type UpdateInferenceExperimentRequest struct {
	Schedule interface{} `json:"Schedule,omitempty"`
	Shadowmodeconfig interface{} `json:"ShadowModeConfig,omitempty"`
	Datastorageconfig interface{} `json:"DataStorageConfig,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Modelvariants interface{} `json:"ModelVariants,omitempty"`
	Name interface{} `json:"Name"`
}

// UserProfileDetails represents the UserProfileDetails schema from the OpenAPI specification
type UserProfileDetails struct {
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Domainid interface{} `json:"DomainId,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Userprofilename interface{} `json:"UserProfileName,omitempty"`
}

// ListUserProfilesRequest represents the ListUserProfilesRequest schema from the OpenAPI specification
type ListUserProfilesRequest struct {
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Userprofilenamecontains interface{} `json:"UserProfileNameContains,omitempty"`
	Domainidequals interface{} `json:"DomainIdEquals,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CreateProjectInput represents the CreateProjectInput schema from the OpenAPI specification
type CreateProjectInput struct {
	Projectdescription interface{} `json:"ProjectDescription,omitempty"`
	Projectname interface{} `json:"ProjectName"`
	Servicecatalogprovisioningdetails interface{} `json:"ServiceCatalogProvisioningDetails"`
	Tags interface{} `json:"Tags,omitempty"`
}

// InferenceRecommendationsJobStep represents the InferenceRecommendationsJobStep schema from the OpenAPI specification
type InferenceRecommendationsJobStep struct {
	Status interface{} `json:"Status"`
	Steptype interface{} `json:"StepType"`
	Inferencebenchmark interface{} `json:"InferenceBenchmark,omitempty"`
	Jobname interface{} `json:"JobName"`
}

// DebugRuleConfiguration represents the DebugRuleConfiguration schema from the OpenAPI specification
type DebugRuleConfiguration struct {
	Ruleevaluatorimage interface{} `json:"RuleEvaluatorImage"`
	Ruleparameters interface{} `json:"RuleParameters,omitempty"`
	S3outputpath interface{} `json:"S3OutputPath,omitempty"`
	Volumesizeingb interface{} `json:"VolumeSizeInGB,omitempty"`
	Instancetype interface{} `json:"InstanceType,omitempty"`
	Localpath interface{} `json:"LocalPath,omitempty"`
	Ruleconfigurationname interface{} `json:"RuleConfigurationName"`
}

// HyperParameterAlgorithmSpecification represents the HyperParameterAlgorithmSpecification schema from the OpenAPI specification
type HyperParameterAlgorithmSpecification struct {
	Algorithmname interface{} `json:"AlgorithmName,omitempty"`
	Metricdefinitions interface{} `json:"MetricDefinitions,omitempty"`
	Trainingimage interface{} `json:"TrainingImage,omitempty"`
	Traininginputmode string `json:"TrainingInputMode"` // <p>The training input mode that the algorithm supports. For more information about input modes, see <a href="https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html">Algorithms</a>.</p> <p> <b>Pipe mode</b> </p> <p>If an algorithm supports <code>Pipe</code> mode, Amazon SageMaker streams data directly from Amazon S3 to the container.</p> <p> <b>File mode</b> </p> <p>If an algorithm supports <code>File</code> mode, SageMaker downloads the training data from S3 to the provisioned ML storage volume, and mounts the directory to the Docker volume for the training container.</p> <p>You must provision the ML storage volume with sufficient capacity to accommodate the data downloaded from S3. In addition to the training data, the ML storage volume also stores the output model. The algorithm container uses the ML storage volume to also store intermediate information, if any.</p> <p>For distributed algorithms, training data is distributed uniformly. Your training duration is predictable if the input data objects sizes are approximately the same. SageMaker does not split the files any further for model training. If the object sizes are skewed, training won't be optimal as the data distribution is also skewed when one host in a training cluster is overloaded, thus becoming a bottleneck in training.</p> <p> <b>FastFile mode</b> </p> <p>If an algorithm supports <code>FastFile</code> mode, SageMaker streams data directly from S3 to the container with no code changes, and provides file system access to the data. Users can author their training script to interact with these files as if they were stored on disk.</p> <p> <code>FastFile</code> mode works best when the data is read sequentially. Augmented manifest files aren't supported. The startup time is lower when there are fewer files in the S3 bucket provided.</p>
}

// ListStudioLifecycleConfigsRequest represents the ListStudioLifecycleConfigsRequest schema from the OpenAPI specification
type ListStudioLifecycleConfigsRequest struct {
	Namecontains interface{} `json:"NameContains,omitempty"`
	Apptypeequals interface{} `json:"AppTypeEquals,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Modifiedtimeafter interface{} `json:"ModifiedTimeAfter,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Modifiedtimebefore interface{} `json:"ModifiedTimeBefore,omitempty"`
}

// AutoMLContainerDefinition represents the AutoMLContainerDefinition schema from the OpenAPI specification
type AutoMLContainerDefinition struct {
	Environment interface{} `json:"Environment,omitempty"`
	Image interface{} `json:"Image"`
	Modeldataurl interface{} `json:"ModelDataUrl"`
}

// CreateModelExplainabilityJobDefinitionRequest represents the CreateModelExplainabilityJobDefinitionRequest schema from the OpenAPI specification
type CreateModelExplainabilityJobDefinitionRequest struct {
	Modelexplainabilitybaselineconfig interface{} `json:"ModelExplainabilityBaselineConfig,omitempty"`
	Modelexplainabilityjobinput interface{} `json:"ModelExplainabilityJobInput"`
	Modelexplainabilityjoboutputconfig MonitoringOutputConfig `json:"ModelExplainabilityJobOutputConfig"` // The output configuration for monitoring jobs.
	Modelexplainabilityappspecification interface{} `json:"ModelExplainabilityAppSpecification"`
	Networkconfig interface{} `json:"NetworkConfig,omitempty"`
	Rolearn interface{} `json:"RoleArn"`
	Stoppingcondition MonitoringStoppingCondition `json:"StoppingCondition,omitempty"` // A time limit for how long the monitoring job is allowed to run before stopping.
	Tags interface{} `json:"Tags,omitempty"`
	Jobdefinitionname interface{} `json:"JobDefinitionName"`
	Jobresources MonitoringResources `json:"JobResources"` // Identifies the resources to deploy for a monitoring job.
}

// ServiceCatalogProvisioningUpdateDetails represents the ServiceCatalogProvisioningUpdateDetails schema from the OpenAPI specification
type ServiceCatalogProvisioningUpdateDetails struct {
	Provisioningparameters interface{} `json:"ProvisioningParameters,omitempty"`
	Provisioningartifactid interface{} `json:"ProvisioningArtifactId,omitempty"`
}

// ListLineageGroupsRequest represents the ListLineageGroupsRequest schema from the OpenAPI specification
type ListLineageGroupsRequest struct {
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Createdafter interface{} `json:"CreatedAfter,omitempty"`
	Createdbefore interface{} `json:"CreatedBefore,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
}

// DomainSettings represents the DomainSettings schema from the OpenAPI specification
type DomainSettings struct {
	Executionroleidentityconfig interface{} `json:"ExecutionRoleIdentityConfig,omitempty"`
	Rstudioserverprodomainsettings interface{} `json:"RStudioServerProDomainSettings,omitempty"`
	Securitygroupids interface{} `json:"SecurityGroupIds,omitempty"`
}

// OutputConfig represents the OutputConfig schema from the OpenAPI specification
type OutputConfig struct {
	S3outputlocation interface{} `json:"S3OutputLocation"`
	Targetdevice interface{} `json:"TargetDevice,omitempty"`
	Targetplatform interface{} `json:"TargetPlatform,omitempty"`
	Compileroptions interface{} `json:"CompilerOptions,omitempty"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
}

// EdgeOutputConfig represents the EdgeOutputConfig schema from the OpenAPI specification
type EdgeOutputConfig struct {
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Presetdeploymentconfig interface{} `json:"PresetDeploymentConfig,omitempty"`
	Presetdeploymenttype interface{} `json:"PresetDeploymentType,omitempty"`
	S3outputlocation interface{} `json:"S3OutputLocation"`
}

// ListMonitoringExecutionsRequest represents the ListMonitoringExecutionsRequest schema from the OpenAPI specification
type ListMonitoringExecutionsRequest struct {
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Endpointname interface{} `json:"EndpointName,omitempty"`
	Monitoringschedulename interface{} `json:"MonitoringScheduleName,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Lastmodifiedtimeafter interface{} `json:"LastModifiedTimeAfter,omitempty"`
	Monitoringjobdefinitionname interface{} `json:"MonitoringJobDefinitionName,omitempty"`
	Statusequals interface{} `json:"StatusEquals,omitempty"`
	Scheduledtimeafter interface{} `json:"ScheduledTimeAfter,omitempty"`
	Monitoringtypeequals interface{} `json:"MonitoringTypeEquals,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Lastmodifiedtimebefore interface{} `json:"LastModifiedTimeBefore,omitempty"`
	Scheduledtimebefore interface{} `json:"ScheduledTimeBefore,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
}

// RStudioServerProDomainSettingsForUpdate represents the RStudioServerProDomainSettingsForUpdate schema from the OpenAPI specification
type RStudioServerProDomainSettingsForUpdate struct {
	Rstudioconnecturl interface{} `json:"RStudioConnectUrl,omitempty"`
	Rstudiopackagemanagerurl interface{} `json:"RStudioPackageManagerUrl,omitempty"`
	Defaultresourcespec ResourceSpec `json:"DefaultResourceSpec,omitempty"` // Specifies the ARN's of a SageMaker image and SageMaker image version, and the instance type that the version runs on.
	Domainexecutionrolearn interface{} `json:"DomainExecutionRoleArn"`
}

// CreateAppResponse represents the CreateAppResponse schema from the OpenAPI specification
type CreateAppResponse struct {
	Apparn interface{} `json:"AppArn,omitempty"`
}

// NotebookInstanceSummary represents the NotebookInstanceSummary schema from the OpenAPI specification
type NotebookInstanceSummary struct {
	Instancetype interface{} `json:"InstanceType,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Notebookinstancearn interface{} `json:"NotebookInstanceArn"`
	Notebookinstancename interface{} `json:"NotebookInstanceName"`
	Notebookinstancestatus interface{} `json:"NotebookInstanceStatus,omitempty"`
	Additionalcoderepositories interface{} `json:"AdditionalCodeRepositories,omitempty"`
	Defaultcoderepository interface{} `json:"DefaultCodeRepository,omitempty"`
	Notebookinstancelifecycleconfigname interface{} `json:"NotebookInstanceLifecycleConfigName,omitempty"`
	Url interface{} `json:"Url,omitempty"`
}

// CreateCompilationJobResponse represents the CreateCompilationJobResponse schema from the OpenAPI specification
type CreateCompilationJobResponse struct {
	Compilationjobarn interface{} `json:"CompilationJobArn"`
}

// HyperParameterSpecification represents the HyperParameterSpecification schema from the OpenAPI specification
type HyperParameterSpecification struct {
	Defaultvalue interface{} `json:"DefaultValue,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Isrequired interface{} `json:"IsRequired,omitempty"`
	Istunable interface{} `json:"IsTunable,omitempty"`
	Name interface{} `json:"Name"`
	RangeField interface{} `json:"Range,omitempty"`
	TypeField interface{} `json:"Type"`
}

// CreateCodeRepositoryOutput represents the CreateCodeRepositoryOutput schema from the OpenAPI specification
type CreateCodeRepositoryOutput struct {
	Coderepositoryarn interface{} `json:"CodeRepositoryArn"`
}

// ProcessingJobSummary represents the ProcessingJobSummary schema from the OpenAPI specification
type ProcessingJobSummary struct {
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Processingendtime interface{} `json:"ProcessingEndTime,omitempty"`
	Processingjobarn interface{} `json:"ProcessingJobArn"`
	Processingjobname interface{} `json:"ProcessingJobName"`
	Processingjobstatus interface{} `json:"ProcessingJobStatus"`
	Creationtime interface{} `json:"CreationTime"`
	Exitmessage interface{} `json:"ExitMessage,omitempty"`
}

// Channel represents the Channel schema from the OpenAPI specification
type Channel struct {
	Channelname interface{} `json:"ChannelName"`
	Compressiontype interface{} `json:"CompressionType,omitempty"`
	Contenttype interface{} `json:"ContentType,omitempty"`
	Datasource interface{} `json:"DataSource"`
	Inputmode interface{} `json:"InputMode,omitempty"`
	Recordwrappertype interface{} `json:"RecordWrapperType,omitempty"`
	Shuffleconfig interface{} `json:"ShuffleConfig,omitempty"`
}

// TrainingJobStatusCounters represents the TrainingJobStatusCounters schema from the OpenAPI specification
type TrainingJobStatusCounters struct {
	Completed interface{} `json:"Completed,omitempty"`
	Inprogress interface{} `json:"InProgress,omitempty"`
	Nonretryableerror interface{} `json:"NonRetryableError,omitempty"`
	Retryableerror interface{} `json:"RetryableError,omitempty"`
	Stopped interface{} `json:"Stopped,omitempty"`
}

// OutputDataConfig represents the OutputDataConfig schema from the OpenAPI specification
type OutputDataConfig struct {
	Compressiontype interface{} `json:"CompressionType,omitempty"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	S3outputpath interface{} `json:"S3OutputPath"`
}

// MonitoringCsvDatasetFormat represents the MonitoringCsvDatasetFormat schema from the OpenAPI specification
type MonitoringCsvDatasetFormat struct {
	Header interface{} `json:"Header,omitempty"`
}

// ListModelMetadataResponse represents the ListModelMetadataResponse schema from the OpenAPI specification
type ListModelMetadataResponse struct {
	Modelmetadatasummaries interface{} `json:"ModelMetadataSummaries"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// SourceAlgorithm represents the SourceAlgorithm schema from the OpenAPI specification
type SourceAlgorithm struct {
	Algorithmname interface{} `json:"AlgorithmName"`
	Modeldataurl interface{} `json:"ModelDataUrl,omitempty"`
}

// CreateContextResponse represents the CreateContextResponse schema from the OpenAPI specification
type CreateContextResponse struct {
	Contextarn interface{} `json:"ContextArn,omitempty"`
}

// LabelingJobDataSource represents the LabelingJobDataSource schema from the OpenAPI specification
type LabelingJobDataSource struct {
	S3datasource interface{} `json:"S3DataSource,omitempty"`
	Snsdatasource interface{} `json:"SnsDataSource,omitempty"`
}

// GetScalingConfigurationRecommendationRequest represents the GetScalingConfigurationRecommendationRequest schema from the OpenAPI specification
type GetScalingConfigurationRecommendationRequest struct {
	Targetcpuutilizationpercore interface{} `json:"TargetCpuUtilizationPerCore,omitempty"`
	Endpointname interface{} `json:"EndpointName,omitempty"`
	Inferencerecommendationsjobname interface{} `json:"InferenceRecommendationsJobName"`
	Recommendationid interface{} `json:"RecommendationId,omitempty"`
	Scalingpolicyobjective interface{} `json:"ScalingPolicyObjective,omitempty"`
}

// DescribePipelineExecutionRequest represents the DescribePipelineExecutionRequest schema from the OpenAPI specification
type DescribePipelineExecutionRequest struct {
	Pipelineexecutionarn interface{} `json:"PipelineExecutionArn"`
}

// NotebookInstanceLifecycleHook represents the NotebookInstanceLifecycleHook schema from the OpenAPI specification
type NotebookInstanceLifecycleHook struct {
	Content interface{} `json:"Content,omitempty"`
}

// DescribeLineageGroupRequest represents the DescribeLineageGroupRequest schema from the OpenAPI specification
type DescribeLineageGroupRequest struct {
	Lineagegroupname interface{} `json:"LineageGroupName"`
}

// ListPipelineExecutionStepsResponse represents the ListPipelineExecutionStepsResponse schema from the OpenAPI specification
type ListPipelineExecutionStepsResponse struct {
	Pipelineexecutionsteps interface{} `json:"PipelineExecutionSteps,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListWorkforcesRequest represents the ListWorkforcesRequest schema from the OpenAPI specification
type ListWorkforcesRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
}

// ModelDataQuality represents the ModelDataQuality schema from the OpenAPI specification
type ModelDataQuality struct {
	Constraints interface{} `json:"Constraints,omitempty"`
	Statistics interface{} `json:"Statistics,omitempty"`
}

// LastUpdateStatus represents the LastUpdateStatus schema from the OpenAPI specification
type LastUpdateStatus struct {
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Status interface{} `json:"Status"`
}

// UpdateWorkteamRequest represents the UpdateWorkteamRequest schema from the OpenAPI specification
type UpdateWorkteamRequest struct {
	Description interface{} `json:"Description,omitempty"`
	Memberdefinitions interface{} `json:"MemberDefinitions,omitempty"`
	Notificationconfiguration interface{} `json:"NotificationConfiguration,omitempty"`
	Workteamname interface{} `json:"WorkteamName"`
}

// ChannelSpecification represents the ChannelSpecification schema from the OpenAPI specification
type ChannelSpecification struct {
	Isrequired interface{} `json:"IsRequired,omitempty"`
	Name interface{} `json:"Name"`
	Supportedcompressiontypes interface{} `json:"SupportedCompressionTypes,omitempty"`
	Supportedcontenttypes interface{} `json:"SupportedContentTypes"`
	Supportedinputmodes interface{} `json:"SupportedInputModes"`
	Description interface{} `json:"Description,omitempty"`
}

// CreatePipelineResponse represents the CreatePipelineResponse schema from the OpenAPI specification
type CreatePipelineResponse struct {
	Pipelinearn interface{} `json:"PipelineArn,omitempty"`
}

// ModelMetadataSummary represents the ModelMetadataSummary schema from the OpenAPI specification
type ModelMetadataSummary struct {
	Model interface{} `json:"Model"`
	Task interface{} `json:"Task"`
	Domain interface{} `json:"Domain"`
	Framework interface{} `json:"Framework"`
	Frameworkversion interface{} `json:"FrameworkVersion"`
}

// ListAliasesResponse represents the ListAliasesResponse schema from the OpenAPI specification
type ListAliasesResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sagemakerimageversionaliases interface{} `json:"SageMakerImageVersionAliases,omitempty"`
}

// ModelCardSummary represents the ModelCardSummary schema from the OpenAPI specification
type ModelCardSummary struct {
	Creationtime interface{} `json:"CreationTime"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Modelcardarn interface{} `json:"ModelCardArn"`
	Modelcardname interface{} `json:"ModelCardName"`
	Modelcardstatus interface{} `json:"ModelCardStatus"`
}

// CreateHumanTaskUiRequest represents the CreateHumanTaskUiRequest schema from the OpenAPI specification
type CreateHumanTaskUiRequest struct {
	Humantaskuiname interface{} `json:"HumanTaskUiName"`
	Tags interface{} `json:"Tags,omitempty"`
	Uitemplate UiTemplate `json:"UiTemplate"` // The Liquid template for the worker user interface.
}

// HyperParameterTuningJobSearchEntity represents the HyperParameterTuningJobSearchEntity schema from the OpenAPI specification
type HyperParameterTuningJobSearchEntity struct {
	Consumedresources interface{} `json:"ConsumedResources,omitempty"`
	Warmstartconfig HyperParameterTuningJobWarmStartConfig `json:"WarmStartConfig,omitempty"` // <p>Specifies the configuration for a hyperparameter tuning job that uses one or more previous hyperparameter tuning jobs as a starting point. The results of previous tuning jobs are used to inform which combinations of hyperparameters to search over in the new tuning job.</p> <p>All training jobs launched by the new hyperparameter tuning job are evaluated by using the objective metric, and the training job that performs the best is compared to the best training jobs from the parent tuning jobs. From these, the training job that performs the best as measured by the objective metric is returned as the overall best training job.</p> <note> <p>All training jobs launched by parent hyperparameter tuning jobs and the new hyperparameter tuning jobs count against the limit of training jobs for the tuning job.</p> </note>
	Trainingjobstatuscounters TrainingJobStatusCounters `json:"TrainingJobStatusCounters,omitempty"` // The numbers of training jobs launched by a hyperparameter tuning job, categorized by status.
	Objectivestatuscounters ObjectiveStatusCounters `json:"ObjectiveStatusCounters,omitempty"` // Specifies the number of training jobs that this hyperparameter tuning job launched, categorized by the status of their objective metric. The objective metric status shows whether the final objective metric for the training job has been evaluated by the tuning job and used in the hyperparameter tuning process.
	Hyperparametertuningjobarn interface{} `json:"HyperParameterTuningJobArn,omitempty"`
	Trainingjobdefinition HyperParameterTrainingJobDefinition `json:"TrainingJobDefinition,omitempty"` // Defines the training jobs launched by a hyperparameter tuning job.
	Hyperparametertuningendtime interface{} `json:"HyperParameterTuningEndTime,omitempty"`
	Overallbesttrainingjob HyperParameterTrainingJobSummary `json:"OverallBestTrainingJob,omitempty"` // The container for the summary information about a training job.
	Tuningjobcompletiondetails interface{} `json:"TuningJobCompletionDetails,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Hyperparametertuningjobname interface{} `json:"HyperParameterTuningJobName,omitempty"`
	Hyperparametertuningjobstatus interface{} `json:"HyperParameterTuningJobStatus,omitempty"`
	Trainingjobdefinitions interface{} `json:"TrainingJobDefinitions,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Besttrainingjob HyperParameterTrainingJobSummary `json:"BestTrainingJob,omitempty"` // The container for the summary information about a training job.
	Tags interface{} `json:"Tags,omitempty"`
	Hyperparametertuningjobconfig HyperParameterTuningJobConfig `json:"HyperParameterTuningJobConfig,omitempty"` // Configures a hyperparameter tuning job.
}

// TrialComponentMetricSummary represents the TrialComponentMetricSummary schema from the OpenAPI specification
type TrialComponentMetricSummary struct {
	Last interface{} `json:"Last,omitempty"`
	Avg interface{} `json:"Avg,omitempty"`
	Count interface{} `json:"Count,omitempty"`
	Min interface{} `json:"Min,omitempty"`
	Max interface{} `json:"Max,omitempty"`
	Metricname interface{} `json:"MetricName,omitempty"`
	Sourcearn interface{} `json:"SourceArn,omitempty"`
	Stddev interface{} `json:"StdDev,omitempty"`
	Timestamp interface{} `json:"TimeStamp,omitempty"`
}

// DescribeEdgeDeploymentPlanRequest represents the DescribeEdgeDeploymentPlanRequest schema from the OpenAPI specification
type DescribeEdgeDeploymentPlanRequest struct {
	Edgedeploymentplanname interface{} `json:"EdgeDeploymentPlanName"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// PipelineExecutionStepMetadata represents the PipelineExecutionStepMetadata schema from the OpenAPI specification
type PipelineExecutionStepMetadata struct {
	Automljob interface{} `json:"AutoMLJob,omitempty"`
	Callback interface{} `json:"Callback,omitempty"`
	Emr interface{} `json:"EMR,omitempty"`
	Model interface{} `json:"Model,omitempty"`
	Processingjob interface{} `json:"ProcessingJob,omitempty"`
	Qualitycheck interface{} `json:"QualityCheck,omitempty"`
	Transformjob interface{} `json:"TransformJob,omitempty"`
	Fail interface{} `json:"Fail,omitempty"`
	Condition interface{} `json:"Condition,omitempty"`
	Clarifycheck interface{} `json:"ClarifyCheck,omitempty"`
	Lambda interface{} `json:"Lambda,omitempty"`
	Registermodel interface{} `json:"RegisterModel,omitempty"`
	Trainingjob interface{} `json:"TrainingJob,omitempty"`
	Tuningjob interface{} `json:"TuningJob,omitempty"`
}

// DeploymentConfig represents the DeploymentConfig schema from the OpenAPI specification
type DeploymentConfig struct {
	Rollingupdatepolicy interface{} `json:"RollingUpdatePolicy,omitempty"`
	Autorollbackconfiguration interface{} `json:"AutoRollbackConfiguration,omitempty"`
	Bluegreenupdatepolicy interface{} `json:"BlueGreenUpdatePolicy,omitempty"`
}

// ScheduleConfig represents the ScheduleConfig schema from the OpenAPI specification
type ScheduleConfig struct {
	Scheduleexpression interface{} `json:"ScheduleExpression"`
}

// ParameterRanges represents the ParameterRanges schema from the OpenAPI specification
type ParameterRanges struct {
	Autoparameters interface{} `json:"AutoParameters,omitempty"`
	Categoricalparameterranges interface{} `json:"CategoricalParameterRanges,omitempty"`
	Continuousparameterranges interface{} `json:"ContinuousParameterRanges,omitempty"`
	Integerparameterranges interface{} `json:"IntegerParameterRanges,omitempty"`
}

// UpdateCodeRepositoryOutput represents the UpdateCodeRepositoryOutput schema from the OpenAPI specification
type UpdateCodeRepositoryOutput struct {
	Coderepositoryarn interface{} `json:"CodeRepositoryArn"`
}

// DomainDetails represents the DomainDetails schema from the OpenAPI specification
type DomainDetails struct {
	Status interface{} `json:"Status,omitempty"`
	Url interface{} `json:"Url,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Domainarn interface{} `json:"DomainArn,omitempty"`
	Domainid interface{} `json:"DomainId,omitempty"`
	Domainname interface{} `json:"DomainName,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
}

// HumanTaskUiSummary represents the HumanTaskUiSummary schema from the OpenAPI specification
type HumanTaskUiSummary struct {
	Creationtime interface{} `json:"CreationTime"`
	Humantaskuiarn interface{} `json:"HumanTaskUiArn"`
	Humantaskuiname interface{} `json:"HumanTaskUiName"`
}

// AlgorithmSpecification represents the AlgorithmSpecification schema from the OpenAPI specification
type AlgorithmSpecification struct {
	Traininginputmode string `json:"TrainingInputMode"` // <p>The training input mode that the algorithm supports. For more information about input modes, see <a href="https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html">Algorithms</a>.</p> <p> <b>Pipe mode</b> </p> <p>If an algorithm supports <code>Pipe</code> mode, Amazon SageMaker streams data directly from Amazon S3 to the container.</p> <p> <b>File mode</b> </p> <p>If an algorithm supports <code>File</code> mode, SageMaker downloads the training data from S3 to the provisioned ML storage volume, and mounts the directory to the Docker volume for the training container.</p> <p>You must provision the ML storage volume with sufficient capacity to accommodate the data downloaded from S3. In addition to the training data, the ML storage volume also stores the output model. The algorithm container uses the ML storage volume to also store intermediate information, if any.</p> <p>For distributed algorithms, training data is distributed uniformly. Your training duration is predictable if the input data objects sizes are approximately the same. SageMaker does not split the files any further for model training. If the object sizes are skewed, training won't be optimal as the data distribution is also skewed when one host in a training cluster is overloaded, thus becoming a bottleneck in training.</p> <p> <b>FastFile mode</b> </p> <p>If an algorithm supports <code>FastFile</code> mode, SageMaker streams data directly from S3 to the container with no code changes, and provides file system access to the data. Users can author their training script to interact with these files as if they were stored on disk.</p> <p> <code>FastFile</code> mode works best when the data is read sequentially. Augmented manifest files aren't supported. The startup time is lower when there are fewer files in the S3 bucket provided.</p>
	Algorithmname interface{} `json:"AlgorithmName,omitempty"`
	Containerarguments interface{} `json:"ContainerArguments,omitempty"`
	Containerentrypoint interface{} `json:"ContainerEntrypoint,omitempty"`
	Enablesagemakermetricstimeseries interface{} `json:"EnableSageMakerMetricsTimeSeries,omitempty"`
	Metricdefinitions interface{} `json:"MetricDefinitions,omitempty"`
	Trainingimage interface{} `json:"TrainingImage,omitempty"`
	Trainingimageconfig interface{} `json:"TrainingImageConfig,omitempty"`
}

// ListProcessingJobsResponse represents the ListProcessingJobsResponse schema from the OpenAPI specification
type ListProcessingJobsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Processingjobsummaries interface{} `json:"ProcessingJobSummaries"`
}

// StopEdgeDeploymentStageRequest represents the StopEdgeDeploymentStageRequest schema from the OpenAPI specification
type StopEdgeDeploymentStageRequest struct {
	Edgedeploymentplanname interface{} `json:"EdgeDeploymentPlanName"`
	Stagename interface{} `json:"StageName"`
}

// ListDevicesResponse represents the ListDevicesResponse schema from the OpenAPI specification
type ListDevicesResponse struct {
	Devicesummaries interface{} `json:"DeviceSummaries"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DescribeModelOutput represents the DescribeModelOutput schema from the OpenAPI specification
type DescribeModelOutput struct {
	Containers interface{} `json:"Containers,omitempty"`
	Enablenetworkisolation interface{} `json:"EnableNetworkIsolation,omitempty"`
	Inferenceexecutionconfig interface{} `json:"InferenceExecutionConfig,omitempty"`
	Primarycontainer interface{} `json:"PrimaryContainer,omitempty"`
	Creationtime interface{} `json:"CreationTime"`
	Executionrolearn interface{} `json:"ExecutionRoleArn"`
	Modelname interface{} `json:"ModelName"`
	Vpcconfig interface{} `json:"VpcConfig,omitempty"`
	Deploymentrecommendation interface{} `json:"DeploymentRecommendation,omitempty"`
	Modelarn interface{} `json:"ModelArn"`
}

// AddTagsOutput represents the AddTagsOutput schema from the OpenAPI specification
type AddTagsOutput struct {
	Tags interface{} `json:"Tags,omitempty"`
}

// RStudioServerProAppSettings represents the RStudioServerProAppSettings schema from the OpenAPI specification
type RStudioServerProAppSettings struct {
	Accessstatus interface{} `json:"AccessStatus,omitempty"`
	Usergroup interface{} `json:"UserGroup,omitempty"`
}

// CreateProcessingJobResponse represents the CreateProcessingJobResponse schema from the OpenAPI specification
type CreateProcessingJobResponse struct {
	Processingjobarn interface{} `json:"ProcessingJobArn"`
}

// AutoMLJobStepMetadata represents the AutoMLJobStepMetadata schema from the OpenAPI specification
type AutoMLJobStepMetadata struct {
	Arn interface{} `json:"Arn,omitempty"`
}

// CreateStudioLifecycleConfigRequest represents the CreateStudioLifecycleConfigRequest schema from the OpenAPI specification
type CreateStudioLifecycleConfigRequest struct {
	Studiolifecycleconfigapptype interface{} `json:"StudioLifecycleConfigAppType"`
	Studiolifecycleconfigcontent interface{} `json:"StudioLifecycleConfigContent"`
	Studiolifecycleconfigname interface{} `json:"StudioLifecycleConfigName"`
	Tags interface{} `json:"Tags,omitempty"`
}

// InferenceExperimentSchedule represents the InferenceExperimentSchedule schema from the OpenAPI specification
type InferenceExperimentSchedule struct {
	Endtime interface{} `json:"EndTime,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
}

// ListExperimentsRequest represents the ListExperimentsRequest schema from the OpenAPI specification
type ListExperimentsRequest struct {
	Createdafter interface{} `json:"CreatedAfter,omitempty"`
	Createdbefore interface{} `json:"CreatedBefore,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
}

// CategoricalParameterRange represents the CategoricalParameterRange schema from the OpenAPI specification
type CategoricalParameterRange struct {
	Values interface{} `json:"Values"`
	Name interface{} `json:"Name"`
}

// TransformDataSource represents the TransformDataSource schema from the OpenAPI specification
type TransformDataSource struct {
	S3datasource interface{} `json:"S3DataSource"`
}

// LabelingJobInputConfig represents the LabelingJobInputConfig schema from the OpenAPI specification
type LabelingJobInputConfig struct {
	Dataattributes interface{} `json:"DataAttributes,omitempty"`
	Datasource interface{} `json:"DataSource"`
}

// ListAlgorithmsOutput represents the ListAlgorithmsOutput schema from the OpenAPI specification
type ListAlgorithmsOutput struct {
	Algorithmsummarylist interface{} `json:"AlgorithmSummaryList"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// MonitoringDatasetFormat represents the MonitoringDatasetFormat schema from the OpenAPI specification
type MonitoringDatasetFormat struct {
	Csv interface{} `json:"Csv,omitempty"`
	Json interface{} `json:"Json,omitempty"`
	Parquet interface{} `json:"Parquet,omitempty"`
}

// DescribeEdgeDeploymentPlanResponse represents the DescribeEdgeDeploymentPlanResponse schema from the OpenAPI specification
type DescribeEdgeDeploymentPlanResponse struct {
	Edgedeploymentfailed interface{} `json:"EdgeDeploymentFailed,omitempty"`
	Edgedeploymentplanarn interface{} `json:"EdgeDeploymentPlanArn"`
	Edgedeploymentsuccess interface{} `json:"EdgeDeploymentSuccess,omitempty"`
	Modelconfigs interface{} `json:"ModelConfigs"`
	Devicefleetname interface{} `json:"DeviceFleetName"`
	Edgedeploymentpending interface{} `json:"EdgeDeploymentPending,omitempty"`
	Edgedeploymentplanname interface{} `json:"EdgeDeploymentPlanName"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Stages interface{} `json:"Stages"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
}

// DataSource represents the DataSource schema from the OpenAPI specification
type DataSource struct {
	Filesystemdatasource interface{} `json:"FileSystemDataSource,omitempty"`
	S3datasource interface{} `json:"S3DataSource,omitempty"`
}

// DataQualityBaselineConfig represents the DataQualityBaselineConfig schema from the OpenAPI specification
type DataQualityBaselineConfig struct {
	Constraintsresource MonitoringConstraintsResource `json:"ConstraintsResource,omitempty"` // The constraints resource for a monitoring job.
	Statisticsresource MonitoringStatisticsResource `json:"StatisticsResource,omitempty"` // The statistics resource for a monitoring job.
	Baseliningjobname interface{} `json:"BaseliningJobName,omitempty"`
}

// ContainerDefinition represents the ContainerDefinition schema from the OpenAPI specification
type ContainerDefinition struct {
	Modeldataurl interface{} `json:"ModelDataUrl,omitempty"`
	Inferencespecificationname interface{} `json:"InferenceSpecificationName,omitempty"`
	Modeldatasource interface{} `json:"ModelDataSource,omitempty"`
	Modelpackagename interface{} `json:"ModelPackageName,omitempty"`
	Multimodelconfig interface{} `json:"MultiModelConfig,omitempty"`
	Mode interface{} `json:"Mode,omitempty"`
	Image interface{} `json:"Image,omitempty"`
	Imageconfig interface{} `json:"ImageConfig,omitempty"`
	Containerhostname interface{} `json:"ContainerHostname,omitempty"`
	Environment interface{} `json:"Environment,omitempty"`
}

// RecommendationJobOutputConfig represents the RecommendationJobOutputConfig schema from the OpenAPI specification
type RecommendationJobOutputConfig struct {
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Compiledoutputconfig interface{} `json:"CompiledOutputConfig,omitempty"`
}

// DeleteContextResponse represents the DeleteContextResponse schema from the OpenAPI specification
type DeleteContextResponse struct {
	Contextarn interface{} `json:"ContextArn,omitempty"`
}

// TransformInput represents the TransformInput schema from the OpenAPI specification
type TransformInput struct {
	Datasource interface{} `json:"DataSource"`
	Splittype interface{} `json:"SplitType,omitempty"`
	Compressiontype interface{} `json:"CompressionType,omitempty"`
	Contenttype interface{} `json:"ContentType,omitempty"`
}

// UiTemplate represents the UiTemplate schema from the OpenAPI specification
type UiTemplate struct {
	Content interface{} `json:"Content"`
}

// DescribeSpaceResponse represents the DescribeSpaceResponse schema from the OpenAPI specification
type DescribeSpaceResponse struct {
	Spacearn interface{} `json:"SpaceArn,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Spacename interface{} `json:"SpaceName,omitempty"`
	Spacesettings interface{} `json:"SpaceSettings,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Homeefsfilesystemuid interface{} `json:"HomeEfsFileSystemUid,omitempty"`
	Domainid interface{} `json:"DomainId,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
}

// ListPipelineExecutionsRequest represents the ListPipelineExecutionsRequest schema from the OpenAPI specification
type ListPipelineExecutionsRequest struct {
	Createdbefore interface{} `json:"CreatedBefore,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Pipelinename interface{} `json:"PipelineName"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Createdafter interface{} `json:"CreatedAfter,omitempty"`
}

// DisassociateTrialComponentResponse represents the DisassociateTrialComponentResponse schema from the OpenAPI specification
type DisassociateTrialComponentResponse struct {
	Trialarn interface{} `json:"TrialArn,omitempty"`
	Trialcomponentarn interface{} `json:"TrialComponentArn,omitempty"`
}

// UpdateEndpointInput represents the UpdateEndpointInput schema from the OpenAPI specification
type UpdateEndpointInput struct {
	Excluderetainedvariantproperties interface{} `json:"ExcludeRetainedVariantProperties,omitempty"`
	Retainallvariantproperties interface{} `json:"RetainAllVariantProperties,omitempty"`
	Retaindeploymentconfig interface{} `json:"RetainDeploymentConfig,omitempty"`
	Deploymentconfig interface{} `json:"DeploymentConfig,omitempty"`
	Endpointconfigname interface{} `json:"EndpointConfigName"`
	Endpointname interface{} `json:"EndpointName"`
}

// ProfilingParameters represents the ProfilingParameters schema from the OpenAPI specification
type ProfilingParameters struct {
}

// CompilationJobSummary represents the CompilationJobSummary schema from the OpenAPI specification
type CompilationJobSummary struct {
	Compilationjobstatus interface{} `json:"CompilationJobStatus"`
	Compilationstarttime interface{} `json:"CompilationStartTime,omitempty"`
	Compilationtargetdevice interface{} `json:"CompilationTargetDevice,omitempty"`
	Compilationtargetplatformaccelerator interface{} `json:"CompilationTargetPlatformAccelerator,omitempty"`
	Compilationtargetplatformarch interface{} `json:"CompilationTargetPlatformArch,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Compilationjobarn interface{} `json:"CompilationJobArn"`
	Compilationjobname interface{} `json:"CompilationJobName"`
	Compilationtargetplatformos interface{} `json:"CompilationTargetPlatformOs,omitempty"`
	Compilationendtime interface{} `json:"CompilationEndTime,omitempty"`
	Creationtime interface{} `json:"CreationTime"`
}

// CategoricalParameterRangeSpecification represents the CategoricalParameterRangeSpecification schema from the OpenAPI specification
type CategoricalParameterRangeSpecification struct {
	Values interface{} `json:"Values"`
}

// DescribeImageVersionResponse represents the DescribeImageVersionResponse schema from the OpenAPI specification
type DescribeImageVersionResponse struct {
	Imageversionarn interface{} `json:"ImageVersionArn,omitempty"`
	Baseimage interface{} `json:"BaseImage,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Processor interface{} `json:"Processor,omitempty"`
	Releasenotes interface{} `json:"ReleaseNotes,omitempty"`
	Imageversionstatus interface{} `json:"ImageVersionStatus,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Version interface{} `json:"Version,omitempty"`
	Containerimage interface{} `json:"ContainerImage,omitempty"`
	Imagearn interface{} `json:"ImageArn,omitempty"`
	Programminglang interface{} `json:"ProgrammingLang,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Jobtype interface{} `json:"JobType,omitempty"`
	Mlframework interface{} `json:"MLFramework,omitempty"`
	Vendorguidance interface{} `json:"VendorGuidance,omitempty"`
	Horovod interface{} `json:"Horovod,omitempty"`
}

// ListLabelingJobsRequest represents the ListLabelingJobsRequest schema from the OpenAPI specification
type ListLabelingJobsRequest struct {
	Statusequals interface{} `json:"StatusEquals,omitempty"`
	Lastmodifiedtimebefore interface{} `json:"LastModifiedTimeBefore,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Lastmodifiedtimeafter interface{} `json:"LastModifiedTimeAfter,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
}

// Alarm represents the Alarm schema from the OpenAPI specification
type Alarm struct {
	Alarmname interface{} `json:"AlarmName,omitempty"`
}

// HyperParameterTuningInstanceConfig represents the HyperParameterTuningInstanceConfig schema from the OpenAPI specification
type HyperParameterTuningInstanceConfig struct {
	Instancecount interface{} `json:"InstanceCount"`
	Instancetype interface{} `json:"InstanceType"`
	Volumesizeingb interface{} `json:"VolumeSizeInGB"`
}

// AutoMLCandidateGenerationConfig represents the AutoMLCandidateGenerationConfig schema from the OpenAPI specification
type AutoMLCandidateGenerationConfig struct {
	Algorithmsconfig interface{} `json:"AlgorithmsConfig,omitempty"`
	Featurespecifications3uri interface{} `json:"FeatureSpecificationS3Uri,omitempty"`
}

// StartEdgeDeploymentStageRequest represents the StartEdgeDeploymentStageRequest schema from the OpenAPI specification
type StartEdgeDeploymentStageRequest struct {
	Edgedeploymentplanname interface{} `json:"EdgeDeploymentPlanName"`
	Stagename interface{} `json:"StageName"`
}

// DescribeModelInput represents the DescribeModelInput schema from the OpenAPI specification
type DescribeModelInput struct {
	Modelname interface{} `json:"ModelName"`
}

// DescribeDataQualityJobDefinitionResponse represents the DescribeDataQualityJobDefinitionResponse schema from the OpenAPI specification
type DescribeDataQualityJobDefinitionResponse struct {
	Dataqualityjoboutputconfig MonitoringOutputConfig `json:"DataQualityJobOutputConfig"` // The output configuration for monitoring jobs.
	Stoppingcondition MonitoringStoppingCondition `json:"StoppingCondition,omitempty"` // A time limit for how long the monitoring job is allowed to run before stopping.
	Dataqualityappspecification interface{} `json:"DataQualityAppSpecification"`
	Dataqualitybaselineconfig interface{} `json:"DataQualityBaselineConfig,omitempty"`
	Networkconfig interface{} `json:"NetworkConfig,omitempty"`
	Jobdefinitionarn interface{} `json:"JobDefinitionArn"`
	Jobdefinitionname interface{} `json:"JobDefinitionName"`
	Jobresources MonitoringResources `json:"JobResources"` // Identifies the resources to deploy for a monitoring job.
	Rolearn interface{} `json:"RoleArn"`
	Creationtime interface{} `json:"CreationTime"`
	Dataqualityjobinput interface{} `json:"DataQualityJobInput"`
}

// UpdateInferenceExperimentResponse represents the UpdateInferenceExperimentResponse schema from the OpenAPI specification
type UpdateInferenceExperimentResponse struct {
	Inferenceexperimentarn interface{} `json:"InferenceExperimentArn"`
}

// ResourceConfigForUpdate represents the ResourceConfigForUpdate schema from the OpenAPI specification
type ResourceConfigForUpdate struct {
	Keepaliveperiodinseconds interface{} `json:"KeepAlivePeriodInSeconds"`
}

// OidcMemberDefinition represents the OidcMemberDefinition schema from the OpenAPI specification
type OidcMemberDefinition struct {
	Groups interface{} `json:"Groups"`
}

// DescribeHyperParameterTuningJobResponse represents the DescribeHyperParameterTuningJobResponse schema from the OpenAPI specification
type DescribeHyperParameterTuningJobResponse struct {
	Creationtime interface{} `json:"CreationTime"`
	Tuningjobcompletiondetails interface{} `json:"TuningJobCompletionDetails,omitempty"`
	Warmstartconfig interface{} `json:"WarmStartConfig,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Objectivestatuscounters interface{} `json:"ObjectiveStatusCounters"`
	Besttrainingjob interface{} `json:"BestTrainingJob,omitempty"`
	Hyperparametertuningjobstatus interface{} `json:"HyperParameterTuningJobStatus"`
	Hyperparametertuningjobconfig interface{} `json:"HyperParameterTuningJobConfig"`
	Trainingjobdefinitions interface{} `json:"TrainingJobDefinitions,omitempty"`
	Consumedresources HyperParameterTuningJobConsumedResources `json:"ConsumedResources,omitempty"` // The total resources consumed by your hyperparameter tuning job.
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Hyperparametertuningjobarn interface{} `json:"HyperParameterTuningJobArn"`
	Autotune interface{} `json:"Autotune,omitempty"`
	Trainingjobdefinition interface{} `json:"TrainingJobDefinition,omitempty"`
	Hyperparametertuningjobname interface{} `json:"HyperParameterTuningJobName"`
	Overallbesttrainingjob interface{} `json:"OverallBestTrainingJob,omitempty"`
	Hyperparametertuningendtime interface{} `json:"HyperParameterTuningEndTime,omitempty"`
	Trainingjobstatuscounters interface{} `json:"TrainingJobStatusCounters"`
}

// DescribeHubContentRequest represents the DescribeHubContentRequest schema from the OpenAPI specification
type DescribeHubContentRequest struct {
	Hubcontenttype interface{} `json:"HubContentType"`
	Hubcontentversion interface{} `json:"HubContentVersion,omitempty"`
	Hubname interface{} `json:"HubName"`
	Hubcontentname interface{} `json:"HubContentName"`
}

// SharingSettings represents the SharingSettings schema from the OpenAPI specification
type SharingSettings struct {
	Notebookoutputoption interface{} `json:"NotebookOutputOption,omitempty"`
	S3kmskeyid interface{} `json:"S3KmsKeyId,omitempty"`
	S3outputpath interface{} `json:"S3OutputPath,omitempty"`
}

// FileSystemDataSource represents the FileSystemDataSource schema from the OpenAPI specification
type FileSystemDataSource struct {
	Filesystemtype interface{} `json:"FileSystemType"`
	Directorypath interface{} `json:"DirectoryPath"`
	Filesystemaccessmode interface{} `json:"FileSystemAccessMode"`
	Filesystemid interface{} `json:"FileSystemId"`
}

// Device represents the Device schema from the OpenAPI specification
type Device struct {
	Description interface{} `json:"Description,omitempty"`
	Devicename interface{} `json:"DeviceName"`
	Iotthingname interface{} `json:"IotThingName,omitempty"`
}

// RecommendationJobCompiledOutputConfig represents the RecommendationJobCompiledOutputConfig schema from the OpenAPI specification
type RecommendationJobCompiledOutputConfig struct {
	S3outputuri interface{} `json:"S3OutputUri,omitempty"`
}

// PutModelPackageGroupPolicyOutput represents the PutModelPackageGroupPolicyOutput schema from the OpenAPI specification
type PutModelPackageGroupPolicyOutput struct {
	Modelpackagegrouparn interface{} `json:"ModelPackageGroupArn"`
}

// ListPipelinesResponse represents the ListPipelinesResponse schema from the OpenAPI specification
type ListPipelinesResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Pipelinesummaries interface{} `json:"PipelineSummaries,omitempty"`
}

// CreateHyperParameterTuningJobRequest represents the CreateHyperParameterTuningJobRequest schema from the OpenAPI specification
type CreateHyperParameterTuningJobRequest struct {
	Autotune interface{} `json:"Autotune,omitempty"`
	Hyperparametertuningjobconfig interface{} `json:"HyperParameterTuningJobConfig"`
	Hyperparametertuningjobname interface{} `json:"HyperParameterTuningJobName"`
	Tags interface{} `json:"Tags,omitempty"`
	Trainingjobdefinition interface{} `json:"TrainingJobDefinition,omitempty"`
	Trainingjobdefinitions interface{} `json:"TrainingJobDefinitions,omitempty"`
	Warmstartconfig interface{} `json:"WarmStartConfig,omitempty"`
}

// ModelClientConfig represents the ModelClientConfig schema from the OpenAPI specification
type ModelClientConfig struct {
	Invocationsmaxretries interface{} `json:"InvocationsMaxRetries,omitempty"`
	Invocationstimeoutinseconds interface{} `json:"InvocationsTimeoutInSeconds,omitempty"`
}

// CreateModelCardRequest represents the CreateModelCardRequest schema from the OpenAPI specification
type CreateModelCardRequest struct {
	Modelcardname interface{} `json:"ModelCardName"`
	Modelcardstatus interface{} `json:"ModelCardStatus"`
	Securityconfig interface{} `json:"SecurityConfig,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Content interface{} `json:"Content"`
}

// ListContextsResponse represents the ListContextsResponse schema from the OpenAPI specification
type ListContextsResponse struct {
	Contextsummaries interface{} `json:"ContextSummaries,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListStudioLifecycleConfigsResponse represents the ListStudioLifecycleConfigsResponse schema from the OpenAPI specification
type ListStudioLifecycleConfigsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Studiolifecycleconfigs interface{} `json:"StudioLifecycleConfigs,omitempty"`
}

// ListImageVersionsRequest represents the ListImageVersionsRequest schema from the OpenAPI specification
type ListImageVersionsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Lastmodifiedtimebefore interface{} `json:"LastModifiedTimeBefore,omitempty"`
	Imagename interface{} `json:"ImageName"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Lastmodifiedtimeafter interface{} `json:"LastModifiedTimeAfter,omitempty"`
}

// SendPipelineExecutionStepFailureResponse represents the SendPipelineExecutionStepFailureResponse schema from the OpenAPI specification
type SendPipelineExecutionStepFailureResponse struct {
	Pipelineexecutionarn interface{} `json:"PipelineExecutionArn,omitempty"`
}

// MonitoringBaselineConfig represents the MonitoringBaselineConfig schema from the OpenAPI specification
type MonitoringBaselineConfig struct {
	Baseliningjobname interface{} `json:"BaseliningJobName,omitempty"`
	Constraintsresource interface{} `json:"ConstraintsResource,omitempty"`
	Statisticsresource interface{} `json:"StatisticsResource,omitempty"`
}

// DeleteDataQualityJobDefinitionRequest represents the DeleteDataQualityJobDefinitionRequest schema from the OpenAPI specification
type DeleteDataQualityJobDefinitionRequest struct {
	Jobdefinitionname interface{} `json:"JobDefinitionName"`
}

// ProductionVariantSummary represents the ProductionVariantSummary schema from the OpenAPI specification
type ProductionVariantSummary struct {
	Currentinstancecount interface{} `json:"CurrentInstanceCount,omitempty"`
	Desiredinstancecount interface{} `json:"DesiredInstanceCount,omitempty"`
	Variantname interface{} `json:"VariantName"`
	Currentweight interface{} `json:"CurrentWeight,omitempty"`
	Deployedimages interface{} `json:"DeployedImages,omitempty"`
	Desiredserverlessconfig interface{} `json:"DesiredServerlessConfig,omitempty"`
	Desiredweight interface{} `json:"DesiredWeight,omitempty"`
	Variantstatus interface{} `json:"VariantStatus,omitempty"`
	Currentserverlessconfig interface{} `json:"CurrentServerlessConfig,omitempty"`
}

// DeleteModelCardRequest represents the DeleteModelCardRequest schema from the OpenAPI specification
type DeleteModelCardRequest struct {
	Modelcardname interface{} `json:"ModelCardName"`
}

// CreateDomainResponse represents the CreateDomainResponse schema from the OpenAPI specification
type CreateDomainResponse struct {
	Domainarn interface{} `json:"DomainArn,omitempty"`
	Url interface{} `json:"Url,omitempty"`
}

// ListModelsInput represents the ListModelsInput schema from the OpenAPI specification
type ListModelsInput struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
}

// AutoMLJobConfig represents the AutoMLJobConfig schema from the OpenAPI specification
type AutoMLJobConfig struct {
	Candidategenerationconfig interface{} `json:"CandidateGenerationConfig,omitempty"`
	Completioncriteria interface{} `json:"CompletionCriteria,omitempty"`
	Datasplitconfig interface{} `json:"DataSplitConfig,omitempty"`
	Mode interface{} `json:"Mode,omitempty"`
	Securityconfig interface{} `json:"SecurityConfig,omitempty"`
}

// AutoMLJobArtifacts represents the AutoMLJobArtifacts schema from the OpenAPI specification
type AutoMLJobArtifacts struct {
	Candidatedefinitionnotebooklocation interface{} `json:"CandidateDefinitionNotebookLocation,omitempty"`
	Dataexplorationnotebooklocation interface{} `json:"DataExplorationNotebookLocation,omitempty"`
}

// DescribeProcessingJobRequest represents the DescribeProcessingJobRequest schema from the OpenAPI specification
type DescribeProcessingJobRequest struct {
	Processingjobname interface{} `json:"ProcessingJobName"`
}

// AsyncInferenceOutputConfig represents the AsyncInferenceOutputConfig schema from the OpenAPI specification
type AsyncInferenceOutputConfig struct {
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Notificationconfig interface{} `json:"NotificationConfig,omitempty"`
	S3failurepath interface{} `json:"S3FailurePath,omitempty"`
	S3outputpath interface{} `json:"S3OutputPath,omitempty"`
}

// MetricDatum represents the MetricDatum schema from the OpenAPI specification
type MetricDatum struct {
	Value interface{} `json:"Value,omitempty"`
	Metricname interface{} `json:"MetricName,omitempty"`
	Set interface{} `json:"Set,omitempty"`
	Standardmetricname interface{} `json:"StandardMetricName,omitempty"`
}

// EndpointInputConfiguration represents the EndpointInputConfiguration schema from the OpenAPI specification
type EndpointInputConfiguration struct {
	Serverlessconfig ProductionVariantServerlessConfig `json:"ServerlessConfig,omitempty"` // Specifies the serverless configuration for an endpoint variant.
	Environmentparameterranges interface{} `json:"EnvironmentParameterRanges,omitempty"`
	Inferencespecificationname interface{} `json:"InferenceSpecificationName,omitempty"`
	Instancetype interface{} `json:"InstanceType,omitempty"`
}

// ListImagesRequest represents the ListImagesRequest schema from the OpenAPI specification
type ListImagesRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Lastmodifiedtimebefore interface{} `json:"LastModifiedTimeBefore,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Lastmodifiedtimeafter interface{} `json:"LastModifiedTimeAfter,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
}

// ModelDashboardModel represents the ModelDashboardModel schema from the OpenAPI specification
type ModelDashboardModel struct {
	Monitoringschedules interface{} `json:"MonitoringSchedules,omitempty"`
	Endpoints interface{} `json:"Endpoints,omitempty"`
	Lastbatchtransformjob TransformJob `json:"LastBatchTransformJob,omitempty"` // A batch transform job. For information about SageMaker batch transform, see <a href="https://docs.aws.amazon.com/sagemaker/latest/dg/batch-transform.html">Use Batch Transform</a>.
	Model interface{} `json:"Model,omitempty"`
	Modelcard interface{} `json:"ModelCard,omitempty"`
}

// OfflineStoreConfig represents the OfflineStoreConfig schema from the OpenAPI specification
type OfflineStoreConfig struct {
	S3storageconfig interface{} `json:"S3StorageConfig"`
	Tableformat interface{} `json:"TableFormat,omitempty"`
	Datacatalogconfig interface{} `json:"DataCatalogConfig,omitempty"`
	Disablegluetablecreation interface{} `json:"DisableGlueTableCreation,omitempty"`
}

// ProfilerRuleEvaluationStatus represents the ProfilerRuleEvaluationStatus schema from the OpenAPI specification
type ProfilerRuleEvaluationStatus struct {
	Statusdetails interface{} `json:"StatusDetails,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Ruleconfigurationname interface{} `json:"RuleConfigurationName,omitempty"`
	Ruleevaluationjobarn interface{} `json:"RuleEvaluationJobArn,omitempty"`
	Ruleevaluationstatus interface{} `json:"RuleEvaluationStatus,omitempty"`
}

// RenderUiTemplateRequest represents the RenderUiTemplateRequest schema from the OpenAPI specification
type RenderUiTemplateRequest struct {
	Task interface{} `json:"Task"`
	Uitemplate interface{} `json:"UiTemplate,omitempty"`
	Humantaskuiarn interface{} `json:"HumanTaskUiArn,omitempty"`
	Rolearn interface{} `json:"RoleArn"`
}

// DeleteModelPackageInput represents the DeleteModelPackageInput schema from the OpenAPI specification
type DeleteModelPackageInput struct {
	Modelpackagename interface{} `json:"ModelPackageName"`
}

// ServiceCatalogProvisionedProductDetails represents the ServiceCatalogProvisionedProductDetails schema from the OpenAPI specification
type ServiceCatalogProvisionedProductDetails struct {
	Provisionedproductid interface{} `json:"ProvisionedProductId,omitempty"`
	Provisionedproductstatusmessage interface{} `json:"ProvisionedProductStatusMessage,omitempty"`
}

// DescribeModelQualityJobDefinitionRequest represents the DescribeModelQualityJobDefinitionRequest schema from the OpenAPI specification
type DescribeModelQualityJobDefinitionRequest struct {
	Jobdefinitionname interface{} `json:"JobDefinitionName"`
}

// DebugRuleEvaluationStatus represents the DebugRuleEvaluationStatus schema from the OpenAPI specification
type DebugRuleEvaluationStatus struct {
	Statusdetails interface{} `json:"StatusDetails,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Ruleconfigurationname interface{} `json:"RuleConfigurationName,omitempty"`
	Ruleevaluationjobarn interface{} `json:"RuleEvaluationJobArn,omitempty"`
	Ruleevaluationstatus interface{} `json:"RuleEvaluationStatus,omitempty"`
}

// DescribeAutoMLJobRequest represents the DescribeAutoMLJobRequest schema from the OpenAPI specification
type DescribeAutoMLJobRequest struct {
	Automljobname interface{} `json:"AutoMLJobName"`
}

// CollectionConfiguration represents the CollectionConfiguration schema from the OpenAPI specification
type CollectionConfiguration struct {
	Collectionname interface{} `json:"CollectionName,omitempty"`
	Collectionparameters interface{} `json:"CollectionParameters,omitempty"`
}

// QueryProperties represents the QueryProperties schema from the OpenAPI specification
type QueryProperties struct {
}

// EndpointPerformance represents the EndpointPerformance schema from the OpenAPI specification
type EndpointPerformance struct {
	Endpointinfo EndpointInfo `json:"EndpointInfo"` // Details about a customer endpoint that was compared in an Inference Recommender job.
	Metrics interface{} `json:"Metrics"`
}

// UpdateTrialComponentRequest represents the UpdateTrialComponentRequest schema from the OpenAPI specification
type UpdateTrialComponentRequest struct {
	Outputartifactstoremove interface{} `json:"OutputArtifactsToRemove,omitempty"`
	Trialcomponentname interface{} `json:"TrialComponentName"`
	Displayname interface{} `json:"DisplayName,omitempty"`
	Endtime interface{} `json:"EndTime,omitempty"`
	Outputartifacts interface{} `json:"OutputArtifacts,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
	Inputartifactstoremove interface{} `json:"InputArtifactsToRemove,omitempty"`
	Parameters interface{} `json:"Parameters,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Parameterstoremove interface{} `json:"ParametersToRemove,omitempty"`
	Inputartifacts interface{} `json:"InputArtifacts,omitempty"`
}

// DriftCheckBias represents the DriftCheckBias schema from the OpenAPI specification
type DriftCheckBias struct {
	Pretrainingconstraints interface{} `json:"PreTrainingConstraints,omitempty"`
	Configfile interface{} `json:"ConfigFile,omitempty"`
	Posttrainingconstraints interface{} `json:"PostTrainingConstraints,omitempty"`
}

// UpdateNotebookInstanceLifecycleConfigInput represents the UpdateNotebookInstanceLifecycleConfigInput schema from the OpenAPI specification
type UpdateNotebookInstanceLifecycleConfigInput struct {
	Onstart interface{} `json:"OnStart,omitempty"`
	Notebookinstancelifecycleconfigname interface{} `json:"NotebookInstanceLifecycleConfigName"`
	Oncreate interface{} `json:"OnCreate,omitempty"`
}

// DisableSagemakerServicecatalogPortfolioInput represents the DisableSagemakerServicecatalogPortfolioInput schema from the OpenAPI specification
type DisableSagemakerServicecatalogPortfolioInput struct {
}

// ImportHubContentResponse represents the ImportHubContentResponse schema from the OpenAPI specification
type ImportHubContentResponse struct {
	Hubcontentarn interface{} `json:"HubContentArn"`
	Hubarn interface{} `json:"HubArn"`
}

// AlgorithmStatusItem represents the AlgorithmStatusItem schema from the OpenAPI specification
type AlgorithmStatusItem struct {
	Status interface{} `json:"Status"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Name interface{} `json:"Name"`
}

// DeleteEdgeDeploymentPlanRequest represents the DeleteEdgeDeploymentPlanRequest schema from the OpenAPI specification
type DeleteEdgeDeploymentPlanRequest struct {
	Edgedeploymentplanname interface{} `json:"EdgeDeploymentPlanName"`
}

// CreateSpaceRequest represents the CreateSpaceRequest schema from the OpenAPI specification
type CreateSpaceRequest struct {
	Spacesettings interface{} `json:"SpaceSettings,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Domainid interface{} `json:"DomainId"`
	Spacename interface{} `json:"SpaceName"`
}

// DescribeFlowDefinitionResponse represents the DescribeFlowDefinitionResponse schema from the OpenAPI specification
type DescribeFlowDefinitionResponse struct {
	Humanlooprequestsource interface{} `json:"HumanLoopRequestSource,omitempty"`
	Outputconfig interface{} `json:"OutputConfig"`
	Rolearn interface{} `json:"RoleArn"`
	Flowdefinitionarn interface{} `json:"FlowDefinitionArn"`
	Flowdefinitionname interface{} `json:"FlowDefinitionName"`
	Humanloopactivationconfig interface{} `json:"HumanLoopActivationConfig,omitempty"`
	Creationtime interface{} `json:"CreationTime"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Flowdefinitionstatus interface{} `json:"FlowDefinitionStatus"`
	Humanloopconfig interface{} `json:"HumanLoopConfig"`
}

// ModelDashboardMonitoringSchedule represents the ModelDashboardMonitoringSchedule schema from the OpenAPI specification
type ModelDashboardMonitoringSchedule struct {
	Lastmonitoringexecutionsummary MonitoringExecutionSummary `json:"LastMonitoringExecutionSummary,omitempty"` // Summary of information about the last monitoring job to run.
	Monitoringtype interface{} `json:"MonitoringType,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Monitoringschedulestatus interface{} `json:"MonitoringScheduleStatus,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Monitoringalertsummaries interface{} `json:"MonitoringAlertSummaries,omitempty"`
	Monitoringschedulearn interface{} `json:"MonitoringScheduleArn,omitempty"`
	Monitoringscheduleconfig MonitoringScheduleConfig `json:"MonitoringScheduleConfig,omitempty"` // Configures the monitoring schedule and defines the monitoring job.
	Monitoringschedulename interface{} `json:"MonitoringScheduleName,omitempty"`
	Endpointname interface{} `json:"EndpointName,omitempty"`
}

// CaptureContentTypeHeader represents the CaptureContentTypeHeader schema from the OpenAPI specification
type CaptureContentTypeHeader struct {
	Csvcontenttypes interface{} `json:"CsvContentTypes,omitempty"`
	Jsoncontenttypes interface{} `json:"JsonContentTypes,omitempty"`
}

// TrainingImageConfig represents the TrainingImageConfig schema from the OpenAPI specification
type TrainingImageConfig struct {
	Trainingrepositoryauthconfig interface{} `json:"TrainingRepositoryAuthConfig,omitempty"`
	Trainingrepositoryaccessmode interface{} `json:"TrainingRepositoryAccessMode"`
}

// ListInferenceExperimentsRequest represents the ListInferenceExperimentsRequest schema from the OpenAPI specification
type ListInferenceExperimentsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Statusequals interface{} `json:"StatusEquals,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Lastmodifiedtimebefore interface{} `json:"LastModifiedTimeBefore,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Lastmodifiedtimeafter interface{} `json:"LastModifiedTimeAfter,omitempty"`
}

// CreateModelInput represents the CreateModelInput schema from the OpenAPI specification
type CreateModelInput struct {
	Vpcconfig interface{} `json:"VpcConfig,omitempty"`
	Containers interface{} `json:"Containers,omitempty"`
	Enablenetworkisolation interface{} `json:"EnableNetworkIsolation,omitempty"`
	Executionrolearn interface{} `json:"ExecutionRoleArn"`
	Inferenceexecutionconfig interface{} `json:"InferenceExecutionConfig,omitempty"`
	Modelname interface{} `json:"ModelName"`
	Primarycontainer interface{} `json:"PrimaryContainer,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// ModelPackageContainerDefinition represents the ModelPackageContainerDefinition schema from the OpenAPI specification
type ModelPackageContainerDefinition struct {
	Environment interface{} `json:"Environment,omitempty"`
	Image interface{} `json:"Image"`
	Imagedigest interface{} `json:"ImageDigest,omitempty"`
	Frameworkversion interface{} `json:"FrameworkVersion,omitempty"`
	Modeldataurl interface{} `json:"ModelDataUrl,omitempty"`
	Nearestmodelname interface{} `json:"NearestModelName,omitempty"`
	Framework interface{} `json:"Framework,omitempty"`
	Containerhostname interface{} `json:"ContainerHostname,omitempty"`
	Modelinput interface{} `json:"ModelInput,omitempty"`
	Productid interface{} `json:"ProductId,omitempty"`
}

// CreateModelPackageGroupOutput represents the CreateModelPackageGroupOutput schema from the OpenAPI specification
type CreateModelPackageGroupOutput struct {
	Modelpackagegrouparn interface{} `json:"ModelPackageGroupArn"`
}

// S3DataSource represents the S3DataSource schema from the OpenAPI specification
type S3DataSource struct {
	S3datadistributiontype interface{} `json:"S3DataDistributionType,omitempty"`
	S3datatype interface{} `json:"S3DataType"`
	S3uri interface{} `json:"S3Uri"`
	Attributenames interface{} `json:"AttributeNames,omitempty"`
	Instancegroupnames interface{} `json:"InstanceGroupNames,omitempty"`
}

// ShuffleConfig represents the ShuffleConfig schema from the OpenAPI specification
type ShuffleConfig struct {
	Seed interface{} `json:"Seed"`
}

// ListSpacesResponse represents the ListSpacesResponse schema from the OpenAPI specification
type ListSpacesResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Spaces interface{} `json:"Spaces,omitempty"`
}

// ListModelMetadataRequest represents the ListModelMetadataRequest schema from the OpenAPI specification
type ListModelMetadataRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Searchexpression interface{} `json:"SearchExpression,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// TrialComponentSource represents the TrialComponentSource schema from the OpenAPI specification
type TrialComponentSource struct {
	Sourcearn interface{} `json:"SourceArn"`
	Sourcetype interface{} `json:"SourceType,omitempty"`
}

// DescribeModelExplainabilityJobDefinitionResponse represents the DescribeModelExplainabilityJobDefinitionResponse schema from the OpenAPI specification
type DescribeModelExplainabilityJobDefinitionResponse struct {
	Rolearn interface{} `json:"RoleArn"`
	Jobdefinitionname interface{} `json:"JobDefinitionName"`
	Jobdefinitionarn interface{} `json:"JobDefinitionArn"`
	Jobresources MonitoringResources `json:"JobResources"` // Identifies the resources to deploy for a monitoring job.
	Modelexplainabilityappspecification interface{} `json:"ModelExplainabilityAppSpecification"`
	Modelexplainabilityjobinput interface{} `json:"ModelExplainabilityJobInput"`
	Networkconfig interface{} `json:"NetworkConfig,omitempty"`
	Stoppingcondition MonitoringStoppingCondition `json:"StoppingCondition,omitempty"` // A time limit for how long the monitoring job is allowed to run before stopping.
	Creationtime interface{} `json:"CreationTime"`
	Modelexplainabilityjoboutputconfig MonitoringOutputConfig `json:"ModelExplainabilityJobOutputConfig"` // The output configuration for monitoring jobs.
	Modelexplainabilitybaselineconfig interface{} `json:"ModelExplainabilityBaselineConfig,omitempty"`
}

// DescribeMonitoringScheduleRequest represents the DescribeMonitoringScheduleRequest schema from the OpenAPI specification
type DescribeMonitoringScheduleRequest struct {
	Monitoringschedulename interface{} `json:"MonitoringScheduleName"`
}

// GetModelPackageGroupPolicyInput represents the GetModelPackageGroupPolicyInput schema from the OpenAPI specification
type GetModelPackageGroupPolicyInput struct {
	Modelpackagegroupname interface{} `json:"ModelPackageGroupName"`
}

// AutoMLChannel represents the AutoMLChannel schema from the OpenAPI specification
type AutoMLChannel struct {
	Datasource interface{} `json:"DataSource"`
	Sampleweightattributename interface{} `json:"SampleWeightAttributeName,omitempty"`
	Targetattributename interface{} `json:"TargetAttributeName"`
	Channeltype interface{} `json:"ChannelType,omitempty"`
	Compressiontype interface{} `json:"CompressionType,omitempty"`
	Contenttype interface{} `json:"ContentType,omitempty"`
}

// DesiredWeightAndCapacity represents the DesiredWeightAndCapacity schema from the OpenAPI specification
type DesiredWeightAndCapacity struct {
	Variantname interface{} `json:"VariantName"`
	Desiredinstancecount interface{} `json:"DesiredInstanceCount,omitempty"`
	Desiredweight interface{} `json:"DesiredWeight,omitempty"`
	Serverlessupdateconfig interface{} `json:"ServerlessUpdateConfig,omitempty"`
}

// DescribePipelineRequest represents the DescribePipelineRequest schema from the OpenAPI specification
type DescribePipelineRequest struct {
	Pipelinename interface{} `json:"PipelineName"`
}

// DeleteTrialComponentResponse represents the DeleteTrialComponentResponse schema from the OpenAPI specification
type DeleteTrialComponentResponse struct {
	Trialcomponentarn interface{} `json:"TrialComponentArn,omitempty"`
}

// ListStageDevicesRequest represents the ListStageDevicesRequest schema from the OpenAPI specification
type ListStageDevicesRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Stagename interface{} `json:"StageName"`
	Edgedeploymentplanname interface{} `json:"EdgeDeploymentPlanName"`
	Excludedevicesdeployedinotherstage interface{} `json:"ExcludeDevicesDeployedInOtherStage,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// ModelBiasJobInput represents the ModelBiasJobInput schema from the OpenAPI specification
type ModelBiasJobInput struct {
	Groundtruths3input interface{} `json:"GroundTruthS3Input"`
	Batchtransforminput interface{} `json:"BatchTransformInput,omitempty"`
	Endpointinput EndpointInput `json:"EndpointInput,omitempty"` // Input object for the endpoint
}

// CollectionParameters represents the CollectionParameters schema from the OpenAPI specification
type CollectionParameters struct {
}

// DeleteHubRequest represents the DeleteHubRequest schema from the OpenAPI specification
type DeleteHubRequest struct {
	Hubname interface{} `json:"HubName"`
}

// ListMonitoringAlertHistoryRequest represents the ListMonitoringAlertHistoryRequest schema from the OpenAPI specification
type ListMonitoringAlertHistoryRequest struct {
	Monitoringschedulename interface{} `json:"MonitoringScheduleName,omitempty"`
	Statusequals interface{} `json:"StatusEquals,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Monitoringalertname interface{} `json:"MonitoringAlertName,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
}

// IntegerParameterRangeSpecification represents the IntegerParameterRangeSpecification schema from the OpenAPI specification
type IntegerParameterRangeSpecification struct {
	Maxvalue interface{} `json:"MaxValue"`
	Minvalue interface{} `json:"MinValue"`
}

// DescribeExperimentResponse represents the DescribeExperimentResponse schema from the OpenAPI specification
type DescribeExperimentResponse struct {
	Createdby interface{} `json:"CreatedBy,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Lastmodifiedby interface{} `json:"LastModifiedBy,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Source interface{} `json:"Source,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Experimentarn interface{} `json:"ExperimentArn,omitempty"`
	Experimentname interface{} `json:"ExperimentName,omitempty"`
	Displayname interface{} `json:"DisplayName,omitempty"`
}

// UpdateImageVersionRequest represents the UpdateImageVersionRequest schema from the OpenAPI specification
type UpdateImageVersionRequest struct {
	Imagename interface{} `json:"ImageName"`
	Mlframework interface{} `json:"MLFramework,omitempty"`
	Programminglang interface{} `json:"ProgrammingLang,omitempty"`
	Alias interface{} `json:"Alias,omitempty"`
	Aliasestoadd interface{} `json:"AliasesToAdd,omitempty"`
	Horovod interface{} `json:"Horovod,omitempty"`
	Jobtype interface{} `json:"JobType,omitempty"`
	Processor interface{} `json:"Processor,omitempty"`
	Releasenotes interface{} `json:"ReleaseNotes,omitempty"`
	Version interface{} `json:"Version,omitempty"`
	Vendorguidance interface{} `json:"VendorGuidance,omitempty"`
	Aliasestodelete interface{} `json:"AliasesToDelete,omitempty"`
}

// DescribeDeviceFleetResponse represents the DescribeDeviceFleetResponse schema from the OpenAPI specification
type DescribeDeviceFleetResponse struct {
	Devicefleetname interface{} `json:"DeviceFleetName"`
	Iotrolealias interface{} `json:"IotRoleAlias,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime"`
	Outputconfig interface{} `json:"OutputConfig"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Creationtime interface{} `json:"CreationTime"`
	Description interface{} `json:"Description,omitempty"`
	Devicefleetarn interface{} `json:"DeviceFleetArn"`
}

// ShadowModelVariantConfig represents the ShadowModelVariantConfig schema from the OpenAPI specification
type ShadowModelVariantConfig struct {
	Samplingpercentage interface{} `json:"SamplingPercentage"`
	Shadowmodelvariantname interface{} `json:"ShadowModelVariantName"`
}

// ListAssociationsRequest represents the ListAssociationsRequest schema from the OpenAPI specification
type ListAssociationsRequest struct {
	Createdafter interface{} `json:"CreatedAfter,omitempty"`
	Createdbefore interface{} `json:"CreatedBefore,omitempty"`
	Destinationarn interface{} `json:"DestinationArn,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Sourcetype interface{} `json:"SourceType,omitempty"`
	Associationtype interface{} `json:"AssociationType,omitempty"`
	Destinationtype interface{} `json:"DestinationType,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sourcearn interface{} `json:"SourceArn,omitempty"`
}

// DescribeNotebookInstanceLifecycleConfigOutput represents the DescribeNotebookInstanceLifecycleConfigOutput schema from the OpenAPI specification
type DescribeNotebookInstanceLifecycleConfigOutput struct {
	Notebookinstancelifecycleconfigarn interface{} `json:"NotebookInstanceLifecycleConfigArn,omitempty"`
	Notebookinstancelifecycleconfigname interface{} `json:"NotebookInstanceLifecycleConfigName,omitempty"`
	Oncreate interface{} `json:"OnCreate,omitempty"`
	Onstart interface{} `json:"OnStart,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
}

// CreateEdgeDeploymentStageRequest represents the CreateEdgeDeploymentStageRequest schema from the OpenAPI specification
type CreateEdgeDeploymentStageRequest struct {
	Edgedeploymentplanname interface{} `json:"EdgeDeploymentPlanName"`
	Stages interface{} `json:"Stages"`
}

// TensorBoardAppSettings represents the TensorBoardAppSettings schema from the OpenAPI specification
type TensorBoardAppSettings struct {
	Defaultresourcespec interface{} `json:"DefaultResourceSpec,omitempty"`
}

// DeleteArtifactRequest represents the DeleteArtifactRequest schema from the OpenAPI specification
type DeleteArtifactRequest struct {
	Source interface{} `json:"Source,omitempty"`
	Artifactarn interface{} `json:"ArtifactArn,omitempty"`
}

// UpdateDomainResponse represents the UpdateDomainResponse schema from the OpenAPI specification
type UpdateDomainResponse struct {
	Domainarn interface{} `json:"DomainArn,omitempty"`
}

// MonitoringAlertSummary represents the MonitoringAlertSummary schema from the OpenAPI specification
type MonitoringAlertSummary struct {
	Creationtime interface{} `json:"CreationTime"`
	Datapointstoalert interface{} `json:"DatapointsToAlert"`
	Evaluationperiod interface{} `json:"EvaluationPeriod"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime"`
	Monitoringalertname interface{} `json:"MonitoringAlertName"`
	Actions interface{} `json:"Actions"`
	Alertstatus interface{} `json:"AlertStatus"`
}

// DescribeEndpointConfigInput represents the DescribeEndpointConfigInput schema from the OpenAPI specification
type DescribeEndpointConfigInput struct {
	Endpointconfigname interface{} `json:"EndpointConfigName"`
}

// DisassociateTrialComponentRequest represents the DisassociateTrialComponentRequest schema from the OpenAPI specification
type DisassociateTrialComponentRequest struct {
	Trialcomponentname interface{} `json:"TrialComponentName"`
	Trialname interface{} `json:"TrialName"`
}

// DescribeCodeRepositoryOutput represents the DescribeCodeRepositoryOutput schema from the OpenAPI specification
type DescribeCodeRepositoryOutput struct {
	Lastmodifiedtime interface{} `json:"LastModifiedTime"`
	Coderepositoryarn interface{} `json:"CodeRepositoryArn"`
	Coderepositoryname interface{} `json:"CodeRepositoryName"`
	Creationtime interface{} `json:"CreationTime"`
	Gitconfig interface{} `json:"GitConfig,omitempty"`
}

// DescribeLabelingJobRequest represents the DescribeLabelingJobRequest schema from the OpenAPI specification
type DescribeLabelingJobRequest struct {
	Labelingjobname interface{} `json:"LabelingJobName"`
}

// ListWorkteamsRequest represents the ListWorkteamsRequest schema from the OpenAPI specification
type ListWorkteamsRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
}

// ListModelQualityJobDefinitionsResponse represents the ListModelQualityJobDefinitionsResponse schema from the OpenAPI specification
type ListModelQualityJobDefinitionsResponse struct {
	Jobdefinitionsummaries interface{} `json:"JobDefinitionSummaries"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// TransformEnvironmentMap represents the TransformEnvironmentMap schema from the OpenAPI specification
type TransformEnvironmentMap struct {
}

// GetDeviceFleetReportRequest represents the GetDeviceFleetReportRequest schema from the OpenAPI specification
type GetDeviceFleetReportRequest struct {
	Devicefleetname interface{} `json:"DeviceFleetName"`
}

// GetLineageGroupPolicyResponse represents the GetLineageGroupPolicyResponse schema from the OpenAPI specification
type GetLineageGroupPolicyResponse struct {
	Lineagegrouparn interface{} `json:"LineageGroupArn,omitempty"`
	Resourcepolicy interface{} `json:"ResourcePolicy,omitempty"`
}

// ResourceCatalog represents the ResourceCatalog schema from the OpenAPI specification
type ResourceCatalog struct {
	Resourcecatalogname interface{} `json:"ResourceCatalogName"`
	Creationtime interface{} `json:"CreationTime"`
	Description interface{} `json:"Description"`
	Resourcecatalogarn interface{} `json:"ResourceCatalogArn"`
}

// ContextSource represents the ContextSource schema from the OpenAPI specification
type ContextSource struct {
	Sourceuri interface{} `json:"SourceUri"`
	Sourceid interface{} `json:"SourceId,omitempty"`
	Sourcetype interface{} `json:"SourceType,omitempty"`
}

// AutoMLPartialFailureReason represents the AutoMLPartialFailureReason schema from the OpenAPI specification
type AutoMLPartialFailureReason struct {
	Partialfailuremessage interface{} `json:"PartialFailureMessage,omitempty"`
}

// ModelVariantConfig represents the ModelVariantConfig schema from the OpenAPI specification
type ModelVariantConfig struct {
	Modelname interface{} `json:"ModelName"`
	Variantname interface{} `json:"VariantName"`
	Infrastructureconfig interface{} `json:"InfrastructureConfig"`
}

// ListHyperParameterTuningJobsRequest represents the ListHyperParameterTuningJobsRequest schema from the OpenAPI specification
type ListHyperParameterTuningJobsRequest struct {
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Lastmodifiedtimebefore interface{} `json:"LastModifiedTimeBefore,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Statusequals interface{} `json:"StatusEquals,omitempty"`
	Lastmodifiedtimeafter interface{} `json:"LastModifiedTimeAfter,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
}

// StopPipelineExecutionRequest represents the StopPipelineExecutionRequest schema from the OpenAPI specification
type StopPipelineExecutionRequest struct {
	Pipelineexecutionarn interface{} `json:"PipelineExecutionArn"`
	Clientrequesttoken interface{} `json:"ClientRequestToken"`
}

// MonitoringInput represents the MonitoringInput schema from the OpenAPI specification
type MonitoringInput struct {
	Endpointinput interface{} `json:"EndpointInput,omitempty"`
	Batchtransforminput interface{} `json:"BatchTransformInput,omitempty"`
}

// DeleteAppRequest represents the DeleteAppRequest schema from the OpenAPI specification
type DeleteAppRequest struct {
	Spacename interface{} `json:"SpaceName,omitempty"`
	Userprofilename interface{} `json:"UserProfileName,omitempty"`
	Appname interface{} `json:"AppName"`
	Apptype interface{} `json:"AppType"`
	Domainid interface{} `json:"DomainId"`
}

// DeleteNotebookInstanceLifecycleConfigInput represents the DeleteNotebookInstanceLifecycleConfigInput schema from the OpenAPI specification
type DeleteNotebookInstanceLifecycleConfigInput struct {
	Notebookinstancelifecycleconfigname interface{} `json:"NotebookInstanceLifecycleConfigName"`
}

// EdgeModel represents the EdgeModel schema from the OpenAPI specification
type EdgeModel struct {
	Latestinference interface{} `json:"LatestInference,omitempty"`
	Latestsampletime interface{} `json:"LatestSampleTime,omitempty"`
	Modelname interface{} `json:"ModelName"`
	Modelversion interface{} `json:"ModelVersion"`
}

// Project represents the Project schema from the OpenAPI specification
type Project struct {
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Projectname interface{} `json:"ProjectName,omitempty"`
	Servicecatalogprovisionedproductdetails ServiceCatalogProvisionedProductDetails `json:"ServiceCatalogProvisionedProductDetails,omitempty"` // Details of a provisioned service catalog product. For information about service catalog, see <a href="https://docs.aws.amazon.com/servicecatalog/latest/adminguide/introduction.html">What is Amazon Web Services Service Catalog</a>.
	Projectstatus interface{} `json:"ProjectStatus,omitempty"`
	Servicecatalogprovisioningdetails ServiceCatalogProvisioningDetails `json:"ServiceCatalogProvisioningDetails,omitempty"` // Details that you specify to provision a service catalog product. For information about service catalog, see <a href="https://docs.aws.amazon.com/servicecatalog/latest/adminguide/introduction.html">What is Amazon Web Services Service Catalog</a>.
	Createdby interface{} `json:"CreatedBy,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Projectdescription interface{} `json:"ProjectDescription,omitempty"`
	Projectid interface{} `json:"ProjectId,omitempty"`
	Lastmodifiedby UserContext `json:"LastModifiedBy,omitempty"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	Projectarn interface{} `json:"ProjectArn,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// CreateHubResponse represents the CreateHubResponse schema from the OpenAPI specification
type CreateHubResponse struct {
	Hubarn interface{} `json:"HubArn"`
}

// ProcessingJob represents the ProcessingJob schema from the OpenAPI specification
type ProcessingJob struct {
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Appspecification AppSpecification `json:"AppSpecification,omitempty"` // Configuration to run a processing job in a specified container image.
	Processingresources ProcessingResources `json:"ProcessingResources,omitempty"` // Identifies the resources, ML compute instances, and ML storage volumes to deploy for a processing job. In distributed training, you specify more than one instance.
	Processingendtime interface{} `json:"ProcessingEndTime,omitempty"`
	Monitoringschedulearn interface{} `json:"MonitoringScheduleArn,omitempty"`
	Processingstarttime interface{} `json:"ProcessingStartTime,omitempty"`
	Processingjobarn interface{} `json:"ProcessingJobArn,omitempty"`
	Processingoutputconfig ProcessingOutputConfig `json:"ProcessingOutputConfig,omitempty"` // Configuration for uploading output from the processing container.
	Networkconfig NetworkConfig `json:"NetworkConfig,omitempty"` // Networking options for a job, such as network traffic encryption between containers, whether to allow inbound and outbound network calls to and from containers, and the VPC subnets and security groups to use for VPC-enabled jobs.
	Processinginputs interface{} `json:"ProcessingInputs,omitempty"`
	Processingjobname interface{} `json:"ProcessingJobName,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Trainingjobarn interface{} `json:"TrainingJobArn,omitempty"`
	Exitmessage interface{} `json:"ExitMessage,omitempty"`
	Experimentconfig ExperimentConfig `json:"ExperimentConfig,omitempty"` // <p>Associates a SageMaker job as a trial component with an experiment and trial. Specified when you call the following APIs:</p> <ul> <li> <p> <a href="https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateProcessingJob.html">CreateProcessingJob</a> </p> </li> <li> <p> <a href="https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTrainingJob.html">CreateTrainingJob</a> </p> </li> <li> <p> <a href="https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTransformJob.html">CreateTransformJob</a> </p> </li> </ul>
	Environment interface{} `json:"Environment,omitempty"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Processingjobstatus interface{} `json:"ProcessingJobStatus,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Stoppingcondition ProcessingStoppingCondition `json:"StoppingCondition,omitempty"` // Configures conditions under which the processing job should be stopped, such as how long the processing job has been running. After the condition is met, the processing job is stopped.
	Automljobarn interface{} `json:"AutoMLJobArn,omitempty"`
}

// ListImageVersionsResponse represents the ListImageVersionsResponse schema from the OpenAPI specification
type ListImageVersionsResponse struct {
	Imageversions interface{} `json:"ImageVersions,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListImagesResponse represents the ListImagesResponse schema from the OpenAPI specification
type ListImagesResponse struct {
	Images interface{} `json:"Images,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// RetentionPolicy represents the RetentionPolicy schema from the OpenAPI specification
type RetentionPolicy struct {
	Homeefsfilesystem interface{} `json:"HomeEfsFileSystem,omitempty"`
}

// ClarifyShapConfig represents the ClarifyShapConfig schema from the OpenAPI specification
type ClarifyShapConfig struct {
	Numberofsamples interface{} `json:"NumberOfSamples,omitempty"`
	Seed interface{} `json:"Seed,omitempty"`
	Shapbaselineconfig interface{} `json:"ShapBaselineConfig"`
	Textconfig interface{} `json:"TextConfig,omitempty"`
	Uselogit interface{} `json:"UseLogit,omitempty"`
}

// SuggestionQuery represents the SuggestionQuery schema from the OpenAPI specification
type SuggestionQuery struct {
	Propertynamequery interface{} `json:"PropertyNameQuery,omitempty"`
}

// ImportHubContentRequest represents the ImportHubContentRequest schema from the OpenAPI specification
type ImportHubContentRequest struct {
	Hubcontentversion interface{} `json:"HubContentVersion,omitempty"`
	Hubname interface{} `json:"HubName"`
	Tags interface{} `json:"Tags,omitempty"`
	Hubcontentdisplayname interface{} `json:"HubContentDisplayName,omitempty"`
	Hubcontentdocument interface{} `json:"HubContentDocument"`
	Hubcontentmarkdown interface{} `json:"HubContentMarkdown,omitempty"`
	Hubcontentname interface{} `json:"HubContentName"`
	Hubcontenttype interface{} `json:"HubContentType"`
	Documentschemaversion interface{} `json:"DocumentSchemaVersion"`
	Hubcontentdescription interface{} `json:"HubContentDescription,omitempty"`
	Hubcontentsearchkeywords interface{} `json:"HubContentSearchKeywords,omitempty"`
}

// StopAutoMLJobRequest represents the StopAutoMLJobRequest schema from the OpenAPI specification
type StopAutoMLJobRequest struct {
	Automljobname interface{} `json:"AutoMLJobName"`
}

// HyperParameterTuningJobWarmStartConfig represents the HyperParameterTuningJobWarmStartConfig schema from the OpenAPI specification
type HyperParameterTuningJobWarmStartConfig struct {
	Parenthyperparametertuningjobs interface{} `json:"ParentHyperParameterTuningJobs"`
	Warmstarttype interface{} `json:"WarmStartType"`
}

// DeleteModelQualityJobDefinitionRequest represents the DeleteModelQualityJobDefinitionRequest schema from the OpenAPI specification
type DeleteModelQualityJobDefinitionRequest struct {
	Jobdefinitionname interface{} `json:"JobDefinitionName"`
}

// WarmPoolStatus represents the WarmPoolStatus schema from the OpenAPI specification
type WarmPoolStatus struct {
	Resourceretainedbillabletimeinseconds interface{} `json:"ResourceRetainedBillableTimeInSeconds,omitempty"`
	Reusedbyjob interface{} `json:"ReusedByJob,omitempty"`
	Status interface{} `json:"Status"`
}

// CreateHubRequest represents the CreateHubRequest schema from the OpenAPI specification
type CreateHubRequest struct {
	Tags interface{} `json:"Tags,omitempty"`
	Hubdescription interface{} `json:"HubDescription"`
	Hubdisplayname interface{} `json:"HubDisplayName,omitempty"`
	Hubname interface{} `json:"HubName"`
	Hubsearchkeywords interface{} `json:"HubSearchKeywords,omitempty"`
	S3storageconfig interface{} `json:"S3StorageConfig,omitempty"`
}

// StopLabelingJobRequest represents the StopLabelingJobRequest schema from the OpenAPI specification
type StopLabelingJobRequest struct {
	Labelingjobname interface{} `json:"LabelingJobName"`
}

// BatchTransformInput represents the BatchTransformInput schema from the OpenAPI specification
type BatchTransformInput struct {
	Probabilityattribute interface{} `json:"ProbabilityAttribute,omitempty"`
	Featuresattribute interface{} `json:"FeaturesAttribute,omitempty"`
	Probabilitythresholdattribute interface{} `json:"ProbabilityThresholdAttribute,omitempty"`
	Starttimeoffset interface{} `json:"StartTimeOffset,omitempty"`
	Datacaptureddestinations3uri interface{} `json:"DataCapturedDestinationS3Uri"`
	S3inputmode interface{} `json:"S3InputMode,omitempty"`
	Datasetformat interface{} `json:"DatasetFormat"`
	Endtimeoffset interface{} `json:"EndTimeOffset,omitempty"`
	Inferenceattribute interface{} `json:"InferenceAttribute,omitempty"`
	Localpath interface{} `json:"LocalPath"`
	S3datadistributiontype interface{} `json:"S3DataDistributionType,omitempty"`
}

// MemberDefinition represents the MemberDefinition schema from the OpenAPI specification
type MemberDefinition struct {
	Cognitomemberdefinition interface{} `json:"CognitoMemberDefinition,omitempty"`
	Oidcmemberdefinition interface{} `json:"OidcMemberDefinition,omitempty"`
}

// DescribeFeatureGroupResponse represents the DescribeFeatureGroupResponse schema from the OpenAPI specification
type DescribeFeatureGroupResponse struct {
	Description interface{} `json:"Description,omitempty"`
	Eventtimefeaturename interface{} `json:"EventTimeFeatureName"`
	Featuregroupname interface{} `json:"FeatureGroupName"`
	Offlinestoreconfig interface{} `json:"OfflineStoreConfig,omitempty"`
	Nexttoken interface{} `json:"NextToken"`
	Onlinestoretotalsizebytes interface{} `json:"OnlineStoreTotalSizeBytes,omitempty"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Featuregrouparn interface{} `json:"FeatureGroupArn"`
	Featuredefinitions interface{} `json:"FeatureDefinitions"`
	Lastupdatestatus interface{} `json:"LastUpdateStatus,omitempty"`
	Offlinestorestatus interface{} `json:"OfflineStoreStatus,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Creationtime interface{} `json:"CreationTime"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Recordidentifierfeaturename interface{} `json:"RecordIdentifierFeatureName"`
	Featuregroupstatus interface{} `json:"FeatureGroupStatus,omitempty"`
	Onlinestoreconfig interface{} `json:"OnlineStoreConfig,omitempty"`
}

// UpdatePipelineRequest represents the UpdatePipelineRequest schema from the OpenAPI specification
type UpdatePipelineRequest struct {
	Parallelismconfiguration interface{} `json:"ParallelismConfiguration,omitempty"`
	Pipelinedefinition interface{} `json:"PipelineDefinition,omitempty"`
	Pipelinedefinitions3location interface{} `json:"PipelineDefinitionS3Location,omitempty"`
	Pipelinedescription interface{} `json:"PipelineDescription,omitempty"`
	Pipelinedisplayname interface{} `json:"PipelineDisplayName,omitempty"`
	Pipelinename interface{} `json:"PipelineName"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
}

// ModelInput represents the ModelInput schema from the OpenAPI specification
type ModelInput struct {
	Datainputconfig interface{} `json:"DataInputConfig"`
}

// ListModelBiasJobDefinitionsResponse represents the ListModelBiasJobDefinitionsResponse schema from the OpenAPI specification
type ListModelBiasJobDefinitionsResponse struct {
	Jobdefinitionsummaries interface{} `json:"JobDefinitionSummaries"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DescribeStudioLifecycleConfigRequest represents the DescribeStudioLifecycleConfigRequest schema from the OpenAPI specification
type DescribeStudioLifecycleConfigRequest struct {
	Studiolifecycleconfigname interface{} `json:"StudioLifecycleConfigName"`
}

// ListPipelineParametersForExecutionRequest represents the ListPipelineParametersForExecutionRequest schema from the OpenAPI specification
type ListPipelineParametersForExecutionRequest struct {
	Pipelineexecutionarn interface{} `json:"PipelineExecutionArn"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// AutoMLDataSource represents the AutoMLDataSource schema from the OpenAPI specification
type AutoMLDataSource struct {
	S3datasource interface{} `json:"S3DataSource"`
}

// ModelQualityBaselineConfig represents the ModelQualityBaselineConfig schema from the OpenAPI specification
type ModelQualityBaselineConfig struct {
	Constraintsresource MonitoringConstraintsResource `json:"ConstraintsResource,omitempty"` // The constraints resource for a monitoring job.
	Baseliningjobname interface{} `json:"BaseliningJobName,omitempty"`
}

// MonitoringParquetDatasetFormat represents the MonitoringParquetDatasetFormat schema from the OpenAPI specification
type MonitoringParquetDatasetFormat struct {
}

// ModelDeployResult represents the ModelDeployResult schema from the OpenAPI specification
type ModelDeployResult struct {
	Endpointname interface{} `json:"EndpointName,omitempty"`
}

// ResourceLimits represents the ResourceLimits schema from the OpenAPI specification
type ResourceLimits struct {
	Maxparalleltrainingjobs interface{} `json:"MaxParallelTrainingJobs"`
	Maxruntimeinseconds interface{} `json:"MaxRuntimeInSeconds,omitempty"`
	Maxnumberoftrainingjobs interface{} `json:"MaxNumberOfTrainingJobs,omitempty"`
}

// S3StorageConfig represents the S3StorageConfig schema from the OpenAPI specification
type S3StorageConfig struct {
	S3uri interface{} `json:"S3Uri"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Resolvedoutputs3uri interface{} `json:"ResolvedOutputS3Uri,omitempty"`
}

// ListHubContentsResponse represents the ListHubContentsResponse schema from the OpenAPI specification
type ListHubContentsResponse struct {
	Hubcontentsummaries interface{} `json:"HubContentSummaries"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// HyperParameterTuningJobStrategyConfig represents the HyperParameterTuningJobStrategyConfig schema from the OpenAPI specification
type HyperParameterTuningJobStrategyConfig struct {
	Hyperbandstrategyconfig interface{} `json:"HyperbandStrategyConfig,omitempty"`
}

// PipelineExecutionSummary represents the PipelineExecutionSummary schema from the OpenAPI specification
type PipelineExecutionSummary struct {
	Pipelineexecutionfailurereason interface{} `json:"PipelineExecutionFailureReason,omitempty"`
	Pipelineexecutionstatus interface{} `json:"PipelineExecutionStatus,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
	Pipelineexecutionarn interface{} `json:"PipelineExecutionArn,omitempty"`
	Pipelineexecutiondescription interface{} `json:"PipelineExecutionDescription,omitempty"`
	Pipelineexecutiondisplayname interface{} `json:"PipelineExecutionDisplayName,omitempty"`
}

// BatchDescribeModelPackageInput represents the BatchDescribeModelPackageInput schema from the OpenAPI specification
type BatchDescribeModelPackageInput struct {
	Modelpackagearnlist interface{} `json:"ModelPackageArnList"`
}

// ListPipelineExecutionStepsRequest represents the ListPipelineExecutionStepsRequest schema from the OpenAPI specification
type ListPipelineExecutionStepsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Pipelineexecutionarn interface{} `json:"PipelineExecutionArn,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
}

// ListModelCardVersionsResponse represents the ListModelCardVersionsResponse schema from the OpenAPI specification
type ListModelCardVersionsResponse struct {
	Modelcardversionsummarylist interface{} `json:"ModelCardVersionSummaryList"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// UpdateArtifactResponse represents the UpdateArtifactResponse schema from the OpenAPI specification
type UpdateArtifactResponse struct {
	Artifactarn interface{} `json:"ArtifactArn,omitempty"`
}

// DescribeStudioLifecycleConfigResponse represents the DescribeStudioLifecycleConfigResponse schema from the OpenAPI specification
type DescribeStudioLifecycleConfigResponse struct {
	Studiolifecycleconfigapptype interface{} `json:"StudioLifecycleConfigAppType,omitempty"`
	Studiolifecycleconfigarn interface{} `json:"StudioLifecycleConfigArn,omitempty"`
	Studiolifecycleconfigcontent interface{} `json:"StudioLifecycleConfigContent,omitempty"`
	Studiolifecycleconfigname interface{} `json:"StudioLifecycleConfigName,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
}

// UpdatePipelineExecutionRequest represents the UpdatePipelineExecutionRequest schema from the OpenAPI specification
type UpdatePipelineExecutionRequest struct {
	Pipelineexecutiondisplayname interface{} `json:"PipelineExecutionDisplayName,omitempty"`
	Parallelismconfiguration interface{} `json:"ParallelismConfiguration,omitempty"`
	Pipelineexecutionarn interface{} `json:"PipelineExecutionArn"`
	Pipelineexecutiondescription interface{} `json:"PipelineExecutionDescription,omitempty"`
}

// ModelCard represents the ModelCard schema from the OpenAPI specification
type ModelCard struct {
	Createdby UserContext `json:"CreatedBy,omitempty"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	Lastmodifiedby UserContext `json:"LastModifiedBy,omitempty"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	Modelpackagegroupname interface{} `json:"ModelPackageGroupName,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Modelcardarn interface{} `json:"ModelCardArn,omitempty"`
	Riskrating interface{} `json:"RiskRating,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Modelcardname interface{} `json:"ModelCardName,omitempty"`
	Securityconfig interface{} `json:"SecurityConfig,omitempty"`
	Content interface{} `json:"Content,omitempty"`
	Modelcardstatus interface{} `json:"ModelCardStatus,omitempty"`
	Modelcardversion interface{} `json:"ModelCardVersion,omitempty"`
	Modelid interface{} `json:"ModelId,omitempty"`
}

// HookParameters represents the HookParameters schema from the OpenAPI specification
type HookParameters struct {
}

// DeleteEndpointInput represents the DeleteEndpointInput schema from the OpenAPI specification
type DeleteEndpointInput struct {
	Endpointname interface{} `json:"EndpointName"`
}

// TrialComponentSimpleSummary represents the TrialComponentSimpleSummary schema from the OpenAPI specification
type TrialComponentSimpleSummary struct {
	Trialcomponentarn interface{} `json:"TrialComponentArn,omitempty"`
	Trialcomponentname interface{} `json:"TrialComponentName,omitempty"`
	Trialcomponentsource TrialComponentSource `json:"TrialComponentSource,omitempty"` // The Amazon Resource Name (ARN) and job type of the source of a trial component.
	Createdby UserContext `json:"CreatedBy,omitempty"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	Creationtime interface{} `json:"CreationTime,omitempty"`
}

// DriftCheckExplainability represents the DriftCheckExplainability schema from the OpenAPI specification
type DriftCheckExplainability struct {
	Configfile interface{} `json:"ConfigFile,omitempty"`
	Constraints interface{} `json:"Constraints,omitempty"`
}

// ImageClassificationJobConfig represents the ImageClassificationJobConfig schema from the OpenAPI specification
type ImageClassificationJobConfig struct {
	Completioncriteria interface{} `json:"CompletionCriteria,omitempty"`
}

// UpdateImageResponse represents the UpdateImageResponse schema from the OpenAPI specification
type UpdateImageResponse struct {
	Imagearn interface{} `json:"ImageArn,omitempty"`
}

// ModelLatencyThreshold represents the ModelLatencyThreshold schema from the OpenAPI specification
type ModelLatencyThreshold struct {
	Percentile interface{} `json:"Percentile,omitempty"`
	Valueinmilliseconds interface{} `json:"ValueInMilliseconds,omitempty"`
}

// InferenceRecommendationsJob represents the InferenceRecommendationsJob schema from the OpenAPI specification
type InferenceRecommendationsJob struct {
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Jobtype interface{} `json:"JobType"`
	Modelpackageversionarn interface{} `json:"ModelPackageVersionArn,omitempty"`
	Jobname interface{} `json:"JobName"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime"`
	Samplepayloadurl interface{} `json:"SamplePayloadUrl,omitempty"`
	Jobarn interface{} `json:"JobArn"`
	Modelname interface{} `json:"ModelName,omitempty"`
	Status interface{} `json:"Status"`
	Completiontime interface{} `json:"CompletionTime,omitempty"`
	Creationtime interface{} `json:"CreationTime"`
	Jobdescription interface{} `json:"JobDescription"`
	Rolearn interface{} `json:"RoleArn"`
}

// GetSagemakerServicecatalogPortfolioStatusInput represents the GetSagemakerServicecatalogPortfolioStatusInput schema from the OpenAPI specification
type GetSagemakerServicecatalogPortfolioStatusInput struct {
}

// DescribeEndpointConfigOutput represents the DescribeEndpointConfigOutput schema from the OpenAPI specification
type DescribeEndpointConfigOutput struct {
	Productionvariants interface{} `json:"ProductionVariants"`
	Shadowproductionvariants interface{} `json:"ShadowProductionVariants,omitempty"`
	Endpointconfigname interface{} `json:"EndpointConfigName"`
	Explainerconfig interface{} `json:"ExplainerConfig,omitempty"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Asyncinferenceconfig interface{} `json:"AsyncInferenceConfig,omitempty"`
	Creationtime interface{} `json:"CreationTime"`
	Datacaptureconfig DataCaptureConfig `json:"DataCaptureConfig,omitempty"` // Configuration to control how SageMaker captures inference data.
	Endpointconfigarn interface{} `json:"EndpointConfigArn"`
}

// ListFeatureGroupsRequest represents the ListFeatureGroupsRequest schema from the OpenAPI specification
type ListFeatureGroupsRequest struct {
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Featuregroupstatusequals interface{} `json:"FeatureGroupStatusEquals,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Offlinestorestatusequals interface{} `json:"OfflineStoreStatusEquals,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
}

// TransformJobStepMetadata represents the TransformJobStepMetadata schema from the OpenAPI specification
type TransformJobStepMetadata struct {
	Arn interface{} `json:"Arn,omitempty"`
}

// MonitoringAlertHistorySummary represents the MonitoringAlertHistorySummary schema from the OpenAPI specification
type MonitoringAlertHistorySummary struct {
	Creationtime interface{} `json:"CreationTime"`
	Monitoringalertname interface{} `json:"MonitoringAlertName"`
	Monitoringschedulename interface{} `json:"MonitoringScheduleName"`
	Alertstatus interface{} `json:"AlertStatus"`
}

// DescribeEdgePackagingJobRequest represents the DescribeEdgePackagingJobRequest schema from the OpenAPI specification
type DescribeEdgePackagingJobRequest struct {
	Edgepackagingjobname interface{} `json:"EdgePackagingJobName"`
}

// S3ModelDataSource represents the S3ModelDataSource schema from the OpenAPI specification
type S3ModelDataSource struct {
	Compressiontype interface{} `json:"CompressionType"`
	S3datatype interface{} `json:"S3DataType"`
	S3uri interface{} `json:"S3Uri"`
}

// DeleteImageRequest represents the DeleteImageRequest schema from the OpenAPI specification
type DeleteImageRequest struct {
	Imagename interface{} `json:"ImageName"`
}

// ListInferenceRecommendationsJobStepsResponse represents the ListInferenceRecommendationsJobStepsResponse schema from the OpenAPI specification
type ListInferenceRecommendationsJobStepsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Steps interface{} `json:"Steps,omitempty"`
}

// WorkspaceSettings represents the WorkspaceSettings schema from the OpenAPI specification
type WorkspaceSettings struct {
	S3artifactpath interface{} `json:"S3ArtifactPath,omitempty"`
	S3kmskeyid interface{} `json:"S3KmsKeyId,omitempty"`
}

// DeleteHumanTaskUiResponse represents the DeleteHumanTaskUiResponse schema from the OpenAPI specification
type DeleteHumanTaskUiResponse struct {
}

// Parent represents the Parent schema from the OpenAPI specification
type Parent struct {
	Experimentname interface{} `json:"ExperimentName,omitempty"`
	Trialname interface{} `json:"TrialName,omitempty"`
}

// ModelInfrastructureConfig represents the ModelInfrastructureConfig schema from the OpenAPI specification
type ModelInfrastructureConfig struct {
	Realtimeinferenceconfig interface{} `json:"RealTimeInferenceConfig"`
	Infrastructuretype interface{} `json:"InfrastructureType"`
}

// ListTrialComponentsRequest represents the ListTrialComponentsRequest schema from the OpenAPI specification
type ListTrialComponentsRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Trialname interface{} `json:"TrialName,omitempty"`
	Createdafter interface{} `json:"CreatedAfter,omitempty"`
	Sourcearn interface{} `json:"SourceArn,omitempty"`
	Createdbefore interface{} `json:"CreatedBefore,omitempty"`
	Experimentname interface{} `json:"ExperimentName,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// Vertex represents the Vertex schema from the OpenAPI specification
type Vertex struct {
	Lineagetype interface{} `json:"LineageType,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
}

// UpdateAppImageConfigRequest represents the UpdateAppImageConfigRequest schema from the OpenAPI specification
type UpdateAppImageConfigRequest struct {
	Appimageconfigname interface{} `json:"AppImageConfigName"`
	Kernelgatewayimageconfig interface{} `json:"KernelGatewayImageConfig,omitempty"`
}

// DescribeTrainingJobRequest represents the DescribeTrainingJobRequest schema from the OpenAPI specification
type DescribeTrainingJobRequest struct {
	Trainingjobname interface{} `json:"TrainingJobName"`
}

// UpdateWorkforceResponse represents the UpdateWorkforceResponse schema from the OpenAPI specification
type UpdateWorkforceResponse struct {
	Workforce interface{} `json:"Workforce"`
}

// RSessionAppSettings represents the RSessionAppSettings schema from the OpenAPI specification
type RSessionAppSettings struct {
	Customimages interface{} `json:"CustomImages,omitempty"`
	Defaultresourcespec ResourceSpec `json:"DefaultResourceSpec,omitempty"` // Specifies the ARN's of a SageMaker image and SageMaker image version, and the instance type that the version runs on.
}

// ModelPackageStatusItem represents the ModelPackageStatusItem schema from the OpenAPI specification
type ModelPackageStatusItem struct {
	Name interface{} `json:"Name"`
	Status interface{} `json:"Status"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
}

// OutputParameter represents the OutputParameter schema from the OpenAPI specification
type OutputParameter struct {
	Name interface{} `json:"Name"`
	Value interface{} `json:"Value"`
}

// SendPipelineExecutionStepSuccessResponse represents the SendPipelineExecutionStepSuccessResponse schema from the OpenAPI specification
type SendPipelineExecutionStepSuccessResponse struct {
	Pipelineexecutionarn interface{} `json:"PipelineExecutionArn,omitempty"`
}

// CreateExperimentResponse represents the CreateExperimentResponse schema from the OpenAPI specification
type CreateExperimentResponse struct {
	Experimentarn interface{} `json:"ExperimentArn,omitempty"`
}

// EnableSagemakerServicecatalogPortfolioInput represents the EnableSagemakerServicecatalogPortfolioInput schema from the OpenAPI specification
type EnableSagemakerServicecatalogPortfolioInput struct {
}

// AutoMLCandidate represents the AutoMLCandidate schema from the OpenAPI specification
type AutoMLCandidate struct {
	Objectivestatus interface{} `json:"ObjectiveStatus"`
	Candidatestatus interface{} `json:"CandidateStatus"`
	Endtime interface{} `json:"EndTime,omitempty"`
	Finalautomljobobjectivemetric FinalAutoMLJobObjectiveMetric `json:"FinalAutoMLJobObjectiveMetric,omitempty"` // The best candidate result from an AutoML training job.
	Candidateproperties interface{} `json:"CandidateProperties,omitempty"`
	Candidatesteps interface{} `json:"CandidateSteps"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Inferencecontainerdefinitions interface{} `json:"InferenceContainerDefinitions,omitempty"`
	Inferencecontainers interface{} `json:"InferenceContainers,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime"`
	Candidatename interface{} `json:"CandidateName"`
	Creationtime interface{} `json:"CreationTime"`
}

// CreateMonitoringScheduleRequest represents the CreateMonitoringScheduleRequest schema from the OpenAPI specification
type CreateMonitoringScheduleRequest struct {
	Monitoringscheduleconfig interface{} `json:"MonitoringScheduleConfig"`
	Monitoringschedulename interface{} `json:"MonitoringScheduleName"`
	Tags interface{} `json:"Tags,omitempty"`
}

// ListWorkforcesResponse represents the ListWorkforcesResponse schema from the OpenAPI specification
type ListWorkforcesResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Workforces interface{} `json:"Workforces"`
}

// Experiment represents the Experiment schema from the OpenAPI specification
type Experiment struct {
	Description interface{} `json:"Description,omitempty"`
	Experimentname interface{} `json:"ExperimentName,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Source ExperimentSource `json:"Source,omitempty"` // The source of the experiment.
	Tags interface{} `json:"Tags,omitempty"`
	Displayname interface{} `json:"DisplayName,omitempty"`
	Experimentarn interface{} `json:"ExperimentArn,omitempty"`
	Lastmodifiedby UserContext `json:"LastModifiedBy,omitempty"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	Createdby interface{} `json:"CreatedBy,omitempty"`
}

// ProcessingOutput represents the ProcessingOutput schema from the OpenAPI specification
type ProcessingOutput struct {
	Featurestoreoutput interface{} `json:"FeatureStoreOutput,omitempty"`
	Outputname interface{} `json:"OutputName"`
	S3output interface{} `json:"S3Output,omitempty"`
	Appmanaged interface{} `json:"AppManaged,omitempty"`
}

// LabelingJobSummary represents the LabelingJobSummary schema from the OpenAPI specification
type LabelingJobSummary struct {
	Creationtime interface{} `json:"CreationTime"`
	Inputconfig interface{} `json:"InputConfig,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Labelingjobarn interface{} `json:"LabelingJobArn"`
	Labelingjobname interface{} `json:"LabelingJobName"`
	Labelingjoboutput interface{} `json:"LabelingJobOutput,omitempty"`
	Labelingjobstatus interface{} `json:"LabelingJobStatus"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime"`
	Prehumantasklambdaarn interface{} `json:"PreHumanTaskLambdaArn"`
	Labelcounters interface{} `json:"LabelCounters"`
	Workteamarn interface{} `json:"WorkteamArn"`
	Annotationconsolidationlambdaarn interface{} `json:"AnnotationConsolidationLambdaArn,omitempty"`
}

// TabularResolvedAttributes represents the TabularResolvedAttributes schema from the OpenAPI specification
type TabularResolvedAttributes struct {
	Problemtype interface{} `json:"ProblemType,omitempty"`
}

// FeatureGroupSummary represents the FeatureGroupSummary schema from the OpenAPI specification
type FeatureGroupSummary struct {
	Featuregroupstatus interface{} `json:"FeatureGroupStatus,omitempty"`
	Offlinestorestatus interface{} `json:"OfflineStoreStatus,omitempty"`
	Creationtime interface{} `json:"CreationTime"`
	Featuregrouparn interface{} `json:"FeatureGroupArn"`
	Featuregroupname interface{} `json:"FeatureGroupName"`
}

// UserSettings represents the UserSettings schema from the OpenAPI specification
type UserSettings struct {
	Tensorboardappsettings interface{} `json:"TensorBoardAppSettings,omitempty"`
	Kernelgatewayappsettings interface{} `json:"KernelGatewayAppSettings,omitempty"`
	Securitygroups interface{} `json:"SecurityGroups,omitempty"`
	Canvasappsettings interface{} `json:"CanvasAppSettings,omitempty"`
	Jupyterserverappsettings interface{} `json:"JupyterServerAppSettings,omitempty"`
	Rstudioserverproappsettings interface{} `json:"RStudioServerProAppSettings,omitempty"`
	Executionrole interface{} `json:"ExecutionRole,omitempty"`
	Rsessionappsettings interface{} `json:"RSessionAppSettings,omitempty"`
	Sharingsettings interface{} `json:"SharingSettings,omitempty"`
}

// AssociationSummary represents the AssociationSummary schema from the OpenAPI specification
type AssociationSummary struct {
	Sourcearn interface{} `json:"SourceArn,omitempty"`
	Associationtype interface{} `json:"AssociationType,omitempty"`
	Sourcetype interface{} `json:"SourceType,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Destinationname interface{} `json:"DestinationName,omitempty"`
	Destinationtype interface{} `json:"DestinationType,omitempty"`
	Sourcename interface{} `json:"SourceName,omitempty"`
	Createdby UserContext `json:"CreatedBy,omitempty"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	Destinationarn interface{} `json:"DestinationArn,omitempty"`
}

// ModelSummary represents the ModelSummary schema from the OpenAPI specification
type ModelSummary struct {
	Modelarn interface{} `json:"ModelArn"`
	Modelname interface{} `json:"ModelName"`
	Creationtime interface{} `json:"CreationTime"`
}

// DescribeDeviceFleetRequest represents the DescribeDeviceFleetRequest schema from the OpenAPI specification
type DescribeDeviceFleetRequest struct {
	Devicefleetname interface{} `json:"DeviceFleetName"`
}

// AutoMLJobChannel represents the AutoMLJobChannel schema from the OpenAPI specification
type AutoMLJobChannel struct {
	Datasource interface{} `json:"DataSource,omitempty"`
	Channeltype interface{} `json:"ChannelType,omitempty"`
	Compressiontype interface{} `json:"CompressionType,omitempty"`
	Contenttype interface{} `json:"ContentType,omitempty"`
}

// ListEndpointsInput represents the ListEndpointsInput schema from the OpenAPI specification
type ListEndpointsInput struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Lastmodifiedtimeafter interface{} `json:"LastModifiedTimeAfter,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
	Statusequals interface{} `json:"StatusEquals,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Lastmodifiedtimebefore interface{} `json:"LastModifiedTimeBefore,omitempty"`
}

// ListLabelingJobsForWorkteamRequest represents the ListLabelingJobsForWorkteamRequest schema from the OpenAPI specification
type ListLabelingJobsForWorkteamRequest struct {
	Jobreferencecodecontains interface{} `json:"JobReferenceCodeContains,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Workteamarn interface{} `json:"WorkteamArn"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
}

// TrialSummary represents the TrialSummary schema from the OpenAPI specification
type TrialSummary struct {
	Trialarn interface{} `json:"TrialArn,omitempty"`
	Trialname interface{} `json:"TrialName,omitempty"`
	Trialsource TrialSource `json:"TrialSource,omitempty"` // The source of the trial.
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Displayname interface{} `json:"DisplayName,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
}

// RepositoryAuthConfig represents the RepositoryAuthConfig schema from the OpenAPI specification
type RepositoryAuthConfig struct {
	Repositorycredentialsproviderarn interface{} `json:"RepositoryCredentialsProviderArn"`
}

// ModelCardSecurityConfig represents the ModelCardSecurityConfig schema from the OpenAPI specification
type ModelCardSecurityConfig struct {
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
}

// FillingTransformationMap represents the FillingTransformationMap schema from the OpenAPI specification
type FillingTransformationMap struct {
}

// GetScalingConfigurationRecommendationResponse represents the GetScalingConfigurationRecommendationResponse schema from the OpenAPI specification
type GetScalingConfigurationRecommendationResponse struct {
	Metric interface{} `json:"Metric,omitempty"`
	Recommendationid interface{} `json:"RecommendationId,omitempty"`
	Scalingpolicyobjective interface{} `json:"ScalingPolicyObjective,omitempty"`
	Targetcpuutilizationpercore interface{} `json:"TargetCpuUtilizationPerCore,omitempty"`
	Dynamicscalingconfiguration interface{} `json:"DynamicScalingConfiguration,omitempty"`
	Endpointname interface{} `json:"EndpointName,omitempty"`
	Inferencerecommendationsjobname interface{} `json:"InferenceRecommendationsJobName,omitempty"`
}

// ListCompilationJobsResponse represents the ListCompilationJobsResponse schema from the OpenAPI specification
type ListCompilationJobsResponse struct {
	Compilationjobsummaries interface{} `json:"CompilationJobSummaries"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DescribeFeatureGroupRequest represents the DescribeFeatureGroupRequest schema from the OpenAPI specification
type DescribeFeatureGroupRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Featuregroupname interface{} `json:"FeatureGroupName"`
}

// AssociateTrialComponentRequest represents the AssociateTrialComponentRequest schema from the OpenAPI specification
type AssociateTrialComponentRequest struct {
	Trialcomponentname interface{} `json:"TrialComponentName"`
	Trialname interface{} `json:"TrialName"`
}

// StopTransformJobRequest represents the StopTransformJobRequest schema from the OpenAPI specification
type StopTransformJobRequest struct {
	Transformjobname interface{} `json:"TransformJobName"`
}

// ListModelCardsResponse represents the ListModelCardsResponse schema from the OpenAPI specification
type ListModelCardsResponse struct {
	Modelcardsummaries interface{} `json:"ModelCardSummaries"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// MetadataProperties represents the MetadataProperties schema from the OpenAPI specification
type MetadataProperties struct {
	Repository interface{} `json:"Repository,omitempty"`
	Commitid interface{} `json:"CommitId,omitempty"`
	Generatedby interface{} `json:"GeneratedBy,omitempty"`
	Projectid interface{} `json:"ProjectId,omitempty"`
}

// ContinuousParameterRange represents the ContinuousParameterRange schema from the OpenAPI specification
type ContinuousParameterRange struct {
	Minvalue interface{} `json:"MinValue"`
	Name interface{} `json:"Name"`
	Scalingtype interface{} `json:"ScalingType,omitempty"`
	Maxvalue interface{} `json:"MaxValue"`
}

// UpdateActionResponse represents the UpdateActionResponse schema from the OpenAPI specification
type UpdateActionResponse struct {
	Actionarn interface{} `json:"ActionArn,omitempty"`
}

// DeleteExperimentResponse represents the DeleteExperimentResponse schema from the OpenAPI specification
type DeleteExperimentResponse struct {
	Experimentarn interface{} `json:"ExperimentArn,omitempty"`
}

// CreateEdgeDeploymentPlanRequest represents the CreateEdgeDeploymentPlanRequest schema from the OpenAPI specification
type CreateEdgeDeploymentPlanRequest struct {
	Devicefleetname interface{} `json:"DeviceFleetName"`
	Edgedeploymentplanname interface{} `json:"EdgeDeploymentPlanName"`
	Modelconfigs interface{} `json:"ModelConfigs"`
	Stages interface{} `json:"Stages,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// DeviceSummary represents the DeviceSummary schema from the OpenAPI specification
type DeviceSummary struct {
	Devicearn interface{} `json:"DeviceArn"`
	Devicefleetname interface{} `json:"DeviceFleetName,omitempty"`
	Devicename interface{} `json:"DeviceName"`
	Iotthingname interface{} `json:"IotThingName,omitempty"`
	Registrationtime interface{} `json:"RegistrationTime,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Agentversion interface{} `json:"AgentVersion,omitempty"`
	Latestheartbeat interface{} `json:"LatestHeartbeat,omitempty"`
	Models interface{} `json:"Models,omitempty"`
}

// SelectiveExecutionResult represents the SelectiveExecutionResult schema from the OpenAPI specification
type SelectiveExecutionResult struct {
	Sourcepipelineexecutionarn interface{} `json:"SourcePipelineExecutionArn,omitempty"`
}

// DescribeHyperParameterTuningJobRequest represents the DescribeHyperParameterTuningJobRequest schema from the OpenAPI specification
type DescribeHyperParameterTuningJobRequest struct {
	Hyperparametertuningjobname interface{} `json:"HyperParameterTuningJobName"`
}

// ClarifyCheckStepMetadata represents the ClarifyCheckStepMetadata schema from the OpenAPI specification
type ClarifyCheckStepMetadata struct {
	Skipcheck interface{} `json:"SkipCheck,omitempty"`
	Violationreport interface{} `json:"ViolationReport,omitempty"`
	Baselineusedfordriftcheckconstraints interface{} `json:"BaselineUsedForDriftCheckConstraints,omitempty"`
	Calculatedbaselineconstraints interface{} `json:"CalculatedBaselineConstraints,omitempty"`
	Checkjobarn interface{} `json:"CheckJobArn,omitempty"`
	Checktype interface{} `json:"CheckType,omitempty"`
	Modelpackagegroupname interface{} `json:"ModelPackageGroupName,omitempty"`
	Registernewbaseline interface{} `json:"RegisterNewBaseline,omitempty"`
}

// DeleteTrialComponentRequest represents the DeleteTrialComponentRequest schema from the OpenAPI specification
type DeleteTrialComponentRequest struct {
	Trialcomponentname interface{} `json:"TrialComponentName"`
}

// ListEndpointsOutput represents the ListEndpointsOutput schema from the OpenAPI specification
type ListEndpointsOutput struct {
	Endpoints interface{} `json:"Endpoints"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DescribeImageRequest represents the DescribeImageRequest schema from the OpenAPI specification
type DescribeImageRequest struct {
	Imagename interface{} `json:"ImageName"`
}

// LabelingJobDataAttributes represents the LabelingJobDataAttributes schema from the OpenAPI specification
type LabelingJobDataAttributes struct {
	Contentclassifiers interface{} `json:"ContentClassifiers,omitempty"`
}

// ScalingPolicy represents the ScalingPolicy schema from the OpenAPI specification
type ScalingPolicy struct {
	Targettracking interface{} `json:"TargetTracking,omitempty"`
}

// Image represents the Image schema from the OpenAPI specification
type Image struct {
	Imagename interface{} `json:"ImageName"`
	Imagestatus interface{} `json:"ImageStatus"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime"`
	Creationtime interface{} `json:"CreationTime"`
	Description interface{} `json:"Description,omitempty"`
	Displayname interface{} `json:"DisplayName,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Imagearn interface{} `json:"ImageArn"`
}

// DataCatalogConfig represents the DataCatalogConfig schema from the OpenAPI specification
type DataCatalogConfig struct {
	Tablename interface{} `json:"TableName"`
	Catalog interface{} `json:"Catalog"`
	Database interface{} `json:"Database"`
}

// DescribeInferenceExperimentResponse represents the DescribeInferenceExperimentResponse schema from the OpenAPI specification
type DescribeInferenceExperimentResponse struct {
	TypeField interface{} `json:"Type"`
	Modelvariants interface{} `json:"ModelVariants"`
	Statusreason interface{} `json:"StatusReason,omitempty"`
	Shadowmodeconfig interface{} `json:"ShadowModeConfig,omitempty"`
	Completiontime interface{} `json:"CompletionTime,omitempty"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Status interface{} `json:"Status"`
	Kmskey interface{} `json:"KmsKey,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Name interface{} `json:"Name"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Endpointmetadata interface{} `json:"EndpointMetadata"`
	Schedule interface{} `json:"Schedule,omitempty"`
	Arn interface{} `json:"Arn"`
	Datastorageconfig interface{} `json:"DataStorageConfig,omitempty"`
}

// CreateWorkteamResponse represents the CreateWorkteamResponse schema from the OpenAPI specification
type CreateWorkteamResponse struct {
	Workteamarn interface{} `json:"WorkteamArn,omitempty"`
}

// ProfilerRuleConfiguration represents the ProfilerRuleConfiguration schema from the OpenAPI specification
type ProfilerRuleConfiguration struct {
	Ruleevaluatorimage interface{} `json:"RuleEvaluatorImage"`
	Ruleparameters interface{} `json:"RuleParameters,omitempty"`
	S3outputpath interface{} `json:"S3OutputPath,omitempty"`
	Volumesizeingb interface{} `json:"VolumeSizeInGB,omitempty"`
	Instancetype interface{} `json:"InstanceType,omitempty"`
	Localpath interface{} `json:"LocalPath,omitempty"`
	Ruleconfigurationname interface{} `json:"RuleConfigurationName"`
}

// BatchDescribeModelPackageOutput represents the BatchDescribeModelPackageOutput schema from the OpenAPI specification
type BatchDescribeModelPackageOutput struct {
	Batchdescribemodelpackageerrormap interface{} `json:"BatchDescribeModelPackageErrorMap,omitempty"`
	Modelpackagesummaries interface{} `json:"ModelPackageSummaries,omitempty"`
}

// RegisterDevicesRequest represents the RegisterDevicesRequest schema from the OpenAPI specification
type RegisterDevicesRequest struct {
	Devices interface{} `json:"Devices"`
	Tags interface{} `json:"Tags,omitempty"`
	Devicefleetname interface{} `json:"DeviceFleetName"`
}

// HumanLoopActivationConditionsConfig represents the HumanLoopActivationConditionsConfig schema from the OpenAPI specification
type HumanLoopActivationConditionsConfig struct {
	Humanloopactivationconditions interface{} `json:"HumanLoopActivationConditions"`
}

// CognitoConfig represents the CognitoConfig schema from the OpenAPI specification
type CognitoConfig struct {
	Clientid interface{} `json:"ClientId"`
	Userpool interface{} `json:"UserPool"`
}

// StartInferenceExperimentResponse represents the StartInferenceExperimentResponse schema from the OpenAPI specification
type StartInferenceExperimentResponse struct {
	Inferenceexperimentarn interface{} `json:"InferenceExperimentArn"`
}

// ListTransformJobsResponse represents the ListTransformJobsResponse schema from the OpenAPI specification
type ListTransformJobsResponse struct {
	Transformjobsummaries interface{} `json:"TransformJobSummaries"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// UiConfig represents the UiConfig schema from the OpenAPI specification
type UiConfig struct {
	Humantaskuiarn interface{} `json:"HumanTaskUiArn,omitempty"`
	Uitemplates3uri interface{} `json:"UiTemplateS3Uri,omitempty"`
}

// EndpointInfo represents the EndpointInfo schema from the OpenAPI specification
type EndpointInfo struct {
	Endpointname interface{} `json:"EndpointName"`
}

// ProcessingEnvironmentMap represents the ProcessingEnvironmentMap schema from the OpenAPI specification
type ProcessingEnvironmentMap struct {
}

// CreateModelCardExportJobRequest represents the CreateModelCardExportJobRequest schema from the OpenAPI specification
type CreateModelCardExportJobRequest struct {
	Outputconfig interface{} `json:"OutputConfig"`
	Modelcardexportjobname interface{} `json:"ModelCardExportJobName"`
	Modelcardname interface{} `json:"ModelCardName"`
	Modelcardversion interface{} `json:"ModelCardVersion,omitempty"`
}

// CanvasAppSettings represents the CanvasAppSettings schema from the OpenAPI specification
type CanvasAppSettings struct {
	Timeseriesforecastingsettings interface{} `json:"TimeSeriesForecastingSettings,omitempty"`
	Workspacesettings interface{} `json:"WorkspaceSettings,omitempty"`
	Modelregistersettings interface{} `json:"ModelRegisterSettings,omitempty"`
}

// ResourceConfig represents the ResourceConfig schema from the OpenAPI specification
type ResourceConfig struct {
	Keepaliveperiodinseconds interface{} `json:"KeepAlivePeriodInSeconds,omitempty"`
	Volumekmskeyid interface{} `json:"VolumeKmsKeyId,omitempty"`
	Volumesizeingb interface{} `json:"VolumeSizeInGB"`
	Instancecount interface{} `json:"InstanceCount,omitempty"`
	Instancegroups interface{} `json:"InstanceGroups,omitempty"`
	Instancetype interface{} `json:"InstanceType,omitempty"`
}

// TimeSeriesForecastingJobConfig represents the TimeSeriesForecastingJobConfig schema from the OpenAPI specification
type TimeSeriesForecastingJobConfig struct {
	Completioncriteria AutoMLJobCompletionCriteria `json:"CompletionCriteria,omitempty"` // How long a job is allowed to run, or how many candidates a job is allowed to generate.
	Featurespecifications3uri interface{} `json:"FeatureSpecificationS3Uri,omitempty"`
	Forecastfrequency interface{} `json:"ForecastFrequency"`
	Forecasthorizon interface{} `json:"ForecastHorizon"`
	Forecastquantiles interface{} `json:"ForecastQuantiles,omitempty"`
	Timeseriesconfig interface{} `json:"TimeSeriesConfig"`
	Transformations interface{} `json:"Transformations,omitempty"`
}

// ParentHyperParameterTuningJob represents the ParentHyperParameterTuningJob schema from the OpenAPI specification
type ParentHyperParameterTuningJob struct {
	Hyperparametertuningjobname interface{} `json:"HyperParameterTuningJobName,omitempty"`
}

// RecommendationMetrics represents the RecommendationMetrics schema from the OpenAPI specification
type RecommendationMetrics struct {
	Modelsetuptime interface{} `json:"ModelSetupTime,omitempty"`
	Costperhour interface{} `json:"CostPerHour"`
	Costperinference interface{} `json:"CostPerInference"`
	Cpuutilization interface{} `json:"CpuUtilization,omitempty"`
	Maxinvocations interface{} `json:"MaxInvocations"`
	Memoryutilization interface{} `json:"MemoryUtilization,omitempty"`
	Modellatency interface{} `json:"ModelLatency"`
}

// CreateInferenceRecommendationsJobRequest represents the CreateInferenceRecommendationsJobRequest schema from the OpenAPI specification
type CreateInferenceRecommendationsJobRequest struct {
	Jobtype interface{} `json:"JobType"`
	Outputconfig interface{} `json:"OutputConfig,omitempty"`
	Rolearn interface{} `json:"RoleArn"`
	Stoppingconditions interface{} `json:"StoppingConditions,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Inputconfig interface{} `json:"InputConfig"`
	Jobdescription interface{} `json:"JobDescription,omitempty"`
	Jobname interface{} `json:"JobName"`
}

// DescribeExperimentRequest represents the DescribeExperimentRequest schema from the OpenAPI specification
type DescribeExperimentRequest struct {
	Experimentname interface{} `json:"ExperimentName"`
}

// DescribeDeviceResponse represents the DescribeDeviceResponse schema from the OpenAPI specification
type DescribeDeviceResponse struct {
	Devicefleetname interface{} `json:"DeviceFleetName"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Agentversion interface{} `json:"AgentVersion,omitempty"`
	Devicename interface{} `json:"DeviceName"`
	Devicearn interface{} `json:"DeviceArn,omitempty"`
	Iotthingname interface{} `json:"IotThingName,omitempty"`
	Maxmodels interface{} `json:"MaxModels,omitempty"`
	Models interface{} `json:"Models,omitempty"`
	Registrationtime interface{} `json:"RegistrationTime"`
	Description interface{} `json:"Description,omitempty"`
	Latestheartbeat interface{} `json:"LatestHeartbeat,omitempty"`
}

// StartNotebookInstanceInput represents the StartNotebookInstanceInput schema from the OpenAPI specification
type StartNotebookInstanceInput struct {
	Notebookinstancename interface{} `json:"NotebookInstanceName"`
}

// Parameter represents the Parameter schema from the OpenAPI specification
type Parameter struct {
	Value interface{} `json:"Value"`
	Name interface{} `json:"Name"`
}

// ListActionsRequest represents the ListActionsRequest schema from the OpenAPI specification
type ListActionsRequest struct {
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Sourceuri interface{} `json:"SourceUri,omitempty"`
	Actiontype interface{} `json:"ActionType,omitempty"`
	Createdafter interface{} `json:"CreatedAfter,omitempty"`
	Createdbefore interface{} `json:"CreatedBefore,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
}

// Edge represents the Edge schema from the OpenAPI specification
type Edge struct {
	Sourcearn interface{} `json:"SourceArn,omitempty"`
	Associationtype interface{} `json:"AssociationType,omitempty"`
	Destinationarn interface{} `json:"DestinationArn,omitempty"`
}

// RealTimeInferenceRecommendation represents the RealTimeInferenceRecommendation schema from the OpenAPI specification
type RealTimeInferenceRecommendation struct {
	Environment interface{} `json:"Environment,omitempty"`
	Instancetype interface{} `json:"InstanceType"`
	Recommendationid interface{} `json:"RecommendationId"`
}

// ModelVariantActionMap represents the ModelVariantActionMap schema from the OpenAPI specification
type ModelVariantActionMap struct {
}

// ListMonitoringAlertHistoryResponse represents the ListMonitoringAlertHistoryResponse schema from the OpenAPI specification
type ListMonitoringAlertHistoryResponse struct {
	Monitoringalerthistory interface{} `json:"MonitoringAlertHistory,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// HyperParameterTuningJobObjective represents the HyperParameterTuningJobObjective schema from the OpenAPI specification
type HyperParameterTuningJobObjective struct {
	TypeField interface{} `json:"Type"`
	Metricname interface{} `json:"MetricName"`
}

// CandidateGenerationConfig represents the CandidateGenerationConfig schema from the OpenAPI specification
type CandidateGenerationConfig struct {
	Algorithmsconfig interface{} `json:"AlgorithmsConfig,omitempty"`
}

// DescribeHumanTaskUiResponse represents the DescribeHumanTaskUiResponse schema from the OpenAPI specification
type DescribeHumanTaskUiResponse struct {
	Humantaskuistatus interface{} `json:"HumanTaskUiStatus,omitempty"`
	Uitemplate UiTemplateInfo `json:"UiTemplate"` // Container for user interface template information.
	Creationtime interface{} `json:"CreationTime"`
	Humantaskuiarn interface{} `json:"HumanTaskUiArn"`
	Humantaskuiname interface{} `json:"HumanTaskUiName"`
}

// DeploymentRecommendation represents the DeploymentRecommendation schema from the OpenAPI specification
type DeploymentRecommendation struct {
	Recommendationstatus interface{} `json:"RecommendationStatus"`
	Realtimeinferencerecommendations interface{} `json:"RealTimeInferenceRecommendations,omitempty"`
}

// InstanceMetadataServiceConfiguration represents the InstanceMetadataServiceConfiguration schema from the OpenAPI specification
type InstanceMetadataServiceConfiguration struct {
	Minimuminstancemetadataserviceversion interface{} `json:"MinimumInstanceMetadataServiceVersion"`
}

// DeleteMonitoringScheduleRequest represents the DeleteMonitoringScheduleRequest schema from the OpenAPI specification
type DeleteMonitoringScheduleRequest struct {
	Monitoringschedulename interface{} `json:"MonitoringScheduleName"`
}

// CreateAutoMLJobRequest represents the CreateAutoMLJobRequest schema from the OpenAPI specification
type CreateAutoMLJobRequest struct {
	Outputdataconfig interface{} `json:"OutputDataConfig"`
	Problemtype interface{} `json:"ProblemType,omitempty"`
	Automljobconfig interface{} `json:"AutoMLJobConfig,omitempty"`
	Modeldeployconfig interface{} `json:"ModelDeployConfig,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Generatecandidatedefinitionsonly interface{} `json:"GenerateCandidateDefinitionsOnly,omitempty"`
	Inputdataconfig interface{} `json:"InputDataConfig"`
	Rolearn interface{} `json:"RoleArn"`
	Automljobname interface{} `json:"AutoMLJobName"`
	Automljobobjective interface{} `json:"AutoMLJobObjective,omitempty"`
}

// CreateArtifactResponse represents the CreateArtifactResponse schema from the OpenAPI specification
type CreateArtifactResponse struct {
	Artifactarn interface{} `json:"ArtifactArn,omitempty"`
}

// GetDeviceFleetReportResponse represents the GetDeviceFleetReportResponse schema from the OpenAPI specification
type GetDeviceFleetReportResponse struct {
	Devicefleetname interface{} `json:"DeviceFleetName"`
	Devicestats interface{} `json:"DeviceStats,omitempty"`
	Modelstats interface{} `json:"ModelStats,omitempty"`
	Outputconfig interface{} `json:"OutputConfig,omitempty"`
	Reportgenerated interface{} `json:"ReportGenerated,omitempty"`
	Agentversions interface{} `json:"AgentVersions,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Devicefleetarn interface{} `json:"DeviceFleetArn"`
}

// UpdateExperimentResponse represents the UpdateExperimentResponse schema from the OpenAPI specification
type UpdateExperimentResponse struct {
	Experimentarn interface{} `json:"ExperimentArn,omitempty"`
}

// ListSubscribedWorkteamsResponse represents the ListSubscribedWorkteamsResponse schema from the OpenAPI specification
type ListSubscribedWorkteamsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Subscribedworkteams interface{} `json:"SubscribedWorkteams"`
}

// UpdatePipelineExecutionResponse represents the UpdatePipelineExecutionResponse schema from the OpenAPI specification
type UpdatePipelineExecutionResponse struct {
	Pipelineexecutionarn interface{} `json:"PipelineExecutionArn,omitempty"`
}

// CacheHitResult represents the CacheHitResult schema from the OpenAPI specification
type CacheHitResult struct {
	Sourcepipelineexecutionarn interface{} `json:"SourcePipelineExecutionArn,omitempty"`
}

// TrainingJobStepMetadata represents the TrainingJobStepMetadata schema from the OpenAPI specification
type TrainingJobStepMetadata struct {
	Arn interface{} `json:"Arn,omitempty"`
}

// ListWorkteamsResponse represents the ListWorkteamsResponse schema from the OpenAPI specification
type ListWorkteamsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Workteams interface{} `json:"Workteams"`
}

// CreateImageResponse represents the CreateImageResponse schema from the OpenAPI specification
type CreateImageResponse struct {
	Imagearn interface{} `json:"ImageArn,omitempty"`
}

// ListAppImageConfigsRequest represents the ListAppImageConfigsRequest schema from the OpenAPI specification
type ListAppImageConfigsRequest struct {
	Namecontains interface{} `json:"NameContains,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Modifiedtimebefore interface{} `json:"ModifiedTimeBefore,omitempty"`
	Modifiedtimeafter interface{} `json:"ModifiedTimeAfter,omitempty"`
}

// CustomizedMetricSpecification represents the CustomizedMetricSpecification schema from the OpenAPI specification
type CustomizedMetricSpecification struct {
	Metricname interface{} `json:"MetricName,omitempty"`
	Namespace interface{} `json:"Namespace,omitempty"`
	Statistic interface{} `json:"Statistic,omitempty"`
}

// DescribeCompilationJobResponse represents the DescribeCompilationJobResponse schema from the OpenAPI specification
type DescribeCompilationJobResponse struct {
	Vpcconfig interface{} `json:"VpcConfig,omitempty"`
	Compilationjobname interface{} `json:"CompilationJobName"`
	Inferenceimage interface{} `json:"InferenceImage,omitempty"`
	Modelpackageversionarn interface{} `json:"ModelPackageVersionArn,omitempty"`
	Stoppingcondition interface{} `json:"StoppingCondition"`
	Compilationendtime interface{} `json:"CompilationEndTime,omitempty"`
	Compilationstarttime interface{} `json:"CompilationStartTime,omitempty"`
	Compilationjobstatus interface{} `json:"CompilationJobStatus"`
	Creationtime interface{} `json:"CreationTime"`
	Modeldigests interface{} `json:"ModelDigests,omitempty"`
	Rolearn interface{} `json:"RoleArn"`
	Compilationjobarn interface{} `json:"CompilationJobArn"`
	Failurereason interface{} `json:"FailureReason"`
	Modelartifacts interface{} `json:"ModelArtifacts"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime"`
	Outputconfig interface{} `json:"OutputConfig"`
	Inputconfig interface{} `json:"InputConfig"`
}

// PropertyNameQuery represents the PropertyNameQuery schema from the OpenAPI specification
type PropertyNameQuery struct {
	Propertynamehint interface{} `json:"PropertyNameHint"`
}

// CreateModelPackageInput represents the CreateModelPackageInput schema from the OpenAPI specification
type CreateModelPackageInput struct {
	Inferencespecification interface{} `json:"InferenceSpecification,omitempty"`
	Modelpackagename interface{} `json:"ModelPackageName,omitempty"`
	Metadataproperties MetadataProperties `json:"MetadataProperties,omitempty"` // Metadata properties of the tracking entity, trial, or trial component.
	Validationspecification interface{} `json:"ValidationSpecification,omitempty"`
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Modelpackagedescription interface{} `json:"ModelPackageDescription,omitempty"`
	Samplepayloadurl interface{} `json:"SamplePayloadUrl,omitempty"`
	Sourcealgorithmspecification interface{} `json:"SourceAlgorithmSpecification,omitempty"`
	Certifyformarketplace interface{} `json:"CertifyForMarketplace,omitempty"`
	Task interface{} `json:"Task,omitempty"`
	Modelmetrics interface{} `json:"ModelMetrics,omitempty"`
	Modelpackagegroupname interface{} `json:"ModelPackageGroupName,omitempty"`
	Modelapprovalstatus interface{} `json:"ModelApprovalStatus,omitempty"`
	Driftcheckbaselines interface{} `json:"DriftCheckBaselines,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Additionalinferencespecifications interface{} `json:"AdditionalInferenceSpecifications,omitempty"`
	Domain interface{} `json:"Domain,omitempty"`
	Customermetadataproperties interface{} `json:"CustomerMetadataProperties,omitempty"`
}

// CreateInferenceRecommendationsJobResponse represents the CreateInferenceRecommendationsJobResponse schema from the OpenAPI specification
type CreateInferenceRecommendationsJobResponse struct {
	Jobarn interface{} `json:"JobArn"`
}

// DescribeArtifactResponse represents the DescribeArtifactResponse schema from the OpenAPI specification
type DescribeArtifactResponse struct {
	Artifactarn interface{} `json:"ArtifactArn,omitempty"`
	Artifacttype interface{} `json:"ArtifactType,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Properties interface{} `json:"Properties,omitempty"`
	Createdby UserContext `json:"CreatedBy,omitempty"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Lastmodifiedby UserContext `json:"LastModifiedBy,omitempty"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	Lineagegrouparn interface{} `json:"LineageGroupArn,omitempty"`
	Source interface{} `json:"Source,omitempty"`
	Artifactname interface{} `json:"ArtifactName,omitempty"`
	Metadataproperties MetadataProperties `json:"MetadataProperties,omitempty"` // Metadata properties of the tracking entity, trial, or trial component.
}

// HumanLoopActivationConfig represents the HumanLoopActivationConfig schema from the OpenAPI specification
type HumanLoopActivationConfig struct {
	Humanloopactivationconditionsconfig interface{} `json:"HumanLoopActivationConditionsConfig"`
}

// CreateTrialComponentRequest represents the CreateTrialComponentRequest schema from the OpenAPI specification
type CreateTrialComponentRequest struct {
	Trialcomponentname interface{} `json:"TrialComponentName"`
	Displayname interface{} `json:"DisplayName,omitempty"`
	Endtime interface{} `json:"EndTime,omitempty"`
	Parameters interface{} `json:"Parameters,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Metadataproperties MetadataProperties `json:"MetadataProperties,omitempty"` // Metadata properties of the tracking entity, trial, or trial component.
	Status interface{} `json:"Status,omitempty"`
	Inputartifacts interface{} `json:"InputArtifacts,omitempty"`
	Outputartifacts interface{} `json:"OutputArtifacts,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
}

// CreateModelBiasJobDefinitionRequest represents the CreateModelBiasJobDefinitionRequest schema from the OpenAPI specification
type CreateModelBiasJobDefinitionRequest struct {
	Networkconfig interface{} `json:"NetworkConfig,omitempty"`
	Rolearn interface{} `json:"RoleArn"`
	Stoppingcondition MonitoringStoppingCondition `json:"StoppingCondition,omitempty"` // A time limit for how long the monitoring job is allowed to run before stopping.
	Jobresources MonitoringResources `json:"JobResources"` // Identifies the resources to deploy for a monitoring job.
	Tags interface{} `json:"Tags,omitempty"`
	Modelbiasappspecification interface{} `json:"ModelBiasAppSpecification"`
	Jobdefinitionname interface{} `json:"JobDefinitionName"`
	Modelbiasbaselineconfig interface{} `json:"ModelBiasBaselineConfig,omitempty"`
	Modelbiasjoboutputconfig MonitoringOutputConfig `json:"ModelBiasJobOutputConfig"` // The output configuration for monitoring jobs.
	Modelbiasjobinput interface{} `json:"ModelBiasJobInput"`
}

// ProcessingStoppingCondition represents the ProcessingStoppingCondition schema from the OpenAPI specification
type ProcessingStoppingCondition struct {
	Maxruntimeinseconds interface{} `json:"MaxRuntimeInSeconds"`
}

// UpdateTrainingJobRequest represents the UpdateTrainingJobRequest schema from the OpenAPI specification
type UpdateTrainingJobRequest struct {
	Profilerconfig interface{} `json:"ProfilerConfig,omitempty"`
	Profilerruleconfigurations interface{} `json:"ProfilerRuleConfigurations,omitempty"`
	Resourceconfig interface{} `json:"ResourceConfig,omitempty"`
	Trainingjobname interface{} `json:"TrainingJobName"`
}

// PutModelPackageGroupPolicyInput represents the PutModelPackageGroupPolicyInput schema from the OpenAPI specification
type PutModelPackageGroupPolicyInput struct {
	Resourcepolicy interface{} `json:"ResourcePolicy"`
	Modelpackagegroupname interface{} `json:"ModelPackageGroupName"`
}

// CandidateProperties represents the CandidateProperties schema from the OpenAPI specification
type CandidateProperties struct {
	Candidatemetrics interface{} `json:"CandidateMetrics,omitempty"`
	Candidateartifactlocations interface{} `json:"CandidateArtifactLocations,omitempty"`
}

// Workforce represents the Workforce schema from the OpenAPI specification
type Workforce struct {
	Oidcconfig interface{} `json:"OidcConfig,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Cognitoconfig interface{} `json:"CognitoConfig,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Lastupdateddate interface{} `json:"LastUpdatedDate,omitempty"`
	Workforcename interface{} `json:"WorkforceName"`
	Workforcevpcconfig interface{} `json:"WorkforceVpcConfig,omitempty"`
	Subdomain interface{} `json:"SubDomain,omitempty"`
	Sourceipconfig interface{} `json:"SourceIpConfig,omitempty"`
	Workforcearn interface{} `json:"WorkforceArn"`
	Createdate interface{} `json:"CreateDate,omitempty"`
}

// ListEndpointConfigsOutput represents the ListEndpointConfigsOutput schema from the OpenAPI specification
type ListEndpointConfigsOutput struct {
	Endpointconfigs interface{} `json:"EndpointConfigs"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListHubContentVersionsResponse represents the ListHubContentVersionsResponse schema from the OpenAPI specification
type ListHubContentVersionsResponse struct {
	Hubcontentsummaries interface{} `json:"HubContentSummaries"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// RenderableTask represents the RenderableTask schema from the OpenAPI specification
type RenderableTask struct {
	Input interface{} `json:"Input"`
}

// CreateCodeRepositoryInput represents the CreateCodeRepositoryInput schema from the OpenAPI specification
type CreateCodeRepositoryInput struct {
	Coderepositoryname interface{} `json:"CodeRepositoryName"`
	Gitconfig interface{} `json:"GitConfig"`
	Tags interface{} `json:"Tags,omitempty"`
}

// CreateMonitoringScheduleResponse represents the CreateMonitoringScheduleResponse schema from the OpenAPI specification
type CreateMonitoringScheduleResponse struct {
	Monitoringschedulearn interface{} `json:"MonitoringScheduleArn"`
}

// CreateExperimentRequest represents the CreateExperimentRequest schema from the OpenAPI specification
type CreateExperimentRequest struct {
	Tags interface{} `json:"Tags,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Displayname interface{} `json:"DisplayName,omitempty"`
	Experimentname interface{} `json:"ExperimentName"`
}

// DescribeModelPackageOutput represents the DescribeModelPackageOutput schema from the OpenAPI specification
type DescribeModelPackageOutput struct {
	Modelpackagedescription interface{} `json:"ModelPackageDescription,omitempty"`
	Modelpackagegroupname interface{} `json:"ModelPackageGroupName,omitempty"`
	Additionalinferencespecifications interface{} `json:"AdditionalInferenceSpecifications,omitempty"`
	Modelpackagestatus interface{} `json:"ModelPackageStatus"`
	Domain interface{} `json:"Domain,omitempty"`
	Metadataproperties MetadataProperties `json:"MetadataProperties,omitempty"` // Metadata properties of the tracking entity, trial, or trial component.
	Modelapprovalstatus interface{} `json:"ModelApprovalStatus,omitempty"`
	Modelpackagestatusdetails interface{} `json:"ModelPackageStatusDetails"`
	Approvaldescription interface{} `json:"ApprovalDescription,omitempty"`
	Lastmodifiedby UserContext `json:"LastModifiedBy,omitempty"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	Certifyformarketplace interface{} `json:"CertifyForMarketplace,omitempty"`
	Modelmetrics interface{} `json:"ModelMetrics,omitempty"`
	Modelpackagename interface{} `json:"ModelPackageName"`
	Modelpackagearn interface{} `json:"ModelPackageArn"`
	Task interface{} `json:"Task,omitempty"`
	Createdby UserContext `json:"CreatedBy,omitempty"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	Creationtime interface{} `json:"CreationTime"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Customermetadataproperties interface{} `json:"CustomerMetadataProperties,omitempty"`
	Validationspecification interface{} `json:"ValidationSpecification,omitempty"`
	Sourcealgorithmspecification interface{} `json:"SourceAlgorithmSpecification,omitempty"`
	Driftcheckbaselines interface{} `json:"DriftCheckBaselines,omitempty"`
	Modelpackageversion interface{} `json:"ModelPackageVersion,omitempty"`
	Inferencespecification interface{} `json:"InferenceSpecification,omitempty"`
	Samplepayloadurl interface{} `json:"SamplePayloadUrl,omitempty"`
}

// StopEdgePackagingJobRequest represents the StopEdgePackagingJobRequest schema from the OpenAPI specification
type StopEdgePackagingJobRequest struct {
	Edgepackagingjobname interface{} `json:"EdgePackagingJobName"`
}

// ListActionsResponse represents the ListActionsResponse schema from the OpenAPI specification
type ListActionsResponse struct {
	Actionsummaries interface{} `json:"ActionSummaries,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DescribeFlowDefinitionRequest represents the DescribeFlowDefinitionRequest schema from the OpenAPI specification
type DescribeFlowDefinitionRequest struct {
	Flowdefinitionname interface{} `json:"FlowDefinitionName"`
}

// ProvisioningParameter represents the ProvisioningParameter schema from the OpenAPI specification
type ProvisioningParameter struct {
	Key interface{} `json:"Key,omitempty"`
	Value interface{} `json:"Value,omitempty"`
}

// UpdateEndpointWeightsAndCapacitiesOutput represents the UpdateEndpointWeightsAndCapacitiesOutput schema from the OpenAPI specification
type UpdateEndpointWeightsAndCapacitiesOutput struct {
	Endpointarn interface{} `json:"EndpointArn"`
}

// MetricSpecification represents the MetricSpecification schema from the OpenAPI specification
type MetricSpecification struct {
	Customized interface{} `json:"Customized,omitempty"`
	Predefined interface{} `json:"Predefined,omitempty"`
}

// UpdateContextResponse represents the UpdateContextResponse schema from the OpenAPI specification
type UpdateContextResponse struct {
	Contextarn interface{} `json:"ContextArn,omitempty"`
}

// AnnotationConsolidationConfig represents the AnnotationConsolidationConfig schema from the OpenAPI specification
type AnnotationConsolidationConfig struct {
	Annotationconsolidationlambdaarn interface{} `json:"AnnotationConsolidationLambdaArn"`
}

// StartInferenceExperimentRequest represents the StartInferenceExperimentRequest schema from the OpenAPI specification
type StartInferenceExperimentRequest struct {
	Name interface{} `json:"Name"`
}

// AddAssociationResponse represents the AddAssociationResponse schema from the OpenAPI specification
type AddAssociationResponse struct {
	Destinationarn interface{} `json:"DestinationArn,omitempty"`
	Sourcearn interface{} `json:"SourceArn,omitempty"`
}

// ListDataQualityJobDefinitionsRequest represents the ListDataQualityJobDefinitionsRequest schema from the OpenAPI specification
type ListDataQualityJobDefinitionsRequest struct {
	Endpointname interface{} `json:"EndpointName,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
}

// UpdateActionRequest represents the UpdateActionRequest schema from the OpenAPI specification
type UpdateActionRequest struct {
	Actionname interface{} `json:"ActionName"`
	Description interface{} `json:"Description,omitempty"`
	Properties interface{} `json:"Properties,omitempty"`
	Propertiestoremove interface{} `json:"PropertiesToRemove,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// DescribeProcessingJobResponse represents the DescribeProcessingJobResponse schema from the OpenAPI specification
type DescribeProcessingJobResponse struct {
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Exitmessage interface{} `json:"ExitMessage,omitempty"`
	Processingjobname interface{} `json:"ProcessingJobName"`
	Stoppingcondition interface{} `json:"StoppingCondition,omitempty"`
	Creationtime interface{} `json:"CreationTime"`
	Processingstarttime interface{} `json:"ProcessingStartTime,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Processingjobarn interface{} `json:"ProcessingJobArn"`
	Automljobarn interface{} `json:"AutoMLJobArn,omitempty"`
	Experimentconfig interface{} `json:"ExperimentConfig,omitempty"`
	Processingendtime interface{} `json:"ProcessingEndTime,omitempty"`
	Processinginputs interface{} `json:"ProcessingInputs,omitempty"`
	Appspecification interface{} `json:"AppSpecification"`
	Environment interface{} `json:"Environment,omitempty"`
	Networkconfig interface{} `json:"NetworkConfig,omitempty"`
	Processingjobstatus interface{} `json:"ProcessingJobStatus"`
	Processingresources interface{} `json:"ProcessingResources"`
	Trainingjobarn interface{} `json:"TrainingJobArn,omitempty"`
	Monitoringschedulearn interface{} `json:"MonitoringScheduleArn,omitempty"`
	Processingoutputconfig interface{} `json:"ProcessingOutputConfig,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
}

// UpdateNotebookInstanceOutput represents the UpdateNotebookInstanceOutput schema from the OpenAPI specification
type UpdateNotebookInstanceOutput struct {
}

// ConditionStepMetadata represents the ConditionStepMetadata schema from the OpenAPI specification
type ConditionStepMetadata struct {
	Outcome interface{} `json:"Outcome,omitempty"`
}

// TransformJobSummary represents the TransformJobSummary schema from the OpenAPI specification
type TransformJobSummary struct {
	Transformjobname interface{} `json:"TransformJobName"`
	Transformjobstatus interface{} `json:"TransformJobStatus"`
	Creationtime interface{} `json:"CreationTime"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Transformendtime interface{} `json:"TransformEndTime,omitempty"`
	Transformjobarn interface{} `json:"TransformJobArn"`
}

// ListModelsOutput represents the ListModelsOutput schema from the OpenAPI specification
type ListModelsOutput struct {
	Models interface{} `json:"Models"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CreateAppImageConfigRequest represents the CreateAppImageConfigRequest schema from the OpenAPI specification
type CreateAppImageConfigRequest struct {
	Appimageconfigname interface{} `json:"AppImageConfigName"`
	Kernelgatewayimageconfig interface{} `json:"KernelGatewayImageConfig,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// CreateDataQualityJobDefinitionResponse represents the CreateDataQualityJobDefinitionResponse schema from the OpenAPI specification
type CreateDataQualityJobDefinitionResponse struct {
	Jobdefinitionarn interface{} `json:"JobDefinitionArn"`
}

// DescribeWorkforceRequest represents the DescribeWorkforceRequest schema from the OpenAPI specification
type DescribeWorkforceRequest struct {
	Workforcename interface{} `json:"WorkforceName"`
}

// StartPipelineExecutionResponse represents the StartPipelineExecutionResponse schema from the OpenAPI specification
type StartPipelineExecutionResponse struct {
	Pipelineexecutionarn interface{} `json:"PipelineExecutionArn,omitempty"`
}

// CreateModelBiasJobDefinitionResponse represents the CreateModelBiasJobDefinitionResponse schema from the OpenAPI specification
type CreateModelBiasJobDefinitionResponse struct {
	Jobdefinitionarn interface{} `json:"JobDefinitionArn"`
}

// DescribeContextResponse represents the DescribeContextResponse schema from the OpenAPI specification
type DescribeContextResponse struct {
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Properties interface{} `json:"Properties,omitempty"`
	Source interface{} `json:"Source,omitempty"`
	Contextname interface{} `json:"ContextName,omitempty"`
	Contexttype interface{} `json:"ContextType,omitempty"`
	Lastmodifiedby UserContext `json:"LastModifiedBy,omitempty"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	Lineagegrouparn interface{} `json:"LineageGroupArn,omitempty"`
	Contextarn interface{} `json:"ContextArn,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Createdby UserContext `json:"CreatedBy,omitempty"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
}

// UpdateTrialResponse represents the UpdateTrialResponse schema from the OpenAPI specification
type UpdateTrialResponse struct {
	Trialarn interface{} `json:"TrialArn,omitempty"`
}

// CreateCompilationJobRequest represents the CreateCompilationJobRequest schema from the OpenAPI specification
type CreateCompilationJobRequest struct {
	Stoppingcondition interface{} `json:"StoppingCondition"`
	Tags interface{} `json:"Tags,omitempty"`
	Vpcconfig interface{} `json:"VpcConfig,omitempty"`
	Compilationjobname interface{} `json:"CompilationJobName"`
	Inputconfig interface{} `json:"InputConfig,omitempty"`
	Modelpackageversionarn interface{} `json:"ModelPackageVersionArn,omitempty"`
	Outputconfig interface{} `json:"OutputConfig"`
	Rolearn interface{} `json:"RoleArn"`
}

// KernelGatewayAppSettings represents the KernelGatewayAppSettings schema from the OpenAPI specification
type KernelGatewayAppSettings struct {
	Customimages interface{} `json:"CustomImages,omitempty"`
	Defaultresourcespec interface{} `json:"DefaultResourceSpec,omitempty"`
	Lifecycleconfigarns interface{} `json:"LifecycleConfigArns,omitempty"`
}

// ListArtifactsResponse represents the ListArtifactsResponse schema from the OpenAPI specification
type ListArtifactsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Artifactsummaries interface{} `json:"ArtifactSummaries,omitempty"`
}

// ModelCardVersionSummary represents the ModelCardVersionSummary schema from the OpenAPI specification
type ModelCardVersionSummary struct {
	Modelcardversion interface{} `json:"ModelCardVersion"`
	Creationtime interface{} `json:"CreationTime"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Modelcardarn interface{} `json:"ModelCardArn"`
	Modelcardname interface{} `json:"ModelCardName"`
	Modelcardstatus interface{} `json:"ModelCardStatus"`
}

// HyperParameterTrainingJobSummary represents the HyperParameterTrainingJobSummary schema from the OpenAPI specification
type HyperParameterTrainingJobSummary struct {
	Trainingendtime interface{} `json:"TrainingEndTime,omitempty"`
	Trainingjobdefinitionname interface{} `json:"TrainingJobDefinitionName,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Trainingjobarn interface{} `json:"TrainingJobArn"`
	Trainingjobname interface{} `json:"TrainingJobName"`
	Finalhyperparametertuningjobobjectivemetric interface{} `json:"FinalHyperParameterTuningJobObjectiveMetric,omitempty"`
	Trainingjobstatus interface{} `json:"TrainingJobStatus"`
	Creationtime interface{} `json:"CreationTime"`
	Objectivestatus interface{} `json:"ObjectiveStatus,omitempty"`
	Trainingstarttime interface{} `json:"TrainingStartTime,omitempty"`
	Tunedhyperparameters interface{} `json:"TunedHyperParameters"`
	Tuningjobname interface{} `json:"TuningJobName,omitempty"`
}

// CreateModelCardResponse represents the CreateModelCardResponse schema from the OpenAPI specification
type CreateModelCardResponse struct {
	Modelcardarn interface{} `json:"ModelCardArn"`
}

// Explainability represents the Explainability schema from the OpenAPI specification
type Explainability struct {
	Report interface{} `json:"Report,omitempty"`
}

// ListModelPackagesOutput represents the ListModelPackagesOutput schema from the OpenAPI specification
type ListModelPackagesOutput struct {
	Modelpackagesummarylist interface{} `json:"ModelPackageSummaryList"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// UpdateSpaceRequest represents the UpdateSpaceRequest schema from the OpenAPI specification
type UpdateSpaceRequest struct {
	Spacename interface{} `json:"SpaceName"`
	Spacesettings interface{} `json:"SpaceSettings,omitempty"`
	Domainid interface{} `json:"DomainId"`
}

// DisableSagemakerServicecatalogPortfolioOutput represents the DisableSagemakerServicecatalogPortfolioOutput schema from the OpenAPI specification
type DisableSagemakerServicecatalogPortfolioOutput struct {
}

// BatchDataCaptureConfig represents the BatchDataCaptureConfig schema from the OpenAPI specification
type BatchDataCaptureConfig struct {
	Generateinferenceid interface{} `json:"GenerateInferenceId,omitempty"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Destinations3uri interface{} `json:"DestinationS3Uri"`
}

// ConvergenceDetected represents the ConvergenceDetected schema from the OpenAPI specification
type ConvergenceDetected struct {
	Completeonconvergence interface{} `json:"CompleteOnConvergence,omitempty"`
}

// LineageEntityParameters represents the LineageEntityParameters schema from the OpenAPI specification
type LineageEntityParameters struct {
}

// ListHubContentVersionsRequest represents the ListHubContentVersionsRequest schema from the OpenAPI specification
type ListHubContentVersionsRequest struct {
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Hubname interface{} `json:"HubName"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Maxschemaversion interface{} `json:"MaxSchemaVersion,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Hubcontentname interface{} `json:"HubContentName"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Hubcontenttype interface{} `json:"HubContentType"`
	Minversion interface{} `json:"MinVersion,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
}

// UpdateModelPackageInput represents the UpdateModelPackageInput schema from the OpenAPI specification
type UpdateModelPackageInput struct {
	Modelapprovalstatus interface{} `json:"ModelApprovalStatus,omitempty"`
	Modelpackagearn interface{} `json:"ModelPackageArn"`
	Additionalinferencespecificationstoadd interface{} `json:"AdditionalInferenceSpecificationsToAdd,omitempty"`
	Approvaldescription interface{} `json:"ApprovalDescription,omitempty"`
	Customermetadataproperties interface{} `json:"CustomerMetadataProperties,omitempty"`
	Customermetadatapropertiestoremove interface{} `json:"CustomerMetadataPropertiesToRemove,omitempty"`
}

// FlowDefinitionSummary represents the FlowDefinitionSummary schema from the OpenAPI specification
type FlowDefinitionSummary struct {
	Flowdefinitionstatus interface{} `json:"FlowDefinitionStatus"`
	Creationtime interface{} `json:"CreationTime"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Flowdefinitionarn interface{} `json:"FlowDefinitionArn"`
	Flowdefinitionname interface{} `json:"FlowDefinitionName"`
}

// Bias represents the Bias schema from the OpenAPI specification
type Bias struct {
	Posttrainingreport interface{} `json:"PostTrainingReport,omitempty"`
	Pretrainingreport interface{} `json:"PreTrainingReport,omitempty"`
	Report interface{} `json:"Report,omitempty"`
}

// RecommendationJobVpcConfig represents the RecommendationJobVpcConfig schema from the OpenAPI specification
type RecommendationJobVpcConfig struct {
	Subnets interface{} `json:"Subnets"`
	Securitygroupids interface{} `json:"SecurityGroupIds"`
}

// ObjectiveStatusCounters represents the ObjectiveStatusCounters schema from the OpenAPI specification
type ObjectiveStatusCounters struct {
	Failed interface{} `json:"Failed,omitempty"`
	Pending interface{} `json:"Pending,omitempty"`
	Succeeded interface{} `json:"Succeeded,omitempty"`
}

// AdditionalInferenceSpecificationDefinition represents the AdditionalInferenceSpecificationDefinition schema from the OpenAPI specification
type AdditionalInferenceSpecificationDefinition struct {
	Supportedtransforminstancetypes interface{} `json:"SupportedTransformInstanceTypes,omitempty"`
	Containers interface{} `json:"Containers"`
	Description interface{} `json:"Description,omitempty"`
	Name interface{} `json:"Name"`
	Supportedcontenttypes interface{} `json:"SupportedContentTypes,omitempty"`
	Supportedrealtimeinferenceinstancetypes interface{} `json:"SupportedRealtimeInferenceInstanceTypes,omitempty"`
	Supportedresponsemimetypes interface{} `json:"SupportedResponseMIMETypes,omitempty"`
}

// StartPipelineExecutionRequest represents the StartPipelineExecutionRequest schema from the OpenAPI specification
type StartPipelineExecutionRequest struct {
	Pipelineexecutiondescription interface{} `json:"PipelineExecutionDescription,omitempty"`
	Pipelineexecutiondisplayname interface{} `json:"PipelineExecutionDisplayName,omitempty"`
	Pipelinename interface{} `json:"PipelineName"`
	Pipelineparameters interface{} `json:"PipelineParameters,omitempty"`
	Selectiveexecutionconfig interface{} `json:"SelectiveExecutionConfig,omitempty"`
	Clientrequesttoken interface{} `json:"ClientRequestToken"`
	Parallelismconfiguration interface{} `json:"ParallelismConfiguration,omitempty"`
}

// ListAppsResponse represents the ListAppsResponse schema from the OpenAPI specification
type ListAppsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Apps interface{} `json:"Apps,omitempty"`
}

// TrialComponentStatus represents the TrialComponentStatus schema from the OpenAPI specification
type TrialComponentStatus struct {
	Primarystatus interface{} `json:"PrimaryStatus,omitempty"`
	Message interface{} `json:"Message,omitempty"`
}

// ListAlgorithmsInput represents the ListAlgorithmsInput schema from the OpenAPI specification
type ListAlgorithmsInput struct {
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
}

// DataQualityAppSpecification represents the DataQualityAppSpecification schema from the OpenAPI specification
type DataQualityAppSpecification struct {
	Environment interface{} `json:"Environment,omitempty"`
	Imageuri interface{} `json:"ImageUri"`
	Postanalyticsprocessorsourceuri interface{} `json:"PostAnalyticsProcessorSourceUri,omitempty"`
	Recordpreprocessorsourceuri interface{} `json:"RecordPreprocessorSourceUri,omitempty"`
	Containerarguments interface{} `json:"ContainerArguments,omitempty"`
	Containerentrypoint interface{} `json:"ContainerEntrypoint,omitempty"`
}

// ImageVersion represents the ImageVersion schema from the OpenAPI specification
type ImageVersion struct {
	Imageversionarn interface{} `json:"ImageVersionArn"`
	Imageversionstatus interface{} `json:"ImageVersionStatus"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime"`
	Version interface{} `json:"Version"`
	Creationtime interface{} `json:"CreationTime"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Imagearn interface{} `json:"ImageArn"`
}

// HyperParameterTuningResourceConfig represents the HyperParameterTuningResourceConfig schema from the OpenAPI specification
type HyperParameterTuningResourceConfig struct {
	Instancecount interface{} `json:"InstanceCount,omitempty"`
	Instancetype interface{} `json:"InstanceType,omitempty"`
	Volumekmskeyid interface{} `json:"VolumeKmsKeyId,omitempty"`
	Volumesizeingb interface{} `json:"VolumeSizeInGB,omitempty"`
	Allocationstrategy interface{} `json:"AllocationStrategy,omitempty"`
	Instanceconfigs interface{} `json:"InstanceConfigs,omitempty"`
}

// TimeSeriesForecastingSettings represents the TimeSeriesForecastingSettings schema from the OpenAPI specification
type TimeSeriesForecastingSettings struct {
	Amazonforecastrolearn interface{} `json:"AmazonForecastRoleArn,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// ModelPackageSummaries represents the ModelPackageSummaries schema from the OpenAPI specification
type ModelPackageSummaries struct {
}

// CustomerMetadataMap represents the CustomerMetadataMap schema from the OpenAPI specification
type CustomerMetadataMap struct {
}

// PipelineExecution represents the PipelineExecution schema from the OpenAPI specification
type PipelineExecution struct {
	Pipelineexecutionstatus interface{} `json:"PipelineExecutionStatus,omitempty"`
	Selectiveexecutionconfig interface{} `json:"SelectiveExecutionConfig,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Pipelineexecutionarn interface{} `json:"PipelineExecutionArn,omitempty"`
	Createdby UserContext `json:"CreatedBy,omitempty"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	Parallelismconfiguration interface{} `json:"ParallelismConfiguration,omitempty"`
	Pipelineexecutiondisplayname interface{} `json:"PipelineExecutionDisplayName,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Lastmodifiedby UserContext `json:"LastModifiedBy,omitempty"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	Pipelinearn interface{} `json:"PipelineArn,omitempty"`
	Pipelineexperimentconfig PipelineExperimentConfig `json:"PipelineExperimentConfig,omitempty"` // Specifies the names of the experiment and trial created by a pipeline.
	Pipelineparameters interface{} `json:"PipelineParameters,omitempty"`
	Pipelineexecutiondescription interface{} `json:"PipelineExecutionDescription,omitempty"`
}

// TrialComponentSummary represents the TrialComponentSummary schema from the OpenAPI specification
type TrialComponentSummary struct {
	Endtime interface{} `json:"EndTime,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
	Trialcomponentarn interface{} `json:"TrialComponentArn,omitempty"`
	Trialcomponentname interface{} `json:"TrialComponentName,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Lastmodifiedby interface{} `json:"LastModifiedBy,omitempty"`
	Trialcomponentsource TrialComponentSource `json:"TrialComponentSource,omitempty"` // The Amazon Resource Name (ARN) and job type of the source of a trial component.
	Createdby interface{} `json:"CreatedBy,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Displayname interface{} `json:"DisplayName,omitempty"`
}

// UpdateHubRequest represents the UpdateHubRequest schema from the OpenAPI specification
type UpdateHubRequest struct {
	Hubsearchkeywords interface{} `json:"HubSearchKeywords,omitempty"`
	Hubdescription interface{} `json:"HubDescription,omitempty"`
	Hubdisplayname interface{} `json:"HubDisplayName,omitempty"`
	Hubname interface{} `json:"HubName"`
}

// ModelArtifacts represents the ModelArtifacts schema from the OpenAPI specification
type ModelArtifacts struct {
	S3modelartifacts interface{} `json:"S3ModelArtifacts"`
}

// DataCaptureConfigSummary represents the DataCaptureConfigSummary schema from the OpenAPI specification
type DataCaptureConfigSummary struct {
	Currentsamplingpercentage interface{} `json:"CurrentSamplingPercentage"`
	Destinations3uri interface{} `json:"DestinationS3Uri"`
	Enablecapture interface{} `json:"EnableCapture"`
	Kmskeyid interface{} `json:"KmsKeyId"`
	Capturestatus interface{} `json:"CaptureStatus"`
}

// ProductionVariantStatus represents the ProductionVariantStatus schema from the OpenAPI specification
type ProductionVariantStatus struct {
	Statusmessage interface{} `json:"StatusMessage,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
	Status interface{} `json:"Status"`
}

// DeleteWorkteamResponse represents the DeleteWorkteamResponse schema from the OpenAPI specification
type DeleteWorkteamResponse struct {
	Success interface{} `json:"Success"`
}

// EdgeModelSummary represents the EdgeModelSummary schema from the OpenAPI specification
type EdgeModelSummary struct {
	Modelname interface{} `json:"ModelName"`
	Modelversion interface{} `json:"ModelVersion"`
}

// AutoMLDataSplitConfig represents the AutoMLDataSplitConfig schema from the OpenAPI specification
type AutoMLDataSplitConfig struct {
	Validationfraction interface{} `json:"ValidationFraction,omitempty"`
}

// UpdateMonitoringAlertRequest represents the UpdateMonitoringAlertRequest schema from the OpenAPI specification
type UpdateMonitoringAlertRequest struct {
	Monitoringschedulename interface{} `json:"MonitoringScheduleName"`
	Datapointstoalert interface{} `json:"DatapointsToAlert"`
	Evaluationperiod interface{} `json:"EvaluationPeriod"`
	Monitoringalertname interface{} `json:"MonitoringAlertName"`
}

// StudioLifecycleConfigDetails represents the StudioLifecycleConfigDetails schema from the OpenAPI specification
type StudioLifecycleConfigDetails struct {
	Studiolifecycleconfigarn interface{} `json:"StudioLifecycleConfigArn,omitempty"`
	Studiolifecycleconfigname interface{} `json:"StudioLifecycleConfigName,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Studiolifecycleconfigapptype interface{} `json:"StudioLifecycleConfigAppType,omitempty"`
}

// ListModelPackagesInput represents the ListModelPackagesInput schema from the OpenAPI specification
type ListModelPackagesInput struct {
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Modelapprovalstatus interface{} `json:"ModelApprovalStatus,omitempty"`
	Modelpackagetype interface{} `json:"ModelPackageType,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Modelpackagegroupname interface{} `json:"ModelPackageGroupName,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
}

// FillingTransformations represents the FillingTransformations schema from the OpenAPI specification
type FillingTransformations struct {
}

// TrialSource represents the TrialSource schema from the OpenAPI specification
type TrialSource struct {
	Sourcetype interface{} `json:"SourceType,omitempty"`
	Sourcearn interface{} `json:"SourceArn"`
}

// DeleteStudioLifecycleConfigRequest represents the DeleteStudioLifecycleConfigRequest schema from the OpenAPI specification
type DeleteStudioLifecycleConfigRequest struct {
	Studiolifecycleconfigname interface{} `json:"StudioLifecycleConfigName"`
}

// ListStageDevicesResponse represents the ListStageDevicesResponse schema from the OpenAPI specification
type ListStageDevicesResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Devicedeploymentsummaries interface{} `json:"DeviceDeploymentSummaries"`
}

// LabelCountersForWorkteam represents the LabelCountersForWorkteam schema from the OpenAPI specification
type LabelCountersForWorkteam struct {
	Pendinghuman interface{} `json:"PendingHuman,omitempty"`
	Total interface{} `json:"Total,omitempty"`
	Humanlabeled interface{} `json:"HumanLabeled,omitempty"`
}

// CreateImageVersionResponse represents the CreateImageVersionResponse schema from the OpenAPI specification
type CreateImageVersionResponse struct {
	Imageversionarn interface{} `json:"ImageVersionArn,omitempty"`
}

// CreateHyperParameterTuningJobResponse represents the CreateHyperParameterTuningJobResponse schema from the OpenAPI specification
type CreateHyperParameterTuningJobResponse struct {
	Hyperparametertuningjobarn interface{} `json:"HyperParameterTuningJobArn"`
}

// MonitoringJobDefinition represents the MonitoringJobDefinition schema from the OpenAPI specification
type MonitoringJobDefinition struct {
	Baselineconfig interface{} `json:"BaselineConfig,omitempty"`
	Environment interface{} `json:"Environment,omitempty"`
	Monitoringappspecification interface{} `json:"MonitoringAppSpecification"`
	Monitoringresources interface{} `json:"MonitoringResources"`
	Networkconfig interface{} `json:"NetworkConfig,omitempty"`
	Rolearn interface{} `json:"RoleArn"`
	Stoppingcondition interface{} `json:"StoppingCondition,omitempty"`
	Monitoringinputs interface{} `json:"MonitoringInputs"`
	Monitoringoutputconfig interface{} `json:"MonitoringOutputConfig"`
}

// CreateWorkteamRequest represents the CreateWorkteamRequest schema from the OpenAPI specification
type CreateWorkteamRequest struct {
	Description interface{} `json:"Description"`
	Memberdefinitions interface{} `json:"MemberDefinitions"`
	Notificationconfiguration interface{} `json:"NotificationConfiguration,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Workforcename interface{} `json:"WorkforceName,omitempty"`
	Workteamname interface{} `json:"WorkteamName"`
}

// DeleteArtifactResponse represents the DeleteArtifactResponse schema from the OpenAPI specification
type DeleteArtifactResponse struct {
	Artifactarn interface{} `json:"ArtifactArn,omitempty"`
}

// TransformResources represents the TransformResources schema from the OpenAPI specification
type TransformResources struct {
	Volumekmskeyid interface{} `json:"VolumeKmsKeyId,omitempty"`
	Instancecount interface{} `json:"InstanceCount"`
	Instancetype interface{} `json:"InstanceType"`
}

// CodeRepository represents the CodeRepository schema from the OpenAPI specification
type CodeRepository struct {
	Repositoryurl interface{} `json:"RepositoryUrl"`
}

// ListModelCardExportJobsResponse represents the ListModelCardExportJobsResponse schema from the OpenAPI specification
type ListModelCardExportJobsResponse struct {
	Modelcardexportjobsummaries interface{} `json:"ModelCardExportJobSummaries"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// StopHyperParameterTuningJobRequest represents the StopHyperParameterTuningJobRequest schema from the OpenAPI specification
type StopHyperParameterTuningJobRequest struct {
	Hyperparametertuningjobname interface{} `json:"HyperParameterTuningJobName"`
}

// BatchDescribeModelPackageSummary represents the BatchDescribeModelPackageSummary schema from the OpenAPI specification
type BatchDescribeModelPackageSummary struct {
	Modelpackagearn interface{} `json:"ModelPackageArn"`
	Modelpackagedescription interface{} `json:"ModelPackageDescription,omitempty"`
	Modelpackagegroupname interface{} `json:"ModelPackageGroupName"`
	Modelpackagestatus interface{} `json:"ModelPackageStatus"`
	Modelpackageversion interface{} `json:"ModelPackageVersion,omitempty"`
	Creationtime interface{} `json:"CreationTime"`
	Inferencespecification InferenceSpecification `json:"InferenceSpecification"` // Defines how to perform inference generation after a training job is run.
	Modelapprovalstatus interface{} `json:"ModelApprovalStatus,omitempty"`
}

// DescribeInferenceExperimentRequest represents the DescribeInferenceExperimentRequest schema from the OpenAPI specification
type DescribeInferenceExperimentRequest struct {
	Name interface{} `json:"Name"`
}

// DescribeModelCardExportJobResponse represents the DescribeModelCardExportJobResponse schema from the OpenAPI specification
type DescribeModelCardExportJobResponse struct {
	Modelcardexportjobname interface{} `json:"ModelCardExportJobName"`
	Modelcardname interface{} `json:"ModelCardName"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Lastmodifiedat interface{} `json:"LastModifiedAt"`
	Modelcardexportjobarn interface{} `json:"ModelCardExportJobArn"`
	Status interface{} `json:"Status"`
	Modelcardversion interface{} `json:"ModelCardVersion"`
	Outputconfig interface{} `json:"OutputConfig"`
	Createdat interface{} `json:"CreatedAt"`
	Exportartifacts interface{} `json:"ExportArtifacts,omitempty"`
}

// TrafficPattern represents the TrafficPattern schema from the OpenAPI specification
type TrafficPattern struct {
	Phases interface{} `json:"Phases,omitempty"`
	Stairs interface{} `json:"Stairs,omitempty"`
	Traffictype interface{} `json:"TrafficType,omitempty"`
}

// CreateArtifactRequest represents the CreateArtifactRequest schema from the OpenAPI specification
type CreateArtifactRequest struct {
	Tags interface{} `json:"Tags,omitempty"`
	Artifactname interface{} `json:"ArtifactName,omitempty"`
	Artifacttype interface{} `json:"ArtifactType"`
	Metadataproperties MetadataProperties `json:"MetadataProperties,omitempty"` // Metadata properties of the tracking entity, trial, or trial component.
	Properties interface{} `json:"Properties,omitempty"`
	Source interface{} `json:"Source"`
}

// CreateDeviceFleetRequest represents the CreateDeviceFleetRequest schema from the OpenAPI specification
type CreateDeviceFleetRequest struct {
	Enableiotrolealias interface{} `json:"EnableIotRoleAlias,omitempty"`
	Outputconfig interface{} `json:"OutputConfig"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Devicefleetname interface{} `json:"DeviceFleetName"`
}

// DeleteDomainRequest represents the DeleteDomainRequest schema from the OpenAPI specification
type DeleteDomainRequest struct {
	Retentionpolicy interface{} `json:"RetentionPolicy,omitempty"`
	Domainid interface{} `json:"DomainId"`
}

// MonitoringScheduleConfig represents the MonitoringScheduleConfig schema from the OpenAPI specification
type MonitoringScheduleConfig struct {
	Monitoringjobdefinition interface{} `json:"MonitoringJobDefinition,omitempty"`
	Monitoringjobdefinitionname interface{} `json:"MonitoringJobDefinitionName,omitempty"`
	Monitoringtype interface{} `json:"MonitoringType,omitempty"`
	Scheduleconfig interface{} `json:"ScheduleConfig,omitempty"`
}

// MonitoringEnvironmentMap represents the MonitoringEnvironmentMap schema from the OpenAPI specification
type MonitoringEnvironmentMap struct {
}

// DeleteCodeRepositoryInput represents the DeleteCodeRepositoryInput schema from the OpenAPI specification
type DeleteCodeRepositoryInput struct {
	Coderepositoryname interface{} `json:"CodeRepositoryName"`
}

// TargetPlatform represents the TargetPlatform schema from the OpenAPI specification
type TargetPlatform struct {
	Os interface{} `json:"Os"`
	Accelerator interface{} `json:"Accelerator,omitempty"`
	Arch interface{} `json:"Arch"`
}

// CreateModelPackageGroupInput represents the CreateModelPackageGroupInput schema from the OpenAPI specification
type CreateModelPackageGroupInput struct {
	Modelpackagegroupdescription interface{} `json:"ModelPackageGroupDescription,omitempty"`
	Modelpackagegroupname interface{} `json:"ModelPackageGroupName"`
	Tags interface{} `json:"Tags,omitempty"`
}

// CreatePresignedDomainUrlRequest represents the CreatePresignedDomainUrlRequest schema from the OpenAPI specification
type CreatePresignedDomainUrlRequest struct {
	Spacename interface{} `json:"SpaceName,omitempty"`
	Userprofilename interface{} `json:"UserProfileName"`
	Domainid interface{} `json:"DomainId"`
	Expiresinseconds interface{} `json:"ExpiresInSeconds,omitempty"`
	Sessionexpirationdurationinseconds interface{} `json:"SessionExpirationDurationInSeconds,omitempty"`
}

// ListDeviceFleetsRequest represents the ListDeviceFleetsRequest schema from the OpenAPI specification
type ListDeviceFleetsRequest struct {
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Lastmodifiedtimebefore interface{} `json:"LastModifiedTimeBefore,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
	Lastmodifiedtimeafter interface{} `json:"LastModifiedTimeAfter,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
}

// CognitoMemberDefinition represents the CognitoMemberDefinition schema from the OpenAPI specification
type CognitoMemberDefinition struct {
	Userpool interface{} `json:"UserPool"`
	Clientid interface{} `json:"ClientId"`
	Usergroup interface{} `json:"UserGroup"`
}

// CreateDataQualityJobDefinitionRequest represents the CreateDataQualityJobDefinitionRequest schema from the OpenAPI specification
type CreateDataQualityJobDefinitionRequest struct {
	Dataqualityappspecification interface{} `json:"DataQualityAppSpecification"`
	Dataqualitybaselineconfig interface{} `json:"DataQualityBaselineConfig,omitempty"`
	Dataqualityjoboutputconfig MonitoringOutputConfig `json:"DataQualityJobOutputConfig"` // The output configuration for monitoring jobs.
	Jobdefinitionname interface{} `json:"JobDefinitionName"`
	Networkconfig interface{} `json:"NetworkConfig,omitempty"`
	Dataqualityjobinput interface{} `json:"DataQualityJobInput"`
	Rolearn interface{} `json:"RoleArn"`
	Stoppingcondition MonitoringStoppingCondition `json:"StoppingCondition,omitempty"` // A time limit for how long the monitoring job is allowed to run before stopping.
	Tags interface{} `json:"Tags,omitempty"`
	Jobresources MonitoringResources `json:"JobResources"` // Identifies the resources to deploy for a monitoring job.
}

// CapacitySize represents the CapacitySize schema from the OpenAPI specification
type CapacitySize struct {
	TypeField interface{} `json:"Type"`
	Value interface{} `json:"Value"`
}

// ResolvedAttributes represents the ResolvedAttributes schema from the OpenAPI specification
type ResolvedAttributes struct {
	Automljobobjective AutoMLJobObjective `json:"AutoMLJobObjective,omitempty"` // Specifies a metric to minimize or maximize as the objective of a job.
	Completioncriteria AutoMLJobCompletionCriteria `json:"CompletionCriteria,omitempty"` // How long a job is allowed to run, or how many candidates a job is allowed to generate.
	Problemtype interface{} `json:"ProblemType,omitempty"`
}

// DescribeModelPackageInput represents the DescribeModelPackageInput schema from the OpenAPI specification
type DescribeModelPackageInput struct {
	Modelpackagename interface{} `json:"ModelPackageName"`
}

// CreateFlowDefinitionRequest represents the CreateFlowDefinitionRequest schema from the OpenAPI specification
type CreateFlowDefinitionRequest struct {
	Humanlooprequestsource interface{} `json:"HumanLoopRequestSource,omitempty"`
	Outputconfig interface{} `json:"OutputConfig"`
	Rolearn interface{} `json:"RoleArn"`
	Tags interface{} `json:"Tags,omitempty"`
	Flowdefinitionname interface{} `json:"FlowDefinitionName"`
	Humanloopactivationconfig interface{} `json:"HumanLoopActivationConfig,omitempty"`
	Humanloopconfig interface{} `json:"HumanLoopConfig"`
}

// CreatePresignedDomainUrlResponse represents the CreatePresignedDomainUrlResponse schema from the OpenAPI specification
type CreatePresignedDomainUrlResponse struct {
	Authorizedurl interface{} `json:"AuthorizedUrl,omitempty"`
}

// InferenceExperimentSummary represents the InferenceExperimentSummary schema from the OpenAPI specification
type InferenceExperimentSummary struct {
	Description interface{} `json:"Description,omitempty"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	TypeField interface{} `json:"Type"`
	Schedule interface{} `json:"Schedule,omitempty"`
	Status interface{} `json:"Status"`
	Creationtime interface{} `json:"CreationTime"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime"`
	Completiontime interface{} `json:"CompletionTime,omitempty"`
	Name interface{} `json:"Name"`
	Statusreason interface{} `json:"StatusReason,omitempty"`
}

// DescribePipelineResponse represents the DescribePipelineResponse schema from the OpenAPI specification
type DescribePipelineResponse struct {
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Pipelinedefinition interface{} `json:"PipelineDefinition,omitempty"`
	Createdby UserContext `json:"CreatedBy,omitempty"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	Parallelismconfiguration interface{} `json:"ParallelismConfiguration,omitempty"`
	Pipelinename interface{} `json:"PipelineName,omitempty"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Pipelinedescription interface{} `json:"PipelineDescription,omitempty"`
	Pipelinestatus interface{} `json:"PipelineStatus,omitempty"`
	Lastruntime interface{} `json:"LastRunTime,omitempty"`
	Pipelinearn interface{} `json:"PipelineArn,omitempty"`
	Pipelinedisplayname interface{} `json:"PipelineDisplayName,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Lastmodifiedby UserContext `json:"LastModifiedBy,omitempty"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
}

// LabelingJobResourceConfig represents the LabelingJobResourceConfig schema from the OpenAPI specification
type LabelingJobResourceConfig struct {
	Vpcconfig VpcConfig `json:"VpcConfig,omitempty"` // Specifies a VPC that your training jobs and hosted models have access to. Control access to and from your training and model containers by configuring the VPC. For more information, see <a href="https://docs.aws.amazon.com/sagemaker/latest/dg/host-vpc.html">Protect Endpoints by Using an Amazon Virtual Private Cloud</a> and <a href="https://docs.aws.amazon.com/sagemaker/latest/dg/train-vpc.html">Protect Training Jobs by Using an Amazon Virtual Private Cloud</a>.
	Volumekmskeyid interface{} `json:"VolumeKmsKeyId,omitempty"`
}

// ListProjectsOutput represents the ListProjectsOutput schema from the OpenAPI specification
type ListProjectsOutput struct {
	Projectsummarylist interface{} `json:"ProjectSummaryList"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// UpdatePipelineResponse represents the UpdatePipelineResponse schema from the OpenAPI specification
type UpdatePipelineResponse struct {
	Pipelinearn interface{} `json:"PipelineArn,omitempty"`
}

// CreateLabelingJobResponse represents the CreateLabelingJobResponse schema from the OpenAPI specification
type CreateLabelingJobResponse struct {
	Labelingjobarn interface{} `json:"LabelingJobArn"`
}

// LabelCounters represents the LabelCounters schema from the OpenAPI specification
type LabelCounters struct {
	Machinelabeled interface{} `json:"MachineLabeled,omitempty"`
	Totallabeled interface{} `json:"TotalLabeled,omitempty"`
	Unlabeled interface{} `json:"Unlabeled,omitempty"`
	Failednonretryableerror interface{} `json:"FailedNonRetryableError,omitempty"`
	Humanlabeled interface{} `json:"HumanLabeled,omitempty"`
}

// ClarifyTextConfig represents the ClarifyTextConfig schema from the OpenAPI specification
type ClarifyTextConfig struct {
	Granularity interface{} `json:"Granularity"`
	Language interface{} `json:"Language"`
}

// GitConfig represents the GitConfig schema from the OpenAPI specification
type GitConfig struct {
	Branch interface{} `json:"Branch,omitempty"`
	Repositoryurl interface{} `json:"RepositoryUrl"`
	Secretarn interface{} `json:"SecretArn,omitempty"`
}

// DescribeTrialRequest represents the DescribeTrialRequest schema from the OpenAPI specification
type DescribeTrialRequest struct {
	Trialname interface{} `json:"TrialName"`
}

// HubContentDependency represents the HubContentDependency schema from the OpenAPI specification
type HubContentDependency struct {
	Dependencycopypath interface{} `json:"DependencyCopyPath,omitempty"`
	Dependencyoriginpath interface{} `json:"DependencyOriginPath,omitempty"`
}

// QualityCheckStepMetadata represents the QualityCheckStepMetadata schema from the OpenAPI specification
type QualityCheckStepMetadata struct {
	Registernewbaseline interface{} `json:"RegisterNewBaseline,omitempty"`
	Skipcheck interface{} `json:"SkipCheck,omitempty"`
	Checkjobarn interface{} `json:"CheckJobArn,omitempty"`
	Calculatedbaselineconstraints interface{} `json:"CalculatedBaselineConstraints,omitempty"`
	Modelpackagegroupname interface{} `json:"ModelPackageGroupName,omitempty"`
	Baselineusedfordriftcheckconstraints interface{} `json:"BaselineUsedForDriftCheckConstraints,omitempty"`
	Calculatedbaselinestatistics interface{} `json:"CalculatedBaselineStatistics,omitempty"`
	Violationreport interface{} `json:"ViolationReport,omitempty"`
	Baselineusedfordriftcheckstatistics interface{} `json:"BaselineUsedForDriftCheckStatistics,omitempty"`
	Checktype interface{} `json:"CheckType,omitempty"`
}

// DeviceStats represents the DeviceStats schema from the OpenAPI specification
type DeviceStats struct {
	Connecteddevicecount interface{} `json:"ConnectedDeviceCount"`
	Registereddevicecount interface{} `json:"RegisteredDeviceCount"`
}

// AddTagsInput represents the AddTagsInput schema from the OpenAPI specification
type AddTagsInput struct {
	Resourcearn interface{} `json:"ResourceArn"`
	Tags interface{} `json:"Tags"`
}

// AutoMLProblemTypeConfig represents the AutoMLProblemTypeConfig schema from the OpenAPI specification
type AutoMLProblemTypeConfig struct {
	Tabularjobconfig interface{} `json:"TabularJobConfig,omitempty"`
	Textclassificationjobconfig interface{} `json:"TextClassificationJobConfig,omitempty"`
	Timeseriesforecastingjobconfig interface{} `json:"TimeSeriesForecastingJobConfig,omitempty"`
	Imageclassificationjobconfig interface{} `json:"ImageClassificationJobConfig,omitempty"`
}

// ModelDashboardIndicatorAction represents the ModelDashboardIndicatorAction schema from the OpenAPI specification
type ModelDashboardIndicatorAction struct {
	Enabled interface{} `json:"Enabled,omitempty"`
}

// ListAssociationsResponse represents the ListAssociationsResponse schema from the OpenAPI specification
type ListAssociationsResponse struct {
	Associationsummaries interface{} `json:"AssociationSummaries,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DeleteFeatureGroupRequest represents the DeleteFeatureGroupRequest schema from the OpenAPI specification
type DeleteFeatureGroupRequest struct {
	Featuregroupname interface{} `json:"FeatureGroupName"`
}

// DeleteContextRequest represents the DeleteContextRequest schema from the OpenAPI specification
type DeleteContextRequest struct {
	Contextname interface{} `json:"ContextName"`
}

// SearchRequest represents the SearchRequest schema from the OpenAPI specification
type SearchRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Resource interface{} `json:"Resource"`
	Searchexpression interface{} `json:"SearchExpression,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Crossaccountfilteroption interface{} `json:"CrossAccountFilterOption,omitempty"`
}

// DeleteNotebookInstanceInput represents the DeleteNotebookInstanceInput schema from the OpenAPI specification
type DeleteNotebookInstanceInput struct {
	Notebookinstancename interface{} `json:"NotebookInstanceName"`
}

// EdgeDeploymentConfig represents the EdgeDeploymentConfig schema from the OpenAPI specification
type EdgeDeploymentConfig struct {
	Failurehandlingpolicy interface{} `json:"FailureHandlingPolicy"`
}

// ListTrainingJobsForHyperParameterTuningJobRequest represents the ListTrainingJobsForHyperParameterTuningJobRequest schema from the OpenAPI specification
type ListTrainingJobsForHyperParameterTuningJobRequest struct {
	Hyperparametertuningjobname interface{} `json:"HyperParameterTuningJobName"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Statusequals interface{} `json:"StatusEquals,omitempty"`
}

// DeleteUserProfileRequest represents the DeleteUserProfileRequest schema from the OpenAPI specification
type DeleteUserProfileRequest struct {
	Domainid interface{} `json:"DomainId"`
	Userprofilename interface{} `json:"UserProfileName"`
}

// ListNotebookInstanceLifecycleConfigsOutput represents the ListNotebookInstanceLifecycleConfigsOutput schema from the OpenAPI specification
type ListNotebookInstanceLifecycleConfigsOutput struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Notebookinstancelifecycleconfigs interface{} `json:"NotebookInstanceLifecycleConfigs,omitempty"`
}

// PipelineExperimentConfig represents the PipelineExperimentConfig schema from the OpenAPI specification
type PipelineExperimentConfig struct {
	Trialname interface{} `json:"TrialName,omitempty"`
	Experimentname interface{} `json:"ExperimentName,omitempty"`
}

// DeleteWorkforceResponse represents the DeleteWorkforceResponse schema from the OpenAPI specification
type DeleteWorkforceResponse struct {
}

// RecommendationJobPayloadConfig represents the RecommendationJobPayloadConfig schema from the OpenAPI specification
type RecommendationJobPayloadConfig struct {
	Samplepayloadurl interface{} `json:"SamplePayloadUrl,omitempty"`
	Supportedcontenttypes interface{} `json:"SupportedContentTypes,omitempty"`
}

// UpdateEndpointOutput represents the UpdateEndpointOutput schema from the OpenAPI specification
type UpdateEndpointOutput struct {
	Endpointarn interface{} `json:"EndpointArn"`
}

// AutoMLJobSummary represents the AutoMLJobSummary schema from the OpenAPI specification
type AutoMLJobSummary struct {
	Automljobsecondarystatus interface{} `json:"AutoMLJobSecondaryStatus"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Automljobarn interface{} `json:"AutoMLJobArn"`
	Partialfailurereasons interface{} `json:"PartialFailureReasons,omitempty"`
	Automljobname interface{} `json:"AutoMLJobName"`
	Automljobstatus interface{} `json:"AutoMLJobStatus"`
	Creationtime interface{} `json:"CreationTime"`
	Endtime interface{} `json:"EndTime,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime"`
}

// DeleteTrialResponse represents the DeleteTrialResponse schema from the OpenAPI specification
type DeleteTrialResponse struct {
	Trialarn interface{} `json:"TrialArn,omitempty"`
}

// FeatureDefinition represents the FeatureDefinition schema from the OpenAPI specification
type FeatureDefinition struct {
	Featurename interface{} `json:"FeatureName,omitempty"`
	Featuretype interface{} `json:"FeatureType,omitempty"`
}

// TrainingEnvironmentMap represents the TrainingEnvironmentMap schema from the OpenAPI specification
type TrainingEnvironmentMap struct {
}

// ActionSource represents the ActionSource schema from the OpenAPI specification
type ActionSource struct {
	Sourceid interface{} `json:"SourceId,omitempty"`
	Sourcetype interface{} `json:"SourceType,omitempty"`
	Sourceuri interface{} `json:"SourceUri"`
}

// ListNotebookInstanceLifecycleConfigsInput represents the ListNotebookInstanceLifecycleConfigsInput schema from the OpenAPI specification
type ListNotebookInstanceLifecycleConfigsInput struct {
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
	Lastmodifiedtimeafter interface{} `json:"LastModifiedTimeAfter,omitempty"`
	Lastmodifiedtimebefore interface{} `json:"LastModifiedTimeBefore,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// StopCompilationJobRequest represents the StopCompilationJobRequest schema from the OpenAPI specification
type StopCompilationJobRequest struct {
	Compilationjobname interface{} `json:"CompilationJobName"`
}

// UpdateFeatureMetadataRequest represents the UpdateFeatureMetadataRequest schema from the OpenAPI specification
type UpdateFeatureMetadataRequest struct {
	Parameterremovals interface{} `json:"ParameterRemovals,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Featuregroupname interface{} `json:"FeatureGroupName"`
	Featurename interface{} `json:"FeatureName"`
	Parameteradditions interface{} `json:"ParameterAdditions,omitempty"`
}

// ListHumanTaskUisResponse represents the ListHumanTaskUisResponse schema from the OpenAPI specification
type ListHumanTaskUisResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Humantaskuisummaries interface{} `json:"HumanTaskUiSummaries"`
}

// NetworkConfig represents the NetworkConfig schema from the OpenAPI specification
type NetworkConfig struct {
	Enableintercontainertrafficencryption interface{} `json:"EnableInterContainerTrafficEncryption,omitempty"`
	Enablenetworkisolation interface{} `json:"EnableNetworkIsolation,omitempty"`
	Vpcconfig VpcConfig `json:"VpcConfig,omitempty"` // Specifies a VPC that your training jobs and hosted models have access to. Control access to and from your training and model containers by configuring the VPC. For more information, see <a href="https://docs.aws.amazon.com/sagemaker/latest/dg/host-vpc.html">Protect Endpoints by Using an Amazon Virtual Private Cloud</a> and <a href="https://docs.aws.amazon.com/sagemaker/latest/dg/train-vpc.html">Protect Training Jobs by Using an Amazon Virtual Private Cloud</a>.
}

// DatasetDefinition represents the DatasetDefinition schema from the OpenAPI specification
type DatasetDefinition struct {
	Localpath interface{} `json:"LocalPath,omitempty"`
	Redshiftdatasetdefinition RedshiftDatasetDefinition `json:"RedshiftDatasetDefinition,omitempty"` // Configuration for Redshift Dataset Definition input.
	Athenadatasetdefinition AthenaDatasetDefinition `json:"AthenaDatasetDefinition,omitempty"` // Configuration for Athena Dataset Definition input.
	Datadistributiontype interface{} `json:"DataDistributionType,omitempty"`
	Inputmode interface{} `json:"InputMode,omitempty"`
}

// UpdateImageVersionResponse represents the UpdateImageVersionResponse schema from the OpenAPI specification
type UpdateImageVersionResponse struct {
	Imageversionarn interface{} `json:"ImageVersionArn,omitempty"`
}

// InstanceGroup represents the InstanceGroup schema from the OpenAPI specification
type InstanceGroup struct {
	Instancecount interface{} `json:"InstanceCount"`
	Instancegroupname interface{} `json:"InstanceGroupName"`
	Instancetype interface{} `json:"InstanceType"`
}

// CreateTransformJobResponse represents the CreateTransformJobResponse schema from the OpenAPI specification
type CreateTransformJobResponse struct {
	Transformjobarn interface{} `json:"TransformJobArn"`
}

// DeploymentStage represents the DeploymentStage schema from the OpenAPI specification
type DeploymentStage struct {
	Deviceselectionconfig interface{} `json:"DeviceSelectionConfig"`
	Stagename interface{} `json:"StageName"`
	Deploymentconfig interface{} `json:"DeploymentConfig,omitempty"`
}

// BestObjectiveNotImproving represents the BestObjectiveNotImproving schema from the OpenAPI specification
type BestObjectiveNotImproving struct {
	Maxnumberoftrainingjobsnotimproving interface{} `json:"MaxNumberOfTrainingJobsNotImproving,omitempty"`
}

// ModelPackageStatusDetails represents the ModelPackageStatusDetails schema from the OpenAPI specification
type ModelPackageStatusDetails struct {
	Imagescanstatuses interface{} `json:"ImageScanStatuses,omitempty"`
	Validationstatuses interface{} `json:"ValidationStatuses"`
}

// AppDetails represents the AppDetails schema from the OpenAPI specification
type AppDetails struct {
	Apptype interface{} `json:"AppType,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Domainid interface{} `json:"DomainId,omitempty"`
	Spacename interface{} `json:"SpaceName,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Userprofilename interface{} `json:"UserProfileName,omitempty"`
	Appname interface{} `json:"AppName,omitempty"`
}

// StoppingCondition represents the StoppingCondition schema from the OpenAPI specification
type StoppingCondition struct {
	Maxruntimeinseconds interface{} `json:"MaxRuntimeInSeconds,omitempty"`
	Maxwaittimeinseconds interface{} `json:"MaxWaitTimeInSeconds,omitempty"`
}

// ExperimentSource represents the ExperimentSource schema from the OpenAPI specification
type ExperimentSource struct {
	Sourcetype interface{} `json:"SourceType,omitempty"`
	Sourcearn interface{} `json:"SourceArn"`
}

// TrialComponentArtifact represents the TrialComponentArtifact schema from the OpenAPI specification
type TrialComponentArtifact struct {
	Mediatype interface{} `json:"MediaType,omitempty"`
	Value interface{} `json:"Value"`
}

// ModelQuality represents the ModelQuality schema from the OpenAPI specification
type ModelQuality struct {
	Statistics interface{} `json:"Statistics,omitempty"`
	Constraints interface{} `json:"Constraints,omitempty"`
}

// ModelConfiguration represents the ModelConfiguration schema from the OpenAPI specification
type ModelConfiguration struct {
	Compilationjobname interface{} `json:"CompilationJobName,omitempty"`
	Environmentparameters interface{} `json:"EnvironmentParameters,omitempty"`
	Inferencespecificationname interface{} `json:"InferenceSpecificationName,omitempty"`
}

// CreateProjectOutput represents the CreateProjectOutput schema from the OpenAPI specification
type CreateProjectOutput struct {
	Projectarn interface{} `json:"ProjectArn"`
	Projectid interface{} `json:"ProjectId"`
}

// ListDevicesRequest represents the ListDevicesRequest schema from the OpenAPI specification
type ListDevicesRequest struct {
	Devicefleetname interface{} `json:"DeviceFleetName,omitempty"`
	Latestheartbeatafter interface{} `json:"LatestHeartbeatAfter,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Modelname interface{} `json:"ModelName,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// UpdateNotebookInstanceLifecycleConfigOutput represents the UpdateNotebookInstanceLifecycleConfigOutput schema from the OpenAPI specification
type UpdateNotebookInstanceLifecycleConfigOutput struct {
}

// TtlDuration represents the TtlDuration schema from the OpenAPI specification
type TtlDuration struct {
	Value interface{} `json:"Value,omitempty"`
	Unit interface{} `json:"Unit,omitempty"`
}

// MonitoringGroundTruthS3Input represents the MonitoringGroundTruthS3Input schema from the OpenAPI specification
type MonitoringGroundTruthS3Input struct {
	S3uri interface{} `json:"S3Uri,omitempty"`
}

// CreateModelOutput represents the CreateModelOutput schema from the OpenAPI specification
type CreateModelOutput struct {
	Modelarn interface{} `json:"ModelArn"`
}

// EndpointInput represents the EndpointInput schema from the OpenAPI specification
type EndpointInput struct {
	Probabilitythresholdattribute interface{} `json:"ProbabilityThresholdAttribute,omitempty"`
	Inferenceattribute interface{} `json:"InferenceAttribute,omitempty"`
	S3datadistributiontype interface{} `json:"S3DataDistributionType,omitempty"`
	Endpointname interface{} `json:"EndpointName"`
	Featuresattribute interface{} `json:"FeaturesAttribute,omitempty"`
	Localpath interface{} `json:"LocalPath"`
	Endtimeoffset interface{} `json:"EndTimeOffset,omitempty"`
	Probabilityattribute interface{} `json:"ProbabilityAttribute,omitempty"`
	S3inputmode interface{} `json:"S3InputMode,omitempty"`
	Starttimeoffset interface{} `json:"StartTimeOffset,omitempty"`
}

// MonitoringSchedule represents the MonitoringSchedule schema from the OpenAPI specification
type MonitoringSchedule struct {
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Monitoringscheduleconfig MonitoringScheduleConfig `json:"MonitoringScheduleConfig,omitempty"` // Configures the monitoring schedule and defines the monitoring job.
	Tags interface{} `json:"Tags,omitempty"`
	Endpointname interface{} `json:"EndpointName,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Lastmonitoringexecutionsummary MonitoringExecutionSummary `json:"LastMonitoringExecutionSummary,omitempty"` // Summary of information about the last monitoring job to run.
	Monitoringschedulearn interface{} `json:"MonitoringScheduleArn,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Monitoringschedulestatus interface{} `json:"MonitoringScheduleStatus,omitempty"`
	Monitoringtype interface{} `json:"MonitoringType,omitempty"`
	Monitoringschedulename interface{} `json:"MonitoringScheduleName,omitempty"`
}

// CreateImageVersionRequest represents the CreateImageVersionRequest schema from the OpenAPI specification
type CreateImageVersionRequest struct {
	Releasenotes interface{} `json:"ReleaseNotes,omitempty"`
	Aliases interface{} `json:"Aliases,omitempty"`
	Mlframework interface{} `json:"MLFramework,omitempty"`
	Vendorguidance interface{} `json:"VendorGuidance,omitempty"`
	Baseimage interface{} `json:"BaseImage"`
	Clienttoken interface{} `json:"ClientToken"`
	Imagename interface{} `json:"ImageName"`
	Jobtype interface{} `json:"JobType,omitempty"`
	Processor interface{} `json:"Processor,omitempty"`
	Programminglang interface{} `json:"ProgrammingLang,omitempty"`
	Horovod interface{} `json:"Horovod,omitempty"`
}

// RedshiftDatasetDefinition represents the RedshiftDatasetDefinition schema from the OpenAPI specification
type RedshiftDatasetDefinition struct {
	Outputs3uri interface{} `json:"OutputS3Uri"`
	Database string `json:"Database"` // The name of the Redshift database used in Redshift query execution.
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Outputcompression string `json:"OutputCompression,omitempty"` // The compression used for Redshift query results.
	Querystring string `json:"QueryString"` // The SQL query statements to be executed.
	Dbuser string `json:"DbUser"` // The database user name used in Redshift query execution.
	Clusterid string `json:"ClusterId"` // The Redshift cluster Identifier.
	Clusterrolearn interface{} `json:"ClusterRoleArn"`
	Outputformat string `json:"OutputFormat"` // The data storage format for Redshift query results.
}

// ListExperimentsResponse represents the ListExperimentsResponse schema from the OpenAPI specification
type ListExperimentsResponse struct {
	Experimentsummaries interface{} `json:"ExperimentSummaries,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CreateAlgorithmOutput represents the CreateAlgorithmOutput schema from the OpenAPI specification
type CreateAlgorithmOutput struct {
	Algorithmarn interface{} `json:"AlgorithmArn"`
}

// WorkforceVpcConfigResponse represents the WorkforceVpcConfigResponse schema from the OpenAPI specification
type WorkforceVpcConfigResponse struct {
	Vpcid interface{} `json:"VpcId"`
	Securitygroupids interface{} `json:"SecurityGroupIds"`
	Subnets interface{} `json:"Subnets"`
	Vpcendpointid interface{} `json:"VpcEndpointId,omitempty"`
}

// DescribeSpaceRequest represents the DescribeSpaceRequest schema from the OpenAPI specification
type DescribeSpaceRequest struct {
	Domainid interface{} `json:"DomainId"`
	Spacename interface{} `json:"SpaceName"`
}

// DynamicScalingConfiguration represents the DynamicScalingConfiguration schema from the OpenAPI specification
type DynamicScalingConfiguration struct {
	Scaleoutcooldown interface{} `json:"ScaleOutCooldown,omitempty"`
	Scalingpolicies interface{} `json:"ScalingPolicies,omitempty"`
	Maxcapacity interface{} `json:"MaxCapacity,omitempty"`
	Mincapacity interface{} `json:"MinCapacity,omitempty"`
	Scaleincooldown interface{} `json:"ScaleInCooldown,omitempty"`
}

// MultiModelConfig represents the MultiModelConfig schema from the OpenAPI specification
type MultiModelConfig struct {
	Modelcachesetting interface{} `json:"ModelCacheSetting,omitempty"`
}

// CandidateArtifactLocations represents the CandidateArtifactLocations schema from the OpenAPI specification
type CandidateArtifactLocations struct {
	Backtestresults interface{} `json:"BacktestResults,omitempty"`
	Explainability interface{} `json:"Explainability"`
	Modelinsights interface{} `json:"ModelInsights,omitempty"`
}

// GitConfigForUpdate represents the GitConfigForUpdate schema from the OpenAPI specification
type GitConfigForUpdate struct {
	Secretarn interface{} `json:"SecretArn,omitempty"`
}

// RetryPipelineExecutionRequest represents the RetryPipelineExecutionRequest schema from the OpenAPI specification
type RetryPipelineExecutionRequest struct {
	Clientrequesttoken interface{} `json:"ClientRequestToken"`
	Parallelismconfiguration interface{} `json:"ParallelismConfiguration,omitempty"`
	Pipelineexecutionarn interface{} `json:"PipelineExecutionArn"`
}

// ListAutoMLJobsRequest represents the ListAutoMLJobsRequest schema from the OpenAPI specification
type ListAutoMLJobsRequest struct {
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Lastmodifiedtimeafter interface{} `json:"LastModifiedTimeAfter,omitempty"`
	Lastmodifiedtimebefore interface{} `json:"LastModifiedTimeBefore,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Statusequals interface{} `json:"StatusEquals,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
}

// ListLineageGroupsResponse represents the ListLineageGroupsResponse schema from the OpenAPI specification
type ListLineageGroupsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Lineagegroupsummaries interface{} `json:"LineageGroupSummaries,omitempty"`
}

// UpdateUserProfileResponse represents the UpdateUserProfileResponse schema from the OpenAPI specification
type UpdateUserProfileResponse struct {
	Userprofilearn interface{} `json:"UserProfileArn,omitempty"`
}

// ListInferenceExperimentsResponse represents the ListInferenceExperimentsResponse schema from the OpenAPI specification
type ListInferenceExperimentsResponse struct {
	Inferenceexperiments interface{} `json:"InferenceExperiments,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CreateNotebookInstanceInput represents the CreateNotebookInstanceInput schema from the OpenAPI specification
type CreateNotebookInstanceInput struct {
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Notebookinstancename interface{} `json:"NotebookInstanceName"`
	Rootaccess interface{} `json:"RootAccess,omitempty"`
	Directinternetaccess interface{} `json:"DirectInternetAccess,omitempty"`
	Lifecycleconfigname interface{} `json:"LifecycleConfigName,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Platformidentifier interface{} `json:"PlatformIdentifier,omitempty"`
	Volumesizeingb interface{} `json:"VolumeSizeInGB,omitempty"`
	Subnetid interface{} `json:"SubnetId,omitempty"`
	Rolearn interface{} `json:"RoleArn"`
	Additionalcoderepositories interface{} `json:"AdditionalCodeRepositories,omitempty"`
	Defaultcoderepository interface{} `json:"DefaultCodeRepository,omitempty"`
	Acceleratortypes interface{} `json:"AcceleratorTypes,omitempty"`
	Instancemetadataserviceconfiguration interface{} `json:"InstanceMetadataServiceConfiguration,omitempty"`
	Securitygroupids interface{} `json:"SecurityGroupIds,omitempty"`
	Instancetype interface{} `json:"InstanceType"`
}

// UpdateExperimentRequest represents the UpdateExperimentRequest schema from the OpenAPI specification
type UpdateExperimentRequest struct {
	Description interface{} `json:"Description,omitempty"`
	Displayname interface{} `json:"DisplayName,omitempty"`
	Experimentname interface{} `json:"ExperimentName"`
}

// SendPipelineExecutionStepFailureRequest represents the SendPipelineExecutionStepFailureRequest schema from the OpenAPI specification
type SendPipelineExecutionStepFailureRequest struct {
	Callbacktoken interface{} `json:"CallbackToken"`
	Clientrequesttoken interface{} `json:"ClientRequestToken,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
}

// EndpointSummary represents the EndpointSummary schema from the OpenAPI specification
type EndpointSummary struct {
	Creationtime interface{} `json:"CreationTime"`
	Endpointarn interface{} `json:"EndpointArn"`
	Endpointname interface{} `json:"EndpointName"`
	Endpointstatus interface{} `json:"EndpointStatus"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime"`
}

// DescribeMonitoringScheduleResponse represents the DescribeMonitoringScheduleResponse schema from the OpenAPI specification
type DescribeMonitoringScheduleResponse struct {
	Monitoringscheduleconfig interface{} `json:"MonitoringScheduleConfig"`
	Creationtime interface{} `json:"CreationTime"`
	Monitoringtype interface{} `json:"MonitoringType,omitempty"`
	Monitoringschedulestatus interface{} `json:"MonitoringScheduleStatus"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime"`
	Lastmonitoringexecutionsummary interface{} `json:"LastMonitoringExecutionSummary,omitempty"`
	Monitoringschedulearn interface{} `json:"MonitoringScheduleArn"`
	Monitoringschedulename interface{} `json:"MonitoringScheduleName"`
	Endpointname interface{} `json:"EndpointName,omitempty"`
}

// DescribePipelineExecutionResponse represents the DescribePipelineExecutionResponse schema from the OpenAPI specification
type DescribePipelineExecutionResponse struct {
	Createdby UserContext `json:"CreatedBy,omitempty"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	Lastmodifiedby UserContext `json:"LastModifiedBy,omitempty"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	Pipelineexecutionarn interface{} `json:"PipelineExecutionArn,omitempty"`
	Pipelineexecutiondescription interface{} `json:"PipelineExecutionDescription,omitempty"`
	Pipelineexperimentconfig PipelineExperimentConfig `json:"PipelineExperimentConfig,omitempty"` // Specifies the names of the experiment and trial created by a pipeline.
	Pipelineexecutiondisplayname interface{} `json:"PipelineExecutionDisplayName,omitempty"`
	Parallelismconfiguration interface{} `json:"ParallelismConfiguration,omitempty"`
	Pipelinearn interface{} `json:"PipelineArn,omitempty"`
	Pipelineexecutionstatus interface{} `json:"PipelineExecutionStatus,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Selectiveexecutionconfig interface{} `json:"SelectiveExecutionConfig,omitempty"`
}

// ListPipelineParametersForExecutionResponse represents the ListPipelineParametersForExecutionResponse schema from the OpenAPI specification
type ListPipelineParametersForExecutionResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Pipelineparameters interface{} `json:"PipelineParameters,omitempty"`
}

// ListAppsRequest represents the ListAppsRequest schema from the OpenAPI specification
type ListAppsRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Spacenameequals interface{} `json:"SpaceNameEquals,omitempty"`
	Userprofilenameequals interface{} `json:"UserProfileNameEquals,omitempty"`
	Domainidequals interface{} `json:"DomainIdEquals,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// UpdateMonitoringScheduleResponse represents the UpdateMonitoringScheduleResponse schema from the OpenAPI specification
type UpdateMonitoringScheduleResponse struct {
	Monitoringschedulearn interface{} `json:"MonitoringScheduleArn"`
}

// WorkforceVpcConfigRequest represents the WorkforceVpcConfigRequest schema from the OpenAPI specification
type WorkforceVpcConfigRequest struct {
	Subnets interface{} `json:"Subnets,omitempty"`
	Vpcid interface{} `json:"VpcId,omitempty"`
	Securitygroupids interface{} `json:"SecurityGroupIds,omitempty"`
}

// UpdateAppImageConfigResponse represents the UpdateAppImageConfigResponse schema from the OpenAPI specification
type UpdateAppImageConfigResponse struct {
	Appimageconfigarn interface{} `json:"AppImageConfigArn,omitempty"`
}

// AutoMLSecurityConfig represents the AutoMLSecurityConfig schema from the OpenAPI specification
type AutoMLSecurityConfig struct {
	Enableintercontainertrafficencryption interface{} `json:"EnableInterContainerTrafficEncryption,omitempty"`
	Volumekmskeyid interface{} `json:"VolumeKmsKeyId,omitempty"`
	Vpcconfig interface{} `json:"VpcConfig,omitempty"`
}

// RecommendationJobStoppingConditions represents the RecommendationJobStoppingConditions schema from the OpenAPI specification
type RecommendationJobStoppingConditions struct {
	Flatinvocations interface{} `json:"FlatInvocations,omitempty"`
	Maxinvocations interface{} `json:"MaxInvocations,omitempty"`
	Modellatencythresholds interface{} `json:"ModelLatencyThresholds,omitempty"`
}

// DescribeEdgePackagingJobResponse represents the DescribeEdgePackagingJobResponse schema from the OpenAPI specification
type DescribeEdgePackagingJobResponse struct {
	Edgepackagingjobarn interface{} `json:"EdgePackagingJobArn"`
	Outputconfig interface{} `json:"OutputConfig,omitempty"`
	Edgepackagingjobstatusmessage interface{} `json:"EdgePackagingJobStatusMessage,omitempty"`
	Modelartifact interface{} `json:"ModelArtifact,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Modelsignature interface{} `json:"ModelSignature,omitempty"`
	Modelversion interface{} `json:"ModelVersion,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Compilationjobname interface{} `json:"CompilationJobName,omitempty"`
	Resourcekey interface{} `json:"ResourceKey,omitempty"`
	Edgepackagingjobstatus interface{} `json:"EdgePackagingJobStatus"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Presetdeploymentoutput interface{} `json:"PresetDeploymentOutput,omitempty"`
	Modelname interface{} `json:"ModelName,omitempty"`
	Edgepackagingjobname interface{} `json:"EdgePackagingJobName"`
}

// EdgePackagingJobSummary represents the EdgePackagingJobSummary schema from the OpenAPI specification
type EdgePackagingJobSummary struct {
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Modelname interface{} `json:"ModelName,omitempty"`
	Modelversion interface{} `json:"ModelVersion,omitempty"`
	Compilationjobname interface{} `json:"CompilationJobName,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Edgepackagingjobarn interface{} `json:"EdgePackagingJobArn"`
	Edgepackagingjobname interface{} `json:"EdgePackagingJobName"`
	Edgepackagingjobstatus interface{} `json:"EdgePackagingJobStatus"`
}

// DescribeModelPackageGroupInput represents the DescribeModelPackageGroupInput schema from the OpenAPI specification
type DescribeModelPackageGroupInput struct {
	Modelpackagegroupname interface{} `json:"ModelPackageGroupName"`
}

// ExperimentConfig represents the ExperimentConfig schema from the OpenAPI specification
type ExperimentConfig struct {
	Experimentname interface{} `json:"ExperimentName,omitempty"`
	Runname interface{} `json:"RunName,omitempty"`
	Trialcomponentdisplayname interface{} `json:"TrialComponentDisplayName,omitempty"`
	Trialname interface{} `json:"TrialName,omitempty"`
}

// OfflineStoreStatus represents the OfflineStoreStatus schema from the OpenAPI specification
type OfflineStoreStatus struct {
	Blockedreason interface{} `json:"BlockedReason,omitempty"`
	Status interface{} `json:"Status"`
}

// ModelPackage represents the ModelPackage schema from the OpenAPI specification
type ModelPackage struct {
	Tags interface{} `json:"Tags,omitempty"`
	Domain interface{} `json:"Domain,omitempty"`
	Modelpackagegroupname interface{} `json:"ModelPackageGroupName,omitempty"`
	Customermetadataproperties interface{} `json:"CustomerMetadataProperties,omitempty"`
	Sourcealgorithmspecification interface{} `json:"SourceAlgorithmSpecification,omitempty"`
	Task interface{} `json:"Task,omitempty"`
	Certifyformarketplace interface{} `json:"CertifyForMarketplace,omitempty"`
	Driftcheckbaselines interface{} `json:"DriftCheckBaselines,omitempty"`
	Approvaldescription interface{} `json:"ApprovalDescription,omitempty"`
	Modelmetrics interface{} `json:"ModelMetrics,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Metadataproperties interface{} `json:"MetadataProperties,omitempty"`
	Modelpackagestatusdetails interface{} `json:"ModelPackageStatusDetails,omitempty"`
	Modelpackagestatus interface{} `json:"ModelPackageStatus,omitempty"`
	Inferencespecification interface{} `json:"InferenceSpecification,omitempty"`
	Validationspecification interface{} `json:"ValidationSpecification,omitempty"`
	Lastmodifiedby interface{} `json:"LastModifiedBy,omitempty"`
	Modelpackagedescription interface{} `json:"ModelPackageDescription,omitempty"`
	Modelpackageversion interface{} `json:"ModelPackageVersion,omitempty"`
	Createdby interface{} `json:"CreatedBy,omitempty"`
	Samplepayloadurl interface{} `json:"SamplePayloadUrl,omitempty"`
	Additionalinferencespecifications interface{} `json:"AdditionalInferenceSpecifications,omitempty"`
	Modelapprovalstatus interface{} `json:"ModelApprovalStatus,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Modelpackagearn interface{} `json:"ModelPackageArn,omitempty"`
	Modelpackagename interface{} `json:"ModelPackageName,omitempty"`
}

// ListMonitoringExecutionsResponse represents the ListMonitoringExecutionsResponse schema from the OpenAPI specification
type ListMonitoringExecutionsResponse struct {
	Monitoringexecutionsummaries interface{} `json:"MonitoringExecutionSummaries"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DebugHookConfig represents the DebugHookConfig schema from the OpenAPI specification
type DebugHookConfig struct {
	Localpath interface{} `json:"LocalPath,omitempty"`
	S3outputpath interface{} `json:"S3OutputPath"`
	Collectionconfigurations interface{} `json:"CollectionConfigurations,omitempty"`
	Hookparameters interface{} `json:"HookParameters,omitempty"`
}

// DescribeAlgorithmInput represents the DescribeAlgorithmInput schema from the OpenAPI specification
type DescribeAlgorithmInput struct {
	Algorithmname interface{} `json:"AlgorithmName"`
}

// UpdateEndpointWeightsAndCapacitiesInput represents the UpdateEndpointWeightsAndCapacitiesInput schema from the OpenAPI specification
type UpdateEndpointWeightsAndCapacitiesInput struct {
	Desiredweightsandcapacities interface{} `json:"DesiredWeightsAndCapacities"`
	Endpointname interface{} `json:"EndpointName"`
}

// ListResourceCatalogsRequest represents the ListResourceCatalogsRequest schema from the OpenAPI specification
type ListResourceCatalogsRequest struct {
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
}

// MonitoringS3Output represents the MonitoringS3Output schema from the OpenAPI specification
type MonitoringS3Output struct {
	Localpath interface{} `json:"LocalPath"`
	S3uploadmode interface{} `json:"S3UploadMode,omitempty"`
	S3uri interface{} `json:"S3Uri"`
}

// SecondaryStatusTransition represents the SecondaryStatusTransition schema from the OpenAPI specification
type SecondaryStatusTransition struct {
	Endtime interface{} `json:"EndTime,omitempty"`
	Starttime interface{} `json:"StartTime"`
	Status interface{} `json:"Status"`
	Statusmessage interface{} `json:"StatusMessage,omitempty"`
}

// ListSpacesRequest represents the ListSpacesRequest schema from the OpenAPI specification
type ListSpacesRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Spacenamecontains interface{} `json:"SpaceNameContains,omitempty"`
	Domainidequals interface{} `json:"DomainIdEquals,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// ArtifactSource represents the ArtifactSource schema from the OpenAPI specification
type ArtifactSource struct {
	Sourceuri interface{} `json:"SourceUri"`
	Sourcetypes interface{} `json:"SourceTypes,omitempty"`
}

// UpdateModelCardRequest represents the UpdateModelCardRequest schema from the OpenAPI specification
type UpdateModelCardRequest struct {
	Modelcardname interface{} `json:"ModelCardName"`
	Modelcardstatus interface{} `json:"ModelCardStatus,omitempty"`
	Content interface{} `json:"Content,omitempty"`
}

// ModelCardExportJobSummary represents the ModelCardExportJobSummary schema from the OpenAPI specification
type ModelCardExportJobSummary struct {
	Modelcardexportjobname interface{} `json:"ModelCardExportJobName"`
	Modelcardname interface{} `json:"ModelCardName"`
	Modelcardversion interface{} `json:"ModelCardVersion"`
	Status interface{} `json:"Status"`
	Createdat interface{} `json:"CreatedAt"`
	Lastmodifiedat interface{} `json:"LastModifiedAt"`
	Modelcardexportjobarn interface{} `json:"ModelCardExportJobArn"`
}

// ListDeviceFleetsResponse represents the ListDeviceFleetsResponse schema from the OpenAPI specification
type ListDeviceFleetsResponse struct {
	Devicefleetsummaries interface{} `json:"DeviceFleetSummaries"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// EdgeDeploymentStatus represents the EdgeDeploymentStatus schema from the OpenAPI specification
type EdgeDeploymentStatus struct {
	Edgedeploymentstagestarttime interface{} `json:"EdgeDeploymentStageStartTime,omitempty"`
	Edgedeploymentstatusmessage interface{} `json:"EdgeDeploymentStatusMessage,omitempty"`
	Edgedeploymentsuccessinstage interface{} `json:"EdgeDeploymentSuccessInStage"`
	Stagestatus interface{} `json:"StageStatus"`
	Edgedeploymentfailedinstage interface{} `json:"EdgeDeploymentFailedInStage"`
	Edgedeploymentpendinginstage interface{} `json:"EdgeDeploymentPendingInStage"`
}

// Stairs represents the Stairs schema from the OpenAPI specification
type Stairs struct {
	Durationinseconds interface{} `json:"DurationInSeconds,omitempty"`
	Numberofsteps interface{} `json:"NumberOfSteps,omitempty"`
	Usersperstep interface{} `json:"UsersPerStep,omitempty"`
}

// HyperParameterTuningJobSummary represents the HyperParameterTuningJobSummary schema from the OpenAPI specification
type HyperParameterTuningJobSummary struct {
	Strategy interface{} `json:"Strategy"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Objectivestatuscounters interface{} `json:"ObjectiveStatusCounters"`
	Hyperparametertuningendtime interface{} `json:"HyperParameterTuningEndTime,omitempty"`
	Hyperparametertuningjobstatus interface{} `json:"HyperParameterTuningJobStatus"`
	Resourcelimits interface{} `json:"ResourceLimits,omitempty"`
	Hyperparametertuningjobarn interface{} `json:"HyperParameterTuningJobArn"`
	Trainingjobstatuscounters interface{} `json:"TrainingJobStatusCounters"`
	Creationtime interface{} `json:"CreationTime"`
	Hyperparametertuningjobname interface{} `json:"HyperParameterTuningJobName"`
}

// TensorBoardOutputConfig represents the TensorBoardOutputConfig schema from the OpenAPI specification
type TensorBoardOutputConfig struct {
	Localpath interface{} `json:"LocalPath,omitempty"`
	S3outputpath interface{} `json:"S3OutputPath"`
}

// UpdateDeviceFleetRequest represents the UpdateDeviceFleetRequest schema from the OpenAPI specification
type UpdateDeviceFleetRequest struct {
	Description interface{} `json:"Description,omitempty"`
	Devicefleetname interface{} `json:"DeviceFleetName"`
	Enableiotrolealias interface{} `json:"EnableIotRoleAlias,omitempty"`
	Outputconfig interface{} `json:"OutputConfig"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
}

// InferenceSpecification represents the InferenceSpecification schema from the OpenAPI specification
type InferenceSpecification struct {
	Containers interface{} `json:"Containers"`
	Supportedcontenttypes interface{} `json:"SupportedContentTypes"`
	Supportedrealtimeinferenceinstancetypes interface{} `json:"SupportedRealtimeInferenceInstanceTypes,omitempty"`
	Supportedresponsemimetypes interface{} `json:"SupportedResponseMIMETypes"`
	Supportedtransforminstancetypes interface{} `json:"SupportedTransformInstanceTypes,omitempty"`
}

// EndpointMetadata represents the EndpointMetadata schema from the OpenAPI specification
type EndpointMetadata struct {
	Endpointconfigname interface{} `json:"EndpointConfigName,omitempty"`
	Endpointname interface{} `json:"EndpointName"`
	Endpointstatus interface{} `json:"EndpointStatus,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
}

// RecommendationJobInputConfig represents the RecommendationJobInputConfig schema from the OpenAPI specification
type RecommendationJobInputConfig struct {
	Jobdurationinseconds interface{} `json:"JobDurationInSeconds,omitempty"`
	Modelpackageversionarn interface{} `json:"ModelPackageVersionArn,omitempty"`
	Resourcelimit interface{} `json:"ResourceLimit,omitempty"`
	Modelname interface{} `json:"ModelName,omitempty"`
	Containerconfig interface{} `json:"ContainerConfig,omitempty"`
	Endpoints interface{} `json:"Endpoints,omitempty"`
	Trafficpattern interface{} `json:"TrafficPattern,omitempty"`
	Endpointconfigurations interface{} `json:"EndpointConfigurations,omitempty"`
	Volumekmskeyid interface{} `json:"VolumeKmsKeyId,omitempty"`
	Vpcconfig interface{} `json:"VpcConfig,omitempty"`
}

// CreateNotebookInstanceOutput represents the CreateNotebookInstanceOutput schema from the OpenAPI specification
type CreateNotebookInstanceOutput struct {
	Notebookinstancearn interface{} `json:"NotebookInstanceArn,omitempty"`
}

// DriftCheckBaselines represents the DriftCheckBaselines schema from the OpenAPI specification
type DriftCheckBaselines struct {
	Bias interface{} `json:"Bias,omitempty"`
	Explainability interface{} `json:"Explainability,omitempty"`
	Modeldataquality interface{} `json:"ModelDataQuality,omitempty"`
	Modelquality interface{} `json:"ModelQuality,omitempty"`
}

// ListTagsOutput represents the ListTagsOutput schema from the OpenAPI specification
type ListTagsOutput struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// ProcessingClusterConfig represents the ProcessingClusterConfig schema from the OpenAPI specification
type ProcessingClusterConfig struct {
	Instancecount interface{} `json:"InstanceCount"`
	Instancetype interface{} `json:"InstanceType"`
	Volumekmskeyid interface{} `json:"VolumeKmsKeyId,omitempty"`
	Volumesizeingb interface{} `json:"VolumeSizeInGB"`
}

// CaptureOption represents the CaptureOption schema from the OpenAPI specification
type CaptureOption struct {
	Capturemode interface{} `json:"CaptureMode"`
}

// AsyncInferenceNotificationConfig represents the AsyncInferenceNotificationConfig schema from the OpenAPI specification
type AsyncInferenceNotificationConfig struct {
	Errortopic interface{} `json:"ErrorTopic,omitempty"`
	Includeinferenceresponsein interface{} `json:"IncludeInferenceResponseIn,omitempty"`
	Successtopic interface{} `json:"SuccessTopic,omitempty"`
}

// ListHyperParameterTuningJobsResponse represents the ListHyperParameterTuningJobsResponse schema from the OpenAPI specification
type ListHyperParameterTuningJobsResponse struct {
	Hyperparametertuningjobsummaries interface{} `json:"HyperParameterTuningJobSummaries"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DescribeHubRequest represents the DescribeHubRequest schema from the OpenAPI specification
type DescribeHubRequest struct {
	Hubname interface{} `json:"HubName"`
}

// IntegerParameterRange represents the IntegerParameterRange schema from the OpenAPI specification
type IntegerParameterRange struct {
	Maxvalue interface{} `json:"MaxValue"`
	Minvalue interface{} `json:"MinValue"`
	Name interface{} `json:"Name"`
	Scalingtype interface{} `json:"ScalingType,omitempty"`
}

// DescribeTransformJobRequest represents the DescribeTransformJobRequest schema from the OpenAPI specification
type DescribeTransformJobRequest struct {
	Transformjobname interface{} `json:"TransformJobName"`
}

// AppImageConfigDetails represents the AppImageConfigDetails schema from the OpenAPI specification
type AppImageConfigDetails struct {
	Appimageconfigname interface{} `json:"AppImageConfigName,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Kernelgatewayimageconfig interface{} `json:"KernelGatewayImageConfig,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Appimageconfigarn interface{} `json:"AppImageConfigArn,omitempty"`
}

// NeoVpcConfig represents the NeoVpcConfig schema from the OpenAPI specification
type NeoVpcConfig struct {
	Securitygroupids interface{} `json:"SecurityGroupIds"`
	Subnets interface{} `json:"Subnets"`
}

// ListFlowDefinitionsResponse represents the ListFlowDefinitionsResponse schema from the OpenAPI specification
type ListFlowDefinitionsResponse struct {
	Flowdefinitionsummaries interface{} `json:"FlowDefinitionSummaries"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListInferenceRecommendationsJobsResponse represents the ListInferenceRecommendationsJobsResponse schema from the OpenAPI specification
type ListInferenceRecommendationsJobsResponse struct {
	Inferencerecommendationsjobs interface{} `json:"InferenceRecommendationsJobs"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListUserProfilesResponse represents the ListUserProfilesResponse schema from the OpenAPI specification
type ListUserProfilesResponse struct {
	Userprofiles interface{} `json:"UserProfiles,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// UpdateDomainRequest represents the UpdateDomainRequest schema from the OpenAPI specification
type UpdateDomainRequest struct {
	Domainid interface{} `json:"DomainId"`
	Domainsettingsforupdate interface{} `json:"DomainSettingsForUpdate,omitempty"`
	Appsecuritygroupmanagement interface{} `json:"AppSecurityGroupManagement,omitempty"`
	Defaultspacesettings interface{} `json:"DefaultSpaceSettings,omitempty"`
	Defaultusersettings interface{} `json:"DefaultUserSettings,omitempty"`
}

// OidcConfig represents the OidcConfig schema from the OpenAPI specification
type OidcConfig struct {
	Jwksuri interface{} `json:"JwksUri"`
	Logoutendpoint interface{} `json:"LogoutEndpoint"`
	Tokenendpoint interface{} `json:"TokenEndpoint"`
	Userinfoendpoint interface{} `json:"UserInfoEndpoint"`
	Authorizationendpoint interface{} `json:"AuthorizationEndpoint"`
	Clientid interface{} `json:"ClientId"`
	Clientsecret interface{} `json:"ClientSecret"`
	Issuer interface{} `json:"Issuer"`
}

// CreateAutoMLJobV2Request represents the CreateAutoMLJobV2Request schema from the OpenAPI specification
type CreateAutoMLJobV2Request struct {
	Tags interface{} `json:"Tags,omitempty"`
	Automljobinputdataconfig interface{} `json:"AutoMLJobInputDataConfig"`
	Datasplitconfig interface{} `json:"DataSplitConfig,omitempty"`
	Rolearn interface{} `json:"RoleArn"`
	Securityconfig interface{} `json:"SecurityConfig,omitempty"`
	Automljobname interface{} `json:"AutoMLJobName"`
	Outputdataconfig interface{} `json:"OutputDataConfig"`
	Automljobobjective interface{} `json:"AutoMLJobObjective,omitempty"`
	Automlproblemtypeconfig interface{} `json:"AutoMLProblemTypeConfig"`
	Modeldeployconfig interface{} `json:"ModelDeployConfig,omitempty"`
}

// DomainSettingsForUpdate represents the DomainSettingsForUpdate schema from the OpenAPI specification
type DomainSettingsForUpdate struct {
	Securitygroupids interface{} `json:"SecurityGroupIds,omitempty"`
	Executionroleidentityconfig interface{} `json:"ExecutionRoleIdentityConfig,omitempty"`
	Rstudioserverprodomainsettingsforupdate interface{} `json:"RStudioServerProDomainSettingsForUpdate,omitempty"`
}

// ListCompilationJobsRequest represents the ListCompilationJobsRequest schema from the OpenAPI specification
type ListCompilationJobsRequest struct {
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Statusequals interface{} `json:"StatusEquals,omitempty"`
	Lastmodifiedtimeafter interface{} `json:"LastModifiedTimeAfter,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Lastmodifiedtimebefore interface{} `json:"LastModifiedTimeBefore,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
}

// AutoMLJobObjective represents the AutoMLJobObjective schema from the OpenAPI specification
type AutoMLJobObjective struct {
	Metricname interface{} `json:"MetricName"`
}

// KernelGatewayImageConfig represents the KernelGatewayImageConfig schema from the OpenAPI specification
type KernelGatewayImageConfig struct {
	Filesystemconfig interface{} `json:"FileSystemConfig,omitempty"`
	Kernelspecs interface{} `json:"KernelSpecs"`
}

// CreateContextRequest represents the CreateContextRequest schema from the OpenAPI specification
type CreateContextRequest struct {
	Description interface{} `json:"Description,omitempty"`
	Properties interface{} `json:"Properties,omitempty"`
	Source interface{} `json:"Source"`
	Tags interface{} `json:"Tags,omitempty"`
	Contextname interface{} `json:"ContextName"`
	Contexttype interface{} `json:"ContextType"`
}

// MonitoringOutputConfig represents the MonitoringOutputConfig schema from the OpenAPI specification
type MonitoringOutputConfig struct {
	Monitoringoutputs interface{} `json:"MonitoringOutputs"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
}

// DescribeHubResponse represents the DescribeHubResponse schema from the OpenAPI specification
type DescribeHubResponse struct {
	Hubname interface{} `json:"HubName"`
	Hubsearchkeywords interface{} `json:"HubSearchKeywords,omitempty"`
	Creationtime interface{} `json:"CreationTime"`
	Hubarn interface{} `json:"HubArn"`
	Hubdisplayname interface{} `json:"HubDisplayName,omitempty"`
	Hubstatus interface{} `json:"HubStatus"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	S3storageconfig interface{} `json:"S3StorageConfig,omitempty"`
	Hubdescription interface{} `json:"HubDescription,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime"`
}

// CreateActionResponse represents the CreateActionResponse schema from the OpenAPI specification
type CreateActionResponse struct {
	Actionarn interface{} `json:"ActionArn,omitempty"`
}

// CategoricalParameter represents the CategoricalParameter schema from the OpenAPI specification
type CategoricalParameter struct {
	Name interface{} `json:"Name"`
	Value interface{} `json:"Value"`
}

// ListMonitoringAlertsRequest represents the ListMonitoringAlertsRequest schema from the OpenAPI specification
type ListMonitoringAlertsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Monitoringschedulename interface{} `json:"MonitoringScheduleName"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ModelPackageValidationSpecification represents the ModelPackageValidationSpecification schema from the OpenAPI specification
type ModelPackageValidationSpecification struct {
	Validationrole interface{} `json:"ValidationRole"`
	Validationprofiles interface{} `json:"ValidationProfiles"`
}

// JupyterServerAppSettings represents the JupyterServerAppSettings schema from the OpenAPI specification
type JupyterServerAppSettings struct {
	Coderepositories interface{} `json:"CodeRepositories,omitempty"`
	Defaultresourcespec interface{} `json:"DefaultResourceSpec,omitempty"`
	Lifecycleconfigarns interface{} `json:"LifecycleConfigArns,omitempty"`
}

// ProcessingResources represents the ProcessingResources schema from the OpenAPI specification
type ProcessingResources struct {
	Clusterconfig interface{} `json:"ClusterConfig"`
}

// TrialComponentParameterValue represents the TrialComponentParameterValue schema from the OpenAPI specification
type TrialComponentParameterValue struct {
	Numbervalue interface{} `json:"NumberValue,omitempty"`
	Stringvalue interface{} `json:"StringValue,omitempty"`
}

// GetSearchSuggestionsRequest represents the GetSearchSuggestionsRequest schema from the OpenAPI specification
type GetSearchSuggestionsRequest struct {
	Suggestionquery interface{} `json:"SuggestionQuery,omitempty"`
	Resource interface{} `json:"Resource"`
}

// DeleteFlowDefinitionRequest represents the DeleteFlowDefinitionRequest schema from the OpenAPI specification
type DeleteFlowDefinitionRequest struct {
	Flowdefinitionname interface{} `json:"FlowDefinitionName"`
}

// DescribeModelExplainabilityJobDefinitionRequest represents the DescribeModelExplainabilityJobDefinitionRequest schema from the OpenAPI specification
type DescribeModelExplainabilityJobDefinitionRequest struct {
	Jobdefinitionname interface{} `json:"JobDefinitionName"`
}

// ListEdgePackagingJobsResponse represents the ListEdgePackagingJobsResponse schema from the OpenAPI specification
type ListEdgePackagingJobsResponse struct {
	Edgepackagingjobsummaries interface{} `json:"EdgePackagingJobSummaries"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// SubscribedWorkteam represents the SubscribedWorkteam schema from the OpenAPI specification
type SubscribedWorkteam struct {
	Listingid interface{} `json:"ListingId,omitempty"`
	Marketplacedescription interface{} `json:"MarketplaceDescription,omitempty"`
	Marketplacetitle interface{} `json:"MarketplaceTitle,omitempty"`
	Sellername interface{} `json:"SellerName,omitempty"`
	Workteamarn interface{} `json:"WorkteamArn"`
}

// DataCaptureConfig represents the DataCaptureConfig schema from the OpenAPI specification
type DataCaptureConfig struct {
	Captureoptions interface{} `json:"CaptureOptions"`
	Destinations3uri interface{} `json:"DestinationS3Uri"`
	Enablecapture interface{} `json:"EnableCapture,omitempty"`
	Initialsamplingpercentage interface{} `json:"InitialSamplingPercentage"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Capturecontenttypeheader interface{} `json:"CaptureContentTypeHeader,omitempty"`
}

// ScalingPolicyMetric represents the ScalingPolicyMetric schema from the OpenAPI specification
type ScalingPolicyMetric struct {
	Modellatency interface{} `json:"ModelLatency,omitempty"`
	Invocationsperinstance interface{} `json:"InvocationsPerInstance,omitempty"`
}

// DeleteInferenceExperimentResponse represents the DeleteInferenceExperimentResponse schema from the OpenAPI specification
type DeleteInferenceExperimentResponse struct {
	Inferenceexperimentarn interface{} `json:"InferenceExperimentArn"`
}

// CreateTransformJobRequest represents the CreateTransformJobRequest schema from the OpenAPI specification
type CreateTransformJobRequest struct {
	Environment interface{} `json:"Environment,omitempty"`
	Modelclientconfig interface{} `json:"ModelClientConfig,omitempty"`
	Modelname interface{} `json:"ModelName"`
	Maxpayloadinmb interface{} `json:"MaxPayloadInMB,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Transforminput interface{} `json:"TransformInput"`
	Batchstrategy interface{} `json:"BatchStrategy,omitempty"`
	Datacaptureconfig interface{} `json:"DataCaptureConfig,omitempty"`
	Maxconcurrenttransforms interface{} `json:"MaxConcurrentTransforms,omitempty"`
	Transformjobname interface{} `json:"TransformJobName"`
	Experimentconfig ExperimentConfig `json:"ExperimentConfig,omitempty"` // <p>Associates a SageMaker job as a trial component with an experiment and trial. Specified when you call the following APIs:</p> <ul> <li> <p> <a href="https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateProcessingJob.html">CreateProcessingJob</a> </p> </li> <li> <p> <a href="https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTrainingJob.html">CreateTrainingJob</a> </p> </li> <li> <p> <a href="https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTransformJob.html">CreateTransformJob</a> </p> </li> </ul>
	Transformoutput interface{} `json:"TransformOutput"`
	Transformresources interface{} `json:"TransformResources"`
	Dataprocessing interface{} `json:"DataProcessing,omitempty"`
}

// DeleteAssociationResponse represents the DeleteAssociationResponse schema from the OpenAPI specification
type DeleteAssociationResponse struct {
	Destinationarn interface{} `json:"DestinationArn,omitempty"`
	Sourcearn interface{} `json:"SourceArn,omitempty"`
}

// ProcessingInput represents the ProcessingInput schema from the OpenAPI specification
type ProcessingInput struct {
	S3input interface{} `json:"S3Input,omitempty"`
	Appmanaged interface{} `json:"AppManaged,omitempty"`
	Datasetdefinition interface{} `json:"DatasetDefinition,omitempty"`
	Inputname interface{} `json:"InputName"`
}

// AutoMLS3DataSource represents the AutoMLS3DataSource schema from the OpenAPI specification
type AutoMLS3DataSource struct {
	S3uri interface{} `json:"S3Uri"`
	S3datatype interface{} `json:"S3DataType"`
}

// DeleteDeviceFleetRequest represents the DeleteDeviceFleetRequest schema from the OpenAPI specification
type DeleteDeviceFleetRequest struct {
	Devicefleetname interface{} `json:"DeviceFleetName"`
}

// Pipeline represents the Pipeline schema from the OpenAPI specification
type Pipeline struct {
	Createdby UserContext `json:"CreatedBy,omitempty"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	Lastruntime interface{} `json:"LastRunTime,omitempty"`
	Parallelismconfiguration interface{} `json:"ParallelismConfiguration,omitempty"`
	Pipelinedescription interface{} `json:"PipelineDescription,omitempty"`
	Lastmodifiedby UserContext `json:"LastModifiedBy,omitempty"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Pipelinearn interface{} `json:"PipelineArn,omitempty"`
	Pipelinedisplayname interface{} `json:"PipelineDisplayName,omitempty"`
	Pipelinename interface{} `json:"PipelineName,omitempty"`
	Pipelinestatus interface{} `json:"PipelineStatus,omitempty"`
}

// AsyncInferenceClientConfig represents the AsyncInferenceClientConfig schema from the OpenAPI specification
type AsyncInferenceClientConfig struct {
	Maxconcurrentinvocationsperinstance interface{} `json:"MaxConcurrentInvocationsPerInstance,omitempty"`
}

// MonitoringExecutionSummary represents the MonitoringExecutionSummary schema from the OpenAPI specification
type MonitoringExecutionSummary struct {
	Monitoringschedulename interface{} `json:"MonitoringScheduleName"`
	Scheduledtime interface{} `json:"ScheduledTime"`
	Creationtime interface{} `json:"CreationTime"`
	Endpointname interface{} `json:"EndpointName,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Monitoringexecutionstatus interface{} `json:"MonitoringExecutionStatus"`
	Monitoringjobdefinitionname interface{} `json:"MonitoringJobDefinitionName,omitempty"`
	Processingjobarn interface{} `json:"ProcessingJobArn,omitempty"`
	Monitoringtype interface{} `json:"MonitoringType,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime"`
}

// DescribeContextRequest represents the DescribeContextRequest schema from the OpenAPI specification
type DescribeContextRequest struct {
	Contextname interface{} `json:"ContextName"`
}

// ModelPackageValidationProfile represents the ModelPackageValidationProfile schema from the OpenAPI specification
type ModelPackageValidationProfile struct {
	Transformjobdefinition interface{} `json:"TransformJobDefinition"`
	Profilename interface{} `json:"ProfileName"`
}

// DeleteTrialRequest represents the DeleteTrialRequest schema from the OpenAPI specification
type DeleteTrialRequest struct {
	Trialname interface{} `json:"TrialName"`
}

// DeleteProjectInput represents the DeleteProjectInput schema from the OpenAPI specification
type DeleteProjectInput struct {
	Projectname interface{} `json:"ProjectName"`
}

// TrafficRoutingConfig represents the TrafficRoutingConfig schema from the OpenAPI specification
type TrafficRoutingConfig struct {
	Canarysize interface{} `json:"CanarySize,omitempty"`
	Linearstepsize interface{} `json:"LinearStepSize,omitempty"`
	TypeField interface{} `json:"Type"`
	Waitintervalinseconds interface{} `json:"WaitIntervalInSeconds"`
}

// HyperParameterTuningJobConsumedResources represents the HyperParameterTuningJobConsumedResources schema from the OpenAPI specification
type HyperParameterTuningJobConsumedResources struct {
	Runtimeinseconds interface{} `json:"RuntimeInSeconds,omitempty"`
}

// DescribeEndpointInput represents the DescribeEndpointInput schema from the OpenAPI specification
type DescribeEndpointInput struct {
	Endpointname interface{} `json:"EndpointName"`
}

// GetSagemakerServicecatalogPortfolioStatusOutput represents the GetSagemakerServicecatalogPortfolioStatusOutput schema from the OpenAPI specification
type GetSagemakerServicecatalogPortfolioStatusOutput struct {
	Status interface{} `json:"Status,omitempty"`
}

// Model represents the Model schema from the OpenAPI specification
type Model struct {
	Containers interface{} `json:"Containers,omitempty"`
	Inferenceexecutionconfig InferenceExecutionConfig `json:"InferenceExecutionConfig,omitempty"` // Specifies details about how containers in a multi-container endpoint are run.
	Modelname interface{} `json:"ModelName,omitempty"`
	Primarycontainer ContainerDefinition `json:"PrimaryContainer,omitempty"` // Describes the container, as part of model definition.
	Tags interface{} `json:"Tags,omitempty"`
	Vpcconfig VpcConfig `json:"VpcConfig,omitempty"` // Specifies a VPC that your training jobs and hosted models have access to. Control access to and from your training and model containers by configuring the VPC. For more information, see <a href="https://docs.aws.amazon.com/sagemaker/latest/dg/host-vpc.html">Protect Endpoints by Using an Amazon Virtual Private Cloud</a> and <a href="https://docs.aws.amazon.com/sagemaker/latest/dg/train-vpc.html">Protect Training Jobs by Using an Amazon Virtual Private Cloud</a>.
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Enablenetworkisolation interface{} `json:"EnableNetworkIsolation,omitempty"`
	Modelarn interface{} `json:"ModelArn,omitempty"`
	Executionrolearn interface{} `json:"ExecutionRoleArn,omitempty"`
	Deploymentrecommendation interface{} `json:"DeploymentRecommendation,omitempty"`
}

// RecommendationJobInferenceBenchmark represents the RecommendationJobInferenceBenchmark schema from the OpenAPI specification
type RecommendationJobInferenceBenchmark struct {
	Metrics RecommendationMetrics `json:"Metrics,omitempty"` // The metrics of recommendations.
	Modelconfiguration ModelConfiguration `json:"ModelConfiguration"` // Defines the model configuration. Includes the specification name and environment parameters.
	Endpointconfiguration EndpointOutputConfiguration `json:"EndpointConfiguration,omitempty"` // The endpoint configuration made by Inference Recommender during a recommendation job.
	Endpointmetrics InferenceMetrics `json:"EndpointMetrics,omitempty"` // The metrics for an existing endpoint compared in an Inference Recommender job.
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Invocationendtime interface{} `json:"InvocationEndTime,omitempty"`
	Invocationstarttime interface{} `json:"InvocationStartTime,omitempty"`
}

// DescribeNotebookInstanceInput represents the DescribeNotebookInstanceInput schema from the OpenAPI specification
type DescribeNotebookInstanceInput struct {
	Notebookinstancename interface{} `json:"NotebookInstanceName"`
}

// TrainingSpecification represents the TrainingSpecification schema from the OpenAPI specification
type TrainingSpecification struct {
	Trainingimage interface{} `json:"TrainingImage"`
	Trainingimagedigest interface{} `json:"TrainingImageDigest,omitempty"`
	Metricdefinitions interface{} `json:"MetricDefinitions,omitempty"`
	Supportedhyperparameters interface{} `json:"SupportedHyperParameters,omitempty"`
	Supportedtraininginstancetypes interface{} `json:"SupportedTrainingInstanceTypes"`
	Supportedtuningjobobjectivemetrics interface{} `json:"SupportedTuningJobObjectiveMetrics,omitempty"`
	Supportsdistributedtraining interface{} `json:"SupportsDistributedTraining,omitempty"`
	Trainingchannels interface{} `json:"TrainingChannels"`
}

// RetryPipelineExecutionResponse represents the RetryPipelineExecutionResponse schema from the OpenAPI specification
type RetryPipelineExecutionResponse struct {
	Pipelineexecutionarn interface{} `json:"PipelineExecutionArn,omitempty"`
}

// UserContext represents the UserContext schema from the OpenAPI specification
type UserContext struct {
	Userprofilearn interface{} `json:"UserProfileArn,omitempty"`
	Userprofilename interface{} `json:"UserProfileName,omitempty"`
	Domainid interface{} `json:"DomainId,omitempty"`
	Iamidentity interface{} `json:"IamIdentity,omitempty"`
}

// MonitoringJsonDatasetFormat represents the MonitoringJsonDatasetFormat schema from the OpenAPI specification
type MonitoringJsonDatasetFormat struct {
	Line interface{} `json:"Line,omitempty"`
}

// RStudioServerProDomainSettings represents the RStudioServerProDomainSettings schema from the OpenAPI specification
type RStudioServerProDomainSettings struct {
	Defaultresourcespec ResourceSpec `json:"DefaultResourceSpec,omitempty"` // Specifies the ARN's of a SageMaker image and SageMaker image version, and the instance type that the version runs on.
	Domainexecutionrolearn interface{} `json:"DomainExecutionRoleArn"`
	Rstudioconnecturl interface{} `json:"RStudioConnectUrl,omitempty"`
	Rstudiopackagemanagerurl interface{} `json:"RStudioPackageManagerUrl,omitempty"`
}

// CreateStudioLifecycleConfigResponse represents the CreateStudioLifecycleConfigResponse schema from the OpenAPI specification
type CreateStudioLifecycleConfigResponse struct {
	Studiolifecycleconfigarn interface{} `json:"StudioLifecycleConfigArn,omitempty"`
}

// ContextSummary represents the ContextSummary schema from the OpenAPI specification
type ContextSummary struct {
	Source interface{} `json:"Source,omitempty"`
	Contextarn interface{} `json:"ContextArn,omitempty"`
	Contextname interface{} `json:"ContextName,omitempty"`
	Contexttype interface{} `json:"ContextType,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
}

// Workteam represents the Workteam schema from the OpenAPI specification
type Workteam struct {
	Workforcearn interface{} `json:"WorkforceArn,omitempty"`
	Workteamname interface{} `json:"WorkteamName"`
	Notificationconfiguration interface{} `json:"NotificationConfiguration,omitempty"`
	Createdate interface{} `json:"CreateDate,omitempty"`
	Description interface{} `json:"Description"`
	Lastupdateddate interface{} `json:"LastUpdatedDate,omitempty"`
	Workteamarn interface{} `json:"WorkteamArn"`
	Productlistingids interface{} `json:"ProductListingIds,omitempty"`
	Subdomain interface{} `json:"SubDomain,omitempty"`
	Memberdefinitions interface{} `json:"MemberDefinitions"`
}

// DescribeNotebookInstanceLifecycleConfigInput represents the DescribeNotebookInstanceLifecycleConfigInput schema from the OpenAPI specification
type DescribeNotebookInstanceLifecycleConfigInput struct {
	Notebookinstancelifecycleconfigname interface{} `json:"NotebookInstanceLifecycleConfigName"`
}

// ModelDataSource represents the ModelDataSource schema from the OpenAPI specification
type ModelDataSource struct {
	S3datasource interface{} `json:"S3DataSource"`
}

// ListTagsInput represents the ListTagsInput schema from the OpenAPI specification
type ListTagsInput struct {
	Resourcearn interface{} `json:"ResourceArn"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListHubsRequest represents the ListHubsRequest schema from the OpenAPI specification
type ListHubsRequest struct {
	Namecontains interface{} `json:"NameContains,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Lastmodifiedtimebefore interface{} `json:"LastModifiedTimeBefore,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Lastmodifiedtimeafter interface{} `json:"LastModifiedTimeAfter,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// DeviceSelectionConfig represents the DeviceSelectionConfig schema from the OpenAPI specification
type DeviceSelectionConfig struct {
	Devicenames interface{} `json:"DeviceNames,omitempty"`
	Devicesubsettype interface{} `json:"DeviceSubsetType"`
	Percentage interface{} `json:"Percentage,omitempty"`
	Devicenamecontains interface{} `json:"DeviceNameContains,omitempty"`
}

// CreateActionRequest represents the CreateActionRequest schema from the OpenAPI specification
type CreateActionRequest struct {
	Actiontype interface{} `json:"ActionType"`
	Description interface{} `json:"Description,omitempty"`
	Metadataproperties MetadataProperties `json:"MetadataProperties,omitempty"` // Metadata properties of the tracking entity, trial, or trial component.
	Properties interface{} `json:"Properties,omitempty"`
	Source interface{} `json:"Source"`
	Status interface{} `json:"Status,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Actionname interface{} `json:"ActionName"`
}

// VpcConfig represents the VpcConfig schema from the OpenAPI specification
type VpcConfig struct {
	Securitygroupids interface{} `json:"SecurityGroupIds"`
	Subnets interface{} `json:"Subnets"`
}

// RetryStrategy represents the RetryStrategy schema from the OpenAPI specification
type RetryStrategy struct {
	Maximumretryattempts interface{} `json:"MaximumRetryAttempts"`
}

// CustomImage represents the CustomImage schema from the OpenAPI specification
type CustomImage struct {
	Appimageconfigname interface{} `json:"AppImageConfigName"`
	Imagename interface{} `json:"ImageName"`
	Imageversionnumber interface{} `json:"ImageVersionNumber,omitempty"`
}

// PipelineSummary represents the PipelineSummary schema from the OpenAPI specification
type PipelineSummary struct {
	Pipelinedisplayname interface{} `json:"PipelineDisplayName,omitempty"`
	Pipelinename interface{} `json:"PipelineName,omitempty"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Lastexecutiontime interface{} `json:"LastExecutionTime,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Pipelinearn interface{} `json:"PipelineArn,omitempty"`
	Pipelinedescription interface{} `json:"PipelineDescription,omitempty"`
}

// ArtifactSummary represents the ArtifactSummary schema from the OpenAPI specification
type ArtifactSummary struct {
	Artifactname interface{} `json:"ArtifactName,omitempty"`
	Artifacttype interface{} `json:"ArtifactType,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Source interface{} `json:"Source,omitempty"`
	Artifactarn interface{} `json:"ArtifactArn,omitempty"`
}

// EdgeDeploymentPlanSummary represents the EdgeDeploymentPlanSummary schema from the OpenAPI specification
type EdgeDeploymentPlanSummary struct {
	Edgedeploymentfailed interface{} `json:"EdgeDeploymentFailed"`
	Edgedeploymentpending interface{} `json:"EdgeDeploymentPending"`
	Edgedeploymentplanarn interface{} `json:"EdgeDeploymentPlanArn"`
	Edgedeploymentplanname interface{} `json:"EdgeDeploymentPlanName"`
	Edgedeploymentsuccess interface{} `json:"EdgeDeploymentSuccess"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Devicefleetname interface{} `json:"DeviceFleetName"`
}

// CreateSpaceResponse represents the CreateSpaceResponse schema from the OpenAPI specification
type CreateSpaceResponse struct {
	Spacearn interface{} `json:"SpaceArn,omitempty"`
}

// FlowDefinitionOutputConfig represents the FlowDefinitionOutputConfig schema from the OpenAPI specification
type FlowDefinitionOutputConfig struct {
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	S3outputpath interface{} `json:"S3OutputPath"`
}

// DescribeFeatureMetadataRequest represents the DescribeFeatureMetadataRequest schema from the OpenAPI specification
type DescribeFeatureMetadataRequest struct {
	Featuregroupname interface{} `json:"FeatureGroupName"`
	Featurename interface{} `json:"FeatureName"`
}

// OnlineStoreSecurityConfig represents the OnlineStoreSecurityConfig schema from the OpenAPI specification
type OnlineStoreSecurityConfig struct {
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
}

// LabelingJobSnsDataSource represents the LabelingJobSnsDataSource schema from the OpenAPI specification
type LabelingJobSnsDataSource struct {
	Snstopicarn interface{} `json:"SnsTopicArn"`
}

// DeleteInferenceExperimentRequest represents the DeleteInferenceExperimentRequest schema from the OpenAPI specification
type DeleteInferenceExperimentRequest struct {
	Name interface{} `json:"Name"`
}

// DescribeProjectOutput represents the DescribeProjectOutput schema from the OpenAPI specification
type DescribeProjectOutput struct {
	Projectid interface{} `json:"ProjectId"`
	Createdby UserContext `json:"CreatedBy,omitempty"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	Projectdescription interface{} `json:"ProjectDescription,omitempty"`
	Projectname interface{} `json:"ProjectName"`
	Projectstatus interface{} `json:"ProjectStatus"`
	Servicecatalogprovisioningdetails interface{} `json:"ServiceCatalogProvisioningDetails"`
	Servicecatalogprovisionedproductdetails interface{} `json:"ServiceCatalogProvisionedProductDetails,omitempty"`
	Creationtime interface{} `json:"CreationTime"`
	Lastmodifiedby UserContext `json:"LastModifiedBy,omitempty"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Projectarn interface{} `json:"ProjectArn"`
}

// SpaceSettings represents the SpaceSettings schema from the OpenAPI specification
type SpaceSettings struct {
	Jupyterserverappsettings JupyterServerAppSettings `json:"JupyterServerAppSettings,omitempty"` // The JupyterServer app settings.
	Kernelgatewayappsettings KernelGatewayAppSettings `json:"KernelGatewayAppSettings,omitempty"` // The KernelGateway app settings.
}

// ModelRegisterSettings represents the ModelRegisterSettings schema from the OpenAPI specification
type ModelRegisterSettings struct {
	Crossaccountmodelregisterrolearn interface{} `json:"CrossAccountModelRegisterRoleArn,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// DescribeAutoMLJobResponse represents the DescribeAutoMLJobResponse schema from the OpenAPI specification
type DescribeAutoMLJobResponse struct {
	Problemtype interface{} `json:"ProblemType,omitempty"`
	Automljobartifacts interface{} `json:"AutoMLJobArtifacts,omitempty"`
	Outputdataconfig interface{} `json:"OutputDataConfig"`
	Rolearn interface{} `json:"RoleArn"`
	Automljobconfig interface{} `json:"AutoMLJobConfig,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime"`
	Automljobobjective interface{} `json:"AutoMLJobObjective,omitempty"`
	Automljobname interface{} `json:"AutoMLJobName"`
	Modeldeployconfig interface{} `json:"ModelDeployConfig,omitempty"`
	Automljobsecondarystatus interface{} `json:"AutoMLJobSecondaryStatus"`
	Endtime interface{} `json:"EndTime,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Resolvedattributes interface{} `json:"ResolvedAttributes,omitempty"`
	Automljobarn interface{} `json:"AutoMLJobArn"`
	Creationtime interface{} `json:"CreationTime"`
	Generatecandidatedefinitionsonly interface{} `json:"GenerateCandidateDefinitionsOnly,omitempty"`
	Modeldeployresult interface{} `json:"ModelDeployResult,omitempty"`
	Partialfailurereasons interface{} `json:"PartialFailureReasons,omitempty"`
	Bestcandidate interface{} `json:"BestCandidate,omitempty"`
	Inputdataconfig interface{} `json:"InputDataConfig"`
	Automljobstatus interface{} `json:"AutoMLJobStatus"`
}

// DeleteImageVersionResponse represents the DeleteImageVersionResponse schema from the OpenAPI specification
type DeleteImageVersionResponse struct {
}

// QueryLineageResponse represents the QueryLineageResponse schema from the OpenAPI specification
type QueryLineageResponse struct {
	Edges interface{} `json:"Edges,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Vertices interface{} `json:"Vertices,omitempty"`
}

// DeleteExperimentRequest represents the DeleteExperimentRequest schema from the OpenAPI specification
type DeleteExperimentRequest struct {
	Experimentname interface{} `json:"ExperimentName"`
}

// MetricDefinition represents the MetricDefinition schema from the OpenAPI specification
type MetricDefinition struct {
	Name interface{} `json:"Name"`
	Regex interface{} `json:"Regex"`
}

// MonitoringStoppingCondition represents the MonitoringStoppingCondition schema from the OpenAPI specification
type MonitoringStoppingCondition struct {
	Maxruntimeinseconds interface{} `json:"MaxRuntimeInSeconds"`
}

// DescribeAppImageConfigResponse represents the DescribeAppImageConfigResponse schema from the OpenAPI specification
type DescribeAppImageConfigResponse struct {
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Appimageconfigarn interface{} `json:"AppImageConfigArn,omitempty"`
	Appimageconfigname interface{} `json:"AppImageConfigName,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Kernelgatewayimageconfig interface{} `json:"KernelGatewayImageConfig,omitempty"`
}

// RealTimeInferenceConfig represents the RealTimeInferenceConfig schema from the OpenAPI specification
type RealTimeInferenceConfig struct {
	Instancetype interface{} `json:"InstanceType"`
	Instancecount interface{} `json:"InstanceCount"`
}

// DescribeWorkteamResponse represents the DescribeWorkteamResponse schema from the OpenAPI specification
type DescribeWorkteamResponse struct {
	Workteam interface{} `json:"Workteam"`
}

// DeleteModelPackageGroupPolicyInput represents the DeleteModelPackageGroupPolicyInput schema from the OpenAPI specification
type DeleteModelPackageGroupPolicyInput struct {
	Modelpackagegroupname interface{} `json:"ModelPackageGroupName"`
}

// DescribeModelCardResponse represents the DescribeModelCardResponse schema from the OpenAPI specification
type DescribeModelCardResponse struct {
	Securityconfig interface{} `json:"SecurityConfig,omitempty"`
	Content interface{} `json:"Content"`
	Modelcardarn interface{} `json:"ModelCardArn"`
	Modelcardstatus interface{} `json:"ModelCardStatus"`
	Modelcardversion interface{} `json:"ModelCardVersion"`
	Creationtime interface{} `json:"CreationTime"`
	Createdby UserContext `json:"CreatedBy"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	Lastmodifiedby UserContext `json:"LastModifiedBy,omitempty"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Modelcardname interface{} `json:"ModelCardName"`
	Modelcardprocessingstatus interface{} `json:"ModelCardProcessingStatus,omitempty"`
}

// ModelExplainabilityJobInput represents the ModelExplainabilityJobInput schema from the OpenAPI specification
type ModelExplainabilityJobInput struct {
	Batchtransforminput interface{} `json:"BatchTransformInput,omitempty"`
	Endpointinput EndpointInput `json:"EndpointInput,omitempty"` // Input object for the endpoint
}

// RenderingError represents the RenderingError schema from the OpenAPI specification
type RenderingError struct {
	Code interface{} `json:"Code"`
	Message interface{} `json:"Message"`
}

// SelectiveExecutionConfig represents the SelectiveExecutionConfig schema from the OpenAPI specification
type SelectiveExecutionConfig struct {
	Selectedsteps interface{} `json:"SelectedSteps"`
	Sourcepipelineexecutionarn interface{} `json:"SourcePipelineExecutionArn"`
}

// AutoMLAlgorithmConfig represents the AutoMLAlgorithmConfig schema from the OpenAPI specification
type AutoMLAlgorithmConfig struct {
	Automlalgorithms interface{} `json:"AutoMLAlgorithms"`
}

// UpdateTrialRequest represents the UpdateTrialRequest schema from the OpenAPI specification
type UpdateTrialRequest struct {
	Displayname interface{} `json:"DisplayName,omitempty"`
	Trialname interface{} `json:"TrialName"`
}

// TargetTrackingScalingPolicyConfiguration represents the TargetTrackingScalingPolicyConfiguration schema from the OpenAPI specification
type TargetTrackingScalingPolicyConfiguration struct {
	Metricspecification interface{} `json:"MetricSpecification,omitempty"`
	Targetvalue interface{} `json:"TargetValue,omitempty"`
}

// EnvironmentMap represents the EnvironmentMap schema from the OpenAPI specification
type EnvironmentMap struct {
}

// ContinuousParameterRangeSpecification represents the ContinuousParameterRangeSpecification schema from the OpenAPI specification
type ContinuousParameterRangeSpecification struct {
	Maxvalue interface{} `json:"MaxValue"`
	Minvalue interface{} `json:"MinValue"`
}

// ResourceSpec represents the ResourceSpec schema from the OpenAPI specification
type ResourceSpec struct {
	Instancetype interface{} `json:"InstanceType,omitempty"`
	Lifecycleconfigarn interface{} `json:"LifecycleConfigArn,omitempty"`
	Sagemakerimagearn interface{} `json:"SageMakerImageArn,omitempty"`
	Sagemakerimageversionarn interface{} `json:"SageMakerImageVersionArn,omitempty"`
}

// EdgeModelStat represents the EdgeModelStat schema from the OpenAPI specification
type EdgeModelStat struct {
	Modelname interface{} `json:"ModelName"`
	Modelversion interface{} `json:"ModelVersion"`
	Offlinedevicecount interface{} `json:"OfflineDeviceCount"`
	Samplingdevicecount interface{} `json:"SamplingDeviceCount"`
	Activedevicecount interface{} `json:"ActiveDeviceCount"`
	Connecteddevicecount interface{} `json:"ConnectedDeviceCount"`
}

// HumanLoopRequestSource represents the HumanLoopRequestSource schema from the OpenAPI specification
type HumanLoopRequestSource struct {
	Awsmanagedhumanlooprequestsource interface{} `json:"AwsManagedHumanLoopRequestSource"`
}

// UpdateArtifactRequest represents the UpdateArtifactRequest schema from the OpenAPI specification
type UpdateArtifactRequest struct {
	Artifactarn interface{} `json:"ArtifactArn"`
	Artifactname interface{} `json:"ArtifactName,omitempty"`
	Properties interface{} `json:"Properties,omitempty"`
	Propertiestoremove interface{} `json:"PropertiesToRemove,omitempty"`
}

// DescribeUserProfileRequest represents the DescribeUserProfileRequest schema from the OpenAPI specification
type DescribeUserProfileRequest struct {
	Userprofilename interface{} `json:"UserProfileName"`
	Domainid interface{} `json:"DomainId"`
}

// Autotune represents the Autotune schema from the OpenAPI specification
type Autotune struct {
	Mode interface{} `json:"Mode"`
}

// CreateInferenceExperimentRequest represents the CreateInferenceExperimentRequest schema from the OpenAPI specification
type CreateInferenceExperimentRequest struct {
	Shadowmodeconfig interface{} `json:"ShadowModeConfig"`
	Kmskey interface{} `json:"KmsKey,omitempty"`
	Modelvariants interface{} `json:"ModelVariants"`
	Tags interface{} `json:"Tags,omitempty"`
	TypeField interface{} `json:"Type"`
	Endpointname interface{} `json:"EndpointName"`
	Schedule interface{} `json:"Schedule,omitempty"`
	Datastorageconfig interface{} `json:"DataStorageConfig,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Name interface{} `json:"Name"`
	Rolearn interface{} `json:"RoleArn"`
}

// ListModelBiasJobDefinitionsRequest represents the ListModelBiasJobDefinitionsRequest schema from the OpenAPI specification
type ListModelBiasJobDefinitionsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Endpointname interface{} `json:"EndpointName,omitempty"`
}

// ModelExplainabilityAppSpecification represents the ModelExplainabilityAppSpecification schema from the OpenAPI specification
type ModelExplainabilityAppSpecification struct {
	Configuri interface{} `json:"ConfigUri"`
	Environment interface{} `json:"Environment,omitempty"`
	Imageuri interface{} `json:"ImageUri"`
}

// ListEdgeDeploymentPlansResponse represents the ListEdgeDeploymentPlansResponse schema from the OpenAPI specification
type ListEdgeDeploymentPlansResponse struct {
	Edgedeploymentplansummaries interface{} `json:"EdgeDeploymentPlanSummaries"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListMonitoringAlertsResponse represents the ListMonitoringAlertsResponse schema from the OpenAPI specification
type ListMonitoringAlertsResponse struct {
	Monitoringalertsummaries interface{} `json:"MonitoringAlertSummaries,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DescribeUserProfileResponse represents the DescribeUserProfileResponse schema from the OpenAPI specification
type DescribeUserProfileResponse struct {
	Homeefsfilesystemuid interface{} `json:"HomeEfsFileSystemUid,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Singlesignonuservalue interface{} `json:"SingleSignOnUserValue,omitempty"`
	Userprofilearn interface{} `json:"UserProfileArn,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Singlesignonuseridentifier interface{} `json:"SingleSignOnUserIdentifier,omitempty"`
	Userprofilename interface{} `json:"UserProfileName,omitempty"`
	Usersettings interface{} `json:"UserSettings,omitempty"`
	Domainid interface{} `json:"DomainId,omitempty"`
}

// InferenceMetrics represents the InferenceMetrics schema from the OpenAPI specification
type InferenceMetrics struct {
	Maxinvocations interface{} `json:"MaxInvocations"`
	Modellatency interface{} `json:"ModelLatency"`
}

// MonitoringClusterConfig represents the MonitoringClusterConfig schema from the OpenAPI specification
type MonitoringClusterConfig struct {
	Volumesizeingb interface{} `json:"VolumeSizeInGB"`
	Instancecount interface{} `json:"InstanceCount"`
	Instancetype interface{} `json:"InstanceType"`
	Volumekmskeyid interface{} `json:"VolumeKmsKeyId,omitempty"`
}

// InputConfig represents the InputConfig schema from the OpenAPI specification
type InputConfig struct {
	Datainputconfig interface{} `json:"DataInputConfig"`
	Framework interface{} `json:"Framework"`
	Frameworkversion interface{} `json:"FrameworkVersion,omitempty"`
	S3uri interface{} `json:"S3Uri"`
}

// LabelingJobStoppingConditions represents the LabelingJobStoppingConditions schema from the OpenAPI specification
type LabelingJobStoppingConditions struct {
	Maxpercentageofinputdatasetlabeled interface{} `json:"MaxPercentageOfInputDatasetLabeled,omitempty"`
	Maxhumanlabeledobjectcount interface{} `json:"MaxHumanLabeledObjectCount,omitempty"`
}

// SourceAlgorithmSpecification represents the SourceAlgorithmSpecification schema from the OpenAPI specification
type SourceAlgorithmSpecification struct {
	Sourcealgorithms interface{} `json:"SourceAlgorithms"`
}

// CreateLabelingJobRequest represents the CreateLabelingJobRequest schema from the OpenAPI specification
type CreateLabelingJobRequest struct {
	Outputconfig interface{} `json:"OutputConfig"`
	Labelattributename interface{} `json:"LabelAttributeName"`
	Labelcategoryconfigs3uri interface{} `json:"LabelCategoryConfigS3Uri,omitempty"`
	Labelingjobname interface{} `json:"LabelingJobName"`
	Labelingjobalgorithmsconfig interface{} `json:"LabelingJobAlgorithmsConfig,omitempty"`
	Rolearn interface{} `json:"RoleArn"`
	Inputconfig interface{} `json:"InputConfig"`
	Stoppingconditions interface{} `json:"StoppingConditions,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Humantaskconfig interface{} `json:"HumanTaskConfig"`
}

// PipelineDefinitionS3Location represents the PipelineDefinitionS3Location schema from the OpenAPI specification
type PipelineDefinitionS3Location struct {
	Bucket interface{} `json:"Bucket"`
	Objectkey interface{} `json:"ObjectKey"`
	Versionid interface{} `json:"VersionId,omitempty"`
}

// CreateModelExplainabilityJobDefinitionResponse represents the CreateModelExplainabilityJobDefinitionResponse schema from the OpenAPI specification
type CreateModelExplainabilityJobDefinitionResponse struct {
	Jobdefinitionarn interface{} `json:"JobDefinitionArn"`
}

// CodeRepositorySummary represents the CodeRepositorySummary schema from the OpenAPI specification
type CodeRepositorySummary struct {
	Lastmodifiedtime interface{} `json:"LastModifiedTime"`
	Coderepositoryarn interface{} `json:"CodeRepositoryArn"`
	Coderepositoryname interface{} `json:"CodeRepositoryName"`
	Creationtime interface{} `json:"CreationTime"`
	Gitconfig interface{} `json:"GitConfig,omitempty"`
}

// DescribeImageResponse represents the DescribeImageResponse schema from the OpenAPI specification
type DescribeImageResponse struct {
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Displayname interface{} `json:"DisplayName,omitempty"`
	Imagearn interface{} `json:"ImageArn,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Imagename interface{} `json:"ImageName,omitempty"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Imagestatus interface{} `json:"ImageStatus,omitempty"`
}

// ListHumanTaskUisRequest represents the ListHumanTaskUisRequest schema from the OpenAPI specification
type ListHumanTaskUisRequest struct {
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// SendPipelineExecutionStepSuccessRequest represents the SendPipelineExecutionStepSuccessRequest schema from the OpenAPI specification
type SendPipelineExecutionStepSuccessRequest struct {
	Clientrequesttoken interface{} `json:"ClientRequestToken,omitempty"`
	Outputparameters interface{} `json:"OutputParameters,omitempty"`
	Callbacktoken interface{} `json:"CallbackToken"`
}

// MonitoringStatisticsResource represents the MonitoringStatisticsResource schema from the OpenAPI specification
type MonitoringStatisticsResource struct {
	S3uri interface{} `json:"S3Uri,omitempty"`
}

// InferenceExperimentDataStorageConfig represents the InferenceExperimentDataStorageConfig schema from the OpenAPI specification
type InferenceExperimentDataStorageConfig struct {
	Destination interface{} `json:"Destination"`
	Kmskey interface{} `json:"KmsKey,omitempty"`
	Contenttype CaptureContentTypeHeader `json:"ContentType,omitempty"` // Configuration specifying how to treat different headers. If no headers are specified SageMaker will by default base64 encode when capturing the data.
}

// ListEndpointConfigsInput represents the ListEndpointConfigsInput schema from the OpenAPI specification
type ListEndpointConfigsInput struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
}

// AutoMLCandidateStep represents the AutoMLCandidateStep schema from the OpenAPI specification
type AutoMLCandidateStep struct {
	Candidatestepname interface{} `json:"CandidateStepName"`
	Candidatesteptype interface{} `json:"CandidateStepType"`
	Candidatesteparn interface{} `json:"CandidateStepArn"`
}

// UiTemplateInfo represents the UiTemplateInfo schema from the OpenAPI specification
type UiTemplateInfo struct {
	Url interface{} `json:"Url,omitempty"`
	Contentsha256 interface{} `json:"ContentSha256,omitempty"`
}

// ActionSummary represents the ActionSummary schema from the OpenAPI specification
type ActionSummary struct {
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Source interface{} `json:"Source,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Actionarn interface{} `json:"ActionArn,omitempty"`
	Actionname interface{} `json:"ActionName,omitempty"`
	Actiontype interface{} `json:"ActionType,omitempty"`
}

// PublicWorkforceTaskPrice represents the PublicWorkforceTaskPrice schema from the OpenAPI specification
type PublicWorkforceTaskPrice struct {
	Amountinusd interface{} `json:"AmountInUsd,omitempty"`
}

// ModelMetrics represents the ModelMetrics schema from the OpenAPI specification
type ModelMetrics struct {
	Modeldataquality interface{} `json:"ModelDataQuality,omitempty"`
	Modelquality interface{} `json:"ModelQuality,omitempty"`
	Bias interface{} `json:"Bias,omitempty"`
	Explainability interface{} `json:"Explainability,omitempty"`
}

// MonitoringResources represents the MonitoringResources schema from the OpenAPI specification
type MonitoringResources struct {
	Clusterconfig interface{} `json:"ClusterConfig"`
}

// UpdateMonitoringScheduleRequest represents the UpdateMonitoringScheduleRequest schema from the OpenAPI specification
type UpdateMonitoringScheduleRequest struct {
	Monitoringscheduleconfig interface{} `json:"MonitoringScheduleConfig"`
	Monitoringschedulename interface{} `json:"MonitoringScheduleName"`
}

// StopTrainingJobRequest represents the StopTrainingJobRequest schema from the OpenAPI specification
type StopTrainingJobRequest struct {
	Trainingjobname interface{} `json:"TrainingJobName"`
}

// StopPipelineExecutionResponse represents the StopPipelineExecutionResponse schema from the OpenAPI specification
type StopPipelineExecutionResponse struct {
	Pipelineexecutionarn interface{} `json:"PipelineExecutionArn,omitempty"`
}

// TrainingJob represents the TrainingJob schema from the OpenAPI specification
type TrainingJob struct {
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Experimentconfig ExperimentConfig `json:"ExperimentConfig,omitempty"` // <p>Associates a SageMaker job as a trial component with an experiment and trial. Specified when you call the following APIs:</p> <ul> <li> <p> <a href="https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateProcessingJob.html">CreateProcessingJob</a> </p> </li> <li> <p> <a href="https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTrainingJob.html">CreateTrainingJob</a> </p> </li> <li> <p> <a href="https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTransformJob.html">CreateTransformJob</a> </p> </li> </ul>
	Billabletimeinseconds interface{} `json:"BillableTimeInSeconds,omitempty"`
	Trainingjobstatus interface{} `json:"TrainingJobStatus,omitempty"`
	Algorithmspecification interface{} `json:"AlgorithmSpecification,omitempty"`
	Inputdataconfig interface{} `json:"InputDataConfig,omitempty"`
	Profilerconfig ProfilerConfig `json:"ProfilerConfig,omitempty"` // Configuration information for Amazon SageMaker Debugger system monitoring, framework profiling, and storage paths.
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Retrystrategy interface{} `json:"RetryStrategy,omitempty"`
	Debughookconfig DebugHookConfig `json:"DebugHookConfig,omitempty"` // Configuration information for the Amazon SageMaker Debugger hook parameters, metric and tensor collections, and storage paths. To learn more about how to configure the <code>DebugHookConfig</code> parameter, see <a href="https://docs.aws.amazon.com/sagemaker/latest/dg/debugger-createtrainingjob-api.html">Use the SageMaker and Debugger Configuration API Operations to Create, Update, and Debug Your Training Job</a>.
	Enablemanagedspottraining interface{} `json:"EnableManagedSpotTraining,omitempty"`
	Finalmetricdatalist interface{} `json:"FinalMetricDataList,omitempty"`
	Hyperparameters interface{} `json:"HyperParameters,omitempty"`
	Checkpointconfig CheckpointConfig `json:"CheckpointConfig,omitempty"` // Contains information about the output location for managed spot training checkpoint data.
	Trainingendtime interface{} `json:"TrainingEndTime,omitempty"`
	Enableintercontainertrafficencryption interface{} `json:"EnableInterContainerTrafficEncryption,omitempty"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Automljobarn interface{} `json:"AutoMLJobArn,omitempty"`
	Debugruleconfigurations interface{} `json:"DebugRuleConfigurations,omitempty"`
	Tuningjobarn interface{} `json:"TuningJobArn,omitempty"`
	Environment interface{} `json:"Environment,omitempty"`
	Vpcconfig interface{} `json:"VpcConfig,omitempty"`
	Modelartifacts interface{} `json:"ModelArtifacts,omitempty"`
	Resourceconfig interface{} `json:"ResourceConfig,omitempty"`
	Secondarystatus interface{} `json:"SecondaryStatus,omitempty"`
	Enablenetworkisolation interface{} `json:"EnableNetworkIsolation,omitempty"`
	Secondarystatustransitions interface{} `json:"SecondaryStatusTransitions,omitempty"`
	Trainingstarttime interface{} `json:"TrainingStartTime,omitempty"`
	Debugruleevaluationstatuses interface{} `json:"DebugRuleEvaluationStatuses,omitempty"`
	Stoppingcondition interface{} `json:"StoppingCondition,omitempty"`
	Trainingjobname interface{} `json:"TrainingJobName,omitempty"`
	Trainingjobarn interface{} `json:"TrainingJobArn,omitempty"`
	Outputdataconfig interface{} `json:"OutputDataConfig,omitempty"`
	Trainingtimeinseconds interface{} `json:"TrainingTimeInSeconds,omitempty"`
	Tensorboardoutputconfig TensorBoardOutputConfig `json:"TensorBoardOutputConfig,omitempty"` // Configuration of storage locations for the Amazon SageMaker Debugger TensorBoard output data.
	Labelingjobarn interface{} `json:"LabelingJobArn,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// LabelingJobOutputConfig represents the LabelingJobOutputConfig schema from the OpenAPI specification
type LabelingJobOutputConfig struct {
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	S3outputpath interface{} `json:"S3OutputPath"`
	Snstopicarn interface{} `json:"SnsTopicArn,omitempty"`
}

// LabelingJobOutput represents the LabelingJobOutput schema from the OpenAPI specification
type LabelingJobOutput struct {
	Finalactivelearningmodelarn interface{} `json:"FinalActiveLearningModelArn,omitempty"`
	Outputdatasets3uri interface{} `json:"OutputDatasetS3Uri"`
}

// CreateHumanTaskUiResponse represents the CreateHumanTaskUiResponse schema from the OpenAPI specification
type CreateHumanTaskUiResponse struct {
	Humantaskuiarn interface{} `json:"HumanTaskUiArn"`
}

// ListPipelineExecutionsResponse represents the ListPipelineExecutionsResponse schema from the OpenAPI specification
type ListPipelineExecutionsResponse struct {
	Pipelineexecutionsummaries interface{} `json:"PipelineExecutionSummaries,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListPipelinesRequest represents the ListPipelinesRequest schema from the OpenAPI specification
type ListPipelinesRequest struct {
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Createdafter interface{} `json:"CreatedAfter,omitempty"`
	Createdbefore interface{} `json:"CreatedBefore,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Pipelinenameprefix interface{} `json:"PipelineNamePrefix,omitempty"`
}

// CreateAutoMLJobV2Response represents the CreateAutoMLJobV2Response schema from the OpenAPI specification
type CreateAutoMLJobV2Response struct {
	Automljobarn interface{} `json:"AutoMLJobArn"`
}

// DescribeTrialComponentResponse represents the DescribeTrialComponentResponse schema from the OpenAPI specification
type DescribeTrialComponentResponse struct {
	Source interface{} `json:"Source,omitempty"`
	Createdby interface{} `json:"CreatedBy,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Inputartifacts interface{} `json:"InputArtifacts,omitempty"`
	Metadataproperties MetadataProperties `json:"MetadataProperties,omitempty"` // Metadata properties of the tracking entity, trial, or trial component.
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Metrics interface{} `json:"Metrics,omitempty"`
	Displayname interface{} `json:"DisplayName,omitempty"`
	Lineagegrouparn interface{} `json:"LineageGroupArn,omitempty"`
	Endtime interface{} `json:"EndTime,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Lastmodifiedby interface{} `json:"LastModifiedBy,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
	Outputartifacts interface{} `json:"OutputArtifacts,omitempty"`
	Parameters interface{} `json:"Parameters,omitempty"`
	Sources interface{} `json:"Sources,omitempty"`
	Trialcomponentname interface{} `json:"TrialComponentName,omitempty"`
	Trialcomponentarn interface{} `json:"TrialComponentArn,omitempty"`
}

// SearchResponse represents the SearchResponse schema from the OpenAPI specification
type SearchResponse struct {
	Results interface{} `json:"Results,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// EndpointOutputConfiguration represents the EndpointOutputConfiguration schema from the OpenAPI specification
type EndpointOutputConfiguration struct {
	Initialinstancecount interface{} `json:"InitialInstanceCount,omitempty"`
	Instancetype interface{} `json:"InstanceType,omitempty"`
	Serverlessconfig ProductionVariantServerlessConfig `json:"ServerlessConfig,omitempty"` // Specifies the serverless configuration for an endpoint variant.
	Variantname interface{} `json:"VariantName"`
	Endpointname interface{} `json:"EndpointName"`
}

// ListTrainingJobsForHyperParameterTuningJobResponse represents the ListTrainingJobsForHyperParameterTuningJobResponse schema from the OpenAPI specification
type ListTrainingJobsForHyperParameterTuningJobResponse struct {
	Trainingjobsummaries interface{} `json:"TrainingJobSummaries"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// NotebookInstanceLifecycleConfigSummary represents the NotebookInstanceLifecycleConfigSummary schema from the OpenAPI specification
type NotebookInstanceLifecycleConfigSummary struct {
	Notebookinstancelifecycleconfigarn interface{} `json:"NotebookInstanceLifecycleConfigArn"`
	Notebookinstancelifecycleconfigname interface{} `json:"NotebookInstanceLifecycleConfigName"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
}

// ModelCardExportArtifacts represents the ModelCardExportArtifacts schema from the OpenAPI specification
type ModelCardExportArtifacts struct {
	S3exportartifacts interface{} `json:"S3ExportArtifacts"`
}

// ModelPackageGroupSummary represents the ModelPackageGroupSummary schema from the OpenAPI specification
type ModelPackageGroupSummary struct {
	Modelpackagegroupname interface{} `json:"ModelPackageGroupName"`
	Modelpackagegroupstatus interface{} `json:"ModelPackageGroupStatus"`
	Creationtime interface{} `json:"CreationTime"`
	Modelpackagegrouparn interface{} `json:"ModelPackageGroupArn"`
	Modelpackagegroupdescription interface{} `json:"ModelPackageGroupDescription,omitempty"`
}

// BatchDescribeModelPackageError represents the BatchDescribeModelPackageError schema from the OpenAPI specification
type BatchDescribeModelPackageError struct {
	Errorcode interface{} `json:"ErrorCode"`
	Errorresponse interface{} `json:"ErrorResponse"`
}

// PendingDeploymentSummary represents the PendingDeploymentSummary schema from the OpenAPI specification
type PendingDeploymentSummary struct {
	Endpointconfigname interface{} `json:"EndpointConfigName"`
	Productionvariants interface{} `json:"ProductionVariants,omitempty"`
	Shadowproductionvariants interface{} `json:"ShadowProductionVariants,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
}

// TabularJobConfig represents the TabularJobConfig schema from the OpenAPI specification
type TabularJobConfig struct {
	Targetattributename interface{} `json:"TargetAttributeName"`
	Candidategenerationconfig interface{} `json:"CandidateGenerationConfig,omitempty"`
	Completioncriteria AutoMLJobCompletionCriteria `json:"CompletionCriteria,omitempty"` // How long a job is allowed to run, or how many candidates a job is allowed to generate.
	Featurespecifications3uri interface{} `json:"FeatureSpecificationS3Uri,omitempty"`
	Generatecandidatedefinitionsonly interface{} `json:"GenerateCandidateDefinitionsOnly,omitempty"`
	Mode interface{} `json:"Mode,omitempty"`
	Problemtype interface{} `json:"ProblemType,omitempty"`
	Sampleweightattributename interface{} `json:"SampleWeightAttributeName,omitempty"`
}

// DescribeEndpointOutput represents the DescribeEndpointOutput schema from the OpenAPI specification
type DescribeEndpointOutput struct {
	Datacaptureconfig DataCaptureConfigSummary `json:"DataCaptureConfig,omitempty"` // The currently active data capture configuration used by your Endpoint.
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Pendingdeploymentsummary interface{} `json:"PendingDeploymentSummary,omitempty"`
	Endpointarn interface{} `json:"EndpointArn"`
	Explainerconfig interface{} `json:"ExplainerConfig,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime"`
	Shadowproductionvariants interface{} `json:"ShadowProductionVariants,omitempty"`
	Endpointconfigname interface{} `json:"EndpointConfigName"`
	Endpointname interface{} `json:"EndpointName"`
	Productionvariants interface{} `json:"ProductionVariants,omitempty"`
	Endpointstatus interface{} `json:"EndpointStatus"`
	Lastdeploymentconfig interface{} `json:"LastDeploymentConfig,omitempty"`
	Asyncinferenceconfig interface{} `json:"AsyncInferenceConfig,omitempty"`
	Creationtime interface{} `json:"CreationTime"`
}

// StopInferenceRecommendationsJobRequest represents the StopInferenceRecommendationsJobRequest schema from the OpenAPI specification
type StopInferenceRecommendationsJobRequest struct {
	Jobname interface{} `json:"JobName"`
}

// BatchDescribeModelPackageErrorMap represents the BatchDescribeModelPackageErrorMap schema from the OpenAPI specification
type BatchDescribeModelPackageErrorMap struct {
}

// CreateProcessingJobRequest represents the CreateProcessingJobRequest schema from the OpenAPI specification
type CreateProcessingJobRequest struct {
	Appspecification interface{} `json:"AppSpecification"`
	Processingresources interface{} `json:"ProcessingResources"`
	Tags interface{} `json:"Tags,omitempty"`
	Environment interface{} `json:"Environment,omitempty"`
	Processingoutputconfig interface{} `json:"ProcessingOutputConfig,omitempty"`
	Stoppingcondition interface{} `json:"StoppingCondition,omitempty"`
	Experimentconfig ExperimentConfig `json:"ExperimentConfig,omitempty"` // <p>Associates a SageMaker job as a trial component with an experiment and trial. Specified when you call the following APIs:</p> <ul> <li> <p> <a href="https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateProcessingJob.html">CreateProcessingJob</a> </p> </li> <li> <p> <a href="https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTrainingJob.html">CreateTrainingJob</a> </p> </li> <li> <p> <a href="https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTransformJob.html">CreateTransformJob</a> </p> </li> </ul>
	Networkconfig interface{} `json:"NetworkConfig,omitempty"`
	Processinginputs interface{} `json:"ProcessingInputs,omitempty"`
	Processingjobname interface{} `json:"ProcessingJobName"`
	Rolearn interface{} `json:"RoleArn"`
}

// ListLabelingJobsResponse represents the ListLabelingJobsResponse schema from the OpenAPI specification
type ListLabelingJobsResponse struct {
	Labelingjobsummarylist interface{} `json:"LabelingJobSummaryList,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CreateNotebookInstanceLifecycleConfigInput represents the CreateNotebookInstanceLifecycleConfigInput schema from the OpenAPI specification
type CreateNotebookInstanceLifecycleConfigInput struct {
	Onstart interface{} `json:"OnStart,omitempty"`
	Notebookinstancelifecycleconfigname interface{} `json:"NotebookInstanceLifecycleConfigName"`
	Oncreate interface{} `json:"OnCreate,omitempty"`
}

// GetLineageGroupPolicyRequest represents the GetLineageGroupPolicyRequest schema from the OpenAPI specification
type GetLineageGroupPolicyRequest struct {
	Lineagegroupname interface{} `json:"LineageGroupName"`
}

// ListTrialsRequest represents the ListTrialsRequest schema from the OpenAPI specification
type ListTrialsRequest struct {
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Trialcomponentname interface{} `json:"TrialComponentName,omitempty"`
	Createdafter interface{} `json:"CreatedAfter,omitempty"`
	Createdbefore interface{} `json:"CreatedBefore,omitempty"`
	Experimentname interface{} `json:"ExperimentName,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DriftCheckModelQuality represents the DriftCheckModelQuality schema from the OpenAPI specification
type DriftCheckModelQuality struct {
	Constraints interface{} `json:"Constraints,omitempty"`
	Statistics interface{} `json:"Statistics,omitempty"`
}

// LabelingJobForWorkteamSummary represents the LabelingJobForWorkteamSummary schema from the OpenAPI specification
type LabelingJobForWorkteamSummary struct {
	Jobreferencecode interface{} `json:"JobReferenceCode"`
	Labelcounters interface{} `json:"LabelCounters,omitempty"`
	Labelingjobname interface{} `json:"LabelingJobName,omitempty"`
	Numberofhumanworkersperdataobject interface{} `json:"NumberOfHumanWorkersPerDataObject,omitempty"`
	Workrequesteraccountid interface{} `json:"WorkRequesterAccountId"`
	Creationtime interface{} `json:"CreationTime"`
}

// CreatePipelineRequest represents the CreatePipelineRequest schema from the OpenAPI specification
type CreatePipelineRequest struct {
	Pipelinedefinition interface{} `json:"PipelineDefinition,omitempty"`
	Pipelinename interface{} `json:"PipelineName"`
	Rolearn interface{} `json:"RoleArn"`
	Clientrequesttoken interface{} `json:"ClientRequestToken"`
	Parallelismconfiguration interface{} `json:"ParallelismConfiguration,omitempty"`
	Pipelinedisplayname interface{} `json:"PipelineDisplayName,omitempty"`
	Pipelinedefinitions3location interface{} `json:"PipelineDefinitionS3Location,omitempty"`
	Pipelinedescription interface{} `json:"PipelineDescription,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// DeleteHumanTaskUiRequest represents the DeleteHumanTaskUiRequest schema from the OpenAPI specification
type DeleteHumanTaskUiRequest struct {
	Humantaskuiname interface{} `json:"HumanTaskUiName"`
}

// DeleteModelInput represents the DeleteModelInput schema from the OpenAPI specification
type DeleteModelInput struct {
	Modelname interface{} `json:"ModelName"`
}

// Endpoint represents the Endpoint schema from the OpenAPI specification
type Endpoint struct {
	Monitoringschedules interface{} `json:"MonitoringSchedules,omitempty"`
	Shadowproductionvariants interface{} `json:"ShadowProductionVariants,omitempty"`
	Datacaptureconfig DataCaptureConfigSummary `json:"DataCaptureConfig,omitempty"` // The currently active data capture configuration used by your Endpoint.
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Productionvariants interface{} `json:"ProductionVariants,omitempty"`
	Endpointconfigname interface{} `json:"EndpointConfigName"`
	Endpointname interface{} `json:"EndpointName"`
	Endpointstatus interface{} `json:"EndpointStatus"`
	Tags interface{} `json:"Tags,omitempty"`
	Endpointarn interface{} `json:"EndpointArn"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime"`
	Creationtime interface{} `json:"CreationTime"`
}

// DeleteFlowDefinitionResponse represents the DeleteFlowDefinitionResponse schema from the OpenAPI specification
type DeleteFlowDefinitionResponse struct {
}

// ListTransformJobsRequest represents the ListTransformJobsRequest schema from the OpenAPI specification
type ListTransformJobsRequest struct {
	Lastmodifiedtimeafter interface{} `json:"LastModifiedTimeAfter,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Lastmodifiedtimebefore interface{} `json:"LastModifiedTimeBefore,omitempty"`
	Statusequals interface{} `json:"StatusEquals,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
}

// TransformS3DataSource represents the TransformS3DataSource schema from the OpenAPI specification
type TransformS3DataSource struct {
	S3datatype interface{} `json:"S3DataType"`
	S3uri interface{} `json:"S3Uri"`
}

// UpdateSpaceResponse represents the UpdateSpaceResponse schema from the OpenAPI specification
type UpdateSpaceResponse struct {
	Spacearn interface{} `json:"SpaceArn,omitempty"`
}

// MonitoringOutput represents the MonitoringOutput schema from the OpenAPI specification
type MonitoringOutput struct {
	S3output interface{} `json:"S3Output"`
}

// Phase represents the Phase schema from the OpenAPI specification
type Phase struct {
	Initialnumberofusers interface{} `json:"InitialNumberOfUsers,omitempty"`
	Spawnrate interface{} `json:"SpawnRate,omitempty"`
	Durationinseconds interface{} `json:"DurationInSeconds,omitempty"`
}

// ListFlowDefinitionsRequest represents the ListFlowDefinitionsRequest schema from the OpenAPI specification
type ListFlowDefinitionsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
}

// CreateEndpointInput represents the CreateEndpointInput schema from the OpenAPI specification
type CreateEndpointInput struct {
	Tags interface{} `json:"Tags,omitempty"`
	Deploymentconfig DeploymentConfig `json:"DeploymentConfig,omitempty"` // The deployment configuration for an endpoint, which contains the desired deployment strategy and rollback configurations.
	Endpointconfigname interface{} `json:"EndpointConfigName"`
	Endpointname interface{} `json:"EndpointName"`
}

// DescribeLineageGroupResponse represents the DescribeLineageGroupResponse schema from the OpenAPI specification
type DescribeLineageGroupResponse struct {
	Lastmodifiedby UserContext `json:"LastModifiedBy,omitempty"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Lineagegrouparn interface{} `json:"LineageGroupArn,omitempty"`
	Lineagegroupname interface{} `json:"LineageGroupName,omitempty"`
	Createdby UserContext `json:"CreatedBy,omitempty"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Displayname interface{} `json:"DisplayName,omitempty"`
}

// QueryLineageRequest represents the QueryLineageRequest schema from the OpenAPI specification
type QueryLineageRequest struct {
	Direction interface{} `json:"Direction,omitempty"`
	Filters interface{} `json:"Filters,omitempty"`
	Includeedges interface{} `json:"IncludeEdges,omitempty"`
	Maxdepth interface{} `json:"MaxDepth,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Startarns interface{} `json:"StartArns,omitempty"`
}

// ModelPackageGroup represents the ModelPackageGroup schema from the OpenAPI specification
type ModelPackageGroup struct {
	Modelpackagegroupname interface{} `json:"ModelPackageGroupName,omitempty"`
	Modelpackagegroupstatus interface{} `json:"ModelPackageGroupStatus,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Createdby UserContext `json:"CreatedBy,omitempty"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Modelpackagegrouparn interface{} `json:"ModelPackageGroupArn,omitempty"`
	Modelpackagegroupdescription interface{} `json:"ModelPackageGroupDescription,omitempty"`
}

// UpdateWorkteamResponse represents the UpdateWorkteamResponse schema from the OpenAPI specification
type UpdateWorkteamResponse struct {
	Workteam interface{} `json:"Workteam"`
}

// AgentVersion represents the AgentVersion schema from the OpenAPI specification
type AgentVersion struct {
	Agentcount interface{} `json:"AgentCount"`
	Version interface{} `json:"Version"`
}

// TrainingRepositoryAuthConfig represents the TrainingRepositoryAuthConfig schema from the OpenAPI specification
type TrainingRepositoryAuthConfig struct {
	Trainingrepositorycredentialsproviderarn interface{} `json:"TrainingRepositoryCredentialsProviderArn"`
}

// CreateModelQualityJobDefinitionRequest represents the CreateModelQualityJobDefinitionRequest schema from the OpenAPI specification
type CreateModelQualityJobDefinitionRequest struct {
	Rolearn interface{} `json:"RoleArn"`
	Tags interface{} `json:"Tags,omitempty"`
	Jobresources MonitoringResources `json:"JobResources"` // Identifies the resources to deploy for a monitoring job.
	Modelqualitybaselineconfig interface{} `json:"ModelQualityBaselineConfig,omitempty"`
	Modelqualityjobinput interface{} `json:"ModelQualityJobInput"`
	Networkconfig interface{} `json:"NetworkConfig,omitempty"`
	Stoppingcondition MonitoringStoppingCondition `json:"StoppingCondition,omitempty"` // A time limit for how long the monitoring job is allowed to run before stopping.
	Jobdefinitionname interface{} `json:"JobDefinitionName"`
	Modelqualityappspecification interface{} `json:"ModelQualityAppSpecification"`
	Modelqualityjoboutputconfig MonitoringOutputConfig `json:"ModelQualityJobOutputConfig"` // The output configuration for monitoring jobs.
}

// DescribeAutoMLJobV2Response represents the DescribeAutoMLJobV2Response schema from the OpenAPI specification
type DescribeAutoMLJobV2Response struct {
	Lastmodifiedtime interface{} `json:"LastModifiedTime"`
	Endtime interface{} `json:"EndTime,omitempty"`
	Modeldeployresult interface{} `json:"ModelDeployResult,omitempty"`
	Modeldeployconfig interface{} `json:"ModelDeployConfig,omitempty"`
	Resolvedattributes interface{} `json:"ResolvedAttributes,omitempty"`
	Automljobsecondarystatus interface{} `json:"AutoMLJobSecondaryStatus"`
	Rolearn interface{} `json:"RoleArn"`
	Automljobinputdataconfig interface{} `json:"AutoMLJobInputDataConfig"`
	Automlproblemtypeconfig interface{} `json:"AutoMLProblemTypeConfig,omitempty"`
	Automljobstatus interface{} `json:"AutoMLJobStatus"`
	Datasplitconfig interface{} `json:"DataSplitConfig,omitempty"`
	Securityconfig interface{} `json:"SecurityConfig,omitempty"`
	Partialfailurereasons interface{} `json:"PartialFailureReasons,omitempty"`
	Bestcandidate interface{} `json:"BestCandidate,omitempty"`
	Automljobartifacts AutoMLJobArtifacts `json:"AutoMLJobArtifacts,omitempty"` // The artifacts that are generated during an AutoML job.
	Automljobname interface{} `json:"AutoMLJobName"`
	Automljobobjective interface{} `json:"AutoMLJobObjective,omitempty"`
	Creationtime interface{} `json:"CreationTime"`
	Outputdataconfig interface{} `json:"OutputDataConfig"`
	Automljobarn interface{} `json:"AutoMLJobArn"`
	Automlproblemtypeconfigname interface{} `json:"AutoMLProblemTypeConfigName,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
}

// DescribePipelineDefinitionForExecutionRequest represents the DescribePipelineDefinitionForExecutionRequest schema from the OpenAPI specification
type DescribePipelineDefinitionForExecutionRequest struct {
	Pipelineexecutionarn interface{} `json:"PipelineExecutionArn"`
}

// DeleteActionRequest represents the DeleteActionRequest schema from the OpenAPI specification
type DeleteActionRequest struct {
	Actionname interface{} `json:"ActionName"`
}

// HyperParameterTuningJobCompletionDetails represents the HyperParameterTuningJobCompletionDetails schema from the OpenAPI specification
type HyperParameterTuningJobCompletionDetails struct {
	Convergencedetectedtime interface{} `json:"ConvergenceDetectedTime,omitempty"`
	Numberoftrainingjobsobjectivenotimproving interface{} `json:"NumberOfTrainingJobsObjectiveNotImproving,omitempty"`
}

// DescribeInferenceRecommendationsJobRequest represents the DescribeInferenceRecommendationsJobRequest schema from the OpenAPI specification
type DescribeInferenceRecommendationsJobRequest struct {
	Jobname interface{} `json:"JobName"`
}

// ListModelCardExportJobsRequest represents the ListModelCardExportJobsRequest schema from the OpenAPI specification
type ListModelCardExportJobsRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Statusequals interface{} `json:"StatusEquals,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Modelcardexportjobnamecontains interface{} `json:"ModelCardExportJobNameContains,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Modelcardname interface{} `json:"ModelCardName"`
	Modelcardversion interface{} `json:"ModelCardVersion,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
}

// UpdateMonitoringAlertResponse represents the UpdateMonitoringAlertResponse schema from the OpenAPI specification
type UpdateMonitoringAlertResponse struct {
	Monitoringalertname interface{} `json:"MonitoringAlertName,omitempty"`
	Monitoringschedulearn interface{} `json:"MonitoringScheduleArn"`
}

// CheckpointConfig represents the CheckpointConfig schema from the OpenAPI specification
type CheckpointConfig struct {
	Localpath interface{} `json:"LocalPath,omitempty"`
	S3uri interface{} `json:"S3Uri"`
}

// ListTrainingJobsRequest represents the ListTrainingJobsRequest schema from the OpenAPI specification
type ListTrainingJobsRequest struct {
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Warmpoolstatusequals interface{} `json:"WarmPoolStatusEquals,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Lastmodifiedtimeafter interface{} `json:"LastModifiedTimeAfter,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Lastmodifiedtimebefore interface{} `json:"LastModifiedTimeBefore,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Statusequals interface{} `json:"StatusEquals,omitempty"`
}

// CreateAlgorithmInput represents the CreateAlgorithmInput schema from the OpenAPI specification
type CreateAlgorithmInput struct {
	Algorithmname interface{} `json:"AlgorithmName"`
	Certifyformarketplace interface{} `json:"CertifyForMarketplace,omitempty"`
	Inferencespecification interface{} `json:"InferenceSpecification,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Trainingspecification interface{} `json:"TrainingSpecification"`
	Validationspecification interface{} `json:"ValidationSpecification,omitempty"`
	Algorithmdescription interface{} `json:"AlgorithmDescription,omitempty"`
}

// Filter represents the Filter schema from the OpenAPI specification
type Filter struct {
	Name interface{} `json:"Name"`
	Operator interface{} `json:"Operator,omitempty"`
	Value interface{} `json:"Value,omitempty"`
}

// AutoMLOutputDataConfig represents the AutoMLOutputDataConfig schema from the OpenAPI specification
type AutoMLOutputDataConfig struct {
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	S3outputpath interface{} `json:"S3OutputPath"`
}

// DeleteAppImageConfigRequest represents the DeleteAppImageConfigRequest schema from the OpenAPI specification
type DeleteAppImageConfigRequest struct {
	Appimageconfigname interface{} `json:"AppImageConfigName"`
}

// MetricsSource represents the MetricsSource schema from the OpenAPI specification
type MetricsSource struct {
	S3uri interface{} `json:"S3Uri"`
	Contentdigest interface{} `json:"ContentDigest,omitempty"`
	Contenttype interface{} `json:"ContentType"`
}

// ListCandidatesForAutoMLJobResponse represents the ListCandidatesForAutoMLJobResponse schema from the OpenAPI specification
type ListCandidatesForAutoMLJobResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Candidates interface{} `json:"Candidates"`
}

// ClarifyInferenceConfig represents the ClarifyInferenceConfig schema from the OpenAPI specification
type ClarifyInferenceConfig struct {
	Maxrecordcount interface{} `json:"MaxRecordCount,omitempty"`
	Probabilityindex interface{} `json:"ProbabilityIndex,omitempty"`
	Featureheaders interface{} `json:"FeatureHeaders,omitempty"`
	Labelheaders interface{} `json:"LabelHeaders,omitempty"`
	Probabilityattribute interface{} `json:"ProbabilityAttribute,omitempty"`
	Featuresattribute interface{} `json:"FeaturesAttribute,omitempty"`
	Featuretypes interface{} `json:"FeatureTypes,omitempty"`
	Maxpayloadinmb interface{} `json:"MaxPayloadInMB,omitempty"`
	Contenttemplate interface{} `json:"ContentTemplate,omitempty"`
	Labelattribute interface{} `json:"LabelAttribute,omitempty"`
	Labelindex interface{} `json:"LabelIndex,omitempty"`
}

// CreateEndpointConfigOutput represents the CreateEndpointConfigOutput schema from the OpenAPI specification
type CreateEndpointConfigOutput struct {
	Endpointconfigarn interface{} `json:"EndpointConfigArn"`
}

// AddAssociationRequest represents the AddAssociationRequest schema from the OpenAPI specification
type AddAssociationRequest struct {
	Sourcearn interface{} `json:"SourceArn"`
	Associationtype interface{} `json:"AssociationType,omitempty"`
	Destinationarn interface{} `json:"DestinationArn"`
}

// CreateWorkforceRequest represents the CreateWorkforceRequest schema from the OpenAPI specification
type CreateWorkforceRequest struct {
	Tags interface{} `json:"Tags,omitempty"`
	Workforcename interface{} `json:"WorkforceName"`
	Workforcevpcconfig interface{} `json:"WorkforceVpcConfig,omitempty"`
	Cognitoconfig interface{} `json:"CognitoConfig,omitempty"`
	Oidcconfig interface{} `json:"OidcConfig,omitempty"`
	Sourceipconfig SourceIpConfig `json:"SourceIpConfig,omitempty"` // A list of IP address ranges (<a href="https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html">CIDRs</a>). Used to create an allow list of IP addresses for a private workforce. Workers will only be able to login to their worker portal from an IP address within this range. By default, a workforce isn't restricted to specific IP addresses.
}

// AssociateTrialComponentResponse represents the AssociateTrialComponentResponse schema from the OpenAPI specification
type AssociateTrialComponentResponse struct {
	Trialarn interface{} `json:"TrialArn,omitempty"`
	Trialcomponentarn interface{} `json:"TrialComponentArn,omitempty"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Key interface{} `json:"Key"`
	Value interface{} `json:"Value"`
}

// UpdateFeatureGroupResponse represents the UpdateFeatureGroupResponse schema from the OpenAPI specification
type UpdateFeatureGroupResponse struct {
	Featuregrouparn interface{} `json:"FeatureGroupArn"`
}

// UpdateImageRequest represents the UpdateImageRequest schema from the OpenAPI specification
type UpdateImageRequest struct {
	Imagename interface{} `json:"ImageName"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Deleteproperties interface{} `json:"DeleteProperties,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Displayname interface{} `json:"DisplayName,omitempty"`
}

// DescribeLabelingJobResponse represents the DescribeLabelingJobResponse schema from the OpenAPI specification
type DescribeLabelingJobResponse struct {
	Labelingjobstatus interface{} `json:"LabelingJobStatus"`
	Outputconfig interface{} `json:"OutputConfig"`
	Tags interface{} `json:"Tags,omitempty"`
	Labelattributename interface{} `json:"LabelAttributeName,omitempty"`
	Labelingjobalgorithmsconfig interface{} `json:"LabelingJobAlgorithmsConfig,omitempty"`
	Inputconfig interface{} `json:"InputConfig"`
	Labelcategoryconfigs3uri interface{} `json:"LabelCategoryConfigS3Uri,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime"`
	Stoppingconditions interface{} `json:"StoppingConditions,omitempty"`
	Labelcounters interface{} `json:"LabelCounters"`
	Labelingjobarn interface{} `json:"LabelingJobArn"`
	Creationtime interface{} `json:"CreationTime"`
	Labelingjobname interface{} `json:"LabelingJobName"`
	Rolearn interface{} `json:"RoleArn"`
	Jobreferencecode interface{} `json:"JobReferenceCode"`
	Labelingjoboutput interface{} `json:"LabelingJobOutput,omitempty"`
	Humantaskconfig interface{} `json:"HumanTaskConfig"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
}

// DeleteAssociationRequest represents the DeleteAssociationRequest schema from the OpenAPI specification
type DeleteAssociationRequest struct {
	Destinationarn interface{} `json:"DestinationArn"`
	Sourcearn interface{} `json:"SourceArn"`
}

// ModelQualityJobInput represents the ModelQualityJobInput schema from the OpenAPI specification
type ModelQualityJobInput struct {
	Endpointinput EndpointInput `json:"EndpointInput,omitempty"` // Input object for the endpoint
	Groundtruths3input interface{} `json:"GroundTruthS3Input"`
	Batchtransforminput interface{} `json:"BatchTransformInput,omitempty"`
}

// DescribeInferenceRecommendationsJobResponse represents the DescribeInferenceRecommendationsJobResponse schema from the OpenAPI specification
type DescribeInferenceRecommendationsJobResponse struct {
	Jobtype interface{} `json:"JobType"`
	Endpointperformances interface{} `json:"EndpointPerformances,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Inputconfig interface{} `json:"InputConfig"`
	Jobdescription interface{} `json:"JobDescription,omitempty"`
	Jobname interface{} `json:"JobName"`
	Inferencerecommendations interface{} `json:"InferenceRecommendations,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime"`
	Rolearn interface{} `json:"RoleArn"`
	Status interface{} `json:"Status"`
	Creationtime interface{} `json:"CreationTime"`
	Jobarn interface{} `json:"JobArn"`
	Stoppingconditions interface{} `json:"StoppingConditions,omitempty"`
	Completiontime interface{} `json:"CompletionTime,omitempty"`
}

// DescribeProjectInput represents the DescribeProjectInput schema from the OpenAPI specification
type DescribeProjectInput struct {
	Projectname interface{} `json:"ProjectName"`
}

// ListTrialComponentsResponse represents the ListTrialComponentsResponse schema from the OpenAPI specification
type ListTrialComponentsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Trialcomponentsummaries interface{} `json:"TrialComponentSummaries,omitempty"`
}

// ProjectSummary represents the ProjectSummary schema from the OpenAPI specification
type ProjectSummary struct {
	Projectname interface{} `json:"ProjectName"`
	Projectstatus interface{} `json:"ProjectStatus"`
	Creationtime interface{} `json:"CreationTime"`
	Projectarn interface{} `json:"ProjectArn"`
	Projectdescription interface{} `json:"ProjectDescription,omitempty"`
	Projectid interface{} `json:"ProjectId"`
}

// PendingProductionVariantSummary represents the PendingProductionVariantSummary schema from the OpenAPI specification
type PendingProductionVariantSummary struct {
	Desiredserverlessconfig interface{} `json:"DesiredServerlessConfig,omitempty"`
	Currentinstancecount interface{} `json:"CurrentInstanceCount,omitempty"`
	Currentserverlessconfig interface{} `json:"CurrentServerlessConfig,omitempty"`
	Variantname interface{} `json:"VariantName"`
	Instancetype interface{} `json:"InstanceType,omitempty"`
	Variantstatus interface{} `json:"VariantStatus,omitempty"`
	Acceleratortype interface{} `json:"AcceleratorType,omitempty"`
	Currentweight interface{} `json:"CurrentWeight,omitempty"`
	Desiredweight interface{} `json:"DesiredWeight,omitempty"`
	Deployedimages interface{} `json:"DeployedImages,omitempty"`
	Desiredinstancecount interface{} `json:"DesiredInstanceCount,omitempty"`
}

// CallbackStepMetadata represents the CallbackStepMetadata schema from the OpenAPI specification
type CallbackStepMetadata struct {
	Callbacktoken interface{} `json:"CallbackToken,omitempty"`
	Outputparameters interface{} `json:"OutputParameters,omitempty"`
	Sqsqueueurl interface{} `json:"SqsQueueUrl,omitempty"`
}

// CreateModelQualityJobDefinitionResponse represents the CreateModelQualityJobDefinitionResponse schema from the OpenAPI specification
type CreateModelQualityJobDefinitionResponse struct {
	Jobdefinitionarn interface{} `json:"JobDefinitionArn"`
}

// ParameterRange represents the ParameterRange schema from the OpenAPI specification
type ParameterRange struct {
	Continuousparameterrangespecification interface{} `json:"ContinuousParameterRangeSpecification,omitempty"`
	Integerparameterrangespecification interface{} `json:"IntegerParameterRangeSpecification,omitempty"`
	Categoricalparameterrangespecification interface{} `json:"CategoricalParameterRangeSpecification,omitempty"`
}

// DeviceDeploymentSummary represents the DeviceDeploymentSummary schema from the OpenAPI specification
type DeviceDeploymentSummary struct {
	Deployedstagename interface{} `json:"DeployedStageName,omitempty"`
	Devicedeploymentstatus interface{} `json:"DeviceDeploymentStatus,omitempty"`
	Devicefleetname interface{} `json:"DeviceFleetName,omitempty"`
	Edgedeploymentplanname interface{} `json:"EdgeDeploymentPlanName"`
	Stagename interface{} `json:"StageName"`
	Devicearn interface{} `json:"DeviceArn"`
	Edgedeploymentplanarn interface{} `json:"EdgeDeploymentPlanArn"`
	Description interface{} `json:"Description,omitempty"`
	Devicename interface{} `json:"DeviceName"`
	Deploymentstarttime interface{} `json:"DeploymentStartTime,omitempty"`
	Devicedeploymentstatusmessage interface{} `json:"DeviceDeploymentStatusMessage,omitempty"`
}

// AutoMLJobCompletionCriteria represents the AutoMLJobCompletionCriteria schema from the OpenAPI specification
type AutoMLJobCompletionCriteria struct {
	Maxautomljobruntimeinseconds interface{} `json:"MaxAutoMLJobRuntimeInSeconds,omitempty"`
	Maxcandidates interface{} `json:"MaxCandidates,omitempty"`
	Maxruntimepertrainingjobinseconds interface{} `json:"MaxRuntimePerTrainingJobInSeconds,omitempty"`
}

// ProcessingS3Output represents the ProcessingS3Output schema from the OpenAPI specification
type ProcessingS3Output struct {
	S3uri interface{} `json:"S3Uri"`
	Localpath interface{} `json:"LocalPath"`
	S3uploadmode interface{} `json:"S3UploadMode"`
}

// MonitoringJobDefinitionSummary represents the MonitoringJobDefinitionSummary schema from the OpenAPI specification
type MonitoringJobDefinitionSummary struct {
	Creationtime interface{} `json:"CreationTime"`
	Endpointname interface{} `json:"EndpointName"`
	Monitoringjobdefinitionarn interface{} `json:"MonitoringJobDefinitionArn"`
	Monitoringjobdefinitionname interface{} `json:"MonitoringJobDefinitionName"`
}

// DescribeFeatureMetadataResponse represents the DescribeFeatureMetadataResponse schema from the OpenAPI specification
type DescribeFeatureMetadataResponse struct {
	Creationtime interface{} `json:"CreationTime"`
	Description interface{} `json:"Description,omitempty"`
	Featuregrouparn interface{} `json:"FeatureGroupArn"`
	Featuregroupname interface{} `json:"FeatureGroupName"`
	Featurename interface{} `json:"FeatureName"`
	Featuretype interface{} `json:"FeatureType"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime"`
	Parameters interface{} `json:"Parameters,omitempty"`
}

// EdgeDeploymentModelConfig represents the EdgeDeploymentModelConfig schema from the OpenAPI specification
type EdgeDeploymentModelConfig struct {
	Modelhandle interface{} `json:"ModelHandle"`
	Edgepackagingjobname interface{} `json:"EdgePackagingJobName"`
}

// ScalingPolicyObjective represents the ScalingPolicyObjective schema from the OpenAPI specification
type ScalingPolicyObjective struct {
	Maxinvocationsperminute interface{} `json:"MaxInvocationsPerMinute,omitempty"`
	Mininvocationsperminute interface{} `json:"MinInvocationsPerMinute,omitempty"`
}

// CreateTrialResponse represents the CreateTrialResponse schema from the OpenAPI specification
type CreateTrialResponse struct {
	Trialarn interface{} `json:"TrialArn,omitempty"`
}

// AggregationTransformations represents the AggregationTransformations schema from the OpenAPI specification
type AggregationTransformations struct {
}

// DescribeAppRequest represents the DescribeAppRequest schema from the OpenAPI specification
type DescribeAppRequest struct {
	Domainid interface{} `json:"DomainId"`
	Spacename interface{} `json:"SpaceName,omitempty"`
	Userprofilename interface{} `json:"UserProfileName,omitempty"`
	Appname interface{} `json:"AppName"`
	Apptype interface{} `json:"AppType"`
}

// HumanTaskConfig represents the HumanTaskConfig schema from the OpenAPI specification
type HumanTaskConfig struct {
	Taskdescription interface{} `json:"TaskDescription"`
	Workteamarn interface{} `json:"WorkteamArn"`
	Maxconcurrenttaskcount interface{} `json:"MaxConcurrentTaskCount,omitempty"`
	Numberofhumanworkersperdataobject interface{} `json:"NumberOfHumanWorkersPerDataObject"`
	Publicworkforcetaskprice interface{} `json:"PublicWorkforceTaskPrice,omitempty"`
	Tasktimelimitinseconds interface{} `json:"TaskTimeLimitInSeconds"`
	Annotationconsolidationconfig interface{} `json:"AnnotationConsolidationConfig"`
	Taskkeywords interface{} `json:"TaskKeywords,omitempty"`
	Prehumantasklambdaarn interface{} `json:"PreHumanTaskLambdaArn"`
	Taskavailabilitylifetimeinseconds interface{} `json:"TaskAvailabilityLifetimeInSeconds,omitempty"`
	Tasktitle interface{} `json:"TaskTitle"`
	Uiconfig interface{} `json:"UiConfig"`
}

// ListNotebookInstancesOutput represents the ListNotebookInstancesOutput schema from the OpenAPI specification
type ListNotebookInstancesOutput struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Notebookinstances interface{} `json:"NotebookInstances,omitempty"`
}

// ListFeatureGroupsResponse represents the ListFeatureGroupsResponse schema from the OpenAPI specification
type ListFeatureGroupsResponse struct {
	Featuregroupsummaries interface{} `json:"FeatureGroupSummaries"`
	Nexttoken interface{} `json:"NextToken"`
}

// DeletePipelineRequest represents the DeletePipelineRequest schema from the OpenAPI specification
type DeletePipelineRequest struct {
	Pipelinename interface{} `json:"PipelineName"`
	Clientrequesttoken interface{} `json:"ClientRequestToken"`
}

// CreateTrainingJobRequest represents the CreateTrainingJobRequest schema from the OpenAPI specification
type CreateTrainingJobRequest struct {
	Trainingjobname interface{} `json:"TrainingJobName"`
	Hyperparameters interface{} `json:"HyperParameters,omitempty"`
	Checkpointconfig interface{} `json:"CheckpointConfig,omitempty"`
	Debughookconfig DebugHookConfig `json:"DebugHookConfig,omitempty"` // Configuration information for the Amazon SageMaker Debugger hook parameters, metric and tensor collections, and storage paths. To learn more about how to configure the <code>DebugHookConfig</code> parameter, see <a href="https://docs.aws.amazon.com/sagemaker/latest/dg/debugger-createtrainingjob-api.html">Use the SageMaker and Debugger Configuration API Operations to Create, Update, and Debug Your Training Job</a>.
	Environment interface{} `json:"Environment,omitempty"`
	Inputdataconfig interface{} `json:"InputDataConfig,omitempty"`
	Algorithmspecification interface{} `json:"AlgorithmSpecification"`
	Debugruleconfigurations interface{} `json:"DebugRuleConfigurations,omitempty"`
	Vpcconfig interface{} `json:"VpcConfig,omitempty"`
	Enablemanagedspottraining interface{} `json:"EnableManagedSpotTraining,omitempty"`
	Profilerruleconfigurations interface{} `json:"ProfilerRuleConfigurations,omitempty"`
	Enableintercontainertrafficencryption interface{} `json:"EnableInterContainerTrafficEncryption,omitempty"`
	Profilerconfig ProfilerConfig `json:"ProfilerConfig,omitempty"` // Configuration information for Amazon SageMaker Debugger system monitoring, framework profiling, and storage paths.
	Tensorboardoutputconfig TensorBoardOutputConfig `json:"TensorBoardOutputConfig,omitempty"` // Configuration of storage locations for the Amazon SageMaker Debugger TensorBoard output data.
	Rolearn interface{} `json:"RoleArn"`
	Outputdataconfig interface{} `json:"OutputDataConfig"`
	Tags interface{} `json:"Tags,omitempty"`
	Stoppingcondition interface{} `json:"StoppingCondition"`
	Enablenetworkisolation interface{} `json:"EnableNetworkIsolation,omitempty"`
	Experimentconfig ExperimentConfig `json:"ExperimentConfig,omitempty"` // <p>Associates a SageMaker job as a trial component with an experiment and trial. Specified when you call the following APIs:</p> <ul> <li> <p> <a href="https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateProcessingJob.html">CreateProcessingJob</a> </p> </li> <li> <p> <a href="https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTrainingJob.html">CreateTrainingJob</a> </p> </li> <li> <p> <a href="https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTransformJob.html">CreateTransformJob</a> </p> </li> </ul>
	Resourceconfig interface{} `json:"ResourceConfig"`
	Retrystrategy interface{} `json:"RetryStrategy,omitempty"`
}

// UpdateContextRequest represents the UpdateContextRequest schema from the OpenAPI specification
type UpdateContextRequest struct {
	Description interface{} `json:"Description,omitempty"`
	Properties interface{} `json:"Properties,omitempty"`
	Propertiestoremove interface{} `json:"PropertiesToRemove,omitempty"`
	Contextname interface{} `json:"ContextName"`
}

// AutoMLInferenceContainerDefinitions represents the AutoMLInferenceContainerDefinitions schema from the OpenAPI specification
type AutoMLInferenceContainerDefinitions struct {
}

// StartMonitoringScheduleRequest represents the StartMonitoringScheduleRequest schema from the OpenAPI specification
type StartMonitoringScheduleRequest struct {
	Monitoringschedulename interface{} `json:"MonitoringScheduleName"`
}

// DeregisterDevicesRequest represents the DeregisterDevicesRequest schema from the OpenAPI specification
type DeregisterDevicesRequest struct {
	Devicefleetname interface{} `json:"DeviceFleetName"`
	Devicenames interface{} `json:"DeviceNames"`
}

// AutoMLResolvedAttributes represents the AutoMLResolvedAttributes schema from the OpenAPI specification
type AutoMLResolvedAttributes struct {
	Completioncriteria AutoMLJobCompletionCriteria `json:"CompletionCriteria,omitempty"` // How long a job is allowed to run, or how many candidates a job is allowed to generate.
	Automljobobjective AutoMLJobObjective `json:"AutoMLJobObjective,omitempty"` // Specifies a metric to minimize or maximize as the objective of a job.
	Automlproblemtyperesolvedattributes interface{} `json:"AutoMLProblemTypeResolvedAttributes,omitempty"`
}

// StopInferenceExperimentRequest represents the StopInferenceExperimentRequest schema from the OpenAPI specification
type StopInferenceExperimentRequest struct {
	Modelvariantactions interface{} `json:"ModelVariantActions"`
	Name interface{} `json:"Name"`
	Reason interface{} `json:"Reason,omitempty"`
	Desiredmodelvariants interface{} `json:"DesiredModelVariants,omitempty"`
	Desiredstate interface{} `json:"DesiredState,omitempty"`
}

// CreateDomainRequest represents the CreateDomainRequest schema from the OpenAPI specification
type CreateDomainRequest struct {
	Defaultspacesettings interface{} `json:"DefaultSpaceSettings,omitempty"`
	Defaultusersettings interface{} `json:"DefaultUserSettings"`
	Domainsettings interface{} `json:"DomainSettings,omitempty"`
	Appnetworkaccesstype interface{} `json:"AppNetworkAccessType,omitempty"`
	Homeefsfilesystemkmskeyid interface{} `json:"HomeEfsFileSystemKmsKeyId,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Appsecuritygroupmanagement interface{} `json:"AppSecurityGroupManagement,omitempty"`
	Authmode interface{} `json:"AuthMode"`
	Domainname interface{} `json:"DomainName"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Vpcid interface{} `json:"VpcId"`
	Subnetids interface{} `json:"SubnetIds"`
}

// DeleteWorkforceRequest represents the DeleteWorkforceRequest schema from the OpenAPI specification
type DeleteWorkforceRequest struct {
	Workforcename interface{} `json:"WorkforceName"`
}

// UpdateTrialComponentResponse represents the UpdateTrialComponentResponse schema from the OpenAPI specification
type UpdateTrialComponentResponse struct {
	Trialcomponentarn interface{} `json:"TrialComponentArn,omitempty"`
}

// AlgorithmSummary represents the AlgorithmSummary schema from the OpenAPI specification
type AlgorithmSummary struct {
	Algorithmarn interface{} `json:"AlgorithmArn"`
	Algorithmdescription interface{} `json:"AlgorithmDescription,omitempty"`
	Algorithmname interface{} `json:"AlgorithmName"`
	Algorithmstatus interface{} `json:"AlgorithmStatus"`
	Creationtime interface{} `json:"CreationTime"`
}

// TransformJobDefinition represents the TransformJobDefinition schema from the OpenAPI specification
type TransformJobDefinition struct {
	Transformresources interface{} `json:"TransformResources"`
	Batchstrategy interface{} `json:"BatchStrategy,omitempty"`
	Environment interface{} `json:"Environment,omitempty"`
	Maxconcurrenttransforms interface{} `json:"MaxConcurrentTransforms,omitempty"`
	Maxpayloadinmb interface{} `json:"MaxPayloadInMB,omitempty"`
	Transforminput interface{} `json:"TransformInput"`
	Transformoutput interface{} `json:"TransformOutput"`
}

// DescribeWorkforceResponse represents the DescribeWorkforceResponse schema from the OpenAPI specification
type DescribeWorkforceResponse struct {
	Workforce interface{} `json:"Workforce"`
}

// ListEdgeDeploymentPlansRequest represents the ListEdgeDeploymentPlansRequest schema from the OpenAPI specification
type ListEdgeDeploymentPlansRequest struct {
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Lastmodifiedtimebefore interface{} `json:"LastModifiedTimeBefore,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
	Devicefleetnamecontains interface{} `json:"DeviceFleetNameContains,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Lastmodifiedtimeafter interface{} `json:"LastModifiedTimeAfter,omitempty"`
}

// ListCodeRepositoriesInput represents the ListCodeRepositoriesInput schema from the OpenAPI specification
type ListCodeRepositoriesInput struct {
	Namecontains interface{} `json:"NameContains,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Lastmodifiedtimebefore interface{} `json:"LastModifiedTimeBefore,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Lastmodifiedtimeafter interface{} `json:"LastModifiedTimeAfter,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// ListMonitoringSchedulesResponse represents the ListMonitoringSchedulesResponse schema from the OpenAPI specification
type ListMonitoringSchedulesResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Monitoringschedulesummaries interface{} `json:"MonitoringScheduleSummaries"`
}

// HyperParameterTrainingJobEnvironmentMap represents the HyperParameterTrainingJobEnvironmentMap schema from the OpenAPI specification
type HyperParameterTrainingJobEnvironmentMap struct {
}

// DescribeDataQualityJobDefinitionRequest represents the DescribeDataQualityJobDefinitionRequest schema from the OpenAPI specification
type DescribeDataQualityJobDefinitionRequest struct {
	Jobdefinitionname interface{} `json:"JobDefinitionName"`
}

// CreateNotebookInstanceLifecycleConfigOutput represents the CreateNotebookInstanceLifecycleConfigOutput schema from the OpenAPI specification
type CreateNotebookInstanceLifecycleConfigOutput struct {
	Notebookinstancelifecycleconfigarn interface{} `json:"NotebookInstanceLifecycleConfigArn,omitempty"`
}

// HubContentInfo represents the HubContentInfo schema from the OpenAPI specification
type HubContentInfo struct {
	Hubcontenttype interface{} `json:"HubContentType"`
	Documentschemaversion interface{} `json:"DocumentSchemaVersion"`
	Hubcontentdescription interface{} `json:"HubContentDescription,omitempty"`
	Hubcontentdisplayname interface{} `json:"HubContentDisplayName,omitempty"`
	Hubcontentname interface{} `json:"HubContentName"`
	Hubcontentversion interface{} `json:"HubContentVersion"`
	Creationtime interface{} `json:"CreationTime"`
	Hubcontentarn interface{} `json:"HubContentArn"`
	Hubcontentsearchkeywords interface{} `json:"HubContentSearchKeywords,omitempty"`
	Hubcontentstatus interface{} `json:"HubContentStatus"`
}

// UpdateUserProfileRequest represents the UpdateUserProfileRequest schema from the OpenAPI specification
type UpdateUserProfileRequest struct {
	Domainid interface{} `json:"DomainId"`
	Userprofilename interface{} `json:"UserProfileName"`
	Usersettings interface{} `json:"UserSettings,omitempty"`
}

// ListDataQualityJobDefinitionsResponse represents the ListDataQualityJobDefinitionsResponse schema from the OpenAPI specification
type ListDataQualityJobDefinitionsResponse struct {
	Jobdefinitionsummaries interface{} `json:"JobDefinitionSummaries"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ArtifactSourceType represents the ArtifactSourceType schema from the OpenAPI specification
type ArtifactSourceType struct {
	Sourceidtype interface{} `json:"SourceIdType"`
	Value interface{} `json:"Value"`
}

// ModelDigests represents the ModelDigests schema from the OpenAPI specification
type ModelDigests struct {
	Artifactdigest interface{} `json:"ArtifactDigest,omitempty"`
}

// ServiceCatalogProvisioningDetails represents the ServiceCatalogProvisioningDetails schema from the OpenAPI specification
type ServiceCatalogProvisioningDetails struct {
	Pathid interface{} `json:"PathId,omitempty"`
	Productid interface{} `json:"ProductId"`
	Provisioningartifactid interface{} `json:"ProvisioningArtifactId,omitempty"`
	Provisioningparameters interface{} `json:"ProvisioningParameters,omitempty"`
}

// CreateTrialComponentResponse represents the CreateTrialComponentResponse schema from the OpenAPI specification
type CreateTrialComponentResponse struct {
	Trialcomponentarn interface{} `json:"TrialComponentArn,omitempty"`
}

// DescribeDeviceRequest represents the DescribeDeviceRequest schema from the OpenAPI specification
type DescribeDeviceRequest struct {
	Devicefleetname interface{} `json:"DeviceFleetName"`
	Devicename interface{} `json:"DeviceName"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// USD represents the USD schema from the OpenAPI specification
type USD struct {
	Cents interface{} `json:"Cents,omitempty"`
	Dollars interface{} `json:"Dollars,omitempty"`
	Tenthfractionsofacent interface{} `json:"TenthFractionsOfACent,omitempty"`
}

// EnvironmentParameterRanges represents the EnvironmentParameterRanges schema from the OpenAPI specification
type EnvironmentParameterRanges struct {
	Categoricalparameterranges interface{} `json:"CategoricalParameterRanges,omitempty"`
}

// AlgorithmStatusDetails represents the AlgorithmStatusDetails schema from the OpenAPI specification
type AlgorithmStatusDetails struct {
	Imagescanstatuses interface{} `json:"ImageScanStatuses,omitempty"`
	Validationstatuses interface{} `json:"ValidationStatuses,omitempty"`
}

// LambdaStepMetadata represents the LambdaStepMetadata schema from the OpenAPI specification
type LambdaStepMetadata struct {
	Arn interface{} `json:"Arn,omitempty"`
	Outputparameters interface{} `json:"OutputParameters,omitempty"`
}

// DescribeCompilationJobRequest represents the DescribeCompilationJobRequest schema from the OpenAPI specification
type DescribeCompilationJobRequest struct {
	Compilationjobname interface{} `json:"CompilationJobName"`
}

// CreateEndpointConfigInput represents the CreateEndpointConfigInput schema from the OpenAPI specification
type CreateEndpointConfigInput struct {
	Explainerconfig interface{} `json:"ExplainerConfig,omitempty"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Productionvariants interface{} `json:"ProductionVariants"`
	Shadowproductionvariants interface{} `json:"ShadowProductionVariants,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Asyncinferenceconfig interface{} `json:"AsyncInferenceConfig,omitempty"`
	Datacaptureconfig DataCaptureConfig `json:"DataCaptureConfig,omitempty"` // Configuration to control how SageMaker captures inference data.
	Endpointconfigname interface{} `json:"EndpointConfigName"`
}

// UpdateModelCardResponse represents the UpdateModelCardResponse schema from the OpenAPI specification
type UpdateModelCardResponse struct {
	Modelcardarn interface{} `json:"ModelCardArn"`
}

// FailStepMetadata represents the FailStepMetadata schema from the OpenAPI specification
type FailStepMetadata struct {
	Errormessage interface{} `json:"ErrorMessage,omitempty"`
}

// ListProcessingJobsRequest represents the ListProcessingJobsRequest schema from the OpenAPI specification
type ListProcessingJobsRequest struct {
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Lastmodifiedtimeafter interface{} `json:"LastModifiedTimeAfter,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
	Statusequals interface{} `json:"StatusEquals,omitempty"`
	Lastmodifiedtimebefore interface{} `json:"LastModifiedTimeBefore,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
}

// RuleParameters represents the RuleParameters schema from the OpenAPI specification
type RuleParameters struct {
}

// StopInferenceExperimentResponse represents the StopInferenceExperimentResponse schema from the OpenAPI specification
type StopInferenceExperimentResponse struct {
	Inferenceexperimentarn interface{} `json:"InferenceExperimentArn"`
}

// PipelineExecutionStep represents the PipelineExecutionStep schema from the OpenAPI specification
type PipelineExecutionStep struct {
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Stepdescription interface{} `json:"StepDescription,omitempty"`
	Stepdisplayname interface{} `json:"StepDisplayName,omitempty"`
	Stepname interface{} `json:"StepName,omitempty"`
	Metadata interface{} `json:"Metadata,omitempty"`
	Selectiveexecutionresult interface{} `json:"SelectiveExecutionResult,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
	Endtime interface{} `json:"EndTime,omitempty"`
	Stepstatus interface{} `json:"StepStatus,omitempty"`
	Attemptcount interface{} `json:"AttemptCount,omitempty"`
	Cachehitresult interface{} `json:"CacheHitResult,omitempty"`
}

// ListAppImageConfigsResponse represents the ListAppImageConfigsResponse schema from the OpenAPI specification
type ListAppImageConfigsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Appimageconfigs interface{} `json:"AppImageConfigs,omitempty"`
}

// HyperbandStrategyConfig represents the HyperbandStrategyConfig schema from the OpenAPI specification
type HyperbandStrategyConfig struct {
	Minresource interface{} `json:"MinResource,omitempty"`
	Maxresource interface{} `json:"MaxResource,omitempty"`
}

// DescribeAppResponse represents the DescribeAppResponse schema from the OpenAPI specification
type DescribeAppResponse struct {
	Status interface{} `json:"Status,omitempty"`
	Userprofilename interface{} `json:"UserProfileName,omitempty"`
	Apptype interface{} `json:"AppType,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Apparn interface{} `json:"AppArn,omitempty"`
	Appname interface{} `json:"AppName,omitempty"`
	Lastuseractivitytimestamp interface{} `json:"LastUserActivityTimestamp,omitempty"`
	Domainid interface{} `json:"DomainId,omitempty"`
	Spacename interface{} `json:"SpaceName,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Lasthealthchecktimestamp interface{} `json:"LastHealthCheckTimestamp,omitempty"`
	Resourcespec interface{} `json:"ResourceSpec,omitempty"`
}

// DescribeTransformJobResponse represents the DescribeTransformJobResponse schema from the OpenAPI specification
type DescribeTransformJobResponse struct {
	Batchstrategy interface{} `json:"BatchStrategy,omitempty"`
	Dataprocessing DataProcessing `json:"DataProcessing,omitempty"` // The data structure used to specify the data to be used for inference in a batch transform job and to associate the data that is relevant to the prediction results in the output. The input filter provided allows you to exclude input data that is not needed for inference in a batch transform job. The output filter provided allows you to include input data relevant to interpreting the predictions in the output from the job. For more information, see <a href="https://docs.aws.amazon.com/sagemaker/latest/dg/batch-transform-data-processing.html">Associate Prediction Results with their Corresponding Input Records</a>.
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Modelname interface{} `json:"ModelName"`
	Transforminput interface{} `json:"TransformInput"`
	Transformstarttime interface{} `json:"TransformStartTime,omitempty"`
	Automljobarn interface{} `json:"AutoMLJobArn,omitempty"`
	Experimentconfig ExperimentConfig `json:"ExperimentConfig,omitempty"` // <p>Associates a SageMaker job as a trial component with an experiment and trial. Specified when you call the following APIs:</p> <ul> <li> <p> <a href="https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateProcessingJob.html">CreateProcessingJob</a> </p> </li> <li> <p> <a href="https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTrainingJob.html">CreateTrainingJob</a> </p> </li> <li> <p> <a href="https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTransformJob.html">CreateTransformJob</a> </p> </li> </ul>
	Transformjobname interface{} `json:"TransformJobName"`
	Transformjobarn interface{} `json:"TransformJobArn"`
	Datacaptureconfig interface{} `json:"DataCaptureConfig,omitempty"`
	Environment interface{} `json:"Environment,omitempty"`
	Labelingjobarn interface{} `json:"LabelingJobArn,omitempty"`
	Maxpayloadinmb interface{} `json:"MaxPayloadInMB,omitempty"`
	Transformjobstatus interface{} `json:"TransformJobStatus"`
	Creationtime interface{} `json:"CreationTime"`
	Transformendtime interface{} `json:"TransformEndTime,omitempty"`
	Transformoutput interface{} `json:"TransformOutput,omitempty"`
	Modelclientconfig interface{} `json:"ModelClientConfig,omitempty"`
	Transformresources interface{} `json:"TransformResources"`
	Maxconcurrenttransforms interface{} `json:"MaxConcurrentTransforms,omitempty"`
}

// ModelDeployConfig represents the ModelDeployConfig schema from the OpenAPI specification
type ModelDeployConfig struct {
	Autogenerateendpointname interface{} `json:"AutoGenerateEndpointName,omitempty"`
	Endpointname interface{} `json:"EndpointName,omitempty"`
}

// DescribeCodeRepositoryInput represents the DescribeCodeRepositoryInput schema from the OpenAPI specification
type DescribeCodeRepositoryInput struct {
	Coderepositoryname interface{} `json:"CodeRepositoryName"`
}

// UpdateDevicesRequest represents the UpdateDevicesRequest schema from the OpenAPI specification
type UpdateDevicesRequest struct {
	Devicefleetname interface{} `json:"DeviceFleetName"`
	Devices interface{} `json:"Devices"`
}

// ProcessingFeatureStoreOutput represents the ProcessingFeatureStoreOutput schema from the OpenAPI specification
type ProcessingFeatureStoreOutput struct {
	Featuregroupname interface{} `json:"FeatureGroupName"`
}

// TrialComponentParameters represents the TrialComponentParameters schema from the OpenAPI specification
type TrialComponentParameters struct {
}

// MonitoringAppSpecification represents the MonitoringAppSpecification schema from the OpenAPI specification
type MonitoringAppSpecification struct {
	Containerarguments interface{} `json:"ContainerArguments,omitempty"`
	Containerentrypoint interface{} `json:"ContainerEntrypoint,omitempty"`
	Imageuri interface{} `json:"ImageUri"`
	Postanalyticsprocessorsourceuri interface{} `json:"PostAnalyticsProcessorSourceUri,omitempty"`
	Recordpreprocessorsourceuri interface{} `json:"RecordPreprocessorSourceUri,omitempty"`
}

// AppSpecification represents the AppSpecification schema from the OpenAPI specification
type AppSpecification struct {
	Containerentrypoint interface{} `json:"ContainerEntrypoint,omitempty"`
	Imageuri interface{} `json:"ImageUri"`
	Containerarguments interface{} `json:"ContainerArguments,omitempty"`
}

// ListCandidatesForAutoMLJobRequest represents the ListCandidatesForAutoMLJobRequest schema from the OpenAPI specification
type ListCandidatesForAutoMLJobRequest struct {
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Statusequals interface{} `json:"StatusEquals,omitempty"`
	Automljobname interface{} `json:"AutoMLJobName"`
	Candidatenameequals interface{} `json:"CandidateNameEquals,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CreateEndpointOutput represents the CreateEndpointOutput schema from the OpenAPI specification
type CreateEndpointOutput struct {
	Endpointarn interface{} `json:"EndpointArn"`
}

// CreateEdgeDeploymentPlanResponse represents the CreateEdgeDeploymentPlanResponse schema from the OpenAPI specification
type CreateEdgeDeploymentPlanResponse struct {
	Edgedeploymentplanarn interface{} `json:"EdgeDeploymentPlanArn"`
}

// ModelBiasAppSpecification represents the ModelBiasAppSpecification schema from the OpenAPI specification
type ModelBiasAppSpecification struct {
	Imageuri interface{} `json:"ImageUri"`
	Configuri interface{} `json:"ConfigUri"`
	Environment interface{} `json:"Environment,omitempty"`
}

// DescribeModelQualityJobDefinitionResponse represents the DescribeModelQualityJobDefinitionResponse schema from the OpenAPI specification
type DescribeModelQualityJobDefinitionResponse struct {
	Stoppingcondition MonitoringStoppingCondition `json:"StoppingCondition,omitempty"` // A time limit for how long the monitoring job is allowed to run before stopping.
	Creationtime interface{} `json:"CreationTime"`
	Modelqualityjoboutputconfig MonitoringOutputConfig `json:"ModelQualityJobOutputConfig"` // The output configuration for monitoring jobs.
	Jobdefinitionname interface{} `json:"JobDefinitionName"`
	Modelqualityjobinput interface{} `json:"ModelQualityJobInput"`
	Networkconfig interface{} `json:"NetworkConfig,omitempty"`
	Rolearn interface{} `json:"RoleArn"`
	Jobdefinitionarn interface{} `json:"JobDefinitionArn"`
	Jobresources MonitoringResources `json:"JobResources"` // Identifies the resources to deploy for a monitoring job.
	Modelqualityappspecification interface{} `json:"ModelQualityAppSpecification"`
	Modelqualitybaselineconfig interface{} `json:"ModelQualityBaselineConfig,omitempty"`
}

// TimeSeriesConfig represents the TimeSeriesConfig schema from the OpenAPI specification
type TimeSeriesConfig struct {
	Groupingattributenames interface{} `json:"GroupingAttributeNames,omitempty"`
	Itemidentifierattributename interface{} `json:"ItemIdentifierAttributeName"`
	Targetattributename interface{} `json:"TargetAttributeName"`
	Timestampattributename interface{} `json:"TimestampAttributeName"`
}

// UpdateTrainingJobResponse represents the UpdateTrainingJobResponse schema from the OpenAPI specification
type UpdateTrainingJobResponse struct {
	Trainingjobarn interface{} `json:"TrainingJobArn"`
}

// UpdateModelPackageOutput represents the UpdateModelPackageOutput schema from the OpenAPI specification
type UpdateModelPackageOutput struct {
	Modelpackagearn interface{} `json:"ModelPackageArn"`
}

// DeleteWorkteamRequest represents the DeleteWorkteamRequest schema from the OpenAPI specification
type DeleteWorkteamRequest struct {
	Workteamname interface{} `json:"WorkteamName"`
}

// ListLabelingJobsForWorkteamResponse represents the ListLabelingJobsForWorkteamResponse schema from the OpenAPI specification
type ListLabelingJobsForWorkteamResponse struct {
	Labelingjobsummarylist interface{} `json:"LabelingJobSummaryList"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// EMRStepMetadata represents the EMRStepMetadata schema from the OpenAPI specification
type EMRStepMetadata struct {
	Stepid interface{} `json:"StepId,omitempty"`
	Stepname interface{} `json:"StepName,omitempty"`
	Clusterid interface{} `json:"ClusterId,omitempty"`
	Logfilepath interface{} `json:"LogFilePath,omitempty"`
}

// CreateWorkforceResponse represents the CreateWorkforceResponse schema from the OpenAPI specification
type CreateWorkforceResponse struct {
	Workforcearn interface{} `json:"WorkforceArn"`
}

// DescribeModelCardExportJobRequest represents the DescribeModelCardExportJobRequest schema from the OpenAPI specification
type DescribeModelCardExportJobRequest struct {
	Modelcardexportjobarn interface{} `json:"ModelCardExportJobArn"`
}

// EdgePresetDeploymentOutput represents the EdgePresetDeploymentOutput schema from the OpenAPI specification
type EdgePresetDeploymentOutput struct {
	TypeField interface{} `json:"Type"`
	Artifact interface{} `json:"Artifact,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Statusmessage interface{} `json:"StatusMessage,omitempty"`
}

// SourceIpConfig represents the SourceIpConfig schema from the OpenAPI specification
type SourceIpConfig struct {
	Cidrs interface{} `json:"Cidrs"`
}

// StopProcessingJobRequest represents the StopProcessingJobRequest schema from the OpenAPI specification
type StopProcessingJobRequest struct {
	Processingjobname interface{} `json:"ProcessingJobName"`
}

// CreateUserProfileResponse represents the CreateUserProfileResponse schema from the OpenAPI specification
type CreateUserProfileResponse struct {
	Userprofilearn interface{} `json:"UserProfileArn,omitempty"`
}

// CreateAppImageConfigResponse represents the CreateAppImageConfigResponse schema from the OpenAPI specification
type CreateAppImageConfigResponse struct {
	Appimageconfigarn interface{} `json:"AppImageConfigArn,omitempty"`
}

// CreateFlowDefinitionResponse represents the CreateFlowDefinitionResponse schema from the OpenAPI specification
type CreateFlowDefinitionResponse struct {
	Flowdefinitionarn interface{} `json:"FlowDefinitionArn"`
}

// OnlineStoreConfigUpdate represents the OnlineStoreConfigUpdate schema from the OpenAPI specification
type OnlineStoreConfigUpdate struct {
	Ttlduration interface{} `json:"TtlDuration,omitempty"`
}

// ImageConfig represents the ImageConfig schema from the OpenAPI specification
type ImageConfig struct {
	Repositoryauthconfig interface{} `json:"RepositoryAuthConfig,omitempty"`
	Repositoryaccessmode interface{} `json:"RepositoryAccessMode"`
}

// AutoParameter represents the AutoParameter schema from the OpenAPI specification
type AutoParameter struct {
	Name interface{} `json:"Name"`
	Valuehint interface{} `json:"ValueHint"`
}

// DeleteEndpointConfigInput represents the DeleteEndpointConfigInput schema from the OpenAPI specification
type DeleteEndpointConfigInput struct {
	Endpointconfigname interface{} `json:"EndpointConfigName"`
}

// DescribeNotebookInstanceOutput represents the DescribeNotebookInstanceOutput schema from the OpenAPI specification
type DescribeNotebookInstanceOutput struct {
	Notebookinstancestatus interface{} `json:"NotebookInstanceStatus,omitempty"`
	Platformidentifier interface{} `json:"PlatformIdentifier,omitempty"`
	Notebookinstancelifecycleconfigname interface{} `json:"NotebookInstanceLifecycleConfigName,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Securitygroups interface{} `json:"SecurityGroups,omitempty"`
	Url interface{} `json:"Url,omitempty"`
	Volumesizeingb interface{} `json:"VolumeSizeInGB,omitempty"`
	Directinternetaccess interface{} `json:"DirectInternetAccess,omitempty"`
	Additionalcoderepositories interface{} `json:"AdditionalCodeRepositories,omitempty"`
	Instancemetadataserviceconfiguration interface{} `json:"InstanceMetadataServiceConfiguration,omitempty"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Rootaccess interface{} `json:"RootAccess,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Defaultcoderepository interface{} `json:"DefaultCodeRepository,omitempty"`
	Acceleratortypes interface{} `json:"AcceleratorTypes,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Notebookinstancename interface{} `json:"NotebookInstanceName,omitempty"`
	Instancetype interface{} `json:"InstanceType,omitempty"`
	Subnetid interface{} `json:"SubnetId,omitempty"`
	Networkinterfaceid interface{} `json:"NetworkInterfaceId,omitempty"`
	Notebookinstancearn interface{} `json:"NotebookInstanceArn,omitempty"`
}

// ProductionVariantServerlessConfig represents the ProductionVariantServerlessConfig schema from the OpenAPI specification
type ProductionVariantServerlessConfig struct {
	Provisionedconcurrency interface{} `json:"ProvisionedConcurrency,omitempty"`
	Maxconcurrency interface{} `json:"MaxConcurrency"`
	Memorysizeinmb interface{} `json:"MemorySizeInMB"`
}

// UpdateCodeRepositoryInput represents the UpdateCodeRepositoryInput schema from the OpenAPI specification
type UpdateCodeRepositoryInput struct {
	Coderepositoryname interface{} `json:"CodeRepositoryName"`
	Gitconfig interface{} `json:"GitConfig,omitempty"`
}

// DeleteModelPackageGroupInput represents the DeleteModelPackageGroupInput schema from the OpenAPI specification
type DeleteModelPackageGroupInput struct {
	Modelpackagegroupname interface{} `json:"ModelPackageGroupName"`
}

// ListModelCardsRequest represents the ListModelCardsRequest schema from the OpenAPI specification
type ListModelCardsRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Modelcardstatus interface{} `json:"ModelCardStatus,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
}

// CreateInferenceExperimentResponse represents the CreateInferenceExperimentResponse schema from the OpenAPI specification
type CreateInferenceExperimentResponse struct {
	Inferenceexperimentarn interface{} `json:"InferenceExperimentArn"`
}

// MonitoringAlertActions represents the MonitoringAlertActions schema from the OpenAPI specification
type MonitoringAlertActions struct {
	Modeldashboardindicator interface{} `json:"ModelDashboardIndicator,omitempty"`
}

// ModelBiasBaselineConfig represents the ModelBiasBaselineConfig schema from the OpenAPI specification
type ModelBiasBaselineConfig struct {
	Baseliningjobname interface{} `json:"BaseliningJobName,omitempty"`
	Constraintsresource MonitoringConstraintsResource `json:"ConstraintsResource,omitempty"` // The constraints resource for a monitoring job.
}

// UpdateNotebookInstanceInput represents the UpdateNotebookInstanceInput schema from the OpenAPI specification
type UpdateNotebookInstanceInput struct {
	Disassociatedefaultcoderepository interface{} `json:"DisassociateDefaultCodeRepository,omitempty"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Rootaccess interface{} `json:"RootAccess,omitempty"`
	Acceleratortypes interface{} `json:"AcceleratorTypes,omitempty"`
	Disassociateacceleratortypes interface{} `json:"DisassociateAcceleratorTypes,omitempty"`
	Disassociatelifecycleconfig interface{} `json:"DisassociateLifecycleConfig,omitempty"`
	Notebookinstancename interface{} `json:"NotebookInstanceName"`
	Disassociateadditionalcoderepositories interface{} `json:"DisassociateAdditionalCodeRepositories,omitempty"`
	Instancemetadataserviceconfiguration interface{} `json:"InstanceMetadataServiceConfiguration,omitempty"`
	Instancetype interface{} `json:"InstanceType,omitempty"`
	Lifecycleconfigname interface{} `json:"LifecycleConfigName,omitempty"`
	Volumesizeingb interface{} `json:"VolumeSizeInGB,omitempty"`
	Additionalcoderepositories interface{} `json:"AdditionalCodeRepositories,omitempty"`
	Defaultcoderepository interface{} `json:"DefaultCodeRepository,omitempty"`
}

// ListContextsRequest represents the ListContextsRequest schema from the OpenAPI specification
type ListContextsRequest struct {
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Sourceuri interface{} `json:"SourceUri,omitempty"`
	Contexttype interface{} `json:"ContextType,omitempty"`
	Createdafter interface{} `json:"CreatedAfter,omitempty"`
	Createdbefore interface{} `json:"CreatedBefore,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
}

// DescribeModelBiasJobDefinitionRequest represents the DescribeModelBiasJobDefinitionRequest schema from the OpenAPI specification
type DescribeModelBiasJobDefinitionRequest struct {
	Jobdefinitionname interface{} `json:"JobDefinitionName"`
}

// DescribeAppImageConfigRequest represents the DescribeAppImageConfigRequest schema from the OpenAPI specification
type DescribeAppImageConfigRequest struct {
	Appimageconfigname interface{} `json:"AppImageConfigName"`
}

// SpaceDetails represents the SpaceDetails schema from the OpenAPI specification
type SpaceDetails struct {
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Spacename interface{} `json:"SpaceName,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Domainid interface{} `json:"DomainId,omitempty"`
}

// ListModelExplainabilityJobDefinitionsResponse represents the ListModelExplainabilityJobDefinitionsResponse schema from the OpenAPI specification
type ListModelExplainabilityJobDefinitionsResponse struct {
	Jobdefinitionsummaries interface{} `json:"JobDefinitionSummaries"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// InferenceRecommendation represents the InferenceRecommendation schema from the OpenAPI specification
type InferenceRecommendation struct {
	Invocationstarttime interface{} `json:"InvocationStartTime,omitempty"`
	Metrics interface{} `json:"Metrics"`
	Modelconfiguration interface{} `json:"ModelConfiguration"`
	Recommendationid interface{} `json:"RecommendationId,omitempty"`
	Endpointconfiguration interface{} `json:"EndpointConfiguration"`
	Invocationendtime interface{} `json:"InvocationEndTime,omitempty"`
}

// ListModelPackageGroupsOutput represents the ListModelPackageGroupsOutput schema from the OpenAPI specification
type ListModelPackageGroupsOutput struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Modelpackagegroupsummarylist interface{} `json:"ModelPackageGroupSummaryList"`
}

// ProcessingJobStepMetadata represents the ProcessingJobStepMetadata schema from the OpenAPI specification
type ProcessingJobStepMetadata struct {
	Arn interface{} `json:"Arn,omitempty"`
}

// AutoRollbackConfig represents the AutoRollbackConfig schema from the OpenAPI specification
type AutoRollbackConfig struct {
	Alarms interface{} `json:"Alarms,omitempty"`
}

// DeleteModelBiasJobDefinitionRequest represents the DeleteModelBiasJobDefinitionRequest schema from the OpenAPI specification
type DeleteModelBiasJobDefinitionRequest struct {
	Jobdefinitionname interface{} `json:"JobDefinitionName"`
}

// AutoMLProblemTypeResolvedAttributes represents the AutoMLProblemTypeResolvedAttributes schema from the OpenAPI specification
type AutoMLProblemTypeResolvedAttributes struct {
	Tabularresolvedattributes interface{} `json:"TabularResolvedAttributes,omitempty"`
}

// SearchExpression represents the SearchExpression schema from the OpenAPI specification
type SearchExpression struct {
	Operator interface{} `json:"Operator,omitempty"`
	Subexpressions interface{} `json:"SubExpressions,omitempty"`
	Filters interface{} `json:"Filters,omitempty"`
	Nestedfilters interface{} `json:"NestedFilters,omitempty"`
}

// ListHubContentsRequest represents the ListHubContentsRequest schema from the OpenAPI specification
type ListHubContentsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Hubcontenttype interface{} `json:"HubContentType"`
	Maxschemaversion interface{} `json:"MaxSchemaVersion,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Hubname interface{} `json:"HubName"`
	Namecontains interface{} `json:"NameContains,omitempty"`
}

// ListSubscribedWorkteamsRequest represents the ListSubscribedWorkteamsRequest schema from the OpenAPI specification
type ListSubscribedWorkteamsRequest struct {
	Namecontains interface{} `json:"NameContains,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// CreateEdgePackagingJobRequest represents the CreateEdgePackagingJobRequest schema from the OpenAPI specification
type CreateEdgePackagingJobRequest struct {
	Rolearn interface{} `json:"RoleArn"`
	Tags interface{} `json:"Tags,omitempty"`
	Compilationjobname interface{} `json:"CompilationJobName"`
	Edgepackagingjobname interface{} `json:"EdgePackagingJobName"`
	Modelname interface{} `json:"ModelName"`
	Modelversion interface{} `json:"ModelVersion"`
	Outputconfig interface{} `json:"OutputConfig"`
	Resourcekey interface{} `json:"ResourceKey,omitempty"`
}

// DeviceFleetSummary represents the DeviceFleetSummary schema from the OpenAPI specification
type DeviceFleetSummary struct {
	Devicefleetarn interface{} `json:"DeviceFleetArn"`
	Devicefleetname interface{} `json:"DeviceFleetName"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
}

// DescribeAlgorithmOutput represents the DescribeAlgorithmOutput schema from the OpenAPI specification
type DescribeAlgorithmOutput struct {
	Validationspecification interface{} `json:"ValidationSpecification,omitempty"`
	Algorithmarn interface{} `json:"AlgorithmArn"`
	Algorithmstatusdetails interface{} `json:"AlgorithmStatusDetails"`
	Algorithmstatus interface{} `json:"AlgorithmStatus"`
	Inferencespecification interface{} `json:"InferenceSpecification,omitempty"`
	Productid interface{} `json:"ProductId,omitempty"`
	Algorithmname interface{} `json:"AlgorithmName"`
	Certifyformarketplace interface{} `json:"CertifyForMarketplace,omitempty"`
	Creationtime interface{} `json:"CreationTime"`
	Algorithmdescription interface{} `json:"AlgorithmDescription,omitempty"`
	Trainingspecification interface{} `json:"TrainingSpecification"`
}

// TransformJob represents the TransformJob schema from the OpenAPI specification
type TransformJob struct {
	Modelclientconfig ModelClientConfig `json:"ModelClientConfig,omitempty"` // Configures the timeout and maximum number of retries for processing a transform job invocation.
	Transformjobname interface{} `json:"TransformJobName,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Labelingjobarn interface{} `json:"LabelingJobArn,omitempty"`
	Dataprocessing DataProcessing `json:"DataProcessing,omitempty"` // The data structure used to specify the data to be used for inference in a batch transform job and to associate the data that is relevant to the prediction results in the output. The input filter provided allows you to exclude input data that is not needed for inference in a batch transform job. The output filter provided allows you to include input data relevant to interpreting the predictions in the output from the job. For more information, see <a href="https://docs.aws.amazon.com/sagemaker/latest/dg/batch-transform-data-processing.html">Associate Prediction Results with their Corresponding Input Records</a>.
	Modelname interface{} `json:"ModelName,omitempty"`
	Transformendtime interface{} `json:"TransformEndTime,omitempty"`
	Transformjobarn interface{} `json:"TransformJobArn,omitempty"`
	Transformjobstatus interface{} `json:"TransformJobStatus,omitempty"`
	Batchstrategy interface{} `json:"BatchStrategy,omitempty"`
	Maxpayloadinmb interface{} `json:"MaxPayloadInMB,omitempty"`
	Automljobarn interface{} `json:"AutoMLJobArn,omitempty"`
	Environment interface{} `json:"Environment,omitempty"`
	Transformresources TransformResources `json:"TransformResources,omitempty"` // Describes the resources, including ML instance types and ML instance count, to use for transform job.
	Tags interface{} `json:"Tags,omitempty"`
	Transformoutput TransformOutput `json:"TransformOutput,omitempty"` // Describes the results of a transform job.
	Transforminput TransformInput `json:"TransformInput,omitempty"` // Describes the input source of a transform job and the way the transform job consumes it.
	Transformstarttime interface{} `json:"TransformStartTime,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Experimentconfig ExperimentConfig `json:"ExperimentConfig,omitempty"` // <p>Associates a SageMaker job as a trial component with an experiment and trial. Specified when you call the following APIs:</p> <ul> <li> <p> <a href="https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateProcessingJob.html">CreateProcessingJob</a> </p> </li> <li> <p> <a href="https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTrainingJob.html">CreateTrainingJob</a> </p> </li> <li> <p> <a href="https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTransformJob.html">CreateTransformJob</a> </p> </li> </ul>
	Datacaptureconfig BatchDataCaptureConfig `json:"DataCaptureConfig,omitempty"` // Configuration to control how SageMaker captures inference data for batch transform jobs.
	Maxconcurrenttransforms interface{} `json:"MaxConcurrentTransforms,omitempty"`
}

// DescribeAutoMLJobV2Request represents the DescribeAutoMLJobV2Request schema from the OpenAPI specification
type DescribeAutoMLJobV2Request struct {
	Automljobname interface{} `json:"AutoMLJobName"`
}

// ListTrainingJobsResponse represents the ListTrainingJobsResponse schema from the OpenAPI specification
type ListTrainingJobsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Trainingjobsummaries interface{} `json:"TrainingJobSummaries"`
}

// ExperimentSummary represents the ExperimentSummary schema from the OpenAPI specification
type ExperimentSummary struct {
	Experimentarn interface{} `json:"ExperimentArn,omitempty"`
	Experimentname interface{} `json:"ExperimentName,omitempty"`
	Experimentsource ExperimentSource `json:"ExperimentSource,omitempty"` // The source of the experiment.
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Displayname interface{} `json:"DisplayName,omitempty"`
}

// ExplainerConfig represents the ExplainerConfig schema from the OpenAPI specification
type ExplainerConfig struct {
	Clarifyexplainerconfig interface{} `json:"ClarifyExplainerConfig,omitempty"`
}

// DescribeTrainingJobResponse represents the DescribeTrainingJobResponse schema from the OpenAPI specification
type DescribeTrainingJobResponse struct {
	Trainingjobarn interface{} `json:"TrainingJobArn"`
	Outputdataconfig interface{} `json:"OutputDataConfig,omitempty"`
	Creationtime interface{} `json:"CreationTime"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Trainingjobname interface{} `json:"TrainingJobName"`
	Warmpoolstatus interface{} `json:"WarmPoolStatus,omitempty"`
	Checkpointconfig CheckpointConfig `json:"CheckpointConfig,omitempty"` // Contains information about the output location for managed spot training checkpoint data.
	Stoppingcondition interface{} `json:"StoppingCondition"`
	Hyperparameters interface{} `json:"HyperParameters,omitempty"`
	Modelartifacts interface{} `json:"ModelArtifacts"`
	Debugruleconfigurations interface{} `json:"DebugRuleConfigurations,omitempty"`
	Experimentconfig ExperimentConfig `json:"ExperimentConfig,omitempty"` // <p>Associates a SageMaker job as a trial component with an experiment and trial. Specified when you call the following APIs:</p> <ul> <li> <p> <a href="https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateProcessingJob.html">CreateProcessingJob</a> </p> </li> <li> <p> <a href="https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTrainingJob.html">CreateTrainingJob</a> </p> </li> <li> <p> <a href="https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTransformJob.html">CreateTransformJob</a> </p> </li> </ul>
	Profilingstatus interface{} `json:"ProfilingStatus,omitempty"`
	Debughookconfig DebugHookConfig `json:"DebugHookConfig,omitempty"` // Configuration information for the Amazon SageMaker Debugger hook parameters, metric and tensor collections, and storage paths. To learn more about how to configure the <code>DebugHookConfig</code> parameter, see <a href="https://docs.aws.amazon.com/sagemaker/latest/dg/debugger-createtrainingjob-api.html">Use the SageMaker and Debugger Configuration API Operations to Create, Update, and Debug Your Training Job</a>.
	Trainingjobstatus interface{} `json:"TrainingJobStatus"`
	Environment interface{} `json:"Environment,omitempty"`
	Profilerconfig ProfilerConfig `json:"ProfilerConfig,omitempty"` // Configuration information for Amazon SageMaker Debugger system monitoring, framework profiling, and storage paths.
	Trainingtimeinseconds interface{} `json:"TrainingTimeInSeconds,omitempty"`
	Tensorboardoutputconfig TensorBoardOutputConfig `json:"TensorBoardOutputConfig,omitempty"` // Configuration of storage locations for the Amazon SageMaker Debugger TensorBoard output data.
	Billabletimeinseconds interface{} `json:"BillableTimeInSeconds,omitempty"`
	Resourceconfig interface{} `json:"ResourceConfig"`
	Retrystrategy interface{} `json:"RetryStrategy,omitempty"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Debugruleevaluationstatuses interface{} `json:"DebugRuleEvaluationStatuses,omitempty"`
	Profilerruleconfigurations interface{} `json:"ProfilerRuleConfigurations,omitempty"`
	Tuningjobarn interface{} `json:"TuningJobArn,omitempty"`
	Automljobarn interface{} `json:"AutoMLJobArn,omitempty"`
	Enablemanagedspottraining interface{} `json:"EnableManagedSpotTraining,omitempty"`
	Algorithmspecification interface{} `json:"AlgorithmSpecification"`
	Secondarystatus interface{} `json:"SecondaryStatus"`
	Secondarystatustransitions interface{} `json:"SecondaryStatusTransitions,omitempty"`
	Profilerruleevaluationstatuses interface{} `json:"ProfilerRuleEvaluationStatuses,omitempty"`
	Trainingendtime interface{} `json:"TrainingEndTime,omitempty"`
	Labelingjobarn interface{} `json:"LabelingJobArn,omitempty"`
	Enableintercontainertrafficencryption interface{} `json:"EnableInterContainerTrafficEncryption,omitempty"`
	Trainingstarttime interface{} `json:"TrainingStartTime,omitempty"`
	Finalmetricdatalist interface{} `json:"FinalMetricDataList,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Vpcconfig interface{} `json:"VpcConfig,omitempty"`
	Inputdataconfig interface{} `json:"InputDataConfig,omitempty"`
	Enablenetworkisolation interface{} `json:"EnableNetworkIsolation,omitempty"`
}

// TuningJobStepMetaData represents the TuningJobStepMetaData schema from the OpenAPI specification
type TuningJobStepMetaData struct {
	Arn interface{} `json:"Arn,omitempty"`
}

// ListModelExplainabilityJobDefinitionsRequest represents the ListModelExplainabilityJobDefinitionsRequest schema from the OpenAPI specification
type ListModelExplainabilityJobDefinitionsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Endpointname interface{} `json:"EndpointName,omitempty"`
}

// DeleteImageVersionRequest represents the DeleteImageVersionRequest schema from the OpenAPI specification
type DeleteImageVersionRequest struct {
	Alias interface{} `json:"Alias,omitempty"`
	Imagename interface{} `json:"ImageName"`
	Version interface{} `json:"Version,omitempty"`
}

// UpdateWorkforceRequest represents the UpdateWorkforceRequest schema from the OpenAPI specification
type UpdateWorkforceRequest struct {
	Workforcename interface{} `json:"WorkforceName"`
	Workforcevpcconfig interface{} `json:"WorkforceVpcConfig,omitempty"`
	Oidcconfig interface{} `json:"OidcConfig,omitempty"`
	Sourceipconfig interface{} `json:"SourceIpConfig,omitempty"`
}

// FeatureGroup represents the FeatureGroup schema from the OpenAPI specification
type FeatureGroup struct {
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Onlinestoreconfig OnlineStoreConfig `json:"OnlineStoreConfig,omitempty"` // <p>Use this to specify the Amazon Web Services Key Management Service (KMS) Key ID, or <code>KMSKeyId</code>, for at rest data encryption. You can turn <code>OnlineStore</code> on or off by specifying the <code>EnableOnlineStore</code> flag at General Assembly.</p> <p>The default value is <code>False</code>.</p>
	Featuregrouparn interface{} `json:"FeatureGroupArn,omitempty"`
	Featuregroupname interface{} `json:"FeatureGroupName,omitempty"`
	Offlinestoreconfig OfflineStoreConfig `json:"OfflineStoreConfig,omitempty"` // <p>The configuration of an <code>OfflineStore</code>.</p> <p>Provide an <code>OfflineStoreConfig</code> in a request to <code>CreateFeatureGroup</code> to create an <code>OfflineStore</code>.</p> <p>To encrypt an <code>OfflineStore</code> using at rest data encryption, specify Amazon Web Services Key Management Service (KMS) key ID, or <code>KMSKeyId</code>, in <code>S3StorageConfig</code>.</p>
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Offlinestorestatus OfflineStoreStatus `json:"OfflineStoreStatus,omitempty"` // The status of <code>OfflineStore</code>.
	Lastupdatestatus interface{} `json:"LastUpdateStatus,omitempty"`
	Featuredefinitions interface{} `json:"FeatureDefinitions,omitempty"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Eventtimefeaturename interface{} `json:"EventTimeFeatureName,omitempty"`
	Recordidentifierfeaturename interface{} `json:"RecordIdentifierFeatureName,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Featuregroupstatus interface{} `json:"FeatureGroupStatus,omitempty"`
}

// DriftCheckModelDataQuality represents the DriftCheckModelDataQuality schema from the OpenAPI specification
type DriftCheckModelDataQuality struct {
	Statistics interface{} `json:"Statistics,omitempty"`
	Constraints interface{} `json:"Constraints,omitempty"`
}

// ClarifyExplainerConfig represents the ClarifyExplainerConfig schema from the OpenAPI specification
type ClarifyExplainerConfig struct {
	Shapconfig interface{} `json:"ShapConfig"`
	Enableexplanations interface{} `json:"EnableExplanations,omitempty"`
	Inferenceconfig interface{} `json:"InferenceConfig,omitempty"`
}

// MonitoringNetworkConfig represents the MonitoringNetworkConfig schema from the OpenAPI specification
type MonitoringNetworkConfig struct {
	Enableintercontainertrafficencryption interface{} `json:"EnableInterContainerTrafficEncryption,omitempty"`
	Enablenetworkisolation interface{} `json:"EnableNetworkIsolation,omitempty"`
	Vpcconfig VpcConfig `json:"VpcConfig,omitempty"` // Specifies a VPC that your training jobs and hosted models have access to. Control access to and from your training and model containers by configuring the VPC. For more information, see <a href="https://docs.aws.amazon.com/sagemaker/latest/dg/host-vpc.html">Protect Endpoints by Using an Amazon Virtual Private Cloud</a> and <a href="https://docs.aws.amazon.com/sagemaker/latest/dg/train-vpc.html">Protect Training Jobs by Using an Amazon Virtual Private Cloud</a>.
}

// CreateTrainingJobResponse represents the CreateTrainingJobResponse schema from the OpenAPI specification
type CreateTrainingJobResponse struct {
	Trainingjobarn interface{} `json:"TrainingJobArn"`
}

// DeleteHubContentRequest represents the DeleteHubContentRequest schema from the OpenAPI specification
type DeleteHubContentRequest struct {
	Hubcontentversion interface{} `json:"HubContentVersion"`
	Hubname interface{} `json:"HubName"`
	Hubcontentname interface{} `json:"HubContentName"`
	Hubcontenttype interface{} `json:"HubContentType"`
}

// DescribeSubscribedWorkteamRequest represents the DescribeSubscribedWorkteamRequest schema from the OpenAPI specification
type DescribeSubscribedWorkteamRequest struct {
	Workteamarn interface{} `json:"WorkteamArn"`
}

// DataQualityJobInput represents the DataQualityJobInput schema from the OpenAPI specification
type DataQualityJobInput struct {
	Endpointinput EndpointInput `json:"EndpointInput,omitempty"` // Input object for the endpoint
	Batchtransforminput interface{} `json:"BatchTransformInput,omitempty"`
}

// ListNotebookInstancesInput represents the ListNotebookInstancesInput schema from the OpenAPI specification
type ListNotebookInstancesInput struct {
	Sortby interface{} `json:"SortBy,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
	Additionalcoderepositoryequals interface{} `json:"AdditionalCodeRepositoryEquals,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Notebookinstancelifecycleconfignamecontains interface{} `json:"NotebookInstanceLifecycleConfigNameContains,omitempty"`
	Defaultcoderepositorycontains interface{} `json:"DefaultCodeRepositoryContains,omitempty"`
	Lastmodifiedtimeafter interface{} `json:"LastModifiedTimeAfter,omitempty"`
	Lastmodifiedtimebefore interface{} `json:"LastModifiedTimeBefore,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Statusequals interface{} `json:"StatusEquals,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// SearchRecord represents the SearchRecord schema from the OpenAPI specification
type SearchRecord struct {
	Modelpackagegroup ModelPackageGroup `json:"ModelPackageGroup,omitempty"` // A group of versioned models in the model registry.
	Trainingjob interface{} `json:"TrainingJob,omitempty"`
	Experiment interface{} `json:"Experiment,omitempty"`
	Featuregroup FeatureGroup `json:"FeatureGroup,omitempty"` // Amazon SageMaker Feature Store stores features in a collection called Feature Group. A Feature Group can be visualized as a table which has rows, with a unique identifier for each row where each column in the table is a feature. In principle, a Feature Group is composed of features and values per features.
	Project interface{} `json:"Project,omitempty"`
	Modelcard interface{} `json:"ModelCard,omitempty"`
	Pipeline Pipeline `json:"Pipeline,omitempty"` // A SageMaker Model Building Pipeline instance.
	Featuremetadata interface{} `json:"FeatureMetadata,omitempty"`
	Model ModelDashboardModel `json:"Model,omitempty"` // A model displayed in the Amazon SageMaker Model Dashboard.
	Trialcomponent interface{} `json:"TrialComponent,omitempty"`
	Modelpackage ModelPackage `json:"ModelPackage,omitempty"` // A versioned model that can be deployed for SageMaker inference.
	Hyperparametertuningjob interface{} `json:"HyperParameterTuningJob,omitempty"`
	Trial interface{} `json:"Trial,omitempty"`
	Pipelineexecution PipelineExecution `json:"PipelineExecution,omitempty"` // An execution of a pipeline.
	Endpoint Endpoint `json:"Endpoint,omitempty"` // A hosted endpoint for real-time inference.
}

// ListDomainsRequest represents the ListDomainsRequest schema from the OpenAPI specification
type ListDomainsRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// ShadowModeConfig represents the ShadowModeConfig schema from the OpenAPI specification
type ShadowModeConfig struct {
	Sourcemodelvariantname interface{} `json:"SourceModelVariantName"`
	Shadowmodelvariants interface{} `json:"ShadowModelVariants"`
}

// RecommendationJobResourceLimit represents the RecommendationJobResourceLimit schema from the OpenAPI specification
type RecommendationJobResourceLimit struct {
	Maxnumberoftests interface{} `json:"MaxNumberOfTests,omitempty"`
	Maxparalleloftests interface{} `json:"MaxParallelOfTests,omitempty"`
}

// NestedFilters represents the NestedFilters schema from the OpenAPI specification
type NestedFilters struct {
	Filters interface{} `json:"Filters"`
	Nestedpropertyname interface{} `json:"NestedPropertyName"`
}

// DescribeWorkteamRequest represents the DescribeWorkteamRequest schema from the OpenAPI specification
type DescribeWorkteamRequest struct {
	Workteamname interface{} `json:"WorkteamName"`
}

// HumanLoopConfig represents the HumanLoopConfig schema from the OpenAPI specification
type HumanLoopConfig struct {
	Tasktimelimitinseconds interface{} `json:"TaskTimeLimitInSeconds,omitempty"`
	Humantaskuiarn interface{} `json:"HumanTaskUiArn"`
	Publicworkforcetaskprice PublicWorkforceTaskPrice `json:"PublicWorkforceTaskPrice,omitempty"` // <p>Defines the amount of money paid to an Amazon Mechanical Turk worker for each task performed. </p> <p>Use one of the following prices for bounding box tasks. Prices are in US dollars and should be based on the complexity of the task; the longer it takes in your initial testing, the more you should offer.</p> <ul> <li> <p>0.036</p> </li> <li> <p>0.048</p> </li> <li> <p>0.060</p> </li> <li> <p>0.072</p> </li> <li> <p>0.120</p> </li> <li> <p>0.240</p> </li> <li> <p>0.360</p> </li> <li> <p>0.480</p> </li> <li> <p>0.600</p> </li> <li> <p>0.720</p> </li> <li> <p>0.840</p> </li> <li> <p>0.960</p> </li> <li> <p>1.080</p> </li> <li> <p>1.200</p> </li> </ul> <p>Use one of the following prices for image classification, text classification, and custom tasks. Prices are in US dollars.</p> <ul> <li> <p>0.012</p> </li> <li> <p>0.024</p> </li> <li> <p>0.036</p> </li> <li> <p>0.048</p> </li> <li> <p>0.060</p> </li> <li> <p>0.072</p> </li> <li> <p>0.120</p> </li> <li> <p>0.240</p> </li> <li> <p>0.360</p> </li> <li> <p>0.480</p> </li> <li> <p>0.600</p> </li> <li> <p>0.720</p> </li> <li> <p>0.840</p> </li> <li> <p>0.960</p> </li> <li> <p>1.080</p> </li> <li> <p>1.200</p> </li> </ul> <p>Use one of the following prices for semantic segmentation tasks. Prices are in US dollars.</p> <ul> <li> <p>0.840</p> </li> <li> <p>0.960</p> </li> <li> <p>1.080</p> </li> <li> <p>1.200</p> </li> </ul> <p>Use one of the following prices for Textract AnalyzeDocument Important Form Key Amazon Augmented AI review tasks. Prices are in US dollars.</p> <ul> <li> <p>2.400 </p> </li> <li> <p>2.280 </p> </li> <li> <p>2.160 </p> </li> <li> <p>2.040 </p> </li> <li> <p>1.920 </p> </li> <li> <p>1.800 </p> </li> <li> <p>1.680 </p> </li> <li> <p>1.560 </p> </li> <li> <p>1.440 </p> </li> <li> <p>1.320 </p> </li> <li> <p>1.200 </p> </li> <li> <p>1.080 </p> </li> <li> <p>0.960 </p> </li> <li> <p>0.840 </p> </li> <li> <p>0.720 </p> </li> <li> <p>0.600 </p> </li> <li> <p>0.480 </p> </li> <li> <p>0.360 </p> </li> <li> <p>0.240 </p> </li> <li> <p>0.120 </p> </li> <li> <p>0.072 </p> </li> <li> <p>0.060 </p> </li> <li> <p>0.048 </p> </li> <li> <p>0.036 </p> </li> <li> <p>0.024 </p> </li> <li> <p>0.012 </p> </li> </ul> <p>Use one of the following prices for Rekognition DetectModerationLabels Amazon Augmented AI review tasks. Prices are in US dollars.</p> <ul> <li> <p>1.200 </p> </li> <li> <p>1.080 </p> </li> <li> <p>0.960 </p> </li> <li> <p>0.840 </p> </li> <li> <p>0.720 </p> </li> <li> <p>0.600 </p> </li> <li> <p>0.480 </p> </li> <li> <p>0.360 </p> </li> <li> <p>0.240 </p> </li> <li> <p>0.120 </p> </li> <li> <p>0.072 </p> </li> <li> <p>0.060 </p> </li> <li> <p>0.048 </p> </li> <li> <p>0.036 </p> </li> <li> <p>0.024 </p> </li> <li> <p>0.012 </p> </li> </ul> <p>Use one of the following prices for Amazon Augmented AI custom human review tasks. Prices are in US dollars.</p> <ul> <li> <p>1.200 </p> </li> <li> <p>1.080 </p> </li> <li> <p>0.960 </p> </li> <li> <p>0.840 </p> </li> <li> <p>0.720 </p> </li> <li> <p>0.600 </p> </li> <li> <p>0.480 </p> </li> <li> <p>0.360 </p> </li> <li> <p>0.240 </p> </li> <li> <p>0.120 </p> </li> <li> <p>0.072 </p> </li> <li> <p>0.060 </p> </li> <li> <p>0.048 </p> </li> <li> <p>0.036 </p> </li> <li> <p>0.024 </p> </li> <li> <p>0.012 </p> </li> </ul>
	Taskavailabilitylifetimeinseconds interface{} `json:"TaskAvailabilityLifetimeInSeconds,omitempty"`
	Workteamarn interface{} `json:"WorkteamArn"`
	Taskdescription interface{} `json:"TaskDescription"`
	Taskkeywords interface{} `json:"TaskKeywords,omitempty"`
	Tasktitle interface{} `json:"TaskTitle"`
	Taskcount interface{} `json:"TaskCount"`
}

// ListTrialsResponse represents the ListTrialsResponse schema from the OpenAPI specification
type ListTrialsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Trialsummaries interface{} `json:"TrialSummaries,omitempty"`
}

// ModelMetadataFilter represents the ModelMetadataFilter schema from the OpenAPI specification
type ModelMetadataFilter struct {
	Value interface{} `json:"Value"`
	Name interface{} `json:"Name"`
}

// ModelStepMetadata represents the ModelStepMetadata schema from the OpenAPI specification
type ModelStepMetadata struct {
	Arn interface{} `json:"Arn,omitempty"`
}

// EnableSagemakerServicecatalogPortfolioOutput represents the EnableSagemakerServicecatalogPortfolioOutput schema from the OpenAPI specification
type EnableSagemakerServicecatalogPortfolioOutput struct {
}

// CreateAppRequest represents the CreateAppRequest schema from the OpenAPI specification
type CreateAppRequest struct {
	Appname interface{} `json:"AppName"`
	Apptype interface{} `json:"AppType"`
	Domainid interface{} `json:"DomainId"`
	Resourcespec interface{} `json:"ResourceSpec,omitempty"`
	Spacename interface{} `json:"SpaceName,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Userprofilename interface{} `json:"UserProfileName,omitempty"`
}

// ModelQualityAppSpecification represents the ModelQualityAppSpecification schema from the OpenAPI specification
type ModelQualityAppSpecification struct {
	Environment interface{} `json:"Environment,omitempty"`
	Imageuri interface{} `json:"ImageUri"`
	Postanalyticsprocessorsourceuri interface{} `json:"PostAnalyticsProcessorSourceUri,omitempty"`
	Problemtype interface{} `json:"ProblemType,omitempty"`
	Recordpreprocessorsourceuri interface{} `json:"RecordPreprocessorSourceUri,omitempty"`
	Containerarguments interface{} `json:"ContainerArguments,omitempty"`
	Containerentrypoint interface{} `json:"ContainerEntrypoint,omitempty"`
}

// PredefinedMetricSpecification represents the PredefinedMetricSpecification schema from the OpenAPI specification
type PredefinedMetricSpecification struct {
	Predefinedmetrictype interface{} `json:"PredefinedMetricType,omitempty"`
}

// FeatureParameter represents the FeatureParameter schema from the OpenAPI specification
type FeatureParameter struct {
	Key interface{} `json:"Key,omitempty"`
	Value interface{} `json:"Value,omitempty"`
}

// RenderUiTemplateResponse represents the RenderUiTemplateResponse schema from the OpenAPI specification
type RenderUiTemplateResponse struct {
	Errors interface{} `json:"Errors"`
	Renderedcontent interface{} `json:"RenderedContent"`
}

// RollingUpdatePolicy represents the RollingUpdatePolicy schema from the OpenAPI specification
type RollingUpdatePolicy struct {
	Maximumbatchsize interface{} `json:"MaximumBatchSize"`
	Maximumexecutiontimeoutinseconds interface{} `json:"MaximumExecutionTimeoutInSeconds,omitempty"`
	Rollbackmaximumbatchsize interface{} `json:"RollbackMaximumBatchSize,omitempty"`
	Waitintervalinseconds interface{} `json:"WaitIntervalInSeconds"`
}

// DescribeImageVersionRequest represents the DescribeImageVersionRequest schema from the OpenAPI specification
type DescribeImageVersionRequest struct {
	Alias interface{} `json:"Alias,omitempty"`
	Imagename interface{} `json:"ImageName"`
	Version interface{} `json:"Version,omitempty"`
}

// DescribeModelCardRequest represents the DescribeModelCardRequest schema from the OpenAPI specification
type DescribeModelCardRequest struct {
	Modelcardname interface{} `json:"ModelCardName"`
	Modelcardversion interface{} `json:"ModelCardVersion,omitempty"`
}

// DeleteTagsOutput represents the DeleteTagsOutput schema from the OpenAPI specification
type DeleteTagsOutput struct {
}

// DescribeActionResponse represents the DescribeActionResponse schema from the OpenAPI specification
type DescribeActionResponse struct {
	Status interface{} `json:"Status,omitempty"`
	Lineagegrouparn interface{} `json:"LineageGroupArn,omitempty"`
	Properties interface{} `json:"Properties,omitempty"`
	Actionarn interface{} `json:"ActionArn,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Source interface{} `json:"Source,omitempty"`
	Actiontype interface{} `json:"ActionType,omitempty"`
	Createdby UserContext `json:"CreatedBy,omitempty"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Lastmodifiedby UserContext `json:"LastModifiedBy,omitempty"` // Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	Metadataproperties MetadataProperties `json:"MetadataProperties,omitempty"` // Metadata properties of the tracking entity, trial, or trial component.
	Actionname interface{} `json:"ActionName,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
}

// ListHubsResponse represents the ListHubsResponse schema from the OpenAPI specification
type ListHubsResponse struct {
	Hubsummaries interface{} `json:"HubSummaries"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// MonitoringScheduleSummary represents the MonitoringScheduleSummary schema from the OpenAPI specification
type MonitoringScheduleSummary struct {
	Creationtime interface{} `json:"CreationTime"`
	Endpointname interface{} `json:"EndpointName,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime"`
	Monitoringjobdefinitionname interface{} `json:"MonitoringJobDefinitionName,omitempty"`
	Monitoringschedulearn interface{} `json:"MonitoringScheduleArn"`
	Monitoringschedulename interface{} `json:"MonitoringScheduleName"`
	Monitoringschedulestatus interface{} `json:"MonitoringScheduleStatus"`
	Monitoringtype interface{} `json:"MonitoringType,omitempty"`
}

// ListModelCardVersionsRequest represents the ListModelCardVersionsRequest schema from the OpenAPI specification
type ListModelCardVersionsRequest struct {
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Modelcardname interface{} `json:"ModelCardName"`
	Modelcardstatus interface{} `json:"ModelCardStatus,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
}

// CreateFeatureGroupRequest represents the CreateFeatureGroupRequest schema from the OpenAPI specification
type CreateFeatureGroupRequest struct {
	Eventtimefeaturename interface{} `json:"EventTimeFeatureName"`
	Featuredefinitions interface{} `json:"FeatureDefinitions"`
	Onlinestoreconfig interface{} `json:"OnlineStoreConfig,omitempty"`
	Recordidentifierfeaturename interface{} `json:"RecordIdentifierFeatureName"`
	Offlinestoreconfig interface{} `json:"OfflineStoreConfig,omitempty"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Featuregroupname interface{} `json:"FeatureGroupName"`
	Tags interface{} `json:"Tags,omitempty"`
	Description interface{} `json:"Description,omitempty"`
}

// TrainingJobDefinition represents the TrainingJobDefinition schema from the OpenAPI specification
type TrainingJobDefinition struct {
	Hyperparameters interface{} `json:"HyperParameters,omitempty"`
	Inputdataconfig interface{} `json:"InputDataConfig"`
	Outputdataconfig interface{} `json:"OutputDataConfig"`
	Resourceconfig interface{} `json:"ResourceConfig"`
	Stoppingcondition interface{} `json:"StoppingCondition"`
	Traininginputmode string `json:"TrainingInputMode"` // <p>The training input mode that the algorithm supports. For more information about input modes, see <a href="https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html">Algorithms</a>.</p> <p> <b>Pipe mode</b> </p> <p>If an algorithm supports <code>Pipe</code> mode, Amazon SageMaker streams data directly from Amazon S3 to the container.</p> <p> <b>File mode</b> </p> <p>If an algorithm supports <code>File</code> mode, SageMaker downloads the training data from S3 to the provisioned ML storage volume, and mounts the directory to the Docker volume for the training container.</p> <p>You must provision the ML storage volume with sufficient capacity to accommodate the data downloaded from S3. In addition to the training data, the ML storage volume also stores the output model. The algorithm container uses the ML storage volume to also store intermediate information, if any.</p> <p>For distributed algorithms, training data is distributed uniformly. Your training duration is predictable if the input data objects sizes are approximately the same. SageMaker does not split the files any further for model training. If the object sizes are skewed, training won't be optimal as the data distribution is also skewed when one host in a training cluster is overloaded, thus becoming a bottleneck in training.</p> <p> <b>FastFile mode</b> </p> <p>If an algorithm supports <code>FastFile</code> mode, SageMaker streams data directly from S3 to the container with no code changes, and provides file system access to the data. Users can author their training script to interact with these files as if they were stored on disk.</p> <p> <code>FastFile</code> mode works best when the data is read sequentially. Augmented manifest files aren't supported. The startup time is lower when there are fewer files in the S3 bucket provided.</p>
}

// DeleteActionResponse represents the DeleteActionResponse schema from the OpenAPI specification
type DeleteActionResponse struct {
	Actionarn interface{} `json:"ActionArn,omitempty"`
}

// EnvironmentParameter represents the EnvironmentParameter schema from the OpenAPI specification
type EnvironmentParameter struct {
	Valuetype interface{} `json:"ValueType"`
	Key interface{} `json:"Key"`
	Value interface{} `json:"Value"`
}

// StopNotebookInstanceInput represents the StopNotebookInstanceInput schema from the OpenAPI specification
type StopNotebookInstanceInput struct {
	Notebookinstancename interface{} `json:"NotebookInstanceName"`
}

// NotificationConfiguration represents the NotificationConfiguration schema from the OpenAPI specification
type NotificationConfiguration struct {
	Notificationtopicarn interface{} `json:"NotificationTopicArn,omitempty"`
}

// BlueGreenUpdatePolicy represents the BlueGreenUpdatePolicy schema from the OpenAPI specification
type BlueGreenUpdatePolicy struct {
	Maximumexecutiontimeoutinseconds interface{} `json:"MaximumExecutionTimeoutInSeconds,omitempty"`
	Terminationwaitinseconds interface{} `json:"TerminationWaitInSeconds,omitempty"`
	Trafficroutingconfiguration interface{} `json:"TrafficRoutingConfiguration"`
}

// DeploymentStageStatusSummary represents the DeploymentStageStatusSummary schema from the OpenAPI specification
type DeploymentStageStatusSummary struct {
	Deviceselectionconfig interface{} `json:"DeviceSelectionConfig"`
	Stagename interface{} `json:"StageName"`
	Deploymentconfig interface{} `json:"DeploymentConfig"`
	Deploymentstatus interface{} `json:"DeploymentStatus"`
}

// HubS3StorageConfig represents the HubS3StorageConfig schema from the OpenAPI specification
type HubS3StorageConfig struct {
	S3outputpath interface{} `json:"S3OutputPath,omitempty"`
}

// TrialComponentArtifacts represents the TrialComponentArtifacts schema from the OpenAPI specification
type TrialComponentArtifacts struct {
}

// DeployedImage represents the DeployedImage schema from the OpenAPI specification
type DeployedImage struct {
	Resolutiontime interface{} `json:"ResolutionTime,omitempty"`
	Resolvedimage interface{} `json:"ResolvedImage,omitempty"`
	Specifiedimage interface{} `json:"SpecifiedImage,omitempty"`
}

// ListModelPackageGroupsInput represents the ListModelPackageGroupsInput schema from the OpenAPI specification
type ListModelPackageGroupsInput struct {
	Namecontains interface{} `json:"NameContains,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// ListAutoMLJobsResponse represents the ListAutoMLJobsResponse schema from the OpenAPI specification
type ListAutoMLJobsResponse struct {
	Automljobsummaries interface{} `json:"AutoMLJobSummaries"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DescribeModelBiasJobDefinitionResponse represents the DescribeModelBiasJobDefinitionResponse schema from the OpenAPI specification
type DescribeModelBiasJobDefinitionResponse struct {
	Jobresources MonitoringResources `json:"JobResources"` // Identifies the resources to deploy for a monitoring job.
	Modelbiasappspecification interface{} `json:"ModelBiasAppSpecification"`
	Modelbiasjoboutputconfig MonitoringOutputConfig `json:"ModelBiasJobOutputConfig"` // The output configuration for monitoring jobs.
	Networkconfig interface{} `json:"NetworkConfig,omitempty"`
	Stoppingcondition MonitoringStoppingCondition `json:"StoppingCondition,omitempty"` // A time limit for how long the monitoring job is allowed to run before stopping.
	Jobdefinitionarn interface{} `json:"JobDefinitionArn"`
	Modelbiasbaselineconfig interface{} `json:"ModelBiasBaselineConfig,omitempty"`
	Jobdefinitionname interface{} `json:"JobDefinitionName"`
	Modelbiasjobinput interface{} `json:"ModelBiasJobInput"`
	Rolearn interface{} `json:"RoleArn"`
	Creationtime interface{} `json:"CreationTime"`
}

// MetricData represents the MetricData schema from the OpenAPI specification
type MetricData struct {
	Value interface{} `json:"Value,omitempty"`
	Metricname interface{} `json:"MetricName,omitempty"`
	Timestamp interface{} `json:"Timestamp,omitempty"`
}

// DeleteAlgorithmInput represents the DeleteAlgorithmInput schema from the OpenAPI specification
type DeleteAlgorithmInput struct {
	Algorithmname interface{} `json:"AlgorithmName"`
}

// TimeSeriesTransformations represents the TimeSeriesTransformations schema from the OpenAPI specification
type TimeSeriesTransformations struct {
	Aggregation interface{} `json:"Aggregation,omitempty"`
	Filling interface{} `json:"Filling,omitempty"`
}

// PropertyNameSuggestion represents the PropertyNameSuggestion schema from the OpenAPI specification
type PropertyNameSuggestion struct {
	Propertyname interface{} `json:"PropertyName,omitempty"`
}

// VariantProperty represents the VariantProperty schema from the OpenAPI specification
type VariantProperty struct {
	Variantpropertytype interface{} `json:"VariantPropertyType"`
}

// CreateTrialRequest represents the CreateTrialRequest schema from the OpenAPI specification
type CreateTrialRequest struct {
	Displayname interface{} `json:"DisplayName,omitempty"`
	Experimentname interface{} `json:"ExperimentName"`
	Metadataproperties MetadataProperties `json:"MetadataProperties,omitempty"` // Metadata properties of the tracking entity, trial, or trial component.
	Tags interface{} `json:"Tags,omitempty"`
	Trialname interface{} `json:"TrialName"`
}

// KernelSpec represents the KernelSpec schema from the OpenAPI specification
type KernelSpec struct {
	Name interface{} `json:"Name"`
	Displayname interface{} `json:"DisplayName,omitempty"`
}

// ModelDashboardEndpoint represents the ModelDashboardEndpoint schema from the OpenAPI specification
type ModelDashboardEndpoint struct {
	Lastmodifiedtime interface{} `json:"LastModifiedTime"`
	Creationtime interface{} `json:"CreationTime"`
	Endpointarn interface{} `json:"EndpointArn"`
	Endpointname interface{} `json:"EndpointName"`
	Endpointstatus interface{} `json:"EndpointStatus"`
}

// DeleteTagsInput represents the DeleteTagsInput schema from the OpenAPI specification
type DeleteTagsInput struct {
	Resourcearn interface{} `json:"ResourceArn"`
	Tagkeys interface{} `json:"TagKeys"`
}

// ListResourceCatalogsResponse represents the ListResourceCatalogsResponse schema from the OpenAPI specification
type ListResourceCatalogsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Resourcecatalogs interface{} `json:"ResourceCatalogs,omitempty"`
}

// HubInfo represents the HubInfo schema from the OpenAPI specification
type HubInfo struct {
	Hubstatus interface{} `json:"HubStatus"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime"`
	Creationtime interface{} `json:"CreationTime"`
	Hubarn interface{} `json:"HubArn"`
	Hubdescription interface{} `json:"HubDescription,omitempty"`
	Hubdisplayname interface{} `json:"HubDisplayName,omitempty"`
	Hubname interface{} `json:"HubName"`
	Hubsearchkeywords interface{} `json:"HubSearchKeywords,omitempty"`
}

// ParallelismConfiguration represents the ParallelismConfiguration schema from the OpenAPI specification
type ParallelismConfiguration struct {
	Maxparallelexecutionsteps interface{} `json:"MaxParallelExecutionSteps"`
}

// ProcessingOutputConfig represents the ProcessingOutputConfig schema from the OpenAPI specification
type ProcessingOutputConfig struct {
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Outputs interface{} `json:"Outputs"`
}

// CreateAutoMLJobResponse represents the CreateAutoMLJobResponse schema from the OpenAPI specification
type CreateAutoMLJobResponse struct {
	Automljobarn interface{} `json:"AutoMLJobArn"`
}

// MonitoringConstraintsResource represents the MonitoringConstraintsResource schema from the OpenAPI specification
type MonitoringConstraintsResource struct {
	S3uri interface{} `json:"S3Uri,omitempty"`
}

// DeleteEdgeDeploymentStageRequest represents the DeleteEdgeDeploymentStageRequest schema from the OpenAPI specification
type DeleteEdgeDeploymentStageRequest struct {
	Edgedeploymentplanname interface{} `json:"EdgeDeploymentPlanName"`
	Stagename interface{} `json:"StageName"`
}

// ModelPackageSummary represents the ModelPackageSummary schema from the OpenAPI specification
type ModelPackageSummary struct {
	Modelpackagename interface{} `json:"ModelPackageName"`
	Modelpackagestatus interface{} `json:"ModelPackageStatus"`
	Modelpackageversion interface{} `json:"ModelPackageVersion,omitempty"`
	Creationtime interface{} `json:"CreationTime"`
	Modelapprovalstatus interface{} `json:"ModelApprovalStatus,omitempty"`
	Modelpackagearn interface{} `json:"ModelPackageArn"`
	Modelpackagedescription interface{} `json:"ModelPackageDescription,omitempty"`
	Modelpackagegroupname interface{} `json:"ModelPackageGroupName,omitempty"`
}

// ListArtifactsRequest represents the ListArtifactsRequest schema from the OpenAPI specification
type ListArtifactsRequest struct {
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Sourceuri interface{} `json:"SourceUri,omitempty"`
	Artifacttype interface{} `json:"ArtifactType,omitempty"`
	Createdafter interface{} `json:"CreatedAfter,omitempty"`
	Createdbefore interface{} `json:"CreatedBefore,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
}

// ProductionVariant represents the ProductionVariant schema from the OpenAPI specification
type ProductionVariant struct {
	Variantname interface{} `json:"VariantName"`
	Volumesizeingb interface{} `json:"VolumeSizeInGB,omitempty"`
	Instancetype interface{} `json:"InstanceType,omitempty"`
	Serverlessconfig interface{} `json:"ServerlessConfig,omitempty"`
	Acceleratortype interface{} `json:"AcceleratorType,omitempty"`
	Containerstartuphealthchecktimeoutinseconds interface{} `json:"ContainerStartupHealthCheckTimeoutInSeconds,omitempty"`
	Coredumpconfig interface{} `json:"CoreDumpConfig,omitempty"`
	Initialinstancecount interface{} `json:"InitialInstanceCount,omitempty"`
	Modeldatadownloadtimeoutinseconds interface{} `json:"ModelDataDownloadTimeoutInSeconds,omitempty"`
	Modelname interface{} `json:"ModelName"`
	Enablessmaccess interface{} `json:"EnableSSMAccess,omitempty"`
	Initialvariantweight interface{} `json:"InitialVariantWeight,omitempty"`
}

// IamIdentity represents the IamIdentity schema from the OpenAPI specification
type IamIdentity struct {
	Arn interface{} `json:"Arn,omitempty"`
	Principalid interface{} `json:"PrincipalId,omitempty"`
	Sourceidentity interface{} `json:"SourceIdentity,omitempty"`
}

// FinalHyperParameterTuningJobObjectiveMetric represents the FinalHyperParameterTuningJobObjectiveMetric schema from the OpenAPI specification
type FinalHyperParameterTuningJobObjectiveMetric struct {
	Metricname interface{} `json:"MetricName"`
	TypeField interface{} `json:"Type,omitempty"`
	Value interface{} `json:"Value"`
}

// StopMonitoringScheduleRequest represents the StopMonitoringScheduleRequest schema from the OpenAPI specification
type StopMonitoringScheduleRequest struct {
	Monitoringschedulename interface{} `json:"MonitoringScheduleName"`
}

// FileSystemConfig represents the FileSystemConfig schema from the OpenAPI specification
type FileSystemConfig struct {
	Defaultgid interface{} `json:"DefaultGid,omitempty"`
	Defaultuid interface{} `json:"DefaultUid,omitempty"`
	Mountpath interface{} `json:"MountPath,omitempty"`
}

// ModelVariantConfigSummary represents the ModelVariantConfigSummary schema from the OpenAPI specification
type ModelVariantConfigSummary struct {
	Infrastructureconfig interface{} `json:"InfrastructureConfig"`
	Modelname interface{} `json:"ModelName"`
	Status interface{} `json:"Status"`
	Variantname interface{} `json:"VariantName"`
}

// RegisterModelStepMetadata represents the RegisterModelStepMetadata schema from the OpenAPI specification
type RegisterModelStepMetadata struct {
	Arn interface{} `json:"Arn,omitempty"`
}

// ListAliasesRequest represents the ListAliasesRequest schema from the OpenAPI specification
type ListAliasesRequest struct {
	Alias interface{} `json:"Alias,omitempty"`
	Imagename interface{} `json:"ImageName"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Version interface{} `json:"Version,omitempty"`
}

// RecommendationJobContainerConfig represents the RecommendationJobContainerConfig schema from the OpenAPI specification
type RecommendationJobContainerConfig struct {
	Payloadconfig interface{} `json:"PayloadConfig,omitempty"`
	Task interface{} `json:"Task,omitempty"`
	Framework interface{} `json:"Framework,omitempty"`
	Supportedinstancetypes interface{} `json:"SupportedInstanceTypes,omitempty"`
	Datainputconfig interface{} `json:"DataInputConfig,omitempty"`
	Domain interface{} `json:"Domain,omitempty"`
	Nearestmodelname interface{} `json:"NearestModelName,omitempty"`
	Supportedendpointtype interface{} `json:"SupportedEndpointType,omitempty"`
	Frameworkversion interface{} `json:"FrameworkVersion,omitempty"`
}

// AlgorithmValidationSpecification represents the AlgorithmValidationSpecification schema from the OpenAPI specification
type AlgorithmValidationSpecification struct {
	Validationprofiles interface{} `json:"ValidationProfiles"`
	Validationrole interface{} `json:"ValidationRole"`
}

// ProductionVariantCoreDumpConfig represents the ProductionVariantCoreDumpConfig schema from the OpenAPI specification
type ProductionVariantCoreDumpConfig struct {
	Destinations3uri interface{} `json:"DestinationS3Uri"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
}

// LineageGroupSummary represents the LineageGroupSummary schema from the OpenAPI specification
type LineageGroupSummary struct {
	Lineagegrouparn interface{} `json:"LineageGroupArn,omitempty"`
	Lineagegroupname interface{} `json:"LineageGroupName,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Displayname interface{} `json:"DisplayName,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
}

// TrialComponentSourceDetail represents the TrialComponentSourceDetail schema from the OpenAPI specification
type TrialComponentSourceDetail struct {
	Processingjob interface{} `json:"ProcessingJob,omitempty"`
	Sourcearn interface{} `json:"SourceArn,omitempty"`
	Trainingjob interface{} `json:"TrainingJob,omitempty"`
	Transformjob interface{} `json:"TransformJob,omitempty"`
}

// EndpointConfigSummary represents the EndpointConfigSummary schema from the OpenAPI specification
type EndpointConfigSummary struct {
	Endpointconfigarn interface{} `json:"EndpointConfigArn"`
	Endpointconfigname interface{} `json:"EndpointConfigName"`
	Creationtime interface{} `json:"CreationTime"`
}

// TuningJobCompletionCriteria represents the TuningJobCompletionCriteria schema from the OpenAPI specification
type TuningJobCompletionCriteria struct {
	Targetobjectivemetricvalue interface{} `json:"TargetObjectiveMetricValue,omitempty"`
	Bestobjectivenotimproving interface{} `json:"BestObjectiveNotImproving,omitempty"`
	Convergencedetected interface{} `json:"ConvergenceDetected,omitempty"`
}

// DescribeDomainRequest represents the DescribeDomainRequest schema from the OpenAPI specification
type DescribeDomainRequest struct {
	Domainid interface{} `json:"DomainId"`
}

// LabelingJobAlgorithmsConfig represents the LabelingJobAlgorithmsConfig schema from the OpenAPI specification
type LabelingJobAlgorithmsConfig struct {
	Initialactivelearningmodelarn interface{} `json:"InitialActiveLearningModelArn,omitempty"`
	Labelingjobalgorithmspecificationarn interface{} `json:"LabelingJobAlgorithmSpecificationArn"`
	Labelingjobresourceconfig interface{} `json:"LabelingJobResourceConfig,omitempty"`
}

// ModelMetadataSearchExpression represents the ModelMetadataSearchExpression schema from the OpenAPI specification
type ModelMetadataSearchExpression struct {
	Filters interface{} `json:"Filters,omitempty"`
}

// DefaultSpaceSettings represents the DefaultSpaceSettings schema from the OpenAPI specification
type DefaultSpaceSettings struct {
	Kernelgatewayappsettings KernelGatewayAppSettings `json:"KernelGatewayAppSettings,omitempty"` // The KernelGateway app settings.
	Securitygroups interface{} `json:"SecurityGroups,omitempty"`
	Executionrole interface{} `json:"ExecutionRole,omitempty"`
	Jupyterserverappsettings JupyterServerAppSettings `json:"JupyterServerAppSettings,omitempty"` // The JupyterServer app settings.
}

// UpdateFeatureGroupRequest represents the UpdateFeatureGroupRequest schema from the OpenAPI specification
type UpdateFeatureGroupRequest struct {
	Featureadditions interface{} `json:"FeatureAdditions,omitempty"`
	Featuregroupname interface{} `json:"FeatureGroupName"`
	Onlinestoreconfig interface{} `json:"OnlineStoreConfig,omitempty"`
}

// TextClassificationJobConfig represents the TextClassificationJobConfig schema from the OpenAPI specification
type TextClassificationJobConfig struct {
	Completioncriteria interface{} `json:"CompletionCriteria,omitempty"`
	Contentcolumn interface{} `json:"ContentColumn"`
	Targetlabelcolumn interface{} `json:"TargetLabelColumn"`
}

// ClarifyShapBaselineConfig represents the ClarifyShapBaselineConfig schema from the OpenAPI specification
type ClarifyShapBaselineConfig struct {
	Mimetype interface{} `json:"MimeType,omitempty"`
	Shapbaseline interface{} `json:"ShapBaseline,omitempty"`
	Shapbaselineuri interface{} `json:"ShapBaselineUri,omitempty"`
}

// TrainingJobSummary represents the TrainingJobSummary schema from the OpenAPI specification
type TrainingJobSummary struct {
	Trainingendtime interface{} `json:"TrainingEndTime,omitempty"`
	Trainingjobarn interface{} `json:"TrainingJobArn"`
	Trainingjobname interface{} `json:"TrainingJobName"`
	Trainingjobstatus interface{} `json:"TrainingJobStatus"`
	Warmpoolstatus interface{} `json:"WarmPoolStatus,omitempty"`
	Creationtime interface{} `json:"CreationTime"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
}

// ListProjectsInput represents the ListProjectsInput schema from the OpenAPI specification
type ListProjectsInput struct {
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Creationtimeafter interface{} `json:"CreationTimeAfter,omitempty"`
	Creationtimebefore interface{} `json:"CreationTimeBefore,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DeleteModelExplainabilityJobDefinitionRequest represents the DeleteModelExplainabilityJobDefinitionRequest schema from the OpenAPI specification
type DeleteModelExplainabilityJobDefinitionRequest struct {
	Jobdefinitionname interface{} `json:"JobDefinitionName"`
}

// ListCodeRepositoriesOutput represents the ListCodeRepositoriesOutput schema from the OpenAPI specification
type ListCodeRepositoriesOutput struct {
	Coderepositorysummarylist interface{} `json:"CodeRepositorySummaryList"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CreateImageRequest represents the CreateImageRequest schema from the OpenAPI specification
type CreateImageRequest struct {
	Rolearn interface{} `json:"RoleArn"`
	Tags interface{} `json:"Tags,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Displayname interface{} `json:"DisplayName,omitempty"`
	Imagename interface{} `json:"ImageName"`
}

// ListDomainsResponse represents the ListDomainsResponse schema from the OpenAPI specification
type ListDomainsResponse struct {
	Domains interface{} `json:"Domains,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ModelExplainabilityBaselineConfig represents the ModelExplainabilityBaselineConfig schema from the OpenAPI specification
type ModelExplainabilityBaselineConfig struct {
	Baseliningjobname interface{} `json:"BaseliningJobName,omitempty"`
	Constraintsresource MonitoringConstraintsResource `json:"ConstraintsResource,omitempty"` // The constraints resource for a monitoring job.
}

// DeletePipelineResponse represents the DeletePipelineResponse schema from the OpenAPI specification
type DeletePipelineResponse struct {
	Pipelinearn interface{} `json:"PipelineArn,omitempty"`
}

// FeatureMetadata represents the FeatureMetadata schema from the OpenAPI specification
type FeatureMetadata struct {
	Parameters interface{} `json:"Parameters,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Featuregrouparn interface{} `json:"FeatureGroupArn,omitempty"`
	Featuregroupname interface{} `json:"FeatureGroupName,omitempty"`
	Featurename interface{} `json:"FeatureName,omitempty"`
	Featuretype interface{} `json:"FeatureType,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
}

// CreateFeatureGroupResponse represents the CreateFeatureGroupResponse schema from the OpenAPI specification
type CreateFeatureGroupResponse struct {
	Featuregrouparn interface{} `json:"FeatureGroupArn"`
}
