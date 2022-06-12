
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
	