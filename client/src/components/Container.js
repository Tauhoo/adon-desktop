import styled from "styled-components"

const Wrapper = styled.div`
width: 100vw;
height: 100vh;
display: grid;
grid-template-rows: max-content calc(100vh - 64px);
`

function Container({ children }) {
    return <Wrapper>{children}</Wrapper>
}

export default Container