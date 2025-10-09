import { Route, Routes } from "react-router-dom";
import Logo from "./components/Logo";
import Prompt from "./components/Prompt";
import ReviewPage from "./components/ReviewPage";

const App = () => {
  return (
    <Routes>
      <Route
        path="/"
        element={
          <div className="p-5 h-screen w-screen relative bg-black text-white font-rubik overflow-hidden">
            <Logo />
            <Prompt />
          </div>
        }
      />
      <Route path="/review" element={<ReviewPage />} />
    </Routes>
  );
};

export default App;
