from aiogram import Bot
from aiogram.fsm.context import FSMContext
from aiogram.types import Message
from api_manager import APIManager
from core.keyboards.default import default_keyboard
from core.utils.stategroup import SG


async def get_start(message: Message, state: FSMContext):
    await state.set_state(SG.profile_update)
    response = ('Введите имя пользователя, возраст, локацию, bio, в следующем формате:\n'
                'Имя пользователя: {}\n'
                'Возраст: {}\n'
                'И так далее...')
    await message.reply(response)


# text: "Отмена"
async def cancel(message: Message, state: FSMContext):
    await state.clear()
    await message.reply("Хорошо", reply_markup=default_keyboard)
