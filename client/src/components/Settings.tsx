import { AnimatePresence, motion } from "motion/react";
import { IoCloseSharp } from "react-icons/io5";
import { General } from "./settings/General";
import { ActivePlan } from "./settings/ActivePlan";
import { UpgradePlan } from "./settings/UpgradePlan";
import { useGetUserInfo } from "../hooks/useGetUserInfo";
import type { Session } from "@supabase/supabase-js";
import type { User } from "../lib/userType";

export const Settings = ({
  show,
  onClose,
  session,
}: {
  show: boolean;
  onClose: () => void;
  session: Session;
}) => {
  const info = useGetUserInfo(session);
  return (
    <AnimatePresence>
      {show && (
        <>
          <motion.div
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
            transition={{ duration: 0.25 }}
            className="fixed inset-0 bg-black/40 backdrop-blur-sm z-[9998]"
            onClick={onClose}
          />
          <motion.div
            initial={{
              opacity: 0,
              scale: 0.7,
            }}
            animate={{
              opacity: 1,
              scale: 1,
            }}
            transition={{
              duration: 0.3,
              type: "tween",
            }}
            exit={{
              opacity: 0,
              scale: 0.7,
            }}
            className="bg-zinc-900 flex flex-col gap-5 absolute z-[9999] top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 rounded-lg border-2 border-white/5 px-4 py-3 w-[90%] sm:w-[600px] lg:w-[800px] h-[450px]"
          >
            {/*Header*/}
            <div className="flex justify-between items-center font-[500] text-white/70">
              Settings
              <IoCloseSharp
                className="text-white bg-red-500 cursor-pointer rounded-full size-6 p-1 flex justify-center items-center"
                onClick={onClose}
              />
            </div>

            {/*List*/}
            <div className="flex flex-col self-center justify-center items-center gap-5 w-full sm:w-[80%]">
              <General session={session} />
              <ActivePlan info={info as User} />
              <UpgradePlan info={info as User} />
            </div>
          </motion.div>
        </>
      )}
    </AnimatePresence>
  );
};
