package vultr

import (
	"github.com/hashicorp/packer/common"
	"github.com/hashicorp/packer/helper/communicator"
	"github.com/hashicorp/packer/template/interpolate"
)

// Config ...
type Config struct {
	common.PackerConfig `mapstructure:",squash"`
	Comm                communicator.Config `mapstructure:",squash"`
	APIKey              string              `mapstructure:"api_key"`
	Description         string              `mapstructure:"description"`
	RegionID            int                 `mapstructure:"region_id"`
	RegionName          string              `mapstructure:"region_name"`
	RegionCode          string              `mapstructure:"region_code"`
	PlanID              int                 `mapstructure:"plan_id"`
	PlanName            string              `mapstructure:"plan_name"`
	OSID                int                 `mapstructure:"os_id"`
	OSName              string              `mapstructure:"os_name"`
	ScriptID            int                 `mapstructure:"script_id"`
	SnapshotID          string              `mapstructure:"snapshot_id"`
	ShutdownCommand     string              `mapstructure:"shutdown_command"`
	interCtx            interpolate.Context
}