export const ActivePlan = () => {
  return (
    <div className="w-full">
      <h4 className="text-white/40 text-sm font-[600]">Active Plan</h4>
      <div className="mt-1 p-3 flex justify-between items-center w-full bg-zinc-800 rounded-xl">
        <div className="flex flex-col">
          <h1 className="font-[500] text-white/80 text-sm">Free Plan</h1>
          <h4 className="font-[400] text-sm text-white/60">Max 3 prompts</h4>
        </div>

        <div className="px-3 py-2 font-rubik text-white/70 rounded-md text-xs font-[500] cursor-pointer">
          2 prompts remaining
        </div>
      </div>
    </div>
  );
};
