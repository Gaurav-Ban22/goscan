/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"

	color "github.com/TwiN/go-color"
	"github.com/spf13/cobra"
)

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
		fileSet := token.NewFileSet()
		astre, err := parser.ParseFile(fileSet, args[0], nil, parser.ParseComments)
		if err != nil {
			panic(err)
		}
		//macro edit ast to do actions extend functionality
		pacName := astre.Name.Name
		deps := astre.Imports
		comments := astre.Comments
		exported := map[string]int{"Variable": 0, "Function": 0}
		ast.Inspect(astre, func(n ast.Node) bool {
			toFind, okay := n.(*ast.Ident) //ast node of all types
			if okay {
				if toFind.IsExported() {
					switch toFind.Obj.Kind {
					case ast.Var:
						exported["Variable"]++
					case ast.Fun:
						exported["Function"]++
					}
				}
			} else {
				return false
			}
			//st extend asty to make amcros or travers eto fgind nodes of each part of code experssions
			return true

		})
		fmt.Println(color.Ize(color.Green+color.Bold, "Package Name: "+pacName))
		fmt.Print(color.Ize(color.Green+color.Bold, "Dependencies:"))
		for _, val := range deps {
			fmt.Println(color.Ize(color.Green, val.Path.Value))
		}
		fmt.Println(color.Ize(color.Green+color.Bold, "Exported:"+fmt.Sprint(exported["Variable"]+exported["Function"])))
		fmt.Println(color.Ize(color.Green+color.Bold, "Variables:"+fmt.Sprint(exported["Variable"])))
		fmt.Println(color.Ize(color.Green+color.Bold, "Function:"+fmt.Sprint(exported["Function"])))
		fmt.Println(color.Ize(color.Green+color.Bold, "Comments:"+fmt.Sprint(len(comments))))

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
}
