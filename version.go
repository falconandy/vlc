package vlc

import (
	"github.com/hashicorp/go-version"
	"regexp"
	"sort"
)

type vlcVersion struct {
	version         *version.Version
	shutdownCommand string
}

type vlcVersionFactory struct {
	versions        []*vlcVersion
	playerVersionRE *regexp.Regexp
}

func newVersionFactory() *vlcVersionFactory {
	vBase := &vlcVersion{version.Must(version.NewVersion("0.0.0")), "quit"}
	v4 := &vlcVersion{version.Must(version.NewVersion("4.0.0")), "shutdown"}

	versions := []*vlcVersion{vBase, v4}
	sort.Slice(versions, func(i, j int) bool {
		return versions[i].version.Compare(versions[j].version) > 0
	})

	return &vlcVersionFactory{
		versions:        versions,
		playerVersionRE: regexp.MustCompile(`VLC media player (\w+.\w+.\w+)`),
	}
}

func (f *vlcVersionFactory) Get(versionStr string) *vlcVersion {
	v, err := version.NewVersion(versionStr)
	if err != nil {
		return f.versions[len(f.versions)-1]
	}

	for _, vs := range f.versions {
		if v.Compare(vs.version) >= 0 {
			return vs
		}
	}

	return f.versions[len(f.versions)-1]
}

func (f *vlcVersionFactory) Detect(lines []string) *vlcVersion {
	for _, line := range lines {
		matches := f.playerVersionRE.FindStringSubmatch(line)
		if matches != nil {
			return f.Get(matches[1])
		}
	}
	return f.Get("")
}
