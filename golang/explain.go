package golang

import (
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"golang.org/x/mod/module"
)

func Explain(gopkg string, promptFilename string) error {
	log.Println("ensure downloaded", gopkg, "...")
	if err := exec.Command("go", "mod", "download", gopkg).Run(); err != nil {
		return err
	}
	parts := strings.Split(gopkg, "@")
	var pkg, version string
	if len(parts) == 2 {
		pkg = parts[0]
		version = parts[1]
	} else {
		if len(parts) == 1 {
			pkg = parts[0]
			version = "latest"
		}
	}
	modpath, err := getModulePath(pkg, version)
	if err != nil {
		log.Fatal(err)
	}
	abspath := path.Join(getGoPath(), modpath)
	pkgname := path.Base(modpath)
	return explainGoPackageIn(abspath, []string{}, pkgname+".md", promptFilename)
}

func getGoPath() string {
	if p := os.Getenv("GOPATH"); p != "" {
		return p
	}
	return path.Join(os.Getenv("HOME"), "go")
}

// https://stackoverflow.com/questions/67211875/how-to-get-the-path-to-a-go-module-dependency
func getModulePath(name, version string) (string, error) {
	// first we need GOMODCACHE
	cache, ok := os.LookupEnv("GOMODCACHE")
	if !ok {
		cache = path.Join(os.Getenv("GOPATH"), "pkg", "mod")
	}

	// then we need to escape path
	escapedPath, err := module.EscapePath(name)
	if err != nil {
		return "", err
	}

	// version also
	escapedVersion, err := module.EscapeVersion(version)
	if err != nil {
		return "", err
	}

	return path.Join(cache, escapedPath+"@"+escapedVersion), nil
}
