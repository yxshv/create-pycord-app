package utils

import (
	"os"
	"os/exec"
	"runtime"
)

func CreateDir(path string) error {
	if (path != ".") && (path != "./") {
		err := os.Mkdir(path, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateFiles(path string, token string) error {
	file, err := os.Create(path + "/main.py")
	if err != nil {
		return err
	}

	_, err = file.Write([]byte(`
from discord.ext import commands
import asyncio, os, dotenv, discord

dotenv.load_dotenv()
TOKEN = os.getenv('TOKEN')
COGDIR = os.getenv('COGDIR') or './cogs'

class DaBot(commands.Bot):

	def __init__(self):
		
		for file in os.listdir(COGDIR):
			if file.endswith('.py'):
				self.load_extension(f"cogs.{file[:-3]}")

		super().__init__(command_prefix="!")

	async def on_ready(self):
		print('Bot is ready')

	async def on_message(self, message: discord.Message):
		await self.process_commands(message)

bot = DaBot()

bot.run(TOKEN)
	`))

	if err != nil {
		return err
	}

	file, err = os.Create(path + "/.env")
	if err != nil {
		return err
	}

	_, err = file.Write([]byte("TOKEN=" + token))

	if err != nil {
		return err
	}

	err = os.Mkdir(path+"/cogs", 0755)
	if err != nil {
		return err
	}

	file, err = os.OpenFile(path+"/cogs/ping.py", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	_, err = file.Write([]byte(`
from discord.ext import commands

class Cog(commands.Cog):

	def __init__(self, bot : commands.Bot):
		self.bot = bot

	@commands.command(name="ping")
	async def ping(self, ctx : commands.Context):
		await ctx.send("Pong!")

def setup(bot : commands.Bot):
	bot.add_cog(Cog(bot))	
	`))

	if err != nil {
		return err
	}

	file, err = os.OpenFile(path+"/requirements.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	_, err = file.Write([]byte(`
git+https://github.com/Pycord-Development/pycord
python-dotenv
	`))

	if err != nil {
		return err
	}

	return nil

}

func InitializeGit(dir string) error {

	c := exec.Command("git", "init")
	c.Dir = dir

	err := c.Run()

	if err != nil {
		return err
	}

	c = exec.Command("git", "add", ".")
	c.Dir = dir

	err = c.Run()

	if err != nil {
		return err
	}

	c = exec.Command("git", "commit", "-m", "'Initial commit from create-pycord-app'")
	c.Dir = dir

	err = c.Run()

	if err != nil {
		return err
	}

	return nil

}

func InitializeVenv(dir string) error {

	p := "python"
	pip := "./env/Scripts/pip3.exe"

	switch runtime.GOOS {
	case "windows":
		p = "python"
		pip = "./env/Scripts/pip3.exe"
	case "darwin":
		p = "python3"
		pip = "./env/bin/pip3"
	case "linux":
		p = "python3"
		pip = "./env/bin/pip3"
	}

	c := exec.Command(p, "-m", "venv", "env")
	c.Dir = dir

	err := c.Run()

	if err != nil {
		return err
	}

	c = exec.Command(pip, "install", "git+https://github.com/Pycord-Development/pycord", "python-dotenv")
	c.Dir = dir

	err = c.Run()

	if err != nil {
		return err
	}

	return nil

}
