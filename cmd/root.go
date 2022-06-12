package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/briandowns/spinner"
	"github.com/kekda-py/create-pycord-app/utils"
	"github.com/spf13/cobra"
)

var qs = []*survey.Question{
	{
		Name:     "project-name",
		Prompt:   &survey.Input{Message: "What is the bot's name?", Default: "the greatest bot"},
		Validate: survey.Required,
	},
	{
		Name:   "dir",
		Prompt: &survey.Input{Message: "What should be the directory's name?", Default: "."},
		Validate: func(ans interface{}) error {
			dir := ans.(string)
			if dir == "." || dir == "./" {
				return nil
			}
			if _, err := os.Stat("./" + dir); err != nil {
				return nil
			}
			return fmt.Errorf("directory `%s` already exists", dir)
		},
	},
	{
		Name: "token",
		Prompt: &survey.Input{
			Message: "Bot's token?",
		},
	},
}

var rootCmd = &cobra.Command{
	Use:   "create-pycord-app",
	Short: "Set up a pycord app by running a single command.",

	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println(utils.Colorize("blue", "Creating a new pycord app...\n"))

		ans := struct {
			Name  string `survey:"project-name"`
			Dir   string `survey:"dir"`
			Token string `survey:"token"`
		}{}

		err := survey.Ask(qs, &ans)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println()

		s := spinner.New(spinner.CharSets[34], 100*time.Millisecond)

		s.Suffix = utils.Colorize("blue", " Creating the directory...")
		s.Start()

		if err = utils.CreateDir(ans.Dir); err != nil {
			fmt.Println(utils.Colorize("red", "Error: "+err.Error()))
			os.Exit(1)
		}

		time.Sleep(1 * time.Second)
		s.Stop()

		s = spinner.New(spinner.CharSets[34], 100*time.Millisecond)

		s.Suffix = utils.Colorize("blue", " Creating the files...")
		s.Start()

		token := ans.Token
		if token == "" {
			token = "TOKEN"
		}

		if err = utils.CreateFiles(ans.Dir, token); err != nil {
			fmt.Println(utils.Colorize("red", "Error: "+err.Error()))
			os.Exit(1)
		}

		time.Sleep(1 * time.Second)
		s.Stop()

		s = spinner.New(spinner.CharSets[34], 100*time.Millisecond)

		s.Suffix = utils.Colorize("blue", " Initializing a github repository...")

		s.Start()

		if err = utils.InitializeGit(ans.Dir); err != nil {
			fmt.Println(utils.Colorize("red", "Error: "+err.Error()))
			os.Exit(1)
		}

		time.Sleep(1 * time.Second)
		s.Stop()

		s = spinner.New(spinner.CharSets[34], 100*time.Millisecond)

		s.Suffix = utils.Colorize("blue", " Creating a virtual environment and installing packages...")

		s.Start()

		if err = utils.InitializeVenv(ans.Dir); err != nil {
			fmt.Println(utils.Colorize("red", "Error: "+err.Error()))
			os.Exit(1)
		}

		time.Sleep(1 * time.Second)
		s.Stop()

		fmt.Println(utils.Colorize("green", "Successfully created a pycord app!\n"))
		fmt.Println("To run the app do -\n ")
		fmt.Println(utils.Colorize("blue", "\tcd "+ans.Dir+"\n"+"\tenv/Scripts/activate\n"+"\tpython main.py"))

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
