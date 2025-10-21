import { useState } from "react";
import { FiLoader } from "react-icons/fi";
import { IoLogOut } from "react-icons/io5";
import { supabaseClient } from "../../lib/supabaseClient";

export const General = () => {
  const [loading, setLoading] = useState<boolean>(false);
  return (
    <div className="w-full">
      <h4 className="text-white/40 text-sm font-[600]">General</h4>
      <div className="mt-1 p-3 flex justify-between items-center w-full bg-zinc-800 rounded-xl">
        <div className="flex flex-col">
          <h1 className="font-[500] text-white/80 text-sm">
            Harshit Raj Pandey
          </h1>
          <h4 className="font-[400] text-sm text-white/60">dummy@gmail.com</h4>
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
            console.log(error);
          }}
          className={`${loading ? "bg-zinc-900 text-white/70" : "bg-gradient-to-b from-red-600 to-red-800 text-white"} active:scale-[0.98] transition-all duration-300 px-3 py-2 flex justify-center items-center gap-2 font-rubik rounded-md text-xs font-[500] cursor-pointer`}
        >
          {loading ? <FiLoader className="animate-spin" /> : <IoLogOut />}
          {loading ? "Hang on..." : "Log Out"}
        </button>
      </div>
    </div>
  );
};
