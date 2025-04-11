import React from 'react'
import LeftLogin from '../../components/LeftLogin/LeftLogin'
import RightLogin from '../../components/RightLogin/RightLogin'
import "./Login.css"

function Login() {
    return (
        <div className="Login">
            <LeftLogin />
            <RightLogin />
        </div>
    )
}

export default Login
