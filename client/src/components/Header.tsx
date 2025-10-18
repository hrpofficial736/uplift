import toast from "react-hot-toast/headless";
import { supabaseClient } from "../lib/supabaseClient";
import Logo from "./Logo";
import type { Session } from "@supabase/supabase-js";
import callAPI from "../helpers/fetchWrapper";

export const Header = ({
  authenticated,
  session,
  callback,
}: {
  authenticated: boolean;
  session?: Session;
  callback: () => void;
}) => {
  const handleGoogleAuth = async () => {
    const { error } = await supabaseClient.auth.signInWithOAuth({
      provider: "google",
      options: {
        redirectTo: "/",
        queryParams: {
          access_type: "offline",
          prompt: "consent",
        },
      },
    });

    if (error) {
      toast.error("Error signing you in");
      return;
    }

    const responseFromServer = await callAPI({
      path: "/api/auth",
      method: "POST",
      headers: {
        Authorization: `Bearer ${session?.access_token}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        email: session?.user.email,
      }),
    });

    if (responseFromServer.status !== 200) {
      toast.error("Error signing you in...");
      return;
    }
    toast.success("Signed in successfully");
  };

  return (
    <div className="flex justify-between items-start">
      <Logo />
      {authenticated ? (
        <div
          onClick={callback}
          className="flex gap-3 justify-center items-center hover:bg-zinc-900 transition-colors duration-200 cursor-pointer rounded-xl overflow-hidden px-3 py-2"
        >
          <img
            src={session?.user?.user_metadata["avatar_url"]}
            className="w-10 h-10 rounded-full border-2 border-zinc-500"
          />
          <h1 className="font-rubik font-[500] text-white/80">
            {session?.user.user_metadata["full_name"]}
          </h1>
        </div>
      ) : (
        <button
          onClick={handleGoogleAuth}
          className="px-3 py-2 font-rubik bg-gradient-to-b from-indigo-600 to-indigo-700 rounded-md text-sm font-[500] cursor-pointer"
        >
          Sign in with Google
        </button>
      )}
    </div>
  );
};
