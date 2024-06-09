from aiogram.types import ReplyKeyboardMarkup, KeyboardButton, KeyboardButtonPollType

journey_keyboard = ReplyKeyboardMarkup(keyboard=[
    [
        KeyboardButton(
            text='Мои путешествия'
        ),
        KeyboardButton(
            text='Редактировать'
        ),
        KeyboardButton(
            text='Отмена'
        )
    ]
], resize_keyboard=True, one_time_keyboard=True
)

edit_journey_keyboard = ReplyKeyboardMarkup(keyboard=[
    [
        KeyboardButton(
            text='Заметки'
        ),
        KeyboardButton(
            text='Локации'
        ),
        KeyboardButton(
            text='Путешественники'
        ),
        KeyboardButton(
            text='Удалить'
        )
    ]
], resize_keyboard=True, one_time_keyboard=True
)
