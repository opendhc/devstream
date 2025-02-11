package python

import (
	"github.com/devstream-io/devstream/internal/pkg/plugininstaller"
	"github.com/devstream-io/devstream/internal/pkg/plugininstaller/github"
)

// Delete remove GitHub Actions workflows.
func Delete(options map[string]interface{}) (bool, error) {
	runner := &plugininstaller.Runner{
		PreExecuteOperations: []plugininstaller.MutableOperation{
			github.Validate,
			github.BuildWorkFlowsWrapper(workflows),
		},
		ExecuteOperations: []plugininstaller.BaseOperation{
			deleteDockerHubInfoForPush,
			github.ProcessAction("delete"),
		},
	}

	_, err := runner.Execute(plugininstaller.RawOptions(options))
	if err != nil {
		return false, err
	}
	return true, nil

}
