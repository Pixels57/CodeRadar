import './header.css';
import { Link } from 'react-router-dom';

function Header() {
  return (
    <header className="header">
        <div className="left">
            <Link to="/" style={{textDecoration: "none"}}>
            </Link>
        </div>
        <div className="right">
            <Link to="/newprogrammer" style={{textDecoration: "none"}}>
                <p className="header-button">Add Programmer</p>
            </Link>
        </div>
    </header>
  );
};

export default Header;
