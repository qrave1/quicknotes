import React from "react"
import styles from './header.css'

export function Header() {
    return (
        <header class="header">
            <div className="header_content">
                {/* make routing to index here */}
                <a href="/" className="logo">
                    <img src="./icon.svg" alt="Logo" />
                    QuickNotes
                </a>

                <div class="search">
                    <input type="text" placeholder="Search" />
                </div>

                <div class="login">
                    <button>Login</button>
                </div>
            </div>
        </header>
    )
}


