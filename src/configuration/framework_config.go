package configuration

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type TracerConfigurations struct {
	Record           RecordConfiguration            `yaml:"record,omitempty"`
	Report           ReportConfiguration            `yaml:"report,omitempty"`
	RulesBreakers    []RulesBreakersConfiguration   `yaml:"rules_breakers"`
	BreakingProfiles []BreakingProfileConfiguration `yaml:"breaking_profiles"`
}

type RecordConfiguration struct {
	ProgConfig      ProgConfiguration `yaml:"prog_config,omitempty"`
	BreakingProfile string            `yaml:"breaking_profile,omitempty"`
}

type ReportConfiguration struct {
	Mode       string            `yaml:"mode,omitempty"`
	ProgConfig ProgConfiguration `yaml:"prog_config,omitempty"`
	Iterations int               `yaml:"iterations,omitempty"`
}

type ProgConfiguration struct {
	ProgBin  string   `yaml:"prog_bin,omitempty"`
	ProgArgs []string `yaml:"prog_args,omitempty"`
}

type RulesBreakersConfiguration struct {
	Rule   string `yaml:"rule,omitempty"`
	RuleId int    `yaml:"rule_id,omitempty"`
}

type BreakingProfileConfiguration struct {
	Name             string  `yaml:"name,omitempty"`
	Sequence         [][]int `yaml:"sequence,omitempty"`
	RollbackSequence []int   `yaml:"rollback_sequence,omitempty"`
	Ratio            int     `yaml:"ratio,omitempty"`
	Limit            int     `yaml:"limit,omitempty"`
}

func (tc *TracerConfigurations) UnmarshalYAML(configFile string) error {

	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, tc)

	return err
}
