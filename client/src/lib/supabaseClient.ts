import { SupabaseClient } from "@supabase/supabase-js";

export const supabaseClient = new SupabaseClient(
  import.meta.env.VITE_SUPABASE_URL,
  import.meta.env.SUPABASE_KEY,
);
