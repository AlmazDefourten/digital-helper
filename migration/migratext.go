package migration

const (
	MigrationText = "INSERT INTO public.grants (names, date_starts, date_ends) VALUES('Грант Президента РФ', '2023-01-30', '2023-05-01');" +
		"			 INSERT INTO public.grants (names, date_starts, date_ends) VALUES('Грант мэрии Москвы', '2023-04-30', '2023-06-01');" +
		"			 INSERT INTO public.answers (questions, answers) VALUES('как подать на грант', 'Сначала тебе нужно создать проект для гранта https://grants.myrosmol.ru/projects/create/386d79e1-1fa9-4e9d-9357-f8777644bcde:1c49a8d0-35c1-43c3-894e-ed03087dceaa');" +
		"			 INSERT INTO public.answers (questions, answers) VALUES('Верификация аккаунта через ЕСИА', 'Чтобы пройти верификацию, вам необходимо нажать на соответствующую кнопку в меню редактирования профиля под вашим аватаром. Для того, чтобы верификация прошла полностью, необходимо чтобы профиль на госуслугах имел статус «Подтвержденная» и в разделе Адреса был заполнен адрес фактического проживания (не галочкой соответствия с пропиской, а текстом).\n\nПосле этого убедитесь о наличии галочки в профиле ФГАИС рядом с ФИО.\n\nПосле зайдите на grants.myrosmol.ru,используя логин и пароль ФГАИС (Если система не запрашивает логин и пароль, то выйдите из аккаунта и зайдите заново).');" +
		"			 INSERT INTO public.answers (questions, answers) VALUES('Как выбрать шаблон проекта', 'Шаблон проекта выбирается по названию конкурса, в котором вы планируете участвовать');" +
		"			 INSERT INTO public.answers (questions, answers) VALUES('Создал проект, но он не выбирается при подаче заявки', 'Убедитесь, что вы создали проект по шаблону конкурса, на который вы подаете заявку. Шаблон выбирается по названию конкурса');" +
		"			 INSERT INTO public.answers (questions, answers) VALUES('У меня имеется другая техническая проблема', 'Опишите пожалуйста вашу проблему подробнее, и пришлите пожалуйста на почту support@myrosmol.ru скриншот экрана где видна ошибка которую вы встречаете, а также укажите ID вашего аккаунта в ФГАИС (для этого кликните по имени пользователя в карточке профиля в левой части экрана');" +
		"			 INSERT INTO public.answers (questions, answers) VALUES('Я стал победителем Всероссийского конкурса молодежных инициатив среди физических лиц, как заключить соглашение', 'Если вы стали победителем Всероссийского конкурса молодежных инициатив среди физических лиц (вышел приказ и в разделе «Мои заявки» сменился статус заявки), то в грантовом модуле ФГАИС Молодёжь России вам стал доступен раздел для заключения соглашения. Для того чтобы начать процесс заключения соглашения, необходимо пройти в обозначенный выше раздел и нажать «Начать заключение договора».\n\n1 шаг — это изменение проекта. После изменения проекта необходимо сохранить изменения ( проект сохраниться только тогда, когда в проекте будет корректная сумма) 2 шаг — это внесение информации о вас. После заполнения всех полей, необходимо сохранить данные и отправить на проверку.\n\nПосле проверки вам будет либо согласованы вкладки, либо придут комментарии от куратора по тем изменениям, которые необходимо внести.\n\nПосле согласования вкладки Проект и вкладки Данные, начнется процесс регистрации проекта соглашения в ГИИС «Электронный бюджет», после окончания данного процесса у вас сменится статус и станет доступно подписание.\n\nОбращаем внимание, что подписание будет доступно только с верифицированного аккаунта. Последовательность подписания: Проект — Данные — Приложения(все) — Соглашение\n\nПо техническим вопросам обращаться на почту support@myrosmol.ru');" +
		"			 INSERT INTO public.answers (questions, answers) VALUES('Что делать, если мне не нужно заполнять все статьи расходов', 'При создании проекта в разделе «Расходы» при отсутствии у вас необходимости в заполнении статьи расходов, оставляйте поле пустым.\n\nНедопустимо заполнять поля нулями, прочерками, пробелами и т.д.');" +
		"			 INSERT INTO public.statuses (names) VALUES('В работе');" +
		"			 INSERT INTO public.statuses (names) VALUES('На проверке');" +
		"			 INSERT INTO public.statuses (names) VALUES('Решение вынесено');" +
		"			 INSERT INTO public.applications (grant_id, status_id, user_id, dates) VALUES(1, 1, 0, '2023-01-05');"
)