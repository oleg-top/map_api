from aiogram.filters.state import State, StatesGroup


class SG(StatesGroup):
    profile = State()
    profile_update = State()

    journey = State()
    edit_journey = State()
    edit_journey_notes = State()
    edit_journey_points = State()
    edit_journey_travelers = State()
