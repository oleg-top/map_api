import logging
from typing import Dict, Any

import requests
import os


class APIManager:
    BASE_URL = os.environ['API_URL']

    @staticmethod
    def get_hello() -> str:
        res = requests.get(f'{APIManager.BASE_URL}/hello/').text
        return res

    @staticmethod
    def get_profile_data(chat_id: int) -> Dict[str, Any]:
        res = requests.get(f'{APIManager.BASE_URL}/profile/about/{str(chat_id)}')
        return res.json()

    @staticmethod
    def update_profile_data(body: Dict[str, Any]) -> str:
        if 'Имя пользователя' in body:
            username = body['Имя пользователя']
        else:
            username = ''

        if 'Возраст' in body:
            age = body['Возраст']
        else:
            age = 0

        if 'Локация' in body:
            location = body['Локация']
        else:
            location = ''

        if 'Bio' in body:
            bio = body['Bio']
        else:
            bio = ''

        data = {
            'username': username,
            'age': age,
            'location': location,
            'bio': bio,
            'id': body['id']
        }

        res = requests.post(f'{APIManager.BASE_URL}/profile/update', json=data)

        if 200 <= res.status_code <= 299:
            return res.json()['message']
        return 'Oops... Something went wrong'

    @staticmethod
    def create_journey(name: str, description: str, chat_id: int) -> str:
        data = {
            'name': name,
            'description': description,
            'chat_id': chat_id
        }
        res = requests.post(f'{APIManager.BASE_URL}/journey/create', json=data)

        return res.text

    @staticmethod
    def add_point_to_journey(
            start_date: str, end_date: str, location: str, description: str, journey_name: str
    ) -> str:
        data = {
            'start_date': start_date,
            'end_date': end_date,
            'location': location,
            'description': description,
            'journey_name': journey_name
        }
        res = requests.post(f'{APIManager.BASE_URL}/journey/point/add')

        return res.text

    @staticmethod
    def remove_point_from_journey(
            start_date: str, end_date: str, location: str, journey_name: str
    ) -> str:
        data = {
            'start_date': start_date,
            'end_date': end_date,
            'location': location,
            'journey_name': journey_name
        }
        res = requests.post(f'{APIManager.BASE_URL}/journey/point/remove', json=data)

        return res.text

    @staticmethod
    def get_journeys(chat_id: int):
        res = requests.get(f'{APIManager.BASE_URL}/journey/chat/{str(chat_id)}')
        return res.text

    @staticmethod
    def remove_journey(name: str) -> str:
        res = requests.delete(f'{APIManager.BASE_URL}/journey/remove/{name}')
        return res.text

    @staticmethod
    def add_note_to_journey(name: str, description: str, journey_name: str, is_public: bool) -> str:
        data = {
            'name': name,
            'description': description,
            'journey_name': journey_name,
            'is_public': is_public
        }

        res = requests.post(f'{APIManager.BASE_URL}/journey/note/create', json=data)
        return res.text
