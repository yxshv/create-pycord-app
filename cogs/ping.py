
from discord.ext import commands

class Cog(commands.Cog):

	def __init__(self, bot : commands.Bot):
		self.bot = bot

	@commands.command(name="ping")
	async def ping(self, ctx : commands.Context):
		await ctx.send("Pong!")

def setup(bot : commands.Bot):
	bot.add_cog(Cog(bot))	
	