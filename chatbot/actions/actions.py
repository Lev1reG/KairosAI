# This files contains your custom actions which can be used to run
# custom Python code.
#
# See this guide on how to implement these action:
# https://rasa.com/docs/rasa/custom-actions


# This is a simple example for a custom action which utters "Hello World!"

# from typing import Any, Text, Dict, List
#
# from rasa_sdk import Action, Tracker
# from rasa_sdk.executor import CollectingDispatcher
#
#
# class ActionHelloWorld(Action):
#
#     def name(self) -> Text:
#         return "action_hello_world"
#
#     def run(self, dispatcher: CollectingDispatcher,
#             tracker: Tracker,
#             domain: Dict[Text, Any]) -> List[Dict[Text, Any]]:
#
#         dispatcher.utter_message(text="Hello World!")
#
#         return []

from typing import Any, Text, Dict, List
from rasa_sdk import Action, Tracker
from rasa_sdk.executor import CollectingDispatcher
from rasa_sdk.events import SlotSet

class ActionAddSchedule(Action):
    def name(self) -> Text:
        return "action_add_schedule"

    def run(self, dispatcher: CollectingDispatcher,
            tracker: Tracker,
            domain: Dict[Text, Any]) -> List[Dict[Text, Any]]:

        event = tracker.get_slot("event-title")
        datetime = tracker.get_slot("date-time")

        # Ucapkan konfirmasi penambahan
        dispatcher.utter_message(text=f"Got it! I’ve added \"{event}\" to your calendar for {datetime}.")
        return [
            SlotSet("event-title", event),
            SlotSet("date-time", datetime)
        ]

class ActionDeleteSchedule(Action):
    def name(self) -> Text:
        return "action_delete_schedule"

    def run(self, dispatcher: CollectingDispatcher,
            tracker: Tracker,
            domain: Dict[Text, Any]) -> List[Dict[Text, Any]]:

        event = tracker.get_slot("event-title")

        dispatcher.utter_message(text=f"The event \"{event}\" has been deleted.")
        return [
            SlotSet("event-title", event)
        ]

class ActionViewSchedule(Action):
    def name(self) -> Text:
        return "action_view_schedule"

    def run(self, dispatcher: CollectingDispatcher,
            tracker: Tracker,
            domain: Dict[Text, Any]) -> List[Dict[Text, Any]]:

        dispatcher.utter_message(text="Here is your schedule: Design Review on Thursday at 10 AM.")
        return []

class ActionUpdateSchedule(Action):
    def name(self) -> Text:
        return "action_update_schedule"

    def run(self, dispatcher: CollectingDispatcher,
            tracker: Tracker,
            domain: Dict[Text, Any]) -> List[Dict[Text, Any]]:

        event = tracker.get_slot("event-title")
        new_datetime = tracker.get_slot("new-date-time")

        dispatcher.utter_message(
            text=f"Got it! \"{event}\" has been updated to {new_datetime}."
        )
        return [
            SlotSet("event-title", event),
            SlotSet("new-date-time", new_datetime)
        ]
