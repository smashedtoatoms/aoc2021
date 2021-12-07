//go:build mage
// +build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/pkg/errors"
)

var Default = Test

// Removes artifact directories and blows the go test cache.
func Clean() error {
	if err := sh.RunV("go", "clean", "-testcache"); err != nil {
		return err
	}
	if err := sh.RunV("rm", "-rf", "out"); err != nil {
		return err
	}
	return nil
}

// Runs tests.
func Test() error {
	mg.Deps(Clean)
	fmt.Println("testing...")
	out, err := sh.Output("go", "test", "-cover", "-v", "./internal/...")
	fmt.Println(out)
	if err != nil {
		return errors.New("error encountered while testing")
	}
	return err
}

// Builds binary.
func Build() error {
	mg.Deps(Test)
	fmt.Println("building...")
	out, err := sh.Output("go", "build", "-o", "out/", "./cmd/...")
	fmt.Println(out)
	if err != nil {
		return errors.New("error encountered while building")
	}
	return err
}
