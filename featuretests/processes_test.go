// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package featuretests

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/juju/cmd"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"
	"gopkg.in/juju/charm.v5"
	goyaml "gopkg.in/yaml.v1"

	"github.com/juju/juju/cmd/envcmd"
	cmdaction "github.com/juju/juju/cmd/juju/action"
	cmdjuju "github.com/juju/juju/cmd/juju/commands"
	cmdenv "github.com/juju/juju/cmd/juju/environment"
	cmdmachine "github.com/juju/juju/cmd/juju/machine"
	cmdservice "github.com/juju/juju/cmd/juju/service"
	"github.com/juju/juju/component/all"
	jjj "github.com/juju/juju/juju"
	"github.com/juju/juju/process"
	"github.com/juju/juju/testcharms"
	coretesting "github.com/juju/juju/testing"
)

func initProcessesSuites() {
	if err := all.RegisterForServer(); err != nil {
		panic(err)
	}

	gc.Suite(&processesHookContextSuite{})
	gc.Suite(&processesWorkerSuite{})
	gc.Suite(&processesCmdJujuSuite{})
}

var (
	repoDir         = testcharms.Repo.Path()
	procsEnv        *procsEnviron
	procsSuiteCount = 3
)

type processesBaseSuite struct {
	env *procsEnviron
}

func (s *processesBaseSuite) SetUpSuite(c *gc.C) {
	// TODO(ericsnow) Run in a test-only local provider
	//  (see https://github.com/juju/docs/issues/35).
	s.env = newProcsEnv(c, "local", procsSuiteCount)
	s.env.bootstrap(c)
}

func (s *processesBaseSuite) TearDownSuite(c *gc.C) {
	if s.env != nil {
		s.env.destroy(c)
	}
}

func (s *processesBaseSuite) SetUpTest(c *gc.C) {
}

func (s *processesBaseSuite) TearDownTest(c *gc.C) {
	if c.Failed() {
		s.env.markFailure(c, s)
	}
}

type processesHookContextSuite struct {
	processesBaseSuite
}

func (s *processesHookContextSuite) TestHookLifecycle(c *gc.C) {
	// start: info, launch, info
	// config-changed: info, set-status, info
	// stop: info, destroy, info

	svc := s.env.addService(c, "proc-hooks", "a-service")

	// Add/start the unit.

	unit := svc.deploy(c, "myproc", "xyz123", "running")

	unit.checkState(c, []process.Info{{
		Process: charm.Process{
			Name: "myproc",
			Type: "myplugin",
		},
		Details: process.Details{
			ID: "xyz123",
			Status: process.PluginStatus{
				Label: "running",
			},
		},
	}})
	unit.checkPluginLog(c, []string{
		`myproc	xyz123	running		added`,
		`myproc	xyz123	running	{"Name":"myproc","Description":"","Type":"myplugin","TypeOptions":{"critical":"true"},"Command":"run-server","Image":"web-server","Ports":[{"External":8080,"Internal":80,"Endpoint":""},{"External":8081,"Internal":443,"Endpoint":""}],"Volumes":[{"ExternalMount":"/var/some-server/html","InternalMount":"/usr/share/some-server/html","Mode":"ro","Name":""},{"ExternalMount":"/var/some-server/conf","InternalMount":"/etc/some-server","Mode":"ro","Name":""}],"EnvVars":{"IMPORTANT":"some value"}}	definition set`,
	})

	// Change the config.

	// TODO(ericsnow) Implement once proc-set-status exists...

	// Stop the unit.

	if c.Failed() {
		return
	}
	unit.destroy(c)

	unit.checkState(c, nil)
	unit.checkPluginLog(c, []string{
		`myproc	xyz123	running		added`,
		`myproc	xyz123	running	{"Name":"myproc","Description":"","Type":"myplugin","TypeOptions":{"critical":"true"},"Command":"run-server","Image":"web-server","Ports":[{"External":8080,"Internal":80,"Endpoint":""},{"External":8081,"Internal":443,"Endpoint":""}],"Volumes":[{"ExternalMount":"/var/some-server/html","InternalMount":"/usr/share/some-server/html","Mode":"ro","Name":""},{"ExternalMount":"/var/some-server/conf","InternalMount":"/etc/some-server","Mode":"ro","Name":""}],"EnvVars":{"IMPORTANT":"some value"}}	definition set`,
		`myproc	xyz123	running		removed`,
	})
}

// TODO(ericsnow) Add a test specifically for each supported plugin
// (e.g. docker)?

func (s *processesHookContextSuite) TestRegister(c *gc.C) {
	//s.addUnit(c, "proc-actions", "a-service")
	// TODO(ericsnow) Finish!
	c.Skip("not finished")
}

func (s *processesHookContextSuite) TestLaunch(c *gc.C) {
	// TODO(ericsnow) Finish!
	c.Skip("not finished")
}

func (s *processesHookContextSuite) TestInfo(c *gc.C) {
	// TODO(ericsnow) Finish!
	c.Skip("not finished")
}

func (s *processesHookContextSuite) TestUnregister(c *gc.C) {
	// TODO(ericsnow) Finish!
	c.Skip("not finished")
}

