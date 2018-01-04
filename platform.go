package make

import (
	"fmt"
	"runtime"
	"strings"
)

// OS represents an operating system.
type OS string

const (
	// Android represents the Android OS.
	Android OS = "android"
	// Darwin represents the Darwin OS.
	Darwin OS = "darwin"
	// Dragonfly represents the Dragonfly OS.
	Dragonfly OS = "dragonfly"
	// FreeBSD represents the FreeBSD OS.
	FreeBSD OS = "freebsd"
	// Linux represents the Linux OS.
	Linux OS = "linux"
	// NetBSD represents the NetBSD OS.
	NetBSD OS = "netbsd"
	// OpenBSD represents the OpenBSD OS.
	OpenBSD OS = "openbsd"
	// Plan9 represents the Plan9 OS.
	Plan9 OS = "plan9"
	// Solaris represents the Solaris OS.
	Solaris OS = "solaris"
	// Windows represents the Windows OS.
	Windows OS = "windows"
)

// ParseOS checks the text and returns an equivalent OS if possible.
func ParseOS(text string) (os OS, err error) {
	text = strings.ToLower(text)

	if text == "native" {
		text = runtime.GOOS
	}

	for _, o := range []OS{Android, Darwin, Dragonfly, FreeBSD, Linux, NetBSD, OpenBSD, Plan9, Solaris, Windows} {
		if string(o) == text {
			os = o
			return
		}
	}
	if os == "" {
		err = fmt.Errorf("invalid OS \"%s\"", text)
	}
	return
}

func (o OS) String() string {
	return string(o)
}

// Arch represents a CPU architecture.
type Arch string

const (
	// Arm represents the Arm architecture.
	Arm Arch = "arm"
	// Arm64 represents the Arm64 architecture.
	Arm64 Arch = "arm64"
	// X386 represents the X386 architecture.
	X386 Arch = "386"
	// Amd64 represents the Amd64 architecture.
	Amd64 Arch = "amd64"
	// Ppc64 represents the Ppc64 architecture.
	Ppc64 Arch = "ppc64"
	// Ppc64LE represents the Ppc64LE architecture.
	Ppc64LE Arch = "ppc64le"
	// Mips represents the Mips architecture.
	Mips Arch = "mips"
	// MipsLE represents the MipsLE architecture.
	MipsLE Arch = "mipsle"
	// Mips64 represents the Mips64 architecture.
	Mips64 Arch = "mips64"
	// Mips64LE represents the Mips64LE architecture.
	Mips64LE Arch = "mips64le"
)

// ParseArch checks the text and returns an equivalent Arch if possible.
func ParseArch(text string) (arch Arch, err error) {
	text = strings.ToLower(text)

	if text == "native" {
		text = runtime.GOARCH
	}

	for _, a := range []Arch{Arm, Arm64, X386, Amd64, Ppc64, Ppc64LE, Mips, MipsLE, Mips64, Mips64LE} {
		if string(a) == text {
			arch = a
			return
		}
	}
	if arch == "" {
		err = fmt.Errorf("invalid architecture \"%s\"", text)
	}
	return
}

func (a Arch) String() string {
	return string(a)
}

// Platform represents a build platform including OS and architecture.
type Platform struct {
	OS        OS
	Arch      Arch
	Extension string
}

// ParsePlatform tries to parse the given OS and Arch and checks if
// the combination of those is supported by go.
func ParsePlatform(osText, archText string) (p *Platform, err error) {
	os, err := ParseOS(osText)
	if err != nil {
		return
	}
	arch, err := ParseArch(archText)
	if err != nil {
		return
	}

	// Note return the platform in SupportedPlatforms because that has more information (e. g. the file extension).
	ok, p := SupportedPlatforms.Contains(&Platform{OS: os, Arch: arch})
	if !ok {
		err = fmt.Errorf("invalid platform, combination of %s and %s is unsupported by go", os, arch)
	}

	return
}

// Equals returns true iff. the OS and architecture of this and the
// other platform are the same.
func (p *Platform) Equals(other *Platform) bool {
	return p.OS == other.OS && p.Arch == other.Arch
}

