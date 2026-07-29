package extractor

import "sync"

var (
	fixtureOnce sync.Once
	fixturePkgs *GoPackageSet
)

// sharedFixturePackages loads the Go AST fixture once and caches it.
// All tests that need loadGoPackages(fixtureDir()) should call this instead.
func sharedFixturePackages() *GoPackageSet {
	fixtureOnce.Do(func() {
		fixturePkgs = loadGoPackages(fixtureDir())
	})
	return fixturePkgs
}
