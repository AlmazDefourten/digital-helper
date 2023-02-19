package models

const (
	START      = 0
	WANT_GRANT = 1
)

func Start() (string, error) {
	result := "Привет! Я Грант бот, задавай мне любые вопросы и я постараюсь на них ответить!"
	applications, err := GetApplications()
	if err != nil {
		return "", err
	}
	if len(applications) == 0 {
		result += "Пока что по конкурсам и грантам у тебя нет заявок, но сейчас отличное время чтоб ее завести :)" +
			"Заинтерисовало? Переходи сюда! https://grants.myrosmol.ru/projects/create/386d79e1-1fa9-4e9d-9357-f8777644bcde:1c49a8d0-35c1-43c3-894e-ed03087dceaa"
	} else {
		result += "Твои текущие заявки: "
		for _, appl := range applications {
			result += appl.Grant.Names + ", статус: " + appl.Status.Names
		}
	}
	return result, nil
}
