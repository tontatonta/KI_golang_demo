import { useEffect, useState } from 'react';
import './App.css';

function App() {
  const [films, setFilms] = useState([]);
  const [selectedFilm, setSelectedFilm] = useState(null);
  const [filmDetails, setFilmDetails] = useState(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  // 映画リストを取得
  useEffect(() => {
    fetch('https://ghibliapi.vercel.app/films')
      .then((response) => response.json())
      .then((data) => setFilms(data))
      .catch((error) => console.error('Error fetching film list:', error));
  }, []);

  // 選択された映画の情報を取得
  const fetchFilmDetails = (filmID) => {
    setLoading(true);
    setError(null);
    fetch(`http://localhost:8000/film?id=${filmID}`)
      .then((response) => {
        if (!response.ok) {
          throw new Error('Failed to fetch film details');
        }
        return response.json();
      })
      .then((data) => {
        setFilmDetails(data);
        setLoading(false);
      })
      .catch((error) => {
        setError(error);
        setLoading(false);
      });
  };

  const handleFilmChange = (event) => {
    const filmID = event.target.value;
    setSelectedFilm(filmID);
    if (filmID) {
      fetchFilmDetails(filmID);
    } else {
      setFilmDetails(null);
    }
  };

  return (
    <div className="App">
      <h1 className="title">Select a Ghibli Film</h1>
      <select value={selectedFilm || ''} onChange={handleFilmChange}>
        <option value="">Select a film</option>
        {films.map((film) => (
          <option key={film.id} value={film.id}>
            {film.title}
          </option>
        ))}
      </select>

      {loading && <div>Loading film details...</div>}
      {error && <div>Error: {error.message}</div>}

      {filmDetails && (
        <div className="card">
          <h2>{filmDetails.title} ({filmDetails.original_title})</h2>
          <p><strong>Description:</strong> {filmDetails.description}</p>
          <p><strong>Director:</strong> {filmDetails.director}</p>
          <p><strong>Producer:</strong> {filmDetails.producer}</p>
          <p><strong>Release Date:</strong> {filmDetails.release_date}</p>
          <p><strong>Running Time:</strong> {filmDetails.running_time} minutes</p>
          <p><strong>Rotten Tomatoes Score:</strong> {filmDetails.rt_score}</p>
        </div>
      )}
    </div>
  );
}

export default App;
