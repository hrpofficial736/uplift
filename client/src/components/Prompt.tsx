import { useState } from "react";
import { FaSearch } from "react-icons/fa";
import { LuLoader } from "react-icons/lu";
import { motion } from "motion/react";
import callAPI from "../helpers/fetchWrapper";
import { useNavigate } from "react-router-dom";
import type { ReviewPageProps } from "./ReviewPage";
import type { Session } from "@supabase/supabase-js";
import toast from "react-hot-toast";

const Prompt = ({
  callback,
  session,
}: {
  callback: (info: ReviewPageProps) => void;
  session: Session;
}) => {
  const [loading, setLoading] = useState<boolean>(false);
  const [prompt, setPrompt] = useState<string>("");

  const [showHeader, setShowHeader] = useState<boolean>(true);
  const [messages, setMessages] = useState<string[]>([]);

  const navigate = useNavigate();

  const submitPrompt = async () => {
    if (!session) {
      toast.error("Please sign in to get started");
      setLoading(false);
      return;
    }
    if (!prompt) {
      toast.error("Please write a prompt");
      setLoading(false);
      return;
    }
    setPrompt("");
    const data = await callAPI({
      method: "POST",
      path: "/api/github",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${session.access_token}`,
      },
      body: JSON.stringify({
        prompt: prompt,
        email: session.user.email,
      }),
    });

    if (data.status === 200 && !data.reviewed) {
      setShowHeader(false);
      setMessages([data.message]);
      setLoading(false);
      return;
    } else if (data.status === 200 && data.reviewed) {
      setLoading(false);
      let reviewInfo: ReviewPageProps = {
        security: "",
        maintainability: "",
        quality: "",
        mentor: "",
        repoName: data.repoInfo.repoName,
        ownerName: data.repoInfo.ownerName,
      };
      data.data.map((section) => {
        reviewInfo = {
          ...reviewInfo,
          [section.agent]: section.data.text,
        };
      });
      callback(reviewInfo);
      navigate("/review");
      return;
    } else {
      toast.error("Prompt limit reached");
      setLoading(false);
      return;
    }
  };
  return (
    <div className="w-full px-5 py-2 absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 font-bold flex flex-col justify-center items-center">
      {!showHeader ? (
        <div className="flex flex-col items-start w-[1000px] gap-3">
          {messages.map((message, index) => (
            <motion.div
              key={index}
              initial={{
                y: 50,
                opacity: 0,
              }}
              animate={{
                y: 0,
                opacity: 1,
              }}
              transition={{
                duration: 0.3,
                type: "tween",
              }}
              className={`rounded-xl px-5 py-3 bg-zinc-900 font-[500] max-w-[500px] self-start text-white/70`}
            >
              {message}
            </motion.div>
          ))}
        </div>
      ) : (
        <>
          <h1 className="bg-gradient-to-t from-indigo-700 to-indigo-500 bg-clip-text text-transparent font-[500] text-2xl sm:text-3xl md:text-4xl">
            Let's improve your project...
          </h1>

          <p className="text-white/50 lg:w-[50%] max-md:text-sm mt-3 text-center font-[400] ">
            Your project deserves honest feedback — Uplift’s AI reviewers
            analyze your project’s security, maintainability, and quality, while
            a mentor turns their criticism into a growth plan you’ll actually
            enjoy following.
          </p>
        </>
      )}

      <div className="flex justify-center w-[90%] lg:w-[80%] xl:w-[60%] gap-3 mt-7 relative">
        <input
          type="text"
          value={prompt}
          onChange={(e) => setPrompt(e.target.value)}
          placeholder="Your project URL"
          className="p-5 w-full font-[500] bg-zinc-900 focus:outline-none text-white/80  border-2 border-white/20 rounded-xl"
        />

        <button
          onClick={() => {
            setLoading(true);
            submitPrompt();
          }}
          className={`rounded-lg cursor-pointer absolute top-1/2 -translate-y-1/2 right-3 transition-all duration-200 font-[600] flex justify-center items-center gap-3 ${
            loading ? "bg-white/5 text-white/60" : "bg-indigo-700 text-white/80"
          } px-4 py-3`}
        >
          {loading ? <LuLoader className="animate-spin" /> : <FaSearch />}
          {loading ? "Hang on..." : ""}
        </button>
      </div>
    </div>
  );
};

export default Prompt;
