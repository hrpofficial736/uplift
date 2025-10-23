import type { Session } from "@supabase/supabase-js";
import { useEffect, useState } from "react";

export const useGetUserInfo = (session: Session) => {
  const [info, setInfo] = useState<object>();
  useEffect(() => {
    const getUserInfo = async () => {
      const response = await fetch(
        `${import.meta.env.VITE_SERVER_URL}/api/get-user-info`,
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${session.access_token}`,
          },
          body: JSON.stringify({
            email: session.user.email,
          }),
        },
      );

      const data = await response.json();
      setInfo(data);
    };

    getUserInfo();
  }, []);

  return info;
};
