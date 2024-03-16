import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faFlag, faMedal, faCog, faHome, faGrinAlt, faRightFromBracket } from "@fortawesome/free-solid-svg-icons";
import { useState } from "react";
import { NavLink, Navigate } from "react-router-dom";

const Sidebar = () => {
  const [redirect, setRedirect] = useState(false)

  const logout = async () => {
    // document.cookie = "jwt=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
    // window.location.href = '/login';
    try {
      const response = await fetch('http://localhost:8000/api/logout', {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
      });

      if (response.ok) {
        setRedirect(true);
      } else {
        throw new Error("Failed to logout");
      }
    } catch (error) {
      console.error(error);
    }
  }

  if(redirect){
    return <Navigate replace to="/login"/>
  }

  return (
    <nav className="sidebar">
      <div className="menu_content">
        {/* Menu Items */}
        <ul className="menu_items">
          {/* Dashboard */}
          <div className="menu_title menu_dashboard"></div>
          <li className="item" >
            <NavLink to="/homepage">
              <div className="nav_link submenu_item" >
                <span className="navlink_icon">
                  <FontAwesomeIcon icon={faHome} />
                </span>
                <span className="navlink">Home</span>
              </div>
            </NavLink>
          </li>

          {/* Overview */}
          <li className="item">
          <NavLink to="/overview">
            <div className="nav_link submenu_item">
              <span className="navlink_icon">
                <FontAwesomeIcon icon={faGrinAlt} />
              </span>
              <span className="navlink">Overview</span>
            </div>
            </NavLink>
          </li>
        </ul>

        {/* Menu Items */}
        <ul className="menu_items">
          {/* Setting */}
          <div className="menu_title menu_setting"></div>
          <li className="item">
            <a href="#" className="nav_link">
              <span className="navlink_icon">
                <FontAwesomeIcon icon={faFlag} />
              </span>
              <span className="navlink">Notice board</span>
            </a>
          </li>
          <li className="item">
            <a href="#" className="nav_link">
              <span className="navlink_icon">
                <FontAwesomeIcon icon={faMedal} />
              </span>
              <span className="navlink">Award</span>
            </a>
          </li>
          <li className="item">
            <a href="#" className="nav_link">
              <span className="navlink_icon">
                <FontAwesomeIcon icon={faCog} />
              </span>
              <span className="navlink">Setting</span>
            </a>
          </li>
          <li className="item">
            <a onClick={logout} className="nav_link" target="">
              <span className="navlink_icon">
                <FontAwesomeIcon icon={faRightFromBracket} />
              </span>
              <span className="navlink" onClick={logout}>Logout</span>
            </a>
          </li>
        </ul>
      </div>
    </nav>
  )
}

export default Sidebar;
