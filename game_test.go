package game

import (
  "go/ast"
  "go/parser"
  "go/token"
  "log"
  "testing"
)

func node() *ast.File {
  fset := token.NewFileSet()
  node, err := parser.ParseFile(fset, "game.go", nil, parser.ParseComments)
  if err != nil {
    log.Fatal(err)
  }
  return node
}

func contains(a []string, x string) bool {
  for _, n := range a {
    if x == n {
      return true
    }
  }
  return false
}

func imports(value string) bool {
  var imports []string
  node := node()
  for _, i := range node.Imports {
    imports = append(imports, i.Path.Value)
  }

  return contains(imports, value)
}

func functions(value string) bool {
  var functions []string
  node := node()
  for _, f := range node.Decls {
    fn, ok := f.(*ast.FuncDecl)
    if !ok {
      continue
    }
    functions = append(functions, fn.Name.Name)
  }
  return contains(functions, value)
}

func TestImports(t *testing.T) {
  if !imports("\"log\"") {
    t.Errorf("Are you importing the `log` package?")
  }
  if !imports("\"github.com/hajimehoshi/ebiten\"") {
    t.Errorf("Are you importing the `ebiten` package?")
  }
  if !imports("\"github.com/hajimehoshi/ebiten/ebitenutil\"") {
    t.Errorf("Are you importing the `ebitenutil` package?")
  }
}

func TestFunctions(t *testing.T) {
  if !functions("main") {
    t.Errorf("Have you created a `main` function?")
  }
  if !functions("update") {
    t.Errorf("Have you created an `update` function?")
  }
}