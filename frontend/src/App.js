import { useEffect, useState } from 'react';
import './App.css';

function App() {
  const [film, setFilm] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    // Goバックエンドからデータを取得
    fetch('http://localhost:8000/film')
      .then((response) => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.json();
      })
      .then((data) => {
        setFilm(data);
        setLoading(false);
      })
      .catch((error) => {
        setError(error);
        setLoading(false);
      });
  }, []);

  if (loading) {
    return <div className="App">Loading...</div>;
  }

  if (error) {
    return <div className="App">Error: {error.message}</div>;
  }

  if (!film) {
    return <div className="App">No film data available</div>;
  }

  return (
    <div className="App">
      <h1 className="title">{film.title} ({film.original_title})</h1>
      <div className="card">
        <p><strong>Description:</strong> {film.description}</p>
        <p><strong>Director:</strong> {film.director}</p>
        <p><strong>Producer:</strong> {film.producer}</p>
        <p><strong>Release Date:</strong> {film.release_date}</p>
        <p><strong>Running Time:</strong> {film.running_time} minutes</p>
        <p><strong>Rotten Tomatoes Score:</strong> {film.rt_score}</p>
      </div>
    </div>
  );
}

export default App;
