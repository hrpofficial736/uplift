import type { User } from "../../lib/userType";

export const ActivePlan = ({ info }: { info: User }) => {
  return (
    <div className="w-full">
      <h4 className="text-white/40 text-sm font-[600]">Active Plan</h4>
      <div className="mt-1 p-3 flex justify-between items-center w-full bg-zinc-800 rounded-xl">
        <div className="flex flex-col">
          <h1 className="font-[500] text-white/80 text-xs sm:text-sm">
            {info.plan === "FREE" ? "Free" : "Pro"} Plan
          </h1>
          <h4 className="font-[400] text-xs sm:text-sm text-white/60">
            Max {info.plan === "FREE" ? 3 : 10} prompts
          </h4>
        </div>

        <div className="sm:px-3 py-2 font-rubik text-white/70 rounded-md text-xs font-[500] cursor-pointer">
          {info.plan === "FREE" ? 3 - info.prompts : 10 - info.prompts} prompts
          remaining
        </div>
      </div>
    </div>
  );
};
