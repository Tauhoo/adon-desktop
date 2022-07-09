import styled from "styled-components"
import Logo from "./Logo"
import { Typography } from "antd"

const Container = styled.div`
    width: 100%;
    border: 1px solid rgb(235, 237, 240);
    padding: 10px 20px;
    display: grid;
    grid-template-columns: max-content  1fr;
    align-items: flex-end;
    gap: 2px;
`

const { Title, Paragraph } = Typography
const NameContainer = styled.div`
  display: grid;
  grid-template-columns: max-content  1fr;
  align-items: flex-end;
  gap: 10px;
`

function Navbar() {
    return <Container>
        <Logo height="40" width="40" />
        <NameContainer>
            <Title style={{ margin: "0px", lineHeight: "33px" }}>don
            </Title>
            <Paragraph style={{ lineHeight: "18px", margin: "0px" }}>v1.0.0</Paragraph>
        </NameContainer>
    </Container>
}

export default Navbar