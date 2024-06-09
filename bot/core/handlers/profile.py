from aiogram import Bot
from aiogram.fsm.context import FSMContext
from aiogram.types import Message
from api_manager import APIManager
from core.keyboards.default import default_keyboard
from core.keyboards.profile import profile_keyboard
from core.utils.stategroup import SG

FIELDS = ['Имя пользователя', 'Возраст', 'Локация', 'Bio']


# text: "Профиль"

async def profile_chain(message: Message, bot: Bot, state: FSMContext):
    await state.set_state(SG.profile)
    await bot.send_message(message.chat.id, 'Выберите действие: получить или обновить данные профиля',
                           reply_markup=profile_keyboard)


# state: profile and text: "Получить"
async def get_profile(message: Message, state: FSMContext):
    data = APIManager.get_profile_data(message.chat.id)
    if 'message' in data:
        response = data['message']
    else:
        response = (f'Имя пользователя: {data['username']}\n'
                    f'Возраст: {str(data['age'])}\n'
                    f'Локация: {data['location']}\n'
                    f'Bio: {data['bio']}\n'
                    f'ID: {data['id']}')
    await state.clear()
    await message.reply(response, reply_markup=default_keyboard)


# state: profile and text: "Обновить"
async def pre_update_profile(message: Message, state: FSMContext):
    await state.set_state(SG.profile_update)
    response = ('Введите имя пользователя, возраст, локацию, bio, в следующем формате:\n'
                'Имя пользователя: {}\n'
                'Возраст: {}\n'
                'И так далее...')
    await message.reply(response)


# state: profile_update
async def update_profile(message: Message, state: FSMContext):
    body = message.text.split('\n')
    data = {}
    for row in body:
        values = row.split(': ')
        if len(values) != 2:
            await message.reply('Неправильный формат данных. Попробуйте еще раз')
            return
        else:
            key, val = values[0].strip(), values[1].strip()
            if key not in FIELDS:
                await message.reply('Неправильный формат данных. Попробуйте еще раз')
                return
            if key == 'Возраст':
                try:
                    data[key] = int(val)
                except:
                    await message.reply('Неправильный формат данных. Попробуйте еще раз')
                    return
            else:
                data[key] = val
    data['id'] = message.chat.id
    res = APIManager.update_profile_data(data)
    await state.clear()
    await message.reply(res, reply_markup=default_keyboard)
