import React, {useState} from 'react';
import axios from 'axios'
import {useNavigate} from "react-router-dom";
import styles from './auth.module.css'

const authKey = "X-Auth-Token"
const apiBaseUrl = "https://quicknotes-ee54.onrender.com/signin";

interface SignInFormState {
    email: string;
    password: string
}

function SignIn() {
    const navigate = useNavigate();

    const [formData, setFormData] = useState<SignInFormState>({
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
            localStorage.setItem(authKey, 'Bearer ' + response.data.token)
            navigate('/')
        } catch (error) {
            console.error(error);
        }
    }

    return (
        <div className={styles.formContainer}>
            <form onSubmit={handleSubmit} className={styles.form}>
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
                    <button className={styles.formButton}>Войти</button>
                </div>
            </form>
        </div>
    );
}

export default SignIn;
