import { useEffect, useState } from "react";
import { supabaseClient } from "../lib/supabaseClient";
import toast from "react-hot-toast";
import type { Session } from "@supabase/supabase-js";

export const useSupabaseAuth = () => {
  const [session, setSession] = useState<Session | null>(null);
  const [authenticated, setAuthenticated] = useState(false);

  useEffect(() => {
    let handled = false;

    const fetchSession = async () => {
      const { data, error } = await supabaseClient.auth.getSession();
      if (error) {
        toast.error("Please sign in first!");
        return;
      }
      if (data.session) {
        setAuthenticated(true);
        setSession(data.session);
      }
    };

    fetchSession();

    const {
      data: { subscription },
    } = supabaseClient.auth.onAuthStateChange((_event, session) => {
      if (!session) {
        toast.error("Error signing you in...");
        setAuthenticated(false);
        return;
      }
      history.replaceState(null, "", window.location.pathname);
      if (handled) return;
      handled = true;
      toast.success("Signed back successfully");
      setAuthenticated(true);
      setSession(session);
    });

    return () => subscription.unsubscribe();
  }, []);

  return { session, authenticated };
};
