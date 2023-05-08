/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	color "github.com/TwiN/go-color"
	"github.com/spf13/cobra"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

var module = false

func LoopFiles(path string) (int, []*ast.CommentGroup, []*ast.ImportSpec, []string) {
	chil, err := os.ReadDir(path)
	if err != nil {
		return 0, nil, nil, nil
	}
	if path == "/" {
		path = ""
	}
	exp := 0
	var com []*ast.CommentGroup
	var imp []*ast.ImportSpec
	var travamnt int = 0
	var name []string

	for _, val := range chil {

		if val.IsDir() {
			travamnt++
			expa, coma, impmqtt, nama := LoopFiles(path + "/" + val.Name()) //me when loop to travcerse ast?\
			com = append(com, coma...)
			imp = append(imp, impmqtt...)
			//always going to be a directory
			//for _, valas := range name {
			//	for _, vala := range nama {
			//		if valas == vala {
			//			continue
			//		}
			//	}
			//}
			name = append(name, nama...)
			exp += expa
			//for k, v := range expa {
			//	exp[k] += v
			//}

		} else {
			valExt := strings.Split(val.Name(), ".")
			if valExt[len(valExt)-1] != "go" {
				continue
			}

			fileSet := token.NewFileSet()
			astre, _ := parser.ParseFile(fileSet, path+"/"+val.Name(), nil, parser.ParseComments)
			co := astre.Comments
			im := astre.Imports
			com = append(com, co...)
			imp = append(imp, im...)
			among := true
			for _, val := range name {
				if val == astre.Name.Name {
					among = false
				}
			}
			if among {
				name = append(name, astre.Name.Name)
			}

			ast.Inspect(astre, func(n ast.Node) bool {
				toFind, okay := n.(*ast.FuncDecl) //ast node of all types
				if okay {
					if toFind.Name.IsExported() {
						//if toFind.Obj == nil {
						//	return true
						//}
						//switch toFind.Obj.Kind {
						//case ast.Var:
						//	exp["Variable"]++
						//case ast.Fun:
						//	exp["Function"]++
						//}
						//fmt.Println(toFind.Obj.Decl)
						exp++
					}
				}
				//st extend asty to make amcros or travers eto fgind nodes of each part of code experssions
				return true
			})
		}

	}

	return exp, com, imp, name
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "main",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		//amqp emsaging protocol
		var pacName []string
		var comments []*ast.CommentGroup
		var deps []*ast.ImportSpec
		exported := 0 //map[string]int{"Variable": 0, "Function": 0}
		if !module {

			//macro edit ast to do actions extend functionality

			fileSet := token.NewFileSet()
			astre, err := parser.ParseFile(fileSet, args[0]+".go", nil, parser.ParseComments)
			pacName = append(pacName, astre.Name.Name)
			deps = astre.Imports
			comments = astre.Comments
			if err != nil {
				panic(err)
			}
			ast.Inspect(astre, func(n ast.Node) bool {
				toFind, okay := n.(*ast.FuncDecl) //ast node of all types
				if okay {
					//if toFind.IsExported() {
					//	if toFind.Obj == nil {
					//		return true
					//	}
					//	switch toFind.Obj.Kind {
					//	case ast.Var:
					//		exported["Variable"]++
					//	case ast.Fun:
					//		exported["Function"]++
					//	}
					//	fmt.Println(toFind.Obj.Decl)
					//}
					if toFind.Name.IsExported() {
						exported++
					}
				}
				//st extend asty to make amcros or travers eto fgind nodes of each part of code experssions
				return true

			})
		} else {
			var x string
			if len(args) == 0 {
				x = "."
			} else {
				x = args[0]
			}
			_, err := os.ReadDir(x)
			if err != nil {
				panic(err)
			}
			exported, comments, deps, pacName = LoopFiles(x)
		}

		fmt.Print(color.Ize(color.Green+color.Bold, "Package Name(s): "))
		for _, val := range pacName {
			fmt.Print(color.Ize(color.Green, strings.TrimSpace(val)+" "))

		}
		fmt.Println()
		fmt.Println(color.Ize(color.Green+color.Bold, "Dependencies:"))
		for _, val := range deps {
			fmt.Println(color.Ize(color.Green, val.Path.Value))
		}
		fmt.Println(color.Ize(color.Green+color.Bold, "Exported: "+fmt.Sprint(exported)))
		//fmt.Println(color.Ize(color.Green+color.Bold, "Variables:"+fmt.Sprint(exported["Variable"])))
		//fmt.Println(color.Ize(color.Green+color.Bold, "Function:"+fmt.Sprint(exported["Function"])))
		fmt.Println(color.Ize(color.Green+color.Bold, "Comments: "+fmt.Sprint(len(comments))))

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.main.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolVar(&module, "module", false, "module")
}
