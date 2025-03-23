import React, { useState } from 'react'
import axios from "axios"
import Cookies from "universal-cookie"
import { useHistory } from "react-router-dom"

type Errors = {
    message: string
}[]

function RightLogin() {

    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [showCreate, setShowCreate] = useState(true)
    const [errors, setErrors] = useState<Errors>([])
    const history = useHistory()

    const handleClick = async (e: any) => {
        try {
            const response = await axios.post('/api/users/signup', {
                username,
                password
            }, {
                headers: {
                    'Content-Type': 'application/json',
                },
                withCredentials: true
            });
            let jwt = response.headers.cookie;
            const cookies = new Cookies();
            cookies.set('JWT_KEY', jwt);
            history.push("/dashboard")
        } catch (error) {
            setErrors(error.response.data.errors)
            setTimeout(() => {
                setErrors([])
            }, 2000)
        }
    }

    const login = async () => {
        try {
            const response = await axios.post('/api/users/signin', {
                username,
                password
            }, {
                headers: {
                    'Content-Type': 'application/json',
                },
                withCredentials: true
            });
            let jwt = response.headers.cookie;
            const cookies = new Cookies();
            cookies.set('JWT_KEY', jwt);
            history.push("/dashboard")
        } catch (error) {
            setErrors(error.response.data.errors)
            setTimeout(() => {
                setErrors([])
            }, 2000)
        }
    }

    return (
        <div className="Login__right">
            {
                showCreate ? (
                    <div className="Login__right--container">
                        <h2>Create an Account</h2>
                        <input 
                            value={username} 
                            onChange={(e) => {
                                setUsername(e.target.value)
                            }} 
                            placeholder="username"
                        />
                        <input 
                            value={password} 
                            onChange={(e) => {
                                setPassword(e.target.value)
                            }} 
                            placeholder="password"
                            type="password"
                        />
                        {errors[0] && (
                            <h4>Oooops.... {errors[0].message}</h4>
                        )}
                        <p>Already have an account? <span onClick={() => setShowCreate(false)}>Sign in</span> instead!</p>
                        <button onClick={handleClick}>Create Account</button>
                    </div>
                ) : (
                    <div className="Login__right--container">
                        <h2>Sign In</h2>
                        <input 
                            value={username} 
                            onChange={(e) => {
                                setUsername(e.target.value)
                            }} 
                            placeholder="email"
                        />
                        <input 
                            value={password} 
                            onChange={(e) => {
                                setPassword(e.target.value)
                            }} 
                            placeholder="password"
                        />
                        {errors[0] && (
                            <h4>Oooops.... {errors[0].message}</h4>
                        )}
                        <p>Don't have an account? <span onClick={() => setShowCreate(true)}>Create</span> one now!</p>
                        <button onClick={login}>Sign In</button>
                    </div>
                )
            }
        </div>
    )
}

export default RightLogin
