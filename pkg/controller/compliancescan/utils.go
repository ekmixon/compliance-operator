package compliancescan

import (
	// we can suppress the gosec warning about sha1 here because we don't use sha1 for crypto
	// purposes, but only as a string shortener
	// #nosec G505
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

const DefaultContentContainerImage string = "quay.io/jhrozek/ocp4-openscap-content:latest"

type ComplianceComponent uint

const (
	LOG_COLLECTOR = iota
	OPENSCAP
)

var componentDefaults = []struct {
	defaultImage string
	envVar       string
}{
	{"quay.io/jhrozek/scapresults-k8s:latest", "LOG_COLLECTOR_IMAGE"},
	{"quay.io/jhrozek/openscap-ocp:latest", "OPENSCAP_IMAGE"},
}

// GetComponentImage returns a full image pull spec for a given component
// based on the component type
func GetComponentImage(component ComplianceComponent) string {
	comp := componentDefaults[component]

	imageTag := os.Getenv(comp.envVar)
	if imageTag == "" {
		imageTag = comp.defaultImage
	}
	return imageTag
}

func dnsLengthName(hashPrefix string, format string, a ...interface{}) string {
	const maxDnsLen = 64

	friendlyName := fmt.Sprintf(format, a...)
	if len(friendlyName) < maxDnsLen {
		return friendlyName
	}

	// If that's too long, just hash the name. It's not very user friendly, but whatever
	//
	// We can suppress the gosec warning about sha1 here because we don't use sha1 for crypto
	// purposes, but only as a string shortener
	// #nosec G401
	hasher := sha1.New()
	io.WriteString(hasher, friendlyName)
	return hashPrefix + fmt.Sprintf("%x", hasher.Sum(nil))
}
