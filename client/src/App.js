import styled from 'styled-components';
import Container from './components/Container';
import Navbar from './components/Navbar';
import PluginMenu from './components/PluginMenu';
import Panel from './components/Panel';
import { useAstilectron } from './provider/astilectronProvider';
import usePageHook from './hook/pageHook';

const Layout = styled.div`
display: grid;
grid-template-columns: max-content 1fr;
height: 100%;
`

function App() {
  const { astilectron } = useAstilectron()
  const { activePage, pages, selectFuctionPage, selectVariablePage } = usePageHook()
  if (astilectron === null) {
    return null
  }
  return (
    <Container>
      <Navbar></Navbar>
      <Layout>
        <PluginMenu
          onClickFunction={selectFuctionPage}
          onClickVariable={selectVariablePage}></PluginMenu>
        <Panel
          activePage={activePage}
          pages={pages}
          selectFuctionPage={selectFuctionPage}
          selectVariablePage={selectVariablePage}
        />
      </Layout>
    </Container>
  );
}

export default App;
