import { Route, Routes } from "react-router-dom";
import { useState } from "react";
import { useSupabaseAuth } from "./hooks/useSupabaseAuth";
import { AppLayout } from "./components/AppLayout";
import Prompt from "./components/Prompt";
import ReviewPage, { type ReviewPageProps } from "./components/ReviewPage";

const App = () => {
  const [info, setInfo] = useState<ReviewPageProps>({
    security: "",
    maintainability: "",
    quality: "",
    mentor: "",
    repoName: "",
    ownerName: "",
  });

  const { authenticated, session } = useSupabaseAuth();

  return (
    <Routes>
      <Route
        path="/"
        element={
          <AppLayout session={session!} authenticated={authenticated}>
            <Prompt callback={setInfo} />
          </AppLayout>
        }
      />
      <Route path="/review" element={<ReviewPage props={info} />} />
    </Routes>
  );
};

export default App;