func (s *processesHookContextSuite) TestDestroy(c *gc.C) {
	// TODO(ericsnow) Finish!
	c.Skip("not finished")
}

type processesWorkerSuite struct {
	processesBaseSuite
}

func (s *processesWorkerSuite) TestSetStatus(c *gc.C) {
	// TODO(ericsnow) Finish!
	c.Skip("not finished")
}

func (s *processesWorkerSuite) TestCleanUp(c *gc.C) {
	// TODO(ericsnow) Finish!
	c.Skip("not finished")
}

type processesCmdJujuSuite struct {
	processesBaseSuite
}

func (s *processesCmdJujuSuite) TestStatus(c *gc.C) {
	// TODO(ericsnow) Finish!
	c.Skip("not finished")
}

type procsEnviron struct {
	name     string
	machine  string
	refCount int
}

func newProcsEnv(c *gc.C, envName string, suiteCount int) *procsEnviron {
	if procsEnv != nil {
		c.Assert(procsEnv.name, gc.Equals, envName)
		return procsEnv
	}
	return &procsEnviron{
		name:     envName,
		machine:  "1",
		refCount: suiteCount,
	}
}

func (env *procsEnviron) markFailure(c *gc.C, s interface{}) {
	env.refCount = -1
}

func (env *procsEnviron) run(c *gc.C, cmd string, args ...string) string {
	envArg := "--environment=" + env.name
	if cmd == "destroy-environment" {
		envArg = env.name
	}
	args = append([]string{envArg}, args...)
	c.Logf(" COMMAND: juju %s %s", cmd, strings.Join(args, " "))

	command := lookUpCommand(cmd)
	ctx, err := coretesting.RunCommand(c, command, args...)
	c.Assert(err, jc.ErrorIsNil)

	return strings.TrimSpace(coretesting.Stdout(ctx))
}

func initJuju(c *gc.C) {
	err := jjj.InitJujuHome()
	c.Assert(err, jc.ErrorIsNil)
}

// TODO(ericsnow) Instead, directly access the command registry...
func lookUpCommand(cmd string) cmd.Command {
	switch cmd {
	case "bootstrap":
		return envcmd.Wrap(&cmdjuju.BootstrapCommand{})
	case "environment set":
		return envcmd.Wrap(&cmdenv.SetCommand{})
	case "destroy-environment":
		return &cmdjuju.DestroyEnvironmentCommand{}
	case "status":
		return envcmd.Wrap(&cmdjuju.StatusCommand{})
	case "add-machine":
		return envcmd.Wrap(&cmdmachine.AddCommand{})
	case "deploy":
		return envcmd.Wrap(&cmdjuju.DeployCommand{})
	case "service set":
		return envcmd.Wrap(&cmdservice.SetCommand{})
	case "service add-unit":
		return envcmd.Wrap(&cmdservice.AddUnitCommand{})
	case "destroy-unit":
		return envcmd.Wrap(&cmdjuju.RemoveUnitCommand{})
	case "action do":
		return envcmd.Wrap(&cmdaction.DoCommand{})
	case "action fetch":
		return envcmd.Wrap(&cmdaction.FetchCommand{})
	default:
		panic("unknown command: " + cmd)
	}
	return nil
}

func (env *procsEnviron) bootstrap(c *gc.C) {
	if procsEnv != nil {
		c.Assert(env, gc.Equals, procsEnv)
		return
	}
	procsEnv = env

	initJuju(c)

	env.run(c, "bootstrap")
	env.run(c, "environment set", "logging-config=<root>=DEBUG")
}

func (env *procsEnviron) addService(c *gc.C, charmName, serviceName string) *procsService {
	if serviceName == "" {
		serviceName = charmName
	}
	charmURL := "local:quantal/" + charmName

	env.run(c, "add-machine", "--series=quantal")
	env.run(c, "deploy", "--to="+env.machine, "--repository="+repoDir, charmURL, serviceName)
	// We leave unit /0 alive to keep the machine alive.

	svc := &procsService{
		env:       env,
		charmName: charmName,
		name:      serviceName,
	}
	svc.dummy = procsUnit{
		svc: svc,
		id:  serviceName + "/0",
	}
	return svc
}

func (env *procsEnviron) destroy(c *gc.C) {
	if env.refCount > 0 {
		env.refCount -= 1
	}
	if env.refCount != 0 || procsEnv == nil {
		return
	}
	env.run(c, "destroy-environment", "--force")
	procsEnv = nil
}

type procsService struct {
	env       *procsEnviron
	charmName string
	name      string
	dummy     procsUnit
	lastUnit  int
}

func (svc *procsService) block(c *gc.C) {
	svc.dummy.runAction(c, "noop", nil)
}

func (svc *procsService) setConfig(c *gc.C, settings map[string]string) {
	if len(settings) == 0 {
		return
	}
	// Try to force at least a little sequentiality on Juju.
	svc.block(c)

	args := []string{svc.name}
	for k, v := range settings {
		args = append(args, fmt.Sprintf("%s=%s", k, v))
	}
	svc.env.run(c, "service set", args...)
	//filename := procsYAMLFile(c, map[string]interface{}{svc.name: settings})
	//svc.env.run(c, "service set", "--config="+filename, svc.name)
}

