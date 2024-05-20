package Mocks

func (database *DB) GetCatigories() map[string][]string {

	list_catigories := map[string][]string{
		"categories": []string{"РСО", "СНО ФКТиПМ", "Точка кипения", "Центр развития карьеры"},
	}

	return list_catigories
}

func (database *DB) GetEvents(catigory string) any {

	type Event struct {
		id_creator string
		catigory   string
		date       string
		address    string
		definition string
	}

	list_events := []Event{
		Event{id_creator: "727473E7-14B3-4580-A086-2199430834DB", catigory: catigory, date: "25.05.2024 10:21",
			address: "г. Краснодар, ул. Ставропольская 149", definition: "<Информация о мероприятии>"},
		Event{id_creator: "717473E7-54B3-4581-B086-2199430834DA", catigory: catigory, date: "02.06.2024 11:35",
			address: "г. Краснодар, ул. Ставропольская 149", definition: "<Информация о мероприятии>"},

		Event{id_creator: "717473E7-54B3-4580-A086-2199430834DA", catigory: catigory, date: "02.06.2024 11:35",
			address: "г. Краснодар, ул. Ставропольская 149", definition: "<Информация о мероприятии>"},
	}

	return list_events
}
