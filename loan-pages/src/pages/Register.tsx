import { useState } from "react";
import { Navigate } from "react-router-dom";

const Register = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [errorMessage, setErrorMessage] = useState('');
    const [redirect, setRedirect] = useState(false);

    const submit = async () => {
        try {
          const response = await fetch("http://localhost:8000/api/register", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ 
                username, 
                password 
            }),
          });
    
          if (!response.ok) {
            const data = await response.json();
            // Cek jika respons berisi pesan kesalahan
            if (data.message) {
              setErrorMessage(data.message);
            } else {
              setErrorMessage("Failed to register");
            }
            return;
          }

          setRedirect(true)
          
        } catch (error) {
          console.error(error);
          setErrorMessage("Failed to register");
        }
    }
    // Jika pendaftaran berhasil, arahkan pengguna ke halaman lain (misalnya, halaman login)
    if(redirect){
        return <Navigate to={'/login'} />;
    }

    return (
        <div className="form-bg">
            {
                errorMessage && 
                <label>
                    <input type="checkbox" className="alertCheckbox" autoComplete="on" />
                    <div className="alert error">
                        <span className="alertClose">X</span>
                        <span className="alertText">{errorMessage}
                        <br className="clear"/></span>
                    </div>
                </label>
            }
            <form onSubmit={(e) => { e.preventDefault(); submit(); }} className="form-horizontal">
                <span className="heading">Register</span>
                <div className="form-group">
                    <input type="text" className="form-control" placeholder="Username" required
                    onChange={(e)=> setUsername(e.target.value)}/>
                </div>
                <div className="form-group help">
                    <input type="password" className="form-control" placeholder="Password" required
                    onChange={(e)=>setPassword(e.target.value)}/>
                </div>
                <div className="form-group">
                    <button type="submit" className="btn btn-default">Submit</button>
                </div>
            </form>
        </div>
    );
};

export default Register;