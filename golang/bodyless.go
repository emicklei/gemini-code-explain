package golang

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
)

func sourceWithoutFunctionBodies(path, source string) (string, error) {
	fset := token.NewFileSet() // positions are relative to fset
	astFile, err := parser.ParseFile(fset, path, source, 0)
	if err != nil {
		return "", err
	}
	ast.Inspect(astFile, func(n ast.Node) bool {
		switch fn := n.(type) {
		case *ast.FuncDecl:
			dropBodyInFunction(fn)
			return false // Don't recurse into function bodies
		default:
			return true
		}
	})
	return generateSourceFromAST(astFile)
}

func dropBodyInFunction(fn *ast.FuncDecl) {
	fn.Body = buildEmptyBlock(fn)
}

func buildEmptyBlock(fn *ast.FuncDecl) *ast.BlockStmt {
	results := []ast.Expr{}
	if fn.Type.Results != nil {
		for range fn.Type.Results.List {
			// TODO return correct zero values
			results = append(results, &ast.Ident{Name: "nil"})
		}
	}
	return &ast.BlockStmt{
		List: []ast.Stmt{
			&ast.ReturnStmt{Results: results},
		},
	}
}

// GenerateSourceFromAST generates Go source code from the given AST file.
func generateSourceFromAST(file *ast.File) (string, error) {
	var buf bytes.Buffer

	// Use format.Node to format the generated code according to Go conventions.
	if err := printer.Fprint(&buf, token.NewFileSet(), file); err != nil {
		return "", err
	}

	// Format the generated code using go/format.
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return "", err
	}

	return string(formatted), nil
}
