import styled from "styled-components"
const Wrapper = styled.div`
width: 100vw;
height: 100vh;
`

function Container({ children }) {
    return <Wrapper>{children}</Wrapper>
}

export default Container