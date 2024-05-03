import Welcome from "./components/welcome/welcome.tsx";
import SignUp from "./components/auth/signUp.tsx";
import Notes from "./components/notes/notes.tsx";
import {BrowserRouter, Route, Routes} from 'react-router-dom';
import SignIn from "./components/auth/signIn.tsx";

function App() {
    return (
        <>
            <BrowserRouter>
                <Routes>
                    <Route path="/" Component={Welcome}/>
                    <Route path="/notes" Component={Notes}/>
                    <Route path="/signup" Component={SignUp}/>
                    <Route  path="/signin" Component={SignIn}/>
                </Routes>
            </BrowserRouter>
        </>
    )
}

export default App
