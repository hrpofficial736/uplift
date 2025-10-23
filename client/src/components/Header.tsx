import toast from "react-hot-toast/headless";
import { supabaseClient } from "../lib/supabaseClient";
import Logo from "./Logo";
import { IoMdSettings } from "react-icons/io";

export const Header = ({
  authenticated,
  callback,
}: {
  authenticated: boolean;
  callback: () => void;
}) => {
  const handleGoogleAuth = async () => {
    const { error } = await supabaseClient.auth.signInWithOAuth({
      provider: "google",
      options: {
        redirectTo: `${window.location.origin}/auth/callback`,
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
  };

  return (
    <div className="flex justify-between items-start">
      <Logo />
      {authenticated ? (
        <div
          onClick={callback}
          className="flex gap-2 justify-center items-center hover:bg-zinc-900 transition-colors duration-200 cursor-pointer rounded-xl overflow-hidden p-3"
        >
          <IoMdSettings className="size-5 text-white/80" />
          <h1 className="font-rubik max-sm:text-xs font-[500] text-white/80">
            Manage your account
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
