import {useState} from 'react'
import Button from 'react-bootstrap/Button'
import Container from 'react-bootstrap/Container'
import Nav from 'react-bootstrap/Nav'
import Navbar from 'react-bootstrap/Navbar'
import {useNavigate, NavLink, Link} from 'react-router-dom'

const Header = () => {
    const navigate = useNavigate();
    const [auth, setAuth] = useState(false);

    return (
        <Navbar bg="dark" variant='dark' expand="lg" className='shadow-sm'>
            <container>
                <Navbar.Brand>
                    Dip Flix
                </Navbar.Brand>
            </container>
            <Navbar.Toggle aria-controls="main-navbar-nav" />
            <Navbar.Collapse>
                <Nav className='me-auto'>
                    <Nav.Link as={NavLink} to="/">Home</Nav.Link>
                    <Nav.Link as={NavLink} to="/recommended">Recommended</Nav.Link>
                </Nav>
            </Navbar.Collapse>
            <Nav className='ms-auto align-items-center'>
                {auth ? (
                <>
                    <span>
                        Hello, <strong>Name</strong>
                    </span>
                    <Button variant="outline-right" size="sm">
                        Logout
                    </Button>
                </>
                ):(
                <>
                    <Button variant="outline-info" size="sm" className="me-2" onclick={() => navigate("/login")}>
                        Login
                    </Button>
                    <Button variant="info" size="sm" onclick={() => navigate("/register")}>
                        Register
                    </Button>
                </>
                )}
            </Nav>
        </Navbar>
    )
}