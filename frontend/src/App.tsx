import "./App.css";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import FileBrowser from "@/screens/FileBrowser";

const queryClient = new QueryClient();

function App() {
	return (
		<QueryClientProvider client={queryClient}>
			<div id="App">
				<FileBrowser />
			</div>
		</QueryClientProvider>
	);
}

export default App;
