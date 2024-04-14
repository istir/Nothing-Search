// import "./style.css";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { useEffect } from "react";
import { LoadAllFilesToDb } from "!/App";
import { FileBrowser } from "./screens/FileBrowser";

const queryClient = new QueryClient();

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <button
        className="absolute top-0 left-0 bg-red-500"
        onClick={() => {
          window.location.reload();
        }}
      >
        RELOAD
      </button>
      <div id="App">
        <FileBrowser />
      </div>
    </QueryClientProvider>
  );
}

export default App;
