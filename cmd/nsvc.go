package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Masterminds/semver"
)

func main() {
	versionType := flag.String("type", "patch", "bump target string.")
	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Println("invalid argument.")
		os.Exit(1)
	}
	str := flag.Args()[0]

	version, err := semver.NewVersion(str)
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
		incremented = version.IncMajor()
	case "minor":
		incremented = version.IncMinor()
	case "patch":
		incremented = version.IncPatch()
	default:
		fmt.Printf("unsupported version type: %+v\n", *versionType)
		os.Exit(1)
	}
	fmt.Println(fmt.Sprintf("%s", incremented.String()))
	//fmt.Println(fmt.Sprintf("%s%s", prefix, incremented.String()))
}
