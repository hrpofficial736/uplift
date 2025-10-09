import { useState } from "react";
import { FaCheckSquare, FaSearch } from "react-icons/fa";
import { LuLoader } from "react-icons/lu";
import { FaAngleDown } from "react-icons/fa";
import { MdCheckBoxOutlineBlank } from "react-icons/md";
import { motion, AnimatePresence } from "motion/react";

type AvailableAgents = {
  security: boolean;
  maintainability: boolean;
  quality: boolean;
  mentor: boolean;
};

const Prompt = () => {
  const [loading, setLoading] = useState<boolean>(false);
  const [openDropDown, setOpenDropDown] = useState<boolean>(false);
  const [selectedAgents, setSelectedAgents] = useState<AvailableAgents>({
    security: true,
    maintainability: true,
    quality: true,
    mentor: true,
  });
  return (
    <div className="w-full absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 font-bold flex flex-col justify-center items-center">
      <h1 className="text-white/80 font-[500] text-4xl">
        Let's improve your project...
      </h1>

      <p className="text-white/50 w-[50%] mt-3 text-center font-[400] ">
        Your project deserves honest feedback — Uplift’s AI reviewers analyze
        your project’s security, maintainability, and quality, while a mentor
        turns their criticism into a growth plan you’ll actually enjoy
        following.
      </p>

      <div className="flex justify-center gap-3 mt-7 relative">
        <input
          type="url"
          className="pl-7 pt-3 pb-20 w-[1000px] font-[500] bg-zinc-900 focus:outline-none text-white/80  placeholder:text-white/50 border-2 border-white/20 rounded-xl"
          placeholder="Enter the prompt with the github url of your project..."
        />
        <div
          onClick={() => setOpenDropDown(!openDropDown)}
          className="font-[500] cursor-pointer bg-zinc-800 rounded-lg px-3 py-2 text-white/70 absolute bottom-3 left-5 flex gap-2 justify-center items-center"
        >
          <FaAngleDown />
          Select Agents
        </div>

        <AnimatePresence>
          {openDropDown && (
            <motion.div
              initial={{
                opacity: 0,
                y: -10,
              }}
              animate={{
                opacity: 1,
                y: 0,
              }}
              exit={{
                opacity: 0,
                y: -10,
              }}
              transition={{
                duration: 0.2,
                type: "tween",
              }}
              className="rounded-lg text-white/70 font-[500] absolute left-4 -bottom-32 px-3 py-2 bg-zinc-800 flex flex-col gap-2 justify-center items-start"
            >
              <div
                onClick={() =>
                  setSelectedAgents((prevData) => ({
                    ...prevData,
                    security: !selectedAgents.security,
                  }))
                }
                className="flex gap-3 items-center cursor-pointer"
              >
                {!selectedAgents.security ? (
                  <MdCheckBoxOutlineBlank className="cursor-pointer" />
                ) : (
                  <FaCheckSquare className="cursor-pointer" />
                )}{" "}
                Security
              </div>
              <div
                onClick={() =>
                  setSelectedAgents((prevData) => ({
                    ...prevData,
                    maintainability: !selectedAgents.maintainability,
                  }))
                }
                className="flex gap-3 items-center cursor-pointer"
              >
                {!selectedAgents.maintainability ? (
                  <MdCheckBoxOutlineBlank className="cursor-pointer" />
                ) : (
                  <FaCheckSquare className="cursor-pointer" />
                )}{" "}
                Maintainability
              </div>
              <div
                onClick={() =>
                  setSelectedAgents((prevData) => ({
                    ...prevData,
                    quality: !selectedAgents.quality,
                  }))
                }
                className="flex gap-3 items-center cursor-pointer"
              >
                {!selectedAgents.quality ? (
                  <MdCheckBoxOutlineBlank className="cursor-pointer" />
                ) : (
                  <FaCheckSquare className="cursor-pointer" />
                )}{" "}
                Quality
              </div>
              <div
                onClick={() =>
                  setSelectedAgents((prevData) => ({
                    ...prevData,
                    mentor: !selectedAgents.mentor,
                  }))
                }
                className="flex gap-3 items-center cursor-pointer"
              >
                {!selectedAgents.mentor ? (
                  <MdCheckBoxOutlineBlank className="cursor-pointer" />
                ) : (
                  <FaCheckSquare className="cursor-pointer" />
                )}{" "}
                Mentor
              </div>
            </motion.div>
          )}
        </AnimatePresence>

        <button
          onClick={() => setLoading(!loading)}
          className={`rounded-4xl cursor-pointer absolute top-3 right-3 transition-all duration-200 font-[600] flex justify-center items-center gap-3 ${
            loading
              ? "bg-white/5 text-white/60"
              : "bg-indigo-700 hover:bg-indigo-600 text-white/80"
          } px-7 py-3`}
        >
          {loading ? <LuLoader className="animate-spin" /> : <FaSearch />}
          {loading ? "Getting your repo..." : "Review"}
        </button>
      </div>
    </div>
  );
};

export default Prompt;
