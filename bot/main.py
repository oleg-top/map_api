from aiogram import Bot, Dispatcher, F
from aiogram.fsm.context import FSMContext
from aiogram.types import Message

from core.handlers.basic import get_start, cancel
from aiogram.filters.command import Command

from core.handlers.profile import profile_chain, get_profile, pre_update_profile, update_profile
from core.keyboards import default
from core.keyboards.default import default_keyboard
from core.keyboards.profile import profile_keyboard
from core.utils.commands import set_commands

import asyncio
import logging
import os

from core.utils.stategroup import SG


async def start_bot(bot: Bot):
    await set_commands(bot)


async def start():
    dp = Dispatcher()

    logging.basicConfig(level=logging.INFO,
                        format='%(asctime)s - [%(levelname)s] - %(name)s - '
                               '(%(filename)s).%(funcName)s(%(lineno)d) - %(message)s'
                        )

    bot = Bot(token=os.environ['TOKEN'])

    dp.startup.register(start_bot)

    dp.message.register(get_start, Command(commands=['start']))
    dp.message.register(profile_chain, F.text == 'Профиль')
    dp.message.register(cancel, F.text == 'Отмена')
    dp.message.register(get_profile, F.text == 'Получить', SG.profile)
    dp.message.register(pre_update_profile, F.text == 'Обновить', SG.profile)
    dp.message.register(update_profile, SG.profile_update)

    try:
        await dp.start_polling(bot)
    finally:
        await bot.session.close()


if __name__ == '__main__':
    asyncio.run(start())
