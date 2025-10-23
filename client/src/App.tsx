import { Route, Routes } from "react-router-dom";
import { useState } from "react";
import { useSupabaseAuth } from "./hooks/useSupabaseAuth";
import { AppLayout } from "./components/AppLayout";
import Prompt from "./components/Prompt";
import ReviewPage, { type ReviewPageProps } from "./components/ReviewPage";
import { Callback } from "./components/Callback";

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
            <Prompt session={session!} callback={setInfo} />
          </AppLayout>
        }
      />
      <Route path="/review" element={<ReviewPage props={info} />} />
      <Route path="/auth/callback" element={<Callback />} />
    </Routes>
  );
};

export default App;
