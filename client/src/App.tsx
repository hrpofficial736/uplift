import { Route, Routes } from "react-router-dom";
import Prompt from "./components/Prompt";
import ReviewPage, { type ReviewPageProps } from "./components/ReviewPage";
import { useEffect, useState } from "react";
import { Header } from "./components/Header";
import { supabaseClient } from "./lib/supabaseClient";
import toast, { Toaster } from "react-hot-toast";

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

  const [authenticated, setAuthenticated] = useState<boolean>(false);

  useEffect(() => {
    const fetchSession = async () => {
      const { data, error } = await supabaseClient.auth.getSession();
      if (error) {
        toast.error("Please sign in first!");
        return;
      }
      if (data.session) {
        setAuthenticated(true);
      }
    };
    fetchSession();
  }, []);
  return (
    <Routes>
      <Route
        path="/"
        element={
          <div className="p-5 h-screen w-screen relative bg-black text-white font-rubik overflow-hidden">
            <Header authenticated={authenticated} />
            <Prompt callback={setReviewInfo} />
          </div>
        }
      />
      <Route path="/review" element={<ReviewPage props={info} />} />
      <Toaster
        containerStyle={{
          backgroundColor: "black",
          color: "white",
        }}
      />
    </Routes>
  );
};

export default App;
