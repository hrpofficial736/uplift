import { Header } from "./Header";
import { Settings } from "./Settings";
import { Toaster } from "react-hot-toast";
import { useState } from "react";
import type { Session } from "@supabase/supabase-js";

interface LayoutProps {
  children: React.ReactNode;
  session?: Session;
  authenticated: boolean;
}

export const AppLayout = ({
  children,
  session,
  authenticated,
}: LayoutProps) => {
  const [openSettings, setOpenSettings] = useState(false);

  const handleSettingsToggle = () => setOpenSettings((prev) => !prev);

  return (
    <div className="p-5 h-screen w-screen relative bg-black text-white font-rubik overflow-hidden">
      <Header callback={handleSettingsToggle} authenticated={authenticated} />
      <Toaster
        toastOptions={{
          style: {
            backgroundColor: "rgb(23, 22, 22)",
            color: "white",
          },
        }}
      />
      {session && (
        <Settings
          session={session!}
          show={openSettings}
          onClose={() => setOpenSettings(false)}
        />
      )}
      {children}
    </div>
  );
};
