// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package process_test

import (
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"
	"gopkg.in/juju/charm.v5"

	"github.com/juju/juju/process"
	"github.com/juju/juju/testing"
)

type infoSuite struct {
	testing.BaseSuite
}

var _ = gc.Suite(&infoSuite{})

func (s *infoSuite) newInfo(name, procType string) *process.Info {
	return &process.Info{
		Process: charm.Process{
			Name: name,
			Type: procType,
		},
	}
}

func (s *infoSuite) TestValidateOkay(c *gc.C) {
	info := s.newInfo("a proc", "docker")
	info.CharmID = "somecharm"
	info.UnitID = "somecharm/0"
	err := info.Validate()

	c.Check(err, jc.ErrorIsNil)
}

func (s *infoSuite) TestValidateBadMetadata(c *gc.C) {
	info := s.newInfo("a proc", "")
	err := info.Validate()

	c.Check(err, gc.ErrorMatches, ".*type: name is required")
}

func (s *infoSuite) TestValidateMissingCharmID(c *gc.C) {
	info := s.newInfo("a proc", "docker")
	err := info.Validate()

	c.Check(err, gc.ErrorMatches, "missing CharmID")
}

func (s *infoSuite) TestValidateMissingUnitID(c *gc.C) {
	info := s.newInfo("a proc", "docker")
	info.CharmID = "somecharm"
	err := info.Validate()
	c.Check(err, jc.ErrorIsNil)

	info.Details.ID = "my-proc"
	err = info.Validate()
	c.Check(err, gc.ErrorMatches, "missing UnitID")
}

func (s *infoSuite) TestIsRegisteredTrue(c *gc.C) {
	info := s.newInfo("a proc", "docker")
	info.Details.ID = "abc123"
	info.Details.Status.Label = "running"
	isRegistered := info.IsRegistered()

	c.Check(isRegistered, jc.IsTrue)
}

func (s *infoSuite) TestIsRegisteredFalse(c *gc.C) {
	info := s.newInfo("a proc", "docker")
	isRegistered := info.IsRegistered()

	c.Check(isRegistered, jc.IsFalse)
}
