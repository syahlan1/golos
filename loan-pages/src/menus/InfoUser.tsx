import React, { useState, useEffect } from 'react';

const InfoUser = () => {
  const [userData, setUserData] = useState({
    username: "",
    role: "",
    permissions: []
  });

  useEffect(() => {
    const fetchUserData = async () => {
      try {
        const response = await fetch('http://localhost:8000/api/user', {
          method: "GET",
          headers: { "Content-Type": "application/json" },
          credentials: "include"
        });

        if (response.ok) {
          const content = await response.json();
          setUserData({
            username: content.username,
            role: content.Role ? content.Role.name : "No Role",
            permissions: content.Role ? content.Role.Permissions.map((permission: { name: any; }) => permission.name) : []
          });
        } else {
          throw new Error('Failed to fetch user data');
        }
      }
      catch (error) {
        console.error(error);
      }
    };
    fetchUserData();
  }, []);

  return (
    <div>
      <h2>User Profile</h2>
      <p>Username: {userData.username}</p>
      <p>Role: {userData.role}</p>
      <p>Permissions: {userData.permissions.join(", ")}</p>
    </div>
  );
};

export default InfoUser;
