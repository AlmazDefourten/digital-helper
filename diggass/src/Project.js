import "./Project.css"
import React from 'react';
function Project(){
    return(
        <div className="Project">
            <div className="bord">
               <h1>Новый проект</h1> 
            </div>
            <div className="ahtung">
                <h3>ОБЩЕЕ</h3>
                <p className="p_kat">Название проекта должно содержать в себе как описательную часть, так и творческую, креативную. При выборе региона - необходимо указать субъект Российской Федерации, в котором будет происходить реализация проекта!</p>
                <p className="p_kat">Учитывайте, что проект не может длиться более 1 календарного года.</p>
                <p className="p_kat">Обратите внимание на актуальность данных, а также на то, что руководителем проекта должно являться лицо, от имени которого подается заявка на Конкурс.</p>
            </div>
            <div className="ein">
                <input id="naz_progect" placeholder="Название проекта *"></input>
                <form>
                    <p className="form_sp">
                        <select id="oblast" className="form_sp" name="proj">
                            <option className="color_text" value={0}>Регион реализации проекта *</option>
                            <option className="color_text" value={1}>Московская область</option>
                            <option className="color_text" value={2}>Иная область</option>
                        </select>
                    </p>
                </form>
            </div>
            <div className="zwei">
                <div className="zagruzka">
                    <p className="p_zaq">+</p>
                    <p className="p_zaq">ДОБАВИТЬ</p>
                </div>
                <div className="text_zag">
                    <h2>Требования к логотипу проекта</h2>
                    <p>Формат — JPEG или PNG<br></br>
                    Разрешение — от 700 до 1600 пикселей по большей стороне (без учета белых полей)<br></br>
                    Размер — не больше 10 МБ</p>
                </div>
            </div>
            <div className="drei">
                <h2>Автор проекта</h2>
                <div className="group">
                    <div className="group1">
                        <input className="in" id="fio" placeholder="Имя"></input>
                        <input className="in" id="number" placeholder="Телефон"></input>
                    </div>
                    <div className="group1">
                        <input className="in" id="email" placeholder="E-mail"></input>
                        <input className="in" id="data" placeholder="Дата рождения"></input>
                    </div>
                </div>
            </div>
            <div className="drei">
                <h2>Автор проекта</h2>
                <div className="group">
                <form>
                    <p className="in form_sp">
                        <select className=" in form_sp" name="proj">
                            <option className="color_text" value={0}>Масштаб реализации проекта *</option>
                            <option className="color_text" value={7}>Муниципальный</option>
                            <option className="color_text" value={8}>Региональный</option>
                            <option className="color_text" value={9}>Межрегиональный</option>
                            <option className="color_text" value={10}>Окружной</option>
                            <option className="color_text" value={11}>Всероссийский</option>
                            <option className="color_text" value={12}>Междунарродный</option>
                        </select>
                    </p>
                </form>
                    <input className="in" id="data_proj" placeholder="Дата начала проекта"></input>  
                </div>
            </div>
            <div className="vier">
                <h1>Руководитель проекта</h1>
                <input id="opit" className="plus" placeholder="Опыт руководителя *"></input>
                <input id="opis" className="plus" placeholder="Описание функцинала руководителя *"></input>
                <p className="p_cat">Какие обязанности есть у него</p>
                <input placeholder="Адрес регистрации руководителя проекта *"></input>
                <p className="p_cat">Добавить резюме</p>
                <button className="but_zag">ЗАГРУЗИТЬ ФАЙЛ *</button>
                <p className="p_kat">Размер загружаемого файла не более 50мб</p>
                <input placeholder="Видео визитка (ссылка на ролик на любом видеохостинге) *"></input>
            </div>
            <div className="funf">
                <p className="p_bold">СОХРАНИТЕ ВНЕСЕННЫЕ ИЗМЕНЕНИЯ ПУТЕМ НАЖАТИЯ КНОПКИ «СОХРАНИТЬ ЧЕРНОВИК» ИЛИ «СОХРАНИТЬ ПРОЕКТ»</p>
                <button className="but_zag zi">СОХРАНИТЬ ЧЕРНОВИК</button>
                <button className="but_zA zi">СОЗДАТЬ ПРОЕКТ</button>
            </div>
        </div>
    );
}

export default Project;