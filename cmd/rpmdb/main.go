package main

import (
	"fmt"
	"golang.org/x/xerrors"
	"log"
	"path/filepath"

	multierror "github.com/hashicorp/go-multierror"
	rpmdb "github.com/knqyf263/go-rpmdb/pkg"

	_ "github.com/glebarez/go-sqlite"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var files []string
	db, err := detectDB()
	if err != nil {
		return err
	}
	pkgList, err := db.ListPackages()
	if err != nil {
		return err
	}

	fmt.Println("Packages:")
	for _, pkg := range pkgList {
		fmt.Println(pkg.Name)
		files, err = pkg.InstalledFileNames()
		if err != nil {
			return xerrors.Errorf("failed to get installed file names:%w", err)
		}
		fmt.Println("Installed File:")
		for _, file := range files {
			fmt.Println(filepath.ToSlash(file))
		}
		fmt.Println()

	}
	fmt.Printf("[Total Packages: %d]\n", len(pkgList))

	return nil
}

func detectDB() (*rpmdb.RpmDB, error) {
	var result error
	db, err := rpmdb.Open("cmd/rpmdb/Packages")
	if err == nil {
		return db, nil
	}
	result = multierror.Append(result, err)

	return nil, result
}
