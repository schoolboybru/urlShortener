import "./App.css";
import { QueryClient, QueryClientProvider } from "react-query";
import { UrlForm } from "./components/UrlForm";

const queryClient = new QueryClient();

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <UrlForm />
    </QueryClientProvider>
  );
}

export default App;