func (p *Platform) String() string {
	return fmt.Sprintf("%s_%s", p.OS, p.Arch)
}

var AndroidArm = &Platform{OS: Android, Arch: Arm, Extension: ""}
var DarwinX386 = &Platform{OS: Darwin, Arch: X386, Extension: ""}
var DarwinAmd64 = &Platform{OS: Darwin, Arch: Amd64, Extension: ""}
var DarwinArm = &Platform{OS: Darwin, Arch: Arm, Extension: ""}
var DarwinArm64 = &Platform{OS: Darwin, Arch: Arm64, Extension: ""}
var DragonflyAmd64 = &Platform{OS: Dragonfly, Arch: Amd64, Extension: ""}
var FreeBSDX386 = &Platform{OS: FreeBSD, Arch: X386, Extension: ""}
var FreeBSDAmd64 = &Platform{OS: FreeBSD, Arch: Amd64, Extension: ""}
var FreeBSDArm = &Platform{OS: FreeBSD, Arch: Arm, Extension: ""}
var LinuxX386 = &Platform{OS: Linux, Arch: X386, Extension: ""}
var LinuxAmd64 = &Platform{OS: Linux, Arch: Amd64, Extension: ""}
var LinuxArm = &Platform{OS: Linux, Arch: Arm, Extension: ""}
var LinuxArm64 = &Platform{OS: Linux, Arch: Arm64, Extension: ""}
var LinuxPpc64 = &Platform{OS: Linux, Arch: Ppc64, Extension: ""}
var LinuxPpc64LE = &Platform{OS: Linux, Arch: Ppc64LE, Extension: ""}
var LinuxMips = &Platform{OS: Linux, Arch: Mips, Extension: ""}
var LinuxMipsLE = &Platform{OS: Linux, Arch: MipsLE, Extension: ""}
var LinuxMips64 = &Platform{OS: Linux, Arch: Mips64, Extension: ""}
var LinuxMips64LE = &Platform{OS: Linux, Arch: Mips64LE, Extension: ""}
var NetBSDX386 = &Platform{OS: NetBSD, Arch: X386, Extension: ""}
var NetBSDAmd64 = &Platform{OS: NetBSD, Arch: Amd64, Extension: ""}
var NetBSDArm = &Platform{OS: NetBSD, Arch: Arm, Extension: ""}
var OpenBSDX386 = &Platform{OS: OpenBSD, Arch: X386, Extension: ""}
var OpenBSDAmd64 = &Platform{OS: OpenBSD, Arch: Amd64, Extension: ""}
var OpenBSDArm = &Platform{OS: OpenBSD, Arch: Arm, Extension: ""}
var Plan9X386 = &Platform{OS: Plan9, Arch: X386, Extension: ""}
var Plan9Amd64 = &Platform{OS: Plan9, Arch: Amd64, Extension: ""}
var SolarisAmd64 = &Platform{OS: Solaris, Arch: Amd64, Extension: ""}
var WindowsX386 = &Platform{OS: Windows, Arch: X386, Extension: ".exe"}
var WindowsAmd64 = &Platform{OS: Windows, Arch: Amd64, Extension: ".exe"}
var PlatformNone = &Platform{OS: "NONE", Arch: "NONE", Extension: ""}

// SupportedPlatforms contains all platforms supported by go.
var SupportedPlatforms = PlatformSet{
	AndroidArm,
	DarwinX386,
	DarwinAmd64,
	DarwinArm,
	DarwinArm64,
	DragonflyAmd64,
	FreeBSDX386,
	FreeBSDAmd64,
	FreeBSDArm,
	LinuxX386,
	LinuxAmd64,
	LinuxArm,
	LinuxArm64,
	LinuxPpc64,
	LinuxPpc64LE,
	LinuxMips,
	LinuxMipsLE,
	LinuxMips64,
	LinuxMips64LE,
	NetBSDX386,
	NetBSDAmd64,
	NetBSDArm,
	OpenBSDX386,
	OpenBSDAmd64,
	OpenBSDArm,
	Plan9X386,
	Plan9Amd64,
	SolarisAmd64,
	WindowsX386,
	WindowsAmd64,
}
