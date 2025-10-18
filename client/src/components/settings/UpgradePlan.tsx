import { useState } from "react";
import { IoFlash } from "react-icons/io5";
import { FiLoader } from "react-icons/fi";

export const UpgradePlan = () => {
  const [loading, setLoading] = useState<boolean>(false);
  return (
    <div className="w-full">
      <h4 className="text-white/40 text-sm font-[600]">Upgrade</h4>
      <div className="mt-1 p-3 flex justify-between items-center w-full bg-zinc-800 rounded-xl">
        <div className="flex flex-col">
          <h1 className="font-[500] text-yellow-500 text-sm">Upgrade to Pro</h1>
          <h4 className="font-[400] text-sm text-white/60">Max 10 prompts</h4>
        </div>

        <button
          onClick={() => setLoading(!loading)}
          className={`${loading ? "bg-zinc-900 text-white/70" : "bg-gradient-to-b from-yellow-400 to-yellow-600 text-black/80"} active:scale-[0.98] transition-all duration-300 px-3 flex justify-center items-center gap-2 py-2 font-rubik rounded-md text-xs font-[500] cursor-pointer`}
        >
          {loading ? <FiLoader className="animate-spin" /> : <IoFlash />}
          {loading ? "Hang on..." : "Upgrade"}
        </button>
      </div>
    </div>
  );
};
