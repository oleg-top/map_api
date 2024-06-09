from aiogram.types import ReplyKeyboardMarkup, KeyboardButton, KeyboardButtonPollType

profile_keyboard = ReplyKeyboardMarkup(keyboard=[
    [
        KeyboardButton(
            text='Обновить'
        ),
        KeyboardButton(
            text='Получить'
        ),
        KeyboardButton(
            text='Отмена'
        )
    ]
], resize_keyboard=True, one_time_keyboard=True
)
