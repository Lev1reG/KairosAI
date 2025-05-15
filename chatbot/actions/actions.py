from typing import Dict, Any, Text, List
from rasa_sdk import Action, Tracker
from rasa_sdk.executor import CollectingDispatcher
from rasa_sdk.events import SlotSet
from datetime import datetime, timedelta
import dateutil.parser
import requests

class ActionCreateSchedule(Action):
    def name(self) -> Text:
        return "action_create_schedule"
    
    def run(self, dispatcher: CollectingDispatcher,
            tracker: Tracker, 
            domain: Dict[Text, Any]) -> List[Dict[Text, Any]]:
        
        title = tracker.get_slot("title")
        description = tracker.get_slot("description")
        start_time_raw = tracker.get_slot("start_time")
        end_time = tracker.get_slot("end_time")

        try: 
            start_time_dt = dateutil.parser.parse(start_time_raw)
        except:
            dispatcher.utter_message(text="Sorry, I couldn't understand the start time you provided.")
            return []
        
        if not end_time:
            end_time_dt = start_time_dt + timedelta(hours=1)
            end_time = end_time_dt.isoformat()

        payload = {
            "title": title,
            "description": description,
            "start_time": start_time_dt.isoformat(),
            "end_time": end_time
        }

        print("title:", title)
        print("description:", description)
        print("start_time:", start_time_raw)
        print("End time:", end_time)

        try: 
            response = requests.post("http://localhost:8080/schedule", json=payload)
            if response.status_code == 200:
                dispatcher.utter_message(text=f"Schedule '{title}' created from {start_time_dt.strftime('%Y-%m-%d %H:%M')}")
            else:
                dispatcher.utter_message(text="Failed to create schedule. Please try again.")
        except Exception:
            dispatcher.utter_message(text="Failed to connect to the scheduling service. Please try again later.")
        
        return[SlotSet("end_time", end_time)]