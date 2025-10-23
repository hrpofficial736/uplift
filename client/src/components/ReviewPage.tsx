import { FaGithub } from "react-icons/fa6";
import Logo from "./Logo";
import { motion } from "motion/react";
import ReactMarkdown from "react-markdown";

export type ReviewPageProps = {
  security?: string;
  maintainability?: string;
  quality?: string;
  mentor?: string;
  repoName: string;
  ownerName: string;
};

export default function ReviewPage({ props }: { props: ReviewPageProps }) {
  return (
    <div className="w-screen flex flex-col px-10 py-7 bg-black text-white font-rubik">
      {/*Header*/}
      <Logo />
      <div className="font-[600] mt-10 text-xl md:text-3xl flex items-center text-white/90 gap-4">
        <FaGithub className="text-white" />
        {props.ownerName}/{props.repoName}
      </div>

      {/*Cards Container*/}
      <div className="mt-10 self-center flex flex-col justify-center items-center xl:w-1/2 gap-5 h-full">
        {props.security?.trim() !== "" && (
          <motion.div
            initial={{
              y: 50,
              opacity: 0,
            }}
            animate={{
              y: 0,
              opacity: 1,
            }}
            transition={{
              duration: 0.5,
              type: "tween",
            }}
            className="rounded-2xl relative flex flex-col gap-4 px-5 py-4 bg-gradient-to-r from-black via-black to-red-950"
          >
            <h1 className="text-2xl font-[600]">Security</h1>
            <div className="relative text-white rounded-xl leading-relaxed">
              <img
                src="/security-2.png"
                alt="Mentor"
                className="float-right ml-5 max-w-[30%] object-cover rounded-lg brightness-60"
              />

              <div className="text-white/80 text-justify whitespace-pre-line">
                <ReactMarkdown>{props.security}</ReactMarkdown>
              </div>
            </div>
          </motion.div>
        )}
        {props.maintainability?.trim() !== "" && (
          <motion.div
            initial={{
              y: 50,
              opacity: 0,
            }}
            animate={{
              y: 0,
              opacity: 1,
            }}
            transition={{
              duration: 0.5,
              type: "tween",
            }}
            className="rounded-2xl flex flex-col gap-4 px-5 py-4 bg-gradient-to-r from-black to-yellow-950 relative"
          >
            <h1 className="text-2xl font-[600]">Maintainability</h1>
            <div className="relative text-white rounded-xl leading-relaxed">
              <img
                src="/maintainability-2.png"
                alt="Mentor"
                className="float-right ml-5 max-w-[30%] object-cover rounded-lg brightness-60"
              />

              <div className="text-white/80 text-justify whitespace-pre-line">
                <ReactMarkdown>{props.maintainability}</ReactMarkdown>
              </div>
            </div>
          </motion.div>
        )}
        {props.quality?.trim() !== "" && (
          <motion.div
            initial={{
              y: 50,
              opacity: 0,
            }}
            animate={{
              y: 0,
              opacity: 1,
            }}
            transition={{
              duration: 0.5,
              type: "tween",
            }}
            className="rounded-2xl relative flex flex-col gap-4 px-5 py-4 bg-gradient-to-r from-black to-blue-950"
          >
            <h1 className="text-2xl font-[600]">Quality</h1>
            <div className="relative text-white rounded-xl leading-relaxed">
              <img
                src="/quality-2.png"
                alt="Mentor"
                className="float-right ml-5 max-w-[30%] object-cover rounded-lg brightness-60"
              />

              <div className="text-white/80 text-justify whitespace-pre-line">
                <ReactMarkdown>{props.quality}</ReactMarkdown>
              </div>
            </div>
          </motion.div>
        )}
        {props.mentor?.trim() !== "" && (
          <motion.div
            initial={{
              y: 50,
              opacity: 0,
            }}
            animate={{
              y: 0,
              opacity: 1,
            }}
            transition={{
              duration: 0.5,
              type: "tween",
            }}
            className="rounded-2xl relative flex flex-col gap-4 px-5 py-4 bg-gradient-to-r from-black to-green-950"
          >
            <h1 className="text-2xl font-[600]">Mentor</h1>
            <div className="relative text-white rounded-xl leading-relaxed">
              <img
                src="/mentor-3.png"
                alt="Mentor"
                className="float-right ml-5 max-w-[30%] object-cover rounded-lg brightness-60"
              />

              <div className="text-white/80 text-justify whitespace-pre-line">
                <ReactMarkdown>{props.mentor}</ReactMarkdown>
              </div>
            </div>
          </motion.div>
        )}
      </div>
    </div>
  );
}
