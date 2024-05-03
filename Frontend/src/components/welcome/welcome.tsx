import "./welcome.css"
import Button from '@mui/material/Button';
import {useNavigate} from "react-router-dom";

function Welcome() {
    const navigate = useNavigate();

    return (
        <>
            <header>
                <div className="header_content">
                    <a href="/" className="logo">
                        <svg xmlns="http://www.w3.org/2000/svg" width="40" height="40" viewBox="0 0 24 24" fill="white">
                            <path
                                d="M19,14.5 L19,5.5 C19,4.67157288 18.3284271,4 17.5,4 L6.5,4 C5.67157288,4 5,4.67157288 5,5.5 L5,18.5 C5,19.3284271 5.67157288,20 6.5,20 L13.5,20 C14.3284271,20 15,19.3284271 15,18.5 C15,17.1192881 16.1192881,16 17.5,16 C18.3284271,16 19,15.3284271 19,14.5 L19,14.5 Z M18.5014408,16.7913481 C18.1948298,16.9255432 17.8561101,17 17.5,17 C16.6715729,17 16,17.6715729 16,18.5 C16,18.8561101 15.9255432,19.1948298 15.7913481,19.5014408 C16.9873685,18.9526013 17.9526013,17.9873685 18.5014408,16.7913481 L18.5014408,16.7913481 Z M4,5.5 C4,4.11928813 5.11928813,3 6.5,3 L17.5,3 C18.8807119,3 20,4.11928813 20,5.5 L20,14.5 C20,18.0898509 17.0898509,21 13.5,21 L6.5,21 C5.11928813,21 4,19.8807119 4,18.5 L4,5.5 Z M8.5,9 C8.22385763,9 8,8.77614237 8,8.5 C8,8.22385763 8.22385763,8 8.5,8 L15.5,8 C15.7761424,8 16,8.22385763 16,8.5 C16,8.77614237 15.7761424,9 15.5,9 L8.5,9 Z M8.5,12 C8.22385763,12 8,11.7761424 8,11.5 C8,11.2238576 8.22385763,11 8.5,11 L15.5,11 C15.7761424,11 16,11.2238576 16,11.5 C16,11.7761424 15.7761424,12 15.5,12 L8.5,12 Z M8.5,15 C8.22385763,15 8,14.7761424 8,14.5 C8,14.2238576 8.22385763,14 8.5,14 L13.5,14 C13.7761424,14 14,14.2238576 14,14.5 C14,14.7761424 13.7761424,15 13.5,15 L8.5,15 Z"/>
                        </svg>

                        QuickNotes
                    </a>

                    <div className="auth">
                        <Button variant="contained" onClick={() => {
                            navigate('/signup')
                        }}>Зарегистрироваться</Button>
                        <Button variant="contained" color="success" onClick={() => {
                            navigate('/signin')
                        }}>Войти</Button>
                    </div>
                </div>
            </header>
            <div className="main-container">
                <div className="intro">
                    <div className="large">Ваш быстрый путь к созданию заметок</div>
                    <div className="mid">От регистрации до создания первой заметки - пара минут!</div>
                    <div className="small">Простота интерфейса и взаимодействия с ним способствует продуктивной работе с
                        заметками
                    </div>
                    <Button className="toNotes" variant="contained" onClick={() => {
                        if (localStorage.getItem("X-Auth-Token") === null) {
                            alert("Вы не авторизованы")
                        } else {
                            navigate('/notes')
                        }
                    }}>К заметкам</Button>
                </div>
                <img src="https://i.pinimg.com/736x/11/13/f6/1113f62b818e33f88264df5494694b1a.jpg" alt="котик"/>
            </div>
        </>
    )
}

export default Welcome
