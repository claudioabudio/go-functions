package simplemath

import "fmt"

type SemanticVersion struct {
	major, minor, patch int
}

func NewSemanticVersion(major int, minor int, patch int) SemanticVersion {
	return SemanticVersion{major: major, minor: minor, patch: patch}
}

func (sv SemanticVersion) String() string {
	return fmt.Sprintf("%d.%d.%d", sv.major, sv.major, sv.patch)
}

func (sv *SemanticVersion) IncrementPatch() {
	sv.patch += 1
}

func (sv *SemanticVersion) IncrementMajor() {
	sv.major += 1
}

func (sv *SemanticVersion) IncrementMinor() {
	sv.minor += 1
}
