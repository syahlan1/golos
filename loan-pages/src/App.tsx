import { BrowserRouter, Routes, Route } from "react-router-dom";
import Home from './pages/Home';
import Login from './pages/Login';
import './App.css';
import Register from "./pages/Register";
import Test from "./menus/Test";
import InfoUser from "./menus/InfoUser";
import ProtectedRoute from "./components/ProtectedRoute";

function App() {

  return (
    <div className="app">
      <BrowserRouter>
        <Routes>
        <Route path="/" element={<ProtectedRoute children={<Home />} />} />
          <Route path="/test" element= { <Test/>} />
          <Route path= "/login" element={<Login />}/>
          <Route path= "/registrasi" element={<Register />}/>
          <Route path= "/info-user" element={<InfoUser />}/>
        </Routes>
      </BrowserRouter>
    </div>
  )
}
export default App
