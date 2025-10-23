import { useState } from "react";
import { FiLoader } from "react-icons/fi";
import { IoLogOut } from "react-icons/io5";
import { supabaseClient } from "../../lib/supabaseClient";
import type { Session } from "@supabase/supabase-js";

export const General = ({ session }: { session: Session }) => {
  const [loading, setLoading] = useState<boolean>(false);
  return (
    <div className="w-full">
      <h4 className="text-white/40 text-sm font-[600]">General</h4>
      <div className="mt-1 p-3 flex justify-between items-center w-full bg-zinc-800 rounded-xl">
        <div className="flex flex-col">
          <h1 className="font-[500] text-white/80 text-xs sm:text-sm">
            {session.user.user_metadata["full_name"]}
          </h1>
          <h4 className="font-[400] text-xs sm:text-sm text-white/60">
            {session.user.email}
          </h4>
        </div>

        <button
          onClick={async () => {
            setLoading(true);
            const { error } = await supabaseClient.auth.signOut();
            if (!error) {
              setLoading(false);
              window.location.reload();
              return;
            }
          }}
          className={`${loading ? "bg-zinc-900 text-white/70" : "bg-gradient-to-b from-red-600 to-red-800 text-white"} active:scale-[0.98] transition-all duration-300 p-2 sm:px-3 sm:py-2 flex justify-center items-center gap-2 font-rubik rounded-md text-xs font-[500] cursor-pointer`}
        >
          {loading ? <FiLoader className="animate-spin" /> : <IoLogOut />}
          {loading ? "Hang on..." : "Log Out"}
        </button>
      </div>
    </div>
  );
};
