from aiogram.types import ReplyKeyboardMarkup, KeyboardButton, KeyboardButtonPollType

default_keyboard = ReplyKeyboardMarkup(keyboard=[
    [
        KeyboardButton(
            text='Профиль'
        ),
        KeyboardButton(
            text='Путешествие'
        )
    ]
], resize_keyboard=True, one_time_keyboard=False
)
