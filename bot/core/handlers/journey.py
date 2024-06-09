from aiogram import Bot
from aiogram.fsm.context import FSMContext
from aiogram.types import Message
from api_manager import APIManager
from core.keyboards.default import default_keyboard
from core.keyboards.journey import journey_keyboard, edit_journey_keyboard
from core.utils.stategroup import SG

FIELDS = ['Название', 'Описание', 'Название путешествия', 'Публична']


async def user_journeys(message: Message, bot: Bot, state: FSMContext):
    journeys = APIManager.get_journeys(message.chat.id)
    await state.clear()
    await bot.send_message(message.chat.id, journeys, reply_markup=default_keyboard)

