import styled from 'styled-components';
import Container from './components/Container';
import Navbar from './components/Navbar';
import PluginMenu from './components/PluginMenu';
import Panel from './components/Panel';
import { useAstilectron } from './provider/astilectronProvider';
import usePageHook from './hook/pageHook';
import plugin from './api/plugin';

const Layout = styled.div`
display: grid;
grid-template-columns: max-content 1fr;
height: 100%;
`

function App() {
  const { activePage, pages, selectFuctionPage, selectVariablePage, onRemoveFunctionPage, onRemoveVariablePage } = usePageHook()

  return (
    <Container>
      <Navbar></Navbar>
      <Layout>
        <PluginMenu
          onClickFunction={selectFuctionPage}
          onClickVariable={selectVariablePage}
          onDeletePlugin={plugin.deletePlugin}
        />
        <Panel
          activePage={activePage}
          pages={pages}
          selectFuctionPage={selectFuctionPage}
          selectVariablePage={selectVariablePage}
          onRemoveFunctionPage={onRemoveFunctionPage}
          onRemoveVariablePage={onRemoveVariablePage}
        />
      </Layout>
    </Container>
  );
}


function AstilectronApp() {
  const { astilectron } = useAstilectron()
  if (astilectron === null) {
    return null
  }
  return <App />
}

export default AstilectronApp;
