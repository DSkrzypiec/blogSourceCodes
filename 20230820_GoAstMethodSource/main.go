package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"log"

	"golang.org/x/tools/go/packages"
)

func main() {
	pkgs := projectPackages()

	fmt.Println(" ------------------------------ Packages basic info: ------------------------------")
	for _, pkg := range pkgs {
		fmt.Printf("Package: %s (%s)\n", pkg.ID, pkg.Name)
		fmt.Printf("  -GoFiles: %v\n", pkg.GoFiles)
		for idx, astFile := range pkg.Syntax {
			fmt.Printf("    -%s has %d top-level declarations\n", pkg.GoFiles[idx], len(astFile.Decls))
		}
		fmt.Printf("  -CompiledGoFiles: %v\n", pkg.CompiledGoFiles)
		fmt.Printf("  -IgnoredFiles: %v\n", pkg.IgnoredFiles)
	}

	fmt.Println(" ------------------------------ Finding A.Hello method source: ------------------------------")
	for _, pkg := range pkgs {
		for idx, astFile := range pkg.Syntax {
			helloMethod := findMethodInAST(astFile, "A", "Hello")
			if helloMethod != nil {
				fmt.Printf("Found A.Hello method in package [%s] and file [%s]!\n", pkg, pkg.GoFiles[idx])
				fmt.Println("A.Hello body AST:")
				ast.Print(nil, helloMethod.Body)
				fmt.Println("Function body source code:")
				var buf bytes.Buffer
				err := printer.Fprint(&buf, pkg.Fset, helloMethod.Body)
				if err != nil {
					log.Panic("Couldn't print A.Hello method body: %s", err.Error())
					return
				}
				src := buf.String()
				fmt.Println(src)
			}
		}
	}
}

func findMethodInAST(astFile *ast.File, typeName, methodName string) *ast.FuncDecl {
	for _, decl := range astFile.Decls {
		funcDecl, isFunc := decl.(*ast.FuncDecl)
		if !isFunc {
			continue
		}
		if funcDecl.Recv == nil || len(funcDecl.Recv.List) != 1 || funcDecl.Name.Name != methodName {
			continue
		}

		ident, isIdent := funcDecl.Recv.List[0].Type.(*ast.Ident)
		if isIdent && ident.Name == typeName {
			return funcDecl
		}

		// Check for *T receivers
		starExpr, isStar := funcDecl.Recv.List[0].Type.(*ast.StarExpr)
		if isStar {
			ident, isIdent := starExpr.X.(*ast.Ident)
			if isIdent && ident.Name == typeName {
				return funcDecl
			}
		}
	}
	return nil
}

func projectPackages() []*packages.Package {
	cfg := &packages.Config{
		Mode:  packages.NeedFiles | packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo,
		Tests: false,
	}
	pkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		log.Panic(err)
	}
	return pkgs
}
