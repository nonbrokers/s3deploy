// Copyright © 2018 Bjørn Erik Pedersen <bjorn.erik.pedersen@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package lib

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFlagsToConfig(t *testing.T) {
	assert := require.New(t)
	flags := flag.NewFlagSet("test", flag.PanicOnError)
	args := []string{
		"-bucket=mybucket",
		"-config=myconfig",
		"-force=true",
		"-key=mykey",
		"-secret=mysecret",
		"-max-delete=42",
		"-acl=public-read",
		"-path=mypath",
		"-quiet=true",
		"-region=myregion",
		"-source=mysource",
		"-distribution-id=mydistro",
		"-try=true",
	}

	cfg, err := flagsToConfig(flags)
	assert.NoError(err)
	assert.NoError(flags.Parse(args))
	assert.Equal("mybucket", cfg.BucketName)
	assert.Equal("myconfig", cfg.ConfigFile)
	assert.Equal(true, cfg.Force)
	assert.Equal("mykey", cfg.AccessKey)
	assert.Equal("mysecret", cfg.SecretKey)
	assert.Equal(42, cfg.MaxDelete)
	assert.Equal("public-read", cfg.ACL)
	assert.Equal("mypath", cfg.BucketPath)
	assert.Equal(true, cfg.Silent)
	assert.Equal("mysource", cfg.SourcePath)
	assert.Equal(true, cfg.Try)
	assert.Equal("myregion", cfg.RegionName)
	assert.Equal("mydistro", cfg.CDNDistributionID)
}

func TestSetAclAndPublicAccessFlag(t *testing.T) {
	assert := require.New(t)
	flags := flag.NewFlagSet("test", flag.PanicOnError)
	args := []string{
		"-bucket=mybucket",
		"-acl=public-read",
		"-public-access=true",
	}

	cfg, err := flagsToConfig(flags)
	assert.NoError(err)
	assert.NoError(flags.Parse(args))

	check_err := cfg.check()
	assert.Error(check_err)
	assert.Contains(check_err.Error(), "you passed a value for the flags public-access and acl")
}
