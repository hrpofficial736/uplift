import { useState } from "react";
import { FaSearch } from "react-icons/fa";
import { LuLoader } from "react-icons/lu";

const Prompt = () => {
  const [loading, setLoading] = useState<boolean>(false);
  return (
    <div className="w-full absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 font-bold flex flex-col justify-center items-center">
      <h1 className="text-white/80 font-[500] text-4xl">
        Let's improve your project...
      </h1>

      <p className="text-white/50 w-[50%] mt-3 text-center font-[400] ">
        Lorem ipsum, dolor sit amet consectetur adipisicing elit. Reprehenderit
        libero officia ad, quas molestiae harum dolor quam mollitia aut corrupti
        numquam neque esse in voluptatem autem fugiat id sint placeat!
      </p>

      <div className="flex w-full justify-center gap-3 mt-7">
        <input
        type="url"
          className="pl-7 py-3 w-[40%] font-[500] focus:outline-none text-white/80  placeholder:text-white/50 border-2 border-white/20 rounded-4xl"
          placeholder="Enter the github url of your project..."
        />
        <button
          onClick={() => setLoading(!loading)}
          className={`rounded-4xl cursor-pointer transition-all duration-200 font-[600] flex justify-center items-center gap-3 ${
            loading
              ? "bg-white/5 text-white/60"
              : "bg-indigo-700 hover:bg-indigo-600 text-white/80"
          } px-7 py-3`}
        >
          {loading ? <LuLoader className="animate-spin" /> : <FaSearch />}
          {loading ? "Getting your repo..." : "Review Project"}
        </button>
      </div>
    </div>
  );
};

export default Prompt;
