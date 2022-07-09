import styled from "styled-components"
import Logo from "./Logo"

const Container = styled.div`
    width: 100%;
    border: 1px solid rgb(235, 237, 240);
    padding: 10px 20px;
`

function Navbar() {
    return <Container>
        <Logo height="40" width="40" />
    </Container>
}

export default Navbar