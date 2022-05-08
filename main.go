package main

import (
	"fmt"
	"path/filepath"

	bzlrule "github.com/bazelbuild/bazel-gazelle/rule"
	bzlbuild "github.com/bazelbuild/buildtools/build"
)

func main() {
	filePath := "./input/BUILD.bazel"
	pkg := filepath.Dir(filePath)

	bzlFile, err := bzlrule.LoadFile(filePath, pkg)
	if err != nil {
		panic(err)
	}

	envvars := map[string]string{}

	for _, r := range bzlFile.Rules {
		if r.Kind() != "deployable" ||
			r.AttrString("target") != "development" ||
			r.Attr("environment") == nil {
			continue
		}

		for _, dictExpr := range r.Attr("environment").(*bzlbuild.DictExpr).List {
			envvars[dictExpr.Key.(*bzlbuild.StringExpr).Value] = dictExpr.Value.(*bzlbuild.StringExpr).Value
		}
	}

	fmt.Println(envvars)
}
