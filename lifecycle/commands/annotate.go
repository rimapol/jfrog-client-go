package commands

import (
	"errors"
	"fmt"
	"github.com/jfrog/gofrog/log"
	"github.com/jfrog/jfrog-cli-core/v2/utils/config"
	"github.com/jfrog/jfrog-client-go/artifactory/services/utils"
	"github.com/jfrog/jfrog-client-go/lifecycle"
	"github.com/jfrog/jfrog-client-go/lifecycle/services"
)

const (
	manifestName     = "release-bundle.json.evd"
	releaseBundlesV2 = "release-bundles-v2"
)

type ReleaseBundleAnnotateCommand struct {
	releaseBundleCmd
	tag   string
	props string
}

func NewReleaseBundleAnnotateCommand() *ReleaseBundleAnnotateCommand {
	return &ReleaseBundleAnnotateCommand{}
}

func (rba *ReleaseBundleAnnotateCommand) SetServerDetails(serverDetails *config.ServerDetails) *ReleaseBundleAnnotateCommand {
	rba.serverDetails = serverDetails
	return rba
}

func (rba *ReleaseBundleAnnotateCommand) SetReleaseBundleName(releaseBundleName string) *ReleaseBundleAnnotateCommand {
	rba.releaseBundleName = releaseBundleName
	return rba
}

func (rba *ReleaseBundleAnnotateCommand) SetReleaseBundleVersion(releaseBundleVersion string) *ReleaseBundleAnnotateCommand {
	rba.releaseBundleVersion = releaseBundleVersion
	return rba
}

func (rba *ReleaseBundleAnnotateCommand) SetReleaseBundleProject(rbProjectKey string) *ReleaseBundleAnnotateCommand {
	rba.rbProjectKey = rbProjectKey
	return rba
}

func (rba *ReleaseBundleAnnotateCommand) SetTag(tag string) *ReleaseBundleAnnotateCommand {
	rba.tag = tag
	return rba
}

func (rba *ReleaseBundleAnnotateCommand) SetProps(props string) *ReleaseBundleAnnotateCommand {
	rba.props = props
	return rba
}

func (rba *ReleaseBundleAnnotateCommand) ServerDetails() (*config.ServerDetails, error) {
	return rba.serverDetails, nil
}

func (rba *ReleaseBundleAnnotateCommand) CommandName() string {
	return "rb_annotate"
}

func (rba *ReleaseBundleAnnotateCommand) Run() error {
	if err := validateArtifactoryVersionSupported(rba.serverDetails); err != nil {
		return err
	}

	servicesManager, rbDetails, queryParams, err := rba.getPrerequisites()
	if err != nil {
		return err
	}

	queryParams.Async = false

	err = rba.annotateReleaseBundle(servicesManager, rbDetails, queryParams, rba.tag, rba.props)
	if err != nil {
		return err
	}
	log.Info("Success, release bundle: ", rbDetails.ReleaseBundleName+"/"+rbDetails.ReleaseBundleVersion)
	return nil
}

func (rba *ReleaseBundleAnnotateCommand) annotateReleaseBundle(manager *lifecycle.LifecycleServicesManager,
	details services.ReleaseBundleDetails, params services.CommonOptionalQueryParams, tag, properties string) error {
	err := validateParameters(tag, properties)
	if err != nil {
		return err
	}

	return manager.AnnotateReleaseBundle(details, rba.serverDetails.ArtifactoryUrl,
		BuildAnnotationOperationParams(tag, properties, details, params))
}

func BuildAnnotationOperationParams(tag, properties string, details services.ReleaseBundleDetails,
	params services.CommonOptionalQueryParams) services.AnnotateOperationParams {
	annotateOperationParams := services.AnnotateOperationParams{
		RbTag: services.RbAnnotationTag{
			Tag: tag,
		},
		RbProps: services.RbAnnotationProps{
			Properties: buildProps(properties),
		},
		RbDetails:   details,
		QueryParams: params,
		PropertyParams: services.CommonPropParams{
			Path:    buildManifestPath(details.ReleaseBundleName, details.ReleaseBundleVersion),
			RepoKey: buildRepoKey(params.ProjectKey),
		},
	}
	return annotateOperationParams
}

func buildProps(properties string) map[string][]string {
	if properties == "" {
		return make(map[string][]string)
	}
	props, err := utils.ParseProperties(properties)
	if err != nil {
		return make(map[string][]string)
	}
	return props.ToMap()
}

func buildRepoKey(project string) string {
	if project == "" || project == "default" {
		return releaseBundlesV2
	}
	return fmt.Sprintf("%s-%s", project, releaseBundlesV2)
}

func buildManifestPath(bundleName, bundleVersion string) string {
	return fmt.Sprintf("%s/%s/%s", bundleName, bundleVersion, manifestName)
}

func validateParameters(tag, properties string) error {
	if tag == "" && properties == "" {
		return errors.New("both tag and properties parameters are empty")
	}
	return nil
}
