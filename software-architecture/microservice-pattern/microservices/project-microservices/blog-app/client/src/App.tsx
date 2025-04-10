import React, { useEffect, useState } from 'react';
import './App.css';
import { Switch, Route } from "react-router-dom"
import Login from './pages/Login/Login';
import Dashboard from './pages/Dashboard/Dashboard';
import Cookies from "universal-cookie"
import axios from "axios"
import spinner from "./assets/spinner.svg"
import { useHistory } from "react-router-dom"

function App() {

  const [loading, setLoading] = useState(true);
  const history = useHistory()

  useEffect(() => {
    fetchUser()
  }, []);

  const fetchUser = async () => {
    try {
      const cookies = new Cookies();
      let jwt: string = cookies.get("JWT_KEY");
      let response = await axios.get('/api/users/currentuser', {
        headers: {
          jwt
        }
      });
      if(jwt) axios.defaults.headers.common['jwt'] = jwt;
      else delete axios.defaults.headers.common['jwt']
      if(response.data.currentUser) {
        history.push("/dashboard")
      } 
      setLoading(false)
    } catch(err) {
      setLoading(false)
    }
  }

  if(loading) return (
    <div>
      <img src={spinner} alt="spinner"/>
    </div>
  )

  return (
    <div className="App">
        <Switch>
          <Route exact strict path="/" component={Login}/>
          <Route exact strict path="/dashboard" component={Dashboard}/>
        </Switch>
    </div>
  );
}

export default App;
