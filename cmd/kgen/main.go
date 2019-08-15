package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/subchen/go-stack/findup"
	"golang.org/x/tools/go/packages"
)

func main() {
	tmpls := map[string]*template.Template{
		"ksvc": template.Must(template.New("ksvc").Parse(`
apiVersion: serving.knative.dev/v1beta1
kind: Service
metadata:
  name: {{.Name}}
  namespace: default
spec:
  template:
    spec:
      containers:
        - image: {{.Image}}
`)),
	}

	cwd := MustS(os.Getwd())
	goMod := MustS(findup.Find("go.mod"))

	p, err := packages.Load(nil, cwd)
	if err != nil {
		log.Fatal(err)
	}

	args := struct {
		Root  string
		Name  string
		Pkg   string
		Image string
	}{
		Root:  filepath.Dir(goMod),
		Name:  filepath.Base(cwd),
		Pkg:   p[0].PkgPath,
		Image: p[0].PkgPath,
	}

	out := filepath.Join(args.Root, "build", "yml", fmt.Sprintf("%s.generated.yaml", args.Name))
	Must(os.MkdirAll(filepath.Dir(out), 0700))
	Must(tmpls["ksvc"].Execute(MustOpen(out), args))

	log.Printf("Wrote %s for %s in %s \n", "ksvc", args.Pkg, out)
}

func Must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func MustS(s string, err error) string {
	if err != nil {
		log.Fatal(err)
	}

	return s
}

func MustOpen(path string) *os.File {
	f, err := os.OpenFile(
		path,
		os.O_CREATE|os.O_WRONLY,
		0700)

	if err != nil {
		log.Fatal(err)
	}

	return f
}
