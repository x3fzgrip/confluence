package reader

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"

	"github.com/x3fzgrip/confluence/pkg/confluence"
)

func ReadConfig(path, parent string) (*confluence.Config, error) {
	file, err := parser.ParseFile(token.NewFileSet(), path, nil, 0)
	if err != nil {
		return nil, err
	}
	var current *confluence.Config
	root := &confluence.Config{}
	ast.Inspect(file, func(n ast.Node) bool {
		if n == nil {
			return true
		}
		// fmt.Println(n)
		// fmt.Println(reflect.TypeOf(n))
		// fmt.Println("----------------------------------------------------------------")
		switch x := n.(type) {
		// case *ast.File:
		// 	fmt.Println(x)
		// 	for _, decl := range x.Imports {
		// 		fmt.Println(decl)
		// 	}
		// merge child & root
		case *ast.StructType:
			fmt.Println("pos: ", x.Pos())
			child := &confluence.Config{}
			child.Name = x.Fields.List[0].Names[0].Name
			for i, field := range x.Fields.List {
				fmt.Println("type: ", field.Type)
				fmt.Println("tag: ", field.Tag)
				fmt.Println("root.Name: ", root.Name)
				// for _, name := range field.Names {
				// 	fmt.Println("field: ", name.Name)
				// }
				child.Fields = append(child.Fields, confluence.Field{
					Name: &field.Names[i].Name,
					Tag: &confluence.Tag{
						ZeroValue: field.Tag.Value,
						Env:       field.Tag.Value, // TODO
					},
					// TODO: get type string
					TypeName: fmt.Sprintf("%v - %v", field.Type.Pos(), field.Type.End()),
				})
				// root.Fields = make([]confluence.Field, len(fields))
				// for _, field := range fields {
				// 	fmt.Println("&field.Tag.Value", &field.Tag.Value)
				// }
				fmt.Println("current: ", current)
				// root.Fields = append(root.Fields, confluence.Field{
				// 	Tag: &field.Tag.Value,
				// })
				fmt.Println("----------------")
			}
		// 	fmt.Println("----------------------------------------------------------------")
		case *ast.TypeSpec:
			// REVIEW
			if x.Name.Name == parent {
				root.Name = x.Name.Name
			}
			// TODO
			current = x.Name.Name
		}

		return true
	})

	return nil, nil
}

func ParseFile(path string) (*ast.File, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return parser.ParseFile(token.NewFileSet(), "", f, 0)
}
