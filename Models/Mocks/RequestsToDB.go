package Mocks

import "EventsList/Models/Mocks/serialize_struct"

type Event = serialize_struct.Event

func (database *DB) GetCatigories() map[string][]string {

	list_catigories := map[string][]string{
		"categories": []string{"РСО", "СНО ФКТиПМ", "Точка кипения", "Центр развития карьеры"},
	}

	return list_catigories
}

func (database *DB) GetEvents(catigory string) []Event {

	list_events := []Event{
		Event{Id_creator: "727473E7-14B3-4580-A086-2199430834DB", Catigory: catigory, Date: "25.05.2024 10:21",
			Address: "г. Краснодар, ул. Ставропольская 149", Definition: "<Информация о мероприятии>"},
		Event{Id_creator: "717473E7-54B3-4581-B086-2199430834DA", Catigory: catigory, Date: "02.06.2024 11:35",
			Address: "г. Краснодар, ул. Ставропольская 149", Definition: "<Информация о мероприятии>"},

		Event{Id_creator: "717473E7-54B3-4580-A086-2199430834DA", Catigory: catigory, Date: "02.06.2024 11:35",
			Address: "г. Краснодар, ул. Ставропольская 149", Definition: "<Информация о мероприятии>"},
	}

	return list_events
}
