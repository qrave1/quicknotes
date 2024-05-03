import React, {useState} from 'react';
import axios from 'axios'
import styles from './auth.module.css'
import {useNavigate} from "react-router-dom";

const authKey = "X-Auth-Token"
const apiBaseUrl = "https://quicknotes-ee54.onrender.com/signup";

interface SignUpFormState {
    username: string;
    email: string;
    password: string
}

function SignUp() {
    const navigate = useNavigate();

    const [formData, setFormData] = useState<SignUpFormState>({
        username: '',
        email: '',
        password: ''
    })

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const {name, value} = e.target;
        setFormData(prevData => ({...prevData, [name]: value}))
    }

    const handleSubmit = async (e: React.ChangeEvent<HTMLFormElement>) => {
        e.preventDefault();
        try {
            const response = await axios.post(apiBaseUrl, formData);
            localStorage.setItem(authKey, response.data)
            navigate('/signin')
        } catch (error) {
            console.error(error);
        }
    }

    console.log(formData);


    return (
        <div className={styles.formContainer}>
            <form onSubmit={handleSubmit} className={styles.form}>
                <div className={styles.formItem}>
                    <label htmlFor="username" className={styles.formLabel}>Имя пользователя</label>
                    <input
                        type="text"
                        name="username"
                        id="username"
                        onChange={handleChange}
                        value={formData.username}
                        required
                        maxLength={20}
                        className={styles.formInput}
                    />
                </div>

                <div className={styles.formItem}>
                    <label htmlFor="email" className={styles.formLabel}>Email</label>
                    <input
                        type="email"
                        name="email"
                        id="email"
                        onChange={handleChange}
                        value={formData.email}
                        required
                        maxLength={40}
                        className={styles.formInput}
                    />
                </div>

                <div className={styles.formItem}>
                    <label htmlFor="password" className={styles.formLabel}>Пароль</label>
                    <input
                        type="password"
                        name="password"
                        id="pass"
                        onChange={handleChange}
                        value={formData.password}
                        required
                        maxLength={20}
                        className={styles.formInput}
                    />
                </div>

                <div className={styles.formButtonContainer}>
                    <button className={styles.formButton}>Зарегистрироваться</button>
                </div>
            </form>
        </div>
    );
}

export default SignUp;
