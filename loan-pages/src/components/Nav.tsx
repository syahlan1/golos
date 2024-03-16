
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faSun, faBell} from "@fortawesome/free-solid-svg-icons";
import { useEffect, useState } from "react";

const Nav = () => {
    const [username, setUsername] = useState("");

    useEffect(()=>{
        const fetchUsername = async() =>{
            try{
                const response = await fetch('http://localhost:8000/api/user', {
                    method:"GET",
                    headers: {"Content-Type" : "application/json"},
                    credentials: "include"
                })

                if (response.ok){
                    const content = await response.json()
                    setUsername(content.username)
                }else{
                    throw new Error('Failed to fetch username')
                }
            }
            catch(error) {
                console.error(Error)
            }
        }
        fetchUsername()
    })

    return(
        <nav className="navbar">
        <div className="logo_item">
          <img src="https://upload.wikimedia.org/wikipedia/commons/3/3f/Symbole_Uchiwa.svg" />Strawberry
        </div>

        <div className="navbar_content">
          <FontAwesomeIcon icon={faSun} id="darkLight" className="navbar-icon"/>
          <FontAwesomeIcon icon={faBell} className="navbar-icon"/>
          <span>{username}</span>
        </div>
      </nav>
    )
}

export default Nav