func (svc *procsService) deploy(c *gc.C, procName, pluginID, status string) *procsUnit {
	settings := map[string]string{
		"plugin-name":   procName,
		"plugin-id":     pluginID,
		"plugin-status": status,
	}
	svc.setConfig(c, settings)

	svc.env.run(c, "service add-unit", "--to="+svc.env.machine, svc.name)

	svc.lastUnit += 1
	u := &procsUnit{
		svc: svc,
		id:  fmt.Sprintf("%s/%d", svc.name, svc.lastUnit),
	}

	u.waitForStatus(c, "started", "pending")
	return u
}

type procsUnit struct {
	svc *procsService
	id  string
}

var procsStatusRegex = regexp.MustCompile(`^- (.*): .* \((.*)\)$`)

func (u *procsUnit) waitForStatus(c *gc.C, target string, okayList ...string) {
	// TODO(ericsnow) Support a timeout?
	for {
		out := u.svc.env.run(c, "status", "--format=short")
		for _, line := range strings.Split(out, "\n") {
			line = strings.TrimSpace(line)
			match := procsStatusRegex.FindStringSubmatch(line)
			if match[0] == "" {
				// Not a match.
				continue
			}
			unitName, status := match[1], match[2]
			if unitName != u.id {
				continue
			}
			if status == target {
				return
			}
			for _, okay := range okayList {
				if status == okay {
					continue
				}
			}
			c.Errorf("invalid status %q", status)
			c.FailNow()
		}

		c.Errorf("no status found for %q", u.id)
		c.FailNow()
	}
}

func (u *procsUnit) setConfigStatus(c *gc.C, status string) {
	settings := map[string]string{"plugin-status": status}
	u.svc.setConfig(c, settings)
}

func (u *procsUnit) destroy(c *gc.C) {
	u.svc.env.run(c, "destroy-unit", u.id)

	u.waitForStatus(c, "stopped", "started")
}

func (u *procsUnit) runAction(c *gc.C, action string, actionArgs map[string]interface{}) map[string]string {
	// Send the command.
	args := []string{
		u.id,
		action,
	}
	for k, v := range actionArgs {
		args = append(args, fmt.Sprintf("%s=%s", k, v))
	}
	doOut := u.svc.env.run(c, "action do", args...)
	//filename := procsYAMLFile(c, actionArgs)
	//doOut := u.svc.env.run(c, "action do", "--params="+filename, u.id, action)
	c.Assert(strings.Split(doOut, ": "), gc.HasLen, 2)
	actionID := strings.Split(doOut, ": ")[1]

	// Get the results.
	fetchOut := u.svc.env.run(c, "action fetch", "--wait=0", actionID)
	result := struct {
		Status  string
		Results map[string]interface{}
	}{}
	err := goyaml.Unmarshal([]byte(fetchOut), &result)
	c.Assert(err, jc.ErrorIsNil)

	// Check and coerce the results.
	if !c.Check(result.Status, gc.Equals, "completed") {
		c.Logf(" got:\n" + fetchOut)
	}
	results := make(map[string]string, len(result.Results))
	for k, v := range result.Results {
		results[k] = v.(string)
	}
	return results
}

func (u *procsUnit) injectStatus(c *gc.C, pluginID, status string) {
	args := map[string]interface{}{
		"id":     pluginID,
		"status": status,
	}
	u.runAction(c, "plugin-setstatus", args)
}

func (u *procsUnit) prepPlugin(c *gc.C, procName, pluginID, status string) {
	args := map[string]interface{}{
		"name":   procName,
		"id":     pluginID,
		"status": status,
	}
	u.runAction(c, "plugin-prep", args)
}

func (u *procsUnit) checkState(c *gc.C, expected []process.Info) {
	var procs []process.Info

	results := u.runAction(c, "list", nil)
	out := strings.TrimSpace(results["out"])
	if out != " [no processes registered]" {
		for _, section := range strings.Split(out, "\n\n") {
			var proc process.Info
			err := goyaml.Unmarshal([]byte(section), &proc)
			c.Assert(err, jc.ErrorIsNil)
			procs = append(procs, proc)
		}
	}

	c.Check(procs, jc.DeepEquals, expected)
}

func (u *procsUnit) checkPluginLog(c *gc.C, expected []string) {
	results := u.runAction(c, "plugin-dump", nil)
	c.Assert(results, gc.HasLen, 1)
	output, ok := results["output"]
	c.Assert(ok, jc.IsTrue)
	lines := strings.Split(output, "\n")
	// TODO(ericsnow) strip header...

	c.Check(lines, jc.DeepEquals, expected)
}

func procsYAMLFile(c *gc.C, value interface{}) string {
	filename := filepath.Join(c.MkDir(), "data.yaml")
	data, err := goyaml.Marshal(value)
	c.Assert(err, jc.ErrorIsNil)
	err = ioutil.WriteFile(filename, []byte(data), 0644)
	c.Assert(err, jc.ErrorIsNil)
	return filename
}
