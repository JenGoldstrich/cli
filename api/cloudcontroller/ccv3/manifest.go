package ccv3

import (
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/internal"
)

// GetApplicationManifest returns a (YAML) manifest for an application and its
// underlying processes.
func (client *Client) GetApplicationManifest(appGUID string) ([]byte, Warnings, error) {
	bytes, warnings, err := client.MakeRequestReceiveRaw(
		internal.GetApplicationManifestRequest,
		internal.Params{"app_guid": appGUID},
		"application/x-yaml",
	)

	return bytes, warnings, err
}

func (client *Client) GetSpaceManifestDiff(spaceGUID string, newManifest []byte) ([]byte, Warnings, error) {
	bytes, warnings, err := client.MakeRequestSendReceiveRaw(
		internal.GetSpaceManifestDiffRequest,
		internal.Params{"space_guid": spaceGUID},
		newManifest,
		"application/x-yaml",
		"application/json",
	)

	return bytes, warnings, err
}
