package ast

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"

	"golang.org/x/tools/go/ast/astutil"
)

type instrumenter struct {
	traceImport string
	tracePkg    string
	traceFunc   string
}

func New(traceImport, tracePkg, traceFunc string) *instrumenter {
	return &instrumenter{
		traceImport: traceImport,
		tracePkg:    tracePkg,
		traceFunc:   traceFunc,
	}
}

func hasFuncDecl(f *ast.File) bool {
	if len(f.Decls) == 0 {
		return false
	}

	for _, decl := range f.Decls {
		_, ok := decl.(*ast.FuncDecl)
		if ok {
			return true
		}
	}

	return false
}

func (a instrumenter) Instrument(filename string) ([]byte, error){
	fset := token.NewFileSet()
	curAst, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("error parsing %s: %w", filename, err)
	}
	if !hasFuncDecl(curAst){
		return nil, nil
	}

	astutil.AddImport(fset, curAst, a.traceImport)
	a.addDeferTraceIntoFuncDecls(curAst)

	buf := &bytes.Buffer{}
	err = format.Node(buf, fset, curAst)
	if err != nil {
		return nil, fmt.Errorf("error formatting new code: %w", err)
	}
	return buf.Bytes(), nil
}

func (a instrumenter) addDeferTraceIntoFuncDecls(f *ast.File){
	for _, decl := range f.Decls {
		fd, ok := decl.(*ast.FuncDecl)
		if ok {
			a.addDeferStmt(fd)
		}
	}
}

func (a instrumenter) addDeferStmt(fd *ast.FuncDecl) (added bool){
	stmts := fd.Body.List

	for _, stmt := range stmts {
		ds, ok := stmt.(*ast.DeferStmt)
		if !ok {
			continue
		}

		ce, ok := ds.Call.Fun.(*ast.CallExpr)
		if !ok {
			continue
		}

		se, ok := ce.Fun.(*ast.SelectorExpr)
		if !ok {
			continue
		}
		
		x, ok := se.X.(*ast.Ident)
		if !ok {
			continue
		}

		if (x.Name == a.tracePkg) && (se.Sel.Name == a.traceFunc){
			return false
		}
	}

	// not found "defer trace.Trace()()"
	// add one
	ds := &ast.DeferStmt{
		Call: &ast.CallExpr{
			Fun: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X: &ast.Ident{
						Name: a.tracePkg,
					},
					Sel: &ast.Ident{
						Name: a.traceFunc,
					},
				},
			},
		},
	}

	newList := make([]ast.Stmt, len(stmts)+1)
	copy(newList[1:], stmts)
	newList[0] = ds
	fd.Body.List = newList
	return true
}
