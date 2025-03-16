import React from 'react'
import "./Nav.css"
import { Container } from "react-bootstrap"
import Cookies from "universal-cookie"
import { useHistory } from "react-router-dom"

function Nav() {

    const history = useHistory()

    const logout = () => {
        const cookie = new Cookies();

        cookie.remove("JWT_KEY");
        history.push("/")
    }

    return (
        <div className="Nav">
            <Container className="Nav__container">
                <h3>Shopigram</h3>
                <p onClick={logout}>Logout</p>
            </Container>
        </div>
    )
}

export default Nav
