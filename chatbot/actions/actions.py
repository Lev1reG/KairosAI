# from typing import Dict, Any, Text, List
# from rasa_sdk import Action, Tracker
# from rasa_sdk.executor import CollectingDispatcher
# from rasa_sdk.events import SlotSet
# from datetime import datetime, timedelta
# import dateutil.parser
# import requests

# class ActionCreateSchedule(Action):
#     def name(self) -> Text:
#         return "action_create_schedule"
    
#     def run(self, dispatcher: CollectingDispatcher,
#             tracker: Tracker, 
#             domain: Dict[Text, Any]) -> List[Dict[Text, Any]]:
        
#         title = tracker.get_slot("title")
#         description = tracker.get_slot("description")
#         start_time_raw = tracker.get_slot("start_time")
#         end_time = tracker.get_slot("end_time")

#         try: 
#             start_time_dt = dateutil.parser.parse(start_time_raw)
#         except:
#             dispatcher.utter_message(text="Sorry, I couldn't understand the start time you provided.")
#             return []
        
#         if not end_time:
#             end_time_dt = start_time_dt + timedelta(hours=1)
#             end_time = end_time_dt.isoformat()

#         payload = {
#             "title": title,
#             "description": description,
#             "start_time": start_time_dt.isoformat(),
#             "end_time": end_time
#         }

#         print("title:", title)
#         print("description:", description)
#         print("start_time:", start_time_raw)
#         print("End time:", end_time)

#         try: 
#             response = requests.post("http://localhost:5055/schedule", json=payload)
#             if response.status_code == 200:
#                 dispatcher.utter_message(text=f"Schedule '{title}' created from {start_time_dt.strftime('%Y-%m-%d %H:%M')}")
#             else:
#                 dispatcher.utter_message(text="Failed to create schedule. Please try again.")
#         except Exception:
#             dispatcher.utter_message(text="Failed to connect to the scheduling service. Please try again later.")
        
#         return[SlotSet("end_time", end_time)]

from typing import Dict, Any, Text, List
from rasa_sdk import Action, Tracker
from rasa_sdk.executor import CollectingDispatcher
from rasa_sdk.events import AllSlotsReset

import dateutil.parser
from datetime import datetime, timedelta
import requests

BASE_URL = "http://localhost:5055"  # Adjust this to your actual backend URL

class ActionCreateSchedule(Action):
    def name(self) -> Text:
        return "action_create_schedule"

    def run(self, dispatcher: CollectingDispatcher, tracker: Tracker, domain: Dict[Text, Any]) -> List[Dict[Text, Any]]:
        title = tracker.get_slot("title")
        description = tracker.get_slot("description") or "No description provided"
        time_raw = tracker.get_slot("time")

        try:
            start_time_dt = dateutil.parser.parse(time_raw)
            end_time_dt = start_time_dt + timedelta(hours=1)
        except:
            dispatcher.utter_message(text="❌ I couldn't understand the time you provided.")
            return []

        payload = {
            "title": title,
            "description": description,
            "start_time": start_time_dt.isoformat(),
            "end_time": end_time_dt.isoformat()
        }

        dispatcher.utter_message(text=f"🧪 TEST MODE:\nHere's the payload I would send:\n{payload}")
        print("🧪 Payload to backend:", payload)

        # try:
        #     res = requests.post(f"{BASE_URL}/schedule", json=payload)
        #     if res.status_code == 200:
        #         dispatcher.utter_message(text=f"✅ Schedule '{title}' created for {start_time_dt.strftime('%Y-%m-%d %H:%M')}")
        #     else:
        #         dispatcher.utter_message(text="❌ Failed to create schedule.")
        # except Exception as e:
        #     dispatcher.utter_message(text=f"🚨 Backend error: {e}")

        return [AllSlotsReset()]

class ActionViewSchedule(Action):
    def name(self) -> Text:
        return "action_view_schedule"
    
    def run(self, dispatcher: CollectingDispatcher, tracker: Tracker, domain: Dict[Text, Any]) -> List[Dict[Text, Any]]:
        try:
            res = requests.get(f"{BASE_URL}/schedule")
            if res.status_code == 200:
                schedules = res.json()
                if not schedules:
                    dispatcher.utter_message(text="📭 No schedules found.")
                else:
                    message = "📅 Your schedules:\n"
                    for s in schedules:
                        message += f"- {s['title']} at {s['start_time']} | {s['description']}\n"
                    dispatcher.utter_message(text=message)
            else:
                dispatcher.utter_message(text="❌ Failed to fetch schedules.")
        except Exception as e:
            dispatcher.utter_message(text=f"🚨 Error connecting to backend: {e}")
        return []

class ActionDeleteSchedule(Action):
    def name(self) -> Text:
        return "action_delete_schedule"

    def run(self, dispatcher: CollectingDispatcher, tracker: Tracker, domain: Dict[Text, Any]) -> List[Dict[Text, Any]]:
        title = tracker.get_slot("title")

        try:
            res = requests.delete(f"{BASE_URL}/schedule", json={"title": title})
            if res.status_code == 200:
                dispatcher.utter_message(text=f"🗑️ Schedule '{title}' deleted.")
            else:
                dispatcher.utter_message(text="❌ Could not delete the schedule.")
        except Exception as e:
            dispatcher.utter_message(text=f"🚨 Error deleting schedule: {e}")

        return [AllSlotsReset()]
    
class ActionUpdateSchedule(Action):
    def name(self) -> Text:
        return "action_update_schedule"

    def run(self, dispatcher: CollectingDispatcher, tracker: Tracker, domain: Dict[Text, Any]) -> List[Dict[Text, Any]]:
        title = tracker.get_slot("title")
        new_description = tracker.get_slot("description")
        time_raw = tracker.get_slot("time")

        payload = {"title": title}
        if new_description:
            payload["description"] = new_description
        if time_raw:
            try:
                start_time_dt = dateutil.parser.parse(time_raw)
                end_time_dt = start_time_dt + timedelta(hours=1)
                payload["start_time"] = start_time_dt.isoformat()
                payload["end_time"] = end_time_dt.isoformat()
            except:
                dispatcher.utter_message(text="❌ I couldn't understand the new time.")
                return []

        try:
            res = requests.put(f"{BASE_URL}/schedule", json=payload)
            if res.status_code == 200:
                dispatcher.utter_message(text=f"✏️ Schedule '{title}' updated.")
            else:
                dispatcher.utter_message(text="❌ Failed to update schedule.")
        except Exception as e:
            dispatcher.utter_message(text=f"🚨 Backend error: {e}")

        return [AllSlotsReset()]