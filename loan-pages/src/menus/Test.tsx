import React, { useState, useEffect } from 'react';

const Test: React.FC = () => {
  const [provinces, setProvinces] = useState<string[]>([]);
  const [cities, setCities] = useState<string[]>([]);
  const [district, setDistrict] = useState<string[]>([]);
  const [subdistrict, setSubdistrict] = useState<string[]>([]);
  const [zipCodes, setZipCodes] = useState<string[]>([]);
  const [selectedProvince, setSelectedProvince] = useState<string>('');
  const [selectedCity, setSelectedCity] = useState<string>('');
  const [selectedDistrict, setSelectedDistrict] = useState<string>('');
  const [selectedSubdistrict, setSelectedSubdistrict] = useState<string>('');

  useEffect(() => {
    // Mendapatkan daftar provinsi dari API
    fetch('http://localhost:8000/api/provinces')
      .then(response => response.json())
      .then(data => setProvinces(data));
  }, []);

  const handleProvinceChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const selectedProvince = e.target.value;
    setSelectedProvince(selectedProvince);
    setSelectedDistrict("")
    setSelectedSubdistrict("")

    // Mendapatkan daftar kota berdasarkan provinsi yang dipilih
    fetch(`http://localhost:8000/api/cities?province=${selectedProvince}`)
      .then(response => response.json())
      .then(data => setCities(data));
  };

  const handleCityChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const selectedCity = e.target.value;
    setSelectedCity(selectedCity);
    setSelectedSubdistrict("")

    // Mendapatkan zip code berdasarkan kota yang dipilih
    fetch(`http://localhost:8000/api/districts?city=${selectedCity}`)
      .then(response => response.json())
      .then(data => setDistrict(data));
  };

  console.log(selectedProvince)

  const handleDistrictChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const selectedDistrict = e.target.value;
    setSelectedDistrict(selectedDistrict);

    // Mendapatkan zip code berdasarkan kota yang dipilih
    fetch(`http://localhost:8000/api/subdistricts?district=${selectedDistrict}`)
      .then(response => response.json())
      .then(data => setSubdistrict(data));
  };

  const handleSubdistrictChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const selectedSubdistrict = e.target.value;
    setSelectedSubdistrict(selectedSubdistrict);

    // Mendapatkan zip code berdasarkan kota yang dipilih
    fetch(`http://localhost:8000/api/zip-codes?subdistrict=${selectedSubdistrict}`)
      .then(response => response.json())
      .then(data => setZipCodes(data));
  };

  return (
    <div>
      <h1>Provinsi:</h1>
      <select value={selectedProvince} onChange={handleProvinceChange}>
        <option value="">Pilih Provinsi</option>
        {provinces.map(province => (
          <option key={province} value={province}>{province}</option>
        ))}
      </select>

      <h1>Kota:</h1>
      <select value={selectedCity} onChange={handleCityChange}>
        <option value="">Pilih Kota</option>
        {cities.map(city => (
          <option key={city} value={city}>{city}</option>
        ))}
      </select>

      <h1>District:</h1>
      <select value={selectedDistrict} onChange={handleDistrictChange}>
        <option value="">Pilih Distrik</option>
        {district.map(district => (
          <option key={district} value={district}>{district}</option>
        ))}
      </select>

      <h1>Sub District:</h1>
      <select value={selectedSubdistrict} onChange={handleSubdistrictChange}>
        <option value="">Pilih Sub Distrik</option>
        {subdistrict.map(subdistrict => (
          <option key={subdistrict} value={subdistrict}>{subdistrict}</option>
        ))}
      </select>

      <h1>Zip Code:</h1>
      <ul>
        {zipCodes.map(zipCode => (
          <li key={zipCode}>{zipCode}</li>
        ))}
      </ul>
    </div>
  );
};

export default Test;
