import styled from 'styled-components';
import Container from './components/Container';
import Navbar from './components/Navbar';
import PluginMenu from './components/PluginMenu';
import Panel from './components/Panel';
import { useAstilectron } from './provider/astilectronProvider';

const Layout = styled.div`
display: grid;
grid-template-columns: max-content 1fr;
height: 100%;
`

function App() {
  const { astilectron } = useAstilectron()

  if (astilectron == null) {
    return null
  }
  return (
    <Container>
      <Navbar></Navbar>
      <Layout>
        <PluginMenu></PluginMenu>
        <Panel></Panel>
      </Layout>
    </Container>
  );
}

export default App;
