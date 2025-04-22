import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Home from './pages/Home';
import NewProgrammer from './pages/NewProgrammer';
import Header from './components/Header';

function App() {
  return (
    <Router>
      <div>
        <Header />
        <Routes>
          <Route exact path="/" element={<Home />} />
          <Route exact path="/newprogrammer" element={<NewProgrammer />} />
        </Routes>
      </div>
    </Router>
  )
}

export default App;