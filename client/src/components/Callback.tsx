import { useEffect } from "react";
import { FiLoader } from "react-icons/fi";
import { supabaseClient } from "../lib/supabaseClient";
import callAPI from "../helpers/fetchWrapper";
import toast from "react-hot-toast/headless";
import { useNavigate } from "react-router-dom";

export const Callback = () => {
  const navigate = useNavigate();
  useEffect(() => {
    const checkSession = async () => {
      const {
        data: { session },
      } = await supabaseClient.auth.getSession();
      if (!session) return;
      const responseFromServer = await callAPI({
        path: "/api/auth",
        method: "POST",
        headers: {
          Authorization: `Bearer ${session?.access_token}`,
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          email: session?.user.email,
          name: session.user.user_metadata["full_name"],
        }),
      });

      if (responseFromServer.status !== 200) {
        toast.error("Error signing you in...");
        return;
      } else navigate("/");
    };
    checkSession();
  }, [navigate]);
  return (
    <div className="flex justify-center items-center gap-2 w-screen h-screen bg-black text-white/80 font-rubik font-[600]">
      <FiLoader className="animate-spin" />
      Hang on
    </div>
  );
};
