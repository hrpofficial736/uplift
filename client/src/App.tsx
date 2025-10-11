import { Route, Routes } from "react-router-dom";
import Logo from "./components/Logo";
import Prompt from "./components/Prompt";
import ReviewPage, { type ReviewPageProps } from "./components/ReviewPage";
import { useState } from "react";

const App = () => {
  const [info, setInfo] = useState<ReviewPageProps>({
    security: "",
    maintainability: "",
    quality: "",
    mentor: "",
    repoName: "",
    ownerName: "",
  });

  const setReviewInfo = (incomingInfo: ReviewPageProps) => {
    setInfo(incomingInfo);
  };
  return (
    <Routes>
      <Route
        path="/"
        element={
          <div className="p-5 h-screen w-screen relative bg-black text-white font-rubik overflow-hidden">
            <Logo />
            <Prompt callback={setReviewInfo} />
          </div>
        }
      />
      <Route path="/review" element={<ReviewPage props={info} />} />
    </Routes>
  );
};

export default App;
