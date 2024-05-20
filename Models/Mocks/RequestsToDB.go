package Mocks

func (database *DB) GetCatigories() map[string][]string {

	list_catigories := map[string][]string{
		"categories": []string{"РСО", "СНО ФКТиПМ", "Точка кипения", "Центр развития карьеры"},
	}

	return list_catigories
}
