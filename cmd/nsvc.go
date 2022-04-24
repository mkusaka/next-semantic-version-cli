package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Masterminds/semver"
)

var version string

func main() {
	versionType := flag.String("type", "patch", "bump target string.")
	showVersion := flag.Bool("version", false, "print current version")
	flag.Parse()

	if *showVersion {
		fmt.Println(version)
		return
	}

	if len(flag.Args()) != 1 {
		fmt.Println("invalid argument.")
		os.Exit(1)
	}
	str := flag.Args()[0]

	v, err := semver.NewVersion(str)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// TODO: prefix
	//prefix := ""
	//if strings.HasPrefix(str, "^") || strings.HasPrefix(str, "~") || strings.HasPrefix(str, ">") || strings.HasPrefix(str, "<") || strings.HasPrefix(str, "v") {
	//	reg := regexp.MustCompile("[^~><v ]+")
	//	prefix = reg.FindStringSubmatch(str)[0]
	//}
	var incremented semver.Version
	switch *versionType {
	case "major":
		incremented = v.IncMajor()
	case "minor":
		incremented = v.IncMinor()
	case "patch":
		incremented = v.IncPatch()
	default:
		fmt.Printf("unsupported version type: %+v\n", *versionType)
		os.Exit(1)
	}
	fmt.Println(fmt.Sprintf("%s", incremented.String()))
	//fmt.Println(fmt.Sprintf("%s%s", prefix, incremented.String()))
}
